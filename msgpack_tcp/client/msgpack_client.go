package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"runtime/pprof"
	"sync"
	"time"

	"github.com/cw9/go/msgpack_tcp/msgpack"
)

const (
	defaultAddress     = "0.0.0.0:50051"
	defaultNumOfMsgs   = 10000000
	defaultPayloadSize = 200
)

var (
	address     = flag.String("address", defaultAddress, "server address")
	payloadSize = flag.Int64("payloadSize", defaultPayloadSize, "payload size")
	numOfMsgs   = flag.Int64("numOfMsgs", defaultNumOfMsgs, "number of messages")
	cpuprofile  = flag.String("cpu", "", "write cpu profile to file")
	receiveAck  = flag.Bool("receiveAck", false, "should receive acks")
)

func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", *address)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatalf("could not open pprof file, %v", err)
		}
		log.Printf("CPU profile will output to %s", f.Name())
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	var wg sync.WaitGroup
	if *receiveAck {
		log.Println("reiceiving acks")
		wg.Add(1)
		go receiveAcks(conn, &wg)
	}

	start := time.Now()
	sendMsgs(conn)
	log.Println("total time for sending messages", time.Now().Sub(start))
	wg.Wait()
}

func sendMsgs(conn net.Conn) {
	var (
		w          = bufio.NewWriterSize(conn, 1024)
		id         = int64(0)
		value      = make([]byte, *payloadSize)
		msgEncoder = msgpack.NewMsgEncoder(msgpack.NewBufferedEncoder())
	)

	log.Printf("start sending msgs")
	start := time.Now()
	for id <= *numOfMsgs {
		id++
		m := msgpack.Msg{
			Offset: id,
			Value:  value,
		}
		err := msgEncoder.EncodeMsg(m)
		if err != nil {
			log.Fatalf("%v", err)
		}
		encoder := msgEncoder.Encoder()
		if _, err := w.Write(encoder.Bytes()); err != nil {
			log.Fatalf("write error, %v", err)
		}
		encoder.Reset()
		if id%1000000 == 0 {
			now := time.Now()
			log.Println("send", id, now.Sub(start))
			start = now
		}
	}
	if err := w.Flush(); err != nil {
		log.Fatalf("could not flush: %v", err)
	}
}

func receiveAcks(conn net.Conn, wg *sync.WaitGroup) {
	var (
		r  = bufio.NewReaderSize(conn, 1024)
		it = msgpack.NewAckIterator(r, nil)
	)
	log.Println("reiceiving acks")
	last := time.Now()
	for {
		if !it.Next() {
			break
		}
		ack := it.Ack()
		if ack.Offset%1000000 == 0 {
			now := time.Now()
			log.Println("ack", ack.Offset, now.Sub(last))
			last = now
		}
		if ack.Offset == *numOfMsgs {
			break
		}
	}
	if err := it.Err(); err != nil && err != io.EOF {
		log.Fatalf("err in reiceiving acks: %v", it.Err())
	}
	log.Println("done receiving acks")
	wg.Done()
}
