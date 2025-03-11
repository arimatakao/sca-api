package storager

type DbCat struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	YearsOfExperience int    `json:"years_of_experience"`
	Breed             string `json:"breed"`
	Salary            int    `json:"salary"`
}

type Cat struct {
	Name              string `json:"name"`
	YearsOfExperience int    `json:"years_of_experience"`
	Breed             string `json:"breed"`
	Salary            int    `json:"salary"`
}

type CatStorager interface {
	Close() error
	GetAllCats() ([]DbCat, error)
	GetSpecificCat(id string) (DbCat, error)
	CreateCat(c Cat) error
	UpdateCat(id string, c Cat) error
	DeleteCat(id string) error
}
