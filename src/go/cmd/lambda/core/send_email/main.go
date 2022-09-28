package main

import (
	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/eventbus/aws"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/core/send_email"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	ssmConfig := send_email.SsmConfig{}
	if err := awsutil.LoadSsmConfig("/common/", &ssmConfig); err != nil {
		log.Fatal(err.Error())
	}

	producer, err := aws.NewAwsProducer(ssmConfig.EventArn)
	if err != nil {
		log.With(zap.Error(err)).Fatal("unable to create publisher")
	}

	s := send_email.NewSendEmailServer(producer, send_email.NewSmtpService(ssmConfig))

	handlers := awsutil.SqsHandlers{
		core.SendEmailServicePathPrefix: core.NewSendEmailServiceServer(s),
	}

	mgr.StartConsumer(awsutil.HandleSqsLambda(handlers, nil))
}
