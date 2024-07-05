package usecases

import (
	"errors"
	"fmt"
	"main/models"
	"sync"
)

type ProcessorInterface interface {
	SubscribeTopic(topicName string) error
	PutRecords(messages []models.Message) error
	GetRecords(consumerGroupId string, consumerId string, count int32) ([]models.Message, error)
}

// for a single topic push the message and invoke subscriber
type Processor struct {
	TopicName   string
	KafkaRepo   *models.KafkaStore
	Subscribers map[models.Consumer]int32
	sync.RWMutex
	sync.WaitGroup
}

func NewProcessor(topicName string, kafkaRepo *models.KafkaStore) *Processor {
	return &Processor{
		TopicName:   topicName,
		KafkaRepo:   kafkaRepo,
		Subscribers: make(map[models.Consumer]int32),
	}
}

func (p *Processor) SubscribeTopic(consumer models.Consumer) error {
	topic, err := p.KafkaRepo.GetTopic(p.TopicName)
	if err != nil {
		return err
	}
	defer p.Unlock()
	p.Lock()

	if _, ok := p.Subscribers[consumer]; ok {

		if topic.Partitions == int(p.Subscribers[consumer]) {
			err := errors.New(fmt.Sprintf("topic %s has %d partitions, and all are filled", p.TopicName, topic.Partitions))
			return err
		} else {
			p.Subscribers[consumer] = p.Subscribers[consumer] + 1
		}
	} else {
		p.Subscribers[consumer] = 1
	}
	return nil

}

func (p *Processor) PutRecords(messages []models.Message) error {
	_, err := p.KafkaRepo.GetTopic(p.TopicName)
	if err != nil {
		return err
	}
	defer p.Wait()

	for _, message := range messages {
		p.Add(1)
		go p.KafkaRepo.PutRecord(message, p.TopicName, &p.WaitGroup)
	}

	return nil
}

func (p *Processor) GetRecords(consumerGroupId string, count int32) ([]models.Message, error) {
	messages, err := p.KafkaRepo.GetRecords(consumerGroupId, count, p.TopicName)

	if err != nil {
		return nil, err
	}
	return messages, nil
}
