package main

import (
	"context"
	// "log"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	getClientBalanceQuery             = "SELECT \"limit\", balance FROM clients WHERE id = $1"
	clientExistsQuery                 = "SELECT EXISTS(SELECT 1 FROM clients WHERE id = $1)"
	inserTransactionQuery             = "INSERT INTO \"transaction\" (value, type, description, client_id) VALUES ($1, $2, $3, $4)"
	getLastTenTransactionOfAUserQuery = "SELECT value, type, description, created_at FROM \"transaction\" WHERE client_id = $1 ORDER BY created_at DESC LIMIT 10"
	selectClientForUpdate             = "SELECT balance, \"limit\" FROM clients WHERE id = $1 FOR UPDATE"
	updateClientBalance               = "UPDATE clients SET balance = $1 WHERE id = $2"
)

type ClientService struct {
	dbPool *pgxpool.Pool
}

func NewClientService(dbPool *pgxpool.Pool) *ClientService {
	return &ClientService{
		dbPool: dbPool,
	}
}

func (cs *ClientService) ClientExists(c context.Context, id int) (bool, error) {
	var exists bool

	idStr := strconv.Itoa(id)
	err := cs.dbPool.QueryRow(c,
		clientExistsQuery,
		idStr,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, err
}

func (cs *ClientService) GetClientBalance(c context.Context, id int) (*BalanceDto, error) {
	cBalance := &BalanceDto{}
	idStr := strconv.Itoa(id)
	err := cs.dbPool.QueryRow(c,
		getClientBalanceQuery,
		idStr,
	).Scan(&cBalance.Limit, &cBalance.Total)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return cBalance, nil
}

type TransactionService struct {
	dbPool *pgxpool.Pool
}

func NewTransactionService(dbPool *pgxpool.Pool) *TransactionService {
	return &TransactionService{
		dbPool: dbPool,
	}
}

func (ts *TransactionService) GetLastTenTransactionOfOneUser(c context.Context, id int) ([]TransactionDto, error) {
	idStr := strconv.Itoa(id)
	rows, err := ts.dbPool.Query(c,
		getLastTenTransactionOfAUserQuery,
		idStr,
	)
	if err != nil {
		return []TransactionDto{}, err
	}

	var tr TransactionDto
	trSlice := make([]TransactionDto, 0, 10)
	for rows.Next() {
		if err := rows.Scan(&tr.Value, &tr.Type, &tr.Description, &tr.Accomplished); err != nil {
			// log.Println("Failed to scan a row")
			continue
		}
		trSlice = append(trSlice, tr)
	}
	return trSlice, nil
}

func (ts *TransactionService) CreateTransaction(c context.Context, uId int, tr CreateTransactionDto) (*TransactionResponseDto, error) {
	tx, err := ts.dbPool.Begin(c)
	idStr := strconv.Itoa(uId)
	if err != nil {
		return nil, ErrDatabaseFailure
	}
	defer tx.Rollback(c)

	var balance BalanceDto
	err = tx.QueryRow(c,
		selectClientForUpdate,
		idStr,
	).Scan(&balance.Total, &balance.Limit)

	newBalance := balance.Total - tr.Value

	if newBalance < 0 && tr.Type == "d" {
		return nil, ErrInsufficientBalance
	}

	if newBalance < -balance.Limit && tr.Type == "c" {
		return nil, ErrInsufficientLimit
	}

	_, err = tx.Exec(c,
		updateClientBalance,
		newBalance, uId,
	)
	if err != nil {
		return nil, ErrDatabaseFailure
	}

	_, err = tx.Exec(c,
		inserTransactionQuery,
		tr.Value, tr.Type, tr.Description, idStr,
	)
	if err != nil {
		return nil, ErrDatabaseFailure
	}

	var trResp TransactionResponseDto
	trResp.Balance = newBalance
	trResp.Limit = balance.Limit

	err = tx.Commit(c)
	if err != nil {
		return nil, ErrDatabaseFailure
	}
	return &trResp, nil
}
