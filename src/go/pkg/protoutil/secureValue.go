package protoutil

import (
	"encoding/base64"
	"errors"

	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/koblas/grpc-todo/pkg/key_manager"
)

func SecureValueEncode(encoder key_manager.Encoder, token string) (*corev1.SecureValue, error) {
	svalue, err := encoder.Encode([]byte(token))
	if err != nil {
		return nil, err
	}
	return &corev1.SecureValue{
		KeyUri:    svalue.KmsUri,
		DataKey:   base64.RawStdEncoding.EncodeToString(svalue.DataKey),
		DataValue: base64.RawStdEncoding.EncodeToString(svalue.Data),
	}, nil

}

func SecureValueDecode(decoder key_manager.Decoder, value *corev1.SecureValue) (string, error) {
	dataKey, err := base64.RawStdEncoding.DecodeString(value.DataKey)
	if err != nil {
		return "", err
	}
	dataValue, err := base64.RawStdEncoding.DecodeString(value.DataValue)
	if err != nil {
		return "", err
	}

	token, err := decoder.Decode(
		key_manager.SecureValue{
			KmsUri:  value.KeyUri,
			DataKey: dataKey,
			Data:    dataValue,
		},
	)

	if err != nil {
		return "", err
	}

	svalue := string(token)
	if svalue == "" {
		return "", errors.New("no token value")
	}

	return svalue, nil
}
