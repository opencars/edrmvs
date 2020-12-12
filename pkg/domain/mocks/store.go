// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/opencars/edrmvs/pkg/domain (interfaces: RegistrationStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	domain "github.com/opencars/edrmvs/pkg/domain"
	reflect "reflect"
)

// MockRegistrationStore is a mock of RegistrationStore interface
type MockRegistrationStore struct {
	ctrl     *gomock.Controller
	recorder *MockRegistrationStoreMockRecorder
}

// MockRegistrationStoreMockRecorder is the mock recorder for MockRegistrationStore
type MockRegistrationStoreMockRecorder struct {
	mock *MockRegistrationStore
}

// NewMockRegistrationStore creates a new mock instance
func NewMockRegistrationStore(ctrl *gomock.Controller) *MockRegistrationStore {
	mock := &MockRegistrationStore{ctrl: ctrl}
	mock.recorder = &MockRegistrationStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegistrationStore) EXPECT() *MockRegistrationStoreMockRecorder {
	return m.recorder
}

// FindByCode mocks base method
func (m *MockRegistrationStore) FindByCode(arg0 context.Context, arg1 string) (*domain.Registration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByCode", arg0, arg1)
	ret0, _ := ret[0].(*domain.Registration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByCode indicates an expected call of FindByCode
func (mr *MockRegistrationStoreMockRecorder) FindByCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByCode", reflect.TypeOf((*MockRegistrationStore)(nil).FindByCode), arg0, arg1)
}

// FindByNumber mocks base method
func (m *MockRegistrationStore) FindByNumber(arg0 context.Context, arg1 string) ([]domain.Registration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByNumber", arg0, arg1)
	ret0, _ := ret[0].([]domain.Registration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByNumber indicates an expected call of FindByNumber
func (mr *MockRegistrationStoreMockRecorder) FindByNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByNumber", reflect.TypeOf((*MockRegistrationStore)(nil).FindByNumber), arg0, arg1)
}

// FindByVIN mocks base method
func (m *MockRegistrationStore) FindByVIN(arg0 context.Context, arg1 string) ([]domain.Registration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByVIN", arg0, arg1)
	ret0, _ := ret[0].([]domain.Registration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByVIN indicates an expected call of FindByVIN
func (mr *MockRegistrationStoreMockRecorder) FindByVIN(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByVIN", reflect.TypeOf((*MockRegistrationStore)(nil).FindByVIN), arg0, arg1)
}

// Health mocks base method
func (m *MockRegistrationStore) Health(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Health indicates an expected call of Health
func (mr *MockRegistrationStoreMockRecorder) Health(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockRegistrationStore)(nil).Health), arg0)
}
