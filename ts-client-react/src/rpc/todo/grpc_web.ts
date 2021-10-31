import * as grpcWeb from "grpc-web";
import { TodoService } from "./index";
import { TodoServiceClient } from "../../genpb/publicapi/todo_grpc_web_pb";
import {
  AddTodoParams,
  DeleteResponse,
  DeleteTodoParams,
  GetTodoParams,
  TodoObject,
  TodoResponse,
} from "../../genpb/publicapi/todo_pb";
import { BASE_URL } from "../utils";
import { handleGrpcError } from "../utils/grpc_helpers";

export function newTodoClient(token: string | null): TodoService {
  const client = new TodoServiceClient(BASE_URL);
  const metadata = {
    Authorization: `Bearer ${token}`,
  };

  return {
    getTodos(options) {
      const req = new GetTodoParams();

      client.getTodos(req, metadata, (err: grpcWeb.RpcError, data: TodoResponse) => {
        if (err) {
          handleGrpcError(err, options);
          return;
        }
        const todos = data.getTodosList().map((item) => item.toObject());
        options.onCompleted?.(todos);
      });
    },
    addTodo(task, options) {
      const req = new AddTodoParams();
      req.setTask(task);

      client.addTodo(req, metadata, (err: grpcWeb.RpcError, data: TodoObject) => {
        if (err) {
          handleGrpcError(err, options);
          return;
        }
        options.onCompleted?.(data.toObject());
      });
    },
    deleteTodo(id, options) {
      const req = new DeleteTodoParams();
      req.setId(id);

      client.deleteTodo(req, metadata, (err: grpcWeb.RpcError, data: DeleteResponse) => {
        if (err) {
          handleGrpcError(err, options);
          return;
        }
        options.onCompleted?.();
      });
    },
  };
}
