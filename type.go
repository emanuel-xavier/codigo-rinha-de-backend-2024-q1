package main

import (
	"log"
	"strings"
	"time"
)

type CreateTransactionDto struct {
	Value       int    `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

func (ctDto *CreateTransactionDto) validate() bool {
	log.Println(*ctDto)
	if ctDto.Value <= 0 {
		log.Println("Value <= 0")
		return false
	}

	if strings.ToLower(ctDto.Type) != "c" && strings.ToLower(ctDto.Type) != "d" {
		log.Println("Wrong type")
		return false
	}

	if decLen := len(ctDto.Description); decLen < 1 || decLen > 10 {
		log.Println("Description size")
		return false
	}

	return true
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
	Balance          BalanceDto       `json:"saldo"`
	LastTransactions []TransactionDto `json:"ultimas_transacoes"`
}
