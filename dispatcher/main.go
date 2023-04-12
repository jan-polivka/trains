package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	pb "proto-api/protos"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello World")
	var opts []grpc.DialOption
	var serverAddr = flag.String("addr", "localhost:8080", "The server address in the format of host:port")
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		fmt.Println("something went fucky wucky")
	}
	defer conn.Close()
	client := pb.NewDispatchServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.DispatchTrains(ctx, &pb.Trains{Train: []*pb.Train{&pb.Train{Name: "RegioSBahn", Drive: "ElectricDrive"}}})
	if err != nil {
		fmt.Println("the response is fucked")
	}
	fmt.Println(resp)

}
