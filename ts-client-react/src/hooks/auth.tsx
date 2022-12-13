import { z } from "zod";
import { useMutation, useQueryClient } from "react-query";

import { RpcMutation, RpcOptions } from "../rpc/errors";
import {
  LoginRegisterResponse,
  AuthOk,
  AuthToken,
  LoginRegisterSuccess,
  OauthLoginUrl,
  AuthOkResponse,
  AuthTokenResponse,
  OauthLoginUrlResponse,
} from "../rpc/auth";
import { useAuthStore } from "../store/useAuthStore";
import { newFetchClient } from "../rpc/utils";
import { buildCallbacksTyped } from "../rpc/utils/helper";
import { Json } from "../types/json";

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
  redirect_url: string;
};

export type OauthAssociateParms = {
  provider: string;
  redirect_url: string;
  code: string;
  state: string;
};

export function useAuth() {
  const { token, setToken } = useAuthStore((s) => s);
  const queryClient = useQueryClient();
  const client = newFetchClient();

  return {
    token,
    subscribe: useAuthStore.subscribe,
    isAuthenticated: !!token,

    mutations: {
      useRegister(): RpcMutation<RegisterParams, LoginRegisterSuccess> {
        // const mutation = useMutation<z.infer<typeof RegisterResponse>, unknown, RegisterParams>(
        const mutation = useMutation(
          (data: RegisterParams) => client.POST<LoginRegisterSuccess>("/v1/auth/register", data),
          {},
        );

        function action(data: RegisterParams, handlers?: RpcOptions<z.infer<typeof LoginRegisterResponse>>) {
          mutation.mutate(
            data,
            buildCallbacksTyped(
              queryClient,
              LoginRegisterResponse,
              {
                onCompleted(result) {
                  setToken(result.token.access_token);
                },
              },
              handlers,
            ),
          );
        }

        return [action, { loading: mutation.isLoading }];
      },
      useEmailConfirm(): RpcMutation<RecoverVerifyParams, AuthOk> {
        const mutation = useMutation<AuthOk, unknown, RecoverVerifyParams>(
          (data: RecoverVerifyParams) => client.POST("/v1/auth/verify_email", data as Json),
          {},
        );

        // function action(data: RecoverVerifyParams, handlers?: RpcOptions<z.infer<typeof AuthOkResponse>>) {
        function action(data: RecoverVerifyParams, handlers?: RpcOptions<AuthOk>) {
          mutation.mutate(data, buildCallbacksTyped(queryClient, AuthOkResponse, handlers));
        }

        return [action, { loading: mutation.isLoading }];
      },

      useLogin(): RpcMutation<LoginParams, AuthToken> {
        const mutation = useMutation<AuthToken, unknown, LoginParams>(
          (data: LoginParams) => client.POST("/v1/auth/authenticate", data as Json),
          {},
        );

        // function action(data: RecoverVerifyParams, handlers?: RpcOptions<z.infer<typeof AuthOkResponse>>) {
        function action(data: LoginParams, handlers?: RpcOptions<AuthToken>) {
          mutation.mutate(
            data,
            buildCallbacksTyped(
              queryClient,
              AuthTokenResponse,
              {
                onCompleted(result) {
                  setToken(result.access_token);
                },
              },
              handlers,
            ),
          );
        }

        return [action, { loading: mutation.isLoading }];
      },
      useRecoverSend(): RpcMutation<RecoverSendParams, void> {
        const mutation = useMutation<void, unknown, RecoverSendParams>(
          (data: RecoverSendParams) => client.POST("/v1/auth/recover_send", data as Json),
          {},
        );

        // function action(data: RecoverVerifyParams, handlers?: RpcOptions<z.infer<typeof AuthOkResponse>>) {
        function action(data: RecoverSendParams, handlers?: RpcOptions<void>) {
          mutation.mutate(data, buildCallbacksTyped(queryClient, AuthOkResponse, handlers));
        }

        return [action, { loading: mutation.isLoading }];
      },
      useRecoveryVerify(): RpcMutation<RecoverVerifyParams, void> {
        const mutation = useMutation<void, unknown, RecoverVerifyParams>(
          (data: RecoverVerifyParams) => client.POST("/v1/auth/recover_verify", data as Json),
          {},
        );

        // function action(data: RecoverVerifyParams, handlers?: RpcOptions<z.infer<typeof AuthOkResponse>>) {
        function action(data: RecoverVerifyParams, handlers?: RpcOptions<void>) {
          mutation.mutate(data, buildCallbacksTyped(queryClient, AuthOkResponse, handlers));
        }

        return [action, { loading: mutation.isLoading }];
      },
      useRecoveryUpdate(): RpcMutation<RecoverUpdateParams, AuthToken> {
        const mutation = useMutation<AuthToken, unknown, RecoverUpdateParams>(
          (data: RecoverUpdateParams) => client.POST("/v1/auth/recover_update", data as Json),
          {},
        );

        // function action(data: RecoverVerifyParams, handlers?: RpcOptions<z.infer<typeof AuthOkResponse>>) {
        function action(data: RecoverUpdateParams, handlers?: RpcOptions<AuthToken>) {
          mutation.mutate(
            data,
            buildCallbacksTyped(
              queryClient,
              AuthTokenResponse,
              {
                onCompleted(result) {
                  setToken(result.access_token);
                },
              },
              handlers,
            ),
          );
        }

        return [action, { loading: mutation.isLoading }];
      },

      useLogout() {
        return (options?: RpcOptions<void>) => {
          setToken(null);

          options?.onCompleted?.(undefined, undefined);
        };
      },

      useOauthRedirect(): RpcMutation<OauthUrlParams, OauthLoginUrl> {
        const mutation = useMutation<OauthLoginUrl, unknown, OauthUrlParams>((data: OauthUrlParams) =>
          client.POST<OauthLoginUrl>("/v1/auth/oauth_url", data),
        );

        // function action(data: RecoverVerifyParams, handlers?: RpcOptions<z.infer<typeof AuthOkResponse>>) {
        function action(data: OauthUrlParams, handlers?: RpcOptions<OauthLoginUrl>) {
          mutation.mutate(data, buildCallbacksTyped(queryClient, OauthLoginUrlResponse, handlers));
        }

        return [action, { loading: mutation.isLoading }];
      },

      useOauthLogin(): RpcMutation<OauthAssociateParms, LoginRegisterSuccess> {
        const mutation = useMutation<LoginRegisterSuccess, unknown, OauthAssociateParms>(
          (data: OauthAssociateParms) => client.POST("/v1/auth/oauth_login", data as Json),
          {},
        );

        // function action(data: RecoverVerifyParams, handlers?: RpcOptions<z.infer<typeof AuthOkResponse>>) {
        function action(data: OauthAssociateParms, handlers?: RpcOptions<LoginRegisterSuccess>) {
          mutation.mutate(
            data,
            buildCallbacksTyped(
              queryClient,
              LoginRegisterResponse,
              {
                onCompleted(result) {
                  setToken(result.token.access_token);
                },
              },
              handlers,
            ),
          );
        }

        return [action, { loading: mutation.isLoading }];
      },
    },
  };
}
