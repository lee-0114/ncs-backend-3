// Code generated by MockGen. DO NOT EDIT.
// Source: app/service/user/money/api/grpc/money.pb.go

// Package grpc is a generated GoMock package.
package grpc

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockMoneyClient is a mock of MoneyClient interface
type MockMoneyClient struct {
	ctrl     *gomock.Controller
	recorder *MockMoneyClientMockRecorder
}

// MockMoneyClientMockRecorder is the mock recorder for MockMoneyClient
type MockMoneyClientMockRecorder struct {
	mock *MockMoneyClient
}

// NewMockMoneyClient creates a new mock instance
func NewMockMoneyClient(ctrl *gomock.Controller) *MockMoneyClient {
	mock := &MockMoneyClient{ctrl: ctrl}
	mock.recorder = &MockMoneyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMoneyClient) EXPECT() *MockMoneyClientMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockMoneyClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*GetResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockMoneyClientMockRecorder) Get(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMoneyClient)(nil).Get), varargs...)
}

// Pay mocks base method
func (m *MockMoneyClient) Pay(ctx context.Context, in *PayReq, opts ...grpc.CallOption) (*PayResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Pay", varargs...)
	ret0, _ := ret[0].(*PayResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pay indicates an expected call of Pay
func (mr *MockMoneyClientMockRecorder) Pay(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pay", reflect.TypeOf((*MockMoneyClient)(nil).Pay), varargs...)
}

// Give mocks base method
func (m *MockMoneyClient) Give(ctx context.Context, in *GiveReq, opts ...grpc.CallOption) (*GiveResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Give", varargs...)
	ret0, _ := ret[0].(*GiveResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Give indicates an expected call of Give
func (mr *MockMoneyClientMockRecorder) Give(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Give", reflect.TypeOf((*MockMoneyClient)(nil).Give), varargs...)
}

// Records mocks base method
func (m *MockMoneyClient) Records(ctx context.Context, in *RecordsReq, opts ...grpc.CallOption) (*RecordsResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Records", varargs...)
	ret0, _ := ret[0].(*RecordsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Records indicates an expected call of Records
func (mr *MockMoneyClientMockRecorder) Records(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Records", reflect.TypeOf((*MockMoneyClient)(nil).Records), varargs...)
}

// MockMoneyServer is a mock of MoneyServer interface
type MockMoneyServer struct {
	ctrl     *gomock.Controller
	recorder *MockMoneyServerMockRecorder
}

// MockMoneyServerMockRecorder is the mock recorder for MockMoneyServer
type MockMoneyServerMockRecorder struct {
	mock *MockMoneyServer
}

// NewMockMoneyServer creates a new mock instance
func NewMockMoneyServer(ctrl *gomock.Controller) *MockMoneyServer {
	mock := &MockMoneyServer{ctrl: ctrl}
	mock.recorder = &MockMoneyServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMoneyServer) EXPECT() *MockMoneyServerMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockMoneyServer) Get(arg0 context.Context, arg1 *GetReq) (*GetResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*GetResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockMoneyServerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMoneyServer)(nil).Get), arg0, arg1)
}

// Pay mocks base method
func (m *MockMoneyServer) Pay(arg0 context.Context, arg1 *PayReq) (*PayResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pay", arg0, arg1)
	ret0, _ := ret[0].(*PayResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pay indicates an expected call of Pay
func (mr *MockMoneyServerMockRecorder) Pay(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pay", reflect.TypeOf((*MockMoneyServer)(nil).Pay), arg0, arg1)
}

// Give mocks base method
func (m *MockMoneyServer) Give(arg0 context.Context, arg1 *GiveReq) (*GiveResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Give", arg0, arg1)
	ret0, _ := ret[0].(*GiveResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Give indicates an expected call of Give
func (mr *MockMoneyServerMockRecorder) Give(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Give", reflect.TypeOf((*MockMoneyServer)(nil).Give), arg0, arg1)
}

// Records mocks base method
func (m *MockMoneyServer) Records(arg0 context.Context, arg1 *RecordsReq) (*RecordsResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Records", arg0, arg1)
	ret0, _ := ret[0].(*RecordsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Records indicates an expected call of Records
func (mr *MockMoneyServerMockRecorder) Records(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Records", reflect.TypeOf((*MockMoneyServer)(nil).Records), arg0, arg1)
}
