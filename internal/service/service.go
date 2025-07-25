package main

import (
	"context"

	pb "path/to/your/protofiles"
)

type server struct {
	pb.UnimplementedYourServiceServer
}

func (s *server) YourServiceMethod(ctx context.Context, req *pb.YourRequest) (*pb.YourResponse, error) {
	// Implement your service logic here
	return &pb.YourResponse{}, nil
}