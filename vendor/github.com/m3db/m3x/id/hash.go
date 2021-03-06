// Copyright (c) 2017 Uber Technologies, Inc.
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

// Package id provides utilities for generating ID's from hash functions.
package id

import (
	"crypto/md5"

	"github.com/spaolacci/murmur3"
)

// Hash represents a form of ID suitable to be used as map keys.
type Hash [md5.Size]byte

// HashFn is the default hashing implementation for IDs.
func HashFn(data []byte) Hash {
	return md5.Sum(data)
}

// Hash128 is a 128-bit hash of an ID consisting of two unsigned 64-bit ints.
type Hash128 [2]uint64

// Murmur3Hash128 computes the 128-bit hash of an id.
func Murmur3Hash128(data []byte) Hash128 {
	h0, h1 := murmur3.Sum128(data)
	return Hash128{h0, h1}
}
