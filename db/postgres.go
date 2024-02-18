package db

import (
	"context"
	"fmt"
	"log"

	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/configs"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func Initialize() {
	err := configs.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("Failed to load configs: %s", err.Error()))
	}

	conf := configs.GetDB()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	config, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		log.Fatal(fmt.Errorf("Failed to parse connection string: %s", err.Error()))
	}

	pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal(fmt.Errorf("Failed to create connection pool: %s", err.Error()))
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal(fmt.Errorf("Failed to ping database: %s", err.Error()))
	}

	log.Println("Successfully connected to database")
}

func GetPool() *pgxpool.Pool {
	return pool
}
