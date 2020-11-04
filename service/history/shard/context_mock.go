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

// Code generated by MockGen. DO NOT EDIT.
// Source: context.go

// Package shard is a generated GoMock package.
package shard

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	v1 "go.temporal.io/api/common/v1"
	cache "go.temporal.io/server/common/cache"
	clock "go.temporal.io/server/common/clock"
	cluster "go.temporal.io/server/common/cluster"
	log "go.temporal.io/server/common/log"
	metrics "go.temporal.io/server/common/metrics"
	persistence "go.temporal.io/server/common/persistence"
	resource "go.temporal.io/server/common/resource"
	configs "go.temporal.io/server/service/history/configs"
	events "go.temporal.io/server/service/history/events"
)

// MockContext is a mock of Context interface.
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext.
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance.
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// GetShardID mocks base method.
func (m *MockContext) GetShardID() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShardID")
	ret0, _ := ret[0].(int32)
	return ret0
}

// GetShardID indicates an expected call of GetShardID.
func (mr *MockContextMockRecorder) GetShardID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShardID", reflect.TypeOf((*MockContext)(nil).GetShardID))
}

// GetService mocks base method.
func (m *MockContext) GetService() resource.Resource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService")
	ret0, _ := ret[0].(resource.Resource)
	return ret0
}

// GetService indicates an expected call of GetService.
func (mr *MockContextMockRecorder) GetService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockContext)(nil).GetService))
}

// GetExecutionManager mocks base method.
func (m *MockContext) GetExecutionManager() persistence.ExecutionManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExecutionManager")
	ret0, _ := ret[0].(persistence.ExecutionManager)
	return ret0
}

// GetExecutionManager indicates an expected call of GetExecutionManager.
func (mr *MockContextMockRecorder) GetExecutionManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExecutionManager", reflect.TypeOf((*MockContext)(nil).GetExecutionManager))
}

// GetHistoryManager mocks base method.
func (m *MockContext) GetHistoryManager() persistence.HistoryManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoryManager")
	ret0, _ := ret[0].(persistence.HistoryManager)
	return ret0
}

// GetHistoryManager indicates an expected call of GetHistoryManager.
func (mr *MockContextMockRecorder) GetHistoryManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistoryManager", reflect.TypeOf((*MockContext)(nil).GetHistoryManager))
}

// GetNamespaceCache mocks base method.
func (m *MockContext) GetNamespaceCache() cache.NamespaceCache {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespaceCache")
	ret0, _ := ret[0].(cache.NamespaceCache)
	return ret0
}

// GetNamespaceCache indicates an expected call of GetNamespaceCache.
func (mr *MockContextMockRecorder) GetNamespaceCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespaceCache", reflect.TypeOf((*MockContext)(nil).GetNamespaceCache))
}

// GetClusterMetadata mocks base method.
func (m *MockContext) GetClusterMetadata() cluster.Metadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterMetadata")
	ret0, _ := ret[0].(cluster.Metadata)
	return ret0
}

// GetClusterMetadata indicates an expected call of GetClusterMetadata.
func (mr *MockContextMockRecorder) GetClusterMetadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterMetadata", reflect.TypeOf((*MockContext)(nil).GetClusterMetadata))
}

// GetConfig mocks base method.
func (m *MockContext) GetConfig() *configs.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*configs.Config)
	return ret0
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockContextMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockContext)(nil).GetConfig))
}

// GetEventsCache mocks base method.
func (m *MockContext) GetEventsCache() events.Cache {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventsCache")
	ret0, _ := ret[0].(events.Cache)
	return ret0
}

// GetEventsCache indicates an expected call of GetEventsCache.
func (mr *MockContextMockRecorder) GetEventsCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventsCache", reflect.TypeOf((*MockContext)(nil).GetEventsCache))
}

// GetLogger mocks base method.
func (m *MockContext) GetLogger() log.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogger")
	ret0, _ := ret[0].(log.Logger)
	return ret0
}

// GetLogger indicates an expected call of GetLogger.
func (mr *MockContextMockRecorder) GetLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogger", reflect.TypeOf((*MockContext)(nil).GetLogger))
}

// GetThrottledLogger mocks base method.
func (m *MockContext) GetThrottledLogger() log.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetThrottledLogger")
	ret0, _ := ret[0].(log.Logger)
	return ret0
}

// GetThrottledLogger indicates an expected call of GetThrottledLogger.
func (mr *MockContextMockRecorder) GetThrottledLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThrottledLogger", reflect.TypeOf((*MockContext)(nil).GetThrottledLogger))
}

// GetMetricsClient mocks base method.
func (m *MockContext) GetMetricsClient() metrics.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetricsClient")
	ret0, _ := ret[0].(metrics.Client)
	return ret0
}

// GetMetricsClient indicates an expected call of GetMetricsClient.
func (mr *MockContextMockRecorder) GetMetricsClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetricsClient", reflect.TypeOf((*MockContext)(nil).GetMetricsClient))
}

// GetTimeSource mocks base method.
func (m *MockContext) GetTimeSource() clock.TimeSource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeSource")
	ret0, _ := ret[0].(clock.TimeSource)
	return ret0
}

// GetTimeSource indicates an expected call of GetTimeSource.
func (mr *MockContextMockRecorder) GetTimeSource() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeSource", reflect.TypeOf((*MockContext)(nil).GetTimeSource))
}

// PreviousShardOwnerWasDifferent mocks base method.
func (m *MockContext) PreviousShardOwnerWasDifferent() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreviousShardOwnerWasDifferent")
	ret0, _ := ret[0].(bool)
	return ret0
}

// PreviousShardOwnerWasDifferent indicates an expected call of PreviousShardOwnerWasDifferent.
func (mr *MockContextMockRecorder) PreviousShardOwnerWasDifferent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreviousShardOwnerWasDifferent", reflect.TypeOf((*MockContext)(nil).PreviousShardOwnerWasDifferent))
}

// GetEngine mocks base method.
func (m *MockContext) GetEngine() Engine {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEngine")
	ret0, _ := ret[0].(Engine)
	return ret0
}

// GetEngine indicates an expected call of GetEngine.
func (mr *MockContextMockRecorder) GetEngine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEngine", reflect.TypeOf((*MockContext)(nil).GetEngine))
}

// SetEngine mocks base method.
func (m *MockContext) SetEngine(arg0 Engine) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetEngine", arg0)
}

// SetEngine indicates an expected call of SetEngine.
func (mr *MockContextMockRecorder) SetEngine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEngine", reflect.TypeOf((*MockContext)(nil).SetEngine), arg0)
}

// GenerateTransferTaskID mocks base method.
func (m *MockContext) GenerateTransferTaskID() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateTransferTaskID")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateTransferTaskID indicates an expected call of GenerateTransferTaskID.
func (mr *MockContextMockRecorder) GenerateTransferTaskID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateTransferTaskID", reflect.TypeOf((*MockContext)(nil).GenerateTransferTaskID))
}

// GenerateTransferTaskIDs mocks base method.
func (m *MockContext) GenerateTransferTaskIDs(number int) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateTransferTaskIDs", number)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateTransferTaskIDs indicates an expected call of GenerateTransferTaskIDs.
func (mr *MockContextMockRecorder) GenerateTransferTaskIDs(number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateTransferTaskIDs", reflect.TypeOf((*MockContext)(nil).GenerateTransferTaskIDs), number)
}

// GetTransferMaxReadLevel mocks base method.
func (m *MockContext) GetTransferMaxReadLevel() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransferMaxReadLevel")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetTransferMaxReadLevel indicates an expected call of GetTransferMaxReadLevel.
func (mr *MockContextMockRecorder) GetTransferMaxReadLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransferMaxReadLevel", reflect.TypeOf((*MockContext)(nil).GetTransferMaxReadLevel))
}

// UpdateTimerMaxReadLevel mocks base method.
func (m *MockContext) UpdateTimerMaxReadLevel(cluster string) time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTimerMaxReadLevel", cluster)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// UpdateTimerMaxReadLevel indicates an expected call of UpdateTimerMaxReadLevel.
func (mr *MockContextMockRecorder) UpdateTimerMaxReadLevel(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTimerMaxReadLevel", reflect.TypeOf((*MockContext)(nil).UpdateTimerMaxReadLevel), cluster)
}

// SetCurrentTime mocks base method.
func (m *MockContext) SetCurrentTime(cluster string, currentTime time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCurrentTime", cluster, currentTime)
}

// SetCurrentTime indicates an expected call of SetCurrentTime.
func (mr *MockContextMockRecorder) SetCurrentTime(cluster, currentTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrentTime", reflect.TypeOf((*MockContext)(nil).SetCurrentTime), cluster, currentTime)
}

// GetCurrentTime mocks base method.
func (m *MockContext) GetCurrentTime(cluster string) time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentTime", cluster)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetCurrentTime indicates an expected call of GetCurrentTime.
func (mr *MockContextMockRecorder) GetCurrentTime(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentTime", reflect.TypeOf((*MockContext)(nil).GetCurrentTime), cluster)
}

// GetLastUpdatedTime mocks base method.
func (m *MockContext) GetLastUpdatedTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastUpdatedTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetLastUpdatedTime indicates an expected call of GetLastUpdatedTime.
func (mr *MockContextMockRecorder) GetLastUpdatedTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastUpdatedTime", reflect.TypeOf((*MockContext)(nil).GetLastUpdatedTime))
}

// GetTimerMaxReadLevel mocks base method.
func (m *MockContext) GetTimerMaxReadLevel(cluster string) time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimerMaxReadLevel", cluster)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTimerMaxReadLevel indicates an expected call of GetTimerMaxReadLevel.
func (mr *MockContextMockRecorder) GetTimerMaxReadLevel(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimerMaxReadLevel", reflect.TypeOf((*MockContext)(nil).GetTimerMaxReadLevel), cluster)
}

// GetTransferAckLevel mocks base method.
func (m *MockContext) GetTransferAckLevel() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransferAckLevel")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetTransferAckLevel indicates an expected call of GetTransferAckLevel.
func (mr *MockContextMockRecorder) GetTransferAckLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransferAckLevel", reflect.TypeOf((*MockContext)(nil).GetTransferAckLevel))
}

// UpdateTransferAckLevel mocks base method.
func (m *MockContext) UpdateTransferAckLevel(ackLevel int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransferAckLevel", ackLevel)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTransferAckLevel indicates an expected call of UpdateTransferAckLevel.
func (mr *MockContextMockRecorder) UpdateTransferAckLevel(ackLevel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransferAckLevel", reflect.TypeOf((*MockContext)(nil).UpdateTransferAckLevel), ackLevel)
}

// GetTransferClusterAckLevel mocks base method.
func (m *MockContext) GetTransferClusterAckLevel(cluster string) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransferClusterAckLevel", cluster)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetTransferClusterAckLevel indicates an expected call of GetTransferClusterAckLevel.
func (mr *MockContextMockRecorder) GetTransferClusterAckLevel(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransferClusterAckLevel", reflect.TypeOf((*MockContext)(nil).GetTransferClusterAckLevel), cluster)
}

// UpdateTransferClusterAckLevel mocks base method.
func (m *MockContext) UpdateTransferClusterAckLevel(cluster string, ackLevel int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransferClusterAckLevel", cluster, ackLevel)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTransferClusterAckLevel indicates an expected call of UpdateTransferClusterAckLevel.
func (mr *MockContextMockRecorder) UpdateTransferClusterAckLevel(cluster, ackLevel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransferClusterAckLevel", reflect.TypeOf((*MockContext)(nil).UpdateTransferClusterAckLevel), cluster, ackLevel)
}

// GetReplicatorAckLevel mocks base method.
func (m *MockContext) GetReplicatorAckLevel() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReplicatorAckLevel")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetReplicatorAckLevel indicates an expected call of GetReplicatorAckLevel.
func (mr *MockContextMockRecorder) GetReplicatorAckLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplicatorAckLevel", reflect.TypeOf((*MockContext)(nil).GetReplicatorAckLevel))
}

// UpdateReplicatorAckLevel mocks base method.
func (m *MockContext) UpdateReplicatorAckLevel(ackLevel int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReplicatorAckLevel", ackLevel)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateReplicatorAckLevel indicates an expected call of UpdateReplicatorAckLevel.
func (mr *MockContextMockRecorder) UpdateReplicatorAckLevel(ackLevel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReplicatorAckLevel", reflect.TypeOf((*MockContext)(nil).UpdateReplicatorAckLevel), ackLevel)
}

// GetReplicatorDLQAckLevel mocks base method.
func (m *MockContext) GetReplicatorDLQAckLevel(sourceCluster string) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReplicatorDLQAckLevel", sourceCluster)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetReplicatorDLQAckLevel indicates an expected call of GetReplicatorDLQAckLevel.
func (mr *MockContextMockRecorder) GetReplicatorDLQAckLevel(sourceCluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplicatorDLQAckLevel", reflect.TypeOf((*MockContext)(nil).GetReplicatorDLQAckLevel), sourceCluster)
}

// UpdateReplicatorDLQAckLevel mocks base method.
func (m *MockContext) UpdateReplicatorDLQAckLevel(sourCluster string, ackLevel int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReplicatorDLQAckLevel", sourCluster, ackLevel)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateReplicatorDLQAckLevel indicates an expected call of UpdateReplicatorDLQAckLevel.
func (mr *MockContextMockRecorder) UpdateReplicatorDLQAckLevel(sourCluster, ackLevel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReplicatorDLQAckLevel", reflect.TypeOf((*MockContext)(nil).UpdateReplicatorDLQAckLevel), sourCluster, ackLevel)
}

// GetClusterReplicationLevel mocks base method.
func (m *MockContext) GetClusterReplicationLevel(cluster string) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterReplicationLevel", cluster)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetClusterReplicationLevel indicates an expected call of GetClusterReplicationLevel.
func (mr *MockContextMockRecorder) GetClusterReplicationLevel(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterReplicationLevel", reflect.TypeOf((*MockContext)(nil).GetClusterReplicationLevel), cluster)
}

// UpdateClusterReplicationLevel mocks base method.
func (m *MockContext) UpdateClusterReplicationLevel(cluster string, lastTaskID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClusterReplicationLevel", cluster, lastTaskID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClusterReplicationLevel indicates an expected call of UpdateClusterReplicationLevel.
func (mr *MockContextMockRecorder) UpdateClusterReplicationLevel(cluster, lastTaskID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClusterReplicationLevel", reflect.TypeOf((*MockContext)(nil).UpdateClusterReplicationLevel), cluster, lastTaskID)
}

// GetTimerAckLevel mocks base method.
func (m *MockContext) GetTimerAckLevel() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimerAckLevel")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTimerAckLevel indicates an expected call of GetTimerAckLevel.
func (mr *MockContextMockRecorder) GetTimerAckLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimerAckLevel", reflect.TypeOf((*MockContext)(nil).GetTimerAckLevel))
}

// UpdateTimerAckLevel mocks base method.
func (m *MockContext) UpdateTimerAckLevel(ackLevel time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTimerAckLevel", ackLevel)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTimerAckLevel indicates an expected call of UpdateTimerAckLevel.
func (mr *MockContextMockRecorder) UpdateTimerAckLevel(ackLevel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTimerAckLevel", reflect.TypeOf((*MockContext)(nil).UpdateTimerAckLevel), ackLevel)
}

// GetTimerClusterAckLevel mocks base method.
func (m *MockContext) GetTimerClusterAckLevel(cluster string) time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimerClusterAckLevel", cluster)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTimerClusterAckLevel indicates an expected call of GetTimerClusterAckLevel.
func (mr *MockContextMockRecorder) GetTimerClusterAckLevel(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimerClusterAckLevel", reflect.TypeOf((*MockContext)(nil).GetTimerClusterAckLevel), cluster)
}

// UpdateTimerClusterAckLevel mocks base method.
func (m *MockContext) UpdateTimerClusterAckLevel(cluster string, ackLevel time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTimerClusterAckLevel", cluster, ackLevel)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTimerClusterAckLevel indicates an expected call of UpdateTimerClusterAckLevel.
func (mr *MockContextMockRecorder) UpdateTimerClusterAckLevel(cluster, ackLevel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTimerClusterAckLevel", reflect.TypeOf((*MockContext)(nil).UpdateTimerClusterAckLevel), cluster, ackLevel)
}

// UpdateTransferFailoverLevel mocks base method.
func (m *MockContext) UpdateTransferFailoverLevel(failoverID string, level persistence.TransferFailoverLevel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransferFailoverLevel", failoverID, level)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTransferFailoverLevel indicates an expected call of UpdateTransferFailoverLevel.
func (mr *MockContextMockRecorder) UpdateTransferFailoverLevel(failoverID, level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransferFailoverLevel", reflect.TypeOf((*MockContext)(nil).UpdateTransferFailoverLevel), failoverID, level)
}

// DeleteTransferFailoverLevel mocks base method.
func (m *MockContext) DeleteTransferFailoverLevel(failoverID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTransferFailoverLevel", failoverID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTransferFailoverLevel indicates an expected call of DeleteTransferFailoverLevel.
func (mr *MockContextMockRecorder) DeleteTransferFailoverLevel(failoverID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransferFailoverLevel", reflect.TypeOf((*MockContext)(nil).DeleteTransferFailoverLevel), failoverID)
}

// GetAllTransferFailoverLevels mocks base method.
func (m *MockContext) GetAllTransferFailoverLevels() map[string]persistence.TransferFailoverLevel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTransferFailoverLevels")
	ret0, _ := ret[0].(map[string]persistence.TransferFailoverLevel)
	return ret0
}

// GetAllTransferFailoverLevels indicates an expected call of GetAllTransferFailoverLevels.
func (mr *MockContextMockRecorder) GetAllTransferFailoverLevels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTransferFailoverLevels", reflect.TypeOf((*MockContext)(nil).GetAllTransferFailoverLevels))
}

// UpdateTimerFailoverLevel mocks base method.
func (m *MockContext) UpdateTimerFailoverLevel(failoverID string, level persistence.TimerFailoverLevel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTimerFailoverLevel", failoverID, level)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTimerFailoverLevel indicates an expected call of UpdateTimerFailoverLevel.
func (mr *MockContextMockRecorder) UpdateTimerFailoverLevel(failoverID, level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTimerFailoverLevel", reflect.TypeOf((*MockContext)(nil).UpdateTimerFailoverLevel), failoverID, level)
}

// DeleteTimerFailoverLevel mocks base method.
func (m *MockContext) DeleteTimerFailoverLevel(failoverID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTimerFailoverLevel", failoverID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTimerFailoverLevel indicates an expected call of DeleteTimerFailoverLevel.
func (mr *MockContextMockRecorder) DeleteTimerFailoverLevel(failoverID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTimerFailoverLevel", reflect.TypeOf((*MockContext)(nil).DeleteTimerFailoverLevel), failoverID)
}

// GetAllTimerFailoverLevels mocks base method.
func (m *MockContext) GetAllTimerFailoverLevels() map[string]persistence.TimerFailoverLevel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTimerFailoverLevels")
	ret0, _ := ret[0].(map[string]persistence.TimerFailoverLevel)
	return ret0
}

// GetAllTimerFailoverLevels indicates an expected call of GetAllTimerFailoverLevels.
func (mr *MockContextMockRecorder) GetAllTimerFailoverLevels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTimerFailoverLevels", reflect.TypeOf((*MockContext)(nil).GetAllTimerFailoverLevels))
}

// GetNamespaceNotificationVersion mocks base method.
func (m *MockContext) GetNamespaceNotificationVersion() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespaceNotificationVersion")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetNamespaceNotificationVersion indicates an expected call of GetNamespaceNotificationVersion.
func (mr *MockContextMockRecorder) GetNamespaceNotificationVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespaceNotificationVersion", reflect.TypeOf((*MockContext)(nil).GetNamespaceNotificationVersion))
}

// UpdateNamespaceNotificationVersion mocks base method.
func (m *MockContext) UpdateNamespaceNotificationVersion(namespaceNotificationVersion int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNamespaceNotificationVersion", namespaceNotificationVersion)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNamespaceNotificationVersion indicates an expected call of UpdateNamespaceNotificationVersion.
func (mr *MockContextMockRecorder) UpdateNamespaceNotificationVersion(namespaceNotificationVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNamespaceNotificationVersion", reflect.TypeOf((*MockContext)(nil).UpdateNamespaceNotificationVersion), namespaceNotificationVersion)
}

// CreateWorkflowExecution mocks base method.
func (m *MockContext) CreateWorkflowExecution(request *persistence.CreateWorkflowExecutionRequest) (*persistence.CreateWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWorkflowExecution", request)
	ret0, _ := ret[0].(*persistence.CreateWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWorkflowExecution indicates an expected call of CreateWorkflowExecution.
func (mr *MockContextMockRecorder) CreateWorkflowExecution(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorkflowExecution", reflect.TypeOf((*MockContext)(nil).CreateWorkflowExecution), request)
}

// UpdateWorkflowExecution mocks base method.
func (m *MockContext) UpdateWorkflowExecution(request *persistence.UpdateWorkflowExecutionRequest) (*persistence.UpdateWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWorkflowExecution", request)
	ret0, _ := ret[0].(*persistence.UpdateWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWorkflowExecution indicates an expected call of UpdateWorkflowExecution.
func (mr *MockContextMockRecorder) UpdateWorkflowExecution(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkflowExecution", reflect.TypeOf((*MockContext)(nil).UpdateWorkflowExecution), request)
}

// ConflictResolveWorkflowExecution mocks base method.
func (m *MockContext) ConflictResolveWorkflowExecution(request *persistence.ConflictResolveWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConflictResolveWorkflowExecution", request)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConflictResolveWorkflowExecution indicates an expected call of ConflictResolveWorkflowExecution.
func (mr *MockContextMockRecorder) ConflictResolveWorkflowExecution(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConflictResolveWorkflowExecution", reflect.TypeOf((*MockContext)(nil).ConflictResolveWorkflowExecution), request)
}

// ResetWorkflowExecution mocks base method.
func (m *MockContext) ResetWorkflowExecution(request *persistence.ResetWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetWorkflowExecution", request)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetWorkflowExecution indicates an expected call of ResetWorkflowExecution.
func (mr *MockContextMockRecorder) ResetWorkflowExecution(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetWorkflowExecution", reflect.TypeOf((*MockContext)(nil).ResetWorkflowExecution), request)
}

// AppendHistoryV2Events mocks base method.
func (m *MockContext) AppendHistoryV2Events(request *persistence.AppendHistoryNodesRequest, namespaceID string, execution v1.WorkflowExecution) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendHistoryV2Events", request, namespaceID, execution)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppendHistoryV2Events indicates an expected call of AppendHistoryV2Events.
func (mr *MockContextMockRecorder) AppendHistoryV2Events(request, namespaceID, execution interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendHistoryV2Events", reflect.TypeOf((*MockContext)(nil).AppendHistoryV2Events), request, namespaceID, execution)
}
