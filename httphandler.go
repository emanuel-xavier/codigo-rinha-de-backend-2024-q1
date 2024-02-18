package main

import (
	"fmt"
	"log"

	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/configs"
	"github.com/gofiber/fiber/v2"
)

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

	app := fiber.New()

	app.Post("/clientes/:id/transacoes", createTransactionHandler)
	app.Get("/clientes/:id/extrato", getStatementHandler)

	log.Fatal(app.Listen(serverPort))
}
