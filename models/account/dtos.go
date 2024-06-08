package account

type CreateAccountDto struct {
	DocumentNumber string `json:"document_number"`
}

type GetAccountDto struct {
	Id             int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}
