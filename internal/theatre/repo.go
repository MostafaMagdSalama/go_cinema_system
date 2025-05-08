package theatre

type Repository interface {
	GetAllTheatres() ([]Theatre, error)
	GetTheatreByID(id int) (Theatre, error)
	CreateTheatre(theatre Theatre) (Theatre, error)
	UpdateTheatre(id int, theatre Theatre) (Theatre, error)
	DeleteTheatre(id int) error
}
