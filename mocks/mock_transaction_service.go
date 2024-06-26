// Code generated by MockGen. DO NOT EDIT.
// Source: services/transactionsService.go

// Package mocks is a generated GoMock package.
package mocks

import (
	transaction "pismo-challenge/models/transaction"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTransactionService is a mock of TransactionService interface.
type MockTransactionService struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionServiceMockRecorder
}

// MockTransactionServiceMockRecorder is the mock recorder for MockTransactionService.
type MockTransactionServiceMockRecorder struct {
	mock *MockTransactionService
}

// NewMockTransactionService creates a new mock instance.
func NewMockTransactionService(ctrl *gomock.Controller) *MockTransactionService {
	mock := &MockTransactionService{ctrl: ctrl}
	mock.recorder = &MockTransactionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionService) EXPECT() *MockTransactionServiceMockRecorder {
	return m.recorder
}

// CreateTransaction mocks base method.
func (m *MockTransactionService) CreateTransaction(dto *transaction.CreateTransactionDto) (transaction.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", dto)
	ret0, _ := ret[0].(transaction.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockTransactionServiceMockRecorder) CreateTransaction(dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockTransactionService)(nil).CreateTransaction), dto)
}

// Discharge mocks base method.
func (m *MockTransactionService) Discharge(t transaction.Transaction) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Discharge", t)
}

// Discharge indicates an expected call of Discharge.
func (mr *MockTransactionServiceMockRecorder) Discharge(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discharge", reflect.TypeOf((*MockTransactionService)(nil).Discharge), t)
}

// GetAmountByOperationType mocks base method.
func (m *MockTransactionService) GetAmountByOperationType(amount float64, operationType transaction.OperationType) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAmountByOperationType", amount, operationType)
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetAmountByOperationType indicates an expected call of GetAmountByOperationType.
func (mr *MockTransactionServiceMockRecorder) GetAmountByOperationType(amount, operationType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAmountByOperationType", reflect.TypeOf((*MockTransactionService)(nil).GetAmountByOperationType), amount, operationType)
}

// OperationExists mocks base method.
func (m *MockTransactionService) OperationExists(operationType transaction.OperationType) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OperationExists", operationType)
	ret0, _ := ret[0].(bool)
	return ret0
}

// OperationExists indicates an expected call of OperationExists.
func (mr *MockTransactionServiceMockRecorder) OperationExists(operationType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OperationExists", reflect.TypeOf((*MockTransactionService)(nil).OperationExists), operationType)
}

// ShouldDischarge mocks base method.
func (m *MockTransactionService) ShouldDischarge(operationType transaction.OperationType) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldDischarge", operationType)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ShouldDischarge indicates an expected call of ShouldDischarge.
func (mr *MockTransactionServiceMockRecorder) ShouldDischarge(operationType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldDischarge", reflect.TypeOf((*MockTransactionService)(nil).ShouldDischarge), operationType)
}

// ValidateTransactionDto mocks base method.
func (m *MockTransactionService) ValidateTransactionDto(dto transaction.CreateTransactionDto) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateTransactionDto", dto)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateTransactionDto indicates an expected call of ValidateTransactionDto.
func (mr *MockTransactionServiceMockRecorder) ValidateTransactionDto(dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateTransactionDto", reflect.TypeOf((*MockTransactionService)(nil).ValidateTransactionDto), dto)
}
