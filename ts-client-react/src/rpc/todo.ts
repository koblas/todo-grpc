import { z } from "zod";

export const TodoItem = z.object({
  id: z.string(),
  task: z.string(),
});

export const TodoList = z.object({
  todos: z.array(TodoItem),
});

export type TodoItemType = z.infer<typeof TodoItem>;
export type TodoListType = z.infer<typeof TodoList>;
