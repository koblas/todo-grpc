import { RpcOptions, RpcResult } from "../errors";

export interface LoginSuccess {
  token: string;
}

export interface AuthService {
  register(
    params: { email: string; password: string; name: string; urlbase?: string },
    options: RpcOptions<LoginSuccess>,
  ): void;
  authenticate(params: { email: string; password: string }, options: RpcOptions<LoginSuccess>): void;
}
