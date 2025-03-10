// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go
//
// Generated by this command:
//
//	mockgen -source=manager.go -destination=mock_manager.go -package=hook
//

// Package hook is a generated GoMock package.
package hook

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/flightctl/flightctl/api/v1alpha1"
	gomock "go.uber.org/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// OnAfterRebooting mocks base method.
func (m *MockManager) OnAfterRebooting(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnAfterRebooting", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnAfterRebooting indicates an expected call of OnAfterRebooting.
func (mr *MockManagerMockRecorder) OnAfterRebooting(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnAfterRebooting", reflect.TypeOf((*MockManager)(nil).OnAfterRebooting), ctx)
}

// OnAfterUpdating mocks base method.
func (m *MockManager) OnAfterUpdating(ctx context.Context, current, desired *v1alpha1.DeviceSpec, systemRebooted bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnAfterUpdating", ctx, current, desired, systemRebooted)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnAfterUpdating indicates an expected call of OnAfterUpdating.
func (mr *MockManagerMockRecorder) OnAfterUpdating(ctx, current, desired, systemRebooted any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnAfterUpdating", reflect.TypeOf((*MockManager)(nil).OnAfterUpdating), ctx, current, desired, systemRebooted)
}

// OnBeforeRebooting mocks base method.
func (m *MockManager) OnBeforeRebooting(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnBeforeRebooting", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnBeforeRebooting indicates an expected call of OnBeforeRebooting.
func (mr *MockManagerMockRecorder) OnBeforeRebooting(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnBeforeRebooting", reflect.TypeOf((*MockManager)(nil).OnBeforeRebooting), ctx)
}

// OnBeforeUpdating mocks base method.
func (m *MockManager) OnBeforeUpdating(ctx context.Context, current, desired *v1alpha1.DeviceSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnBeforeUpdating", ctx, current, desired)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnBeforeUpdating indicates an expected call of OnBeforeUpdating.
func (mr *MockManagerMockRecorder) OnBeforeUpdating(ctx, current, desired any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnBeforeUpdating", reflect.TypeOf((*MockManager)(nil).OnBeforeUpdating), ctx, current, desired)
}

// Sync mocks base method.
func (m *MockManager) Sync(current, desired *v1alpha1.DeviceSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync", current, desired)
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync.
func (mr *MockManagerMockRecorder) Sync(current, desired any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockManager)(nil).Sync), current, desired)
}
