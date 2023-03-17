import { z } from "zod";
import { RpcOptions } from "./errors";

export const TodoObject = z.object({
  id: z.string(),
  task: z.string(),
});
export type TodoObjectT = z.infer<typeof TodoObject>;

// todo_add
export const TodoAddRequest = z.object({
  task: z.string(),
});
export const TodoAddResponse = z.object({
  todo: TodoObject,
});
export type TodoAddRequestT = z.infer<typeof TodoAddRequest>;
export type TodoAddResponseT = z.infer<typeof TodoAddResponse>;

// todo_delete
export const TodoDeleteRequest = z.object({
  id: z.string(),
});
export const TodoDeleteResponse = z.object({
  message: z.string(),
});
export type TodoDeleteRequestT = z.infer<typeof TodoDeleteRequest>;
export type TodoDeleteResponseT = z.infer<typeof TodoDeleteResponse>;

// todo_add
export const TodoListRequest = z.object({});
export const TodoListResponse = z.object({
  todos: z.array(TodoObject),
});
export type TodoListRequestT = z.infer<typeof TodoListRequest>;
export type TodoListResponseT = z.infer<typeof TodoListResponse>;

export interface TodoService {
  todo_add(params: TodoAddRequestT, options: RpcOptions<TodoAddResponseT>): void;
  todo_delete(params: TodoListRequestT, options: RpcOptions<TodoListResponseT>): void;
  todo_list(params: TodoDeleteRequestT, options: RpcOptions<TodoDeleteResponseT>): void;
}
