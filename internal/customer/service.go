package customer

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllCustomers() ([]Customer, error) {
	return s.repo.GetAllCustomers()
}

func (s *Service) GetCustomerByID(id int) (Customer, error) {
	return s.repo.GetCustomerByID(id)
}

func (s *Service) CreateCustomer(customer Customer) (Customer, error) {
	return s.repo.CreateCustomer(customer)
}

func (s *Service) UpdateCustomer(id int, customer Customer) (Customer, error) {
	return s.repo.UpdateCustomer(id, customer)
}

func (s *Service) DeleteCustomer(id int) error {
	return s.repo.DeleteCustomer(id)
}
