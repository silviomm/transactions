package services

var (
	Transactions TransactionService
	Accounts     AccountService
)

func InitServices() {
	Accounts = &DefaultAccountService{}
	Transactions = &DefaultTransactionService{}
}
