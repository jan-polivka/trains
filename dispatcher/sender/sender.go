package sender

import (
	"context"
	"fmt"
	pb "proto-api/protos"
)

func DispatchTrains(client pb.DispatchServiceClient, ctx context.Context) int64 {

	resp, err := client.DispatchTrains(ctx, &pb.Trains{Train: []*pb.Train{{Name: "RegioSBahn", Drive: "ElectricDrive"}}})
	if err != nil {
		fmt.Println("the response is fucked")
	}
	return resp.Response
}
