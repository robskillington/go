package proto

import (
	"encoding/binary"
	"io"
)

const (
	sizeBufferSize        = 4
	defaultDataBufferSize = 1024
)

type encoder struct {
	w          io.Writer
	sizeBuffer []byte
	dataBuffer []byte
}

func NewEncoder(w io.Writer) Encoder {
	e := encoder{
		w:          w,
		sizeBuffer: make([]byte, sizeBufferSize),
		dataBuffer: make([]byte, defaultDataBufferSize),
	}
	return &e
}

func (e *encoder) Encode(m Marshaller) error {
	size, err := m.MarshalTo(e.dataBuffer)
	if err != nil {
		return err
	}
	e.growDataBuffer(size)
	e.encodeSize(size)
	if _, err := e.w.Write(e.dataBuffer[:size]); err != nil {
		return err
	}
	return nil
}

func (e *encoder) encodeSize(size int) error {
	binary.BigEndian.PutUint32(e.sizeBuffer, uint32(size))
	if _, err := e.w.Write(e.sizeBuffer); err != nil {
		return err
	}
	return nil
}

func (e *encoder) growDataBuffer(size int) {
	if len(e.dataBuffer) < size {
		// TODO return old bytes to pool, get new bytes from pool
		e.dataBuffer = make([]byte, size)
	}
}
