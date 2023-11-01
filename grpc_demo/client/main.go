package main

import (
	
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/kailash-kangne/grpc-demo/proto"
)

const(
	port = ":8080"
)

func main(){

	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("did not connect: %v ",err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names:= &pb.NameList{
		Names: []string{"bob","kailu","alice"},
	}

	//callSayHello(client) //unary api

	//callSayHelloServerStreaming(client,names)

	//callSayHelloClientStream(client,names)

	callHelloBidirecationalStream(client,names)


}