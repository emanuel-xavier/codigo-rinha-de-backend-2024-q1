package main

import (
	"fmt"
	"log"

	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/configs"
	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/db"
	"github.com/gofiber/fiber/v2"
)

var cr *ClientRepository
var tr *TransactionRepository

func createTransactionHandler(c *fiber.Ctx) error {
	return nil
}

func getStatementHandler(c *fiber.Ctx) error {
	return nil
}

func Handle() {
	err := configs.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("Failed to load configs %s", err.Error()))
	}
	serverPort := ":" + configs.GetServerPort()

	db.Initialize()
	pool := db.GetPool()

	cr = NewClientRepository(pool)
	tr = NewTransactionRepository(pool)

	app := fiber.New()

	app.Post("/clientes/:id/transacoes", createTransactionHandler)
	app.Get("/clientes/:id/extrato", getStatementHandler)

	log.Fatal(app.Listen(serverPort))
}
