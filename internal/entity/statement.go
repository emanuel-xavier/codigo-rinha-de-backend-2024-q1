package entity

import "time"

type StatementBalance struct {
	Amount int       `json:"total"`
	Date   time.Time `json:"data_extrato"`
	Limit  int       `json:"limite"`
}

type StatementTransaction struct {
	Value        int       `json:"valor"`
	Type         string    `json:"tipo"`
	Description  string    `json:"descricao"`
	Accomplished time.Time `json:"realizado_em"`
}

type StatementLastTransactions []StatementTransaction

type Statement struct {
	Balance             StatementBalance          `json:"saldo"`
	LastTenTransactions StatementLastTransactions `json:"ultimas_transacoes"`
}
