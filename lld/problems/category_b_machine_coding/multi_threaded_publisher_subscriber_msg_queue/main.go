package main

import (
	"fmt"
	"main/models"
	"main/usecases"
	"time"
)

func main() {
	topicName := "topic_main"
	partition := 4
	kafkaRepo := models.NewKafkastore()
	kafkaRepo.CreateTopic(topicName, partition)
	topic, errT := kafkaRepo.GetTopic(topicName)
	if errT == nil {
		fmt.Printf("%+v\n\n", topic)
	}

	processor := usecases.NewProcessor(topicName, kafkaRepo)
	consumerGroup := "consumer_group_main"

	messages := make([]models.Message, 0)
	messages = append(messages, models.Message{Msg: "message_1", Key: "key_1", Timestamp: time.Now().Unix()})
	err := processor.PutRecords(messages)

	if err != nil {
		fmt.Printf("error while putting records %s", err.Error())
	}

	processor.SubscribeTopic(models.Consumer{ConsumerGroupId: consumerGroup})

	res, err := processor.GetRecords(consumerGroup, 1)
	if err != nil {
		fmt.Printf("Error: Got records : %+v", err)
	} else {
		fmt.Printf("Got records: %+v", res)
	}

	res2, err2 := processor.GetRecords(consumerGroup, 1)
	if err2 != nil {
		fmt.Printf("\n\nError: Got records 2 : %+v", err2)
	} else {
		fmt.Printf("Got records: 2  %+v", res2)
	}

}
