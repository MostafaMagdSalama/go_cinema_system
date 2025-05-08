package customer

import (
	"database/sql"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) GetAllCustomers() ([]Customer, error) {
	rows, err := r.db.Query("SELECT id, name, phone_number FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.ID, &c.Name, &c.PhoneNumber); err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}

	if len(customers) == 0 {
		return []Customer{}, nil
	}

	return customers, nil
}

func (r *PostgresRepository) GetCustomerByID(id int) (Customer, error) {
	row := r.db.QueryRow("SELECT id, name, phone_number FROM customers WHERE id=$1", id)
	var customer Customer
	err := row.Scan(&customer.ID, &customer.Name, &customer.PhoneNumber)
	if err != nil {
		return Customer{}, err
	}
	return customer, nil
}

func (r *PostgresRepository) GetCustomerByPhoneNumber(phoneNumber string) (Customer, error) {
	row := r.db.QueryRow("SELECT id, name, phone_number FROM customers WHERE phone_number=$1", phoneNumber)
	var customer Customer
	err := row.Scan(&customer.ID, &customer.Name, &customer.PhoneNumber)
	if err != nil {
		return Customer{}, err
	}
	return customer, nil
}

func (r *PostgresRepository) CreateCustomer(customer Customer) (Customer, error) {
	err := r.db.QueryRow(
		"INSERT INTO customers (name, phone_number) VALUES ($1, $2) RETURNING id",
		customer.Name, customer.PhoneNumber,
	).Scan(&customer.ID)
	if err != nil {
		return Customer{}, err
	}
	return customer, nil
}

func (r *PostgresRepository) UpdateCustomer(id int, customer Customer) (Customer, error) {
	_, err := r.db.Exec(
		"UPDATE customers SET name=$1, phone_number=$2 WHERE id=$3",
		customer.Name, customer.PhoneNumber, id,
	)
	if err != nil {
		return Customer{}, err
	}
	customer.ID = id
	return customer, nil
}

func (r *PostgresRepository) DeleteCustomer(id int) error {
	_, err := r.db.Exec("DELETE FROM customers WHERE id=$1", id)
	return err
}
