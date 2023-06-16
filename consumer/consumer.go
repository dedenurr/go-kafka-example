package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "topic_test"
	partition := 0
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	conn.SetReadDeadline(time.Now().Add(time.Second * 8))

	//read one message
/* 	maxByte := 1e3
	message, _ := conn.ReadMessage(int(maxByte))
	fmt.Println(string(message.Value)) */

	// read all message
	minByte := 1e3
	maxByte := 1e9
	batch:= conn.ReadBatch(int(minByte), int(maxByte))
	bytes := make([]byte, 1e3)
	for{
		_, err := batch.Read(bytes)
		if err != nil {
			break
		}
		fmt.Println(string(bytes)) 
	}
}