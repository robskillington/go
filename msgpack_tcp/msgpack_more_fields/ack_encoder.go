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

type encodeAckFn func(ack Ack)

type ackEncoder struct {
	encoderBase

	encodeRootObjectFn encodeRootObjectFn
	encodeAckFn        encodeAckFn
}

// NewAckEncoder creates an ack encoder.
func NewAckEncoder(encoder BufferedEncoder) AckEncoder {
	enc := &ackEncoder{
		encoderBase: newBaseEncoder(encoder),
	}

	enc.encodeRootObjectFn = enc.encodeRootObject
	enc.encodeAckFn = enc.encodeAck

	return enc
}

func (enc *ackEncoder) Encoder() BufferedEncoder      { return enc.encoder() }
func (enc *ackEncoder) Reset(encoder BufferedEncoder) { enc.reset(encoder) }

func (enc *ackEncoder) EncodeAck(
	ack Ack,
) error {
	if err := enc.err(); err != nil {
		return err
	}
	enc.encodeRootObjectFn(ackType)
	enc.encodeAckFn(ack)
	return enc.err()
}

func (enc *ackEncoder) encodeRootObject(objType objectType) {
	enc.encodeVersion(ackVersion)
	enc.encodeNumObjectFields(numFieldsForType(rootObjectType))
	enc.encodeObjectType(objType)
}

func (enc *ackEncoder) encodeAck(
	ack Ack,
) {
	enc.encodeNumObjectFields(numFieldsForType(ackType))
	enc.encodeVarint(ack.Offset)
}
