import { Json } from "../../types/json";
import { RpcOptions } from "../errors";

export interface TodoItem {
  id: string;
  task: string;
}

export interface TodoList {
  todos: TodoItem[];
}

export interface TodoService {
  getTodos(params: Json, options: RpcOptions<TodoList>): void;
  addTodo(params: Pick<TodoItem, "task">, options: RpcOptions<TodoItem>): void;
  deleteTodo(params: Pick<TodoItem, "id">, options: RpcOptions<unknown>): void;
  // getTodos(options: RpcOptions<TodoItem[]>): void;
  // addTodo(task: string, options: RpcOptions<TodoItem>): void;
  // deleteTodo(id: string, options: RpcOptions<void>): void;
}
