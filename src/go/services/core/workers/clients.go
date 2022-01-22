package workers

import (
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/key_manager"
	"github.com/koblas/grpc-todo/pkg/logger"
	"github.com/koblas/grpc-todo/twpb/core"
)

func getEmailService(log logger.Logger) (core.SendEmailService, error) {
	svc := core.NewSendEmailServiceProtobufClient("sqs://send-email", awsutil.NewTwirpCallLambda())

	return svc, nil
}

func getKeyManager(log logger.Logger) (key_manager.Decoder, error) {
	return key_manager.NewSecureClear(), nil
}
