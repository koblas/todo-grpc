import { QueryClient, useMutation, useQueryClient } from "react-query";
import { useEffect, useMemo } from "react";
import { z } from "zod";
import { create } from "zustand";
import * as rpcMessage from "../../rpc/message";
import { useAuth } from "../auth";
import { newFetchClient } from "../../rpc/utils";
import { useWebsocketUpdates } from "../../rpc/websocket";
import { buildCallbacksTyped } from "../../rpc/utils/helper";
import { RpcMutation, RpcOptions } from "../../rpc/errors";

const CACHE_KEY = "messages";

const useMessageStore = create<{
  messages: rpcMessage.MessageItemT[];
  add(item: rpcMessage.MessageItemT): void;
  update(item: rpcMessage.MessageItemT): void;
  delete(id: rpcMessage.MessageItemT["id"]): void;
}>((set) => ({
  messages: [],
  add: (item: rpcMessage.MessageItemT) =>
    set((state) => {
      if (state.messages.some(({ id }) => id === item.id)) {
        return {};
      }

      return { messages: state.messages.concat(item) };
    }),
  update: (item: rpcMessage.MessageItemT) =>
    set((state) => ({
      messages: state.messages.map((v) => (v.id === item.id ? item : v)),
    })),
  delete: (id: rpcMessage.MessageItemT["id"]) =>
    set((state) => ({
      messages: state.messages.filter((v) => v.id !== id),
    })),
}));

export function useMessages() {
  const { token } = useAuth();
  const queryClient = useQueryClient();
  const messageStore = useMessageStore();

  const mutations = useMemo(() => {
    const client = newFetchClient({ token });

    return {
      useListMessages(): RpcMutation<rpcMessage.ListRequestT, rpcMessage.ListResponseT> {
        const mutation = useMutation(
          CACHE_KEY,
          (data: rpcMessage.ListRequestT) => client.POST<rpcMessage.ListResponseT>("/v1/message/list", data),
          {},
        );

        function action(data: rpcMessage.ListRequestT, handlers?: RpcOptions<rpcMessage.ListResponseT>) {
          mutation.mutate(
            data,
            buildCallbacksTyped(queryClient, rpcMessage.ListResponse, {
              ...handlers,
              onCompleted(payload, variables) {
                payload.messages.forEach((item) => messageStore.add(item));
                handlers?.onCompleted?.(payload, variables);
              },
            }),
          );
        }

        return [action, { loading: mutation.isLoading }];
      },

      useAddMessage(): RpcMutation<rpcMessage.AddRequestT, rpcMessage.AddResponseT> {
        const mutation = useMutation(
          CACHE_KEY,
          (data: rpcMessage.AddRequestT) => client.POST<rpcMessage.AddResponseT>("/v1/message/add", data),
          {},
        );

        // function action(data: RecoverVerifyParams, handlers?: RpcOptions<z.infer<typeof AuthOkResponse>>) {
        function action(data: rpcMessage.AddRequestT, handlers?: RpcOptions<rpcMessage.AddResponseT>) {
          mutation.mutate(
            data,
            buildCallbacksTyped(queryClient, rpcMessage.AddResponse, {
              ...handlers,
              onCompleted(payload, variables) {
                messageStore.add(payload.message);
                handlers?.onCompleted?.(payload, variables);
              },
            }),
          );
        }

        return [action, { loading: mutation.isLoading }];
      },
    };
  }, [messageStore, queryClient, token]);

  return {
    messages: messageStore.messages,
    mutations,
  };
}

export function useMessageListener(queryClient: QueryClient) {
  const { addListener } = useWebsocketUpdates();
  const store = useMessageStore();

  useEffect(
    () =>
      addListener("message", (event: z.infer<typeof rpcMessage.MessageEvent>) => {
        if (event.action === "delete") {
          store.delete(event.object_id);
        } else if (event.action === "update") {
          store.update(event.body);
        } else if (event.action === "create" && event.body !== null) {
          store.add(event.body);
        } else {
          console.log("UNKNOWN TODO EVENT", event);
        }
      }),
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [queryClient],
  );
}