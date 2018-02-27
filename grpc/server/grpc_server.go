package main

import (
	"io"
	"log"
	"net"

	"github.com/cw9/go/grpc/msgpb"
	"google.golang.org/grpc"
)

const (
	port = "0.0.0.0:50051"
)

// server is used to implement customer.CustomerServer.
type server struct {
}

// CreateCustomer creates a new Customer
func (s *server) Send(stream msgpb.Queue_SendServer) error {
	last := int64(-1)
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("stream received err, %v", err)
			return err
		}
		if msg.Id != last+1 {
			log.Printf("last: %d, msg: %d", last, msg.Id)
		}
		last = msg.Id
		if last%1000000 == 0 {
			log.Println(last)
		}
	}
	log.Printf("done")
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	msgpb.RegisterQueueServer(s, &server{})
	s.Serve(lis)
}
