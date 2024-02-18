package main

import "github.com/jackc/pgx/v5/pgxpool"

type ClientRepository struct {
	dbPool *pgxpool.Pool
}

type TransactionRepository struct {
	dbPool *pgxpool.Pool
}
