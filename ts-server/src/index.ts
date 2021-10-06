import * as grpc from "@grpc/grpc-js";
import { todo } from "./grpcjs/protos/todo";
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
interface Todo extends todo.todoObject {}

let todos: Todo[] = [];

/**
 * Really handy function to convert callback style unary gRPC functions
 * into Promise based functions.
 */
function unaryWrap<I, O, CALL extends grpc.ServerUnaryCall<I, O>, CB extends grpc.requestCallback<O>>(
  handler: (params: I, metadata: grpc.Metadata) => Promise<O>,
) {
  type EType = Parameters<grpc.requestCallback<O>>[0];
  //  export declare type sendUnaryData<ResponseType> = (error: ServerErrorResponse | ServerStatusResponse | null, value?: ResponseType | null, trailer?: Metadata, flags?: number) => void;

  return (call: CALL, callback: CB): void => {
    handler(call.request, call.metadata)
      .then((r) => {
        callback(null, r);
      })
      .catch((error) => {
        if (error instanceof Error) {
          callback(error as EType);
        } else {
          callback(new Error(error) as EType);
        }
      });
  };
}

const TodoServer = {
  addTodo: unaryWrap(async ({ task }: todo.addTodoParams) => {
    const item = new todo.todoObject({
      id: uuidv4(),
      task: task,
    });

    todos.push(item);

    return item;
  }),

  deleteTodo: unaryWrap(async ({ id }: todo.deleteTodoParams) => {
    todos = todos.filter((todo) => todo.id !== id);

    return new todo.deleteResponse({ message: "Success" });
  }),

  getTodos: unaryWrap(async (input: todo.getTodoParams) => {
    return new todo.todoResponse({ todos });
  }),
};

function main(): void {
  const server = new grpc.Server();
  server.addService(todo.UnimplementedtodoServiceService.definition, TodoServer);
  server.bindAsync(`localhost:${process.env.PORT ?? 14586}`, grpc.ServerCredentials.createInsecure(), (err, port) => {
    if (err) {
      throw err;
    }
    console.log(`Listening on ${port}`);
    server.start();
  });
}

main();
