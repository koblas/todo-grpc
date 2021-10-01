import * as grpc from "@grpc/grpc-js";
import { todo } from "./grpcjs/protos/todo";
import { promisify } from "util";

// Based on
//   https://github.com/badsyntax/grpc-js-typescript/tree/master/examples/grpc_tools_node_protoc_ts

const host = "localhost:14586";

type UnaryCall<Params, Response> = (
  request: Params,
  callback: (error: grpc.ServiceError | null, response: Response) => void,
) => grpc.ClientUnaryCall;

export function gpromisify<T1, TResult, TClean = Exclude<TResult, undefined>>(
  fn: (arg1: T1, callback: (err: any, result: TResult) => void) => void,
): (arg1: T1) => Promise<TClean> {
  return (arg: T1): Promise<TClean> =>
    new Promise<TClean>((resolve, reject) => {
      fn(arg, (error: any, response) => {
        if (error || response === undefined) {
          return reject(error);
        }

        return resolve(response as unknown as TClean);
      });
    });
}

async function main(argv: string[]) {
  const task = argv.join(" ");
  const client = new todo.todoServiceClient(host, grpc.credentials.createInsecure());

  // We're class style, so we need the bind
  // const addTodo = gpromisify(client.addTodo.bind(client));

  // It's still returns a possibly undefined result, but it's better
  const addTodo = promisify(client.addTodo.bind(client));

  try {
    // const response = await new Promise<todo.todoObject>((resolve, reject) => {
    //   client.addTodo(new todo.addTodoParams({ task }), (error, response) => {
    //     if (error || response === undefined) {
    //       return reject(error);
    //     }
    //     return resolve(response);
    //   });
    // });
    // const response = await wrapper<todoObject>(client.addTodo, request);

    const response = await addTodo(new todo.addTodoParams({ task }));

    console.log("Success ", response?.id);
  } catch (error) {
    console.log(error);
  }
}

main(process.argv.slice(2))
  .then(() => {})
  .catch((err) => console.log(err));
