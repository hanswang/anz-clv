// Code generated by MockGen. DO NOT EDIT.
// Source: aggregate.go

// Package mock_aggregator is a generated GoMock package.
package mock_aggregator

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/hanswang/clv/internal/types"
)

// MockAggregatorManager is a mock of AggregatorManager interface.
type MockAggregatorManager struct {
	ctrl     *gomock.Controller
	recorder *MockAggregatorManagerMockRecorder
}

// MockAggregatorManagerMockRecorder is the mock recorder for MockAggregatorManager.
type MockAggregatorManagerMockRecorder struct {
	mock *MockAggregatorManager
}

// NewMockAggregatorManager creates a new mock instance.
func NewMockAggregatorManager(ctrl *gomock.Controller) *MockAggregatorManager {
	mock := &MockAggregatorManager{ctrl: ctrl}
	mock.recorder = &MockAggregatorManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAggregatorManager) EXPECT() *MockAggregatorManagerMockRecorder {
	return m.recorder
}

// GenerateReport mocks base method.
func (m *MockAggregatorManager) GenerateReport(entities *map[string]types.Entity) []*types.Report {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateReport", entities)
	ret0, _ := ret[0].([]*types.Report)
	return ret0
}

// GenerateReport indicates an expected call of GenerateReport.
func (mr *MockAggregatorManagerMockRecorder) GenerateReport(entities interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateReport", reflect.TypeOf((*MockAggregatorManager)(nil).GenerateReport), entities)
}