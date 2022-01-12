package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/koblas/grpc-todo/services/core/todo"
)

func main() {
	handler := todo.HandleLambda()
	lambda.Start(handler)
}
