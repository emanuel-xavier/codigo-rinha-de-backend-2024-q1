package postgres

import (
	"context"
	"errors"

	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/internal/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClientRepo struct {
	pool *pgxpool.Pool
}

func NewClientRepo(pool *pgxpool.Pool) *ClientRepo {
	return &ClientRepo{pool: pool}
}

func (repo *ClientRepo) GetClientById(ctx context.Context, id int) (entity.Client, error) {
	var client entity.Client
	err := repo.pool.QueryRow(ctx,
		"SELECT balance, \"limit\", name, id FROM clients WHERE id = $1",
		id,
	).Scan(&client.Balance, &client.Limit, &client.Name, &client.Id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return client, errors.New("not found")
		}
		return client, err
	}

	return client, nil
}

func (repo *ClientRepo) GetClientByIdAndLock(ctx context.Context, id int) (client entity.Client, tx pgx.Tx, err error) {
	tx, err = repo.pool.Begin(ctx)
	if err != nil {
		return
	}

	err = tx.QueryRow(ctx,
		"SELECT balance, \"limit\", name, id FROM clients WHERE id = $1 FOR UPDATE",
		id,
	).Scan(&client.Balance, &client.Limit, &client.Name, &client.Id)

	if err != nil {
		if err == pgx.ErrNoRows {
			return client, nil, errors.New("not found")
		}
		return client, nil, err
	}

	return client, tx, nil
}

func (repo *ClientRepo) UpdateBalance(ctx context.Context, tx pgx.Tx, client entity.Client) error {
	_, err := tx.Exec(ctx,
		"UPDATE \"clients\" SET balance = $1 WHERE id = $2",
		client.Balance, client.Id,
	)
	if err != nil {
		return err
	}

	return nil

}
