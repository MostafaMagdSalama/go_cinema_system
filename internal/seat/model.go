package seat

type Seat struct {
	ID         int    `json:"id" db:"id"`
	SeatNumber string `json:"seat_number" db:"seat_number"`
	SeatType   string `json:"seat_type" db:"seat_type"`
	IsBooked   bool   `json:"is_booked" db:"is_booked"`
	TheatreID  int    `json:"theatre_id" db:"theatre_id"`
}
