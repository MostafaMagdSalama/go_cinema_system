package show

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllShows() ([]Show, error) {
	return s.repo.GetAllShows()
}

func (s *Service) GetShowByID(id int) (Show, error) {
	return s.repo.GetShowByID(id)
}

func (s *Service) CreateShow(show Show) (Show, error) {
	return s.repo.CreateShow(show)
}

func (s *Service) UpdateShow(id int, show Show) (Show, error) {
	return s.repo.UpdateShow(id, show)
}

func (s *Service) DeleteShow(id int) error {
	return s.repo.DeleteShow(id)
}
