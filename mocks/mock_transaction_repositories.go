// Code generated by MockGen. DO NOT EDIT.
// Source: database/repositories/transactionsRepository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	transaction "pismo-challenge/models/transaction"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockTransactionsRepository is a mock of TransactionsRepository interface.
type MockTransactionsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionsRepositoryMockRecorder
}

// MockTransactionsRepositoryMockRecorder is the mock recorder for MockTransactionsRepository.
type MockTransactionsRepositoryMockRecorder struct {
	mock *MockTransactionsRepository
}

// NewMockTransactionsRepository creates a new mock instance.
func NewMockTransactionsRepository(ctrl *gomock.Controller) *MockTransactionsRepository {
	mock := &MockTransactionsRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionsRepository) EXPECT() *MockTransactionsRepositoryMockRecorder {
	return m.recorder
}

// GetOperation mocks base method.
func (m *MockTransactionsRepository) GetOperation(operationType transaction.OperationType) *transaction.Operation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperation", operationType)
	ret0, _ := ret[0].(*transaction.Operation)
	return ret0
}

// GetOperation indicates an expected call of GetOperation.
func (mr *MockTransactionsRepositoryMockRecorder) GetOperation(operationType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperation", reflect.TypeOf((*MockTransactionsRepository)(nil).GetOperation), operationType)
}

// InitializeTransactionsRepository mocks base method.
func (m *MockTransactionsRepository) InitializeTransactionsRepository(db gorm.DB) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitializeTransactionsRepository", db)
}

// InitializeTransactionsRepository indicates an expected call of InitializeTransactionsRepository.
func (mr *MockTransactionsRepositoryMockRecorder) InitializeTransactionsRepository(db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeTransactionsRepository", reflect.TypeOf((*MockTransactionsRepository)(nil).InitializeTransactionsRepository), db)
}
