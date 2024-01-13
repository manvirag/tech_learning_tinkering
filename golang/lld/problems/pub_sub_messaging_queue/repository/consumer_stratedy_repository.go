package repository

import "fmt"

type HttpConsumerStrategyRepo struct {
	EndPointUrl string
}

func (http *HttpConsumerStrategyRepo) SendEvent(message string) error {
	fmt.Printf("Successfully notify the subscriber about message: %+v , at end point :%+v", message, http.EndPointUrl)
	return nil
}
