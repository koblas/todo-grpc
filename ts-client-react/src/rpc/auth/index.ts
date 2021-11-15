import { RpcOptions } from "../errors";

export interface AuthOk {}

export interface LoginSuccess {
  token: string;
}

export interface AuthService {
  register(
    params: { email: string; password: string; name: string; urlbase?: string },
    options: RpcOptions<LoginSuccess>,
  ): void;
  authenticate(params: { email: string; password: string }, options: RpcOptions<LoginSuccess>): void;
  verifyEmail(params: { token: string }, options: RpcOptions<AuthOk>): void;
  recoverSend(params: { email: string }, options: RpcOptions<AuthOk>): void;
  recoverVerify(params: { token: string }, options: RpcOptions<AuthOk>): void;
  recoverUpdate(params: { token: string; password: string }, options: RpcOptions<LoginSuccess>): void;
}
