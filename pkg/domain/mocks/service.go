// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/opencars/edrmvs/pkg/domain (interfaces: RegistrationService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/opencars/edrmvs/pkg/domain/model"
)

// MockRegistrationService is a mock of RegistrationService interface.
type MockRegistrationService struct {
	ctrl     *gomock.Controller
	recorder *MockRegistrationServiceMockRecorder
}

// MockRegistrationServiceMockRecorder is the mock recorder for MockRegistrationService.
type MockRegistrationServiceMockRecorder struct {
	mock *MockRegistrationService
}

// NewMockRegistrationService creates a new mock instance.
func NewMockRegistrationService(ctrl *gomock.Controller) *MockRegistrationService {
	mock := &MockRegistrationService{ctrl: ctrl}
	mock.recorder = &MockRegistrationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistrationService) EXPECT() *MockRegistrationServiceMockRecorder {
	return m.recorder
}

// FindByCode mocks base method.
func (m *MockRegistrationService) FindByCode(arg0 context.Context, arg1 string) (*model.Registration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByCode", arg0, arg1)
	ret0, _ := ret[0].(*model.Registration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByCode indicates an expected call of FindByCode.
func (mr *MockRegistrationServiceMockRecorder) FindByCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByCode", reflect.TypeOf((*MockRegistrationService)(nil).FindByCode), arg0, arg1)
}

// FindByNumber mocks base method.
func (m *MockRegistrationService) FindByNumber(arg0 context.Context, arg1 string) ([]model.Registration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByNumber", arg0, arg1)
	ret0, _ := ret[0].([]model.Registration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByNumber indicates an expected call of FindByNumber.
func (mr *MockRegistrationServiceMockRecorder) FindByNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByNumber", reflect.TypeOf((*MockRegistrationService)(nil).FindByNumber), arg0, arg1)
}

// FindByVIN mocks base method.
func (m *MockRegistrationService) FindByVIN(arg0 context.Context, arg1 string, arg2 bool) ([]model.Registration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByVIN", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.Registration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByVIN indicates an expected call of FindByVIN.
func (mr *MockRegistrationServiceMockRecorder) FindByVIN(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByVIN", reflect.TypeOf((*MockRegistrationService)(nil).FindByVIN), arg0, arg1, arg2)
}

// Health mocks base method.
func (m *MockRegistrationService) Health(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Health indicates an expected call of Health.
func (mr *MockRegistrationServiceMockRecorder) Health(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockRegistrationService)(nil).Health), arg0)
}
