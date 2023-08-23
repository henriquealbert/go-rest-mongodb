package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func healthcheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {
	app := fiber.New()

	app.Get("/healthcheck", healthcheck)

	log.Fatal(app.Listen(":3000"))
}
