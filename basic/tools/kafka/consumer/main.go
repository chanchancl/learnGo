package main

import (
	"context"
	"fmt"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "time",
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})

	r.SetOffset(kafka.LastOffset)
	for {
		fmt.Printf("Reading messages :%s\n", time.Now())
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Printf("After Reading messages :%s\n", time.Now())
		fmt.Printf("Message at partition %d, offset %d: %s = %s\n", m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
