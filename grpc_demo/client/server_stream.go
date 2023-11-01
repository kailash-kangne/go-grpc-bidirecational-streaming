package main

import (
	"context"
	"log"
	"io"

	pb "github.com/kailash-kangne/grpc-demo/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NameList){
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)

	if err != nil{
		log.Fatalf("could not send names: %v",err)
	}

	for {
		message , err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Fatalf("error while streaming %v",err)
		}
		log.Println(message)	
	} 
	log.Println("Streaming finished")
}