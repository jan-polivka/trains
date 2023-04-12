package main

import (
	"fmt"
	"log"
	"net"

	pb "proto-api/protos"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello World")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterDispatchServiceServer(grpcServer, pb.UnimplementedDispatchServiceServer{})
	grpcServer.Serve(lis)
}
