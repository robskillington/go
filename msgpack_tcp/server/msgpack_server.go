package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"time"

	"github.com/cw9/go/msgpack_tcp/msgpack"
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

func process(conn net.Conn, shouldAck bool) error {
	var (
		r          = bufio.NewReaderSize(conn, 1024)
		w          = bufio.NewWriterSize(conn, 1024)
		it         = msgpack.NewMsgIterator(r, nil)
		ackEncoder = msgpack.NewAckEncoder(msgpack.NewBufferedEncoder())
	)
	defer w.Flush()

	for {
		if !it.Next() {
			break
		}
		msg := it.Msg()
		if !shouldAck {
			continue
		}
		if err := ack(ackEncoder, w, msg.Offset); err != nil {
			return err
		}
	}
	if err := it.Err(); err != nil && err != io.EOF {
		log.Printf("err: %v", it.Err())
		return it.Err()
	}
	log.Println("done receiving msgs")
	return nil
}

func ack(ackEncoder msgpack.AckEncoder, writer io.Writer, offset int64) error {
	ackEncoder.EncodeAck(msgpack.Ack{
		Offset: offset,
	})
	encoder := ackEncoder.Encoder()
	if _, err := writer.Write(encoder.Bytes()); err != nil {
		return err
	}
	encoder.Reset()
	if offset%1000000 == 0 {
		log.Println("ack", offset, time.Now())
	}
	return nil
}
