package models

import (
	"errors"
	"fmt"
	"sync"
)

type KafkaStoreInterface interface {
	GetTopic(topicName string) (Topic, error)
	CreateTopic(topicName string, partition int) error
	PutRecord(message Message, topicName string, wg *sync.WaitGroup)
	GetRecords(consumerGroupId string, count int32, topicName string) ([]Message, error)
}

type KafkaStore struct {
	Topics         map[string]Topic              // topic name , topic
	Messages       map[string][]Message          // topic name , message
	ConsumerOffset map[string]map[Consumer]int32 // topic name, consumer id , offset
	sync.RWMutex
}

func NewKafkastore() *KafkaStore {
	return &KafkaStore{
		Topics:         make(map[string]Topic),
		Messages:       make(map[string][]Message),
		ConsumerOffset: make(map[string]map[Consumer]int32),
	}
}

func (ks *KafkaStore) CreateTopic(topicName string, partition int) error {

	if _, ok := ks.Topics[topicName]; ok {
		fmt.Printf("topic %s already exists", topicName)

		return nil
	}
	ks.Lock()
	defer ks.Unlock()
	ks.Topics[topicName] = Topic{
		TopicName:  topicName,
		Partitions: partition,
	}
	ks.Messages[topicName] = make([]Message, 0)
	ks.ConsumerOffset[topicName] = make(map[Consumer]int32)
	return nil
}

func (ks *KafkaStore) GetTopic(topicName string) (Topic, error) {
	defer ks.RUnlock()
	ks.RLock()
	topic, ok := ks.Topics[topicName]
	if !ok {
		return Topic{}, errors.New(fmt.Sprintf("topic %s not found", topicName))
	}
	return topic, nil
}

func (ks *KafkaStore) PutRecord(message Message, topicName string, wg *sync.WaitGroup) {

	defer ks.Unlock()
	ks.Lock()
	if messages, ok := ks.Messages[topicName]; ok {
		messages = append(messages, message)
		ks.Messages[topicName] = messages
	} else {
		messages = make([]Message, 0)
		messages = append(messages, message)
		ks.Messages[topicName] = messages
	}
	wg.Done()
}

func (ks *KafkaStore) GetRecords(consumerGroupId string, count int32, topicName string) ([]Message, error) {
	defer ks.RUnlock()
	ks.RLock()
	offset := ks.ConsumerOffset[topicName][Consumer{consumerGroupId}]
	maxOffset := int32(len(ks.Messages[topicName]))
	if ((maxOffset - offset + 1) < count) || offset == maxOffset {
		return nil, errors.New("not enough messages")
	}
	messages := ks.Messages[topicName][offset : offset+count]
	ks.ConsumerOffset[topicName][Consumer{consumerGroupId}] = offset + count
	return messages, nil
}

func (ks *KafkaStore) GetMessges(topicName string) ([]Message, error) {
	defer ks.RUnlock()
	ks.RLock()
	messages, ok := ks.Messages[topicName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("topic %s not found", topicName))
	}
	return messages, nil
}
