import { useState } from "react";
import { newAuthClient } from "../rpc/auth/factory";
import { useNetworkContextErrors } from "./network";
import { RpcMutation, RpcOptions, RpcError } from "../rpc/errors";
import { randomString } from "../util/randomeString";
import { AuthToken, LoginRegisterSuccess, OauthLoginUrl } from "../rpc/auth";
import { useAuthStore } from "../store/useAuthStore";

const authClient = newAuthClient("json");

export type LoginParams = {
  email: string;
  password: string;
};

export type RecoverSendParams = {
  email: string;
};

export type RecoverVerifyParams = {
  userId: string;
  token: string;
};

export type RecoverUpdateParams = {
  userId: string;
  token: string;
  password: string;
};

export type RegisterParams = {
  name: string;
  email: string;
  password: string;
};

export type OauthUrlParams = {
  provider: string;
  returnUrl: string;
};

export type OauthAssociateParms = {
  provider: string;
  redirectUrl: string;
  code: string;
  state: string;
};

export function useAuth() {
  const { token, setToken } = useAuthStore((s) => s);
  const addHandler = useNetworkContextErrors();

  return {
    token,
    subscribe: useAuthStore.subscribe,
    isAuthenticated: !!token,
    mutations: {
      useRegister(): RpcMutation<RegisterParams, LoginRegisterSuccess> {
        const [data, setData] = useState<LoginRegisterSuccess | undefined>();
        const [loading, setLoading] = useState(false);
        const [error, setError] = useState<RpcError | undefined>(undefined);

        const func = (params: RegisterParams, options?: RpcOptions<LoginRegisterSuccess>) => {
          setLoading(true);
          authClient.register(
            params,
            addHandler(
              {
                onCompleted(input) {
                  setData(input);
                  setToken(input.token.access_token);
                },
                onError(err: RpcError) {
                  setError(err);
                },
                onFinished() {
                  setLoading(false);
                },
              },
              options,
            ),
          );
        };

        return [func, { data, loading, error }];
      },
      useEmailConfirm(): RpcMutation<RecoverVerifyParams, void> {
        const [loading, setLoading] = useState(false);
        const [error, setError] = useState<RpcError | undefined>(undefined);

        const func = (params: RecoverVerifyParams, options?: RpcOptions<void>) => {
          setLoading(true);
          authClient.verifyEmail(
            params,
            addHandler(
              {
                onError(err: RpcError) {
                  setError(err);
                },
                onFinished() {
                  setLoading(false);
                },
              },
              options,
            ),
          );
        };

        return [func, { data: undefined, loading, error }];
      },
      useLogin(): RpcMutation<LoginParams, AuthToken> {
        const [data, setData] = useState<AuthToken | undefined>();
        const [loading, setLoading] = useState(false);
        const [error, setError] = useState<RpcError | undefined>(undefined);

        const func = (params: LoginParams, options?: RpcOptions<AuthToken>) => {
          setLoading(true);
          authClient.authenticate(
            params,
            addHandler(
              {
                onError(err: RpcError) {
                  setError(err);
                },
                onFinished() {
                  setLoading(false);
                },
                onCompleted: (input) => {
                  setData(input);
                  setToken(input.access_token);
                },
              },
              options,
            ),
          );
        };

        return [func, { data, loading, error }];
      },
      useRecoverSend(): RpcMutation<RecoverSendParams, void> {
        const [loading, setLoading] = useState(false);
        const [error, setError] = useState<RpcError | undefined>(undefined);

        const func = (params: RecoverSendParams, options?: RpcOptions<void>) => {
          setLoading(true);
          authClient.recoverSend(
            params,
            addHandler(
              {
                onError(err: RpcError) {
                  setError(err);
                },
                onFinished() {
                  setLoading(false);
                },
              },
              options,
            ),
          );
        };

        return [func, { data: undefined, loading, error }];
      },
      useRecoveryVerify(): RpcMutation<RecoverVerifyParams, void> {
        const [loading, setLoading] = useState(false);
        const [error, setError] = useState<RpcError | undefined>(undefined);

        const func = (params: RecoverVerifyParams, options?: RpcOptions<void>) => {
          setLoading(true);
          authClient.recoverVerify(
            params,
            addHandler(
              {
                onError(err: RpcError) {
                  setError(err);
                },
                onFinished() {
                  setLoading(false);
                },
              },
              options,
            ),
          );
        };

        return [func, { data: undefined, loading, error }];
      },
      useRecoveryUpdate(): RpcMutation<RecoverUpdateParams, AuthToken> {
        const [data, setData] = useState<AuthToken | undefined>();
        const [loading, setLoading] = useState(false);
        const [error, setError] = useState<RpcError | undefined>(undefined);

        const func = (params: RecoverUpdateParams, options?: RpcOptions<AuthToken>) => {
          setLoading(true);
          authClient.recoverUpdate(
            params,
            addHandler(
              {
                onCompleted: (input) => {
                  setData(input);
                  setToken(input.access_token);
                },
                onError(err: RpcError) {
                  setError(err);
                },
                onFinished() {
                  setLoading(false);
                },
              },
              options,
            ),
          );
        };

        return [func, { data, loading, error }];
      },
      useLogout() {
        return (options?: RpcOptions<void>) => {
          setToken(null);

          options?.onCompleted?.();
        };
      },

      // OAuth Functionality
      useOauthRedirect(): RpcMutation<OauthUrlParams, { url: string }> {
        const [data, setData] = useState<OauthLoginUrl>({ url: "" });
        const [loading, setLoading] = useState(false);
        const [error, setError] = useState<RpcError | undefined>(undefined);

        const func = (params: OauthUrlParams, options?: RpcOptions<OauthLoginUrl>) => {
          const st = randomString(20);

          authClient.oauthRedirect(
            {
              provider: params.provider,
              redirectUrl: params.returnUrl,
              state: st,
            },
            addHandler(
              {
                onCompleted: (dvalue) => {
                  setData(dvalue);
                },
                onError(err: RpcError) {
                  setError(err);
                },
                onFinished() {
                  setLoading(false);
                },
              },
              options,
            ),
          );
        };

        return [func, { data, loading, error }];
      },

      useOauthLogin(): RpcMutation<OauthAssociateParms, LoginRegisterSuccess> {
        const [data, setData] = useState<LoginRegisterSuccess | undefined>();
        const [loading, setLoading] = useState(false);
        const [error, setError] = useState<RpcError | undefined>(undefined);

        const func = (
          params: { provider: string; code: string; state: string; redirectUrl: string },
          options?: RpcOptions<LoginRegisterSuccess>,
        ) => {
          authClient.oauthLogin(
            params,
            addHandler(
              {
                onCompleted: (input) => {
                  setData(data);
                  setToken(input.token.access_token);
                },
                onError(err: RpcError) {
                  setError(err);
                },
                onFinished() {
                  setLoading(false);
                },
              },
              options,
            ),
          );
        };

        return [func, { data, loading, error }];
      },
    },
  };
}
