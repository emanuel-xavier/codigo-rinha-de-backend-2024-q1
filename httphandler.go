package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/configs"
	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/db"
	"github.com/gofiber/fiber/v2"
)

var (
	ts *TransactionService
	cs *ClientService
)

func createTransactionHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Failed to parse id to string: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var trDto CreateTransactionDto
	if err := c.BodyParser(&trDto); err != nil {
		log.Println("Error parsing request body: ", err)
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	if !trDto.validate() {
		log.Println("Cant validate request body content")
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	trResp, err := ts.CreateTransaction(c.Context(), idInt, trDto)
	if err != nil {
		log.Println(err)
		switch err {
		case ErrInsufficientBalance:
			return c.SendStatus(fiber.StatusUnprocessableEntity)
		case ErrNotFound:
			return c.SendStatus(fiber.StatusNotFound)
		default:
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	return c.Status(fiber.StatusOK).JSON(trResp)
}

func getStatementHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Failed to parse id to string: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	balance, err := cs.GetClientBalance(c.Context(), idInt)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if balance == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	transactions, err := ts.GetLastTenTransactionOfOneUser(c.Context(), idInt)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	balance.BalanceDate = time.Now().Format("2006-01-02T15:04:05.999999Z")
	return c.Status(200).JSON(StatementResponseDto{
		Balance:          *balance,
		LastTransactions: transactions,
	})
}

func Handle() {
	err := configs.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("Failed to load configs %s", err.Error()))
	}
	serverPort := ":" + configs.GetServerPort()

	db.Initialize()
	pool := db.GetPool()

	cs = NewClientService(pool)
	ts = NewTransactionService(pool)

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Post("/clientes/:id/transacoes", createTransactionHandler)
	app.Get("/clientes/:id/extrato", getStatementHandler)

	log.Println("Server runing on port", serverPort)
	log.Fatal(app.Listen(serverPort))
}
