package key_manager

type SecureValue struct {
	// The URI that references the parent key that encoded the value
	KmsUri string
	// The DataKey that will be unlocked to encode/decode the the Data
	DataKey []byte
	// The Data value (either input or output)
	Data []byte
}

type Encoder interface {
	// Encode the given value and return a appropratly wrapped payload
	Encode(value []byte) (SecureValue, error)
}

type Decoder interface {
	// Decode the given value and return the original value
	Decode(value SecureValue) ([]byte, error)
}

type KeyManager interface {
	Encoder
	Decoder
}
