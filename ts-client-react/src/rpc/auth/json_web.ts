import { AuthService, LoginSuccess } from "./index";
import { newFetchClient } from "../utils";
import { RpcOptions } from "../errors";
import { handleJsonError } from "../utils/json_helpers";

type TokenResponse = {
  accessToken: string;
};

export function newAuthClient(): AuthService {
  const client = newFetchClient();

  return {
    register(params, options: RpcOptions<LoginSuccess>): void {
      client
        .POST<TokenResponse>("/auth/register", {
          email: params.email,
          password: params.password,
          name: params.name,
          urlbase: params.urlbase ?? null,
        })
        .then((data) => {
          options.onCompleted?.({ token: data.accessToken });
        })
        .catch((err) => {
          handleJsonError(err, options);
        });
    },
    authenticate(params, options: RpcOptions<LoginSuccess>): void {
      client
        .POST<TokenResponse>("/auth/authenticate", params)
        .then((data) => {
          options.onCompleted?.({ token: data.accessToken });
        })
        .catch((err) => {
          handleJsonError(err, options);
        });
    },
  };
}
