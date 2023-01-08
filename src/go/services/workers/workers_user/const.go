package workers_user

import "github.com/koblas/grpc-todo/gen/corepb"

const APP_NAME = "Test App"
const APP_SENDER_EMAIL = "david@koblas.com"
const APP_SENDER_NAME = "David Koblas"

func buildAppInfo(config Config) *corepb.EmailAppInfo {
	return &corepb.EmailAppInfo{
		UrlBase:     config.UrlBase,
		AppName:     APP_NAME,
		SenderName:  APP_SENDER_NAME,
		SenderEmail: APP_SENDER_EMAIL,
	}
}
