package show

import (
	"database/sql"
	"fmt"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) GetAllShows() ([]Show, error) {
	rows, err := r.db.Query("SELECT id, show_time, movie_id, theatre_id FROM shows")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shows []Show
	for rows.Next() {
		var s Show
		if err := rows.Scan(&s.ID, &s.ShowTime, &s.MovieID, &s.TheatreID); err != nil {
			return nil, err
		}
		shows = append(shows, s)
	}

	if len(shows) == 0 {
		return []Show{}, nil
	}

	fmt.Println("get all shows ", shows)
	return shows, nil
}

func (r *PostgresRepository) GetShowByID(id int) (Show, error) {
	row := r.db.QueryRow("SELECT id, show_time, movie_id, theatre_id FROM shows WHERE id=$1", id)
	var show Show
	err := row.Scan(&show.ID, &show.ShowTime, &show.MovieID, &show.TheatreID)
	if err != nil {
		return Show{}, err
	}
	return show, nil
}

func (r *PostgresRepository) CreateShow(show Show) (Show, error) {
	err := r.db.QueryRow(
		"INSERT INTO shows (show_time, movie_id, theatre_id) VALUES ($1, $2, $3) RETURNING id",
		show.ShowTime, show.MovieID, show.TheatreID,
	).Scan(&show.ID)
	if err != nil {
		return Show{}, err
	}
	return show, nil
}

func (r *PostgresRepository) UpdateShow(id int, show Show) (Show, error) {
	_, err := r.db.Exec(
		"UPDATE shows SET show_time=$1, movie_id=$2, theatre_id=$3 WHERE id=$4",
		show.ShowTime, show.MovieID, show.TheatreID, id,
	)
	if err != nil {
		return Show{}, err
	}
	show.ID = id
	return show, nil
}

func (r *PostgresRepository) DeleteShow(id int) error {
	_, err := r.db.Exec("DELETE FROM shows WHERE id=$1", id)
	return err
}
