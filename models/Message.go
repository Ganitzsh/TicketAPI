package models

// Message represents a ticket's message
type Message struct {
	ID       int
	TicketID int
	AuthorID int
	Body     string
}
