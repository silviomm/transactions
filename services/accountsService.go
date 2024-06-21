package services

import (
	"pismo-challenge/database"
	"pismo-challenge/database/repositories"
	"pismo-challenge/models/account"
)

type AccountService interface {
	AccountExists(accountId int) bool
	ExistsByDocumentNumber(documentNumber string) bool
}

type DefaultAccountService struct{}

// AccountExists todo: add cache layer
func (s DefaultAccountService) AccountExists(accountId int) bool {
	ac := repositories.Accounts.GetAccount(accountId)
	if ac == nil {
		return false
	}
	return true
}

func (s DefaultAccountService) ExistsByDocumentNumber(documentNumber string) bool {
	var count int64 = 0
	database.DB.Model(&account.Account{}).Where("accounts.\"DocumentNumber\" = ?", documentNumber).Count(&count)
	if count == 0 {
		return false
	}
	return true
}
