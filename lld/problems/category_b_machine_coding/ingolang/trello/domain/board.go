package domain

type BoardMetaData struct {
	Id     string
	Name   string
	Access string
	Url    string
}

type Board struct {
	MetaData     BoardMetaData
	Members      map[string]Member
	ProjectLists map[string]ProjectList
}

type IBoardUsecase interface {
	CreateBoard(metaData BoardMetaData) (string, error)
	DeleteBoard(boardID string) error
	UpdateBoardMetadata(boardID string, newMetadata BoardMetaData) error
	AddMemberToBoard(boardID string, memberID string) error
	RemoveMemberFromBoard(boardID string, memberID string) error
	AddProjectListToBoard(boardID string, projectListID string) error
	RemoveProjectListFromBoard(boardID string, projectListID string) error
}

type IBoardRepository interface {
	GetBoardByID(boardID string) (*Board, error)
	SaveBoard(boardID string, board Board) error
	DeleteBoard(boardID string) error
}
