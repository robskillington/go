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
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

var testMsg = Msg{
	Offset: 1,
	Value:  []byte{'a', 'b', 'c'},
}

func validateMsgDecodeResults(
	t *testing.T,
	it MsgIterator,
	expectedResults []Msg,
	expectedErr error,
) {
	var results []Msg
	for it.Next() {
		msg := it.Msg()
		results = append(results, msg)
	}
	require.Equal(t, expectedErr, it.Err())
	require.Equal(t, expectedResults, results)
}

func testMsgEncoder() MsgEncoder {
	return NewMsgEncoder(NewBufferedEncoder())
}

func testMsgIterator(reader io.Reader) MsgIterator {
	return NewMsgIterator(reader, NewMsgIteratorOptions())
}

func testEncodeMsg(
	encoder MsgEncoder,
	m Msg,
) error {
	return encoder.EncodeMsg(m)
}

func TestMsgIteratorDecodeNewerVersionThanSupported(t *testing.T) {
	input := testMsg
	enc := testMsgEncoder().(*msgEncoder)

	// Version encoded is higher than supported version.
	enc.encodeRootObjectFn = func(objType objectType) {
		enc.encodeVersion(msgVersion + 1)
		enc.encodeNumObjectFields(numFieldsForType(rootObjectType))
		enc.encodeObjectType(objType)
	}
	require.NoError(t, testEncodeMsg(enc, input))

	// Now restore the encode top-level function and encode another metric.
	enc.encodeRootObjectFn = enc.encodeRootObject
	require.NoError(t, testEncodeMsg(enc, input))

	it := testMsgIterator(enc.Encoder().Buffer())
	it.(*msgIterator).ignoreHigherVersion = true

	// Check that we skipped the first metric and successfully decoded the second metric.
	validateMsgDecodeResults(t, it, []Msg{input}, io.EOF)
}

func TestMsgIteratorDecodeRootObjectMoreFieldsThanExpected(t *testing.T) {
	input := testMsg
	enc := testMsgEncoder().(*msgEncoder)

	// Pretend we added an extra int field to the root object.
	enc.encodeRootObjectFn = func(objType objectType) {
		enc.encodeVersion(msgVersion)
		enc.encodeNumObjectFields(numFieldsForType(rootObjectType) + 1)
		enc.encodeObjectType(objType)
	}
	err := testEncodeMsg(enc, input)
	require.NoError(t, err)
	enc.encodeVarint(0)
	require.NoError(t, enc.err())

	it := testMsgIterator(enc.Encoder().Buffer())

	// Check that we successfully decoded the metric.
	validateMsgDecodeResults(t, it, []Msg{input}, io.EOF)
}

func TestMsgIteratorClose(t *testing.T) {
	it := NewMsgIterator(nil, nil)
	it.Close()
	require.False(t, it.Next())
	require.NoError(t, it.Err())
	require.True(t, it.(*msgIterator).closed)
}
