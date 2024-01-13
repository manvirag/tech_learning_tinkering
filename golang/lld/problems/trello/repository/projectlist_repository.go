package repository

import (
	"errors"
	"fmt"
	"trello/domain"
)

type ProjectListRepository struct {
	projectLists map[string]domain.ProjectList

	CardRepo *CardRepository
}

func NewProjectListRepository(cardRepository *CardRepository) *ProjectListRepository {
	return &ProjectListRepository{
		projectLists: make(map[string]domain.ProjectList),
		CardRepo:     cardRepository,
	}
}

func (repo *ProjectListRepository) GetProjectListByID(projectListID string) (*domain.ProjectList, error) {
	projectList, ok := repo.projectLists[projectListID]
	if !ok {
		return nil, errors.New("project list not found")
	}
	return &projectList, nil
}

func (repo *ProjectListRepository) SaveProjectList(projectListID string, projectList domain.ProjectList) error {
	repo.projectLists[projectListID] = projectList
	return nil
}

func (repo *ProjectListRepository) DeleteProjectList(projectListID string) error {
	if _, ok := repo.projectLists[projectListID]; !ok {
		return errors.New("project list not found")
	}
	projectList, _ := repo.projectLists[projectListID]
	for cardID := range projectList.Cards {
		fmt.Println("cardID")

		fmt.Println(cardID)
		repo.CardRepo.DeleteCard(cardID)
	}

	delete(repo.projectLists, projectListID)
	return nil
}
func (repo *ProjectListRepository) UpdateCard(projectListID string, cardID string, updatedCard domain.Card) error {
	projectList, ok := repo.projectLists[projectListID]
	if !ok {
		return errors.New("project list not found")
	}
	projectList.Cards[cardID] = updatedCard
	repo.projectLists[projectListID] = projectList
	fmt.Println(repo.projectLists[projectListID])
	return nil
}
