package theatre

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllTheatres() ([]Theatre, error) {
	return s.repo.GetAllTheatres()
}

func (s *Service) GetTheatreByID(id int) (Theatre, error) {
	return s.repo.GetTheatreByID(id)
}

func (s *Service) CreateTheatre(theatre Theatre) (Theatre, error) {
	return s.repo.CreateTheatre(theatre)
}

func (s *Service) UpdateTheatre(id int, theatre Theatre) (Theatre, error) {
	return s.repo.UpdateTheatre(id, theatre)
}

func (s *Service) DeleteTheatre(id int) error {
	return s.repo.DeleteTheatre(id)
}
