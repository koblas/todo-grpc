import { BrowserHeaders } from "browser-headers";
import { TodoService } from "./index";
import { todoServiceClientImpl, GrpcWebImpl } from "../../models/todo";
import { BASE_URL } from "../utils";
import { ErrorUnauthenticated } from "../errors";

const rpc = new GrpcWebImpl(BASE_URL, {
  debug: false,
});

const client = new todoServiceClientImpl(rpc);

function wrapErrors<T, A extends unknown[]>(func: (...args: A) => Promise<T>) {
  return async (...args: A): Promise<T> => {
    try {
      return await func(...args);
    } catch (err) {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      if (typeof err === "object" && err !== null && (err as any).code === 16) {
        throw new ErrorUnauthenticated();
      } else {
        throw err;
      }
    }
  };
}

export function newTodoClient(token: string | null): TodoService {
  const metadata = new BrowserHeaders({
    Authorization: [`Bearer ${token}`],
  });

  return {
    getTodos: wrapErrors(async () => (await client.getTodos({}, metadata)).todos),
    addTodo: wrapErrors(async (task: string) => await client.addTodo({ task }, metadata)),
    deleteTodo: wrapErrors(async (id: string) => {
      await client.deleteTodo({ id }, metadata);
    }),
  };
}
