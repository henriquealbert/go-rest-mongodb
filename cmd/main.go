package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/henriquealbert/go-rest-mongodb/internal/handlers"
)

func healthcheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {
	app := fiber.New()

	app.Get("/healthcheck", healthcheck)

	app.Get("/api/v1/products", handlers.GetAllProducts)
	app.Post("/api/v1/products", handlers.CreateProduct)

	log.Fatal(app.Listen(":3000"))
}
