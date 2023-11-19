package main

import (
	"log"

	pb "github.com/manvirag982/tech_learning_tinkering/golang/go_rpc_microservice/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloWorldClient(conn)
	Hello(client)
}
