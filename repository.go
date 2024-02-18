package main

import (
	"context"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ClientRepository struct {
	dbPool *pgxpool.Pool
}

func NewClientRepository(dbPool *pgxpool.Pool) *ClientRepository {
	return &ClientRepository{
		dbPool: dbPool,
	}
}

func (cm *ClientRepository) checkIfClientExists(c context.Context, id int) (bool, error) {
	var exists bool

	idStr := strconv.Itoa(id)
	err := cm.dbPool.QueryRow(c,
		"SELECT EXISTS(SELECT 1 FROM clients WHERE id = $1)",
		idStr,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
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

func NewTransactionRepository(dbPool *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{
		dbPool: dbPool,
	}
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

func (tm *TransactionRepository) getLastTenTransactionsOfOneUser(c context.Context, id int) ([]TransactionDto, error) {
	var tr []TransactionDto

	idStr := strconv.Itoa(id)

	rows, err := tm.dbPool.Query(c,
		"SELECT value, type, description, created_at FROM \"transaction\" WHERE client_id = $1 ORDER BY created_at DESC LIMIT 10",
		idStr,
	)
	defer rows.Close()

	if err != nil {
		return tr, err
	}

	var t TransactionDto
	for rows.Next() {
		if err := rows.Scan(&t.Value, &t.Type, &t.Description, &t.Accomplished); err != nil {
			log.Println("Failed to scan a row")
		}
		tr = append(tr, t)
	}
	if err := rows.Err(); err != nil {
		return tr, err
	}

	return tr, nil
}
