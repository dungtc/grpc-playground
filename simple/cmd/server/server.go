package main

import (
	"context"
	"log"
	"net"

	"github.com/dungtc/grpc-playground/simple/helloworld"
	"google.golang.org/grpc"
)

var (
	address = "10000"
)

type handler struct {
}

func (h *handler) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: "hello world",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":"+address)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &handler{})

	log.Printf("GRPC Server Listening in %s", address)
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
