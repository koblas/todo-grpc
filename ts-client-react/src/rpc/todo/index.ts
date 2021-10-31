import { RpcOptions } from "../errors";

export interface TodoItem {
  id: string;
  task: string;
}

export interface TodoService {
  getTodos(options: RpcOptions<TodoItem[]>): void;
  addTodo(task: string, options: RpcOptions<TodoItem>): void;
  deleteTodo(id: string, options: RpcOptions<void>): void;
}
