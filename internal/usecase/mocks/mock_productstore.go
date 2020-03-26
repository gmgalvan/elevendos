// Code generated by MockGen. DO NOT EDIT.
// Source: lab/productLab/internal/usecase (interfaces: ProductStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "lab/productLab/internal/entity"
	reflect "reflect"
)

// MockProductStore is a mock of ProductStore interface
type MockProductStore struct {
	ctrl     *gomock.Controller
	recorder *MockProductStoreMockRecorder
}

// MockProductStoreMockRecorder is the mock recorder for MockProductStore
type MockProductStoreMockRecorder struct {
	mock *MockProductStore
}

// NewMockProductStore creates a new mock instance
func NewMockProductStore(ctrl *gomock.Controller) *MockProductStore {
	mock := &MockProductStore{ctrl: ctrl}
	mock.recorder = &MockProductStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProductStore) EXPECT() *MockProductStoreMockRecorder {
	return m.recorder
}

// ByID mocks base method
func (m *MockProductStore) ByID(arg0 context.Context, arg1 int) (*entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByID", arg0, arg1)
	ret0, _ := ret[0].(*entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ByID indicates an expected call of ByID
func (mr *MockProductStoreMockRecorder) ByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByID", reflect.TypeOf((*MockProductStore)(nil).ByID), arg0, arg1)
}

// Create mocks base method
func (m *MockProductStore) Create(arg0 context.Context, arg1 *entity.Product) (*entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockProductStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockProductStore) Delete(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockProductStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductStore)(nil).Delete), arg0, arg1)
}

// List mocks base method
func (m *MockProductStore) List(arg0 context.Context, arg1, arg2 int) ([]*entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockProductStoreMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProductStore)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method
func (m *MockProductStore) Update(arg0 context.Context, arg1 int, arg2 *entity.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockProductStoreMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductStore)(nil).Update), arg0, arg1, arg2)
}
