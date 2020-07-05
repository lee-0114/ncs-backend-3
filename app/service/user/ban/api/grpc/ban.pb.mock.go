// Code generated by MockGen. DO NOT EDIT.
// Source: app/service/user/ban/api/grpc/ban.pb.go

// Package grpc is a generated GoMock package.
package grpc

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockBanClient is a mock of BanClient interface
type MockBanClient struct {
	ctrl     *gomock.Controller
	recorder *MockBanClientMockRecorder
}

// MockBanClientMockRecorder is the mock recorder for MockBanClient
type MockBanClientMockRecorder struct {
	mock *MockBanClient
}

// NewMockBanClient creates a new mock instance
func NewMockBanClient(ctrl *gomock.Controller) *MockBanClient {
	mock := &MockBanClient{ctrl: ctrl}
	mock.recorder = &MockBanClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBanClient) EXPECT() *MockBanClientMockRecorder {
	return m.recorder
}

// Info mocks base method
func (m *MockBanClient) Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Info", varargs...)
	ret0, _ := ret[0].(*InfoResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockBanClientMockRecorder) Info(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockBanClient)(nil).Info), varargs...)
}

// Add mocks base method
func (m *MockBanClient) Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Add", varargs...)
	ret0, _ := ret[0].(*AddResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockBanClientMockRecorder) Add(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockBanClient)(nil).Add), varargs...)
}

// Remove mocks base method
func (m *MockBanClient) Remove(ctx context.Context, in *RemoveReq, opts ...grpc.CallOption) (*RemoveResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Remove", varargs...)
	ret0, _ := ret[0].(*RemoveResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Remove indicates an expected call of Remove
func (mr *MockBanClientMockRecorder) Remove(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockBanClient)(nil).Remove), varargs...)
}

// BanCheck mocks base method
func (m *MockBanClient) BanCheck(ctx context.Context, in *Info2Req, opts ...grpc.CallOption) (*InfoResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BanCheck", varargs...)
	ret0, _ := ret[0].(*InfoResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BanCheck indicates an expected call of BanCheck
func (mr *MockBanClientMockRecorder) BanCheck(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BanCheck", reflect.TypeOf((*MockBanClient)(nil).BanCheck), varargs...)
}

// MockBanServer is a mock of BanServer interface
type MockBanServer struct {
	ctrl     *gomock.Controller
	recorder *MockBanServerMockRecorder
}

// MockBanServerMockRecorder is the mock recorder for MockBanServer
type MockBanServerMockRecorder struct {
	mock *MockBanServer
}

// NewMockBanServer creates a new mock instance
func NewMockBanServer(ctrl *gomock.Controller) *MockBanServer {
	mock := &MockBanServer{ctrl: ctrl}
	mock.recorder = &MockBanServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBanServer) EXPECT() *MockBanServerMockRecorder {
	return m.recorder
}

// Info mocks base method
func (m *MockBanServer) Info(arg0 context.Context, arg1 *InfoReq) (*InfoResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0, arg1)
	ret0, _ := ret[0].(*InfoResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockBanServerMockRecorder) Info(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockBanServer)(nil).Info), arg0, arg1)
}

// Add mocks base method
func (m *MockBanServer) Add(arg0 context.Context, arg1 *AddReq) (*AddResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(*AddResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockBanServerMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockBanServer)(nil).Add), arg0, arg1)
}

// Remove mocks base method
func (m *MockBanServer) Remove(arg0 context.Context, arg1 *RemoveReq) (*RemoveResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0, arg1)
	ret0, _ := ret[0].(*RemoveResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Remove indicates an expected call of Remove
func (mr *MockBanServerMockRecorder) Remove(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockBanServer)(nil).Remove), arg0, arg1)
}

// BanCheck mocks base method
func (m *MockBanServer) BanCheck(arg0 context.Context, arg1 *Info2Req) (*InfoResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BanCheck", arg0, arg1)
	ret0, _ := ret[0].(*InfoResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BanCheck indicates an expected call of BanCheck
func (mr *MockBanServerMockRecorder) BanCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BanCheck", reflect.TypeOf((*MockBanServer)(nil).BanCheck), arg0, arg1)
}
