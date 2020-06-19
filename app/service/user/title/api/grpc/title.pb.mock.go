// Code generated by MockGen. DO NOT EDIT.
// Source: app/service/user/title/api/grpc/title.pb.go

// Package grpc is a generated GoMock package.
package grpc

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockTitleClient is a mock of TitleClient interface
type MockTitleClient struct {
	ctrl     *gomock.Controller
	recorder *MockTitleClientMockRecorder
}

// MockTitleClientMockRecorder is the mock recorder for MockTitleClient
type MockTitleClientMockRecorder struct {
	mock *MockTitleClient
}

// NewMockTitleClient creates a new mock instance
func NewMockTitleClient(ctrl *gomock.Controller) *MockTitleClient {
	mock := &MockTitleClient{ctrl: ctrl}
	mock.recorder = &MockTitleClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTitleClient) EXPECT() *MockTitleClientMockRecorder {
	return m.recorder
}

// GetTitle mocks base method
func (m *MockTitleClient) GetTitle(ctx context.Context, in *GetTitleReq, opts ...grpc.CallOption) (*GetTitleResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTitle", varargs...)
	ret0, _ := ret[0].(*GetTitleResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTitle indicates an expected call of GetTitle
func (mr *MockTitleClientMockRecorder) GetTitle(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTitle", reflect.TypeOf((*MockTitleClient)(nil).GetTitle), varargs...)
}

// SetTitle mocks base method
func (m *MockTitleClient) SetTitle(ctx context.Context, in *SetTitleReq, opts ...grpc.CallOption) (*SetTitleResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetTitle", varargs...)
	ret0, _ := ret[0].(*SetTitleResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetTitle indicates an expected call of SetTitle
func (mr *MockTitleClientMockRecorder) SetTitle(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTitle", reflect.TypeOf((*MockTitleClient)(nil).SetTitle), varargs...)
}

// MockTitleServer is a mock of TitleServer interface
type MockTitleServer struct {
	ctrl     *gomock.Controller
	recorder *MockTitleServerMockRecorder
}

// MockTitleServerMockRecorder is the mock recorder for MockTitleServer
type MockTitleServerMockRecorder struct {
	mock *MockTitleServer
}

// NewMockTitleServer creates a new mock instance
func NewMockTitleServer(ctrl *gomock.Controller) *MockTitleServer {
	mock := &MockTitleServer{ctrl: ctrl}
	mock.recorder = &MockTitleServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTitleServer) EXPECT() *MockTitleServerMockRecorder {
	return m.recorder
}

// GetTitle mocks base method
func (m *MockTitleServer) GetTitle(arg0 context.Context, arg1 *GetTitleReq) (*GetTitleResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTitle", arg0, arg1)
	ret0, _ := ret[0].(*GetTitleResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTitle indicates an expected call of GetTitle
func (mr *MockTitleServerMockRecorder) GetTitle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTitle", reflect.TypeOf((*MockTitleServer)(nil).GetTitle), arg0, arg1)
}

// SetTitle mocks base method
func (m *MockTitleServer) SetTitle(arg0 context.Context, arg1 *SetTitleReq) (*SetTitleResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTitle", arg0, arg1)
	ret0, _ := ret[0].(*SetTitleResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetTitle indicates an expected call of SetTitle
func (mr *MockTitleServerMockRecorder) SetTitle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTitle", reflect.TypeOf((*MockTitleServer)(nil).SetTitle), arg0, arg1)
}