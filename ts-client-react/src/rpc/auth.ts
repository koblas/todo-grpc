import { z } from "zod";
import { RpcOptions } from "../errors";

export const AuthOkResponse = z.object({});

export const AuthTokenResponse = z.object({
  access_token: z.string(),
});

export const LoginRegisterResponse = z.object({
  created: z.boolean(),
  token: AuthTokenResponse,
});

export const OauthLoginUrlResponse = z.object({
  url: z.string(),
});

export type LoginRegisterSuccess = z.infer<typeof LoginRegisterResponse>;
export type OauthLoginUrl = z.infer<typeof OauthLoginUrlResponse>;
export type AuthToken = z.infer<typeof AuthTokenResponse>;
export type AuthOk = z.infer<typeof AuthOkResponse>;

export interface AuthService {
  register(
    params: { email: string; password: string; name: string; urlbase?: string },
    options: RpcOptions<LoginRegisterSuccess>,
  ): void;
  authenticate(params: { email: string; password: string }, options: RpcOptions<AuthToken>): void;
  verifyEmail(params: { token: string }, options: RpcOptions<AuthOk>): void;
  recoverSend(params: { email: string }, options: RpcOptions<AuthOk>): void;
  recoverVerify(params: { token: string }, options: RpcOptions<AuthOk>): void;
  recoverUpdate(params: { token: string; password: string }, options: RpcOptions<AuthToken>): void;
  oauthLogin(
    params: { provider: string; redirectUrl: string; code: string; state: string },
    options: RpcOptions<LoginRegisterSuccess>,
  ): void;
  oauthRedirect(
    params: { provider: string; redirectUrl: string; state: string },
    options: RpcOptions<OauthLoginUrl>,
  ): void;
}
