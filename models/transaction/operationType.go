package transaction

type OperationType int

type Operation struct {
	ID   OperationType `gorm:"primaryKey;column:Id"`
	Name string        `gorm:"column:Name;unique"`
}

const (
	BuyInCash         OperationType = 1
	BuyInInstallments OperationType = 2
	Withdraw          OperationType = 3
	Payment           OperationType = 4
)
