package workers_user

import (
	emailv1 "github.com/koblas/grpc-todo/gen/core/send_email/v1"
)

const APP_NAME = "Test App"
const APP_SENDER_EMAIL = "david@koblas.com"
const APP_SENDER_NAME = "David Koblas"

func buildAppInfo(urlBase string) *emailv1.EmailAppInfo {
	return &emailv1.EmailAppInfo{
		UrlBase:     urlBase,
		AppName:     APP_NAME,
		SenderName:  APP_SENDER_NAME,
		SenderEmail: APP_SENDER_EMAIL,
	}
}
