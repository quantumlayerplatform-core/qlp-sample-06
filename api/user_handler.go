package handler

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"

	pb "path/to/your/protofiles"
	"path/to/your/service"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	userService *service.UserService
}

func NewServer(redisClient *redis.Client, kafkaWriter *kafka.Writer) *Server {
	return &Server{
		userService: service.NewUserService(redisClient, kafkaWriter),
	}
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	err := s.userService.CreateUser(ctx, req.GetUsername(), req.GetEmail())
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{Success: true}, nil
}

func StartGRPCServer(address string, redisClient *redis.Client, kafkaWriter *kafka.Writer) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, NewServer(redisClient, kafkaWriter))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}