// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/zanzibar/examples/example-gateway/build/clients/corge-http (interfaces: Client)

// Package clientmock is a generated GoMock package.
package clientmock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	corge "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients-idl/clients/corge/corge"
	zanzibar "github.com/uber/zanzibar/runtime"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CorgeNoContentOnException mocks base method
func (m *MockClient) CorgeNoContentOnException(arg0 context.Context, arg1 map[string]string, arg2 *corge.Corge_NoContentOnException_Args) (*corge.Foo, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CorgeNoContentOnException", arg0, arg1, arg2)
	ret0, _ := ret[0].(*corge.Foo)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CorgeNoContentOnException indicates an expected call of CorgeNoContentOnException
func (mr *MockClientMockRecorder) CorgeNoContentOnException(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CorgeNoContentOnException", reflect.TypeOf((*MockClient)(nil).CorgeNoContentOnException), arg0, arg1, arg2)
}

// EchoString mocks base method
func (m *MockClient) EchoString(arg0 context.Context, arg1 map[string]string, arg2 *corge.Corge_EchoString_Args) (string, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EchoString", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EchoString indicates an expected call of EchoString
func (mr *MockClientMockRecorder) EchoString(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoString", reflect.TypeOf((*MockClient)(nil).EchoString), arg0, arg1, arg2)
}

// HTTPClient mocks base method
func (m *MockClient) HTTPClient() *zanzibar.HTTPClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPClient")
	ret0, _ := ret[0].(*zanzibar.HTTPClient)
	return ret0
}

// HTTPClient indicates an expected call of HTTPClient
func (mr *MockClientMockRecorder) HTTPClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPClient", reflect.TypeOf((*MockClient)(nil).HTTPClient))
}

// NoContent mocks base method
func (m *MockClient) NoContent(arg0 context.Context, arg1 map[string]string, arg2 *corge.Corge_NoContent_Args) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NoContent", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NoContent indicates an expected call of NoContent
func (mr *MockClientMockRecorder) NoContent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoContent", reflect.TypeOf((*MockClient)(nil).NoContent), arg0, arg1, arg2)
}

// NoContentNoException mocks base method
func (m *MockClient) NoContentNoException(arg0 context.Context, arg1 map[string]string, arg2 *corge.Corge_NoContentNoException_Args) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NoContentNoException", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NoContentNoException indicates an expected call of NoContentNoException
func (mr *MockClientMockRecorder) NoContentNoException(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoContentNoException", reflect.TypeOf((*MockClient)(nil).NoContentNoException), arg0, arg1, arg2)
}
