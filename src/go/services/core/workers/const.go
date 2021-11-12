package workers

import (
	"log"
	"os"

	genpb "github.com/koblas/grpc-todo/genpb/core"
)

const APP_NAME = "Test App"
const APP_SENDER_EMAIL = "david@koblas.com"
const APP_SENDER_NAME = "David Koblas"

var APP_URL_BASE string

var appInfo *genpb.EmailAppInfo

func init() {
	APP_URL_BASE := os.Getenv("URL_BASE_UI")
	if APP_URL_BASE == "" {
		log.Fatal("environment variable URL_BASE_UI must be set")
	}

	appInfo = &genpb.EmailAppInfo{
		UrlBase:     APP_URL_BASE,
		AppName:     APP_NAME,
		SenderName:  APP_SENDER_NAME,
		SenderEmail: APP_SENDER_EMAIL,
	}
}
