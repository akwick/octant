// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/heptio/developer-dash/internal/portforward (interfaces: PortForwardInterface)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	portforward "github.com/heptio/developer-dash/internal/portforward"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	reflect "reflect"
)

// MockPortForwardInterface is a mock of PortForwardInterface interface
type MockPortForwardInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPortForwardInterfaceMockRecorder
}

// MockPortForwardInterfaceMockRecorder is the mock recorder for MockPortForwardInterface
type MockPortForwardInterfaceMockRecorder struct {
	mock *MockPortForwardInterface
}

// NewMockPortForwardInterface creates a new mock instance
func NewMockPortForwardInterface(ctrl *gomock.Controller) *MockPortForwardInterface {
	mock := &MockPortForwardInterface{ctrl: ctrl}
	mock.recorder = &MockPortForwardInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPortForwardInterface) EXPECT() *MockPortForwardInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockPortForwardInterface) Create(arg0 context.Context, arg1 schema.GroupVersionKind, arg2, arg3 string, arg4 uint16) (portforward.PortForwardCreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(portforward.PortForwardCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockPortForwardInterfaceMockRecorder) Create(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPortForwardInterface)(nil).Create), arg0, arg1, arg2, arg3, arg4)
}

// Find mocks base method
func (m *MockPortForwardInterface) Find(arg0 string, arg1 schema.GroupVersionKind, arg2 string) (portforward.PortForwardState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1, arg2)
	ret0, _ := ret[0].(portforward.PortForwardState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockPortForwardInterfaceMockRecorder) Find(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockPortForwardInterface)(nil).Find), arg0, arg1, arg2)
}

// Get mocks base method
func (m *MockPortForwardInterface) Get(arg0 string) (portforward.PortForwardState, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(portforward.PortForwardState)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockPortForwardInterfaceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPortForwardInterface)(nil).Get), arg0)
}

// List mocks base method
func (m *MockPortForwardInterface) List() []portforward.PortForwardState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]portforward.PortForwardState)
	return ret0
}

// List indicates an expected call of List
func (mr *MockPortForwardInterfaceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPortForwardInterface)(nil).List))
}

// Stop mocks base method
func (m *MockPortForwardInterface) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockPortForwardInterfaceMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockPortForwardInterface)(nil).Stop))
}

// StopForwarder mocks base method
func (m *MockPortForwardInterface) StopForwarder(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopForwarder", arg0)
}

// StopForwarder indicates an expected call of StopForwarder
func (mr *MockPortForwardInterfaceMockRecorder) StopForwarder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopForwarder", reflect.TypeOf((*MockPortForwardInterface)(nil).StopForwarder), arg0)
}
