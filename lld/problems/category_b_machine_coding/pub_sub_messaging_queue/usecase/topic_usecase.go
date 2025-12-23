package usecase

import (
	"pub_sub_msg_queue/domain"
	"pub_sub_msg_queue/repository"
)

type ITopicUsecase interface {
	PublishMessage(topicId string, message string) error
	UnsubscribeTopic(topicId string, subscriberId string) error
	SubscribeTopic(topicId string, subscriber domain.Subscriber) error
}

type TopicUsecase struct {
	TopicRepo             *repository.TopicRepository
	PubSubSubSubscriberUC *PubSubSubSubscriberUsecase
}

func NewTopicUsecase(topicRepo *repository.TopicRepository, pubSubSubSubscriberUC *PubSubSubSubscriberUsecase) *TopicUsecase {
	return &TopicUsecase{
		TopicRepo:             topicRepo,
		PubSubSubSubscriberUC: pubSubSubSubscriberUC,
	}
}

func (uc *TopicUsecase) PublishMessage(topicId string, message string) error {
	topic, err := uc.TopicRepo.GetTopicById(topicId)
	if err != nil {
		return err
	}

	topic.Messages = append(topic.Messages, message)
	errN := uc.PubSubSubSubscriberUC.SendNotification(&topic)
	if errN != nil {
		return errN
	}

	uc.TopicRepo.UpdateTopicById(topic)
	return nil
}

func (uc *TopicUsecase) UnsubscribeTopic(topicId string, subscriberId string) error {
	topic, err := uc.TopicRepo.GetTopicById(topicId)
	if err != nil {
		return err
	}
	delete(topic.Subscribers, subscriberId)
	uc.TopicRepo.UpdateTopicById(topic)

	return nil
}

func (uc *TopicUsecase) SubscribeTopic(topicId string, subscriber domain.Subscriber) error {
	topic, err := uc.TopicRepo.GetTopicById(topicId)
	if err != nil {
		return err
	}

	topic.Subscribers[subscriber.SubscriberId] = subscriber

	uc.TopicRepo.UpdateTopicById(topic)
	return nil
}
