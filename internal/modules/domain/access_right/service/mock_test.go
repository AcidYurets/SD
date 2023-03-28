// Code generated by MockGen. DO NOT EDIT.
// Source: calend/internal/modules/domain/access_right/service (interfaces: IAccessRightRepo)

// Package service is a generated GoMock package.
package service

import (
	dto "calend/internal/modules/domain/access_right/dto"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIAccessRightRepo is a mock of IAccessRightRepo interface.
type MockIAccessRightRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIAccessRightRepoMockRecorder
}

// MockIAccessRightRepoMockRecorder is the mock recorder for MockIAccessRightRepo.
type MockIAccessRightRepoMockRecorder struct {
	mock *MockIAccessRightRepo
}

// NewMockIAccessRightRepo creates a new mock instance.
func NewMockIAccessRightRepo(ctrl *gomock.Controller) *MockIAccessRightRepo {
	mock := &MockIAccessRightRepo{ctrl: ctrl}
	mock.recorder = &MockIAccessRightRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAccessRightRepo) EXPECT() *MockIAccessRightRepoMockRecorder {
	return m.recorder
}

// GetByCode mocks base method.
func (m *MockIAccessRightRepo) GetByCode(arg0 context.Context, arg1 string) (*dto.AccessRight, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCode", arg0, arg1)
	ret0, _ := ret[0].(*dto.AccessRight)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCode indicates an expected call of GetByCode.
func (mr *MockIAccessRightRepoMockRecorder) GetByCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCode", reflect.TypeOf((*MockIAccessRightRepo)(nil).GetByCode), arg0, arg1)
}

// List mocks base method.
func (m *MockIAccessRightRepo) List(arg0 context.Context) (dto.AccessRights, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(dto.AccessRights)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockIAccessRightRepoMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIAccessRightRepo)(nil).List), arg0)
}
