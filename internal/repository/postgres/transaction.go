package postgres

import (
	"context"
	"strconv"

	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/internal/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Implements TransactionRepository
type TransactionRepo struct {
	pool *pgxpool.Pool
}

func NewTransactionRepo(pool *pgxpool.Pool) *TransactionRepo {
	return &TransactionRepo{pool: pool}
}

func (repo *TransactionRepo) GetTenLastTransactionsByUserId(ctx context.Context, id int) ([]entity.Transaction, error) {
	idStr := strconv.Itoa(id)
	rows, err := repo.pool.Query(ctx,
		"SELECT value, type, description, created_at FROM \"transaction\" WHERE client_id = $1 ORDER BY created_at DESC LIMIT 10",
		idStr,
	)
	if err != nil {
		return []entity.Transaction{}, err
	}

	var tr entity.Transaction
	trSlice := make([]entity.Transaction, 0, 10)
	for rows.Next() {
		if err := rows.Scan(&tr.Value, &tr.Type, &tr.Description, &tr.Accomplished); err != nil {
			continue
		}
		trSlice = append(trSlice, tr)
	}

	return trSlice, nil
}

func (repo *TransactionRepo) CreateTransaction(ctx context.Context, tx pgx.Tx, transaction entity.Transaction) error {
	_, err := tx.Exec(ctx,
		"INSERT INTO \"transaction\" (value, type, description, client_id) VALUES ($1, $2, $3, $4)",
		transaction.Value, transaction.Type, transaction.Description, transaction.ClientId,
	)
	if err != nil {
		return err
	}

	return nil
}
