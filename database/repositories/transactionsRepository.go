package repositories

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"pismo-challenge/models/transaction"
)

type TransactionsRepository interface {
	GetOperation(operationType transaction.OperationType) *transaction.Operation
	InitializeTransactionsRepository(db gorm.DB)
	GetTransactionsToDischarge(accountId int) []transaction.Transaction
	UpdateBalance(trId int, balance float64) error
	InsertTransaction(transaction *transaction.Transaction) error
}

type DefaultTransactionsRepository struct{}

func (d DefaultTransactionsRepository) InsertTransaction(transaction *transaction.Transaction) error {
	result := DB.Create(&transaction)
	if result.Error != nil {
		return errors.New("failed to create transaction")
	}
	return nil
}

func (d DefaultTransactionsRepository) UpdateBalance(trId int, balance float64) error {
	result := DB.Model(&transaction.Transaction{}).Where("\"Id\" = ?", trId).Update("\"Balance\"", balance)
	return result.Error
}

func (d DefaultTransactionsRepository) GetTransactionsToDischarge(accountId int) []transaction.Transaction {
	transactionsToPayQuery := DB.Where("\"Balance\" < 0")
	transactionsToPayQuery = transactionsToPayQuery.Where("\"AccountId\" = ?", accountId)
	transactionsToPayQuery = transactionsToPayQuery.Where("\"OperationType\" = ?", 1)
	transactionsToPayQuery = transactionsToPayQuery.Or("\"OperationType\" = ?", 2)
	transactionsToPayQuery = transactionsToPayQuery.Or("\"OperationType\" = ?", 3)
	transactionsToPayQuery = transactionsToPayQuery.Order("\"EventDate\"")
	var ts []transaction.Transaction
	transactionsToPayQuery.Find(&ts)
	return ts
}

func (d DefaultTransactionsRepository) InitializeTransactionsRepository(db gorm.DB) {
	err := db.AutoMigrate(&transaction.Operation{})
	if err != nil {
		log.Panic("Error migrating Operations table", err)
	}
	seedData(db)
	err = db.AutoMigrate(&transaction.Transaction{})
	db.Model(&transaction.Transaction{}).Where("\"Balance\" is NULL").Update("\"Balance\"", 0)
	if err != nil {
		log.Panic("Error migrating Transactions table", err)
	}
}

func seedData(db gorm.DB) {
	seedData := []transaction.Operation{
		{ID: transaction.BuyInCash, Name: "BuyInCash"},
		{ID: transaction.BuyInInstallments, Name: "BuyInInstallments"},
		{ID: transaction.Withdraw, Name: "Withdraw"},
		{ID: transaction.Payment, Name: "Payment"},
	}

	for _, op := range seedData {
		if err := db.FirstOrCreate(&op).Error; err != nil {
			log.Fatalf("Failed to seed Operation data: %v", err)
		}
	}
}

func (d DefaultTransactionsRepository) GetOperation(operationType transaction.OperationType) *transaction.Operation {
	var op transaction.Operation
	err := DB.First(&op, operationType)
	if err.Error != nil {
		return nil
	}
	return &op
}
