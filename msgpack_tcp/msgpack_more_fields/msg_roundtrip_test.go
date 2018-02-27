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
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func validateMsgRoundtrip(t *testing.T, inputs ...Msg) {
	encoder := testMsgEncoder()
	it := testMsgIterator(nil)
	validateMsgRoundtripWithEncoderAndIterator(t, encoder, it, inputs...)
}

func validateMsgRoundtripWithEncoderAndIterator(
	t *testing.T,
	encoder MsgEncoder,
	it MsgIterator,
	inputs ...Msg,
) {
	var (
		results []Msg
	)

	// Encode the batch of metrics.
	encoder.Reset(NewBufferedEncoder())
	for _, input := range inputs {
		require.NoError(t, testEncodeMsg(encoder, input))
	}

	// Decode the batch of metrics.
	encodedBytes := bytes.NewBuffer(encoder.Encoder().Bytes())
	it.Reset(encodedBytes)
	for it.Next() {
		msg := it.Msg()
		results = append(results, msg)
	}

	// Assert the results match expectations.
	require.Equal(t, io.EOF, it.Err())
	require.Equal(t, inputs, results)
}

func TestMsgEncodeDecodeMetricWithPolicy(t *testing.T) {
	validateMsgRoundtrip(t, testMsg)
}

func TestMsgEncodeDecodeStress(t *testing.T) {
	var (
		numIter    = 10
		numMetrics = 10000
		encoder    = testMsgEncoder()
		iterator   = testMsgIterator(nil)
	)

	for i := 0; i < numIter; i++ {
		var inputs []Msg
		for j := 0; j < numMetrics; j++ {
			inputs = append(inputs, testMsg)
		}
		validateMsgRoundtripWithEncoderAndIterator(t, encoder, iterator, inputs...)
	}
}
