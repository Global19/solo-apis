// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1"
	controller "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockAuthConfigReconciler is a mock of AuthConfigReconciler interface
type MockAuthConfigReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockAuthConfigReconcilerMockRecorder
}

// MockAuthConfigReconcilerMockRecorder is the mock recorder for MockAuthConfigReconciler
type MockAuthConfigReconcilerMockRecorder struct {
	mock *MockAuthConfigReconciler
}

// NewMockAuthConfigReconciler creates a new mock instance
func NewMockAuthConfigReconciler(ctrl *gomock.Controller) *MockAuthConfigReconciler {
	mock := &MockAuthConfigReconciler{ctrl: ctrl}
	mock.recorder = &MockAuthConfigReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthConfigReconciler) EXPECT() *MockAuthConfigReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAuthConfig mocks base method
func (m *MockAuthConfigReconciler) ReconcileAuthConfig(obj *v1.AuthConfig) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAuthConfig", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileAuthConfig indicates an expected call of ReconcileAuthConfig
func (mr *MockAuthConfigReconcilerMockRecorder) ReconcileAuthConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAuthConfig", reflect.TypeOf((*MockAuthConfigReconciler)(nil).ReconcileAuthConfig), obj)
}

// MockAuthConfigDeletionReconciler is a mock of AuthConfigDeletionReconciler interface
type MockAuthConfigDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockAuthConfigDeletionReconcilerMockRecorder
}

// MockAuthConfigDeletionReconcilerMockRecorder is the mock recorder for MockAuthConfigDeletionReconciler
type MockAuthConfigDeletionReconcilerMockRecorder struct {
	mock *MockAuthConfigDeletionReconciler
}

// NewMockAuthConfigDeletionReconciler creates a new mock instance
func NewMockAuthConfigDeletionReconciler(ctrl *gomock.Controller) *MockAuthConfigDeletionReconciler {
	mock := &MockAuthConfigDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockAuthConfigDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthConfigDeletionReconciler) EXPECT() *MockAuthConfigDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAuthConfigDeletion mocks base method
func (m *MockAuthConfigDeletionReconciler) ReconcileAuthConfigDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAuthConfigDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileAuthConfigDeletion indicates an expected call of ReconcileAuthConfigDeletion
func (mr *MockAuthConfigDeletionReconcilerMockRecorder) ReconcileAuthConfigDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAuthConfigDeletion", reflect.TypeOf((*MockAuthConfigDeletionReconciler)(nil).ReconcileAuthConfigDeletion), req)
}

// MockAuthConfigFinalizer is a mock of AuthConfigFinalizer interface
type MockAuthConfigFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthConfigFinalizerMockRecorder
}

// MockAuthConfigFinalizerMockRecorder is the mock recorder for MockAuthConfigFinalizer
type MockAuthConfigFinalizerMockRecorder struct {
	mock *MockAuthConfigFinalizer
}

// NewMockAuthConfigFinalizer creates a new mock instance
func NewMockAuthConfigFinalizer(ctrl *gomock.Controller) *MockAuthConfigFinalizer {
	mock := &MockAuthConfigFinalizer{ctrl: ctrl}
	mock.recorder = &MockAuthConfigFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthConfigFinalizer) EXPECT() *MockAuthConfigFinalizerMockRecorder {
	return m.recorder
}

// ReconcileAuthConfig mocks base method
func (m *MockAuthConfigFinalizer) ReconcileAuthConfig(obj *v1.AuthConfig) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAuthConfig", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileAuthConfig indicates an expected call of ReconcileAuthConfig
func (mr *MockAuthConfigFinalizerMockRecorder) ReconcileAuthConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAuthConfig", reflect.TypeOf((*MockAuthConfigFinalizer)(nil).ReconcileAuthConfig), obj)
}

// AuthConfigFinalizerName mocks base method
func (m *MockAuthConfigFinalizer) AuthConfigFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthConfigFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// AuthConfigFinalizerName indicates an expected call of AuthConfigFinalizerName
func (mr *MockAuthConfigFinalizerMockRecorder) AuthConfigFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthConfigFinalizerName", reflect.TypeOf((*MockAuthConfigFinalizer)(nil).AuthConfigFinalizerName))
}

// FinalizeAuthConfig mocks base method
func (m *MockAuthConfigFinalizer) FinalizeAuthConfig(obj *v1.AuthConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeAuthConfig", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeAuthConfig indicates an expected call of FinalizeAuthConfig
func (mr *MockAuthConfigFinalizerMockRecorder) FinalizeAuthConfig(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeAuthConfig", reflect.TypeOf((*MockAuthConfigFinalizer)(nil).FinalizeAuthConfig), obj)
}

// MockAuthConfigReconcileLoop is a mock of AuthConfigReconcileLoop interface
type MockAuthConfigReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockAuthConfigReconcileLoopMockRecorder
}

// MockAuthConfigReconcileLoopMockRecorder is the mock recorder for MockAuthConfigReconcileLoop
type MockAuthConfigReconcileLoopMockRecorder struct {
	mock *MockAuthConfigReconcileLoop
}

// NewMockAuthConfigReconcileLoop creates a new mock instance
func NewMockAuthConfigReconcileLoop(ctrl *gomock.Controller) *MockAuthConfigReconcileLoop {
	mock := &MockAuthConfigReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockAuthConfigReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthConfigReconcileLoop) EXPECT() *MockAuthConfigReconcileLoopMockRecorder {
	return m.recorder
}

// RunAuthConfigReconciler mocks base method
func (m *MockAuthConfigReconcileLoop) RunAuthConfigReconciler(ctx context.Context, rec controller.AuthConfigReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunAuthConfigReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunAuthConfigReconciler indicates an expected call of RunAuthConfigReconciler
func (mr *MockAuthConfigReconcileLoopMockRecorder) RunAuthConfigReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunAuthConfigReconciler", reflect.TypeOf((*MockAuthConfigReconcileLoop)(nil).RunAuthConfigReconciler), varargs...)
}
