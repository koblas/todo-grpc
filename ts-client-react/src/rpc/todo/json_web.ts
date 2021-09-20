import { TodoService, TodoItem } from "./index";

const BASE = "http://localhost:8080";

export function newTodoClient(): TodoService {
  return {
    async getTodos(): Promise<TodoItem[]> {
      const response = await fetch(`${BASE}/v1/todo/list`, {});

      return (await response.json()).todos;
    },
    async addTodo(task: string): Promise<TodoItem> {
      const response = await fetch(`${BASE}/v1/todo/add`, {
        method: "POST",
        body: JSON.stringify({ task }),
      });

      return response.json();
    },
    async deleteTodo(id: string): Promise<void> {
      await fetch(`${BASE}/v1/todo/delete/${id}`, {
        method: "DELETE",
      });
    },
  };
}
