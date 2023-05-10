import { z } from "zod";
import { RpcOptions } from "./errors";

export const User = z.object({
  id: z.string(),
  name: z.string(),
  email: z.string(),
  avatar_url: z.optional(z.nullable(z.string())),
});
export type UserT = z.infer<typeof User>;

// get_user
export const GetUserRequest = z.object({});
export const GetUserResponse = z.object({
  user: User,
});
export type GetUserRequestT = z.infer<typeof GetUserRequest>;
export type GetUserResponseT = z.infer<typeof GetUserResponse>;

// update_user
export const UpdateUserRequest = z.object({
  email: z.optional(z.string()),
  name: z.optional(z.string()),
  password: z.optional(z.string()),
  password_new: z.optional(z.string()),
});
export const UpdateUserResponse = z.object({
  user: User,
});
export type UpdateUserRequestT = z.infer<typeof UpdateUserRequest>;
export type UpdateUserResponseT = z.infer<typeof UpdateUserResponse>;

export interface UserService {
  get_user(params: GetUserRequestT, options: RpcOptions<GetUserResponseT>): void;
  update_user(params: UpdateUserRequestT, options: RpcOptions<UpdateUserResponseT>): void;
}
