package main

import (
	"fmt"
	"log"

	"github.com/emanuel-xavier/codigo-rinha-de-backend-2024-q1/configs"
	"github.com/gofiber/fiber/v2"
)

func Handle() {
	err := configs.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("Failed to load configs %s", err.Error()))
	}
	serverPort := ":" + configs.GetServerPort()

	app := fiber.New()

	log.Fatal(app.Listen(serverPort))
}
