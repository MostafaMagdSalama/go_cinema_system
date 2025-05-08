package ticket

type Repository interface {
	GetAllTickets() ([]Ticket, error)
	GetTicketByID(id int) (Ticket, error)
	GetTicketsByShowID(showID int) ([]Ticket, error)
	GetTicketsByCustomerID(customerID int) ([]Ticket, error)
	CreateTicket(ticket Ticket) (Ticket, error)
	UpdateTicket(id int, ticket Ticket) (Ticket, error)
	DeleteTicket(id int) error
}
