LDFLAGS="-s -w"

build:
	go build cmd/auth/main.go
	go build cmd/extauth/main.go
	go build cmd/todo/main.go 
	go build cmd/user/main.go
	go build cmd/workers/main.go
	go build cmd/send_email/main.go

start:
	# Front end
	PORT=14002 go run cmd/auth/main.go
	PORT=14001 go run cmd/todo/main.go 
	PORT=14010 go run cmd/extauth/main.go
	# not a gRPC server
	go run cmd/workers/main.go
	# Backend
	PORT=13001 go run cmd/user/main.go
	PORT=13009 go run cmd/send_email/main.go


build-lambda:
	GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o build/publicapi-auth ./lambda/publicapi/auth/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o build/publicapi-todo ./lambda/publicapi/todo/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o build/core-todo ./lambda/core/todo/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o build/core-user ./lambda/core/user/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o build/core-oauth-user ./lambda/core/oauth_user/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o build/core-send-email ./lambda/core/send_email/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o build/core-workers ./lambda/core/workers/main.go

run-lambda: build-lambda
	sam local start-api --template-file ./lambda/template.yaml --env-vars ../../dev-variables.json

deploy: build-lambda
	( cd deploy && npx cdk deploy)
