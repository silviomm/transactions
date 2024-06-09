package services

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"pismo-challenge/database/repositories"
	"pismo-challenge/mocks"
	"pismo-challenge/models/account"
	"pismo-challenge/models/transaction"
	"testing"
)

func TestValidateTransactionDto_InvalidOperation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountRepo := mocks.NewMockAccountsRepository(ctrl)
	repositories.Accounts = mockAccountRepo
	mockTransactionsRepo := mocks.NewMockTransactionsRepository(ctrl)
	repositories.Transactions = mockTransactionsRepo

	// Test case 3: Invalid OperationTypeId
	invalidOperationDTO := transaction.CreateTransactionDto{
		AccountId:       1,
		OperationTypeId: 999, // Assuming this operation does not exist
	}
	mockAccountRepo.EXPECT().GetAccount(gomock.Any()).Return(&account.Account{}).AnyTimes()
	mockTransactionsRepo.EXPECT().GetOperation(gomock.Any()).Return(nil).AnyTimes()
	err := ValidateTransactionDto(invalidOperationDTO)
	assert.EqualError(t, err, fmt.Sprintf("Operation '%d' does not exists", invalidOperationDTO.OperationTypeId))
}

func TestValidateTransactionDto_InvalidAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountRepo := mocks.NewMockAccountsRepository(ctrl)
	repositories.Accounts = mockAccountRepo
	mockTransactionsRepo := mocks.NewMockTransactionsRepository(ctrl)
	repositories.Transactions = mockTransactionsRepo

	invalidAccountDTO := transaction.CreateTransactionDto{
		AccountId:       999, // Assuming this account does not exist
		OperationTypeId: 1,
	}
	mockAccountRepo.EXPECT().GetAccount(gomock.Any()).Return(nil).AnyTimes()
	mockTransactionsRepo.EXPECT().GetOperation(gomock.Any()).Return(&transaction.Operation{}).AnyTimes()
	err := ValidateTransactionDto(invalidAccountDTO)
	assert.EqualError(t, err, fmt.Sprintf("Account '%d' does not exists", invalidAccountDTO.AccountId))
}

func TestValidateTransactionDto_Valid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountRepo := mocks.NewMockAccountsRepository(ctrl)
	repositories.Accounts = mockAccountRepo
	mockTransactionsRepo := mocks.NewMockTransactionsRepository(ctrl)
	repositories.Transactions = mockTransactionsRepo

	validDTO := transaction.CreateTransactionDto{
		AccountId:       1,
		OperationTypeId: 1,
	}
	mockAccountRepo.EXPECT().GetAccount(gomock.Any()).Return(&account.Account{}).AnyTimes()
	mockTransactionsRepo.EXPECT().GetOperation(gomock.Any()).Return(&transaction.Operation{}).AnyTimes()
	err := ValidateTransactionDto(validDTO)
	assert.NoError(t, err)
}

func TestOperationDoesNotExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTransactionsRepository(ctrl)
	repositories.Transactions = mockRepo

	mockRepo.EXPECT().GetOperation(gomock.Any()).Return(nil)
	exists := OperationExists(transaction.OperationType(1))
	assert.False(t, exists)
}

func TestOperationExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTransactionsRepository(ctrl)
	repositories.Transactions = mockRepo

	mockRepo.EXPECT().GetOperation(transaction.OperationType(2)).Return(&transaction.Operation{})
	exists := OperationExists(transaction.OperationType(2))
	assert.True(t, exists)
}

func TestGetAmountByTransactionType(t *testing.T) {
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
		result := GetAmountByOperationType(tt.amount, tt.operationType)
		if result != tt.expected {
			t.Errorf("GetAmountByTransactionType(%v, %v) = %v; expected %v", tt.amount, tt.operationType, result, tt.expected)
		}
	}
}
