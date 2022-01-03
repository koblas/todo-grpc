package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/koblas/grpc-todo/services/publicapi/extauth"
)

func main() {
	handler := extauth.HandleLambda()
	lambda.Start(handler)
}
