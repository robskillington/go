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

type msgIterator struct {
	iteratorBase

	ignoreHigherVersion bool
	closed              bool
	iteratorPool        MsgIteratorPool
	msg                 Msg
	encodedAtNanos      int64
	value               []byte
}

// NewMsgIterator creates a new msg iterator.
func NewMsgIterator(reader io.Reader, opts MsgIteratorOptions) MsgIterator {
	if opts == nil {
		opts = NewMsgIteratorOptions()
	}
	readerBufferSize := opts.ReaderBufferSize()
	return &msgIterator{
		ignoreHigherVersion: opts.IgnoreHigherVersion(),
		iteratorBase:        newBaseIterator(reader, readerBufferSize),
		msg:                 Msg{},
		iteratorPool:        opts.IteratorPool(),
	}
}

func (it *msgIterator) Err() error { return it.err() }

func (it *msgIterator) Reset(reader io.Reader) {
	it.closed = false
	it.reset(reader)
}

func (it *msgIterator) Msg() Msg {
	return it.msg
}

func (it *msgIterator) Next() bool {
	if it.err() != nil || it.closed {
		return false
	}
	return it.decodeRootObject()
}

func (it *msgIterator) Close() {
	if it.closed {
		return
	}
	it.closed = true
	it.reset(emptyReader)
	// TODO: Reset Msg
	if it.iteratorPool != nil {
		it.iteratorPool.Put(it)
	}
}

func (it *msgIterator) decodeRootObject() bool {
	version := it.decodeVersion()
	if it.err() != nil {
		return false
	}
	// If the actual version is higher than supported version, we skip
	// the data for this metric and continue to the next.
	if version > msgVersion {
		if it.ignoreHigherVersion {
			it.skip(it.decodeNumObjectFields())
			return it.Next()
		}
		it.setErr(fmt.Errorf("received version %d is higher than supported version %d", version, msgVersion))
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
	case msgType:
		it.decodeMsg()
	default:
		it.setErr(fmt.Errorf("unrecognized object type %v", objType))
	}
	it.skip(numActualFields - numExpectedFields)

	return it.err() == nil
}

func (it *msgIterator) decodeMsg() {
	numExpectedFields, numActualFields, ok := it.checkNumFieldsForType(msgType)
	if !ok {
		return
	}
	it.msg.Offset = it.decodeVarint()
	it.msg.Offset2 = it.decodeVarint()
	it.msg.Offset3 = it.decodeVarint()
	it.msg.Offset4 = it.decodeVarint()
	it.msg.Offset5 = it.decodeVarint()
	it.msg.Offset6 = it.decodeVarint()
	it.msg.Offset7 = it.decodeVarint()
	it.msg.Offset8 = it.decodeVarint()
	it.msg.Offset9 = it.decodeVarint()
	// it.msg.Value = it.decodeBytes()
	it.msg.Value = it.decodeBytesZeroAlloc()
	it.skip(numActualFields - numExpectedFields)
}

func (it *msgIterator) decodeBytesZeroAlloc() []byte {
	idLen := it.decodeBytesLen()
	if it.err() != nil {
		return nil
	}
	// NB(xichen): DecodeBytesLen() returns -1 if the byte slice is nil.
	if idLen == -1 {
		return nil
	}
	if cap(it.value) < idLen {
		it.value = make([]byte, idLen)
	} else {
		it.value = it.value[:idLen]
	}
	if _, err := io.ReadFull(it.reader(), it.value); err != nil {
		it.setErr(err)
		return nil
	}
	return it.value
}
