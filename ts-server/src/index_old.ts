import * as grpc from "@grpc/grpc-js";
import * as jspb from "google-protobuf";
import {
  addTodoParams,
  deleteResponse,
  deleteTodoParams,
  getTodoParams,
  todoObject,
  todoResponse,
} from "./grpcjs/todo_pb";
import { todoServiceService, ItodoServiceServer } from "./grpcjs/todo_grpc_pb";
import { v4 as uuidv4 } from "uuid";

//
//  To build the correct gRPC definitions
//
//	$(TS_SERVER)/node_modules/.bin/grpc_tools_node_protoc \
//		--js_out=import_style=commonjs,binary:$(TS_SERVER)/src/grpcjs \
//		--grpc_out=grpc_js:$(TS_SERVER)/src/grpcjs \
//		-I ./protos \
//		protos/todo.proto
//	$(TS_SERVER)/node_modules/.bin/grpc_tools_node_protoc \
//		--plugin=protoc-gen-ts=$(TS_SERVER)/node_modules/.bin/protoc-gen-ts \
//		--ts_out=grpc_js:$(TS_SERVER)/src/grpcjs \
//		-I ./protos \
//		protos/todo.proto
//

let todos: todoObject[] = [];

interface asObjectExtractor<T> {
  toObject(includeInstance?: boolean): T;
}

function proxyObj<T extends jspb.Message & asObjectExtractor<A>, A = unknown>(input: T): T & ReturnType<T["toObject"]> {
  return new Proxy<T>(input, {
    set(obj: T, prop: string | symbol, value: any) {
      if (typeof prop === "string") {
        const name = `set${prop[0].toUpperCase()}${prop.substr(1)}`;
        const fn = (obj as any)[name];
        if (typeof fn === "function") {
          fn.call(obj, value);

          return true;
        }
      }

      (obj as any)[prop] = value;

      return true;
    },

    get(obj: T, prop: string | symbol) {
      if (typeof prop === "string") {
        const name = `get${prop[0].toUpperCase()}${prop.substr(1)}`;
        const fn = (obj as any)[name];
        if (typeof fn === "function") {
          return fn.call(obj);
        }
      }

      return (obj as any)[prop];
    },
  }) as T & ReturnType<T["toObject"]>;
}

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

const TodoServer: ItodoServiceServer = {
  addTodo: unaryWrap(async (request: addTodoParams) => {
    const input = request.toObject();
    const result = proxyObj(new todoObject());

    result.id = uuidv4();
    result.task = input.task;

    todos.push(result);

    return result;
  }),
  //   addTodo(call: grpc.ServerUnaryCall<addTodoParams, todoObject>, callback: grpc.sendUnaryData<todoObject>): void {
  //     const item = new todoObject();

  //     item.setId(uuidv4()).setTask(call.request.getTask());
  //     todos.push(item);

  //     callback(null, item);
  //   },

  deleteTodo(
    call: grpc.ServerUnaryCall<deleteTodoParams, deleteResponse>,
    callback: grpc.sendUnaryData<deleteResponse>,
  ): void {
    const id = call.request.getId();

    todos = todos.filter((todo) => todo.getId() !== id);

    const response = new deleteResponse();
    response.setMessage("Success");

    callback(null, response);
  },

  getTodos(call: grpc.ServerUnaryCall<getTodoParams, todoResponse>, callback: grpc.sendUnaryData<todoResponse>): void {
    const response = new todoResponse();

    response.setTodosList(todos);

    callback(null, response);
  },
};

function main(): void {
  const server = new grpc.Server();
  // @ts-ignore
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
