package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Hello World")
	_, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
