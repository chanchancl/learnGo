package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Printf("I'm starting now!!!!!!!!!!!")

	time.AfterFunc(time.Hour, func() {
		log.Println("Timeout")
	})

	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		for {
			<-s
			log.Printf("Received SIGTERM signal, but do nothing")
			os.Exit(0)
		}
	}()

	select {}
}
