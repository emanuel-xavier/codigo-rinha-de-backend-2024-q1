package main

import (
	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/configs"
	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/db"
	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/internal/repository/postgres"
	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/internal/router"
	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/internal/service"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic("Failed to load configs")
	}
	db.Initialize()
	pool := db.GetPool()

	cRepo := postgres.NewClientRepo(pool)
	tRepo := postgres.NewTransactionRepo(pool)

	cService := service.NewClientService(cRepo, tRepo)
	tService := service.NewTransactionService(cRepo, tRepo)

	r := router.NewFiberRouter(cService, tService)
	r.Init()
}
