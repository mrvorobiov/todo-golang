package main

import (
	"github.com/gofiber/fiber/v3"

	"github.com/mrvorobiov/todo/pkg/tickets/handlers"
)

func main() {
	app := fiber.New()

	handlers.Use(app)

	app.Listen(":80")
}
