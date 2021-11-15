package workers

import (
	genpb "github.com/koblas/grpc-todo/genpb/core"
	"github.com/koblas/grpc-todo/pkg/key_manager"
	"github.com/koblas/grpc-todo/pkg/logger"
)

func decodeSecure(log logger.Logger, value *genpb.SecureValue) (string, error) {
	smanager, err := getKeyManager(log)
	if err != nil {
		log.Fatal("Unable to get key manager")
	}

	token, err := smanager.Decode(
		key_manager.SecureValue{
			KmsUri:  value.KeyUri,
			DataKey: []byte(value.DataKey),
			Data:    []byte(value.Value),
		},
	)

	return string(token), nil
}
