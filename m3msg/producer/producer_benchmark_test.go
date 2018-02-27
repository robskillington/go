package producer

import (
	"container/list"
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkAtomic(b *testing.B) {
	var a uint64
	for n := 0; n < b.N; n++ {
		atomic.AddUint64(&a, 1)
	}
}

func BenchmarkLock(b *testing.B) {
	var a uint64
	var l sync.Mutex
	for n := 0; n < b.N; n++ {
		l.Lock()
		a++
		l.Unlock()
	}
}

func BenchmarkSliceCreate(b *testing.B) {
	b.ReportAllocs()
	var s []interface{}
	for n := 0; n < b.N; n++ {
		s = make([]interface{}, 1024)
		for i := 0; i < 1024; i++ {
			a := int32(0)
			s[i] = a
		}
	}
}

func BenchmarkListCreate2(b *testing.B) {
	b.ReportAllocs()
	var l *list.List
	for n := 0; n < b.N; n++ {
		l = list.New()
		for i := 0; i < 1024; i++ {
			l.PushFront(i)
		}
	}
}

func BenchmarkSliceRemove(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := make([]interface{}, 1024)
		for i := 0; i < 1024; i++ {
			s[i] = i
		}
		for i := 0; i < 1024; i++ {
			size := len(s)
			s[0] = s[size-1]
			s = s[:size-1]
		}
	}
}

func BenchmarkListRemove(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := list.New()
		for i := 0; i < 1024; i++ {
			l.PushFront(i)
		}
		for i := 0; i < 1024; i++ {
			l.Remove(l.Back())
		}
	}
}
