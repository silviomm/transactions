package repositories

import (
	"gorm.io/gorm"
	"log"
	"pismo-challenge/models/transaction"
)

type TransactionsRepository interface {
	GetOperation(operationType transaction.OperationType) *transaction.Operation
	InitializeTransactionsRepository(db gorm.DB)
}

type DefaultTransactionsRepository struct{}

func (d DefaultTransactionsRepository) InitializeTransactionsRepository(db gorm.DB) {
	err := db.AutoMigrate(&transaction.Operation{})
	if err != nil {
		log.Panic("Error migrating Operations table", err)
	}
	seedData(db)
	err = db.AutoMigrate(&transaction.Transaction{})
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
