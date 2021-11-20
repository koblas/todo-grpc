import { AuthOk, AuthService, AuthToken, LoginRegisterSuccess, OauthLoginUrl } from "./index";
import { newFetchPOST } from "../utils";
import { RpcOptions } from "../errors";

type TokenResponse = {
  accessToken: string;
};

export function newAuthClient(): AuthService {
  const client = newFetchPOST();

  return {
    register(params, options: RpcOptions<LoginRegisterSuccess>): void {
      client.call<TokenResponse, LoginRegisterSuccess>("/auth/register", params, options);
    },
    authenticate(params, options: RpcOptions<AuthToken>): void {
      client.call<TokenResponse, AuthToken>("/auth/authenticate", params, options);
    },
    verifyEmail(params, options: RpcOptions<AuthOk>): void {
      client.call<AuthOk, AuthOk>("/auth/verify/email", params, options);
    },
    recoverSend(params, options: RpcOptions<AuthOk>): void {
      client.call<AuthOk, AuthOk>("/auth/recover/send", params, options);
    },
    recoverVerify(params, options: RpcOptions<AuthOk>): void {
      client.call<AuthOk, AuthOk>("/auth/recover/verify", params, options);
    },
    recoverUpdate(params, options: RpcOptions<AuthToken>): void {
      client.call<TokenResponse, AuthToken>("/auth/recover/update", params, options);
    },
    oauthRedirect(
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      params: { provider: string; redirectUrl: string; state: string },
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      options: RpcOptions<OauthLoginUrl>,
    ): void {
      client.call<OauthLoginUrl, OauthLoginUrl>("/auth/oauth/url", params, options);
    },
    oauthLogin(
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      params: { provider: string; redirectUrl: string; code: string; state: string },
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      options: RpcOptions<LoginRegisterSuccess>,
    ): void {
      client.call<TokenResponse, LoginRegisterSuccess>("/auth/oauth/login", params, options);
    },
  };
}
