// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package recordactivitytaskstarted

import (
	"context"
	"fmt"

	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/api/serviceerror"
	"go.temporal.io/server/api/historyservice/v1"
	"go.temporal.io/server/common"
	"go.temporal.io/server/common/definition"
	"go.temporal.io/server/common/metrics"
	"go.temporal.io/server/common/namespace"
	serviceerrors "go.temporal.io/server/common/serviceerror"
	"go.temporal.io/server/common/tqid"
	"go.temporal.io/server/common/worker_versioning"
	"go.temporal.io/server/service/history/api"
	"go.temporal.io/server/service/history/consts"
	"go.temporal.io/server/service/history/shard"
	"go.temporal.io/server/service/history/workflow"
)

func Invoke(
	ctx context.Context,
	request *historyservice.RecordActivityTaskStartedRequest,
	shardContext shard.Context,
	workflowConsistencyChecker api.WorkflowConsistencyChecker,
) (resp *historyservice.RecordActivityTaskStartedResponse, retError error) {

	var err error
	response := &historyservice.RecordActivityTaskStartedResponse{}

	err = api.GetAndUpdateWorkflowWithNew(
		ctx,
		request.Clock,
		definition.NewWorkflowKey(
			request.NamespaceId,
			request.WorkflowExecution.WorkflowId,
			request.WorkflowExecution.RunId,
		),
		func(workflowLease api.WorkflowLease) (*api.UpdateWorkflowAction, error) {
			mutableState := workflowLease.GetMutableState()
			if !mutableState.IsWorkflowExecutionRunning() {
				return nil, consts.ErrWorkflowCompleted
			}

			response, err = recordActivityTaskStarted(ctx, shardContext, mutableState, request)
			if err != nil {
				return nil, err
			}

			return &api.UpdateWorkflowAction{
				Noop:               false,
				CreateWorkflowTask: false,
			}, nil
		},
		nil,
		shardContext,
		workflowConsistencyChecker,
	)

	if err != nil {
		return nil, err
	}

	return response, err
}

func recordActivityTaskStarted(
	ctx context.Context,
	shardContext shard.Context,
	mutableState workflow.MutableState,
	request *historyservice.RecordActivityTaskStartedRequest,
) (*historyservice.RecordActivityTaskStartedResponse, error) {
	namespaceEntry, err := api.GetActiveNamespace(shardContext, namespace.ID(request.GetNamespaceId()))
	if err != nil {
		return nil, err
	}
	namespaceName := namespaceEntry.Name().String()

	scheduledEventID := request.GetScheduledEventId()
	requestID := request.GetRequestId()
	ai, isRunning := mutableState.GetActivityInfo(scheduledEventID)

	taggedMetrics := shardContext.GetMetricsHandler().WithTags(metrics.OperationTag(metrics.HistoryRecordActivityTaskStartedScope))

	// First check to see if cache needs to be refreshed as we could potentially have stale workflow execution in
	// some extreme cassandra failure cases.
	if !isRunning && scheduledEventID >= mutableState.GetNextEventID() {
		metrics.StaleMutableStateCounter.With(taggedMetrics).Record(1)
		return nil, consts.ErrStaleState
	}

	// Check execution state to make sure task is in the list of outstanding tasks and it is not yet started.  If
	// task is not outstanding than it is most probably a duplicate and complete the task.
	if !isRunning {
		// Looks like ActivityTask already completed as a result of another call.
		// It is OK to drop the task at this point.
		return nil, consts.ErrActivityTaskNotFound
	}

	scheduledEvent, err := mutableState.GetActivityScheduledEvent(ctx, scheduledEventID)
	if err != nil {
		return nil, err
	}

	response := &historyservice.RecordActivityTaskStartedResponse{
		ScheduledEvent:              scheduledEvent,
		CurrentAttemptScheduledTime: ai.ScheduledTime,
	}

	if ai.StartedEventId != common.EmptyEventID {
		// If activity is started as part of the current request scope then return a positive response
		if ai.RequestId == requestID {
			response.StartedTime = ai.StartedTime
			response.Attempt = ai.Attempt
			return response, nil
		}

		// Looks like ActivityTask already started as a result of another call.
		// It is OK to drop the task at this point.
		return nil, serviceerrors.NewTaskAlreadyStarted("Activity")
	}

	if ai.Stamp != request.Stamp {
		// activity has changes before task is started.
		// ErrActivityStampMismatch is the error to indicate that requested activity has mismatched stamp
		errorMessage := fmt.Sprintf(
			"Activity task with this stamp not found. Id: %s,: type: %s, current stamp: %d",
			ai.ActivityId, ai.ActivityType.Name, ai.Stamp)
		return nil, serviceerror.NewNotFound(errorMessage)
	}

	versioningStamp := worker_versioning.StampFromCapabilities(request.PollRequest.WorkerVersionCapabilities)
	if _, err := mutableState.AddActivityTaskStartedEvent(
		ai, scheduledEventID, requestID, request.PollRequest.GetIdentity(),
		versioningStamp, request.GetBuildIdRedirectInfo(),
	); err != nil {
		return nil, err
	}

	scheduleToStartLatency := ai.GetStartedTime().AsTime().Sub(ai.GetScheduledTime().AsTime())
	metrics.TaskScheduleToStartLatency.With(
		metrics.GetPerTaskQueuePartitionTypeScope(
			taggedMetrics,
			namespaceName,
			// passing the root partition all the time as we don't care about partition ID in this metric
			tqid.UnsafeTaskQueueFamily(namespaceEntry.ID().String(), ai.GetTaskQueue()).TaskQueue(enumspb.TASK_QUEUE_TYPE_ACTIVITY).RootPartition(),
			shardContext.GetConfig().BreakdownMetricsByTaskQueue(namespaceName, ai.GetTaskQueue(), enumspb.TASK_QUEUE_TYPE_ACTIVITY),
		),
	).Record(scheduleToStartLatency)

	response.StartedTime = ai.StartedTime
	response.Attempt = ai.Attempt
	response.HeartbeatDetails = ai.LastHeartbeatDetails
	response.Version = ai.Version

	response.WorkflowType = mutableState.GetWorkflowType()
	response.WorkflowNamespace = namespaceName

	return response, nil
}
