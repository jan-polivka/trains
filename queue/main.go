package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/jan-polivka/trains/proto-api/protos"
	"github.com/jan-polivka/trains/queue/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Hello World")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterQueueServiceServer(grpcServer, &server.QueueServiceServer{})
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}
