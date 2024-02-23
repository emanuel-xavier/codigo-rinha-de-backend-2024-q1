package entity

import (
	"errors"
	"time"
)

type Transaction struct {
	Id           int       // `json:"id,omitempty"`
	Type         string    `json:"tipo"`
	Description  string    `json:"descricao"`
	Value        int       `json:"valor,omitempty"`
	Accomplished time.Time `json:"realizado_em,omitempty"`
	ClientId     int       // `json:"id_do_cliente,omitempty"`
}

func (tr *Transaction) Validate() error {
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
