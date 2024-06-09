package transaction

import "time"

type CreateTransactionDto struct {
	Amount          float64       `json:"amount" binding:"required"`
	AccountId       int           `json:"account_id" binding:"required,gt=0"`
	OperationTypeId OperationType `json:"operation_type_id" binding:"required"`
}

type GetTransactionsDto struct {
	AccountId       int       `form:"account_id"`
	OperationTypeId int       `form:"operation_type_id"`
	After           time.Time `form:"after"`
	Before          time.Time `form:"before"`
}
