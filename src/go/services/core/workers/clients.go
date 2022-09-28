package workers

import (
	"github.com/koblas/grpc-todo/pkg/key_manager"
	"github.com/koblas/grpc-todo/pkg/logger"
)

func getKeyManager(log logger.Logger) (key_manager.Decoder, error) {
	return key_manager.NewSecureClear(), nil
}
