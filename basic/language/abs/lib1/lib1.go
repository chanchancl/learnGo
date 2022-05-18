package lib1

import (
	"fmt"
	"strings"
)

type callback func(topic, msg string)
type Consumer struct {
	topics []string
	cb     callback
}

func CreateConsumer(cb callback, topics ...string) *Consumer {
	fmt.Println("Start listening topics " + strings.Join(topics, " "))

	return &Consumer{
		topics: topics,
		cb:     cb,
	}
}
