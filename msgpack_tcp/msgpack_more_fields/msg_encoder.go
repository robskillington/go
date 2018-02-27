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

type encodeRootObjectFn func(objType objectType)
type encodeMsgFn func(msg Msg)

type msgEncoder struct {
	encoderBase

	encodeRootObjectFn encodeRootObjectFn
	encodeMsgFn        encodeMsgFn
}

// NewMsgEncoder creates a msg encoder.
func NewMsgEncoder(encoder BufferedEncoder) MsgEncoder {
	enc := &msgEncoder{
		encoderBase: newBaseEncoder(encoder),
	}

	enc.encodeRootObjectFn = enc.encodeRootObject
	enc.encodeMsgFn = enc.encodeMsg

	return enc
}

func (enc *msgEncoder) Encoder() BufferedEncoder      { return enc.encoder() }
func (enc *msgEncoder) Reset(encoder BufferedEncoder) { enc.reset(encoder) }

func (enc *msgEncoder) EncodeMsg(
	msg Msg,
) error {
	if err := enc.err(); err != nil {
		return err
	}
	enc.encodeRootObjectFn(msgType)
	enc.encodeMsgFn(msg)
	return enc.err()
}

func (enc *msgEncoder) encodeRootObject(objType objectType) {
	enc.encodeVersion(msgVersion)
	enc.encodeNumObjectFields(numFieldsForType(rootObjectType))
	enc.encodeObjectType(objType)
}

func (enc *msgEncoder) encodeMsg(
	msg Msg,
) {
	enc.encodeNumObjectFields(numFieldsForType(msgType))
	enc.encodeVarint(msg.Offset)
	enc.encodeVarint(msg.Offset2)
	enc.encodeVarint(msg.Offset3)
	enc.encodeVarint(msg.Offset4)
	enc.encodeVarint(msg.Offset5)
	enc.encodeVarint(msg.Offset6)
	enc.encodeVarint(msg.Offset7)
	enc.encodeVarint(msg.Offset8)
	enc.encodeVarint(msg.Offset9)
	enc.encodeBytes(msg.Value)
}
