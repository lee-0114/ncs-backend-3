// Code generated by MockGen. DO NOT EDIT.
// Source: app/game/user/api/grpc/user.pb.go

// Package grpc is a generated GoMock package.
package grpc

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockGameClient is a mock of GameClient interface
type MockGameClient struct {
	ctrl     *gomock.Controller
	recorder *MockGameClientMockRecorder
}

// MockGameClientMockRecorder is the mock recorder for MockGameClient
type MockGameClientMockRecorder struct {
	mock *MockGameClient
}

// NewMockGameClient creates a new mock instance
func NewMockGameClient(ctrl *gomock.Controller) *MockGameClient {
	mock := &MockGameClient{ctrl: ctrl}
	mock.recorder = &MockGameClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGameClient) EXPECT() *MockGameClientMockRecorder {
	return m.recorder
}

// PlayerConnect mocks base method
func (m *MockGameClient) PlayerConnect(ctx context.Context, in *PlayerConnectReq, opts ...grpc.CallOption) (*PlayerConnectResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PlayerConnect", varargs...)
	ret0, _ := ret[0].(*PlayerConnectResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PlayerConnect indicates an expected call of PlayerConnect
func (mr *MockGameClientMockRecorder) PlayerConnect(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlayerConnect", reflect.TypeOf((*MockGameClient)(nil).PlayerConnect), varargs...)
}

// PlayerDisconnect mocks base method
func (m *MockGameClient) PlayerDisconnect(ctx context.Context, in *PlayerDisconnectReq, opts ...grpc.CallOption) (*PlayerDisconnectResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PlayerDisconnect", varargs...)
	ret0, _ := ret[0].(*PlayerDisconnectResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PlayerDisconnect indicates an expected call of PlayerDisconnect
func (mr *MockGameClientMockRecorder) PlayerDisconnect(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlayerDisconnect", reflect.TypeOf((*MockGameClient)(nil).PlayerDisconnect), varargs...)
}

// MockGameServer is a mock of GameServer interface
type MockGameServer struct {
	ctrl     *gomock.Controller
	recorder *MockGameServerMockRecorder
}

// MockGameServerMockRecorder is the mock recorder for MockGameServer
type MockGameServerMockRecorder struct {
	mock *MockGameServer
}

// NewMockGameServer creates a new mock instance
func NewMockGameServer(ctrl *gomock.Controller) *MockGameServer {
	mock := &MockGameServer{ctrl: ctrl}
	mock.recorder = &MockGameServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGameServer) EXPECT() *MockGameServerMockRecorder {
	return m.recorder
}

// PlayerConnect mocks base method
func (m *MockGameServer) PlayerConnect(arg0 context.Context, arg1 *PlayerConnectReq) (*PlayerConnectResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlayerConnect", arg0, arg1)
	ret0, _ := ret[0].(*PlayerConnectResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PlayerConnect indicates an expected call of PlayerConnect
func (mr *MockGameServerMockRecorder) PlayerConnect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlayerConnect", reflect.TypeOf((*MockGameServer)(nil).PlayerConnect), arg0, arg1)
}

// PlayerDisconnect mocks base method
func (m *MockGameServer) PlayerDisconnect(arg0 context.Context, arg1 *PlayerDisconnectReq) (*PlayerDisconnectResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlayerDisconnect", arg0, arg1)
	ret0, _ := ret[0].(*PlayerDisconnectResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PlayerDisconnect indicates an expected call of PlayerDisconnect
func (mr *MockGameServerMockRecorder) PlayerDisconnect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlayerDisconnect", reflect.TypeOf((*MockGameServer)(nil).PlayerDisconnect), arg0, arg1)
}