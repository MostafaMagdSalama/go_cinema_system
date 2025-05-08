package movie

import (
	"database/sql"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) GetAllMovies() ([]Movie, error) {
	rows, err := r.db.Query("SELECT id, name, genre, duration_minutes FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var m Movie
		if err := rows.Scan(&m.ID, &m.Name, &m.Genre, &m.Duration); err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}
	return movies, nil
}

func (r *PostgresRepository) GetMovieByID(id int) (Movie, error) {
	row := r.db.QueryRow("SELECT id, name, genre, duration FROM movies WHERE id=$1", id)
	var movie Movie
	err := row.Scan(&movie.ID, &movie.Name, &movie.Genre, &movie.Duration)
	if err != nil {
		return Movie{}, err
	}
	return movie, nil
}

func (r *PostgresRepository) CreateMovie(movie Movie) (Movie, error) {
	err := r.db.QueryRow(
		"INSERT INTO movies (name genre, duration) VALUES ($1, $2, $3) RETURNING id",
		movie.Name, movie.Genre, movie.Duration,
	).Scan(&movie.ID)
	if err != nil {
		return Movie{}, err
	}
	return movie, nil
}

func (r *PostgresRepository) UpdateMovie(id int, movie Movie) (Movie, error) {
	_, err := r.db.Exec(
		"UPDATE movies SET name=$1, genre=$2, duration=$3 WHERE id=$4",
		movie.Name, movie.Genre, movie.Duration, id,
	)
	if err != nil {
		return Movie{}, err
	}
	movie.ID = id
	return movie, nil
}

func (r *PostgresRepository) DeleteMovie(id int) error {
	_, err := r.db.Exec("DELETE FROM movies WHERE id=$1", id)
	return err
}
