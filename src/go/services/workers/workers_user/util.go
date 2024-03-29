package workers_user

import (
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/protoutil"
)

func decodeSecure(log logger.Logger, value *corev1.SecureValue) (string, error) {
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
