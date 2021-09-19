import * as grpc from "@grpc/grpc-js";
import {
  addTodoParams,
  deleteTodoParams,
  getTodoParams,
  todoObject,
  todoServiceServer,
  todoServiceService,
} from "./grpcjs/todo";
import { v4 as uuidv4 } from "uuid";

// To build gRPC definitions
//
// protoc --proto_path=$(PROTO_DIR) \
//   --ts_proto_opt=outputServices=grpc-js \
//   --ts_proto_opt=esModuleInterop=true \
//   --plugin=$(TS_SERVER)/node_modules/.bin/protoc-gen-ts_proto \
//   --ts_proto_out=import_style=commonjs,binary:$(TS_SERVER)/src/grpcjs/ \
//   protos/todo.proto
//

// Our internal Todo fufills the interface of the gRPC data
interface Todo extends todoObject {}

let todos: Todo[] = [];

/**
 * Really handy function to convert callback style unary gRPC functions
 * into Promise based functions.
 */
function unaryWrap<I, O, CALL extends grpc.ServerUnaryCall<I, O>, CB extends grpc.sendUnaryData<O>>(
  handler: (params: CALL["request"], metadata: CALL["metadata"]) => Promise<O>,
) {
  return (call: CALL, callback: CB): void => {
    handler(call.request, call.metadata)
      .then((r) => {
        callback(null, r);
      })
      .catch((error) => {
        if (error instanceof Error) {
          callback(error, null);
        } else {
          callback(new Error(JSON.stringify(error)), null);
        }
      });
  };
}

const TodoServer: todoServiceServer = {
  addTodo: unaryWrap(async ({ task }: addTodoParams) => {
    const item = {
      id: uuidv4(),
      task: task,
    };

    todos.push(item);

    return item;
  }),
  deleteTodo: unaryWrap(async ({ id }: deleteTodoParams) => {
    todos = todos.filter((todo) => todo.id !== id);

    return { message: "Success" };
  }),
  getTodos: unaryWrap(async (input: getTodoParams) => {
    return { todos };
  }),
};

function main(): void {
  const server = new grpc.Server();
  server.addService(todoServiceService, TodoServer);
  server.bindAsync(`localhost:${process.env.PORT ?? 14586}`, grpc.ServerCredentials.createInsecure(), (err, port) => {
    if (err) {
      throw err;
    }
    console.log(`Listening on ${port}`);
    server.start();
  });
}

main();
