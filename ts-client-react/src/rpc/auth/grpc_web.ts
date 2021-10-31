import * as grpcWeb from "grpc-web";
import { AuthService, LoginSuccess } from "./index";
import { BASE_URL } from "../utils";
import { AuthenticationServiceClient } from "../../genpb/publicapi/auth_grpc_web_pb";
import { RegisterParams, LoginParams, Token } from "../../genpb/publicapi/auth_pb";
import { handleGrpcError } from "../utils/grpc_helpers";
import { RpcOptions } from "../errors";

export function newAuthClient(): AuthService {
  const client = new AuthenticationServiceClient(BASE_URL);

  return {
    register(params, options: RpcOptions<LoginSuccess>): void {
      const req = new RegisterParams();
      req.setEmail(params.email);
      req.setName(params.name);
      req.setPassword(params.password);
      if (params.urlbase) {
        req.setUrlbase(params.urlbase);
      }

      client.register(req, undefined, (err: grpcWeb.RpcError, data: Token) => {
        if (err) {
          handleGrpcError(err, options);
          return;
        }
        const token = data.getAccessToken();
        options.onCompleted?.({ token });
      });
    },

    authenticate(params: { email: string; password: string }, options: RpcOptions<LoginSuccess>): void {
      const req = new LoginParams();
      req.setEmail(params.email);
      req.setPassword(params.password);

      client.authenticate(req, undefined, (err: grpcWeb.RpcError, data: Token) => {
        if (err) {
          handleGrpcError(err, options);
          return;
        }
        const token = data.getAccessToken();
        options.onCompleted?.({ token });
      });
    },
  };
}
