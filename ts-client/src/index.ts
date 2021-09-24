import * as grpc from "@grpc/grpc-js";
import { addTodoParams, todoObject } from "./grpcjs/protos/todo_pb";
import { todoServiceClient } from "./grpcjs/protos/todo_grpc_pb";

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
  const client = new todoServiceClient(host, grpc.credentials.createInsecure());

  const request = new addTodoParams();
  request.setTask(task);

  try {
    const response = await new Promise<todoObject>((resolve, reject) => {
      client.addTodo(request, (error, response) => {
        if (error) {
          return reject(error);
        }
        return resolve(response);
      });
    });
    // const response = await wrapper<todoObject>(client.addTodo, request);

    console.log("Success ", response.getId());
  } catch (error) {
    console.log(error);
  }
}

main(process.argv.slice(2))
  .then(() => {})
  .catch((err) => console.log(err));
