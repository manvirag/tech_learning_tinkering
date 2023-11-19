package main

import (
	pb "github.com/manvirag982/tech_learning_tinkering/golang/go_rpc_microservice/proto"
	"log"
	"time"
)

func (s *Server) HelloServerStreaming(req *pb.NamesList, stream pb.HelloWorld_HelloServerStreamingServer) error {
	log.Printf("Got request with names : %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		// 2 second delay to simulate a long running process
		time.Sleep(2 * time.Second)
	}
	return nil
}
