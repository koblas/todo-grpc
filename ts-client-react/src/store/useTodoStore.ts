import create from "zustand";
import { Draft, produce } from "immer";
import { zimmer } from "../util/zimmer";

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

export const getTodoStore = create(
  zimmer<TodoState>((set) => ({
    // create<TodoState>(
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
          if (!draft.todos.some((v) => v.id === todo.id)) {
            draft.todos.push(todo);
          }
        }),
      );
    },
  })),
);
