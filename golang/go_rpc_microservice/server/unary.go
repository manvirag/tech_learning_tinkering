package main

import (
	"context"
	"fmt"

	pb "github.com/manvirag982/tech_learning_tinkering/golang/go_rpc_microservice/proto"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.HelloWorldServer
}

func (s *Server) Hello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	fmt.Print("hello world")
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
