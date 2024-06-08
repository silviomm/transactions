package transaction

type OperationType int

const (
	RESERVED OperationType = iota
	PAY_IN_CASH
	INSTALLMENTS
	WITHDRAW
	PAYMENT
)
