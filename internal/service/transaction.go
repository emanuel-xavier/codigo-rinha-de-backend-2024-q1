package service

import (
	"context"

	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/internal/entity"
	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/internal/repository"
)

type TransactionService struct {
	cRepo repository.ClientRepository
	tRepo repository.TransactionRepository
}

func NewTransactionService(cRepo repository.ClientRepository, tRepo repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		cRepo: cRepo,
		tRepo: tRepo,
	}
}

func (serv *TransactionService) CreateTransaction(ctx context.Context, transaction entity.Transaction, balance int) error {
	var newBalance int
	if transaction.Type == "d" {
		newBalance = balance - transaction.Value
	} else {
		newBalance = balance + transaction.Value
	}

	return serv.tRepo.CreateTransaction(ctx, transaction, newBalance)
}
