package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mrvorobiov/todo/internal/handler"
	"github.com/mrvorobiov/todo/internal/repository"
	"github.com/mrvorobiov/todo/internal/service"
)

func main() {
	app := fiber.New()

	repository := repository.NewTickets()
	service := service.NewTickets(repository)
	handlers := handler.NewHandler(service)

	handlers.InitRoutes(app)

	app.Listen(":80")
}
