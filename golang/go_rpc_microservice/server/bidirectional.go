package main

import (
	pb "github.com/manvirag982/tech_learning_tinkering/golang/go_rpc_microservice/proto"
	"io"
	"log"
)

func (s *Server) HelloBidirectionalStreaming(stream pb.HelloWorld_HelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name : %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
