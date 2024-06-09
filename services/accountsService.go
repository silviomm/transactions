package services

import (
	"pismo-challenge/database"
	"pismo-challenge/database/repositories"
	"pismo-challenge/models/account"
)

// AccountExists todo: add cache layer
func AccountExists(accountId int) bool {
	ac := repositories.Accounts.GetAccount(accountId)
	if ac == nil {
		return false
	}
	return true
}

func ExistsByDocumentNumber(documentNumber string) bool {
	var count int64 = 0
	database.DB.Model(&account.Account{}).Where("accounts.\"DocumentNumber\" = ?", documentNumber).Count(&count)
	if count == 0 {
		return false
	}
	return true
}
