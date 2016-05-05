package models

// Ticket represents a ticket opened by a user
type Ticket struct {
	ID         int
	EmployeeID int
	UserID     int
	Subject    string
	// Messages   []*Message `sql_table:"messages"`
}
