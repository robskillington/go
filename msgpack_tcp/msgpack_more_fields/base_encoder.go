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

type encodeVarintFn func(value int64)
type encodeBoolFn func(value bool)
type encodeFloat64Fn func(value float64)
type encodeBytesFn func(value []byte)
type encodeBytesLenFn func(value int)
type encodeArrayLenFn func(value int)

// baseEncoder is the base encoder that provides common encoding APIs.
type baseEncoder struct {
	bufEncoder       BufferedEncoder
	encodeErr        error
	encodeVarintFn   encodeVarintFn
	encodeBoolFn     encodeBoolFn
	encodeFloat64Fn  encodeFloat64Fn
	encodeBytesFn    encodeBytesFn
	encodeBytesLenFn encodeBytesLenFn
	encodeArrayLenFn encodeArrayLenFn
}

func newBaseEncoder(encoder BufferedEncoder) encoderBase {
	enc := &baseEncoder{bufEncoder: encoder}

	enc.encodeVarintFn = enc.encodeVarintInternal
	enc.encodeBoolFn = enc.encodeBoolInternal
	enc.encodeFloat64Fn = enc.encodeFloat64Internal
	enc.encodeBytesFn = enc.encodeBytesInternal
	enc.encodeBytesLenFn = enc.encodeBytesLenInternal
	enc.encodeArrayLenFn = enc.encodeArrayLenInternal

	return enc
}

func (enc *baseEncoder) encoder() BufferedEncoder            { return enc.bufEncoder }
func (enc *baseEncoder) err() error                          { return enc.encodeErr }
func (enc *baseEncoder) resetData()                          { enc.bufEncoder.Reset() }
func (enc *baseEncoder) encodeVersion(version int)           { enc.encodeVarint(int64(version)) }
func (enc *baseEncoder) encodeObjectType(objType objectType) { enc.encodeVarint(int64(objType)) }
func (enc *baseEncoder) encodeNumObjectFields(numFields int) { enc.encodeArrayLen(numFields) }
func (enc *baseEncoder) encodeOffset(offset int64)           { enc.encodeVarintFn(offset) }
func (enc *baseEncoder) encodeValue(value []byte)            { enc.encodeBytesFn(value) }
func (enc *baseEncoder) encodeVarint(value int64)            { enc.encodeVarintFn(value) }
func (enc *baseEncoder) encodeBool(value bool)               { enc.encodeBoolFn(value) }
func (enc *baseEncoder) encodeFloat64(value float64)         { enc.encodeFloat64Fn(value) }
func (enc *baseEncoder) encodeBytes(value []byte)            { enc.encodeBytesFn(value) }
func (enc *baseEncoder) encodeBytesLen(value int)            { enc.encodeBytesLenFn(value) }
func (enc *baseEncoder) encodeArrayLen(value int)            { enc.encodeArrayLenFn(value) }

func (enc *baseEncoder) reset(encoder BufferedEncoder) {
	enc.bufEncoder = encoder
	enc.encodeErr = nil
}

// NB(xichen): the underlying msgpack encoder implementation
// always cast an integer value to an int64 and encodes integer
// values as varints, regardless of the actual integer type.
func (enc *baseEncoder) encodeVarintInternal(value int64) {
	if enc.encodeErr != nil {
		return
	}
	enc.encodeErr = enc.bufEncoder.EncodeInt64(value)
}

func (enc *baseEncoder) encodeBoolInternal(value bool) {
	if enc.encodeErr != nil {
		return
	}
	enc.encodeErr = enc.bufEncoder.EncodeBool(value)
}

func (enc *baseEncoder) encodeFloat64Internal(value float64) {
	if enc.encodeErr != nil {
		return
	}
	enc.encodeErr = enc.bufEncoder.EncodeFloat64(value)
}

func (enc *baseEncoder) encodeBytesInternal(value []byte) {
	if enc.encodeErr != nil {
		return
	}
	enc.encodeErr = enc.bufEncoder.EncodeBytes(value)
}

func (enc *baseEncoder) encodeBytesLenInternal(value int) {
	if enc.encodeErr != nil {
		return
	}
	enc.encodeErr = enc.bufEncoder.EncodeBytesLen(value)
}

func (enc *baseEncoder) encodeArrayLenInternal(value int) {
	if enc.encodeErr != nil {
		return
	}
	enc.encodeErr = enc.bufEncoder.EncodeArrayLen(value)
}

func (enc *baseEncoder) writeRaw(buf []byte) {
	if enc.encodeErr != nil {
		return
	}
	_, enc.encodeErr = enc.bufEncoder.Buffer().Write(buf)
}
