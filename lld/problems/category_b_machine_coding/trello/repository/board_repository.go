package repository

import (
	"errors"
	"trello/domain"
)

type BoardRepository struct {
	boards map[string]domain.Board
}

func NewBoardRepository() *BoardRepository {
	return &BoardRepository{
		boards: make(map[string]domain.Board),
	}
}

func (repo *BoardRepository) GetBoardByID(boardID string) (domain.Board, error) {
	board, ok := repo.boards[boardID]
	if !ok {
		return domain.Board{}, errors.New("board not found")
	}
	return board, nil
}

func (repo *BoardRepository) SaveBoard(boardID string, board domain.Board) error {
	repo.boards[boardID] = board
	return nil
}

func (repo *BoardRepository) DeleteBoard(boardID string) error {
	if _, ok := repo.boards[boardID]; !ok {
		return errors.New("board not found")
	}
	delete(repo.boards, boardID)
	return nil
}
