import { TodoService, TodoItem } from "./index";
import { newFetchClient } from "../utils";

export function newTodoClient(token: string | null): TodoService {
  const client = newFetchClient({ token });

  return {
    async getTodos(): Promise<TodoItem[]> {
      return (await client.GET<{ todos: TodoItem[] }>("/v1/todo/list")).todos;
    },
    async addTodo(task: string): Promise<TodoItem> {
      return await client.POST<TodoItem>(`/v1/todo/add`, { task });
    },
    async deleteTodo(id: string): Promise<void> {
      await client.DELETE(`/v1/todo/delete/${id}`);
    },
  };
}
