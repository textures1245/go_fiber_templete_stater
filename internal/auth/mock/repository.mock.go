// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/phakh/go/src/github.com/textures1245/go-template/internal/auth/repository.go
//
// Generated by this command:
//
//	mockgen -source=/Users/phakh/go/src/github.com/textures1245/go-template/internal/auth/repository.go AuthRepository
//

// Package mock_auth is a generated GoMock package.
package mock_auth

import (
	reflect "reflect"

	dtos "github.com/textures1245/go-template/internal/auth/dtos"
	entities "github.com/textures1245/go-template/internal/auth/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthRepository is a mock of AuthRepository interface.
type MockAuthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepositoryMockRecorder
}

// MockAuthRepositoryMockRecorder is the mock recorder for MockAuthRepository.
type MockAuthRepositoryMockRecorder struct {
	mock *MockAuthRepository
}

// NewMockAuthRepository creates a new mock instance.
func NewMockAuthRepository(ctrl *gomock.Controller) *MockAuthRepository {
	mock := &MockAuthRepository{ctrl: ctrl}
	mock.recorder = &MockAuthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRepository) EXPECT() *MockAuthRepositoryMockRecorder {
	return m.recorder
}

// SignUsersAccessToken mocks base method.
func (m *MockAuthRepository) SignUsersAccessToken(req *entities.UserSignToken) (*dtos.UserTokenRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUsersAccessToken", req)
	ret0, _ := ret[0].(*dtos.UserTokenRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUsersAccessToken indicates an expected call of SignUsersAccessToken.
func (mr *MockAuthRepositoryMockRecorder) SignUsersAccessToken(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUsersAccessToken", reflect.TypeOf((*MockAuthRepository)(nil).SignUsersAccessToken), req)
}
