import { Reducer, useEffect, createContext, useContext, PropsWithChildren, useReducer } from "react";
import { must, assert } from "../util/assert";
import { newTodoClient } from "../rpc/todo/factory";
import { TodoItem, TodoService } from "../rpc/todo";

const todoClient = newTodoClient("grpc");

const TodoContext = createContext<{
  client: TodoService;
  todos: TodoItem[];
  refresh: React.Dispatch<React.SetStateAction<void>>;
  dispatch: React.Dispatch<DispatchAction<TodoItem>>;
}>({
  client: todoClient,
  todos: [],
  dispatch() {
    throw new Error("not initialized");
  },
  refresh() {
    throw new Error("not initialized");
  },
});

interface ListBase {
  id: string;
}

type DispatchAction<T> = {
  type: "set" | "update" | "delete" | "append";
  id?: string;
  value?: T;
  list?: T[];
};

function listReducer<T extends ListBase>(state: T[], action: DispatchAction<T>): T[] {
  switch (action.type) {
    case "set":
      assert(action.list);
      return action.list;
    case "update":
      assert(action.id, "ID missing");
      assert(action.value, "Value missing");
      return state.map((todo) => (todo.id === action.id ? must(action.value) : todo));
    case "delete":
      assert(action.id, "ID missing");
      return state.filter((todo) => todo.id !== action.id);
    case "append":
      assert(action.value, "Value missing");
      return [...state, action.value];
  }
}

function useListReducer<T extends ListBase>(initial: T[] = []) {
  return useReducer<Reducer<T[], DispatchAction<T>>>(listReducer, initial);
}

export function TodoContextProvider({ children }: PropsWithChildren<unknown>) {
  // The reason we need to use a reducer here rather that a setState is
  // that when you are doing optimistic updates your not able to get a handle
  // on the "current" set of items in your list.  You might have done 2..3 actions
  // but your local closure will only have the state at the time you dispatched
  // your asyncronist event.  Thus you need to bump the world into a reducer...
  const [state, dispatch] = useListReducer<TodoItem>();

  function refresh() {
    todoClient.getTodos().then((todos) => {
      dispatch({ type: "set", list: todos });
    });
  }

  useEffect(() => {
    refresh();
  }, [todoClient]);

  return (
    <TodoContext.Provider value={{ todos: state, refresh, dispatch, client: todoClient }}>
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
        .addTodo(task)
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
      dispatch({ type: "delete", id });
      client
        .deleteTodo(id)
        .then(() => {
          refresh();
        })
        .catch((err) => {
          console.log(err);
          dispatch({ type: "append", value: obj });
        });
    },
  };
}
