import { useState } from "react";
import { newTodoClient } from "../rpc/todo/factory";
import { TodoItem, TodoList } from "../rpc/todo";
import { useAuth } from "./auth";
import { useNetworkContextErrors } from "./network";
import { RpcError, RpcMutation, RpcOptions } from "../rpc/errors";
import { useTodoStore } from "../store/useTodoStore";

export function useTodos() {
  const { token } = useAuth();
  const client = newTodoClient(token, "json");
  const addHandler = useNetworkContextErrors();
  const state = useTodoStore((s) => s);

  return {
    todos: state.todos,
    mutations: {
      useLoadTodos(): RpcMutation<unknown, TodoList> {
        const [data, setData] = useState<TodoList | undefined>();
        const [loading, setLoading] = useState(false);
        const [error, setError] = useState<RpcError | undefined>(undefined);

        const func = (params: unknown, options?: RpcOptions<TodoList>) => {
          setLoading(true);
          client?.getTodos(
            {},
            addHandler(
              {
                onCompleted(input) {
                  setData(input);
                  state.setTodos(input.todos);
                },
                onError(err: RpcError) {
                  setError(err);
                },
                onFinished() {
                  setLoading(false);
                },
              },
              options,
            ),
          );
        };

        return [func, { data, loading, error }];
      },

      useAddTodo(): RpcMutation<Pick<TodoItem, "task">, TodoItem> {
        const [data, setData] = useState<TodoItem | undefined>();
        const [loading, setLoading] = useState(false);
        const [error, setError] = useState<RpcError | undefined>(undefined);

        const func = (params: Pick<TodoItem, "task">, options?: RpcOptions<TodoItem>) => {
          if (params.task === "") {
            return;
          }
          const id = new Date().toISOString();

          state.appendTodo({ id, task: params.task });

          setLoading(true);
          client?.addTodo(
            params,
            addHandler(
              {
                onCompleted(input) {
                  state.updateTodo(id, input);
                  setData(input);
                },
                onError(err: RpcError) {
                  state.deleteTodo(id);
                  setError(err);
                },
                onFinished() {
                  setLoading(false);
                },
              },
              options,
            ),
          );
        };

        return [func, { data, loading, error }];
      },

      useDeleteTodo(): RpcMutation<Pick<TodoItem, "id">, void> {
        const [loading, setLoading] = useState(false);
        const [error, setError] = useState<RpcError | undefined>(undefined);

        const func = (params: Pick<TodoItem, "id">, options?: RpcOptions<void>) => {
          const obj = state.todos.find((todo) => todo.id === params.id);

          if (!obj) {
            return;
          }

          state.deleteTodo(params.id);

          setLoading(true);
          client?.deleteTodo(
            params,
            addHandler(
              {
                onCompleted() {
                  // nothing
                },
                onError(err: RpcError) {
                  if (obj) {
                    state.appendTodo(obj);
                  }
                  setError(err);
                },
                onFinished() {
                  setLoading(false);
                },
              },
              options,
            ),
          );
        };

        return [func, { data: undefined, loading, error }];
      },

      useClearTodos() {
        const func = () => {
          state.resetTodos();
        };
        return [func];
      },
    },
  };
}
