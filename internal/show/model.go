package show

type Show struct {
	ID        int    `json:"id" db:"id"`
	ShowTime  string `json:"show_time" db:"show_time"`
	MovieID   int    `json:"movie_id" db:"movie_id"`
	TheatreID int    `json:"theatre_id" db:"theatre_id"`
}
