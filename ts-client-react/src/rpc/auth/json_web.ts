import { AuthOk, AuthService, LoginSuccess } from "./index";
import { newFetchPOST } from "../utils";
import { RpcOptions } from "../errors";

type TokenResponse = {
  accessToken: string;
};

export function newAuthClient(): AuthService {
  const client = newFetchPOST();

  return {
    register(params, options: RpcOptions<LoginSuccess>): void {
      client.call<TokenResponse, LoginSuccess>("/auth/register", params, options, (data) => ({
        token: data.accessToken,
      }));
    },
    authenticate(params, options: RpcOptions<LoginSuccess>): void {
      client.call<TokenResponse, LoginSuccess>("/auth/authenticate", params, options, (data) => ({
        token: data.accessToken,
      }));
    },
    verifyEmail(params, options: RpcOptions<AuthOk>): void {
      client.call<AuthOk, AuthOk>("/auth/verify/email", params, options, () => ({}));
    },
    recoverSend(params, options: RpcOptions<AuthOk>): void {
      client.call<AuthOk, AuthOk>("/auth/recover/send", params, options, () => ({}));
    },
    recoverVerify(params, options: RpcOptions<AuthOk>): void {
      client.call<AuthOk, AuthOk>("/auth/recover/verify", params, options, () => ({}));
    },
    recoverUpdate(params, options: RpcOptions<AuthOk>): void {
      client.call<TokenResponse, LoginSuccess>("/auth/recover/update", params, options, (data) => ({
        token: data.accessToken,
      }));
    },
  };
}
