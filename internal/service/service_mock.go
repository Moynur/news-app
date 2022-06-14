// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/moynur/news-app/internal/models"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetArticles mocks base method.
func (m *MockService) GetArticles(request models.GetArticlesRequest) (models.GetArticlesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticles", request)
	ret0, _ := ret[0].(models.GetArticlesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticles indicates an expected call of GetArticles.
func (mr *MockServiceMockRecorder) GetArticles(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticles", reflect.TypeOf((*MockService)(nil).GetArticles), request)
}
