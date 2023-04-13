package sender

import (
	"context"
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "proto-api/protos"
)

func Test_DispatchTrainsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("integration test")
	}
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

	res := DispatchTrains(client, ctx)
	assert.Equal(t, int64(0), res)
}
