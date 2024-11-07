package domain

import (
	"errors"
	"time"
)

var (
	ErrTicketNotFound = errors.New("ticket is not found")
	ErrMissingTitle   = errors.New("title is missing")
)

type Ticket struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	IsDone      bool       `json:"isDone"`
	Description *string    `json:"description"`
	Deadline    *time.Time `json:"deadline"`
	CreatedAt   time.Time  `json:"createdAt"`
}

type PatchTicketBody struct {
	Title       *string    `json:"title"`
	IsDone      *bool      `json:"isDone"`
	Description *string    `json:"description"`
	Deadline    *time.Time `json:"deadline"`
}
