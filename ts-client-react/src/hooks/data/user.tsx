import { QueryClient, useMutation, useQuery, useQueryClient } from "react-query";
import { useEffect } from "react";
import { z } from "zod";
import * as rpcUser from "../../rpc/user";
import { useAuth } from "../auth";
import { newFetchClient } from "../../rpc/utils";
import { useWebsocketUpdates } from "../../rpc/websocket";
import { buildCallbacksTyped, buildCallbacksTypedQuery } from "../../rpc/utils/helper";
import { RpcMutation, RpcOptions } from "../../rpc/errors";

export function useUser() {
  const { token } = useAuth();
  const client = newFetchClient({ token });
  const queryClient = useQueryClient();

  const result = useQuery("user", () => client.POST<rpcUser.GetUserResponseT>("/v1/user/get_user", {}), {
    staleTime: 300_000,
    ...buildCallbacksTypedQuery(queryClient, rpcUser.GetUserResponse, {}),
  });

  // const updateUser = useMutation(
  //   "user",
  //   (data: rpcUser.UpdateUserRequestT) => client.POST<rpcUser.UpdateUserResponseT>("/v1/user/update_user", data),

  //   function action(data: rpcUser.UpdateUserRequestT, handlers?: RpcOptions<rpcUser.UpdateUserResponseT>) {
  //     updateUser.mutate(data, buildCallbacksTyped(queryClient, rpcAuth.UpdateUserResponse, handlers));
  //   }

  //   buildCallbacksTyped(queryClient, rpcUser.UpdateUserResponse, {
  //     onCompleted(data) {
  //       queryClient.setQueryData("user", data);
  //     },
  //   }),
  // );

  const parsed = rpcUser.GetUserResponse.safeParse(result.data);

  return {
    user: parsed.success ? parsed.data.user : null,
    isLoading: result.isLoading,
    isError: result.isError,
    mutations: {
      useUpdateUser(): RpcMutation<rpcUser.UpdateUserRequestT, rpcUser.UpdateUserResponseT> {
        const mutation = useMutation(
          "user",
          (data: rpcUser.UpdateUserRequestT) => client.POST<rpcUser.UpdateUserResponseT>("/v1/user/update_user", data),
          {},
        );

        // function action(data: RecoverVerifyParams, handlers?: RpcOptions<z.infer<typeof AuthOkResponse>>) {
        function action(data: rpcUser.UpdateUserRequestT, handlers?: RpcOptions<rpcUser.UpdateUserResponseT>) {
          mutation.mutate(data, buildCallbacksTyped(queryClient, rpcUser.UpdateUserResponse, handlers));
        }

        return [action, { loading: mutation.isLoading }];
      },
    },
  };
}

const UserEvent = z.object({
  object_id: z.string(),
  action: z.enum(["delete", "create", "update"]),
  topic: z.literal("user"),
  body: z.nullable(rpcUser.User),
});

export function useUserListener(queryClient: QueryClient) {
  const { addListener } = useWebsocketUpdates();

  useEffect(
    () =>
      addListener("user", (event: z.infer<typeof UserEvent>) => {
        queryClient.setQueriesData("user", { user: event.body });
      }),
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [],
  );
}
