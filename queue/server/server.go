package server

import (
	"context"
	"fmt"

	pb "proto-api/protos"
)

type queueServer struct {
}

func test() {
	fmt.Println()
}

func (s *queueServer) GetDrive(ctx context.Context, drive *pb.Trains) *pb.DispatchAck {
	return &pb.DispatchAck{Response: 0}
}
