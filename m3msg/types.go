package m3msg

import (
	"github.com/m3db/m3cluster/services"
)

type Topic interface {
	Name() string
	NumberOfShards() int64
	DefaultHashType() HashType
	ConsumerServices() []ConsumerService
}

type ConsumerService interface {
	ServiceID() services.ServiceID
	ConsumptionType() ConsumptionType
}

// HashType defines the hashing function for the data.
type HashType int

const (
	Murmur32 HashType = iota
)

// ConsumptionType defines how the consumer consumes data.
type ConsumptionType int

const (
	// Contending means the data for each shard will be
	// contended by all the responsible instances.
	Contending ConsumptionType = iota

	// Replicate means the data for each shard will be
	// replicated to all the responsible instances.
	Replicate
)

type ConsumerListener interface {
	Accept() (Consumer, error)
	Addr() string
	Close() error
}

func NewConsumerListener(addr string) ConsumerListener {
	return nil
}

type Consumer interface {
	Consume() (<-chan Message, error)
	Ack(msg Message) error
	Close() error
}

type Message struct {
	Offset int64
	Data   []byte
}

type TopicConfig interface {
	Name() string
}

type TopicService interface {
	Topic(config TopicConfig) Topic
}
