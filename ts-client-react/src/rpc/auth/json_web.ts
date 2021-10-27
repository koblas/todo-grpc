import { AuthService, LoginSuccess } from "./index";
import { newFetchClient } from "../utils";
import { must } from "../../util/assert";

export function newAuthClient(): AuthService {
  const client = newFetchClient();

  return {
    async register(params): Promise<LoginSuccess> {
      const data = await client.POST<{ token?: { accessToken: string } }>("/auth/register", {
        email: params.email,
        password: params.password,
        name: params.name,
        urlbase: params.urlbase ?? null,
      });

      return { token: must(data.token).accessToken };
    },

    async authenticate(email: string, password: string): Promise<LoginSuccess> {
      const data = await client.POST<{ token?: { accessToken: string } }>("/auth/authenticate", { email, password });

      return { token: must(data.token).accessToken };
    },
  };
}
