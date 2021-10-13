package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"

	todo "github.com/koblas/grpc-todo/genpb"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(":14586", opts...)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	client := todo.NewTodoServiceClient(conn)

	ctx := context.Background()
	params := todo.AddTodoParams{Task: os.Args[1]}
	response, err := client.AddTodo(ctx, &params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response.Id)
}
