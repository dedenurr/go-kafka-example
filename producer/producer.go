package main

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "topic_test"
	partition := 0
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

	conn.WriteMessages(kafka.Message{Value: []byte("Hello Kafka again dede")})
}