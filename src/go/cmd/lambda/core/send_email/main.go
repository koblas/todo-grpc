package main

import (
	"net/http"

	"github.com/koblas/grpc-todo/pkg/awsutil"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/services/core/send_email"
	"github.com/koblas/grpc-todo/twpb/core"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	ssmConfig := send_email.SsmConfig{}
	if err := awsutil.LoadSsmConfig("/common/", &ssmConfig); err != nil {
		log.Fatal(err.Error())
	}

	producer := core.NewSendEmailEventsProtobufClient(
		ssmConfig.EventArn,
		awsutil.NewTwirpCallLambda(),
	)

	s := send_email.NewSendEmailServer(producer, send_email.NewSmtpService(ssmConfig))
	mux := http.NewServeMux()
	mux.Handle(core.SendEmailServicePathPrefix, core.NewSendEmailServiceServer(s))

	mgr.StartConsumer(awsutil.HandleSqsLambda(mux))
}
