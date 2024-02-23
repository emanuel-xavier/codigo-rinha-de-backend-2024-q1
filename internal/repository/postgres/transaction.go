package postgres

import (
	"context"
	"strconv"

	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Implements TransactionRepository
type PostgresRepo struct {
	pool pgxpool.Pool
}

func (repo *PostgresRepo) GetTenLastTransactionsByUserId(ctx context.Context, id int) ([]entity.Transaction, error) {
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
