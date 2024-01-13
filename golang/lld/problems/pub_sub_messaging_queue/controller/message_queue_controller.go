package controller

type IMessageQueueController interface {
	PublishMessage(topicId string, message string) error
	UnsubscribeTopic(topicId string, subscriberId string) error
	SubscribeTopi(topicId string, subscriberId string) error
}
