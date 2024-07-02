package domain

type ProjectList struct {
	Id    string
	Name  string
	Cards map[string]Card
}

type IProjectListRepository interface {
	GetProjectListByID(projectListID string) (*ProjectList, error)
	SaveProjectList(projectListID string, projectList ProjectList) error
	DeleteProjectList(projectListID string) error
	UpdateCard(projectListID string, cardID string) error
}
