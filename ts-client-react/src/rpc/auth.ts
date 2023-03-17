import { z } from "zod";
import { RpcOptions } from "./errors";

export const Token = z.object({
  accessToken: z.string(),
  tokenType: z.string(),
  expiresIn: z.number(),
  refreshToken: z.optional(z.string()),
});

// register()

export const RegisterRequest = z.object({
  email: z.string(),
  password: z.string(),
  name: z.string(),
  invite: z.optional(z.string()),
});
export const RegisterResponse = z.object({
  token: Token,
  created: z.boolean(),
});

export type RegisterRequestT = z.infer<typeof RegisterRequest>;
export type RegisterResponseT = z.infer<typeof RegisterResponse>;

// export type RegisterRequest = z.infer<typeof RegisterRequest>;

// authenticate()

export const AuthenticateRequest = z.object({
  email: z.string(),
  password: z.string(),
  tfaOtp: z.optional(z.string()),
  tfaType: z.optional(z.string()),
});
export const AuthenticateResponse = z.object({
  token: Token,
});

export type AuthenticateRequestT = z.infer<typeof AuthenticateRequest>;
export type AuthenticateResponseT = z.infer<typeof AuthenticateResponse>;

// verify_email()
export const VerifyEmailRequest = z.object({
  userId: z.string(),
  token: z.string(),
});
export const VerifyEmailResponse = z.object({});

export type VerifyEmailRequestT = z.infer<typeof VerifyEmailRequest>;
export type VerifyEmailResponseT = z.infer<typeof VerifyEmailResponse>;

// recover_send()
export const RecoverSendRequest = z.object({
  email: z.string(),
});
export const RecoverSendResponse = z.object({});

export type RecoverSendRequestT = z.infer<typeof RecoverSendRequest>;
export type RecoverSendResponseT = z.infer<typeof RecoverSendResponse>;

// recover_verify()
export const RecoverVerifyRequest = z.object({
  userId: z.string(),
  token: z.string(),
});
export const RecoverVerifyResponse = z.object({});

export type RecoverVerifyRequestT = z.infer<typeof RecoverVerifyRequest>;
export type RecoverVerifyResponseT = z.infer<typeof RecoverVerifyResponse>;
// recover_update()
export const RecoverUpdateRequest = z.object({
  userId: z.string(),
  token: z.string(),
  password: z.string(),
});
export const RecoverUpdateResponse = z.object({
  token: Token,
});

export type RecoverUpdateRequestT = z.infer<typeof RecoverUpdateRequest>;
export type RecoverUpdateResponseT = z.infer<typeof RecoverUpdateResponse>;
// oauth_login()
export const OauthLoginRequest = z.object({
  provider: z.string(),
  redirectUrl: z.string(),
  code: z.string(),
  state: z.optional(z.string()),
});
export const OauthLoginResponse = z.object({
  token: Token,
  created: z.boolean(),
});

export type OauthLoginRequestT = z.infer<typeof OauthLoginRequest>;
export type OauthLoginResponseT = z.infer<typeof OauthLoginResponse>;

// oauth_url()
export const OauthUrlRequest = z.object({
  provider: z.string(),
  redirectUrl: z.string(),
  state: z.optional(z.string()),
});
export const OauthUrlResponse = z.object({
  url: z.string(),
});

export type OauthUrlRequestT = z.infer<typeof OauthUrlRequest>;
export type OauthUrlResponseT = z.infer<typeof OauthUrlResponse>;

export interface AuthService {
  register(params: RegisterRequestT, options: RpcOptions<RegisterResponseT>): void;

  authenticate(params: AuthenticateRequestT, options: RpcOptions<AuthenticateResponseT>): void;

  verifyEmail(params: VerifyEmailRequestT, options: RpcOptions<VerifyEmailResponseT>): void;

  recoverSend(params: RecoverSendRequestT, options: RpcOptions<RecoverSendResponseT>): void;
  recoverVerify(params: RecoverVerifyRequestT, options: RpcOptions<RecoverVerifyResponseT>): void;
  recoverUpdate(params: RecoverUpdateRequestT, options: RpcOptions<RecoverUpdateResponseT>): void;

  oauthLogin(params: OauthLoginRequestT, options: RpcOptions<OauthLoginResponseT>): void;
  oauthUrl(params: OauthUrlRequestT, options: RpcOptions<OauthUrlResponseT>): void;
}
