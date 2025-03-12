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
		err := rows.Scan(&c.ID, &c.Name, &c.Breed, &c.Salary, &c.YearsOfExperience)
		if err != nil {
			return []DbCat{}, err
		}
		results = append(results, c)
	}

	return results, nil
}

func (p *PostgreSQL) GetSpecificCat(id string) (DbCat, error) {
	row := p.db.QueryRow("SELECT * FROM cats WHERE id=?", id)

	if row.Err() != nil {
		return DbCat{}, row.Err()
	}

	var cat DbCat
	if err := row.Scan(&cat); err != nil {
		return DbCat{}, err
	}

	return cat, nil
}

func (p *PostgreSQL) CreateCat(c Cat) error {
	_, err := p.db.Exec("INSERT INTO cats(name, years_of_experience, breed, salary) VALUES ($1, $2, $3, $4)",
		c.Name, c.YearsOfExperience, c.Breed, c.Salary)
	return err
}

func (p *PostgreSQL) UpdateCat(id string, c Cat) error {
	_, err := p.db.Exec("UPDATE cats SET name=$1, years_of_experience=$2, breed=$3, salary=$4 WHERE id=$5",
		c.Name, c.YearsOfExperience, c.Breed, c.Salary, id)
	return err
}

func (p *PostgreSQL) DeleteCat(id string) error {
	_, err := p.db.Exec("DELETE FROM cats WHERE id=$1", id)
	return err
}
