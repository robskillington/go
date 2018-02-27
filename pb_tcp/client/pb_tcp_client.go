package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"sync"
	"time"

	"github.com/cw9/go/pb_tcp"

	"github.com/cw9/go/pb_tcp/msgpb"
)

const (
	defaultAddress     = "0.0.0.0:50051"
	defaultPayloadSize = 200
	defaultNumOfMsgs   = 10000000
)

var (
	address     = flag.String("address", defaultAddress, "server address")
	payloadSize = flag.Int64("payloadSize", defaultPayloadSize, "payload size")
	numOfMsgs   = flag.Int64("numOfMsgs", defaultNumOfMsgs, "number of messages")
	receiveAck  = flag.Bool("receiveAck", false, "should receive acks")
)

func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", *address)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	var wg sync.WaitGroup
	if *receiveAck {
		wg.Add(1)
		go receiveAcks(conn, &wg)
	}

	sendMsgs(conn)
	wg.Wait()

}

func sendMsgs(conn net.Conn) {
	var (
		w          = bufio.NewWriterSize(conn, 1024)
		id         = int64(0)
		dataBuffer = make([]byte, 20480)
		sizeBuffer = make([]byte, 4)
		msg        = msgpb.Message{
			Value: make([]byte, *payloadSize),
		}
	)

	log.Printf("start sending msgs")

	start := time.Now()
	last := start
	for id < *numOfMsgs {
		id++
		msg.Offset = id
		if err := pb_tcp.Encode(&msg, sizeBuffer, dataBuffer, w); err != nil {
			log.Fatal(err.Error())
		}
		if id%1000000 == 0 {
			now := time.Now()
			log.Println("sent", id, now.Sub(last))
			last = now
		}
		if id == *numOfMsgs {
			break
		}
	}
	if err := w.Flush(); err != nil {
		log.Fatalf("could not flush: %v", err)
	}
	log.Println("total time for sending messages", time.Now().Sub(start))
}

func receiveAcks(conn net.Conn, wg *sync.WaitGroup) {
	var (
		r          = bufio.NewReaderSize(conn, 1024)
		sizeBuffer = make([]byte, 4)
		dataBuffer = make([]byte, 1024)
		ack        msgpb.Ack
	)

	log.Println("reiceiving acks")
	start := time.Now()
	last := start
	for {
		if err := pb_tcp.Decode(&ack, sizeBuffer, dataBuffer, r); err != nil {
			log.Fatal(err.Error())
		}
		if ack.Offset%1000000 == 0 {
			now := time.Now()
			log.Println("ack", ack.Offset, now.Sub(last))
			last = now
		}
		if ack.Offset == *numOfMsgs {
			break
		}
	}
	log.Println("total time for receiving acks", time.Now().Sub(start))
	wg.Done()
}
