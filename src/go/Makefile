LDFLAGS="-s -w"

REGISTRY=`minikube ip`:5000

all: build-lambda build-compose

#build:
#	go build cmd/auth/main.go
#	go build cmd/extauth/main.go
#	go build cmd/todo/main.go 
#	go build cmd/user/main.go
#	go build cmd/workers/main.go
#	go build cmd/send_email/main.go

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
	GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o build/publicapi-auth ./cmd/lambda/publicapi/websocket/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o build/publicapi-auth ./cmd/lambda/publicapi/auth/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o build/websocket-todo ./cmd/lambda/websocket/todo/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o build/core-todo ./cmd/lambda/core/todo/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o build/core-user ./cmd/lambda/core/user/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o build/core-oauth-user ./cmd/lambda/core/oauth_user/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o build/core-send-email ./cmd/lambda/core/send_email/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o build/core-workers ./cmd/lambda/core/workers/main.go

build-compose:
	go build -o build/publicapi-auth ./cmd/compose/publicapi/auth/main.go
	go build -o build/publicapi-todo ./cmd/compose/publicapi/todo/main.go
	go build -o build/publicapi-websocket ./cmd/compose/publicapi/websocket/main.go
	go build -o build/websocket-todo ./cmd/compose/websocket/todo/main.go
	go build -o build/core-todo ./cmd/compose/core/todo/main.go
	go build -o build/core-user ./cmd/compose/core/user/main.go
	go build -o build/core-oauth-user ./cmd/compose/core/oauth_user/main.go
	go build -o build/core-send-email ./cmd/compose/core/send_email/main.go
	go build -o build/core-workers ./cmd/compose/core/workers/main.go

run-lambda: build-lambda
	sam local start-api --template-file ./lambda/template.yaml --env-vars ../../dev-variables.json

deploy: 
	( cd deploy && npx cdk deploy)

minikube:
	minikube image build -f ./cmd/compose/publicapi/todo/Dockerfile --tag todo-grpc/publicapi-todo  . 
	minikube image build -f ./cmd/compose/publicapi/auth/Dockerfile --tag todo-grpc/publicapi-auth .
	minikube image build -f ./cmd/compose/publicapi/websocket/Dockerfile --tag todo-grpc/publicapi-websocket .
	minikube image build -f ./cmd/compose/websocket/todo/Dockerfile --tag todo-grpc/websocket-todo .
	minikube image build -f ./cmd/compose/core/todo/Dockerfile --tag todo-grpc/core-todo .
	minikube image build -f ./cmd/compose/core/user/Dockerfile --tag todo-grpc/core-user .
	minikube image build -f ./cmd/compose/core/oauth_user/Dockerfile --tag todo-grpc/core-oauth-user .
	minikube image build -f ./cmd/compose/core/send_email/Dockerfile --tag todo-grpc/core-send-email .
	minikube image build -f ./cmd/compose/core/workers/Dockerfile --tag todo-grpc/core-workers .


.PHONY: deploy
