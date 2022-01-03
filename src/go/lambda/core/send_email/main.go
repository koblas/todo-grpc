package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/koblas/grpc-todo/services/core/send_email"
)

func main() {
	handler := send_email.HandleLambda()
	lambda.Start(handler)
}
