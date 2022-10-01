package workers

import (
	"encoding/base64"
	"errors"

	"github.com/koblas/grpc-todo/pkg/key_manager"
	"github.com/koblas/grpc-todo/pkg/logger"
	genpb "github.com/koblas/grpc-todo/twpb/core"
)

func decodeSecure(log logger.Logger, value *genpb.SecureValue) (string, error) {
	if value.KeyUri == "" || len(value.DataValue) == 0 || len(value.DataKey) == 0 {
		return "", nil
	}

	dataKey, err := base64.RawStdEncoding.DecodeString(value.DataKey)
	if err != nil {
		return "", err
	}
	dataValue, err := base64.RawStdEncoding.DecodeString(value.DataValue)
	if err != nil {
		return "", err
	}

	smanager, err := getKeyManager(log)
	if err != nil {
		log.Fatal("Unable to get key manager")
	}

	token, err := smanager.Decode(
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
