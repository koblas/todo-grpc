import { QueryClient, useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useEffect, useMemo } from "react";
import * as rpcMessage from "../../rpc/message";
import { useAuth } from "../auth";
import { newFetchClient } from "../../rpc/utils";
import { useWebsocketUpdates } from "../../rpc/websocket";
import { buildCallbacksTyped, buildCallbacksTypedQuery } from "../../rpc/utils/helper";
import { RpcMutation, RpcOptions } from "../../rpc/errors";
import { sharedEventCrud } from "./shared";
import { must } from "../../util/assert";

const CACHE_KEY = ["messages"];

function handleEvent(queryClient: QueryClient, event: Omit<rpcMessage.MessageEventT, "topic">) {
  const key = event.body?.room_id ? [...CACHE_KEY, event.body.room_id] : CACHE_KEY;

  queryClient.setQueriesData<rpcMessage.MsgListResponseT>(key, (old) => ({
    messages: sharedEventCrud(old?.messages, event),
  }));
}

export function useRoomsList() {
  const { token } = useAuth();
  const queryClient = useQueryClient();
  const client = newFetchClient({ token });

  const result = useQuery<rpcMessage.RoomListResponseT>(CACHE_KEY, () => client.POST("/v1/message/room_list", {}), {
    staleTime: Infinity,
    enabled: !!token,
    suspense: true,
    ...buildCallbacksTypedQuery(queryClient, rpcMessage.RoomListResponse, {}),
  });

  return must(result?.data?.rooms);
}

export function useMessageList(roomId: string) {
  const { token } = useAuth();
  const queryClient = useQueryClient();
  const client = newFetchClient({ token });

  const result = useQuery<rpcMessage.MsgListResponseT>(
    [...CACHE_KEY, roomId],
    () =>
      client.POST<rpcMessage.MsgListResponseT>("/v1/message/msg_list", {
        room_id: roomId,
      }),
    {
      staleTime: Infinity,
      enabled: !!token,
      suspense: true,
      ...buildCallbacksTypedQuery(queryClient, rpcMessage.MsgListResponse, {}),
    },
  );

  return must(result.data?.messages);
}

export function useMessageMutations() {
  const { token } = useAuth();
  const queryClient = useQueryClient();

  const mutations = useMemo(() => {
    const client = newFetchClient({ token });

    return {
      useRoomJoin(): RpcMutation<rpcMessage.RoomJoinRequestT, rpcMessage.RoomJoinResponseT> {
        const mutation = useMutation(
          CACHE_KEY,
          (data: rpcMessage.RoomJoinRequestT) =>
            client.POST<rpcMessage.RoomJoinResponseT>("/v1/message/room_join", data),
          {},
        );

        function action(data: rpcMessage.RoomJoinRequestT, handlers?: RpcOptions<rpcMessage.RoomJoinResponseT>) {
          mutation.mutate(data, buildCallbacksTyped(queryClient, rpcMessage.RoomJoinResponse, handlers));
        }

        return [action, { loading: mutation.isLoading }];
      },

      useAddMessage(): RpcMutation<rpcMessage.MsgCreateRequestT, rpcMessage.MsgCreateResponseT> {
        const mutation = useMutation(
          CACHE_KEY,
          (data: rpcMessage.MsgCreateRequestT) =>
            client.POST<rpcMessage.MsgCreateResponseT>("/v1/message/msg_create", data),
          {},
        );

        function action(data: rpcMessage.MsgCreateRequestT, handlers?: RpcOptions<rpcMessage.MsgCreateResponseT>) {
          mutation.mutate(
            data,
            buildCallbacksTyped(queryClient, rpcMessage.MsgCreateResponse, {
              ...handlers,
              onCompleted(payload, variables) {
                handleEvent(queryClient, {
                  action: "create",
                  object_id: payload.message.id,
                  body: payload.message,
                });
                handlers?.onCompleted?.(payload, variables);
              },
            }),
          );
        }

        return [action, { loading: mutation.isLoading }];
      },
    };
  }, [token]);

  return mutations;
}

export function useMessageListener() {
  const { addListener } = useWebsocketUpdates();
  const queryClient = useQueryClient();

  useEffect(
    () =>
      addListener("message", (event: rpcMessage.MessageEventT) => {
        const data = rpcMessage.MessageEvent.parse(event);

        handleEvent(queryClient, data);
      }),
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [queryClient],
  );
}
