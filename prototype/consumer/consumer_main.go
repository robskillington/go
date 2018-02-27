package prototype

import (
	"sync"

	"github.com/cw9/go/prototype"
)

type app struct {
	lis prototype.ConsumerListener
}

func main() {
	a := app{lis: prototype.NewConsumerListener("0.0.0.0:1234")}

	for {
		consumer, _ := a.lis.Accept()
		go a.work(consumer)
	}
	// Safe when ungracefully shutdown.
}

func (a app) work(consumer prototype.Consumer) {
	ch, _ := consumer.Subscribe()

	var wg sync.WaitGroup
	for msg := range ch {
		msg := msg
		wg.Add(1)
		go func() {
			// Do work on msg.
			_ = msg.Value()
			// Ack as soon as the msg is done processing, no need to maintain order.
			consumer.Ack(msg)
			wg.Done()
		}()
	}
	wg.Wait()
	consumer.Close()
}
