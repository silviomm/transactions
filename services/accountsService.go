package services

import (
	"pismo-challenge/database"
	"pismo-challenge/models/account"
)

// todo: add cache layer
func Exists(accountId int) bool {
	var count int64 = 0
	database.DB.Model(&account.Account{}).Where("accounts.\"Id\" = ?", accountId).Count(&count)
	if count == 0 {
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
