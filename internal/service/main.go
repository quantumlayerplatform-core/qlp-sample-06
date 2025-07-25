package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
)

const (
	port = ":50051"
	redisAddr = "localhost:6379"
	kafkaAddr = "localhost:9092"
)

type server struct {
	UnimplementedYourServiceServer
	redisClient *redis.Client
	kafkaWriter *kafka.Writer
}

func NewServer() *server {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	kw := &kafka.Writer{
		Addr: kafka.TCP(kafkaAddr),
		Topic: "your-topic",
	}
	return &server{
		redisClient: rdb,
		kafkaWriter: kw,
	}
}

func (s *server) YourMethod(ctx context.Context, in *YourRequest) (*YourResponse, error) {
	// Implement your gRPC method logic here
	return &YourResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterYourServiceServer(s, NewServer())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}