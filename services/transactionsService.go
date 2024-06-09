package services

import (
	"errors"
	"fmt"
	"math"
	"pismo-challenge/database/repositories"
	"pismo-challenge/models/transaction"
)

func ValidateTransactionDto(dto transaction.CreateTransactionDto) error {
	if !AccountExists(dto.AccountId) {
		return errors.New(fmt.Sprintf("Account '%d' does not exists", dto.AccountId))
	}
	if !OperationExists(dto.OperationTypeId) {
		return errors.New(fmt.Sprintf("Operation '%d' does not exists", dto.OperationTypeId))
	}
	return nil
}

// OperationExists todo: add cache layer
func OperationExists(operationType transaction.OperationType) bool {
	op := repositories.Transactions.GetOperation(operationType)
	if op == nil {
		return false
	}
	return true
}

func GetAmountByOperationType(amount float64, operationType transaction.OperationType) float64 {
	if operationType == transaction.Payment {
		return math.Abs(amount)
	}
	return -math.Abs(amount)
}
