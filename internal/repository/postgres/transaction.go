package postgres

import (
	"context"
	"errors"
	"strconv"

	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Implements TransactionRepository
type TransactionRepo struct {
	pool pgxpool.Pool
}

func (repo *TransactionRepo) GetTenLastTransactionsByUserId(ctx context.Context, id int) ([]entity.Transaction, error) {
	idStr := strconv.Itoa(id)
	rows, err := repo.pool.Query(ctx,
		"SELECT value, type, description, created_ad FROM \"transaction\" WHERE client_id = $1 ORDER BY created_at DEC LIMIT 10",
		idStr,
	)
	if err != nil {
		return []entity.Transaction{}, err
	}

	var tr entity.Transaction
	trSlice := make([]entity.Transaction, 0, 10)
	for rows.Next() {
		if err := rows.Scan(&tr.Value, &tr.Type, &tr.Descrition, &tr.Accomplished); err != nil {
			continue
		}
		trSlice = append(trSlice, tr)
	}

	return trSlice, nil
}

func (repo *TransactionRepo) CreateTransaction(ctx context.Context, transaction entity.Transaction, balance int) error {
	tx, err := repo.pool.Begin(ctx)
	if err != nil {
		return err
	}

	cmt, err := tx.Exec(ctx,
		"UPDATE \"clients\" SET balance = $1 WHERE id = $2",
		balance, transaction.ClientId,
	)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	if cmt.RowsAffected() < 1 {
		tx.Rollback(ctx)
		return errors.New("client not found")
	}

	_, err = tx.Exec(ctx,
		"INSERT INTO \"transaction\" (value, type, description, client_id) VALUES ($1, $2, $3, $4)",
		transaction.Value, transaction.Type, transaction.Descrition, transaction.ClientId,
	)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	tx.Commit(ctx)
	return nil
}
