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

type DbTarget struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Country    string `json:"country"`
	Notes      string `json:"notes"`
	IsComplete bool   `json:"is_complete"`
}

type Target struct {
	Name       string `json:"name"`
	Country    string `json:"country"`
	Notes      string `json:"notes"`
	IsComplete bool   `json:"is_complete"`
}

type DbMission struct {
	ID         int    `json:"id"`
	AssigneeId int    `json:"assignee_id"`
	TargetIds  []int  `json:"target_ids"`
	Note       string `json:"note"`
	IsComplete bool   `json:"is_complete"`
}

type Mission struct {
	AssigneeId       string   `json:"assignee_id"`
	Targets          []Target `json:"targets"`
	Note             string   `json:"note"`
	IsComplete       bool     `json:"is_complete"`
	createdTargetIds []string
}

func (m *Mission) AddCreatedTargetId(id string) {
	m.createdTargetIds = append(m.createdTargetIds, id)
}

func (m *Mission) GetCreatedTargetIds() []string {
	return m.createdTargetIds
}

type MissionStorager interface {
	GetAllMissions() ([]DbMission, error)
	GetSpecificMission(id string) (DbMission, error)
	CreateMission(m Mission) error
	UpdateMission(m Mission) error
	DeleteMission(id string) error
	GetMissionTargets(id string) ([]DbTarget, error)
	CreateMissionTarget(Target) (targetId string, err error)
	GetTarget(id string) (DbTarget, error)
	AddTargetToMission(idTarget string, idMission string) error
	UpdateMissionTarget(id string, t Target) error
}

type DbStorager interface {
	CatStorager
	MissionStorager
}
