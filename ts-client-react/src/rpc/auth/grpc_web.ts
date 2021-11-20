import * as grpcWeb from "grpc-web";
import { AuthOk, AuthService, AuthToken, LoginRegisterSuccess, OauthLoginUrl } from "./index";
import { BASE_URL } from "../utils";
import { AuthenticationServiceClient } from "../../genpb/publicapi/auth_grpc_web_pb";
import { RegisterParams, LoginParams, Token, TokenRegister } from "../../genpb/publicapi/auth_pb";
import { handleGrpcError } from "../utils/grpc_helpers";
import { RpcOptions } from "../errors";

export function newAuthClient(): AuthService {
  const client = new AuthenticationServiceClient(BASE_URL);

  return {
    register(params, options: RpcOptions<LoginRegisterSuccess>): void {
      const req = new RegisterParams();
      req.setEmail(params.email);
      req.setName(params.name);
      req.setPassword(params.password);

      client.register(req, undefined, (err: grpcWeb.RpcError, data: TokenRegister) => {
        if (err) {
          handleGrpcError(err, options);
          return;
        }
        const value = data.toObject() as Required<TokenRegister.AsObject>;
        options.onCompleted?.(value);
      });
    },

    authenticate(params: { email: string; password: string }, options: RpcOptions<AuthToken>): void {
      const req = new LoginParams();
      req.setEmail(params.email);
      req.setPassword(params.password);

      client.authenticate(req, undefined, (err: grpcWeb.RpcError, data: Token) => {
        if (err) {
          handleGrpcError(err, options);
          return;
        }
        options.onCompleted?.(data.toObject());
      });
    },

    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    verifyEmail(params: { token: string }, options: RpcOptions<AuthOk>): void {
      throw new Error("Not Impelmented");
    },
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    recoverSend(params: { email: string }, options: RpcOptions<AuthOk>): void {
      throw new Error("Not Impelmented");
    },
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    recoverVerify(params: { token: string }, options: RpcOptions<AuthOk>): void {
      throw new Error("Not Impelmented");
    },
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    recoverUpdate(params: { token: string; password: string }, options: RpcOptions<AuthOk>): void {
      throw new Error("Not Impelmented");
    },
    oauthRedirect(
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      params: { provider: string; redirectUrl: string; state: string },
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      options: RpcOptions<OauthLoginUrl>,
    ): void {
      throw new Error("Not Impelmented");
    },
    oauthLogin(
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      params: { provider: string; redirectUrl: string; code: string; state: string },
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      options: RpcOptions<LoginRegisterSuccess>,
    ): void {
      throw new Error("Not Impelmented");
    },
  };
}
