package usecase

import (
	"fmt"
	"pub_sub_msg_queue/domain"
)

type ISubscriberTypeUsecase interface {
	SendNotification(topic *domain.Topic) error
}

type PubSubSubSubscriberUsecase struct {
}

func (psc *PubSubSubSubscriberUsecase) SendNotification(topic *domain.Topic) error {

	for subscriberId, subscriber := range topic.Subscribers {
		err := subscriber.Strategy.SendEvent(topic.Messages[subscriber.Offset])
		if err == nil {
			subscriber.Offset = subscriber.Offset + 1
			topic.Subscribers[subscriberId] = subscriber
		} else {
			fmt.Printf("Failed to notify subscriber %+v", subscriberId)
		}
	}
	return nil
}
