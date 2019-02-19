package main

import (
	"context"
	"fmt"
	timeProto "learnGo/grpc/time"
	"log"
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	address string = "127.0.0.1:9001"
)

func main() {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Err %v", err)
	}
	defer con.Close()

	client := timeProto.NewTimeNowClient(con)

	for {
		now, err := client.GetTime(context.Background(), &empty.Empty{})
		if err != nil {
			log.Printf("%v", err)
		} else {
			fmt.Printf("%v\n", time.Now())
			log.Printf("%v\n", now.Now)
		}
		time.Sleep(time.Duration(rand.Float32() * float32(time.Second)))
	}

}