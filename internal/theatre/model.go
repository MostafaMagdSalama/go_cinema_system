package theatre

type Theatre struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Location   string `json:"location" db:"location"`
	TotalSeats int    `json:"total_seats" db:"total_seats"`
}
