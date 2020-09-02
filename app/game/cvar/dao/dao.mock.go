// Code generated by MockGen. DO NOT EDIT.
// Source: dao.go

// Package dao is a generated GoMock package.
package dao

import (
	model "backend/app/game/cvar/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDao is a mock of Dao interface
type MockDao struct {
	ctrl     *gomock.Controller
	recorder *MockDaoMockRecorder
}

// MockDaoMockRecorder is the mock recorder for MockDao
type MockDaoMockRecorder struct {
	mock *MockDao
}

// NewMockDao creates a new mock instance
func NewMockDao(ctrl *gomock.Controller) *MockDao {
	mock := &MockDao{ctrl: ctrl}
	mock.recorder = &MockDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDao) EXPECT() *MockDaoMockRecorder {
	return m.recorder
}

// GetCVars mocks base method
func (m *MockDao) GetCVars() ([]*model.CVar, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCVars")
	ret0, _ := ret[0].([]*model.CVar)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCVars indicates an expected call of GetCVars
func (mr *MockDaoMockRecorder) GetCVars() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCVars", reflect.TypeOf((*MockDao)(nil).GetCVars))
}

// UpdatedCVar mocks base method
func (m *MockDao) UpdatedCVar(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatedCVar", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatedCVar indicates an expected call of UpdatedCVar
func (mr *MockDaoMockRecorder) UpdatedCVar(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatedCVar", reflect.TypeOf((*MockDao)(nil).UpdatedCVar), id)
}

// Healthy mocks base method
func (m *MockDao) Healthy() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Healthy")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Healthy indicates an expected call of Healthy
func (mr *MockDaoMockRecorder) Healthy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Healthy", reflect.TypeOf((*MockDao)(nil).Healthy))
}

// Close mocks base method
func (m *MockDao) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockDaoMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDao)(nil).Close))
}