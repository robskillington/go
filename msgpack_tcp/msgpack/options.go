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
	"math"

	xpool "github.com/m3db/m3x/pool"
)

const (
	// The maximum capacity of buffers that can be returned to the buffered
	// encoder pool.
	defaultBufferedEncoderPoolMaxCapacity = math.MaxInt64

	// Whether the iterator should ignore higher-than-supported version
	// by default for unaggregated iterator.
	defaultUnaggregatedIgnoreHigherVersion = false

	// Default reader buffer size for the unaggregated iterator.
	defaultUnaggregatedReaderBufferSize = 1440

	// Whether a float slice is considered a "large" slice and therefore
	// resort to the pool for allocating that slice.
	defaultLargeFloatsSize = 1024

	// Whether the iterator should ignore higher-than-supported version
	// by default for aggregated iterator.
	defaultAggregatedIgnoreHigherVersion = false

	// Default reader buffer size for the aggregated iterator.
	defaultAggregatedReaderBufferSize = 1440
)

type bufferedEncoderPoolOptions struct {
	maxCapacity int
	poolOpts    xpool.ObjectPoolOptions
}

// NewBufferedEncoderPoolOptions creates a new set of buffered encoder pool options.
func NewBufferedEncoderPoolOptions() BufferedEncoderPoolOptions {
	return &bufferedEncoderPoolOptions{
		maxCapacity: defaultBufferedEncoderPoolMaxCapacity,
		poolOpts:    xpool.NewObjectPoolOptions(),
	}
}

func (o *bufferedEncoderPoolOptions) SetMaxCapacity(value int) BufferedEncoderPoolOptions {
	opts := *o
	opts.maxCapacity = value
	return &opts
}

func (o *bufferedEncoderPoolOptions) MaxCapacity() int {
	return o.maxCapacity
}

func (o *bufferedEncoderPoolOptions) SetObjectPoolOptions(value xpool.ObjectPoolOptions) BufferedEncoderPoolOptions {
	opts := *o
	opts.poolOpts = value
	return &opts
}

func (o *bufferedEncoderPoolOptions) ObjectPoolOptions() xpool.ObjectPoolOptions {
	return o.poolOpts
}

type msgIteratorOptions struct {
	ignoreHigherVersion bool
	readerBufferSize    int
	iteratorPool        MsgIteratorPool
}

// NewMsgIteratorOptions creates a new set of aggregated iterator options.
func NewMsgIteratorOptions() MsgIteratorOptions {
	return &msgIteratorOptions{
		ignoreHigherVersion: defaultAggregatedIgnoreHigherVersion,
		readerBufferSize:    defaultAggregatedReaderBufferSize,
	}
}

func (o *msgIteratorOptions) SetIgnoreHigherVersion(value bool) MsgIteratorOptions {
	opts := *o
	opts.ignoreHigherVersion = value
	return &opts
}

func (o *msgIteratorOptions) IgnoreHigherVersion() bool {
	return o.ignoreHigherVersion
}

func (o *msgIteratorOptions) SetReaderBufferSize(value int) MsgIteratorOptions {
	opts := *o
	opts.readerBufferSize = value
	return &opts
}

func (o *msgIteratorOptions) ReaderBufferSize() int {
	return o.readerBufferSize
}

func (o *msgIteratorOptions) SetIteratorPool(value MsgIteratorPool) MsgIteratorOptions {
	opts := *o
	opts.iteratorPool = value
	return &opts
}

func (o *msgIteratorOptions) IteratorPool() MsgIteratorPool {
	return o.iteratorPool
}
