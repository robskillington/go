package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/cw9/go/grpc/msgpb"
	"google.golang.org/grpc"
)

const (
	defaultAddress = "0.0.0.0:50051"
	payloadSize    = 200
)

var (
	address = flag.String("address", defaultAddress, "server address")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := msgpb.NewQueueClient(conn)

	stream, err := client.Send(context.Background())
	if err != nil {
		log.Fatalf("could not read: %v", err)
	}

	id := int64(0)
	value := make([]byte, payloadSize)
	log.Printf("start")
	start := time.Now()
	msg := &msgpb.Message{
		Id:    id,
		Value: value,
	}
	last := start
	for id < 10000000 {
		msg.Id = id
		if err := stream.Send(msg); err != nil {
			log.Fatalf("could not read: %v", err)
		}
		id++
		if id%1000000 == 0 {
			now := time.Now()
			log.Println(id, now.Sub(last))
			last = now
		}
	}
	log.Println("done", time.Now().Sub(start))
}
