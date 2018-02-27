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
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	errTestVarint   = errors.New("test varint error")
	errTestFloat64  = errors.New("test float64 error")
	errTestBytes    = errors.New("test bytes error")
	errTestArrayLen = errors.New("test array len error")
)

func testCapturingBaseEncoder(encoder encoderBase) *[]interface{} {
	baseEncoder := encoder.(*baseEncoder)

	var result []interface{}
	baseEncoder.encodeVarintFn = func(value int64) {
		result = append(result, value)
	}
	baseEncoder.encodeBoolFn = func(value bool) {
		result = append(result, value)
	}
	baseEncoder.encodeFloat64Fn = func(value float64) {
		result = append(result, value)
	}
	baseEncoder.encodeBytesFn = func(value []byte) {
		result = append(result, value)
	}
	baseEncoder.encodeBytesLenFn = func(value int) {
		result = append(result, value)
	}
	baseEncoder.encodeArrayLenFn = func(value int) {
		result = append(result, value)
	}

	return &result
}

func expectedResultsForMsg(p Msg) []interface{} {
	results := []interface{}{
		int64(msgVersion),
		numFieldsForType(rootObjectType),
		int64(msgType),
		numFieldsForType(msgType),
		p.Offset,
		p.Value,
	}

	return results
}

func testCapturingAggregatedEncoder() (MsgEncoder, *[]interface{}) {
	encoder := testMsgEncoder().(*msgEncoder)
	result := testCapturingBaseEncoder(encoder.encoderBase)
	return encoder, result
}

func TestEncodeMsg(t *testing.T) {
	encoder, results := testCapturingAggregatedEncoder()
	require.NoError(t, testEncodeMsg(encoder, testMsg))
	expected := expectedResultsForMsg(testMsg)
	require.Equal(t, expected, *results)
}

func TestEncodeMsgError(t *testing.T) {
	// Intentionally return an error when encoding varint.
	encoder := testMsgEncoder().(*msgEncoder)
	baseEncoder := encoder.encoderBase.(*baseEncoder)
	baseEncoder.encodeVarintFn = func(value int64) {
		baseEncoder.encodeErr = errTestVarint
	}

	// Assert the error is expected.
	require.Equal(t, errTestVarint, testEncodeMsg(encoder, testMsg))

	// Assert re-encoding doesn't change the error.
	require.Equal(t, errTestVarint, testEncodeMsg(encoder, testMsg))
}

func TestAggregatedEncoderReset(t *testing.T) {
	encoder := testMsgEncoder().(*msgEncoder)
	baseEncoder := encoder.encoderBase.(*baseEncoder)
	baseEncoder.encodeErr = errTestVarint
	require.Equal(t, errTestVarint, testEncodeMsg(encoder, testMsg))

	encoder.Reset(NewBufferedEncoder())
	require.NoError(t, testEncodeMsg(encoder, testMsg))
}
