// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/phakh/go/src/github.com/textures1245/go-template/internal/auth/usecase.go
//
// Generated by this command:
//
//	mockgen -source=/Users/phakh/go/src/github.com/textures1245/go-template/internal/auth/usecase.go
//

// Package mock_auth is a generated GoMock package.
package mock_auth

import (
	context "context"
	reflect "reflect"

	dtos "github.com/textures1245/go-template/internal/auth/dtos"
	entities "github.com/textures1245/go-template/internal/auth/entities"
	entities0 "github.com/textures1245/go-template/internal/user/entities"
	apperror "github.com/textures1245/go-template/pkg/apperror"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthUsecase is a mock of AuthUsecase interface.
type MockAuthUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockAuthUsecaseMockRecorder
}

// MockAuthUsecaseMockRecorder is the mock recorder for MockAuthUsecase.
type MockAuthUsecaseMockRecorder struct {
	mock *MockAuthUsecase
}

// NewMockAuthUsecase creates a new mock instance.
func NewMockAuthUsecase(ctrl *gomock.Controller) *MockAuthUsecase {
	mock := &MockAuthUsecase{ctrl: ctrl}
	mock.recorder = &MockAuthUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthUsecase) EXPECT() *MockAuthUsecaseMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuthUsecase) Login(ctx context.Context, req *entities.UsersCredentials) (*dtos.UserTokenRes, int, *apperror.CErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, req)
	ret0, _ := ret[0].(*dtos.UserTokenRes)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(*apperror.CErr)
	return ret0, ret1, ret2
}

// Login indicates an expected call of Login.
func (mr *MockAuthUsecaseMockRecorder) Login(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthUsecase)(nil).Login), ctx, req)
}

// Register mocks base method.
func (m *MockAuthUsecase) Register(ctx context.Context, req *entities0.UserCreatedReq) (*dtos.UsersRegisteredRes, int, *apperror.CErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, req)
	ret0, _ := ret[0].(*dtos.UsersRegisteredRes)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(*apperror.CErr)
	return ret0, ret1, ret2
}

// Register indicates an expected call of Register.
func (mr *MockAuthUsecaseMockRecorder) Register(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthUsecase)(nil).Register), ctx, req)
}
