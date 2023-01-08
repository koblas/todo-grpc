package workers_user

import (
	genpb "github.com/koblas/grpc-todo/twpb/core"
)

const APP_NAME = "Test App"
const APP_SENDER_EMAIL = "david@koblas.com"
const APP_SENDER_NAME = "David Koblas"

func buildAppInfo(config Config) *genpb.EmailAppInfo {
	return &genpb.EmailAppInfo{
		UrlBase:     config.UrlBase,
		AppName:     APP_NAME,
		SenderName:  APP_SENDER_NAME,
		SenderEmail: APP_SENDER_EMAIL,
	}
}
