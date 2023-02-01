package workers_user

import (
	corepbv1 "github.com/koblas/grpc-todo/gen/corepb/v1"
)

const APP_NAME = "Test App"
const APP_SENDER_EMAIL = "david@koblas.com"
const APP_SENDER_NAME = "David Koblas"

func buildAppInfo(config Config) *corepbv1.EmailAppInfo {
	return &corepbv1.EmailAppInfo{
		UrlBase:     config.UrlBase,
		AppName:     APP_NAME,
		SenderName:  APP_SENDER_NAME,
		SenderEmail: APP_SENDER_EMAIL,
	}
}
