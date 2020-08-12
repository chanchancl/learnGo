package main

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.Dial("tcp", ":9092")
	if err != nil {
		fmt.Println(err)
		return
	}
	controller, _ := conn.Controller()
	conncontroller, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = conncontroller.CreateTopics(kafka.TopicConfig{
		Topic:         "time",
		NumPartitions: 1,
	})
	if err != nil {
		fmt.Println(err)
	}

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "time",
	})
	for {
		bt, _ := time.Now().MarshalText()
		fmt.Println("Writing Messages")
		fmt.Printf("Before write : %s\n", time.Now())
		err := w.WriteMessages(context.Background(),
			kafka.Message{
				Value: bt,
			},
		)
		fmt.Printf("After write : %s\n", time.Now())
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(w.Stats())
		time.Sleep(time.Second)
	}
	w.Close()
}
