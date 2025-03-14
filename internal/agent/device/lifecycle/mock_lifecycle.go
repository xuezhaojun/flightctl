// Code generated by MockGen. DO NOT EDIT.
// Source: lifecycle.go
//
// Generated by this command:
//
//	mockgen -source=lifecycle.go -destination=mock_lifecycle.go -package=lifecycle
//

// Package lifecycle is a generated GoMock package.
package lifecycle

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

// AfterUpdate mocks base method.
func (m *MockManager) AfterUpdate(ctx context.Context, current, desired *v1alpha1.DeviceSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterUpdate", ctx, current, desired)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterUpdate indicates an expected call of AfterUpdate.
func (mr *MockManagerMockRecorder) AfterUpdate(ctx, current, desired any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterUpdate", reflect.TypeOf((*MockManager)(nil).AfterUpdate), ctx, current, desired)
}

// Sync mocks base method.
func (m *MockManager) Sync(ctx context.Context, current, desired *v1alpha1.DeviceSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync", ctx, current, desired)
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync.
func (mr *MockManagerMockRecorder) Sync(ctx, current, desired any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockManager)(nil).Sync), ctx, current, desired)
}

// MockInitializer is a mock of Initializer interface.
type MockInitializer struct {
	ctrl     *gomock.Controller
	recorder *MockInitializerMockRecorder
}

// MockInitializerMockRecorder is the mock recorder for MockInitializer.
type MockInitializerMockRecorder struct {
	mock *MockInitializer
}

// NewMockInitializer creates a new mock instance.
func NewMockInitializer(ctrl *gomock.Controller) *MockInitializer {
	mock := &MockInitializer{ctrl: ctrl}
	mock.recorder = &MockInitializerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInitializer) EXPECT() *MockInitializerMockRecorder {
	return m.recorder
}

// Initialize mocks base method.
func (m *MockInitializer) Initialize(ctx context.Context, status *v1alpha1.DeviceStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", ctx, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockInitializerMockRecorder) Initialize(ctx, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockInitializer)(nil).Initialize), ctx, status)
}

// IsInitialized mocks base method.
func (m *MockInitializer) IsInitialized() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInitialized")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsInitialized indicates an expected call of IsInitialized.
func (mr *MockInitializerMockRecorder) IsInitialized() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInitialized", reflect.TypeOf((*MockInitializer)(nil).IsInitialized))
}
