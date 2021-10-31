import * as grpcWeb from 'grpc-web';

import * as publicapi_todo_pb from '../publicapi/todo_pb';


export class TodoServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  addTodo(
    request: publicapi_todo_pb.AddTodoParams,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: publicapi_todo_pb.TodoObject) => void
  ): grpcWeb.ClientReadableStream<publicapi_todo_pb.TodoObject>;

  deleteTodo(
    request: publicapi_todo_pb.DeleteTodoParams,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: publicapi_todo_pb.DeleteResponse) => void
  ): grpcWeb.ClientReadableStream<publicapi_todo_pb.DeleteResponse>;

  getTodos(
    request: publicapi_todo_pb.GetTodoParams,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: publicapi_todo_pb.TodoResponse) => void
  ): grpcWeb.ClientReadableStream<publicapi_todo_pb.TodoResponse>;

}

export class TodoServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  addTodo(
    request: publicapi_todo_pb.AddTodoParams,
    metadata?: grpcWeb.Metadata
  ): Promise<publicapi_todo_pb.TodoObject>;

  deleteTodo(
    request: publicapi_todo_pb.DeleteTodoParams,
    metadata?: grpcWeb.Metadata
  ): Promise<publicapi_todo_pb.DeleteResponse>;

  getTodos(
    request: publicapi_todo_pb.GetTodoParams,
    metadata?: grpcWeb.Metadata
  ): Promise<publicapi_todo_pb.TodoResponse>;

}

