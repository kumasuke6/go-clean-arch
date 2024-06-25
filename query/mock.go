// Code generated by MockGen. DO NOT EDIT.
// Source: ./query/interface.go
//
// Generated by this command:
//
//	mockgen -source=./query/interface.go -destination=./query/mock.go -package=query
//

// Package query is a generated GoMock package.
package query

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUserQuery is a mock of UserQuery interface.
type MockUserQuery struct {
	ctrl     *gomock.Controller
	recorder *MockUserQueryMockRecorder
}

// MockUserQueryMockRecorder is the mock recorder for MockUserQuery.
type MockUserQueryMockRecorder struct {
	mock *MockUserQuery
}

// NewMockUserQuery creates a new mock instance.
func NewMockUserQuery(ctrl *gomock.Controller) *MockUserQuery {
	mock := &MockUserQuery{ctrl: ctrl}
	mock.recorder = &MockUserQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserQuery) EXPECT() *MockUserQueryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserQuery) Create() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(string)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserQueryMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserQuery)(nil).Create))
}

// Delete mocks base method.
func (m *MockUserQuery) Delete() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(string)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserQueryMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserQuery)(nil).Delete))
}

// Read mocks base method.
func (m *MockUserQuery) Read() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read")
	ret0, _ := ret[0].(string)
	return ret0
}

// Read indicates an expected call of Read.
func (mr *MockUserQueryMockRecorder) Read() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockUserQuery)(nil).Read))
}

// Update mocks base method.
func (m *MockUserQuery) Update() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update")
	ret0, _ := ret[0].(string)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserQueryMockRecorder) Update() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserQuery)(nil).Update))
}

// MockMessageQuery is a mock of MessageQuery interface.
type MockMessageQuery struct {
	ctrl     *gomock.Controller
	recorder *MockMessageQueryMockRecorder
}

// MockMessageQueryMockRecorder is the mock recorder for MockMessageQuery.
type MockMessageQueryMockRecorder struct {
	mock *MockMessageQuery
}

// NewMockMessageQuery creates a new mock instance.
func NewMockMessageQuery(ctrl *gomock.Controller) *MockMessageQuery {
	mock := &MockMessageQuery{ctrl: ctrl}
	mock.recorder = &MockMessageQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageQuery) EXPECT() *MockMessageQueryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMessageQuery) Create() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(string)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMessageQueryMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMessageQuery)(nil).Create))
}

// Delete mocks base method.
func (m *MockMessageQuery) Delete() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(string)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMessageQueryMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMessageQuery)(nil).Delete))
}

// Read mocks base method.
func (m *MockMessageQuery) Read() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read")
	ret0, _ := ret[0].(string)
	return ret0
}

// Read indicates an expected call of Read.
func (mr *MockMessageQueryMockRecorder) Read() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockMessageQuery)(nil).Read))
}
