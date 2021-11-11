
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
