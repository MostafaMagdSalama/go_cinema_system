package movie 


type Movie struct {
	ID 		int    `json:"id" db:"id"`
	Name 	string `json:"Name" db:"name"`
	Genre 	string `json:"genre" db:"genre"`
	Duration int    `json:"duration" db:"duration_minutes"`
	Release_Date int    `json:"Release_Date" db:"release_date"`
}
