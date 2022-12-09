import { UserService, User } from "./index";
import { newFetchPOST } from "../utils";
import { RpcOptions } from "../errors";

export function newUserClient(token: string | null): UserService {
  const client = newFetchPOST({ token });
  const base = "/v1/user";

  return {
    getUser(params, options: RpcOptions<User>): void {
      client.call(`${base}/get_user`, params, options);
    },
    updateUser(params, options: RpcOptions<User>): void {
      client.call(`${base}/update_user`, params, options);
    },
  };
}
