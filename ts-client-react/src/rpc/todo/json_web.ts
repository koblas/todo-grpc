import { TodoService, TodoItem, TodoList } from "./index";
import { newFetchPOST } from "../utils";
import { RpcOptions } from "../errors";

export function newTodoClient(token: string | null): TodoService {
  const client = newFetchPOST({ token });
  const base = "/v1/todo";

  return {
    getTodos(params, options: RpcOptions<TodoList>): void {
      client.call(`${base}/get_todos`, params, options);
    },
    addTodo(params, options: RpcOptions<TodoItem>): void {
      client.call(`${base}/add_todo`, params, options);
    },
    deleteTodo(params, options: RpcOptions<unknown>): void {
      client.call(`${base}/delete_todo`, params, options);
    },
    // getTodos(options) {
    //   client
    //     .POST<{ todos: TodoItem[] }>(`${base}/get_todos`, {})
    //     .then((data) => {
    //       options.onCompleted?.(data.todos);
    //     })
    //     .catch((err) => handleJsonError(err, options));
    // },
    // async addTodo(task, options) {
    //   client
    //     .POST<TodoItem>(`/v1/todo/add_todo`, { task })
    //     .then((todo) => {
    //       options.onCompleted?.(todo);
    //     })
    //     .catch((err) => handleJsonError(err, options));
    // },
    // async deleteTodo(id, options) {
    //   client
    //     .POST(`${base}/delete_todo`, { id })
    //     .then(() => {
    //       options.onCompleted?.();
    //     })
    //     .catch((err) => handleJsonError(err, options));
    // },
  };
}
