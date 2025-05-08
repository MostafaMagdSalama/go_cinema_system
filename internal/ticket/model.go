package ticket

type Ticket struct {
	ID          int     `json:"id" db:"id"`
	Price       float64 `json:"price" db:"price"`
	SeatID      int     `json:"seat_id" db:"seat_id"`
	ShowID      int     `json:"show_id" db:"show_id"`
	CustomerID  int     `json:"customer_id" db:"customer_id"`
	PhoneNumber string  `json:"phone_number" db:"phone_number"`
}
