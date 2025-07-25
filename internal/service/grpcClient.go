package store

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GRPCClient holds the client connection and provides methods to interact with gRPC services
type GRPCClient struct {
	Connection *grpc.ClientConn
}

// NewGRPCClient creates a new gRPC client connection
func NewGRPCClient(serverAddr string) *GRPCClient {
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	return &GRPCClient{
		Connection: conn,
	}
}

// Close closes the gRPC client connection
func (client *GRPCClient) Close() {
	if err := client.Connection.Close(); err != nil {
		log.Printf("Error closing gRPC connection: %v", err)
	}
}