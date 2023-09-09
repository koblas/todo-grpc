import { useEffect } from "react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { z } from "zod";
import * as rpcFile from "../../rpc/file";
import { useAuth } from "../auth";
import { newFetchClient } from "../../rpc/utils";
import { buildCallbacksTyped } from "../../rpc/utils/helper";
import { RpcMutationNew } from "../../rpc/errors";
import { useWebsocketUpdates } from "../../rpc/websocket";

export function useUploadFile() {
  const { token } = useAuth();
  const client = newFetchClient({ token });
  const fileClient = newFetchClient({ base: "" });
  const queryClient = useQueryClient();

  return {
    useUploadUrl(): RpcMutationNew<rpcFile.UploadUrlRequestT, rpcFile.UploadUrlResponseT> {
      const mutation = useMutation(
        (data: rpcFile.UploadUrlRequestT) =>
          client.POST<rpcFile.UploadUrlResponseT>("/v1/file/upload_url", {
            type: data.type,
            contentType: data.contentType,
          }),
        {},
      );

      return [
        (data, handlers?) => {
          mutation.mutate(data, buildCallbacksTyped(queryClient, rpcFile.UploadUrlResponse, handlers));
        },
        mutation,
      ];
    },
    useUploadSend(): RpcMutationNew<rpcFile.UploadFileRequestT, rpcFile.UploadFileResponseT> {
      const mutation = useMutation(
        (data: rpcFile.UploadFileRequestT) => fileClient.PUT_FILE<rpcFile.UploadFileResponseT>(data.url, data.file),
        {},
      );

      return [
        (data, handlers?) => {
          mutation.mutate(data, buildCallbacksTyped(queryClient, rpcFile.UploadFileResponse, handlers));
        },
        mutation,
      ];
    },
  };
}

export function useFileListener(action: (id: string, error?: string) => void) {
  const { addListener } = useWebsocketUpdates();

  useEffect(
    () =>
      addListener("file", (event: z.infer<typeof rpcFile.FileEvent>) => {
        if (event.action === "error") {
          if (event.body && event.body.id) {
            action(event.body.id, event.body?.error);
          }
          console.log("ERROR", event.body?.id, event.body?.error);
        } else if (event.action === "create") {
          if (event.body && event.body.id) {
            action(event.body.id);
          }
          // cacheAddTodo(queryClient, event.body);
          console.log("CREATED", event.body?.id);
        }
      }),
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [],
  );
}
