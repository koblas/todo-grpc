LDFLAGS="-s -w"

REGISTRY=`minikube ip`:5000

all: build-lambda build-compose

proto:
	buf generate ../../protos/

start:
	# Front end
	PORT=14002 go run cmd/auth/main.go
	PORT=14001 go run cmd/todo/main.go 
	PORT=14010 go run cmd/extauth/main.go
	# not a gRPC server
	go run cmd/workers_user/main.go
	# Backend
	PORT=13001 go run cmd/user/main.go
	PORT=13009 go run cmd/send_email/main.go


build-lambda:
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o build/publicapi-websocket ./cmd/lambda/publicapi/websocket/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o build/publicapi-auth ./cmd/lambda/publicapi/auth/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o build/publicapi-file ./cmd/lambda/publicapi/file/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o build/publicapi-gpt ./cmd/lambda/publicapi/user/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o build/publicapi-todo ./cmd/lambda/publicapi/todo/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o build/publicapi-user ./cmd/lambda/publicapi/user/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o build/websocket-todo ./cmd/lambda/websocket/todo/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o build/websocket-user ./cmd/lambda/websocket/user/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o build/websocket-broadcast ./cmd/lambda/websocket/broadcast/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o build/core-todo ./cmd/lambda/core/todo/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o build/core-user ./cmd/lambda/core/user/main.go
	# GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o build/core-file ./cmd/lambda/core/file/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o build/core-oauth-user ./cmd/lambda/core/oauth_user/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o build/core-send-email ./cmd/lambda/core/send_email/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o build/workers-user ./cmd/lambda/workers/workers_user/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o build/workers-file ./cmd/lambda/workers/workers_file/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o build/trigger-s3 ./cmd/lambda/trigger/s3/main.go

build-compose:
	go build -o build/publicapi-auth ./cmd/compose/publicapi/auth/main.go
	go build -o build/publicapi-file ./cmd/compose/publicapi/file/main.go
	go build -o build/publicapi-gpt ./cmd/compose/publicapi/gpt/main.go
	go build -o build/publicapi-user ./cmd/compose/publicapi/user/main.go
	go build -o build/publicapi-todo ./cmd/compose/publicapi/todo/main.go
	go build -o build/publicapi-websocket ./cmd/compose/publicapi/websocket/main.go
	go build -o build/websocket-todo ./cmd/compose/websocket/todo/main.go
	go build -o build/websocket-user ./cmd/compose/websocket/user/main.go
	go build -o build/websocket-broadcast ./cmd/compose/websocket/broadcast/main.go
	go build -o build/core-todo ./cmd/compose/core/todo/main.go
	go build -o build/core-user ./cmd/compose/core/user/main.go
	# go build -o build/core-file ./cmd/compose/core/file/main.go
	go build -o build/core-oauth-user ./cmd/compose/core/oauth_user/main.go
	go build -o build/core-send-email ./cmd/compose/core/send_email/main.go
	go build -o build/workers-user ./cmd/compose/workers/workers_user/main.go
	go build -o build/workers-file ./cmd/compose/workers/workers_file/main.go

run-lambda: build-lambda
	sam local start-api --template-file ./lambda/template.yaml --env-vars ../../dev-variables.json

deploy: 
	( cd deploy && npx cdk deploy)

minikube:
	minikube image build -f ./cmd/compose/publicapi/todo/Dockerfile --tag todo-grpc/publicapi-todo  . 
	minikube image build -f ./cmd/compose/publicapi/auth/Dockerfile --tag todo-grpc/publicapi-auth .
	minikube image build -f ./cmd/compose/publicapi/websocket/Dockerfile --tag todo-grpc/publicapi-websocket .
	minikube image build -f ./cmd/compose/websocket/broadcast/Dockerfile --tag todo-grpc/websocket-broadcast .
	minikube image build -f ./cmd/compose/websocket/todo/Dockerfile --tag todo-grpc/websocket-todo .
	minikube image build -f ./cmd/compose/websocket/user/Dockerfile --tag todo-grpc/websocket-user .
	minikube image build -f ./cmd/compose/core/todo/Dockerfile --tag todo-grpc/core-todo .
	minikube image build -f ./cmd/compose/core/user/Dockerfile --tag todo-grpc/core-user .
	minikube image build -f ./cmd/compose/core/oauth_user/Dockerfile --tag todo-grpc/core-oauth-user .
	minikube image build -f ./cmd/compose/core/send_email/Dockerfile --tag todo-grpc/core-send-email .
	minikube image build -f ./cmd/compose/workers/workers_user/Dockerfile --tag todo-grpc/workers-user .

test:
	go test ./...


.PHONY: deploy
