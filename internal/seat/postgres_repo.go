package seat

import (
	"database/sql"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) GetAllSeats() ([]Seat, error) {
	rows, err := r.db.Query("SELECT id, seat_number, seat_type, is_booked, theatre_id FROM seats")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seats []Seat
	for rows.Next() {
		var s Seat
		if err := rows.Scan(&s.ID, &s.SeatNumber, &s.SeatType, &s.IsBooked, &s.TheatreID); err != nil {
			return nil, err
		}
		seats = append(seats, s)
	}

	if len(seats) == 0 {
		return []Seat{}, nil
	}

	return seats, nil
}

func (r *PostgresRepository) GetSeatByID(id int) (Seat, error) {
	row := r.db.QueryRow("SELECT id, seat_number, seat_type, is_booked, theatre_id FROM seats WHERE id=$1", id)
	var seat Seat
	err := row.Scan(&seat.ID, &seat.SeatNumber, &seat.SeatType, &seat.IsBooked, &seat.TheatreID)
	if err != nil {
		return Seat{}, err
	}
	return seat, nil
}

func (r *PostgresRepository) GetSeatsByTheatreID(theatreID int) ([]Seat, error) {
	rows, err := r.db.Query("SELECT id, seat_number, seat_type, is_booked, theatre_id FROM seats WHERE theatre_id=$1", theatreID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seats []Seat
	for rows.Next() {
		var s Seat
		if err := rows.Scan(&s.ID, &s.SeatNumber, &s.SeatType, &s.IsBooked, &s.TheatreID); err != nil {
			return nil, err
		}
		seats = append(seats, s)
	}

	if len(seats) == 0 {
		return []Seat{}, nil
	}

	return seats, nil
}

func (r *PostgresRepository) CreateSeat(seat Seat) (Seat, error) {
	err := r.db.QueryRow(
		"INSERT INTO seats (seat_number, seat_type, is_booked, theatre_id) VALUES ($1, $2, $3, $4) RETURNING id",
		seat.SeatNumber, seat.SeatType, seat.IsBooked, seat.TheatreID,
	).Scan(&seat.ID)
	if err != nil {
		return Seat{}, err
	}
	return seat, nil
}

func (r *PostgresRepository) UpdateSeat(id int, seat Seat) (Seat, error) {
	_, err := r.db.Exec(
		"UPDATE seats SET seat_number=$1, seat_type=$2, is_booked=$3, theatre_id=$4 WHERE id=$5",
		seat.SeatNumber, seat.SeatType, seat.IsBooked, seat.TheatreID, id,
	)
	if err != nil {
		return Seat{}, err
	}
	seat.ID = id
	return seat, nil
}

func (r *PostgresRepository) DeleteSeat(id int) error {
	_, err := r.db.Exec("DELETE FROM seats WHERE id=$1", id)
	return err
}
