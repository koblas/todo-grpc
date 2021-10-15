import { AuthService, LoginSuccess } from "./index";
import { newFetchClient } from "../utils";

export function newAuthClient(): AuthService {
  const client = newFetchClient();

  return {
    async login(username: string, password: string): Promise<LoginSuccess> {
      const data = await client.POST<{ accessToken: string }>("/auth/login", { username, password });

      return { token: data.accessToken };
    },
  };
}
