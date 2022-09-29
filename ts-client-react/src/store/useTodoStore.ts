import create from "zustand";

import { TodoItem } from "../rpc/todo";

export interface TodoState {
  // readonly client: TodoService | null;
  readonly todos: TodoItem[];

  // Actions
  setTodos(todos: TodoItem[]): void;
  deleteTodo(id: TodoItem["id"]): void;
  updateTodo(id: TodoItem["id"], todo: TodoItem): void;
  appendTodo(todo: TodoItem): void;
  resetTodos(): void;
}

export const getTodoStore = create<TodoState>((set) => ({
  todos: [],

  setTodos(todos: TodoItem[]): void {
    return set((state) => ({ ...state, todos: [...todos] }));
  },
  deleteTodo(id: TodoItem["id"]): void {
    return set((state) => ({ ...state, todos: state.todos.filter((v) => v.id !== id) }));
  },
  updateTodo(id: TodoItem["id"], todo: TodoItem): void {
    return set((state) => {
      const index = state.todos.findIndex((item: TodoItem) => item.id === id);
      if (index < 0) {
        return state;
      }

      const front = state.todos.slice(0, index);
      const end = state.todos.slice(index + 1);

      return {
        ...state,
        todos: [...front, todo, ...end],
      };
    });
  },
  appendTodo(todo: TodoItem): void {
    set((state) => {
      if (!state.todos.some((v) => v.id === todo.id)) {
        return { ...state, todos: [...state.todos, todo] };
      }
      return state;
    });
  },
  resetTodos(): void {
    set((state) => ({ ...state, todos: [] }));
  },
}));
