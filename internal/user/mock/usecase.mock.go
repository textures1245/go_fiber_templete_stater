// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/phakh/go/src/github.com/textures1245/go-template/internal/user/usecase.go
//
// Generated by this command:
//
//	mockgen -source=/Users/phakh/go/src/github.com/textures1245/go-template/internal/user/usecase.go
//

// Package mock_user is a generated GoMock package.
package mock_user

import (
	context "context"
	reflect "reflect"

	dtos "github.com/textures1245/go-template/internal/user/dtos"
	entities "github.com/textures1245/go-template/internal/user/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockUserUsecase is a mock of UserUsecase interface.
type MockUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaseMockRecorder
}

// MockUserUsecaseMockRecorder is the mock recorder for MockUserUsecase.
type MockUserUsecaseMockRecorder struct {
	mock *MockUserUsecase
}

// NewMockUserUsecase creates a new mock instance.
func NewMockUserUsecase(ctrl *gomock.Controller) *MockUserUsecase {
	mock := &MockUserUsecase{ctrl: ctrl}
	mock.recorder = &MockUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUsecase) EXPECT() *MockUserUsecaseMockRecorder {
	return m.recorder
}

// OnFetchUserById mocks base method.
func (m *MockUserUsecase) OnFetchUserById(ctx context.Context, userId int64) (*dtos.UserDetailRespond, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnFetchUserById", ctx, userId)
	ret0, _ := ret[0].(*dtos.UserDetailRespond)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// OnFetchUserById indicates an expected call of OnFetchUserById.
func (mr *MockUserUsecaseMockRecorder) OnFetchUserById(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnFetchUserById", reflect.TypeOf((*MockUserUsecase)(nil).OnFetchUserById), ctx, userId)
}

// OnFetchUsers mocks base method.
func (m *MockUserUsecase) OnFetchUsers(ctx context.Context) ([]*dtos.UserDetailRespond, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnFetchUsers", ctx)
	ret0, _ := ret[0].([]*dtos.UserDetailRespond)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// OnFetchUsers indicates an expected call of OnFetchUsers.
func (mr *MockUserUsecaseMockRecorder) OnFetchUsers(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnFetchUsers", reflect.TypeOf((*MockUserUsecase)(nil).OnFetchUsers), ctx)
}

// OnUpdateUserById mocks base method.
func (m *MockUserUsecase) OnUpdateUserById(ctx context.Context, userId int64, req *entities.UserUpdateReq) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnUpdateUserById", ctx, userId, req)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OnUpdateUserById indicates an expected call of OnUpdateUserById.
func (mr *MockUserUsecaseMockRecorder) OnUpdateUserById(ctx, userId, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnUpdateUserById", reflect.TypeOf((*MockUserUsecase)(nil).OnUpdateUserById), ctx, userId, req)
}

// OnUserLogin mocks base method.
func (m *MockUserUsecase) OnUserLogin(ctx context.Context, req *entities.UserLoginReq) (*dtos.UserLoginResponse, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnUserLogin", ctx, req)
	ret0, _ := ret[0].(*dtos.UserLoginResponse)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// OnUserLogin indicates an expected call of OnUserLogin.
func (mr *MockUserUsecaseMockRecorder) OnUserLogin(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnUserLogin", reflect.TypeOf((*MockUserUsecase)(nil).OnUserLogin), ctx, req)
}

// UserDeleted mocks base method.
func (m *MockUserUsecase) UserDeleted(ctx context.Context, userId int64) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserDeleted", ctx, userId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserDeleted indicates an expected call of UserDeleted.
func (mr *MockUserUsecaseMockRecorder) UserDeleted(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserDeleted", reflect.TypeOf((*MockUserUsecase)(nil).UserDeleted), ctx, userId)
}