/* eslint-disable @typescript-eslint/no-unused-vars */
import { RpcOptions } from "../../errors";
import { TodoService, TodoItem } from "../index";

export function newTodoClient(): TodoService {
  let todos: TodoItem[] = [];

  return {
    async getTodos(): Promise<TodoItem[]> {
      return [...todos];
    },
    addTodo(params: Pick<TodoItem, "task">, options: RpcOptions<TodoItem>): void {
      const item = { task: params.task, id: new Date().toISOString() };
      todos = [...todos, item];
    },
    deleteTodo(params: Pick<TodoItem, "id">, options: RpcOptions<unknown>): void {
      todos = todos.filter((todo) => todo.id !== params.id);
    },
  };
}
