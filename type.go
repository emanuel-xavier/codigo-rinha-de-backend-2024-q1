package main

import "time"

type CreateTransactionDto struct {
	Value       int    `json:"valor"`
	T           rune   `json:"tipo"`
	Description string `json:"descricao"`
}

type TransactionResponseDto struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}

type TransactionDto struct {
	Value        int       `json:"valor"`
	Type         string    `json:"tipo"`
	Description  string    `json:"descricao"`
	Accomplished time.Time `json:"realizada_em"`
}

type BalanceDto struct {
	Total       int    `json:"total"`
	BalanceDate string `json:"data_extrato"`
	Limit       int    `json:"limite"`
}

type StatementResponseDto struct {
	B                BalanceDto       `json:"saldo"`
	LastTransactions []TransactionDto `json:"ultimas_transacoes"`
}
