package main

import (
	"context"
	"log"
	"time"

	pb "github.com/manvirag982/tech_learning_tinkering/golang/go_rpc_microservice/proto"
)

func HelloClientStream(client pb.HelloWorldClient, names *pb.NamesList) {
	log.Printf("Client Streaming started")
	stream, err := client.HelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sent request with name: %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client Streaming finished")
	if err != nil {
		log.Fatalf("Error while receiving %v", err)
	}
	log.Printf("%v", res.Messages)
}
