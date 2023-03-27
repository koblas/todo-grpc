package workers_user

import (
	corev1 "github.com/koblas/grpc-todo/gen/core/v1"
)

const APP_NAME = "Test App"
const APP_SENDER_EMAIL = "david@koblas.com"
const APP_SENDER_NAME = "David Koblas"

func buildAppInfo(urlBase string) *corev1.EmailAppInfo {
	return &corev1.EmailAppInfo{
		UrlBase:     urlBase,
		AppName:     APP_NAME,
		SenderName:  APP_SENDER_NAME,
		SenderEmail: APP_SENDER_EMAIL,
	}
}
