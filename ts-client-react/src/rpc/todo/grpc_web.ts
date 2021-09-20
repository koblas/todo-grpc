import { TodoService, TodoItem } from "./index";
import { todoServiceClientImpl, GrpcWebImpl } from "../../models/todo";

const rpc = new GrpcWebImpl("http://localhost:8080", {
  debug: false,
});

export function newTodoClient(): TodoService {
  const client = new todoServiceClientImpl(rpc);

  return {
    async getTodos(): Promise<TodoItem[]> {
      return (await client.getTodos({})).todos;
    },
    async addTodo(task: string): Promise<TodoItem> {
      return client.addTodo({ task });
    },
    async deleteTodo(id: string): Promise<void> {
      client.deleteTodo({ id });
    },
  };
}
