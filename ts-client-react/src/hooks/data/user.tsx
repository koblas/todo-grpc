import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useEffect } from "react";
import { z } from "zod";
import * as rpcUser from "../../rpc/user";
import { useAuth } from "../auth";
import { newFetchClient } from "../../rpc/utils";
import { useWebsocketUpdates } from "../../rpc/websocket";
import { buildCallbacksTyped, buildCallbacksTypedQuery } from "../../rpc/utils/helper";
import { RpcMutation, RpcOptions } from "../../rpc/errors";

const CACHE_KEY = ["user"];

export function useUser(): rpcUser.UserT {
  const { token } = useAuth();
  const client = newFetchClient({ token });
  const queryClient = useQueryClient();

  const result = useQuery(CACHE_KEY, () => client.POST<rpcUser.GetUserResponseT>("/v1/user/get_user", {}), {
    staleTime: Infinity,
    suspense: true,
    enabled: true,
    ...buildCallbacksTypedQuery(queryClient, rpcUser.GetUserResponse, {}),
  });

  const parsed = rpcUser.GetUserResponse.safeParse(result.data);

  if (!parsed.success) {
    throw new Error("Unable to parse user");
  }

  return parsed.data.user;
}

export function useUserMutations() {
  const { token } = useAuth();
  const client = newFetchClient({ token });
  const queryClient = useQueryClient();

  return {
    useUpdateUser(): RpcMutation<rpcUser.UpdateUserRequestT, rpcUser.UpdateUserResponseT> {
      const mutation = useMutation(
        CACHE_KEY,
        (data: rpcUser.UpdateUserRequestT) => client.POST<rpcUser.UpdateUserResponseT>("/v1/user/update_user", data),
        {},
      );

      // function action(data: RecoverVerifyParams, handlers?: RpcOptions<z.infer<typeof AuthOkResponse>>) {
      function action(data: rpcUser.UpdateUserRequestT, handlers?: RpcOptions<rpcUser.UpdateUserResponseT>) {
        mutation.mutate(data, buildCallbacksTyped(queryClient, rpcUser.UpdateUserResponse, handlers));
      }

      return [action, { loading: mutation.isLoading }];
    },
  };
}

export function useUserListener() {
  const queryClient = useQueryClient();
  const { addListener } = useWebsocketUpdates();

  useEffect(
    () =>
      addListener("user", (event: rpcUser.UserEventT) => {
        const data = rpcUser.UserEvent.parse(event);

        if (data.body) {
          queryClient.setQueriesData<rpcUser.GetUserResponseT>(CACHE_KEY, { user: data.body });
        }
      }),
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [],
  );
}
