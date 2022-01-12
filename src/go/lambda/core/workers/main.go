package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/koblas/grpc-todo/services/core/workers"
)

func main() {
	handler := workers.HandleLambda
	lambda.Start(handler)
}
