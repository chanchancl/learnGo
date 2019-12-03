package main

import (
	"context"
	"fmt"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "time",
		Balancer: &kafka.LeastBytes{},
	})
	for {
		bt, _ := time.Now().MarshalText()
		fmt.Println("Writing Messages")
		w.WriteMessages(context.Background(),
			kafka.Message{
				Value: bt,
			},
		)
		time.Sleep(time.Second)
	}
	w.Close()
}
