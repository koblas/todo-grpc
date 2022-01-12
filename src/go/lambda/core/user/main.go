package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/koblas/grpc-todo/services/core/user"
)

func main() {
	handler := user.HandleLambda()
	lambda.Start(handler)
}
