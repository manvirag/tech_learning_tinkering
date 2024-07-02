package domain

type ConsumerStrategyRepo interface {
	SendEvent(message string) error
}
