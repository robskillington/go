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
	"bufio"
	"bytes"
	"fmt"
	"io"

	msgpack "gopkg.in/vmihailenco/msgpack.v2"
)

var (
	emptyReader *bytes.Buffer
)

// baseIterator is the base iterator that provides common decoding APIs.
type baseIterator struct {
	readerBufferSize int
	bufReader        bufReader
	decoder          *msgpack.Decoder
	decodeErr        error
}

func newBaseIterator(reader io.Reader, readerBufferSize int) iteratorBase {
	// NB(xichen): if reader is not a bufReader, the underlying msgpack decoder
	// creates a bufio.Reader wrapping the reader. By converting the reader to a
	// bufReader, it is guaranteed that the reader passed to the decoder is the one
	// used for reading and buffering the underlying data.
	bufReader := toBufReader(reader, readerBufferSize)
	return &baseIterator{
		readerBufferSize: readerBufferSize,
		bufReader:        bufReader,
		decoder:          msgpack.NewDecoder(bufReader),
	}
}

func (it *baseIterator) reset(reader io.Reader) {
	bufReader := toBufReader(reader, it.readerBufferSize)
	it.bufReader = bufReader
	it.decoder.Reset(bufReader) // nolint: errcheck
	it.decodeErr = nil
}

func (it *baseIterator) err() error        { return it.decodeErr }
func (it *baseIterator) setErr(err error)  { it.decodeErr = err }
func (it *baseIterator) reader() bufReader { return it.bufReader }

func (it *baseIterator) decodeVersion() int {
	return int(it.decodeVarint())
}

func (it *baseIterator) decodeObjectType() objectType {
	return objectType(it.decodeVarint())
}

func (it *baseIterator) decodeNumObjectFields() int {
	return it.decodeArrayLen()
}

func (it *baseIterator) decodeOffset() int64 {
	return it.decodeVarint()
}

// NB(xichen): the underlying msgpack decoder implementation
// always decodes an int64 and looks at the actual decoded
// value to determine the width of the integer (a.k.a. varint
// decoding).
func (it *baseIterator) decodeVarint() int64 {
	if it.decodeErr != nil {
		return 0
	}
	value, err := it.decoder.DecodeInt64()
	it.decodeErr = err
	return value
}

func (it *baseIterator) decodeBool() bool {
	if it.decodeErr != nil {
		return false
	}
	value, err := it.decoder.DecodeBool()
	it.decodeErr = err
	return value
}

func (it *baseIterator) decodeFloat64() float64 {
	if it.decodeErr != nil {
		return 0.0
	}
	value, err := it.decoder.DecodeFloat64()
	it.decodeErr = err
	return value
}

func (it *baseIterator) decodeBytes() []byte {
	if it.decodeErr != nil {
		return nil
	}
	value, err := it.decoder.DecodeBytes()
	it.decodeErr = err
	return value
}

func (it *baseIterator) decodeBytesLen() int {
	if it.decodeErr != nil {
		return 0
	}
	bytesLen, err := it.decoder.DecodeBytesLen()
	it.decodeErr = err
	return bytesLen
}

func (it *baseIterator) decodeArrayLen() int {
	if it.decodeErr != nil {
		return 0
	}
	value, err := it.decoder.DecodeArrayLen()
	it.decodeErr = err
	return value
}

func (it *baseIterator) skip(numFields int) {
	if it.decodeErr != nil {
		return
	}
	if numFields < 0 {
		it.decodeErr = fmt.Errorf("number of fields to skip is %d", numFields)
		return
	}
	// Otherwise we skip any unexpected extra fields.
	for i := 0; i < numFields; i++ {
		if err := it.decoder.Skip(); err != nil {
			it.decodeErr = err
			return
		}
	}
}

func (it *baseIterator) checkNumFieldsForType(objType objectType) (int, int, bool) {
	numActualFields := it.decodeNumObjectFields()
	numExpectedFields, ok := it.checkExpectedNumFieldsForType(objType, numActualFields)
	return numExpectedFields, numActualFields, ok
}

func (it *baseIterator) checkExpectedNumFieldsForType(
	objType objectType,
	numActualFields int,
) (int, bool) {
	if it.decodeErr != nil {
		return 0, false
	}
	numExpectedFields := numFieldsForType(objType)
	if numExpectedFields > numActualFields {
		it.decodeErr = fmt.Errorf("number of fields mismatch: expected %d actual %d", numExpectedFields, numActualFields)
		return 0, false
	}
	return numExpectedFields, true
}

// bufReader is a buffered reader.
type bufReader interface {
	io.Reader

	ReadByte() (byte, error)
	UnreadByte() error
}

func toBufReader(reader io.Reader, readerBufferSize int) bufReader {
	bufReader, ok := reader.(bufReader)
	if ok {
		return bufReader
	}
	return bufio.NewReaderSize(reader, readerBufferSize)
}
