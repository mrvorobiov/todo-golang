package main

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
)

type Ticket struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Deadline    *int    `json:"deadline"`
	CreatedAt   int     `json:"createdAt"`
}

func (self *Ticket) Merge(ticket Ticket) {
	if ticket.Title != "" {
		self.Title = ticket.Title
	}
	if ticket.Description != nil {
		self.Description = ticket.Description
	}
	if ticket.Deadline != nil {
		self.Deadline = ticket.Deadline
	}
}

func FindTicket(id int, repository *[]Ticket) (*Ticket, int) {
	var ticket *Ticket
	var index int

	for i, entity := range *repository {
		if entity.Id == id {
			ticket = &entity
			index = i
		}
	}

	return ticket, index
}

func main() {
	app := fiber.New()

	tickets := make([]Ticket, 0)

	app.Get("/tickets", func(ctx fiber.Ctx) error {
		return ctx.JSON(tickets)
	})

	app.Post("/tickets", func(ctx fiber.Ctx) error {
		ctx.Accepts("application/json")

		ticket := new(Ticket)

		if err := ctx.Bind().Body(&ticket); err != nil {
			return err
		}

		if ticket.Title == "" {
			return fiber.NewError(fiber.StatusBadRequest, "A title is required!")
		}

		ticket.Id = len(tickets)
		ticket.CreatedAt = int(time.Now().Unix())

		tickets = append(tickets, *ticket)

		return ctx.JSON(ticket)
	})

	app.Put("/tickets/:id", func(ctx fiber.Ctx) error {
		ctx.Accepts("application/json")

		id, err := strconv.Atoi(ctx.Params("id"))

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "An identifier should be an integer!")
		}

		match, _ := FindTicket(id, &tickets)

		if match == nil {
			return fiber.NewError(fiber.StatusNotFound, "No ticket matched!")
		}

		ticket := new(Ticket)

		if err := ctx.Bind().Body(&ticket); err != nil {
			return err
		}

		match.Merge(*ticket)

		return ctx.JSON(match)
	})

	app.Get("/tickets/:id", func(ctx fiber.Ctx) error {
		id, err := strconv.Atoi(ctx.Params("id"))

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "An identifier should be an integer!")
		}

		ticket, _ := FindTicket(id, &tickets)

		if ticket == nil {
			return fiber.NewError(fiber.StatusNotFound, "Ticket doesn't exist!")
		}

		return ctx.JSON(ticket)
	})

	app.Delete("/tickets/:id", func(ctx fiber.Ctx) error {
		id, err := strconv.Atoi(ctx.Params("id"))

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "An identifier should be an integer!")
		}

		_, index := FindTicket(id, &tickets)

		tickets = append(tickets[:index], tickets[index+1:]...)

		return ctx.JSON(true)
	})

	app.Listen(":80")
}
