package ticket

import (
	"database/sql"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) GetAllTickets() ([]Ticket, error) {
	rows, err := r.db.Query("SELECT id, price, seat_id, show_id, customer_id, phone_number FROM tickets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []Ticket
	for rows.Next() {
		var t Ticket
		if err := rows.Scan(&t.ID, &t.Price, &t.SeatID, &t.ShowID, &t.CustomerID, &t.PhoneNumber); err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}

	if len(tickets) == 0 {
		return []Ticket{}, nil
	}

	return tickets, nil
}

func (r *PostgresRepository) GetTicketByID(id int) (Ticket, error) {
	row := r.db.QueryRow("SELECT id, price, seat_id, show_id, customer_id, phone_number FROM tickets WHERE id=$1", id)
	var ticket Ticket
	err := row.Scan(&ticket.ID, &ticket.Price, &ticket.SeatID, &ticket.ShowID, &ticket.CustomerID, &ticket.PhoneNumber)
	if err != nil {
		return Ticket{}, err
	}
	return ticket, nil
}

func (r *PostgresRepository) GetTicketsByShowID(showID int) ([]Ticket, error) {
	rows, err := r.db.Query("SELECT id, price, seat_id, show_id, customer_id, phone_number FROM tickets WHERE show_id=$1", showID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []Ticket
	for rows.Next() {
		var t Ticket
		if err := rows.Scan(&t.ID, &t.Price, &t.SeatID, &t.ShowID, &t.CustomerID, &t.PhoneNumber); err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}

	if len(tickets) == 0 {
		return []Ticket{}, nil
	}

	return tickets, nil
}

func (r *PostgresRepository) GetTicketsByCustomerID(customerID int) ([]Ticket, error) {
	rows, err := r.db.Query("SELECT id, price, seat_id, show_id, customer_id, phone_number FROM tickets WHERE customer_id=$1", customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []Ticket
	for rows.Next() {
		var t Ticket
		if err := rows.Scan(&t.ID, &t.Price, &t.SeatID, &t.ShowID, &t.CustomerID, &t.PhoneNumber); err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}

	if len(tickets) == 0 {
		return []Ticket{}, nil
	}

	return tickets, nil
}

func (r *PostgresRepository) CreateTicket(ticket Ticket) (Ticket, error) {
	err := r.db.QueryRow(
		"INSERT INTO tickets (price, seat_id, show_id, customer_id, phone_number) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		ticket.Price, ticket.SeatID, ticket.ShowID, ticket.CustomerID, ticket.PhoneNumber,
	).Scan(&ticket.ID)
	if err != nil {
		return Ticket{}, err
	}
	return ticket, nil
}

func (r *PostgresRepository) UpdateTicket(id int, ticket Ticket) (Ticket, error) {
	_, err := r.db.Exec(
		"UPDATE tickets SET price=$1, seat_id=$2, show_id=$3, customer_id=$4, phone_number=$5 WHERE id=$6",
		ticket.Price, ticket.SeatID, ticket.ShowID, ticket.CustomerID, ticket.PhoneNumber, id,
	)
	if err != nil {
		return Ticket{}, err
	}
	ticket.ID = id
	return ticket, nil
}

func (r *PostgresRepository) DeleteTicket(id int) error {
	_, err := r.db.Exec("DELETE FROM tickets WHERE id=$1", id)
	return err
}
