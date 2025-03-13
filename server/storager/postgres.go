package storager

import (
	"database/sql"
	"strconv"

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

func (p *PostgreSQL) GetAllMissions() ([]DbMission, error) {
	rows, err := p.db.Query("SELECT * FROM missions")
	if err != nil {
		return []DbMission{}, err
	}
	defer rows.Close()

	var results []DbMission

	for rows.Next() {
		var m DbMission
		err := rows.Scan(&m.ID, &m.AssigneeId, &m.TargetIds, &m.Note, &m.IsComplete)
		if err != nil {
			return []DbMission{}, err
		}
		results = append(results, m)
	}

	return results, nil
}

func (p *PostgreSQL) GetSpecificMission(id string) (DbMission, error) {
	row := p.db.QueryRow("SELECT * FROM missions WHERE id=?", id)

	if row.Err() != nil {
		return DbMission{}, row.Err()
	}

	var mission DbMission
	if err := row.Scan(&mission); err != nil {
		return DbMission{}, err
	}

	return mission, nil
}

func (p *PostgreSQL) CreateMission(m Mission) error {
	_, err := p.db.Exec("INSERT INTO missions(assignee_id, target_ids, note, is_complete) VALUES ($1, $2, $3, $4)",
		m.AssigneeId, m.GetCreatedTargetIds(), m.Note, m.IsComplete)
	return err
}
func (p *PostgreSQL) UpdateMission(id string, m Mission) error {
	_, err := p.db.Exec("UPDATE missions SET assignee_id=$1, target_ids=$2, note=$3, is_complete=$4 WHERE id=$5",
		m.AssigneeId, m.GetCreatedTargetIds(), m.Note, m.IsComplete, id)
	return err
}

func (p *PostgreSQL) DeleteMission(id string) error {
	_, err := p.db.Exec("DELETE FROM missions WHERE id=$1", id)
	return err
}

func (p *PostgreSQL) GetMissionTargets(id string) ([]DbTarget, error) {
	rows, err := p.db.Query(`
	SELECT t.*
	FROM targets t
	JOIN missions m ON m.target_ids @> ARRAY[t.id]
	WHERE m.id = 1;
	`)
	if err != nil {
		return []DbTarget{}, err
	}
	defer rows.Close()

	var results []DbTarget

	for rows.Next() {
		var t DbTarget
		err := rows.Scan(&t.ID, &t.Name, &t.Country, &t.Notes, &t.IsComplete)
		if err != nil {
			return []DbTarget{}, err
		}
		results = append(results, t)
	}

	return results, nil
}

func (p *PostgreSQL) CreateMissionTarget(t Target) (string, error) {
	res, err := p.db.Exec("INSERT INTO targets(name, country, is_complete) VALUES ($1, $2, $3)",
		t.Name, t.Country, t.IsComplete)
	if err != nil {
		return "", err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return "", err
	}

	return strconv.Itoa(int(lastId)), nil
}

func (p *PostgreSQL) GetTarget(id string) (DbTarget, error) {
	row := p.db.QueryRow("SELECT * FROM targets WHERE id=?", id)

	if row.Err() != nil {
		return DbTarget{}, row.Err()
	}

	var target DbTarget
	if err := row.Scan(&target); err != nil {
		return DbTarget{}, err
	}

	return target, nil
}

func (p *PostgreSQL) AddTargetToMission(idTarget string, idMission string) error {
	_, err := p.db.Exec(`
	UPDATE missions
	SET target_ids = array_append(target_ids, $1)
	WHERE id = $2;
	`, idTarget, idMission)
	return err
}

func (p *PostgreSQL) UpdateMissionTarget(id string, t Target) error {
	return nil
}
