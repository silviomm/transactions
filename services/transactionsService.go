package services

import (
	"errors"
	"fmt"
	"math"
	"pismo-challenge/database/repositories"
	"pismo-challenge/models/transaction"
	"time"
)

type TransactionService interface {
	ValidateTransactionDto(dto transaction.CreateTransactionDto) error
	OperationExists(operationType transaction.OperationType) bool
	GetAmountByOperationType(amount float64, operationType transaction.OperationType) float64
	Discharge(t transaction.Transaction)
	ShouldDischarge(operationType transaction.OperationType) bool
	CreateTransaction(dto *transaction.CreateTransactionDto) (transaction.Transaction, error)
}

type DefaultTransactionService struct{}

func (s DefaultTransactionService) CreateTransaction(dto *transaction.CreateTransactionDto) (transaction.Transaction, error) {
	value := s.GetAmountByOperationType(dto.Amount, dto.OperationTypeId)
	t := transaction.Transaction{
		AccountId:     dto.AccountId,
		OperationType: dto.OperationTypeId,
		Amount:        value,
		EventDate:     time.Now(),
		Balance:       value,
	}
	err := repositories.Transactions.InsertTransaction(&t)
	if err != nil {
		return transaction.Transaction{}, err
	}
	return t, nil
}

func (s DefaultTransactionService) ShouldDischarge(operationType transaction.OperationType) bool {
	if operationType == transaction.Payment {
		return true
	}
	return false
}

func (s DefaultTransactionService) Discharge(t transaction.Transaction) {
	ts := repositories.Transactions.GetTransactionsToDischarge(t.AccountId)
	total := t.Amount
	var err error = nil
	for _, tr := range ts {
		if total > 0 {
			if total < math.Abs(tr.Balance) {
				err = repositories.Transactions.UpdateBalance(tr.Id, total+tr.Balance)
				if err != nil {
					break
				}
				total = 0
			}
			if total >= math.Abs(tr.Balance) {
				total = total + tr.Balance
				err = repositories.Transactions.UpdateBalance(tr.Id, 0)
				if err != nil {
					break
				}
			}
		}
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = repositories.Transactions.UpdateBalance(t.Id, total)
}

func (s DefaultTransactionService) ValidateTransactionDto(dto transaction.CreateTransactionDto) error {
	if !Accounts.AccountExists(dto.AccountId) {
		return errors.New(fmt.Sprintf("Account '%d' does not exists", dto.AccountId))
	}
	if !s.OperationExists(dto.OperationTypeId) {
		return errors.New(fmt.Sprintf("Operation '%d' does not exists", dto.OperationTypeId))
	}
	return nil
}

// OperationExists todo: add cache layer
func (s DefaultTransactionService) OperationExists(operationType transaction.OperationType) bool {
	op := repositories.Transactions.GetOperation(operationType)
	if op == nil {
		return false
	}
	return true
}

func (s DefaultTransactionService) GetAmountByOperationType(amount float64, operationType transaction.OperationType) float64 {
	if operationType == transaction.Payment {
		return math.Abs(amount)
	}
	return -math.Abs(amount)
}
