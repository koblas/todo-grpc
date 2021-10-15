import { AuthService, LoginSuccess } from "./index";
import { todoServiceClientImpl, GrpcWebImpl } from "../../models/auth";

const rpc = new GrpcWebImpl("http://localhost:8080", {
  debug: false,
});

export function newAuthClient(): AuthService {
  const client = new todoServiceClientImpl(rpc);

  return {
    async login(username: string, password: string): Promise<LoginSuccess> {
      return await client.login({});
    },
  };
}
