package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	pb "github.com/kailash-kangne/grpc-demo/proto"
)

const(
	port = ":8080"
)

type helloServer struct{
	pb.GreetServiceServer
}

func main(){

	lis, err := net.Listen("tcp",port)
	if err != nil {
		log.Fatalf("failed to start server %v",err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start : %v",err)
	}
 }