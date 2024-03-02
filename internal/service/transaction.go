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
	c, tx, err := serv.cRepo.GetClientByIdAndLock(ctx, transaction.ClientId)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	*client = c

	if transaction.Type == "d" {
		client.Balance -= transaction.Value
		if client.Balance < -client.Limit {
			return errors.New("insufficient funds")
		}
	} else {
		client.Balance += transaction.Value
	}

	if err = serv.cRepo.UpdateBalance(ctx, tx, *client); err != nil {
		return err
	}

	err = serv.tRepo.CreateTransaction(ctx, tx, transaction)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
