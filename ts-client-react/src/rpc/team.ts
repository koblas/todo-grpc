import { z } from "zod";
import { RpcOptions } from "./errors";

export const TeamMember = z.object({
  id: z.string(),
  team_name: z.optional(z.string()),
  user_id: z.string(),
  role: z.string(),
  status: z.string(),
});
export type TeamMemberT = z.infer<typeof TeamMember>;

// team_invite
export const TeamInviteRequest = z.object({
  team_id: z.string(),
  user_id: z.optional(z.string()),
  email: z.optional(z.string()),
});
export const TeamInviteResponse = z.object({});
export type TeamInviteRequestT = z.infer<typeof TeamInviteRequest>;
export type TeamInviteResponseT = z.infer<typeof TeamInviteResponse>;

// team_create
export const TeamCreateRequest = z.object({
  name: z.string(),
});
export const TeamCreateResponse = z.object({
  team: TeamMember,
});
export type TeamCreateRequestT = z.infer<typeof TeamCreateRequest>;
export type TeamCreateResponseT = z.infer<typeof TeamCreateResponse>;

// team_delete
export const TeamDeleteRequest = z.object({
  team_id: z.string(),
});
export const TeamDeleteResponse = z.object({});
export type TeamDeleteRequestT = z.infer<typeof TeamDeleteRequest>;
export type TeamDeleteResponseT = z.infer<typeof TeamDeleteResponse>;

// team_create
export const TeamListRequest = z.object({});
export const TeamListResponse = z.object({
  teams: z.array(TeamMember),
});
export type TeamListRequestT = z.infer<typeof TeamListRequest>;
export type TeamListResponseT = z.infer<typeof TeamListResponse>;

export interface TeamService {
  team_create(params: TeamCreateRequestT, options: RpcOptions<TeamCreateResponseT>): void;
  team_delete(params: TeamDeleteRequestT, options: RpcOptions<TeamDeleteResponseT>): void;
  team_list(params: TeamListRequestT, options: RpcOptions<TeamListResponseT>): void;
  team_invite(params: TeamInviteRequestT, options: RpcOptions<TeamInviteResponseT>): void;
}

export const TeamMemberEvent = z.object({
  object_id: z.string(),
  action: z.enum(["delete", "create", "update"]),
  topic: z.literal("todo"),
  body: TeamMember,
});
