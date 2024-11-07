package entity

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
