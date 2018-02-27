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

type objectType int

const (
	// Current version for encoding msg.
	msgVersion int = 1

	// Current version for encoding ack.
	ackVersion int = 1
)

// nolint: deadcode
const (
	unknownType objectType = iota

	// Root object type.
	rootObjectType

	msgType

	ackType

	// Total number of object types.
	numObjectTypes = iota
)

const (
	numRootObjectFields = 2
	numMsgFields        = 10
	numAckFields        = 1
)

// NB(xichen): use a slice instead of a map to avoid lookup overhead.
var numObjectFields []int

func numFieldsForType(objType objectType) int {
	return numObjectFields[int(objType)-1]
}

func setNumFieldsForType(objType objectType, numFields int) {
	numObjectFields[int(objType)-1] = numFields
}

func init() {
	numObjectFields = make([]int, int(numObjectTypes))

	setNumFieldsForType(rootObjectType, numRootObjectFields)
	setNumFieldsForType(msgType, numMsgFields)
	setNumFieldsForType(ackType, numAckFields)
}
