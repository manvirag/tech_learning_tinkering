package usecase

import (
	"errors"
	"trello/domain"
	"trello/repository"
	"trello/utils"
)

// BoardUseCase implements the IBoardUsecase interface.
type BoardUseCase struct {
	MemberRepo      *repository.MemberRepository
	BoardRepo       *repository.BoardRepository
	ProjectListRepo *repository.ProjectListRepository
}

func (uc *BoardUseCase) CreateBoard(metaData domain.BoardMetaData) (string, error) {
	boardID := utils.GenerateID()
	if _, err := uc.GetBoardByID(boardID); err == nil {
		return "", errors.New("board with the same ID already exists. Try again in some time")
	}
	metaData.Id = boardID
	newBoard := domain.Board{
		MetaData:     metaData,
		Members:      make(map[string]domain.Member),
		ProjectLists: make(map[string]domain.ProjectList),
	}

	if err := uc.SaveBoard(boardID, newBoard); err != nil {
		return "", err
	}

	return boardID, nil
}

func (uc *BoardUseCase) DeleteBoard(boardID string) error {
	board, err := uc.GetBoardByID(boardID)
	if err != nil {
		return errors.New("board not found")
	}
	if err := uc.BoardRepo.DeleteBoard(boardID); err != nil {
		return err
	}
	for projectListID := range board.ProjectLists {
		if err := uc.ProjectListRepo.DeleteProjectList(projectListID); err != nil {
			return errors.New("found error while deleting board")
		}
	}
	return nil
}

func (uc *BoardUseCase) UpdateBoardMetadata(boardID string, newMetadata domain.BoardMetaData) error {
	board, err := uc.GetBoardByID(boardID)
	if err != nil {
		return errors.New("board not found")
	}
	board.MetaData = newMetadata
	if err := uc.SaveBoard(boardID, *board); err != nil {
		return err
	}
	return nil
}

func (uc *BoardUseCase) AddMemberToBoard(boardID string, memberID string) error {
	board, err := uc.GetBoardByID(boardID)
	if err != nil {
		return errors.New("board not found")
	}
	err, member := uc.MemberRepo.GetMemberById(memberID)
	if err != nil {
		return errors.New("member not found")
	}

	board.Members[memberID] = member
	if err := uc.SaveBoard(boardID, *board); err != nil {
		return err
	}
	return nil
}

func (uc *BoardUseCase) RemoveMemberFromBoard(boardID string, memberID string) error {
	board, err := uc.GetBoardByID(boardID)
	if err != nil {
		return errors.New("board not found")
	}

	if err, _ := uc.MemberRepo.GetMemberById(memberID); err != nil {
		return errors.New("member not found")
	}

	delete(board.Members, memberID)

	if err := uc.SaveBoard(boardID, *board); err != nil {
		return err
	}

	return nil
}

func (uc *BoardUseCase) AddProjectListToBoard(boardID string, projectListID string) error {
	board, err := uc.GetBoardByID(boardID)
	if err != nil {
		return errors.New("board not found")
	}

	projectList, err := uc.ProjectListRepo.GetProjectListByID(projectListID)
	if err != nil {
		return errors.New("project list not found")
	}

	board.ProjectLists[projectListID] = *projectList

	if err := uc.SaveBoard(boardID, *board); err != nil {
		return err
	}
	return nil
}

func (uc *BoardUseCase) RemoveProjectListFromBoard(boardID string, projectListID string) error {
	board, err := uc.GetBoardByID(boardID)
	if err != nil {
		return errors.New("board not found")
	}

	if _, err := uc.ProjectListRepo.GetProjectListByID(projectListID); err != nil {
		return errors.New("project list not found")
	}

	delete(board.ProjectLists, projectListID)

	if err := uc.SaveBoard(boardID, *board); err != nil {
		return err
	}

	return nil
}

func (uc *BoardUseCase) GetBoardByID(boardID string) (*domain.Board, error) {
	board, ok := uc.BoardRepo.GetBoardByID(boardID)
	if ok != nil {
		return nil, errors.New("board not found")
	}
	return &board, nil
}

func (uc *BoardUseCase) SaveBoard(boardID string, board domain.Board) error {
	return uc.BoardRepo.SaveBoard(boardID, board)
}
