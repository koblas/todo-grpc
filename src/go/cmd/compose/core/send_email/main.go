package main

import (
	"net/http"
	"strings"

	"github.com/koblas/grpc-todo/cmd/compose/shared_config"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/manager"
	"github.com/koblas/grpc-todo/pkg/redisutil"
	"github.com/koblas/grpc-todo/services/core/send_email"
	"github.com/koblas/grpc-todo/twpb/core"
	"go.uber.org/zap"
)

func main() {
	mgr := manager.NewManager()
	log := mgr.Logger()

	config := send_email.Config{}
	if err := confmgr.Parse(&config, confmgr.NewJsonReader(strings.NewReader(shared_config.CONFIG))); err != nil {
		log.With(zap.Error(err)).Fatal("failed to load configuration")
	}

	redis := redisutil.NewTwirpRedis(config.RedisAddr)

	producer := core.NewSendEmailEventsProtobufClient(
		"topic://"+config.EmailSentTopic,
		redis,
	)

	s := send_email.NewSendEmailServer(producer, send_email.NewSmtpService(config))
	mux := http.NewServeMux()
	mux.Handle(core.SendEmailServicePathPrefix, core.NewSendEmailServiceServer(s))

	mgr.StartConsumer(redis.QueueConsumer(mgr.Context(), "send-email", mux))
}
