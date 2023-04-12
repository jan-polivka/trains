package main

import (
	"flag"
	"fmt"

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
	// client := pb.
}
