LDFLAGS="-s -w"

REGISTRY=`minikube ip`:5000

BUILDDIR=build

all: build-lambda build-compose

proto:
	buf generate ../../protos/
	# (cd gen/core/v1/corev1connect && minimock -s _mock.go )
	(cd gen/core/eventbus/v1/eventbusv1connect && minimock -s _mock.go )
	(cd gen/core/file/v1/filev1connect && minimock -s _mock.go )
	(cd gen/core/message/v1/messagev1connect && minimock -s _mock.go )
	(cd gen/core/oauth_user/v1/oauth_userv1connect && minimock -s _mock.go )
	(cd gen/core/send_email/v1/send_emailv1connect && minimock -s _mock.go )
	(cd gen/core/todo/v1/todov1connect && minimock -s _mock.go )
	(cd gen/core/user/v1/userv1connect && minimock -s _mock.go )
	# (cd gen/core/websocket/v1/websocketv1connect && minimock -s _mock.go )


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
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/publicapi-auth ./cmd/lambda/publicapi/auth/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/publicapi-file ./cmd/lambda/publicapi/file/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/publicapi-gpt ./cmd/lambda/publicapi/user/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/publicapi-message ./cmd/lambda/publicapi/message/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/publicapi-todo ./cmd/lambda/publicapi/todo/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/publicapi-user ./cmd/lambda/publicapi/user/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/publicapi-websocket ./cmd/lambda/publicapi/websocket/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/websocket-file ./cmd/lambda/websocket/file/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/websocket-message ./cmd/lambda/websocket/message/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/websocket-todo ./cmd/lambda/websocket/todo/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/websocket-user ./cmd/lambda/websocket/user/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/websocket-broadcast ./cmd/lambda/websocket/broadcast/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/core-oauth-user ./cmd/lambda/core/oauth_user/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/core-send-email ./cmd/lambda/core/send_email/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/core-message ./cmd/lambda/core/message/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/core-todo ./cmd/lambda/core/todo/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/core-user ./cmd/lambda/core/user/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/workers-user ./cmd/lambda/workers/workers_user/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/workers-file ./cmd/lambda/workers/workers_file/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILDDIR)/trigger-s3 ./cmd/lambda/trigger/s3/main.go

build-compose:
	go build -o $(BUILDDIR)/publicapi-auth ./cmd/compose/publicapi/auth/main.go
	go build -o $(BUILDDIR)/publicapi-file ./cmd/compose/publicapi/file/main.go
	go build -o $(BUILDDIR)/publicapi-gpt ./cmd/compose/publicapi/gpt/main.go
	go build -o $(BUILDDIR)/publicapi-message ./cmd/compose/publicapi/message/main.go
	go build -o $(BUILDDIR)/publicapi-todo ./cmd/compose/publicapi/todo/main.go
	go build -o $(BUILDDIR)/publicapi-user ./cmd/compose/publicapi/user/main.go
	go build -o $(BUILDDIR)/publicapi-websocket ./cmd/compose/publicapi/websocket/main.go
	go build -o $(BUILDDIR)/websocket-file ./cmd/compose/websocket/file/main.go
	go build -o $(BUILDDIR)/websocket-message ./cmd/compose/websocket/message/main.go
	go build -o $(BUILDDIR)/websocket-todo ./cmd/compose/websocket/todo/main.go
	go build -o $(BUILDDIR)/websocket-user ./cmd/compose/websocket/user/main.go
	go build -o $(BUILDDIR)/websocket-broadcast ./cmd/compose/websocket/broadcast/main.go
	go build -o $(BUILDDIR)/core-oauth-user ./cmd/compose/core/oauth_user/main.go
	go build -o $(BUILDDIR)/core-message ./cmd/compose/core/message/main.go
	go build -o $(BUILDDIR)/core-send-email ./cmd/compose/core/send_email/main.go
	go build -o $(BUILDDIR)/core-todo ./cmd/compose/core/todo/main.go
	go build -o $(BUILDDIR)/core-user ./cmd/compose/core/user/main.go
	go build -o $(BUILDDIR)/workers-user ./cmd/compose/workers/workers_user/main.go
	go build -o $(BUILDDIR)/workers-file ./cmd/compose/workers/workers_file/main.go

build-lambda-assets: build-lambda
	( cd $(BUILDDIR) ; ln -f publicapi-auth bootstrap ; zip publicapi-auth.zip bootstrap ; rm bootstrap )
	( cd $(BUILDDIR) ; ln -f publicapi-file bootstrap ; zip publicapi-file.zip bootstrap ; rm bootstrap )
	( cd $(BUILDDIR) ; ln -f publicapi-gpt bootstrap ; zip publicapi-gpt.zip bootstrap ; rm bootstrap )
	( cd $(BUILDDIR) ; ln -f publicapi-message bootstrap ; zip publicapi-message.zip bootstrap ; rm bootstrap )
	( cd $(BUILDDIR) ; ln -f publicapi-todo bootstrap ; zip publicapi-todo.zip bootstrap ; rm bootstrap )
	( cd $(BUILDDIR) ; ln -f publicapi-user bootstrap ; zip publicapi-user.zip bootstrap ; rm bootstrap )
	( cd $(BUILDDIR) ; ln -f publicapi-websocket bootstrap ; zip publicapi-websocket.zip bootstrap ; rm bootstrap )
	( cd $(BUILDDIR) ; ln -f websocket-file bootstrap ; zip websocket-file.zip bootstrap ; rm bootstrap )
	( cd $(BUILDDIR) ; ln -f websocket-message bootstrap ; zip websocket-message.zip bootstrap ; rm bootstrap )
	( cd $(BUILDDIR) ; ln -f websocket-todo bootstrap ; zip websocket-todo.zip bootstrap ; rm bootstrap )
	( cd $(BUILDDIR) ; ln -f websocket-user bootstrap ; zip websocket-user.zip bootstrap ; rm bootstrap )
	( cd $(BUILDDIR) ; ln -f websocket-broadcast bootstrap ; zip websocket-broadcast.zip bootstrap ; rm bootstrap )
	( cd $(BUILDDIR) ; ln -f core-oauth-user bootstrap ; zip core-oauth-user.zip bootstrap ; rm bootstrap )
	( cd $(BUILDDIR) ; ln -f core-send-email bootstrap ; zip core-send-email.zip bootstrap ; rm bootstrap )
	( cd $(BUILDDIR) ; ln -f core-message bootstrap ; zip core-message.zip bootstrap ; rm bootstrap )
	( cd $(BUILDDIR) ; ln -f core-todo bootstrap ; zip core-todo.zip bootstrap ; rm bootstrap )
	( cd $(BUILDDIR) ; ln -f core-user bootstrap ; zip core-user.zip bootstrap ; rm bootstrap )
	( cd $(BUILDDIR) ; ln -f workers-user bootstrap ; zip workers-user.zip bootstrap ; rm bootstrap )
	( cd $(BUILDDIR) ; ln -f workers-file bootstrap ; zip workers-file.zip bootstrap ; rm bootstrap )
	( cd $(BUILDDIR) ; ln -f trigger-s3 bootstrap ; zip trigger-s3.zip bootstrap ; rm bootstrap )


run-lambda: build-lambda
	sam local start-api --template-file ./lambda/template.yaml --env-vars ../../dev-variables.json

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
