package account

import (
	"pismo-challenge/models/transaction"
)

type Account struct {
	Id             int                     `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:Id;unique"`
	DocumentNumber string                  `json:"document_number" gorm:"column:DocumentNumber;unique"`
	Transactions   transaction.Transaction `json:"transactions" gorm:"foreignKey:AccountId"`
}
