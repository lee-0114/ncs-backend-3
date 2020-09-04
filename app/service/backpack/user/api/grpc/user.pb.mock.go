// Code generated by MockGen. DO NOT EDIT.
// Source: app/service/backpack/user/api/grpc/user.pb.go

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

// Init mocks base method
func (m *MockUserClient) Init(ctx context.Context, in *InitReq, opts ...grpc.CallOption) (*InitResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Init", varargs...)
	ret0, _ := ret[0].(*InitResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Init indicates an expected call of Init
func (mr *MockUserClientMockRecorder) Init(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockUserClient)(nil).Init), varargs...)
}

// GetItems mocks base method
func (m *MockUserClient) GetItems(ctx context.Context, in *GetItemsReq, opts ...grpc.CallOption) (*GetItemsResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetItems", varargs...)
	ret0, _ := ret[0].(*GetItemsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItems indicates an expected call of GetItems
func (mr *MockUserClientMockRecorder) GetItems(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItems", reflect.TypeOf((*MockUserClient)(nil).GetItems), varargs...)
}

// AddItems mocks base method
func (m *MockUserClient) AddItems(ctx context.Context, in *AddItemsReq, opts ...grpc.CallOption) (*AddItemsResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddItems", varargs...)
	ret0, _ := ret[0].(*AddItemsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddItems indicates an expected call of AddItems
func (mr *MockUserClientMockRecorder) AddItems(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddItems", reflect.TypeOf((*MockUserClient)(nil).AddItems), varargs...)
}

// RemoveItem mocks base method
func (m *MockUserClient) RemoveItem(ctx context.Context, in *RemoveItemReq, opts ...grpc.CallOption) (*RemoveItemResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveItem", varargs...)
	ret0, _ := ret[0].(*RemoveItemResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveItem indicates an expected call of RemoveItem
func (mr *MockUserClientMockRecorder) RemoveItem(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveItem", reflect.TypeOf((*MockUserClient)(nil).RemoveItem), varargs...)
}

// HaveItem mocks base method
func (m *MockUserClient) HaveItem(ctx context.Context, in *HaveItemReq, opts ...grpc.CallOption) (*HaveItemResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HaveItem", varargs...)
	ret0, _ := ret[0].(*HaveItemResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HaveItem indicates an expected call of HaveItem
func (mr *MockUserClientMockRecorder) HaveItem(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HaveItem", reflect.TypeOf((*MockUserClient)(nil).HaveItem), varargs...)
}

// AddAllItems mocks base method
func (m *MockUserClient) AddAllItems(ctx context.Context, in *AddAllItemsReq, opts ...grpc.CallOption) (*AddAllItemsResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddAllItems", varargs...)
	ret0, _ := ret[0].(*AddAllItemsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAllItems indicates an expected call of AddAllItems
func (mr *MockUserClientMockRecorder) AddAllItems(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAllItems", reflect.TypeOf((*MockUserClient)(nil).AddAllItems), varargs...)
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

// Init mocks base method
func (m *MockUserServer) Init(arg0 context.Context, arg1 *InitReq) (*InitResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0, arg1)
	ret0, _ := ret[0].(*InitResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Init indicates an expected call of Init
func (mr *MockUserServerMockRecorder) Init(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockUserServer)(nil).Init), arg0, arg1)
}

// GetItems mocks base method
func (m *MockUserServer) GetItems(arg0 context.Context, arg1 *GetItemsReq) (*GetItemsResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItems", arg0, arg1)
	ret0, _ := ret[0].(*GetItemsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItems indicates an expected call of GetItems
func (mr *MockUserServerMockRecorder) GetItems(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItems", reflect.TypeOf((*MockUserServer)(nil).GetItems), arg0, arg1)
}

// AddItems mocks base method
func (m *MockUserServer) AddItems(arg0 context.Context, arg1 *AddItemsReq) (*AddItemsResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddItems", arg0, arg1)
	ret0, _ := ret[0].(*AddItemsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddItems indicates an expected call of AddItems
func (mr *MockUserServerMockRecorder) AddItems(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddItems", reflect.TypeOf((*MockUserServer)(nil).AddItems), arg0, arg1)
}

// RemoveItem mocks base method
func (m *MockUserServer) RemoveItem(arg0 context.Context, arg1 *RemoveItemReq) (*RemoveItemResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveItem", arg0, arg1)
	ret0, _ := ret[0].(*RemoveItemResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveItem indicates an expected call of RemoveItem
func (mr *MockUserServerMockRecorder) RemoveItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveItem", reflect.TypeOf((*MockUserServer)(nil).RemoveItem), arg0, arg1)
}

// HaveItem mocks base method
func (m *MockUserServer) HaveItem(arg0 context.Context, arg1 *HaveItemReq) (*HaveItemResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HaveItem", arg0, arg1)
	ret0, _ := ret[0].(*HaveItemResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HaveItem indicates an expected call of HaveItem
func (mr *MockUserServerMockRecorder) HaveItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HaveItem", reflect.TypeOf((*MockUserServer)(nil).HaveItem), arg0, arg1)
}

// AddAllItems mocks base method
func (m *MockUserServer) AddAllItems(arg0 context.Context, arg1 *AddAllItemsReq) (*AddAllItemsResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAllItems", arg0, arg1)
	ret0, _ := ret[0].(*AddAllItemsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAllItems indicates an expected call of AddAllItems
func (mr *MockUserServerMockRecorder) AddAllItems(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAllItems", reflect.TypeOf((*MockUserServer)(nil).AddAllItems), arg0, arg1)
}
