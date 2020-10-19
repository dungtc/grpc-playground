package main

import (
	"context"
	"log"
	"time"

	"github.com/dungtc/grpc-playground/simple/helloworld"
	"google.golang.org/grpc"
)

var (
	address = "localhost:10000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := helloworld.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "adat"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Greeting: %s", res.GetMessage())
}
