import { z } from "zod";
import { RpcOptions } from "./errors";

export const MessageItem = z.object({
  id: z.string(),
  room_id: z.string(),
  sender: z.string(),
  text: z.string(),
});
export type MessageItemT = z.infer<typeof MessageItem>;

export const RoomItem = z.object({
  id: z.string(),
  name: z.string(),
});
export type RoomItemT = z.infer<typeof RoomItem>;

// message create
export const MsgCreateRequest = z.object({
  room_id: z.string(),
  text: z.string(),
});
export const MsgCreateResponse = z.object({
  message: MessageItem,
});
export type MsgCreateRequestT = z.infer<typeof MsgCreateRequest>;
export type MsgCreateResponseT = z.infer<typeof MsgCreateResponse>;

// todo_delete
export const DeleteRequest = z.object({
  msg_id: z.string(),
  room_id: z.string(),
});
export const DeleteResponse = z.object({});
export type DeleteRequestT = z.infer<typeof DeleteRequest>;
export type DeleteResponseT = z.infer<typeof DeleteResponse>;

// todo_add
export const MsgListRequest = z.object({
  room_id: z.string(),
});
export const MsgListResponse = z.object({
  messages: z.array(MessageItem),
});
export type MsgListRequestT = z.infer<typeof MsgListRequest>;
export type MsgListResponseT = z.infer<typeof MsgListResponse>;

//
export const RoomListRequest = z.object({});
export const RoomListResponse = z.object({
  rooms: z.array(RoomItem),
});
export type RoomListRequestT = z.infer<typeof RoomListRequest>;
export type RoomListResponseT = z.infer<typeof RoomListResponse>;

//
export const RoomJoinRequest = z.object({});
export const RoomJoinResponse = z.object({
  room: RoomItem,
});
export type RoomJoinRequestT = z.infer<typeof RoomJoinRequest>;
export type RoomJoinResponseT = z.infer<typeof RoomJoinResponse>;

export interface MessageService {
  add(params: MsgCreateRequestT, options: RpcOptions<MsgCreateResponseT>): void;
  list(params: DeleteRequestT, options: RpcOptions<DeleteResponseT>): void;
}

export const MessageEvent = z.object({
  object_id: z.string(),
  action: z.enum(["delete", "create", "update"]),
  topic: z.literal("todo"),
  body: MessageItem,
});
