import { RpcOptions } from "../errors";

export interface AuthOk {}

export interface AuthToken {
  accessToken: string;
}

export interface LoginSuccess {
  token: AuthToken;
}

export interface LoginRegisterSuccess {
  created: boolean;
  token: AuthToken;
}

export interface OauthLoginUrl {
  url: string;
}

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
