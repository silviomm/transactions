package account

type CreateAccountDto struct {
	DocumentNumber string `json:"document_number" binding:"required,max=32"`
}

type GetAccountDto struct {
	Id             int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}
