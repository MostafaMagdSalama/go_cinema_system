package movie 

type Repository interface {
	GetAllMovies() ([]Movie, error)
	GetMovieByID(id int) (Movie, error)
	CreateMovie(movie Movie) (Movie, error)
	UpdateMovie(id int, movie Movie) (Movie, error)
	DeleteMovie(id int) error
}

