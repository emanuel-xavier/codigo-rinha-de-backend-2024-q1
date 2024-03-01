package repository

import (
	"context"

	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/internal/entity"
	"github.com/jackc/pgx/v5"
)

type ClientRepository interface {
	GetClientById(ctx context.Context, id int) (entity.Client, error)
	GetClientByIdAndLock(ctx context.Context, id int) (entity.Client, pgx.Tx, error)
	UpdateBalance(ctx context.Context, tx pgx.Tx, Client entity.Client) error
}

type TransactionRepository interface {
	GetTenLastTransactionsByUserId(ctx context.Context, id int) ([]entity.Transaction, error)
	CreateTransaction(ctx context.Context, tx pgx.Tx, transaction entity.Transaction) error
}
