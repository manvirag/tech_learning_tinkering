package main

import (
	"context"
	"log"
	"time"

	pb "github.com/manvirag982/tech_learning_tinkering/golang/go_rpc_microservice/proto"
)

func Hello(client pb.HelloWorldClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Hello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not get reponse: %v", err)
	}
	log.Printf("%s", res.Message)
}
