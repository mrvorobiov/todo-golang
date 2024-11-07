package repository

import (
	"time"

	"github.com/mrvorobiov/todo/internal/domain"
)

type Tickets struct {
	Tickets []domain.Ticket
}

func NewTickets() *Tickets {
	return &Tickets{
		Tickets: make([]domain.Ticket, 0),
	}
}

func (tickets *Tickets) Create(ticket domain.Ticket) error {
	if ticket.Title == "" {
		return domain.ErrMissingTitle
	}
	ticket.Id = len(tickets.Tickets)
	ticket.CreatedAt = time.Now()
	tickets.Tickets = append(tickets.Tickets, ticket)
	return nil
}

func (tickets *Tickets) FindUnique(id int) (ticket *domain.Ticket, index int) {
	for i, entity := range tickets.Tickets {
		if entity.Id == id {
			ticket = &entity
			index = i
		}
	}
	return
}

func (tickets *Tickets) FindAll() []domain.Ticket {
	return tickets.Tickets
}

func (tickets *Tickets) Delete(id int) error {
	_, index := tickets.FindUnique(id)
	if index == -1 {
		return domain.ErrTicketNotFound
	}
	tickets.Tickets = append(tickets.Tickets[:index], tickets.Tickets[index+1:]...)
	return nil
}

func (tickets *Tickets) Patch(id int, body domain.PatchTicketBody) (*domain.Ticket, error) {
	ticket, index := tickets.FindUnique(id)
	if index == -1 {
		return nil, domain.ErrTicketNotFound
	}
	if body.Title != nil {
		ticket.Title = *body.Title
	}
	if ticket.Description != nil {
		ticket.Description = body.Description
	}
	if body.Deadline != nil {
		ticket.Deadline = body.Deadline
	}
	return ticket, nil
}
