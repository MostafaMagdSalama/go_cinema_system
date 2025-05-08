package seat

type Repository interface {
	GetAllSeats() ([]Seat, error)
	GetSeatByID(id int) (Seat, error)
	GetSeatsByTheatreID(theatreID int) ([]Seat, error)
	CreateSeat(seat Seat) (Seat, error)
	UpdateSeat(id int, seat Seat) (Seat, error)
	DeleteSeat(id int) error
}
