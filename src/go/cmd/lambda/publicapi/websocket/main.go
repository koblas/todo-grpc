package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/koblas/grpc-todo/services/publicapi/websocket"
)

func main() {
	handler := websocket.HandleLambda()
	lambda.Start(handler)
}
