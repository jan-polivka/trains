package server

import (
	"context"
	"fmt"

	pb "github.com/jan-polivka/trains/proto-api/protos"
)

type queueServer struct {
}

func test() {
	fmt.Println()
}

func (s *queueServer) GetTrain(ctx context.Context, drive *pb.Trains) *pb.DispatchAck {
	return &pb.DispatchAck{Response: 0}
}
