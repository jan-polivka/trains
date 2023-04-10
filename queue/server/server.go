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

func (s *queueServer) GetDrive(ctx context.Context, drive *pb.Trains) {
	fmt.Println("test")
}
