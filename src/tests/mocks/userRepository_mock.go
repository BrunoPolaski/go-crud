package mocks

import (
	reflect "reflect"

	rest_err "github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	user "github.com/BrunoPolaski/go-crud/src/model/user"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUserRepository mocks base method.
func (m *MockUserRepository) CreateUserRepository(userDomain user.UserDomainInterface) (user.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserRepository", userDomain)
	ret0, _ := ret[0].(user.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// CreateUserRepository indicates an expected call of CreateUserRepository.
func (mr *MockUserRepositoryMockRecorder) CreateUserRepository(userDomain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserRepository", reflect.TypeOf((*MockUserRepository)(nil).CreateUserRepository), userDomain)
}

// DeleteUserRepository mocks base method.
func (m *MockUserRepository) DeleteUserRepository(id string) *rest_err.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserRepository", id)
	ret0, _ := ret[0].(*rest_err.RestErr)
	return ret0
}

// DeleteUserRepository indicates an expected call of DeleteUserRepository.
func (mr *MockUserRepositoryMockRecorder) DeleteUserRepository(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserRepository", reflect.TypeOf((*MockUserRepository)(nil).DeleteUserRepository), id)
}

// FindUserByEmailRepository mocks base method.
func (m *MockUserRepository) FindUserByEmailRepository(email string) (user.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmailRepository", email)
	ret0, _ := ret[0].(user.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// FindUserByEmailRepository indicates an expected call of FindUserByEmailRepository.
func (mr *MockUserRepositoryMockRecorder) FindUserByEmailRepository(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmailRepository", reflect.TypeOf((*MockUserRepository)(nil).FindUserByEmailRepository), email)
}

// FindUserByIDRepository mocks base method.
func (m *MockUserRepository) FindUserByIDRepository(id string) (user.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByIDRepository", id)
	ret0, _ := ret[0].(user.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// FindUserByIDRepository indicates an expected call of FindUserByIDRepository.
func (mr *MockUserRepositoryMockRecorder) FindUserByIDRepository(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByIDRepository", reflect.TypeOf((*MockUserRepository)(nil).FindUserByIDRepository), id)
}

// UpdateUserRepository mocks base method.
func (m *MockUserRepository) UpdateUserRepository(userDomain user.UserDomainInterface, id string) *rest_err.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserRepository", userDomain, id)
	ret0, _ := ret[0].(*rest_err.RestErr)
	return ret0
}

// UpdateUserRepository indicates an expected call of UpdateUserRepository.
func (mr *MockUserRepositoryMockRecorder) UpdateUserRepository(userDomain, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserRepository", reflect.TypeOf((*MockUserRepository)(nil).UpdateUserRepository), userDomain, id)
}
