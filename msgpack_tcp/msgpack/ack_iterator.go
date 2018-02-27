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
	"fmt"
	"io"
)

type ackIterator struct {
	iteratorBase

	ignoreHigherVersion bool
	closed              bool
	ack                 Ack
	encodedAtNanos      int64
}

// NewAckIterator creates a new ack iterator.
func NewAckIterator(reader io.Reader, opts MsgIteratorOptions) AckIterator {
	if opts == nil {
		opts = NewMsgIteratorOptions()
	}
	readerBufferSize := opts.ReaderBufferSize()
	return &ackIterator{
		ignoreHigherVersion: opts.IgnoreHigherVersion(),
		iteratorBase:        newBaseIterator(reader, readerBufferSize),
		ack:                 Ack{},
	}
}

func (it *ackIterator) Err() error { return it.err() }

func (it *ackIterator) Reset(reader io.Reader) {
	it.closed = false
	it.reset(reader)
}

func (it *ackIterator) Ack() Ack {
	return it.ack
}

func (it *ackIterator) Next() bool {
	if it.err() != nil || it.closed {
		return false
	}
	return it.decodeRootObject()
}

func (it *ackIterator) Close() {
	if it.closed {
		return
	}
	it.closed = true
	it.reset(emptyReader)
}

func (it *ackIterator) decodeRootObject() bool {
	version := it.decodeVersion()
	if it.err() != nil {
		return false
	}
	// If the actual version is higher than supported version, we skip
	// the data for this metric and continue to the next.
	if version > ackVersion {
		if it.ignoreHigherVersion {
			it.skip(it.decodeNumObjectFields())
			return it.Next()
		}
		it.setErr(fmt.Errorf("received version %d is higher than supported version %d", version, ackVersion))
		return false
	}
	// Otherwise we proceed to decoding normally.
	numExpectedFields, numActualFields, ok := it.checkNumFieldsForType(rootObjectType)
	if !ok {
		return false
	}
	objType := it.decodeObjectType()
	if it.err() != nil {
		return false
	}
	switch objType {
	case ackType:
		it.decodeAck()
	default:
		it.setErr(fmt.Errorf("unrecognized object type %v", objType))
	}
	it.skip(numActualFields - numExpectedFields)

	return it.err() == nil
}

func (it *ackIterator) decodeAck() {
	numExpectedFields, numActualFields, ok := it.checkNumFieldsForType(ackType)
	if !ok {
		return
	}
	it.ack.Offset = it.decodeVarint()
	it.skip(numActualFields - numExpectedFields)
}
