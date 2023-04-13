package main

import (
	"context"
	"dispatcher/sender"
	"flag"
	"fmt"
	"time"

	pb "proto-api/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var serverAddr = flag.String("addr", "localhost:8080", "The server address in the format of host:port")
	var dialOption = grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(*serverAddr, dialOption)
	if err != nil {
		fmt.Println(err)
		fmt.Println("something went fucky wucky")
	}
	defer conn.Close()
	client := pb.NewDispatchServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res := sender.DispatchTrains(client, ctx)
	fmt.Println(res)

}
