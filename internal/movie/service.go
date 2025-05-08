package movie

import "fmt"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllMovies() ([]Movie, error) {
	movies, err := s.repo.GetAllMovies()
	fmt.Println("Movies fetched from DB:", movies)

	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (s *Service) EditMovie(id int, movie Movie) (*Movie, error) {
	updatedMovie, err := s.repo.UpdateMovie(id, movie)
	if err != nil {
		return nil, err
	}
	return &updatedMovie, nil
}

func (s *Service) RemoveMovie(id int) error {
	return s.repo.DeleteMovie(id)
}

func (s *Service) AddMovie(movie Movie) (*Movie, error) {
	newMovie, err := s.repo.CreateMovie(movie)
	if err != nil {
		return nil, err
	}
	return &newMovie, nil
}
