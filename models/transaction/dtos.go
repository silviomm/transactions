package transaction

import "time"

type CreateTransactionDto struct {
	Amount          float64       `json:"amount"`
	AccountId       int           `json:"account_id"`
	OperationTypeId OperationType `json:"operation_type_id"`
}

type GetTransactionsDto struct {
	AccountId       int       `form:"account_id"`
	OperationTypeId int       `form:"operation_type_id"`
	After           time.Time `form:"after"`
	Before          time.Time `form:"before"`
}
