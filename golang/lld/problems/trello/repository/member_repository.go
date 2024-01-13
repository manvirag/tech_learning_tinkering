package repository

import (
	"errors"
	"trello/domain"
)

type MemberRepository struct {
	InMemoryMember map[string]domain.Member
}

func NewMemberRepository() *MemberRepository {
	return &MemberRepository{
		InMemoryMember: make(map[string]domain.Member),
	}
}

func (mr *MemberRepository) GetMemberById(id string) (error, domain.Member) {
	if member, ok := mr.InMemoryMember[id]; ok {
		return nil, member
	} else {
		return errors.New("error member with this id doesn't exist"), domain.Member{}
	}
}

func (mr *MemberRepository) UpdateMember(member domain.Member) {
	mr.InMemoryMember[member.Id] = member
}
