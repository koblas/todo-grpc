import { AuthService, LoginSuccess } from "./index";
import { AuthenticationServiceClientImpl, GrpcWebImpl } from "../../models/auth";
import { BASE_URL } from "../utils";

const rpc = new GrpcWebImpl(BASE_URL, {
  debug: false,
});

export function newAuthClient(): AuthService {
  const client = new AuthenticationServiceClientImpl(rpc);

  return {
    async login(username: string, password: string): Promise<LoginSuccess> {
      const response = await client.login({ username, password });

      return { token: response.accessToken };
    },
  };
}
