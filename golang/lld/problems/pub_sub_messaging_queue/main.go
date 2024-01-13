package main

import (
	"fmt"
	"pub_sub_msg_queue/domain"
	"pub_sub_msg_queue/repository"
	"pub_sub_msg_queue/usecase"
)

func main() {
	httpStrategy := &repository.HttpConsumerStrategyRepo{EndPointUrl: "http://example.com"}

	subscriber := domain.Subscriber{
		SubscriberId: "123",
		Offset:       0,
		Strategy:     httpStrategy,
	}

	topic := domain.Topic{
		TopicId:     "topic123",
		TopicName:   "Example Topic",
		Subscribers: map[string]domain.Subscriber{"123": subscriber},
		Messages:    []string{"Message 1", "Message 2", "Message 3"},
	}

	topicUsecase := &usecase.TopicUsecase{
		TopicRepo: &repository.TopicRepository{
			InMemoryData: map[string]domain.Topic{topic.TopicId: topic},
		},
		PubSubSubSubscriberUC: &usecase.PubSubSubSubscriberUsecase{},
	}

	err := topicUsecase.PublishMessage("topic123", "new message")
	if err != nil {
		fmt.Println("Error sending notification:", err)
		return
	}
	fmt.Println("Notification sent successfully!")
}
