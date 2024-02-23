package repository

import (
	"context"

	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/internal/entity"
)

type ClientRepository interface {
	GetClientById(ctx context.Context, id int) (entity.Client, error)
}

type TransactionRepository interface {
	GetTenLastTransactionsByUserId(ctx context.Context, id int) ([]entity.Transaction, error)
	CreateTransaction(ctx context.Context, transaction entity.Transaction, balance int) error
}
