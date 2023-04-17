package sender

import (
	"context"
	"fmt"

	pb "github.com/jan-polivka/trains/proto-api/protos"
)

func DispatchTrains(client pb.QueueServiceClient, ctx context.Context) int64 {

	resp, err := client.DispatchTrains(ctx, &pb.Trains{Train: []*pb.Train{{Name: "RegioSBahn", Drive: "ElectricDrive"}}})
	if err != nil {
		fmt.Println("the response is fucked")
	}
	return resp.Response
}
