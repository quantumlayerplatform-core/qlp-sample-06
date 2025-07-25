```go
package test

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "gen-bef1314b-eae5-4ec5-8c2b-8bc9f1d6db06/proto"
)

const (
	serverAddress = "localhost:50051"
	testDuration  = 10 * time.Second
	requestsPerSecond = 100
)

func TestLoad(t *testing.T) {
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewYourServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), testDuration)
	defer cancel()

	var wg sync.WaitGroup
	tick := time.NewTicker(time.Second / requestsPerSecond)
	defer tick.Stop()

	startTime := time.Now()
	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			fmt.Printf("Load test completed in %v\n", time.Since(startTime))
			return
		case <-tick.C:
			wg.Add(1)
			go func() {
				defer wg.Done()
				// Replace with actual request data
				req := &pb.YourRequest{}
				if _, err := client.YourRpcMethod(ctx, req); err != nil {
					log.Printf("Failed to send request: %v", err)
				}
			}()
		}
	}
}
```