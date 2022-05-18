package lib2

import "fmt"

type PublishInterface interface {
	GetConsumerMessageChan() <-chan *ConsumerMessage
}

type ConsumerMessage struct {
	Topic string
	Msg   string
}

type Lib2 struct {
	inf     PublishInterface
	msgChan <-chan *ConsumerMessage

	Exit chan struct{}
}

func Init(publisher PublishInterface) *Lib2 {
	lib2 := &Lib2{
		inf:     publisher,
		msgChan: publisher.GetConsumerMessageChan(),
		Exit:    make(chan struct{}),
	}
	go lib2.loop()

	return lib2
}

func (c *Lib2) WaitUntilExit() {
	<-c.Exit
}

func (c *Lib2) loop() {
	for msg := range c.msgChan {
		c.OnMessage(msg.Topic, msg.Msg)
	}
	fmt.Println("Lib2 : loop exit")
	close(c.Exit)
}

func (c *Lib2) OnMessage(topic, msg string) {
	fmt.Printf("lib2 : Received topic %v, msg %v\n", topic, msg)
}
