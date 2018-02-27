package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"github.com/cw9/go/pb_tcp"
	"github.com/cw9/go/pb_tcp/msgpb"
)

const (
	port = "0.0.0.0:50051"
)

var (
	sendAck = flag.Bool("sendAck", false, "should send acks")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if *sendAck {
		log.Println("will send acks")
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("failed to accept: %v", err)
		}

		go process(conn, *sendAck)
	}
}

func process(conn net.Conn, shouldAck bool) {
	log.Println("processing a new connection")
	var (
		r          = bufio.NewReaderSize(conn, 1024)
		w          = bufio.NewWriterSize(conn, 1024)
		sizeBytes  = make([]byte, 4)
		dataBuffer = make([]byte, 20480)
		msg        msgpb.Message
		ack        msgpb.Ack
		l          sync.Mutex
		doneCh     = make(chan struct{})
		wg         sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		for {
			select {
			case <-time.Tick(1 * time.Second):
				l.Lock()
				w.Flush()
				l.Unlock()
			case <-doneCh:
				wg.Done()
				return
			}
		}
	}()

	for {
		if err := pb_tcp.Decode(&msg, sizeBytes, dataBuffer, r); err != nil {
			if err == io.EOF {
				log.Println("connection closed")
				break
			}
			log.Println("decode error", err.Error())
			break
		}
		if !shouldAck {
			continue
		}
		ack.Offset = msg.Offset
		l.Lock()
		if err := pb_tcp.Encode(&ack, sizeBytes, dataBuffer, w); err != nil {
			l.Unlock()
			log.Println("ack error", err.Error())
			break
		}
		l.Unlock()
		if ack.Offset%1000000 == 0 {
			log.Println("ack", ack.Offset, time.Now())
		}
	}
	close(doneCh)
	wg.Wait()
}
