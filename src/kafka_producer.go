package main

import (
	"context"
	"github.com/segmentio/kafka-go"
)

func NewKafkaWriter(addr, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr: kafka.TCP(addr),
		Topic: topic,
	}
}

func SendMessage(ctx context.Context, writer *kafka.Writer, key, message []byte) error {
	msg := kafka.Message{
		Key: key,
		Value: message,
	}
	return writer.WriteMessages(ctx, msg)
}