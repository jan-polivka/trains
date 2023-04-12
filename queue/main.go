package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "proto-api/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type dispatchServiceSever struct {
	pb.UnimplementedDispatchServiceServer
}

func (s *dispatchServiceSever) DispatchTrains(ctx context.Context, trains *pb.Trains) pb.DispatchAck {
	fmt.Println("what")
	return pb.DispatchAck{Response: 0}
}

func main() {
	fmt.Println("Hello World")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterDispatchServiceServer(grpcServer, pb.UnimplementedDispatchServiceServer{})
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}
