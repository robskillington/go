package proto

import (
	"encoding/binary"
	"io"
)

type decoder struct {
	r          io.Reader
	sizeBuffer []byte
	dataBuffer []byte
}

func NewDecoder(r io.Reader) Decoder {
	d := decoder{
		r:          r,
		sizeBuffer: make([]byte, sizeBufferSize),
		dataBuffer: make([]byte, defaultDataBufferSize),
	}
	return &d
}

func (d *decoder) Decode(m Unmarshaller) error {
	size, err := d.decodeSize()
	if err != nil {
		return err
	}
	d.growDataBuffer(int(size))
	if _, err := io.ReadFull(d.r, d.dataBuffer[:size]); err != nil {
		return err
	}
	if err := m.Unmarshal(d.dataBuffer[:size]); err != nil {
		return err
	}
	return nil
}

func (d *decoder) decodeSize() (int, error) {
	if _, err := io.ReadFull(d.r, d.sizeBuffer); err != nil {
		return 0, err
	}
	size := binary.BigEndian.Uint32(d.sizeBuffer)
	return int(size), nil
}

func (d *decoder) growDataBuffer(size int) {
	if len(d.dataBuffer) < size {
		// TODO return old bytes to pool, get new bytes from pool
		d.dataBuffer = make([]byte, size)
	}
}
