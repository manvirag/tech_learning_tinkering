package repository

import (
	"errors"
	"fmt"
	"trello/domain"
)

type CardRepository struct {
	InMemoryCard map[string]domain.Card
}

func NewCardRepository() *CardRepository {
	return &CardRepository{
		InMemoryCard: make(map[string]domain.Card),
	}
}

func (cr *CardRepository) GetCardById(id string) (error, domain.Card) {
	if card, ok := cr.InMemoryCard[id]; ok {
		return nil, card
	} else {
		return errors.New("error card with this id doesn't exist"), domain.Card{}
	}
}

func (cr *CardRepository) UpdateCard(card domain.Card) {
	cr.InMemoryCard[card.Id] = card
}

func (cr *CardRepository) DeleteCard(cardId string) {
	fmt.Println("DeleteCard")
	fmt.Println(cr.InMemoryCard)
	fmt.Println(cardId)
	if _, ok := cr.InMemoryCard[cardId]; ok {
		delete(cr.InMemoryCard, cardId)
	}
}
