package dto

import "errors"

type TransactionRequest struct {
	Value       int    `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

type TransactionResponse struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}

func (tr *TransactionRequest) Validate() error {
	if tr.Value <= 0 {
		return errors.New("the value must be greater than 0")
	}

	if len(tr.Description) > 10 || tr.Description == "" {
		return errors.New("the description must not be empty or have more than 10 characters")
	}

	if tr.Type != "c" && tr.Type != "d" {
		return errors.New("the type value must be 'c' or 'd'")
	}

	return nil
}
