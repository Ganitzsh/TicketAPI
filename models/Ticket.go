package models

// Ticket represents a ticket opened by a user
type Ticket struct {
	ID        int
	EmplyeeID int
	UserID    int
	Subject   string
}
