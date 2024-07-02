package repository

import (
	"errors"
	"pub_sub_msg_queue/domain"
)

type TopicRepository struct {
	InMemoryData map[string]domain.Topic
}

func (tr *TopicRepository) GetTopicById(topicId string) (domain.Topic, error) {
	if _, ok := tr.InMemoryData[topicId]; ok {
		return tr.InMemoryData[topicId], nil
	} else {
		return domain.Topic{}, errors.New("not found this topic id please create")
	}
}

func (tr *TopicRepository) UpdateTopicById(topic domain.Topic) {
	tr.InMemoryData[topic.TopicId] = topic
}
