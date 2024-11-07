package repository

import (
	"errors"
	"time"

	"github.com/mrvorobiov/todo/pkg/tickets/entity"
)

var repository []entity.Ticket = make([]entity.Ticket, 0)

func GetAll() *[]entity.Ticket {
	return &repository
}

func NewTicket(ticket *entity.Ticket) error {
	if ticket.Title == "" {
		return errors.New("A title is required!")
	}
	ticket.Id = len(repository)
	ticket.CreatedAt = int(time.Now().Unix())
	repository = append(repository, *ticket)
	return nil
}

func FindUnique(id int) (*entity.Ticket, int) {
	var ticket *entity.Ticket
	var index int

	for i, entity := range repository {
		if entity.Id == id {
			ticket = &entity
			index = i
		}
	}

	return ticket, index
}

func PatchTicket(id int, patch *entity.Ticket) *entity.Ticket {
	ticket, _ := FindUnique(id)
	ticket.Merge(*patch)
	return ticket
}

func DeleteTicket(id int) bool {
	_, index := FindUnique(id)
	repository = append(repository[:index], repository[index+1:]...)
	return true
}
