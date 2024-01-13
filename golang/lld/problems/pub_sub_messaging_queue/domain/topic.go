package domain

type Topic struct {
	TopicId     string
	TopicName   string
	Subscribers map[string]Subscriber
	Messages    []string
}
