package server

import (
	"context"

	pb "github.com/jan-polivka/trains/proto-api/protos"
)

type queueServiceServer struct {
	pb.UnimplementedQueueServiceServer
}

func (s *queueServiceServer) DispatchTrains(ctx context.Context, trains *pb.Trains) (*pb.DispatchAck, error) {
	return &pb.DispatchAck{Response: 0}, nil
}
