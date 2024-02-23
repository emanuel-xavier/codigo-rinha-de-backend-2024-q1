package router

import (
	"encoding/json"
	"log"

	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/configs"
	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/db"
	fiber "github.com/gofiber/fiber/v2"
)

type FiberRouter struct {
}

func (fr *FiberRouter) Init() {
	err := configs.Load()
	if err != nil {
		panic("Falied to load configs")
	}

	serverPort := ":" + configs.GetServerPort()

	db.Initialize()

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Get("health-check", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	log.Fatal(app.Listen(serverPort))
}
