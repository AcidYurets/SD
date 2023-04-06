// Code generated by MockGen. DO NOT EDIT.
// Source: calend/internal/modules/domain/tag/service (interfaces: ITagRepo)

// Package service is a generated GoMock package.
package service

import (
	dto "calend/internal/modules/domain/tag/dto"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockITagRepo is a mock of ITagRepo interface.
type MockITagRepo struct {
	ctrl     *gomock.Controller
	recorder *MockITagRepoMockRecorder
}

// MockITagRepoMockRecorder is the mock recorder for MockITagRepo.
type MockITagRepoMockRecorder struct {
	mock *MockITagRepo
}

// NewMockITagRepo creates a new mock instance.
func NewMockITagRepo(ctrl *gomock.Controller) *MockITagRepo {
	mock := &MockITagRepo{ctrl: ctrl}
	mock.recorder = &MockITagRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITagRepo) EXPECT() *MockITagRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockITagRepo) Create(arg0 context.Context, arg1 *dto.CreateTag) (*dto.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*dto.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockITagRepoMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockITagRepo)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockITagRepo) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockITagRepoMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockITagRepo)(nil).Delete), arg0, arg1)
}

// GetByUuid mocks base method.
func (m *MockITagRepo) GetByUuid(arg0 context.Context, arg1 string) (*dto.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUuid", arg0, arg1)
	ret0, _ := ret[0].(*dto.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUuid indicates an expected call of GetByUuid.
func (mr *MockITagRepoMockRecorder) GetByUuid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUuid", reflect.TypeOf((*MockITagRepo)(nil).GetByUuid), arg0, arg1)
}

// List mocks base method.
func (m *MockITagRepo) List(arg0 context.Context) (dto.Tags, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(dto.Tags)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockITagRepoMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockITagRepo)(nil).List), arg0)
}

// Update mocks base method.
func (m *MockITagRepo) Update(arg0 context.Context, arg1 string, arg2 *dto.UpdateTag) (*dto.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*dto.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockITagRepoMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockITagRepo)(nil).Update), arg0, arg1, arg2)
}