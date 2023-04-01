import { z } from "zod";
import { RpcOptions } from "./errors";

export const MessageItem = z.object({
  id: z.string(),
  roomId: z.string(),
  sender: z.string(),
  text: z.string(),
});
export type MessageItemT = z.infer<typeof MessageItem>;

// todo_add
export const AddRequest = z.object({
  roomId: z.string(),
  text: z.string(),
});
export const AddResponse = z.object({
  message: MessageItem,
});
export type AddRequestT = z.infer<typeof AddRequest>;
export type AddResponseT = z.infer<typeof AddResponse>;

// todo_delete
export const DeleteRequest = z.object({
  msgId: z.string(),
  roomId: z.string(),
});
export const DeleteResponse = z.object({});
export type DeleteRequestT = z.infer<typeof DeleteRequest>;
export type DeleteResponseT = z.infer<typeof DeleteResponse>;

// todo_add
export const ListRequest = z.object({
  roomId: z.string(),
});
export const ListResponse = z.object({
  messages: z.array(MessageItem),
});
export type ListRequestT = z.infer<typeof ListRequest>;
export type ListResponseT = z.infer<typeof ListResponse>;

export interface MessageService {
  add(params: AddRequestT, options: RpcOptions<AddResponseT>): void;
  list(params: DeleteRequestT, options: RpcOptions<DeleteResponseT>): void;
}

export const MessageEvent = z.object({
  object_id: z.string(),
  action: z.enum(["delete", "create", "update"]),
  topic: z.literal("todo"),
  body: MessageItem,
});
