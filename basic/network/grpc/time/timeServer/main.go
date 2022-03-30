package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"time"

	timeProto "learnGo/basic/network/grpc/time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	port string = ":9001"
)

type server struct{}

func (sv *server) GetTime(context.Context, *empty.Empty) (*timeProto.Time, error) {
	ret := fmt.Sprintf("现在已经 %v 啦!", time.Now())
	time.Sleep(time.Duration(5) * time.Second)
	log.Println(ret)
	return &timeProto.Time{Now: ret}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	grpcServer := grpc.NewServer()
	timeProto.RegisterTimeNowServer(grpcServer, &server{})
	//reflection.Register(grpcServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
