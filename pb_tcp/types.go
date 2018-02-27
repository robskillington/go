package pb_tcp

import (
	"encoding/binary"
	"io"
)

type Marshaller interface {
	MarshalTo(data []byte) (int, error)
}

type Unmarshaller interface {
	Unmarshal(data []byte) error
}

func Encode(m Marshaller, sizeBuffer []byte, dataBuffer []byte, w io.Writer) error {
	size, err := m.MarshalTo(dataBuffer)
	if err != nil {
		return err
	}
	binary.BigEndian.PutUint32(sizeBuffer, uint32(size))
	if _, err := w.Write(sizeBuffer); err != nil {
		return err
	}
	if _, err := w.Write(dataBuffer[:size]); err != nil {
		return err
	}
	return nil
}

func Decode(m Unmarshaller, sizeBuffer []byte, dataBuffer []byte, r io.Reader) error {
	if _, err := io.ReadFull(r, sizeBuffer); err != nil {
		return err
	}
	size := binary.BigEndian.Uint32(sizeBuffer)
	if _, err := io.ReadFull(r, dataBuffer[:size]); err != nil {
		return err
	}
	if err := m.Unmarshal(dataBuffer[:size]); err != nil {
		return err
	}
	return nil
}
