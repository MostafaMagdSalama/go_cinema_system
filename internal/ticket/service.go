package ticket

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllTickets() ([]Ticket, error) {
	return s.repo.GetAllTickets()
}

func (s *Service) GetTicketByID(id int) (Ticket, error) {
	return s.repo.GetTicketByID(id)
}

func (s *Service) CreateTicket(ticket Ticket) (Ticket, error) {
	return s.repo.CreateTicket(ticket)
}

func (s *Service) UpdateTicket(id int, ticket Ticket) (Ticket, error) {
	return s.repo.UpdateTicket(id, ticket)
}

func (s *Service) DeleteTicket(id int) error {
	return s.repo.DeleteTicket(id)
}
