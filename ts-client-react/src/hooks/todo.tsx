import React, { useEffect, createContext, useContext, PropsWithChildren } from "react";
import { Draft } from "immer";
import { useImmerReducer } from "use-immer";
import { assert } from "../util/assert";
import { newTodoClient } from "../rpc/todo/factory";
import { TodoItem, TodoService } from "../rpc/todo";
import { useAuth } from "./auth";
import { useNetworkContextErrors } from "./network";

const TodoContext = createContext<{
  client: TodoService | null;
  todos: TodoItem[];
  refresh: React.Dispatch<React.SetStateAction<void>>;
  dispatch: React.Dispatch<DispatchAction>;
}>({
  client: null,
  todos: [],
  dispatch() {
    throw new Error("not initialized");
  },
  refresh() {
    throw new Error("not initialized");
  },
});

interface TodoState {
  readonly client: TodoService | null;
  readonly todos: TodoItem[];
}

enum ActionType {
  SET,
  SET_CLIENT,
  DELETE,
  APPEND,
  UPDATE,
}

type DispatchAction =
  | {
      type: ActionType.SET;
      list: TodoItem[];
    }
  | {
      type: ActionType.SET_CLIENT;
      client: TodoService;
    }
  | {
      type: ActionType.DELETE;
      id: string;
    }
  | {
      type: ActionType.APPEND;
      value: TodoItem;
    }
  | {
      type: ActionType.UPDATE;
      id: string;
      value: TodoItem;
    };

function listReducer(draft: Draft<TodoState>, action: DispatchAction) {
  const handlers: Record<ActionType, () => void> = {
    [ActionType.SET_CLIENT]() {
      assert(action.type === ActionType.SET_CLIENT);
      assert(action.client);
      draft.client = action.client;
    },
    [ActionType.SET]() {
      assert(action.type === ActionType.SET);
      assert(action.list);
      draft.todos = action.list;
    },
    [ActionType.UPDATE]() {
      assert(action.type === ActionType.UPDATE);
      assert(action.id, "ID missing");
      assert(action.value, "Value missing");
      const itemIdx = draft.todos.findIndex((todo) => todo.id === action.id);
      draft.todos[itemIdx] = action.value;
    },
    [ActionType.DELETE]() {
      assert(action.type === ActionType.DELETE);
      assert(action.id, "ID missing");
      draft.todos = draft.todos.filter((todo) => todo.id !== action.id);
    },
    [ActionType.APPEND]() {
      assert(action.type === ActionType.APPEND);
      assert(action.value, "Value missing");
      draft.todos.push(action.value);
    },
  };

  handlers[action.type]();
}

const initialState: TodoState = {
  client: null,
  todos: [],
};

export function TodoContextProvider({ children }: PropsWithChildren<unknown>) {
  // The reason we need to use a reducer here rather that a setState is
  // that when you are doing optimistic updates your not able to get a handle
  // on the "current" set of items in your list.  You might have done 2..3 actions
  // but your local closure will only have the state at the time you dispatched
  // your asyncronist event.  Thus you need to bump the world into a reducer...
  const [state, dispatch] = useImmerReducer(listReducer, initialState);
  const { token, isAuthenticated, mutations } = useAuth();
  const addHandlers = useNetworkContextErrors();
  const logout = mutations.useLogout();

  useEffect(() => {
    const todoClient = newTodoClient(token, "grpc");

    dispatch({ type: ActionType.SET_CLIENT, client: todoClient });
  }, [token, dispatch]);

  function refresh() {
    if (state.client && isAuthenticated) {
      state.client.getTodos(
        addHandlers({
          onCompleted(todos) {
            dispatch({ type: ActionType.SET, list: todos });
          },
          onErrorAuthentication() {
            logout();
          },
        }),
      );
    } else {
      dispatch({ type: ActionType.SET, list: [] });
    }
  }

  /* eslint-disable react-hooks/exhaustive-deps */
  useEffect(() => {
    refresh();
  }, [state.client]);
  /* eslint-enable react-hooks/exhaustive-deps */

  return (
    <TodoContext.Provider value={{ todos: state.todos, refresh, dispatch, client: state.client }}>
      {children}
    </TodoContext.Provider>
  );
}

export function useTodos() {
  const { client, refresh, todos, dispatch } = useContext(TodoContext);
  const addHandlers = useNetworkContextErrors();

  return {
    todos,
    addTodo(task: TodoItem["task"]) {
      if (task === "") {
        return;
      }
      const id = new Date().toISOString();
      // optimistic creation
      dispatch({
        type: ActionType.APPEND,
        value: {
          task,
          id,
        },
      });
      client?.addTodo(
        task,
        addHandlers({
          onCompleted(data) {
            dispatch({ type: ActionType.UPDATE, id, value: data });
            refresh();
          },
          onError() {
            dispatch({ type: ActionType.DELETE, id });
          },
        }),
      );
    },
    deleteTodo(id: TodoItem["id"]) {
      const obj = todos.find((todo) => todo.id === id);
      if (obj) {
        dispatch({ type: ActionType.DELETE, id });
        client?.deleteTodo(
          id,
          addHandlers({
            onCompleted() {
              refresh();
            },
            onError() {
              dispatch({ type: ActionType.APPEND, value: obj });
            },
          }),
        );
      }
    },
  };
}
