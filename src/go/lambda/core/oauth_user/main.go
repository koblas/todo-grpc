package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	ou "github.com/koblas/grpc-todo/services/core/oauth_user"
)

func main() {
	handler := ou.HandleLambda()
	lambda.Start(handler)
}
