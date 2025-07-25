package components

import (
	"context"
	"log"

	pb "path/to/project/grpc/protos"
	"path/to/project/frontend/store"
)

// UserComponent handles the user-related UI logic
type UserComponent struct {
	GrpcClient *store.GRPCClient
}

// NewUserComponent creates a new instance of UserComponent
func NewUserComponent(client *store.GRPCClient) *UserComponent {
	return &UserComponent{
		GrpcClient: client,
	}
}

// FetchUser fetches user data from the server
func (uc *UserComponent) FetchUser(userID string) (*pb.UserResponse, error) {
	client := pb.NewUserServiceClient(uc.GrpcClient.Connection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.GetUser(ctx, &pb.UserRequest{Id: userID})
	if err != nil {
		log.Printf("Failed to fetch user: %v", err)
		return nil, err
	}
	return response, nil
}