// Code generated by MockGen. DO NOT EDIT.
// Source: i_point.go

// Package mock is a generated GoMock package.
package mock

import (
	goneMock "github.com/gone-io/gone"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIPoint is a mock of IPoint interface.
type MockIPoint struct {
	goneMock.Flag
	ctrl     *gomock.Controller
	recorder *MockIPointMockRecorder
}

// MockIPointMockRecorder is the mock recorder for MockIPoint.
type MockIPointMockRecorder struct {
	mock *MockIPoint
}

// NewMockIPoint creates a new mock instance.
func NewMockIPoint(ctrl *gomock.Controller) *MockIPoint {
	mock := &MockIPoint{ctrl: ctrl}
	mock.recorder = &MockIPointMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPoint) EXPECT() *MockIPointMockRecorder {
	return m.recorder
}

// GetX mocks base method.
func (m *MockIPoint) GetX() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetX")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetX indicates an expected call of GetX.
func (mr *MockIPointMockRecorder) GetX() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetX", reflect.TypeOf((*MockIPoint)(nil).GetX))
}

// GetY mocks base method.
func (m *MockIPoint) GetY() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetY")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetY indicates an expected call of GetY.
func (mr *MockIPointMockRecorder) GetY() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetY", reflect.TypeOf((*MockIPoint)(nil).GetY))
}
