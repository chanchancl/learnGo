package main

import (
	"fmt"
	"learnGo/basic/language/abs/lib1"
	"learnGo/basic/language/abs/lib2"
	"time"
)

// Lib1Agent use function from lib1, and implement the interface from lib2
// So lib2 didn't need to import lib1, but only need to declare the interface
type Lib1Agent struct {
	msgChan chan *lib2.ConsumerMessage
}

func InitLib1Agent() *Lib1Agent {
	agent := &Lib1Agent{msgChan: make(chan *lib2.ConsumerMessage)}

	lib1.CreateConsumer(agent.Callback)

	return agent
}

func (c *Lib1Agent) GetConsumerMessageChan() <-chan *lib2.ConsumerMessage {
	return c.msgChan
}

func (c *Lib1Agent) Callback(topic, msg string) {
	if topic == "" {
		close(c.msgChan)
		return
	}
	fmt.Printf("Lib1Agent : Received message from topic %v, msg %v\n", topic, msg)

	c.msgChan <- &lib2.ConsumerMessage{
		Topic: topic,
		Msg:   msg,
	}
}

func (c *Lib1Agent) SimulateSomeMessage() {
	// simulate receive some msg
	time.Sleep(time.Second)

	c.Callback("apple", "We get 3 apples")
	c.Callback("apple", "We lost 2 apples")

	time.Sleep(time.Second)

	c.Callback("banana", "We get 5 bananas")
	time.Sleep(time.Second)

	close(c.msgChan)
}

func main() {
	// take function PublishString from Lib1 to Init Lib2
	lib1Agent := InitLib1Agent()
	go lib1Agent.SimulateSomeMessage()

	l2 := lib2.Init(lib1Agent)

	l2.WaitUntilExit()
}
