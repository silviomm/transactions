package repositories

import "gorm.io/gorm"

var (
	DB           gorm.DB
	Transactions TransactionsRepository
	Accounts     AccountsRepository
)

func InitRepositories(db gorm.DB) {
	DB = db
	Transactions = &DefaultTransactionsRepository{}
	Transactions.InitializeTransactionsRepository(db)

	Accounts = &DefaultAccountsRepository{}
	Accounts.InitializeAccountsRepository(db)
}
