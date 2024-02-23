package service

import (
	"context"
	"errors"

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

func (serv *TransactionService) CreateTransaction(ctx context.Context, transaction entity.Transaction, client *entity.Client) error {
	if transaction.Type == "d" {
		client.Balance -= transaction.Value
		if client.Balance < -client.Limit {
			return errors.New("insufficient funds")
		}
	} else {
		client.Balance += transaction.Value
	}

	return serv.tRepo.CreateTransaction(ctx, transaction, client.Balance)
}
