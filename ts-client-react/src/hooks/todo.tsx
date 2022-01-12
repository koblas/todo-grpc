import React, { useEffect, PropsWithChildren, useState } from "react";
import create from "zustand";
import { Draft, produce } from "immer";
import { zimmer } from "../util/zimmer";
import { newTodoClient } from "../rpc/todo/factory";
import { TodoItem, TodoList } from "../rpc/todo";
import { useAuth } from "./auth";
import { useNetworkContextErrors } from "./network";
import { RpcError, RpcMutation, RpcOptions } from "../rpc/errors";

interface TodoState {
  // readonly client: TodoService | null;
  readonly todos: TodoItem[];

  // Actions
  setTodos(todos: TodoItem[]): void;
  deleteTodo(id: TodoItem["id"]): void;
  updateTodo(id: TodoItem["id"], todo: TodoItem): void;
  appendTodo(todo: TodoItem): void;
  resetTodos(): void;
}

const useTodoStore = create<TodoState>(
  zimmer((set) => ({
    todos: [],
    resetTodos() {
      set(
        produce((draft: Draft<TodoState>) => {
          draft.todos = [];
        }),
      );
    },
    setTodos(todos: TodoItem[]) {
      set(
        produce((draft: Draft<TodoState>) => {
          draft.todos = todos;
        }),
      );
    },
    deleteTodo(id: TodoItem["id"]) {
      set(
        produce((draft: Draft<TodoState>) => {
          const index = draft.todos.findIndex((todo: TodoItem) => todo.id === id);

          if (index !== -1) {
            draft.todos.splice(index, 1);
          }
        }),
      );
    },
    updateTodo(id: TodoItem["id"], todo: TodoItem) {
      set(
        produce((draft: Draft<TodoState>) => {
          const index = draft.todos.findIndex((item: TodoItem) => item.id === id);

          if (index !== -1) {
            draft.todos[index] = todo;
          }
        }),
      );
    },
    appendTodo(todo: TodoItem) {
      set(
        produce((draft: Draft<TodoState>) => {
          draft.todos.push(todo);
        }),
      );
    },
  })),
);

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

const TodoContext = React.createContext({});

export function TodoContextProvider({ children }: PropsWithChildren<unknown>) {
  const { isAuthenticated } = useAuth();
  const { mutations } = useTodos();
  const [clearTodos] = mutations.useClearTodos();
  const [loadTodos] = mutations.useLoadTodos();

  useEffect(() => {
    if (isAuthenticated) {
      loadTodos({});
    } else {
      clearTodos();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isAuthenticated]);

  return <TodoContext.Provider value={{}}>{children}</TodoContext.Provider>;
}
