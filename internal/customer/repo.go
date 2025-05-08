package customer

type Repository interface {
	GetAllCustomers() ([]Customer, error)
	GetCustomerByID(id int) (Customer, error)
	GetCustomerByPhoneNumber(phoneNumber string) (Customer, error)
	CreateCustomer(customer Customer) (Customer, error)
	UpdateCustomer(id int, customer Customer) (Customer, error)
	DeleteCustomer(id int) error
}
