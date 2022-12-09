import { Json, JsonObject } from "../../types/json";
import { RpcOptions } from "../errors";

export interface User extends JsonObject {
  name: string;
  email: string;
}

export interface UserService {
  getUser(params: Json, options: RpcOptions<User>): void;
  updateUser(params: Json, options: RpcOptions<User>): void;
}
