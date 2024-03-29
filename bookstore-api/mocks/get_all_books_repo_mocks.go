// Code generated by MockGen. DO NOT EDIT.
// Source: get_all_books_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	model "bookstore-api/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBookStoreRepository is a mock of BookStoreRepository interface.
type MockBookStoreRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBookStoreRepositoryMockRecorder
}

// MockBookStoreRepositoryMockRecorder is the mock recorder for MockBookStoreRepository.
type MockBookStoreRepositoryMockRecorder struct {
	mock *MockBookStoreRepository
}

// NewMockBookStoreRepository creates a new mock instance.
func NewMockBookStoreRepository(ctrl *gomock.Controller) *MockBookStoreRepository {
	mock := &MockBookStoreRepository{ctrl: ctrl}
	mock.recorder = &MockBookStoreRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookStoreRepository) EXPECT() *MockBookStoreRepositoryMockRecorder {
	return m.recorder
}

// GetAllBooks mocks base method.
func (m *MockBookStoreRepository) GetAllBooks(ctx context.Context, searchText string, pageSize, pageNum int) ([]model.Books, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBooks", ctx, searchText, pageSize, pageNum)
	ret0, _ := ret[0].([]model.Books)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBooks indicates an expected call of GetAllBooks.
func (mr *MockBookStoreRepositoryMockRecorder) GetAllBooks(ctx, searchText, pageSize, pageNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBooks", reflect.TypeOf((*MockBookStoreRepository)(nil).GetAllBooks), ctx, searchText, pageSize, pageNum)
}

// GetBookById mocks base method.
func (m *MockBookStoreRepository) GetBookById(ctx context.Context, bookId string) (model.Books, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookById", ctx, bookId)
	ret0, _ := ret[0].(model.Books)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookById indicates an expected call of GetBookById.
func (mr *MockBookStoreRepositoryMockRecorder) GetBookById(ctx, bookId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookById", reflect.TypeOf((*MockBookStoreRepository)(nil).GetBookById), ctx, bookId)
}
