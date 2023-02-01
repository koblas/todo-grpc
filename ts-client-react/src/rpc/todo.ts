import { z } from "zod";

export const TodoItem = z.object({
  id: z.string(),
  task: z.string(),
});

export const TodoAdd = z.object({
  todo: TodoItem,
});

export const TodoList = z.object({
  todos: z.array(TodoItem),
});

export type TodoItemType = z.infer<typeof TodoItem>;
export type TodoAddType = z.infer<typeof TodoAdd>;
export type TodoListType = z.infer<typeof TodoList>;
