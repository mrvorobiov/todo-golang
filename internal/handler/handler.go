package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/mrvorobiov/todo/internal/domain"
)

type TicketsService interface {
	Create(ticket domain.Ticket) error
	FindUnique(id int) (*domain.Ticket, int)
	FindAll() []domain.Ticket
	Delete(id int) error
	Patch(id int, body domain.PatchTicketBody) (*domain.Ticket, error)
}

type Handler struct {
	Service TicketsService
}

func (handler *Handler) InitRoutes(app *fiber.App) {
	app.Get("/tickets", handler.FindAll)
	app.Post("/tickets", handler.NewTicket)
	app.Patch("/tickets/:id", handler.PatchTicket)
	app.Get("/tickets/:id", handler.FindTicket)
	app.Delete("/tickets/:id", handler.DeleteTicket)
}

func NewHandler(service TicketsService) Handler {
	return Handler{
		Service: service,
	}
}

func (handler *Handler) FindAll(ctx fiber.Ctx) error {
	return ctx.JSON(handler.Service.FindAll())
}

func (handler *Handler) NewTicket(ctx fiber.Ctx) error {
	ctx.Accepts("application/json")

	ticket := new(domain.Ticket)

	if err := ctx.Bind().Body(&ticket); err != nil {
		return err
	}

	err := handler.Service.Create(*ticket)

	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, fmt.Sprint(err))
	}

	return ctx.JSON(ticket)
}

func (handler *Handler) PatchTicket(ctx fiber.Ctx) error {
	ctx.Accepts("application/json")

	id, _ := strconv.Atoi(ctx.Params("id"))

	body := new(domain.PatchTicketBody)

	if err := ctx.Bind().Body(&body); err != nil {
		return err
	}

	ticket, _ := handler.Service.Patch(id, *body)

	return ctx.JSON(ticket)
}

func (handler *Handler) FindTicket(ctx fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	ticket, _ := handler.Service.FindUnique(id)

	return ctx.JSON(ticket)
}

func (handler *Handler) DeleteTicket(ctx fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	handler.Service.Delete(id)

	return ctx.JSON(true)
}
