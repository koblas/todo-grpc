package workers

import (
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/protoutil"
	genpb "github.com/koblas/grpc-todo/twpb/core"
)

func decodeSecure(log logger.Logger, value *genpb.SecureValue) (string, error) {
	decoder, err := getKeyManager(log)
	if err != nil {
		log.Fatal("Unable to get key manager")
	}
	token, err := protoutil.DecodeSecure(decoder, value)
	if err != nil {
		return "", err
	}

	return token, nil
}
