package main

import (
	"net/http"

	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/redisutil"
	"github.com/koblas/grpc-todo/pkg/util"
	"github.com/koblas/grpc-todo/services/core/send_email"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	ssmConfig := send_email.SsmConfig{}
	if err := confmgr.Parse(&ssmConfig); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	redis := redisutil.NewTwirpRedis(util.Getenv("REDIS_ADDR", "redis:6379"))

	producer := core.NewSendEmailEventsProtobufClient(
		"topic://send-email-complete",
		redis,
	)

	s := send_email.NewSendEmailServer(producer, send_email.NewSmtpService(ssmConfig))
	mux := http.NewServeMux()
	mux.Handle(core.SendEmailServicePathPrefix, core.NewSendEmailServiceServer(s))

	mgr.StartConsumer(redis.QueueConsumer(mgr.Context(), "send-email", mux))
}
