package entity

import (
	"time"
)

type StatementBalance struct {
	Amount int       `json:"total"`
	Date   time.Time `json:"data_extrato"`
	Limit  int       `json:"limite"`
}

type StatementLastTransactions []Transaction

type Statement struct {
	Balance             StatementBalance          `json:"saldo"`
	LastTenTransactions StatementLastTransactions `json:"ultimas_transacoes"`
}
