package service

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/segmentio/kafka-go"
)

type UserService struct {
	redisClient *redis.Client
	kafkaWriter *kafka.Writer
}

func NewUserService(redisClient *redis.Client, kafkaWriter *kafka.Writer) *UserService {
	return &UserService{
		redisClient: redisClient,
		kafkaWriter: kafkaWriter,
	}
}

func (s *UserService) CreateUser(ctx context.Context, username string, email string) error {
	if err := s.validateUserInput(username, email); err != nil {
		return status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	if err := s.redisClient.Set(ctx, "user:"+username, email, 0).Err(); err != nil {
		return status.Errorf(codes.Internal, "failed to save user to Redis: %v", err)
	}

	message := kafka.Message{
		Key:   []byte(username),
		Value: []byte(email),
	}
	if err := s.kafkaWriter.WriteMessages(ctx, message); err != nil {
		return status.Errorf(codes.Internal, "failed to send user creation event to Kafka: %v", err)
	}

	return nil
}

func (s *UserService) validateUserInput(username string, email string) error {
	if username == "" || email == "" {
		return status.Error(codes.InvalidArgument, "username and email cannot be empty")
	}
	return nil
}