package main 

import (
	"io"
	"log"
	pb "github.com/kailash-kangne/grpc-demo/proto"
)

func (s *helloServer) SayHelloBidirecationalStreaming(stream pb.GreetService_SayHelloBidirecationalStreamingServer) error{

	for {
		req, err := stream.Recv()
		
		if err == io.EOF{
			return nil
		}
		if err != nil{
			return err
		}
		log.Printf("got request with name: %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello "+req.Name,
		}
		if err := stream.Send(res); err != nil{
			return err
		}
	}
}