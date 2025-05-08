package theatre

import (
	"database/sql"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) GetAllTheatres() ([]Theatre, error) {
	rows, err := r.db.Query("SELECT id, name, location, total_seats FROM theatres")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var theatres []Theatre
	for rows.Next() {
		var t Theatre
		if err := rows.Scan(&t.ID, &t.Name, &t.Location, &t.TotalSeats); err != nil {
			return nil, err
		}
		theatres = append(theatres, t)
	}

	if len(theatres) == 0 {
		return []Theatre{}, nil
	}

	return theatres, nil
}

func (r *PostgresRepository) GetTheatreByID(id int) (Theatre, error) {
	row := r.db.QueryRow("SELECT id, name, location, total_seats FROM theatres WHERE id=$1", id)
	var theatre Theatre
	err := row.Scan(&theatre.ID, &theatre.Name, &theatre.Location, &theatre.TotalSeats)
	if err != nil {
		return Theatre{}, err
	}
	return theatre, nil
}

func (r *PostgresRepository) CreateTheatre(theatre Theatre) (Theatre, error) {
	err := r.db.QueryRow(
		"INSERT INTO theatres (name, location, total_seats) VALUES ($1, $2, $3) RETURNING id",
		theatre.Name, theatre.Location, theatre.TotalSeats,
	).Scan(&theatre.ID)
	if err != nil {
		return Theatre{}, err
	}
	return theatre, nil
}

func (r *PostgresRepository) UpdateTheatre(id int, theatre Theatre) (Theatre, error) {
	_, err := r.db.Exec(
		"UPDATE theatres SET name=$1, location=$2, total_seats=$3 WHERE id=$4",
		theatre.Name, theatre.Location, theatre.TotalSeats, id,
	)
	if err != nil {
		return Theatre{}, err
	}
	theatre.ID = id
	return theatre, nil
}

func (r *PostgresRepository) DeleteTheatre(id int) error {
	_, err := r.db.Exec("DELETE FROM theatres WHERE id=$1", id)
	return err
}
