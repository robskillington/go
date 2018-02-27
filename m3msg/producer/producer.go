package producer

import (
	"container/list"
	"sync"
	"sync/atomic"

	"github.com/m3db/m3x/checked"
)

// type Producer interface {
// 	Produce(key []byte, data Data) error

// 	ProduceWithShard(shard int, data Data) error

// 	RegistConsumerFilter(req RegistConsumerFilterRequest) error

// 	AddConsumer(req AddConsumerRequest) error

// 	RemoveConsumer(req RemoveConsumerRequest) error

// 	Close() error
// }

type producer struct {
	b buffer
}

func (p *producer) Produce(key []byte, data Data) error {
	if err := p.b.push(data); err != nil {
		return err
	}
	// TODO: route the data to message writer.
	return nil
}

// func NewProducer(topic Topic) Producer {
// 	var sd services.Services
// 	for _, consumer := range topic.ConsumerMetadata() {
// 		watch, _ := sd.Watch(consumer, nil)
// 		// Listen to the watch for placement changes
// 		_ = watch
// 	}
// 	return nil
// }

func (p *producer) purgeData() {

}

type refCountedElement interface {
	checked.Ref
	next() refCountedElement
	prev() refCountedElement
}

type refCountedData struct {
	ref  checked.Ref
	data Data
}

type buffer interface {
	push(d Data) error
	remove(e refCountedElement)
	front() refCountedElement
	back() refCountedElement
	cleanup()
}

type bufferOpts struct {
	limit uint64
}

type lists struct {
	sync.Mutex

	opts     bufferOpts
	buffered uint64
	new      *list.List
	old      [][]refCountedData
}

func (l *lists) push(d Data) error {
	if atomic.LoadUint64(&l.buffered) < l.opts.limit {
		l.Lock()
		l.new.PushFront(d)
		l.Unlock()
	}
	// TODO: drop and push
	return nil
}

// type listBasedBuffer struct {
// 	sync.Mutex

// 	buffered int
// 	l        *list.List
// }

// func newListBasedBuffer() buffer {
// 	d := listBasedBuffer{
// 		l: list.New(),
// 	}
// 	return &d
// }

// func (b *listBasedBuffer) push(d Data) error {
// 	b.Lock()
// 	defer b.Unlock()
// 	b.l.Pu
// }
