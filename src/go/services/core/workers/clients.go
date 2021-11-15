package workers

import (
	"sync"

	"github.com/koblas/grpc-todo/genpb/core"
	"github.com/koblas/grpc-todo/pkg/key_manager"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/pkg/util"
	"google.golang.org/grpc"
)

var (
	onceEmail sync.Once
)

func getEmailService(log logger.Logger) (core.SendEmailServiceClient, error) {
	var svc core.SendEmailServiceClient

	if svc == nil {
		opts := []grpc.DialOption{
			grpc.WithInsecure(),
		}

		host := util.Getenv("EMAILSERVICE_ADDR", ":13009")
		conn, err := grpc.Dial(host, opts...)
		if err != nil {
			log.With("error", err).Info("Failed to connect to email-service")

			return nil, err
		}

		svc = core.NewSendEmailServiceClient(conn)
	}

	return svc, nil
}

func getKeyManager(log logger.Logger) (key_manager.Decoder, error) {
	return key_manager.NewSecureClear(), nil
}
