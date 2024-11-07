package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/mrvorobiov/todo/pkg/tickets/entity"
	"github.com/mrvorobiov/todo/pkg/tickets/repository"
)

func Use(app *fiber.App) {
	app.Get("/tickets", GetAll)
	app.Post("/tickets", NewTicket)
	app.Patch("/tickets/:id", PatchTicket)
	app.Get("/tickets/:id", FindTicket)
	app.Delete("/tickets/:id", DeleteTicket)
}

func GetAll(ctx fiber.Ctx) error {
	return ctx.JSON(repository.GetAll())
}

func NewTicket(ctx fiber.Ctx) error {
	ctx.Accepts("application/json")

	ticket := new(entity.Ticket)

	if err := ctx.Bind().Body(&ticket); err != nil {
		return err
	}

	err := repository.NewTicket(ticket)

	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, fmt.Sprint(err))
	}

	return ctx.JSON(ticket)
}

func PatchTicket(ctx fiber.Ctx) error {
	ctx.Accepts("application/json")

	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "An identifier should be an integer!")
	}

	ticket := new(entity.Ticket)

	if err := ctx.Bind().Body(&ticket); err != nil {
		return err
	}

	ticket = repository.PatchTicket(id, ticket)

	return ctx.JSON(ticket)
}

func FindTicket(ctx fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "An identifier should be an integer!")
	}

	ticket, _ := repository.FindUnique(id)

	return ctx.JSON(ticket)
}

func DeleteTicket(ctx fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "An identifier should be an integer!")
	}

	repository.DeleteTicket(id)

	return ctx.JSON(true)
}
