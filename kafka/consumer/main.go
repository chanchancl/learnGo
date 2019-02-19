package main

import (
	"context"
	"fmt"

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
	fmt.Printf("%v", r.Stats())

	r.SetOffset(kafka.LastOffset)
	for {
		fmt.Println("Reading messages")
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("Message at partition %d, offset %d: %s = %s\n", m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
