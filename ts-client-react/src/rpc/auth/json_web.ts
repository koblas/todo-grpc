import { AuthOk, AuthService, AuthToken, LoginRegisterSuccess, OauthLoginUrl } from "./index";
import { newFetchPOST } from "../utils";
import { RpcOptions } from "../errors";

type TokenResponse = {
  access_token: string;
};

export function newAuthClient(): AuthService {
  const client = newFetchPOST();
  const base = "/v1/auth";

  return {
    register(params, options: RpcOptions<LoginRegisterSuccess>): void {
      client.call<TokenResponse, LoginRegisterSuccess>(`${base}/register`, params, options);
    },
    authenticate(params, options: RpcOptions<AuthToken>): void {
      client.call<TokenResponse, AuthToken>(`${base}/authenticate`, params, options);
    },
    verifyEmail(params, options: RpcOptions<AuthOk>): void {
      client.call<AuthOk, AuthOk>(`${base}/verify_email`, params, options);
    },
    recoverSend(params, options: RpcOptions<AuthOk>): void {
      client.call<AuthOk, AuthOk>(`${base}/recover_send`, params, options);
    },
    recoverVerify(params, options: RpcOptions<AuthOk>): void {
      client.call<AuthOk, AuthOk>(`${base}/recover_verify`, params, options);
    },
    recoverUpdate(params, options: RpcOptions<AuthToken>): void {
      client.call<TokenResponse, AuthToken>(`${base}/recover_update`, params, options);
    },
    oauthRedirect(
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      params: { provider: string; redirectUrl: string; state: string },
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      options: RpcOptions<OauthLoginUrl>,
    ): void {
      client.call<OauthLoginUrl, OauthLoginUrl>(`${base}/oauth_url`, params, options);
    },
    oauthLogin(
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      params: { provider: string; redirectUrl: string; code: string; state: string },
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      options: RpcOptions<LoginRegisterSuccess>,
    ): void {
      client.call<TokenResponse, LoginRegisterSuccess>(`${base}/oauth_login`, params, options);
    },
  };
}
