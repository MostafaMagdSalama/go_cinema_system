package seat

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllSeats() ([]Seat, error) {
	return s.repo.GetAllSeats()
}

func (s *Service) GetSeatByID(id int) (Seat, error) {
	return s.repo.GetSeatByID(id)
}

func (s *Service) GetSeatsByTheatreID(theatreID int) ([]Seat, error) {
	return s.repo.GetSeatsByTheatreID(theatreID)
}

func (s *Service) CreateSeat(seat Seat) (Seat, error) {
	return s.repo.CreateSeat(seat)
}

func (s *Service) UpdateSeat(id int, seat Seat) (Seat, error) {
	return s.repo.UpdateSeat(id, seat)
}

func (s *Service) DeleteSeat(id int) error {
	return s.repo.DeleteSeat(id)
}
