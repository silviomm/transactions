package services

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"pismo-challenge/models/account"
	"pismo-challenge/models/transaction"
	"testing"
)

func TestShouldNotDischargeOperationDifferentThanPayment(t *testing.T) {
	ctrl := gomock.NewController(t)
	InitMocks(ctrl)
	defer ctrl.Finish()

	tests := []struct {
		operationType transaction.OperationType
	}{
		{transaction.BuyInCash},
		{transaction.BuyInInstallments},
		{transaction.Withdraw},
	}
	for _, tt := range tests {
		result := Transactions.ShouldDischarge(tt.operationType)
		assert.False(t, result)
	}
}

func TestShouldDischargeOnlyPaymentOperation(t *testing.T) {
	ctrl := gomock.NewController(t)
	InitMocks(ctrl)
	defer ctrl.Finish()

	result := Transactions.ShouldDischarge(transaction.Payment)
	assert.True(t, result)
}

func TestValidateTransactionDto_InvalidOperation(t *testing.T) {
	ctrl := gomock.NewController(t)
	InitMocks(ctrl)
	defer ctrl.Finish()

	invalidOperationDTO := transaction.CreateTransactionDto{
		AccountId:       1,
		OperationTypeId: 999, // Assuming this operation does not exist
	}
	mockAccountsRepository.EXPECT().GetAccount(gomock.Any()).Return(&account.Account{})
	mockTransactionsRepository.EXPECT().GetOperation(gomock.Any()).Return(nil).Times(1)
	err := Transactions.ValidateTransactionDto(invalidOperationDTO)
	assert.EqualError(t, err, fmt.Sprintf("Operation '%d' does not exists", invalidOperationDTO.OperationTypeId))
}

func TestValidateTransactionDto_InvalidAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	InitMocks(ctrl)
	defer ctrl.Finish()

	invalidAccountDTO := transaction.CreateTransactionDto{
		AccountId:       999, // Assuming this account does not exist
		OperationTypeId: 1,
	}
	mockAccountsRepository.EXPECT().GetAccount(gomock.Any()).Return(nil)
	mockTransactionsRepository.EXPECT().GetOperation(gomock.Any()).Return(&transaction.Operation{}).AnyTimes()
	err := Transactions.ValidateTransactionDto(invalidAccountDTO)
	assert.EqualError(t, err, fmt.Sprintf("Account '%d' does not exists", invalidAccountDTO.AccountId))
}

func TestValidateTransactionDto_Valid(t *testing.T) {
	ctrl := gomock.NewController(t)
	InitMocks(ctrl)
	defer ctrl.Finish()

	validDTO := transaction.CreateTransactionDto{
		AccountId:       1,
		OperationTypeId: 1,
	}
	mockAccountsRepository.EXPECT().GetAccount(gomock.Any()).Return(&account.Account{})
	mockTransactionsRepository.EXPECT().GetOperation(gomock.Any()).Return(&transaction.Operation{}).Times(1)
	err := Transactions.ValidateTransactionDto(validDTO)
	assert.NoError(t, err)
}

func TestOperationDoesNotExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	InitMocks(ctrl)
	defer ctrl.Finish()

	mockTransactionsRepository.EXPECT().GetOperation(gomock.Any()).Return(nil).Times(1)
	exists := Transactions.OperationExists(transaction.OperationType(1))
	assert.False(t, exists)
}

func TestOperationExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	InitMocks(ctrl)
	defer ctrl.Finish()

	mockTransactionsRepository.EXPECT().GetOperation(transaction.OperationType(2)).Return(&transaction.Operation{}).Times(1)
	exists := Transactions.OperationExists(transaction.OperationType(2))
	assert.True(t, exists)
}

func TestGetAmountByTransactionType(t *testing.T) {
	ctrl := gomock.NewController(t)
	InitMocks(ctrl)
	defer ctrl.Finish()

	tests := []struct {
		amount        float64
		operationType transaction.OperationType
		expected      float64
	}{
		{10, transaction.Payment, 10},
		{-10, transaction.Payment, 10},
		{10, transaction.Withdraw, -10},
		{-10, transaction.Withdraw, -10},
		{10, transaction.BuyInCash, -10},
		{-10, transaction.BuyInCash, -10},
		{10, transaction.BuyInInstallments, -10},
		{-10, transaction.BuyInInstallments, -10},
	}

	for _, tt := range tests {
		result := Transactions.GetAmountByOperationType(tt.amount, tt.operationType)
		if result != tt.expected {
			t.Errorf("GetAmountByTransactionType(%v, %v) = %v; expected %v", tt.amount, tt.operationType, result, tt.expected)
		}
	}
}
