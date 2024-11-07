package service

import (
	"github.com/mrvorobiov/todo/internal/domain"
)

type TicketsRepository interface {
	Create(ticket domain.Ticket) error
	FindUnique(id int) (*domain.Ticket, int)
	FindAll() []domain.Ticket
	Delete(id int) error
	Patch(id int, body domain.PatchTicketBody) (*domain.Ticket, error)
}

type Tickets struct {
	repository TicketsRepository
}

func NewTickets(repository TicketsRepository) Tickets {
	return Tickets{
		repository: repository,
	}
}

func (tickets Tickets) Create(ticket domain.Ticket) error {
	return tickets.repository.Create(ticket)
}

func (tickets Tickets) FindUnique(id int) (*domain.Ticket, int) {
	return tickets.repository.FindUnique(id)
}

func (tickets Tickets) FindAll() []domain.Ticket {
	return tickets.repository.FindAll()
}

func (tickets Tickets) Delete(id int) error {
	return tickets.repository.Delete(id)
}

func (tickets Tickets) Patch(id int, body domain.PatchTicketBody) (*domain.Ticket, error) {
	return tickets.repository.Patch(id, body)
}
