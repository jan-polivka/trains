package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/jan-polivka/trains/proto-api/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type dispatchServiceSever struct {
	pb.UnimplementedQueueServiceServer
}

func (s *dispatchServiceSever) DispatchTrains(ctx context.Context, trains *pb.Trains) (*pb.DispatchAck, error) {
	fmt.Println("what")
	return &pb.DispatchAck{Response: 0}, nil
}

func main() {
	fmt.Println("Hello World")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterQueueServiceServer(grpcServer, &dispatchServiceSever{})
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}
