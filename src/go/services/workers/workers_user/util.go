package workers_user

import (
	"github.com/koblas/grpc-todo/gen/corepb"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/protoutil"
)

func decodeSecure(log logger.Logger, value *corepb.SecureValue) (string, error) {
	decoder, err := getKeyManager(log)
	if err != nil {
		log.Fatal("Unable to get key manager")
	}
	token, err := protoutil.SecureValueDecode(decoder, value)
	if err != nil {
		return "", err
	}

	return token, nil
}
