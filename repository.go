package main

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ClientRepository struct {
	dbPool *pgxpool.Pool
}

func (cm *ClientRepository) getClientBalance(c context.Context, id int) (*BalanceDto, error) {
	var b BalanceDto

	idStr := strconv.Itoa(id)
	row := cm.dbPool.QueryRow(c,
		"SELECT \"limit\", balance FROM \"clients\" WHERE id = $1",
		idStr,
	)

	err := row.Scan(&b.Limit, &b.Total)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

type TransactionRepository struct {
	dbPool *pgxpool.Pool
}

func (tm *TransactionRepository) createTransaction(c context.Context, value, uId int, t, desc string) error {
	_, err := tm.dbPool.Exec(c,
		"INSERT INTO \"transaction\" (value, type, description, client_id) VALUES ($1, $2, $3, $4)",
		value, t, desc, uId,
	)
	if err != nil {
		return err
	}

	return nil
}
