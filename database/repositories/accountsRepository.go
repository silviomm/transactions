package repositories

import (
	"gorm.io/gorm"
	"log"
	"pismo-challenge/models/account"
)

type AccountsRepository interface {
	GetAccount(accountId int) *account.Account
	InitializeAccountsRepository(db gorm.DB)
}

type DefaultAccountsRepository struct{}

func (d DefaultAccountsRepository) InitializeAccountsRepository(db gorm.DB) {
	err := db.AutoMigrate(&account.Account{})
	if err != nil {
		log.Panic("Error migrating Accounts table", err)
	}
}

func (d DefaultAccountsRepository) GetAccount(accountId int) *account.Account {
	var ac account.Account
	err := DB.First(&ac, accountId)
	if err.Error != nil {
		return nil
	}
	return &ac
}
