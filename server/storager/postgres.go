package storager

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	db *sql.DB
}

func NewPostgreSQL(url string) (*PostgreSQL, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgreSQL{
		db: db,
	}, nil
}

func (p *PostgreSQL) Close() error {
	return p.db.Close()
}

func (p *PostgreSQL) GetAllCats() ([]DbCat, error) {
	rows, err := p.db.Query("SELECT * FROM cats")
	if err != nil {
		return []DbCat{}, err
	}
	defer rows.Close()

	var results []DbCat

	for rows.Next() {
		var c DbCat
		err := rows.Scan(&c.ID, &c.Name, &c.YearsOfExperience, &c.Breed, &c.Salary)
		if err != nil {
			return []DbCat{}, err
		}
		results = append(results, c)
	}

	return results, nil
}

func (p *PostgreSQL) GetSpecificCat(id string) (DbCat, error) {
	return DbCat{}, nil
}

func (p *PostgreSQL) CreateCat(c Cat) error {
	return nil
}

func (p *PostgreSQL) UpdateCat(id string, c Cat) error {
	return nil
}

func (p *PostgreSQL) DeleteCat(id string) error {
	return nil
}
