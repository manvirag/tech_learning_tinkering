package domain

type Subscriber struct {
	SubscriberId string
	Offset       int
	Strategy     ConsumerStrategyRepo
}
