package domain

type Member struct {
	Id    string
	Email string
	Name  string
}

type IMemberRepository interface {
	GetMemberById(id string) (error, Member)
	UpdateMember(member Member)
}
