package domain

type Card struct {
	Id             string
	Name           string
	Description    string
	AssignedUserId Member
}

type ICardRepository interface {
	GetCardById(id string) (error, Card)
	UpdateCard(card Card)
}

type ICardUsecase interface {
	GetCardById(id string) (error, Card)
}
