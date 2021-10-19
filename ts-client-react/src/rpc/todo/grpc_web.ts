import { BrowserHeaders } from "browser-headers";
import { TodoService, TodoItem } from "./index";
import { todoServiceClientImpl, GrpcWebImpl } from "../../models/todo";
import { BASE_URL } from "../utils";

const rpc = new GrpcWebImpl(BASE_URL, {
  debug: false,
});

const client = new todoServiceClientImpl(rpc);

export function newTodoClient(token: string | null): TodoService {
  const metadata = new BrowserHeaders({
    Authorization: [`Bearer ${token}`],
  });

  return {
    async getTodos(): Promise<TodoItem[]> {
      return (await client.getTodos({}, metadata)).todos;
    },
    async addTodo(task: string): Promise<TodoItem> {
      return client.addTodo({ task }, metadata);
    },
    async deleteTodo(id: string): Promise<void> {
      client.deleteTodo({ id }, metadata);
    },
  };
}
