// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package msgpack

import (
	"testing"
	"time"

	"github.com/m3db/m3metrics/metric/aggregated"
	"github.com/m3db/m3metrics/policy"
	"github.com/m3db/m3metrics/protocol/msgpack"
	xtime "github.com/m3db/m3x/time"
)

func BenchmarkRoundTripMsgpackWithDecodeInOneGo(b *testing.B) {
	var (
		encoder  = testMsgEncoder()
		iterator = testMsgIterator(nil)
	)

	id := "01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789"
	p := policy.NewStoragePolicy(10*time.Second, xtime.Second, time.Hour)
	enc := msgpack.NewAggregatedEncoder(msgpack.NewBufferedEncoder())
	enc.EncodeMetricWithStoragePolicy(aggregated.MetricWithStoragePolicy{
		Metric: aggregated.Metric{
			ID:        []byte(id),
			Value:     1,
			TimeNanos: 1,
		},
		StoragePolicy: p,
	})

	m := Msg{
		Offset: 1,
		Value:  enc.Encoder().Bytes(),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		encoder.EncodeMsg(m)
		iterator.Reset(encoder.Encoder().Buffer())
		n := iterator.Next()
		if !n {
			b.FailNow()
		}
		encoder.Encoder().Reset()

		// rm, sp, _ := iterator.(*msgIterator).Value()
		// if sp != p {
		// 	b.FailNow()
		// }
		// m, _ := rm.Metric()
		// if m.Value != float64(1) {
		// 	b.FailNow()
		// }
		// if m.TimeNanos != int64(1) {
		// 	b.FailNow()
		// }
		// if string(m.ID) != id {
		// 	b.FailNow()
		// }
	}
}
