package show

type Repository interface {
	GetAllShows() ([]Show, error)
	GetShowByID(id int) (Show, error)
	CreateShow(show Show) (Show, error)
	UpdateShow(id int, show Show) (Show, error)
	DeleteShow(id int) error
}
