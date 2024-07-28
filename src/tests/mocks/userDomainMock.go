// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/model/user/userDomainInterface.go
//
// Generated by this command:
//
//	mockgen -source=./src/model/user/userDomainInterface.go -destination=./src/tests/mocks/userDomainMock.go
//

// Package mock_model is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	rest_err "github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	gomock "github.com/golang/mock/gomock"
)

// MockUserDomainInterface is a mock of UserDomainInterface interface.
type MockUserDomainInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserDomainInterfaceMockRecorder
}

// MockUserDomainInterfaceMockRecorder is the mock recorder for MockUserDomainInterface.
type MockUserDomainInterfaceMockRecorder struct {
	mock *MockUserDomainInterface
}

// NewMockUserDomainInterface creates a new mock instance.
func NewMockUserDomainInterface(ctrl *gomock.Controller) *MockUserDomainInterface {
	mock := &MockUserDomainInterface{ctrl: ctrl}
	mock.recorder = &MockUserDomainInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDomainInterface) EXPECT() *MockUserDomainInterfaceMockRecorder {
	return m.recorder
}

// ComparePassword mocks base method.
func (m *MockUserDomainInterface) ComparePassword(password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComparePassword", password)
	ret0, _ := ret[0].(error)
	return ret0
}

// ComparePassword indicates an expected call of ComparePassword.
func (mr *MockUserDomainInterfaceMockRecorder) ComparePassword(password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComparePassword", reflect.TypeOf((*MockUserDomainInterface)(nil).ComparePassword), password)
}

// EncryptPassword mocks base method.
func (m *MockUserDomainInterface) EncryptPassword() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncryptPassword")
	ret0, _ := ret[0].(error)
	return ret0
}

// EncryptPassword indicates an expected call of EncryptPassword.
func (mr *MockUserDomainInterfaceMockRecorder) EncryptPassword() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptPassword", reflect.TypeOf((*MockUserDomainInterface)(nil).EncryptPassword))
}

// GenerateToken mocks base method.
func (m *MockUserDomainInterface) GenerateToken() (string, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockUserDomainInterfaceMockRecorder) GenerateToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockUserDomainInterface)(nil).GenerateToken))
}

// GetAge mocks base method.
func (m *MockUserDomainInterface) GetAge() int8 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAge")
	ret0, _ := ret[0].(int8)
	return ret0
}

// GetAge indicates an expected call of GetAge.
func (mr *MockUserDomainInterfaceMockRecorder) GetAge() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAge", reflect.TypeOf((*MockUserDomainInterface)(nil).GetAge))
}

// GetEmail mocks base method.
func (m *MockUserDomainInterface) GetEmail() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmail")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEmail indicates an expected call of GetEmail.
func (mr *MockUserDomainInterfaceMockRecorder) GetEmail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmail", reflect.TypeOf((*MockUserDomainInterface)(nil).GetEmail))
}

// GetID mocks base method.
func (m *MockUserDomainInterface) GetID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockUserDomainInterfaceMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockUserDomainInterface)(nil).GetID))
}

// GetName mocks base method.
func (m *MockUserDomainInterface) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockUserDomainInterfaceMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockUserDomainInterface)(nil).GetName))
}

// GetPassword mocks base method.
func (m *MockUserDomainInterface) GetPassword() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPassword")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPassword indicates an expected call of GetPassword.
func (mr *MockUserDomainInterfaceMockRecorder) GetPassword() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPassword", reflect.TypeOf((*MockUserDomainInterface)(nil).GetPassword))
}

// SetAge mocks base method.
func (m *MockUserDomainInterface) SetAge(arg0 int8) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAge", arg0)
}

// SetAge indicates an expected call of SetAge.
func (mr *MockUserDomainInterfaceMockRecorder) SetAge(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAge", reflect.TypeOf((*MockUserDomainInterface)(nil).SetAge), arg0)
}

// SetEmail mocks base method.
func (m *MockUserDomainInterface) SetEmail(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetEmail", arg0)
}

// SetEmail indicates an expected call of SetEmail.
func (mr *MockUserDomainInterfaceMockRecorder) SetEmail(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEmail", reflect.TypeOf((*MockUserDomainInterface)(nil).SetEmail), arg0)
}

// SetID mocks base method.
func (m *MockUserDomainInterface) SetID(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetID", arg0)
}

// SetID indicates an expected call of SetID.
func (mr *MockUserDomainInterfaceMockRecorder) SetID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetID", reflect.TypeOf((*MockUserDomainInterface)(nil).SetID), arg0)
}

// SetName mocks base method.
func (m *MockUserDomainInterface) SetName(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetName", arg0)
}

// SetName indicates an expected call of SetName.
func (mr *MockUserDomainInterfaceMockRecorder) SetName(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetName", reflect.TypeOf((*MockUserDomainInterface)(nil).SetName), arg0)
}

// SetPassword mocks base method.
func (m *MockUserDomainInterface) SetPassword(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPassword", arg0)
}

// SetPassword indicates an expected call of SetPassword.
func (mr *MockUserDomainInterfaceMockRecorder) SetPassword(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPassword", reflect.TypeOf((*MockUserDomainInterface)(nil).SetPassword), arg0)
}