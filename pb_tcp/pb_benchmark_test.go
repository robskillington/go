package pb_tcp

import (
	"bytes"
	"testing"
	"time"

	"github.com/cw9/go/pb_tcp/msgpb"
	"github.com/m3db/m3metrics/metric/aggregated"
	"github.com/m3db/m3metrics/metric/unaggregated"
	"github.com/m3db/m3metrics/policy"
	"github.com/m3db/m3metrics/protocol/msgpack"
	xtime "github.com/m3db/m3x/time"
)

var (
	id = "01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789"
)

func BenchmarkRoundTripProtobuf(b *testing.B) {
	mimicTCP := bytes.NewBuffer(nil)
	encodeMsg := msgpb.Message{
		Value: make([]byte, 200),
	}
	encodeSizeBuffer := make([]byte, 4)
	encodeDataBuffer := make([]byte, 20480)
	decodeMsg := msgpb.Message{}
	decodeSizeBuffer := make([]byte, 4)
	decodeDataBuffer := make([]byte, 20480)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		encodeMsg.Offset = int64(n)
		if err := Encode(&encodeMsg, encodeSizeBuffer, encodeDataBuffer, mimicTCP); err != nil {
			b.FailNow()
		}
		if err := Decode(&decodeMsg, decodeSizeBuffer, decodeDataBuffer, mimicTCP); err != nil {
			b.FailNow()
		}
		// if decodeMsg.Offset != int64(n) {
		// 	b.FailNow()
		// }
	}
}

func BenchmarkRoundTripProtobufWithDecodeMsgpack(b *testing.B) {
	mimicTCP := bytes.NewBuffer(nil)
	encodeSizeBuffer := make([]byte, 4)
	encodeDataBuffer := make([]byte, 20480)
	decodeMsg := msgpb.Message{}
	decodeSizeBuffer := make([]byte, 4)
	decodeDataBuffer := make([]byte, 20480)

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

	encodeMsg := msgpb.Message{
		Value: enc.Encoder().Bytes(),
	}

	it := msgpack.NewAggregatedIterator(nil, nil)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		encodeMsg.Offset = int64(n)
		if err := Encode(&encodeMsg, encodeSizeBuffer, encodeDataBuffer, mimicTCP); err != nil {
			b.FailNow()
		}
		if err := Decode(&decodeMsg, decodeSizeBuffer, decodeDataBuffer, mimicTCP); err != nil {
			b.FailNow()
		}
		it.Reset(bytes.NewBuffer(decodeMsg.Value))
		n := it.Next()
		if !n {
			b.FailNow()
		}
		// if decodeMsg.Offset != int64(n) {
		// 	b.FailNow()
		// }
		// rm, _, _ := it.Value()
		// m, _ := rm.Metric()
		// if m.Value != 1 {
		// 	b.FailNow()
		// }
		// if string(m.ID) != id {
		// 	b.FailNow()
		// }
	}
}

func BenchmarkRoundTripProtobufWithDecodePB(b *testing.B) {
	mimicTCP := bytes.NewBuffer(nil)
	encodeSizeBuffer := make([]byte, 4)
	encodeDataBuffer := make([]byte, 20480)
	decodeMsg := msgpb.Message{}
	decodeSizeBuffer := make([]byte, 4)
	decodeDataBuffer := make([]byte, 20480)

	pb := msgpb.AggregatedMetric{
		RawMetric: &msgpb.RawMetric{
			Id:    []byte(id),
			Value: 1,
			Nanos: 1,
		},
		StoragePolicy: &msgpb.StoragePolicy{
			Resolution: &msgpb.Resolution{
				Window:    10 * int64(time.Second),
				Precision: int64(time.Second),
			},
			Retention: &msgpb.Retention{
				Retention: int64(time.Hour),
			},
		},
	}
	bytes, _ := pb.Marshal()
	encodeMsg := msgpb.Message{
		Value: bytes,
	}

	var decodePb msgpb.AggregatedMetric

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		encodeMsg.Offset = int64(n)
		if err := Encode(&encodeMsg, encodeSizeBuffer, encodeDataBuffer, mimicTCP); err != nil {
			b.FailNow()
		}
		if err := Decode(&decodeMsg, decodeSizeBuffer, decodeDataBuffer, mimicTCP); err != nil {
			b.FailNow()
		}
		err := decodePb.Unmarshal(decodeMsg.Value)
		if err != nil {
			b.FailNow()
		}
		// if decodeMsg.Offset != int64(n) {
		// 	b.FailNow()
		// }
		// if string(decodePb.RawMetric.Id) != id {
		// 	b.FailNow()
		// }
	}
}

func BenchmarkRoundTripProtobufWithMoreFields(b *testing.B) {
	mimicTCP := bytes.NewBuffer(nil)
	encodeMsg := msgpb.MessageBig{
		Value: make([]byte, 200),
	}
	encodeSizeBuffer := make([]byte, 4)
	encodeDataBuffer := make([]byte, 20480)
	decodeMsg := msgpb.MessageBig{}
	decodeSizeBuffer := make([]byte, 4)
	decodeDataBuffer := make([]byte, 20480)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		encodeMsg.Offset1 = int64(n)
		encodeMsg.Offset2 = int64(n)
		encodeMsg.Offset3 = int64(n)
		encodeMsg.Offset4 = int64(n)
		encodeMsg.Offset5 = int64(n)
		encodeMsg.Offset6 = int64(n)
		encodeMsg.Offset7 = int64(n)
		encodeMsg.Offset8 = int64(n)
		encodeMsg.Offset9 = int64(n)
		if err := Encode(&encodeMsg, encodeSizeBuffer, encodeDataBuffer, mimicTCP); err != nil {
			b.FailNow()
		}
		if err := Decode(&decodeMsg, decodeSizeBuffer, decodeDataBuffer, mimicTCP); err != nil {
			b.FailNow()
		}
		// if decodeMsg.Offset1 != int64(n) {
		// 	b.FailNow()
		// }
	}
}

func BenchmarkAggregatedMetricRoundTripInProtobuf(b *testing.B) {
	encodePb := msgpb.AggregatedMetric{
		RawMetric: &msgpb.RawMetric{
			Id:    []byte(id),
			Value: 1,
			Nanos: 1,
		},
		StoragePolicy: &msgpb.StoragePolicy{
			Resolution: &msgpb.Resolution{
				Window:    10 * int64(time.Second),
				Precision: int64(time.Second),
			},
			Retention: &msgpb.Retention{
				Retention: int64(time.Hour),
			},
		},
	}

	decodePb := msgpb.AggregatedMetric{}

	bytes := make([]byte, 2048)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		size, err := encodePb.MarshalTo(bytes)
		if err != nil {
			b.FailNow()
		}
		err = decodePb.Unmarshal(bytes[:size])
		if err != nil {
			b.FailNow()
		}
		// if string(decodePb.RawMetric.Id) != id {
		// 	b.FailNow()
		// }
	}
}

func BenchmarkAggregatedMetricRoundTripInMsgpack(b *testing.B) {
	p := policy.NewStoragePolicy(10*time.Second, xtime.Second, time.Hour)
	m := aggregated.MetricWithStoragePolicy{
		Metric: aggregated.Metric{
			ID:        []byte(id),
			Value:     1,
			TimeNanos: 1,
		},
		StoragePolicy: p,
	}
	enc := msgpack.NewAggregatedEncoder(msgpack.NewBufferedEncoder())
	it := msgpack.NewAggregatedIterator(nil, nil)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		enc.EncodeMetricWithStoragePolicy(m)
		it.Reset(enc.Encoder().Buffer())
		n := it.Next()
		if !n {
			b.FailNow()
		}
		enc.Encoder().Buffer().Reset()
		// rm, _, _ := it.Value()
		// m, _ := rm.Metric()
		// // if string(m.ID) != id {
		// 	b.FailNow()
		// }
	}
}

func BenchmarkUnaggregatedMetricRoundTripInProtobuf(b *testing.B) {
	values := make([]float64, 140)
	for i := 0; i < 140; i++ {
		values[i] = float64(i + 10000000)
	}
	encodePb := msgpb.UnAggregatedMetric{
		Id:     []byte(id),
		Values: values,
	}
	decodePb := msgpb.UnAggregatedMetric{}

	bytes := make([]byte, 2048)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		size, err := encodePb.MarshalTo(bytes)
		if err != nil {
			b.FailNow()
		}
		err = decodePb.Unmarshal(bytes[:size])
		if err != nil {
			b.FailNow()
		}
		// if string(decodePb.Id) != id {
		// 	b.FailNow()
		// }
	}
}

func BenchmarkUnaggregatedMetricRoundTripInMsgpack(b *testing.B) {
	values := make([]float64, 140)
	for i := 0; i < 140; i++ {
		values[i] = float64(i + 10000000)
	}
	enc := msgpack.NewUnaggregatedEncoder(msgpack.NewBufferedEncoder())
	it := msgpack.NewUnaggregatedIterator(nil, nil)
	bt := unaggregated.BatchTimer{
		ID:     []byte(id),
		Values: values,
	}
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		enc.EncodeBatchTimer(bt)
		it.Reset(enc.Encoder().Buffer())
		n := it.Next()
		if !n {
			b.FailNow()
		}
		enc.Encoder().Buffer().Reset()

		// mu := it.Metric()
		// if string(mu.ID) != id {
		// 	b.FailNow()
		// }
	}
}
