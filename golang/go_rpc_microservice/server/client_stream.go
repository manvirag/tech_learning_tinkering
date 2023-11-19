package main

import (
	pb "github.com/manvirag982/tech_learning_tinkering/golang/go_rpc_microservice/proto"
	"io"
	"log"
)

func (s *Server) HelloClientStreaming(stream pb.HelloWorld_HelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name : %v", req.Name)
		messages = append(messages, "Hello "+req.Name)
	}
}
