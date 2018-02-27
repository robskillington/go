package producer

import (
	"github.com/cw9/go/prototype"
)

type p struct {
	producer prototype.Producer
}

func main() {
	var topic prototype.Topic
	p := p{producer: prototype.NewProducer(topic)}

	for {
		var data []byte
		p.producer.PublishWithShard(1, data)
	}
}
