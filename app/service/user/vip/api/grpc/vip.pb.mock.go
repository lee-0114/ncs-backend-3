// Code generated by MockGen. DO NOT EDIT.
// Source: app/service/user/vip/api/grpc/vip.pb.go

// Package grpc is a generated GoMock package.
package grpc

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockVIPClient is a mock of VIPClient interface
type MockVIPClient struct {
	ctrl     *gomock.Controller
	recorder *MockVIPClientMockRecorder
}

// MockVIPClientMockRecorder is the mock recorder for MockVIPClient
type MockVIPClientMockRecorder struct {
	mock *MockVIPClient
}

// NewMockVIPClient creates a new mock instance
func NewMockVIPClient(ctrl *gomock.Controller) *MockVIPClient {
	mock := &MockVIPClient{ctrl: ctrl}
	mock.recorder = &MockVIPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVIPClient) EXPECT() *MockVIPClientMockRecorder {
	return m.recorder
}

// Info mocks base method
func (m *MockVIPClient) Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
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
func (mr *MockVIPClientMockRecorder) Info(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockVIPClient)(nil).Info), varargs...)
}

// Renewal mocks base method
func (m *MockVIPClient) Renewal(ctx context.Context, in *RenewalReq, opts ...grpc.CallOption) (*RenewalResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Renewal", varargs...)
	ret0, _ := ret[0].(*RenewalResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Renewal indicates an expected call of Renewal
func (mr *MockVIPClientMockRecorder) Renewal(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Renewal", reflect.TypeOf((*MockVIPClient)(nil).Renewal), varargs...)
}

// AddPoint mocks base method
func (m *MockVIPClient) AddPoint(ctx context.Context, in *AddPointReq, opts ...grpc.CallOption) (*AddPointResp, error) {
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
func (mr *MockVIPClientMockRecorder) AddPoint(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPoint", reflect.TypeOf((*MockVIPClient)(nil).AddPoint), varargs...)
}

// MockVIPServer is a mock of VIPServer interface
type MockVIPServer struct {
	ctrl     *gomock.Controller
	recorder *MockVIPServerMockRecorder
}

// MockVIPServerMockRecorder is the mock recorder for MockVIPServer
type MockVIPServerMockRecorder struct {
	mock *MockVIPServer
}

// NewMockVIPServer creates a new mock instance
func NewMockVIPServer(ctrl *gomock.Controller) *MockVIPServer {
	mock := &MockVIPServer{ctrl: ctrl}
	mock.recorder = &MockVIPServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVIPServer) EXPECT() *MockVIPServerMockRecorder {
	return m.recorder
}

// Info mocks base method
func (m *MockVIPServer) Info(arg0 context.Context, arg1 *InfoReq) (*InfoResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0, arg1)
	ret0, _ := ret[0].(*InfoResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockVIPServerMockRecorder) Info(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockVIPServer)(nil).Info), arg0, arg1)
}

// Renewal mocks base method
func (m *MockVIPServer) Renewal(arg0 context.Context, arg1 *RenewalReq) (*RenewalResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Renewal", arg0, arg1)
	ret0, _ := ret[0].(*RenewalResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Renewal indicates an expected call of Renewal
func (mr *MockVIPServerMockRecorder) Renewal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Renewal", reflect.TypeOf((*MockVIPServer)(nil).Renewal), arg0, arg1)
}

// AddPoint mocks base method
func (m *MockVIPServer) AddPoint(arg0 context.Context, arg1 *AddPointReq) (*AddPointResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPoint", arg0, arg1)
	ret0, _ := ret[0].(*AddPointResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPoint indicates an expected call of AddPoint
func (mr *MockVIPServerMockRecorder) AddPoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPoint", reflect.TypeOf((*MockVIPServer)(nil).AddPoint), arg0, arg1)
}
