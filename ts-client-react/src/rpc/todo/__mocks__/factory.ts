import { TodoService, TodoItem } from "../index";

export function newTodoClient(): TodoService {
  let todos: TodoItem[] = [];

  return {
    async getTodos(): Promise<TodoItem[]> {
      return [...todos];
    },
    async addTodo(task: string): Promise<TodoItem> {
      const item = { task, id: new Date().toISOString() };
      todos = [...todos, item];
      return item;
    },
    async deleteTodo(id: string): Promise<void> {
      todos = todos.filter((todo) => todo.id !== id);
    },
  };
}
