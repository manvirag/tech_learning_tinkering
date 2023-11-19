package main

import (
	"context"
	pb "github.com/manvirag982/tech_learning_tinkering/golang/go_rpc_microservice/proto"
	"io"
	"log"
)

func HelloServerStream(client pb.HelloWorldClient, names *pb.NamesList) {
	log.Printf("Client Streaming started")
	stream, err := client.HelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming %v", err)
		}
		log.Println(message)
	}

	log.Printf("Streaming finished")
}
