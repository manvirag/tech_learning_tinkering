package main

import (
	"fmt"
	"trello/domain"
	"trello/repository"
	"trello/usecase"
	"trello/utils"
)

func main() {
	cardRepository := repository.NewCardRepository()
	projectListRepository := repository.NewProjectListRepository(cardRepository)
	memberRepository := repository.NewMemberRepository()
	boardRepository := repository.NewBoardRepository()

	boardUseCase := &usecase.BoardUseCase{
		ProjectListRepo: projectListRepository,
		MemberRepo:      memberRepository,
		BoardRepo:       boardRepository,
	}

	memberID := utils.GenerateID()
	member := domain.Member{Id: memberID, Name: "John Doe"}
	memberRepository.UpdateMember(member)

	projectListID := utils.GenerateID()
	projectList := domain.ProjectList{Id: projectListID, Name: "To Do List", Cards: make(map[string]domain.Card)}
	_ = projectListRepository.SaveProjectList(projectListID, projectList)

	boardMetaData := domain.BoardMetaData{Name: "Project Board", Access: "Private", Url: "https://example.com/board"}
	boardID, _ := boardUseCase.CreateBoard(boardMetaData)

	_ = boardUseCase.AddMemberToBoard(boardID, memberID)

	_ = boardUseCase.AddProjectListToBoard(boardID, projectListID)

	cardID := utils.GenerateID()
	card := domain.Card{Id: cardID, Name: "Task 1"}
	cardRepository.UpdateCard(card)

	_ = projectListRepository.UpdateCard(projectListID, cardID, card)

	retrievedBoard, err := boardUseCase.GetBoardByID(boardID)

	if err != nil {
		fmt.Printf("error in getting board: %+v", err)
	}

	fmt.Println(retrievedBoard)

	_ = boardUseCase.DeleteBoard(boardID)

}
