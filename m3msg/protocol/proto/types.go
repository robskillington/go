package proto

type Marshaller interface {
	MarshalTo(data []byte) (int, error)
}

type Unmarshaller interface {
	Unmarshal(data []byte) error
}

type Encoder interface {
	Encode(m Marshaller) error
}

type Decoder interface {
	Decode(m Unmarshaller) error
}
