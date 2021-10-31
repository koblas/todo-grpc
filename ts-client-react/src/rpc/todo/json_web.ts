import { TodoService, TodoItem } from "./index";
import { newFetchClient } from "../utils";
import { RpcOptions } from "../errors";
import { handleJsonError } from "../utils/json_helpers";

export function newTodoClient(token: string | null): TodoService {
  const client = newFetchClient({ token });

  return {
    getTodos(options) {
      client
        .GET<{ todos: TodoItem[] }>("/v1/todo/list")
        .then((data) => {
          options.onCompleted?.(data.todos);
        })
        .catch((err) => handleJsonError(err, options));
    },
    async addTodo(task, options) {
      client
        .POST<TodoItem>(`/v1/todo/add`, { task })
        .then((todo) => {
          options.onCompleted?.(todo);
        })
        .catch((err) => handleJsonError(err, options));
    },
    async deleteTodo(id, options) {
      client
        .DELETE(`/v1/todo/delete/${id}`)
        .then(() => {
          options.onCompleted?.();
        })
        .catch((err) => handleJsonError(err, options));
    },
  };
}
