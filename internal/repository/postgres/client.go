package postgres

import (
	"context"
	"errors"

	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/internal/entity"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClientRepo struct {
	pool pgxpool.Pool
}

func (repo *ClientRepo) GetClientById(ctx context.Context, id int) (entity.Client, error) {
	var client entity.Client
	err := repo.pool.QueryRow(ctx,
		"SELECT balance, \"limit\", name FROM clients WHERE id = $1",
		id,
	).Scan(&client.Balance, &client.Limit, &client.Name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return client, errors.New("not found")
		}
		return client, err
	}

	return client, nil
}
