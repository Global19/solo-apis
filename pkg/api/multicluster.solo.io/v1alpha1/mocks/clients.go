// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock_v1alpha1 is a generated GoMock package.
package mock_v1alpha1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/solo-apis/pkg/api/multicluster.solo.io/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockMulticlusterClientset is a mock of MulticlusterClientset interface
type MockMulticlusterClientset struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterClientsetMockRecorder
}

// MockMulticlusterClientsetMockRecorder is the mock recorder for MockMulticlusterClientset
type MockMulticlusterClientsetMockRecorder struct {
	mock *MockMulticlusterClientset
}

// NewMockMulticlusterClientset creates a new mock instance
func NewMockMulticlusterClientset(ctrl *gomock.Controller) *MockMulticlusterClientset {
	mock := &MockMulticlusterClientset{ctrl: ctrl}
	mock.recorder = &MockMulticlusterClientsetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterClientset) EXPECT() *MockMulticlusterClientsetMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockMulticlusterClientset) Cluster(cluster string) (v1alpha1.Clientset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1alpha1.Clientset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster
func (mr *MockMulticlusterClientsetMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterClientset)(nil).Cluster), cluster)
}

// MockClientset is a mock of Clientset interface
type MockClientset struct {
	ctrl     *gomock.Controller
	recorder *MockClientsetMockRecorder
}

// MockClientsetMockRecorder is the mock recorder for MockClientset
type MockClientsetMockRecorder struct {
	mock *MockClientset
}

// NewMockClientset creates a new mock instance
func NewMockClientset(ctrl *gomock.Controller) *MockClientset {
	mock := &MockClientset{ctrl: ctrl}
	mock.recorder = &MockClientsetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientset) EXPECT() *MockClientsetMockRecorder {
	return m.recorder
}

// MultiClusterRoles mocks base method
func (m *MockClientset) MultiClusterRoles() v1alpha1.MultiClusterRoleClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiClusterRoles")
	ret0, _ := ret[0].(v1alpha1.MultiClusterRoleClient)
	return ret0
}

// MultiClusterRoles indicates an expected call of MultiClusterRoles
func (mr *MockClientsetMockRecorder) MultiClusterRoles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiClusterRoles", reflect.TypeOf((*MockClientset)(nil).MultiClusterRoles))
}

// MockMultiClusterRoleReader is a mock of MultiClusterRoleReader interface
type MockMultiClusterRoleReader struct {
	ctrl     *gomock.Controller
	recorder *MockMultiClusterRoleReaderMockRecorder
}

// MockMultiClusterRoleReaderMockRecorder is the mock recorder for MockMultiClusterRoleReader
type MockMultiClusterRoleReaderMockRecorder struct {
	mock *MockMultiClusterRoleReader
}

// NewMockMultiClusterRoleReader creates a new mock instance
func NewMockMultiClusterRoleReader(ctrl *gomock.Controller) *MockMultiClusterRoleReader {
	mock := &MockMultiClusterRoleReader{ctrl: ctrl}
	mock.recorder = &MockMultiClusterRoleReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMultiClusterRoleReader) EXPECT() *MockMultiClusterRoleReaderMockRecorder {
	return m.recorder
}

// GetMultiClusterRole mocks base method
func (m *MockMultiClusterRoleReader) GetMultiClusterRole(ctx context.Context, key client.ObjectKey) (*v1alpha1.MultiClusterRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMultiClusterRole", ctx, key)
	ret0, _ := ret[0].(*v1alpha1.MultiClusterRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMultiClusterRole indicates an expected call of GetMultiClusterRole
func (mr *MockMultiClusterRoleReaderMockRecorder) GetMultiClusterRole(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleReader)(nil).GetMultiClusterRole), ctx, key)
}

// ListMultiClusterRole mocks base method
func (m *MockMultiClusterRoleReader) ListMultiClusterRole(ctx context.Context, opts ...client.ListOption) (*v1alpha1.MultiClusterRoleList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMultiClusterRole", varargs...)
	ret0, _ := ret[0].(*v1alpha1.MultiClusterRoleList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMultiClusterRole indicates an expected call of ListMultiClusterRole
func (mr *MockMultiClusterRoleReaderMockRecorder) ListMultiClusterRole(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleReader)(nil).ListMultiClusterRole), varargs...)
}

// MockMultiClusterRoleWriter is a mock of MultiClusterRoleWriter interface
type MockMultiClusterRoleWriter struct {
	ctrl     *gomock.Controller
	recorder *MockMultiClusterRoleWriterMockRecorder
}

// MockMultiClusterRoleWriterMockRecorder is the mock recorder for MockMultiClusterRoleWriter
type MockMultiClusterRoleWriterMockRecorder struct {
	mock *MockMultiClusterRoleWriter
}

// NewMockMultiClusterRoleWriter creates a new mock instance
func NewMockMultiClusterRoleWriter(ctrl *gomock.Controller) *MockMultiClusterRoleWriter {
	mock := &MockMultiClusterRoleWriter{ctrl: ctrl}
	mock.recorder = &MockMultiClusterRoleWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMultiClusterRoleWriter) EXPECT() *MockMultiClusterRoleWriterMockRecorder {
	return m.recorder
}

// CreateMultiClusterRole mocks base method
func (m *MockMultiClusterRoleWriter) CreateMultiClusterRole(ctx context.Context, obj *v1alpha1.MultiClusterRole, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMultiClusterRole", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMultiClusterRole indicates an expected call of CreateMultiClusterRole
func (mr *MockMultiClusterRoleWriterMockRecorder) CreateMultiClusterRole(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleWriter)(nil).CreateMultiClusterRole), varargs...)
}

// DeleteMultiClusterRole mocks base method
func (m *MockMultiClusterRoleWriter) DeleteMultiClusterRole(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMultiClusterRole", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMultiClusterRole indicates an expected call of DeleteMultiClusterRole
func (mr *MockMultiClusterRoleWriterMockRecorder) DeleteMultiClusterRole(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleWriter)(nil).DeleteMultiClusterRole), varargs...)
}

// UpdateMultiClusterRole mocks base method
func (m *MockMultiClusterRoleWriter) UpdateMultiClusterRole(ctx context.Context, obj *v1alpha1.MultiClusterRole, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateMultiClusterRole", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMultiClusterRole indicates an expected call of UpdateMultiClusterRole
func (mr *MockMultiClusterRoleWriterMockRecorder) UpdateMultiClusterRole(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleWriter)(nil).UpdateMultiClusterRole), varargs...)
}

// PatchMultiClusterRole mocks base method
func (m *MockMultiClusterRoleWriter) PatchMultiClusterRole(ctx context.Context, obj *v1alpha1.MultiClusterRole, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchMultiClusterRole", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchMultiClusterRole indicates an expected call of PatchMultiClusterRole
func (mr *MockMultiClusterRoleWriterMockRecorder) PatchMultiClusterRole(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleWriter)(nil).PatchMultiClusterRole), varargs...)
}

// DeleteAllOfMultiClusterRole mocks base method
func (m *MockMultiClusterRoleWriter) DeleteAllOfMultiClusterRole(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfMultiClusterRole", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfMultiClusterRole indicates an expected call of DeleteAllOfMultiClusterRole
func (mr *MockMultiClusterRoleWriterMockRecorder) DeleteAllOfMultiClusterRole(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleWriter)(nil).DeleteAllOfMultiClusterRole), varargs...)
}

// UpsertMultiClusterRole mocks base method
func (m *MockMultiClusterRoleWriter) UpsertMultiClusterRole(ctx context.Context, obj *v1alpha1.MultiClusterRole, transitionFuncs ...v1alpha1.MultiClusterRoleTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertMultiClusterRole", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertMultiClusterRole indicates an expected call of UpsertMultiClusterRole
func (mr *MockMultiClusterRoleWriterMockRecorder) UpsertMultiClusterRole(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleWriter)(nil).UpsertMultiClusterRole), varargs...)
}

// MockMultiClusterRoleStatusWriter is a mock of MultiClusterRoleStatusWriter interface
type MockMultiClusterRoleStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockMultiClusterRoleStatusWriterMockRecorder
}

// MockMultiClusterRoleStatusWriterMockRecorder is the mock recorder for MockMultiClusterRoleStatusWriter
type MockMultiClusterRoleStatusWriterMockRecorder struct {
	mock *MockMultiClusterRoleStatusWriter
}

// NewMockMultiClusterRoleStatusWriter creates a new mock instance
func NewMockMultiClusterRoleStatusWriter(ctrl *gomock.Controller) *MockMultiClusterRoleStatusWriter {
	mock := &MockMultiClusterRoleStatusWriter{ctrl: ctrl}
	mock.recorder = &MockMultiClusterRoleStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMultiClusterRoleStatusWriter) EXPECT() *MockMultiClusterRoleStatusWriterMockRecorder {
	return m.recorder
}

// UpdateMultiClusterRoleStatus mocks base method
func (m *MockMultiClusterRoleStatusWriter) UpdateMultiClusterRoleStatus(ctx context.Context, obj *v1alpha1.MultiClusterRole, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateMultiClusterRoleStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMultiClusterRoleStatus indicates an expected call of UpdateMultiClusterRoleStatus
func (mr *MockMultiClusterRoleStatusWriterMockRecorder) UpdateMultiClusterRoleStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMultiClusterRoleStatus", reflect.TypeOf((*MockMultiClusterRoleStatusWriter)(nil).UpdateMultiClusterRoleStatus), varargs...)
}

// PatchMultiClusterRoleStatus mocks base method
func (m *MockMultiClusterRoleStatusWriter) PatchMultiClusterRoleStatus(ctx context.Context, obj *v1alpha1.MultiClusterRole, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchMultiClusterRoleStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchMultiClusterRoleStatus indicates an expected call of PatchMultiClusterRoleStatus
func (mr *MockMultiClusterRoleStatusWriterMockRecorder) PatchMultiClusterRoleStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchMultiClusterRoleStatus", reflect.TypeOf((*MockMultiClusterRoleStatusWriter)(nil).PatchMultiClusterRoleStatus), varargs...)
}

// MockMultiClusterRoleClient is a mock of MultiClusterRoleClient interface
type MockMultiClusterRoleClient struct {
	ctrl     *gomock.Controller
	recorder *MockMultiClusterRoleClientMockRecorder
}

// MockMultiClusterRoleClientMockRecorder is the mock recorder for MockMultiClusterRoleClient
type MockMultiClusterRoleClientMockRecorder struct {
	mock *MockMultiClusterRoleClient
}

// NewMockMultiClusterRoleClient creates a new mock instance
func NewMockMultiClusterRoleClient(ctrl *gomock.Controller) *MockMultiClusterRoleClient {
	mock := &MockMultiClusterRoleClient{ctrl: ctrl}
	mock.recorder = &MockMultiClusterRoleClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMultiClusterRoleClient) EXPECT() *MockMultiClusterRoleClientMockRecorder {
	return m.recorder
}

// GetMultiClusterRole mocks base method
func (m *MockMultiClusterRoleClient) GetMultiClusterRole(ctx context.Context, key client.ObjectKey) (*v1alpha1.MultiClusterRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMultiClusterRole", ctx, key)
	ret0, _ := ret[0].(*v1alpha1.MultiClusterRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMultiClusterRole indicates an expected call of GetMultiClusterRole
func (mr *MockMultiClusterRoleClientMockRecorder) GetMultiClusterRole(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleClient)(nil).GetMultiClusterRole), ctx, key)
}

// ListMultiClusterRole mocks base method
func (m *MockMultiClusterRoleClient) ListMultiClusterRole(ctx context.Context, opts ...client.ListOption) (*v1alpha1.MultiClusterRoleList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMultiClusterRole", varargs...)
	ret0, _ := ret[0].(*v1alpha1.MultiClusterRoleList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMultiClusterRole indicates an expected call of ListMultiClusterRole
func (mr *MockMultiClusterRoleClientMockRecorder) ListMultiClusterRole(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleClient)(nil).ListMultiClusterRole), varargs...)
}

// CreateMultiClusterRole mocks base method
func (m *MockMultiClusterRoleClient) CreateMultiClusterRole(ctx context.Context, obj *v1alpha1.MultiClusterRole, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMultiClusterRole", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMultiClusterRole indicates an expected call of CreateMultiClusterRole
func (mr *MockMultiClusterRoleClientMockRecorder) CreateMultiClusterRole(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleClient)(nil).CreateMultiClusterRole), varargs...)
}

// DeleteMultiClusterRole mocks base method
func (m *MockMultiClusterRoleClient) DeleteMultiClusterRole(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMultiClusterRole", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMultiClusterRole indicates an expected call of DeleteMultiClusterRole
func (mr *MockMultiClusterRoleClientMockRecorder) DeleteMultiClusterRole(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleClient)(nil).DeleteMultiClusterRole), varargs...)
}

// UpdateMultiClusterRole mocks base method
func (m *MockMultiClusterRoleClient) UpdateMultiClusterRole(ctx context.Context, obj *v1alpha1.MultiClusterRole, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateMultiClusterRole", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMultiClusterRole indicates an expected call of UpdateMultiClusterRole
func (mr *MockMultiClusterRoleClientMockRecorder) UpdateMultiClusterRole(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleClient)(nil).UpdateMultiClusterRole), varargs...)
}

// PatchMultiClusterRole mocks base method
func (m *MockMultiClusterRoleClient) PatchMultiClusterRole(ctx context.Context, obj *v1alpha1.MultiClusterRole, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchMultiClusterRole", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchMultiClusterRole indicates an expected call of PatchMultiClusterRole
func (mr *MockMultiClusterRoleClientMockRecorder) PatchMultiClusterRole(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleClient)(nil).PatchMultiClusterRole), varargs...)
}

// DeleteAllOfMultiClusterRole mocks base method
func (m *MockMultiClusterRoleClient) DeleteAllOfMultiClusterRole(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfMultiClusterRole", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfMultiClusterRole indicates an expected call of DeleteAllOfMultiClusterRole
func (mr *MockMultiClusterRoleClientMockRecorder) DeleteAllOfMultiClusterRole(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleClient)(nil).DeleteAllOfMultiClusterRole), varargs...)
}

// UpsertMultiClusterRole mocks base method
func (m *MockMultiClusterRoleClient) UpsertMultiClusterRole(ctx context.Context, obj *v1alpha1.MultiClusterRole, transitionFuncs ...v1alpha1.MultiClusterRoleTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertMultiClusterRole", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertMultiClusterRole indicates an expected call of UpsertMultiClusterRole
func (mr *MockMultiClusterRoleClientMockRecorder) UpsertMultiClusterRole(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertMultiClusterRole", reflect.TypeOf((*MockMultiClusterRoleClient)(nil).UpsertMultiClusterRole), varargs...)
}

// UpdateMultiClusterRoleStatus mocks base method
func (m *MockMultiClusterRoleClient) UpdateMultiClusterRoleStatus(ctx context.Context, obj *v1alpha1.MultiClusterRole, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateMultiClusterRoleStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMultiClusterRoleStatus indicates an expected call of UpdateMultiClusterRoleStatus
func (mr *MockMultiClusterRoleClientMockRecorder) UpdateMultiClusterRoleStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMultiClusterRoleStatus", reflect.TypeOf((*MockMultiClusterRoleClient)(nil).UpdateMultiClusterRoleStatus), varargs...)
}

// PatchMultiClusterRoleStatus mocks base method
func (m *MockMultiClusterRoleClient) PatchMultiClusterRoleStatus(ctx context.Context, obj *v1alpha1.MultiClusterRole, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchMultiClusterRoleStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchMultiClusterRoleStatus indicates an expected call of PatchMultiClusterRoleStatus
func (mr *MockMultiClusterRoleClientMockRecorder) PatchMultiClusterRoleStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchMultiClusterRoleStatus", reflect.TypeOf((*MockMultiClusterRoleClient)(nil).PatchMultiClusterRoleStatus), varargs...)
}

// MockMulticlusterMultiClusterRoleClient is a mock of MulticlusterMultiClusterRoleClient interface
type MockMulticlusterMultiClusterRoleClient struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterMultiClusterRoleClientMockRecorder
}

// MockMulticlusterMultiClusterRoleClientMockRecorder is the mock recorder for MockMulticlusterMultiClusterRoleClient
type MockMulticlusterMultiClusterRoleClientMockRecorder struct {
	mock *MockMulticlusterMultiClusterRoleClient
}

// NewMockMulticlusterMultiClusterRoleClient creates a new mock instance
func NewMockMulticlusterMultiClusterRoleClient(ctrl *gomock.Controller) *MockMulticlusterMultiClusterRoleClient {
	mock := &MockMulticlusterMultiClusterRoleClient{ctrl: ctrl}
	mock.recorder = &MockMulticlusterMultiClusterRoleClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterMultiClusterRoleClient) EXPECT() *MockMulticlusterMultiClusterRoleClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockMulticlusterMultiClusterRoleClient) Cluster(cluster string) (v1alpha1.MultiClusterRoleClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1alpha1.MultiClusterRoleClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster
func (mr *MockMulticlusterMultiClusterRoleClientMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterMultiClusterRoleClient)(nil).Cluster), cluster)
}
