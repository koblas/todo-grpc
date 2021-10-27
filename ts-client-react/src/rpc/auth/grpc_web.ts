import { AuthService, LoginSuccess } from "./index";
import { AuthenticationServiceClientImpl, GrpcWebImpl } from "../../models/auth";
import { BASE_URL } from "../utils";
import { must } from "../../util/assert";

const rpc = new GrpcWebImpl(BASE_URL, {
  debug: false,
});

export function newAuthClient(): AuthService {
  const client = new AuthenticationServiceClientImpl(rpc);

  return {
    async register(params): Promise<LoginSuccess> {
      const response = await client.register({
        email: params.email,
        name: params.name,
        password: params.password,
        urlbase: params.urlbase,
      });

      if (response.errors) {
        throw new Error("Unable to authenticate");
      }

      return { token: must(response.token).accessToken };
    },

    async authenticate(email: string, password: string): Promise<LoginSuccess> {
      const response = await client.authenticate({ email, password });

      if (response.errors) {
        throw new Error("Unable to authenticate");
      }

      return { token: must(response.token).accessToken };
    },
  };
}
