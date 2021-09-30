import * as grpc from "@grpc/grpc-js";
import { todo } from "./grpcjs/protos/todo";

// Based on
//   https://github.com/badsyntax/grpc-js-typescript/tree/master/examples/grpc_tools_node_protoc_ts

const host = "localhost:14586";

type UnaryCall<Params, Response> = (
  request: Params,
  callback: (error: grpc.ServiceError | null, response: Response) => void,
) => grpc.ClientUnaryCall;

// async function wrapper<Result, Params = any, Caller extends UnaryCall<Params, Result> = any>(
//   call: Caller,
//   request: Params,
// ): Promise<Result> {
//   return await new Promise<Result>((resolve, reject) => {
//     call(request, (error: grpc.ServiceError | null, response: Result) => {
//       if (error) {
//         return reject(error);
//       }
//       return resolve(response);
//     });
//   });
// }

async function main(argv: string[]) {
  const task = argv.join(" ");
  const client = new todo.todoServiceClient(host, grpc.credentials.createInsecure());

  try {
    const response = await new Promise<todo.todoObject>((resolve, reject) => {
      client.addTodo(new todo.addTodoParams({ task }), (error, response) => {
        if (error || response === undefined) {
          return reject(error);
        }
        return resolve(response);
      });
    });
    // const response = await wrapper<todoObject>(client.addTodo, request);

    console.log("Success ", response.id);
  } catch (error) {
    console.log(error);
  }
}

main(process.argv.slice(2))
  .then(() => {})
  .catch((err) => console.log(err));
