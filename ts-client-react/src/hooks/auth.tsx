import { useMutation, useQueryClient } from "@tanstack/react-query";

import { RpcMutation, RpcOptions } from "../rpc/errors";
import * as rpcAuth from "../rpc/auth";
import { useAuthStore } from "../store/useAuthStore";
import { newFetchClient } from "../rpc/utils";
import { buildCallbacksTyped } from "../rpc/utils/helper";
import { Json } from "../types/json";

export function useAuth() {
  const { token, setToken } = useAuthStore((s) => s);
  const queryClient = useQueryClient();
  const client = newFetchClient();

  return {
    token,
    subscribe: useAuthStore.subscribe,
    isAuthenticated: !!token,

    mutations: {
      useRegister(): RpcMutation<rpcAuth.RegisterRequestT, rpcAuth.RegisterResponseT> {
        const mutation = useMutation(
          (data: rpcAuth.RegisterRequestT) => client.POST<rpcAuth.RegisterResponseT>("/v1/auth/register", data),
          {},
        );

        function action(data: rpcAuth.RegisterRequestT, handlers?: RpcOptions<rpcAuth.RegisterResponseT>) {
          mutation.mutate(
            data,
            buildCallbacksTyped(
              queryClient,
              rpcAuth.RegisterResponse,
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
      useEmailConfirm(): RpcMutation<rpcAuth.VerifyEmailRequestT, rpcAuth.VerifyEmailResponseT> {
        const mutation = useMutation(
          (data: rpcAuth.VerifyEmailRequestT) =>
            client.POST<rpcAuth.VerifyEmailResponseT>("/v1/auth/verify_email", data),
          {},
        );

        // function action(data: RecoverVerifyParams, handlers?: RpcOptions<z.infer<typeof AuthOkResponse>>) {
        function action(data: rpcAuth.VerifyEmailRequestT, handlers?: RpcOptions<rpcAuth.VerifyEmailResponseT>) {
          mutation.mutate(data, buildCallbacksTyped(queryClient, rpcAuth.VerifyEmailResponse, handlers));
        }

        return [action, { loading: mutation.isLoading }];
      },

      useLogin(): RpcMutation<rpcAuth.AuthenticateRequestT, rpcAuth.AuthenticateResponseT> {
        const mutation = useMutation(
          (data: rpcAuth.AuthenticateRequestT) =>
            client.POST<rpcAuth.AuthenticateResponseT>("/v1/auth/authenticate", data),
          {},
        );

        // function action(data: RecoverVerifyParams, handlers?: RpcOptions<z.infer<typeof AuthOkResponse>>) {
        function action(data: rpcAuth.AuthenticateRequestT, handlers?: RpcOptions<rpcAuth.AuthenticateResponseT>) {
          mutation.mutate(
            data,
            buildCallbacksTyped(
              queryClient,
              rpcAuth.AuthenticateResponse,
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
      useRecoverSend(): RpcMutation<rpcAuth.RecoverSendRequestT, rpcAuth.RecoverSendResponseT> {
        const mutation = useMutation(
          (data: rpcAuth.RecoverSendRequestT) =>
            client.POST<rpcAuth.RecoverSendResponseT>("/v1/auth/recover_send", data as Json),
          {},
        );

        // function action(data: RecoverVerifyParams, handlers?: RpcOptions<z.infer<typeof AuthOkResponse>>) {
        function action(data: rpcAuth.RecoverSendRequestT, handlers?: RpcOptions<rpcAuth.RecoverSendResponseT>) {
          mutation.mutate(data, buildCallbacksTyped(queryClient, rpcAuth.RecoverSendResponse, handlers));
        }

        return [action, { loading: mutation.isLoading }];
      },
      useRecoveryVerify(): RpcMutation<rpcAuth.RecoverVerifyRequestT, rpcAuth.RecoverVerifyResponseT> {
        const mutation = useMutation(
          (data: rpcAuth.RecoverVerifyRequestT) =>
            client.POST<rpcAuth.RecoverVerifyResponseT>("/v1/auth/recover_verify", data as Json),
          {},
        );

        // function action(data: RecoverVerifyParams, handlers?: RpcOptions<z.infer<typeof AuthOkResponse>>) {
        function action(data: rpcAuth.RecoverVerifyRequestT, handlers?: RpcOptions<rpcAuth.RecoverVerifyResponseT>) {
          mutation.mutate(data, buildCallbacksTyped(queryClient, rpcAuth.RecoverVerifyResponse, handlers));
        }

        return [action, { loading: mutation.isLoading }];
      },
      useRecoveryUpdate(): RpcMutation<rpcAuth.RecoverUpdateRequestT, rpcAuth.RecoverUpdateResponseT> {
        const mutation = useMutation(
          (data: rpcAuth.RecoverUpdateRequestT) =>
            client.POST<rpcAuth.RecoverUpdateResponseT>("/v1/auth/recover_update", data as Json),
          {},
        );

        // function action(data: RecoverVerifyParams, handlers?: RpcOptions<z.infer<typeof AuthOkResponse>>) {
        function action(data: rpcAuth.RecoverUpdateRequestT, handlers?: RpcOptions<rpcAuth.RecoverUpdateResponseT>) {
          mutation.mutate(
            data,
            buildCallbacksTyped(
              queryClient,
              rpcAuth.RecoverUpdateResponse,
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

      useLogout() {
        return (options?: RpcOptions<void>) => {
          setToken(null);

          options?.onCompleted?.(undefined, undefined);
        };
      },

      useOauthRedirect(): RpcMutation<rpcAuth.OauthUrlRequestT, rpcAuth.OauthUrlResponseT> {
        const mutation = useMutation((data: rpcAuth.OauthUrlRequestT) =>
          client.POST<rpcAuth.OauthUrlResponseT>("/v1/auth/oauth_url", data),
        );

        // function action(data: RecoverVerifyParams, handlers?: RpcOptions<z.infer<typeof AuthOkResponse>>) {
        function action(data: rpcAuth.OauthUrlRequestT, handlers?: RpcOptions<rpcAuth.OauthUrlResponseT>) {
          mutation.mutate(data, buildCallbacksTyped(queryClient, rpcAuth.OauthUrlResponse, handlers));
        }

        return [action, { loading: mutation.isLoading }];
      },

      useOauthLogin(): RpcMutation<rpcAuth.OauthLoginRequestT, rpcAuth.OauthLoginResponseT> {
        const mutation = useMutation(
          (data: rpcAuth.OauthLoginRequestT) => client.POST<rpcAuth.OauthLoginResponseT>("/v1/auth/oauth_login", data),
          {},
        );

        // function action(data: RecoverVerifyParams, handlers?: RpcOptions<z.infer<typeof AuthOkResponse>>) {
        function action(data: rpcAuth.OauthLoginRequestT, handlers?: RpcOptions<rpcAuth.OauthLoginResponseT>) {
          mutation.mutate(
            data,
            buildCallbacksTyped(
              queryClient,
              rpcAuth.OauthLoginResponse,
              {
                ...(handlers ?? {}),
                onCompleted(result, vars) {
                  setToken(result.token.access_token);
                  handlers?.onCompleted?.(result, vars);
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
