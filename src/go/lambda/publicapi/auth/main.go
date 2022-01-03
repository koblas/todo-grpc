package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/koblas/grpc-todo/services/publicapi/auth"
)

func main() {
	handler := auth.HandleLambda()
	lambda.Start(handler)
}
