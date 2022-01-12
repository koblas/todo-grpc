package workers

import (
	"sync"

	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/key_manager"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/twpb/core"
)

var (
	onceEmail sync.Once
)

func getEmailService(log logger.Logger) (core.SendEmailService, error) {
	svc := core.NewSendEmailServiceProtobufClient("lambda://core-send-email", awsutil.NewTwirpCallLambda())

	return svc, nil
}

// func getEmailService(log logger.Logger) (core.SendEmailServiceClient, error) {
// 	var svc core.SendEmailServiceClient

// 	if svc == nil {
// 		opts := []grpc.DialOption{
// 			grpc.WithInsecure(),
// 		}

// 		host := util.Getenv("EMAILSERVICE_ADDR", ":13009")
// 		conn, err := grpc.Dial(host, opts...)
// 		if err != nil {
// 			log.With("error", err).Info("Failed to connect to email-service")

// 			return nil, err
// 		}

// 		svc = core.NewSendEmailServiceClient(conn)
// 	}

// 	return svc, nil
// }

func getKeyManager(log logger.Logger) (key_manager.Decoder, error) {
	return key_manager.NewSecureClear(), nil
}
