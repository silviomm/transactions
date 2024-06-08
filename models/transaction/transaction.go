package transaction

import "time"

type Transaction struct {
	Id            int           `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:Id;unique"`
	AccountId     int           `json:"account_id" gorm:"column:AccountId"`
	OperationType OperationType `json:"operation_type_id" gorm:"column:OperationType"`
	Amount        float64       `json:"amount"`
	EventDate     time.Time     `json:"event_date" gorm:"column:EventDate"`
}
