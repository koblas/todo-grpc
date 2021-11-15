package key_manager

type clearKms struct {
	Encoder
	Decoder
}

func NewSecureClear() KeyManager {
	return &clearKms{}
}

func (*clearKms) Encode(value []byte) (SecureValue, error) {
	return SecureValue{
		KmsUri:  "",
		DataKey: []byte{},
		Data:    value,
	}, nil
}

func (*clearKms) Decode(value SecureValue) ([]byte, error) {
	return value.Data, nil
}
