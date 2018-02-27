package producer

// DataFinalizeReason defines the reason why the data is being finalized by Producer.
type DataFinalizeReason int

const (
	// Consumed means the data has been fully consumed.
	Consumed DataFinalizeReason = iota

	// Expired means the data has been expired.
	Expired
)

type Data interface {
	Bytes() []byte
	Finalize(DataFinalizeReason) error
}

type Producer interface {
	Produce(key []byte, data Data) error

	ProduceWithShard(shard int, data Data) error

	RegistConsumerFilter(req RegistConsumerFilterRequest) error

	AddConsumer(req AddConsumerRequest) error

	RemoveConsumer(req RemoveConsumerRequest) error

	Close() error
}

type RegistConsumerFilterRequest interface{}
type AddConsumerRequest interface{}
type RemoveConsumerRequest interface{}
