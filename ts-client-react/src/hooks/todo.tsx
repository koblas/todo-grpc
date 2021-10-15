import { useEffect, createContext, useContext, PropsWithChildren, useReducer } from "react";
import { Draft } from "immer";
import { useImmerReducer } from "use-immer";
import { must, assert } from "../util/assert";
import { newTodoClient } from "../rpc/todo/factory";
import { TodoItem, TodoService } from "../rpc/todo";
import { useAuth } from "./auth";

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

type DispatchAction =
  | {
      type: "set";
      list: TodoItem[];
    }
  | {
      type: "setClient";
      client: TodoService;
    }
  | {
      type: "delete";
      id: string;
    }
  | {
      type: "append";
      value: TodoItem;
    }
  | {
      type: "update";
      id: string;
      value: TodoItem;
    };

function listReducer(draft: Draft<TodoState>, action: DispatchAction) {
  switch (action.type) {
    case "setClient":
      assert(action.client);
      draft.client = action.client;
      break;
    case "set":
      assert(action.list);
      draft.todos = action.list;
      break;
    case "update":
      assert(action.id, "ID missing");
      assert(action.value, "Value missing");
      const itemIdx = draft.todos.findIndex((todo) => todo.id === action.id);
      draft.todos[itemIdx] = action.value;
      break;
    case "delete":
      assert(action.id, "ID missing");
      draft.todos = draft.todos.filter((todo) => todo.id !== action.id);
      break;
    case "append":
      assert(action.value, "Value missing");
      draft.todos.push(action.value);
      break;
  }
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
  const { token, isAuthenticated } = useAuth();

  useEffect(() => {
    const todoClient = newTodoClient(token, "json");

    dispatch({ type: "setClient", client: todoClient });
  }, [token]);

  function refresh() {
    if (state.client && isAuthenticated) {
      state.client.getTodos().then((todos) => {
        dispatch({ type: "set", list: todos });
      });
    } else {
      dispatch({ type: "set", list: [] });
    }
  }

  useEffect(() => {
    refresh();
  }, [state.client]);

  return (
    <TodoContext.Provider value={{ todos: state.todos, refresh, dispatch, client: state.client }}>
      {children}
    </TodoContext.Provider>
  );
}

export function useTodos() {
  const { client, refresh, todos, dispatch } = useContext(TodoContext);

  return {
    todos,
    addTodo(task: TodoItem["task"]) {
      if (task === "") {
        return;
      }
      const id = new Date().toISOString();
      // optimistic creation
      dispatch({
        type: "append",
        value: {
          task,
          id,
        },
      });
      client
        ?.addTodo(task)
        .then((obj) => {
          dispatch({ type: "update", id, value: obj });
          refresh();
        })
        .catch((err) => {
          console.log(err);
          // TODO -- Display error
          dispatch({ type: "delete", id });
        });
    },
    deleteTodo(id: TodoItem["id"]) {
      const obj = todos.find((todo) => todo.id === id);
      if (obj) {
        dispatch({ type: "delete", id });
        client
          ?.deleteTodo(id)
          .then(() => {
            refresh();
          })
          .catch((err) => {
            console.log(err);
            dispatch({ type: "append", value: obj });
          });
      }
    },
  };
}
