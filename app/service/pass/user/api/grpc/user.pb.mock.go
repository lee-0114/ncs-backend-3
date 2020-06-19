// Code generated by MockGen. DO NOT EDIT.
// Source: app/service/pass/user/api/grpc/user.pb.go

// Package grpc is a generated GoMock package.
package grpc

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockUserClient is a mock of UserClient interface
type MockUserClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserClientMockRecorder
}

// MockUserClientMockRecorder is the mock recorder for MockUserClient
type MockUserClientMockRecorder struct {
	mock *MockUserClient
}

// NewMockUserClient creates a new mock instance
func NewMockUserClient(ctrl *gomock.Controller) *MockUserClient {
	mock := &MockUserClient{ctrl: ctrl}
	mock.recorder = &MockUserClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserClient) EXPECT() *MockUserClientMockRecorder {
	return m.recorder
}

// Info mocks base method
func (m *MockUserClient) Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
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
func (mr *MockUserClientMockRecorder) Info(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockUserClient)(nil).Info), varargs...)
}

// AddPoint mocks base method
func (m *MockUserClient) AddPoint(ctx context.Context, in *AddPointReq, opts ...grpc.CallOption) (*AddPointResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddPoint", varargs...)
	ret0, _ := ret[0].(*AddPointResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPoint indicates an expected call of AddPoint
func (mr *MockUserClientMockRecorder) AddPoint(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPoint", reflect.TypeOf((*MockUserClient)(nil).AddPoint), varargs...)
}

// Update mocks base method
func (m *MockUserClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*UpdateResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockUserClientMockRecorder) Update(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserClient)(nil).Update), varargs...)
}

// MockUserServer is a mock of UserServer interface
type MockUserServer struct {
	ctrl     *gomock.Controller
	recorder *MockUserServerMockRecorder
}

// MockUserServerMockRecorder is the mock recorder for MockUserServer
type MockUserServerMockRecorder struct {
	mock *MockUserServer
}

// NewMockUserServer creates a new mock instance
func NewMockUserServer(ctrl *gomock.Controller) *MockUserServer {
	mock := &MockUserServer{ctrl: ctrl}
	mock.recorder = &MockUserServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserServer) EXPECT() *MockUserServerMockRecorder {
	return m.recorder
}

// Info mocks base method
func (m *MockUserServer) Info(arg0 context.Context, arg1 *InfoReq) (*InfoResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0, arg1)
	ret0, _ := ret[0].(*InfoResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockUserServerMockRecorder) Info(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockUserServer)(nil).Info), arg0, arg1)
}

// AddPoint mocks base method
func (m *MockUserServer) AddPoint(arg0 context.Context, arg1 *AddPointReq) (*AddPointResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPoint", arg0, arg1)
	ret0, _ := ret[0].(*AddPointResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPoint indicates an expected call of AddPoint
func (mr *MockUserServerMockRecorder) AddPoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPoint", reflect.TypeOf((*MockUserServer)(nil).AddPoint), arg0, arg1)
}

// Update mocks base method
func (m *MockUserServer) Update(arg0 context.Context, arg1 *UpdateReq) (*UpdateResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*UpdateResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockUserServerMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserServer)(nil).Update), arg0, arg1)
}