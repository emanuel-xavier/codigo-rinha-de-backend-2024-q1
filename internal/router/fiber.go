package router

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/configs"
	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/internal/dto"
	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/internal/entity"
	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/internal/service"
	"github.com/gofiber/fiber/v2"
)

type FiberRouter struct {
	cs *service.ClientService
	ts *service.TransactionService
}

func NewFiberRouter(cs *service.ClientService, ts *service.TransactionService) *FiberRouter {
	return &FiberRouter{
		cs: cs,
		ts: ts,
	}
}

func (fr *FiberRouter) Init() {
	err := configs.Load()
	if err != nil {
		panic("Falied to load configs")
	}

	serverPort := ":" + configs.GetServerPort()

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Get("health-check", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Post("/clientes/:id/transacoes", fr.createTransactionHandler)
	app.Get("/clientes/:id/extrato", fr.getStatementHandler)

	log.Fatal(app.Listen(serverPort))
}

func (fr *FiberRouter) createTransactionHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	var trDto dto.TransactionRequest
	if err := ctx.BodyParser(&trDto); err != nil {
		return ctx.SendStatus(fiber.StatusUnprocessableEntity)
	}

	if err := trDto.Validate(); err != nil {
		return ctx.SendStatus(fiber.StatusUnprocessableEntity)
	}

	client, err := fr.cs.GetClientById(context.Background(), idInt)
	if err != nil {
		if err.Error() == "not found" {
			return ctx.SendStatus(fiber.StatusNotFound)
		}
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	transaction := entity.Transaction{
		Type:       trDto.Type,
		Descrition: trDto.Description,
		Value:      trDto.Value,
		ClientId:   idInt,
	}

	if err := fr.ts.CreateTransaction(ctx.Context(), transaction, client.Balance); err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return nil
}

func (fr *FiberRouter) getStatementHandler(ctx *fiber.Ctx) error {
	return nil
}
