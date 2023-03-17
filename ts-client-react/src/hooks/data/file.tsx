import { useMutation, useQueryClient } from "react-query";
import { z } from "zod";
import { useAuth } from "../auth";
import { newFetchClient } from "../../rpc/utils";
import { buildCallbacksTyped } from "../../rpc/utils/helper";
import { RpcMutationNew } from "../../rpc/errors";

type UploadUrlParam = {
  type: string;
  contentType: string;
};

const UploadUrlResponse = z.object({
  url: z.string(),
  id: z.string(),
});

type UploadFileParam = {
  url: string;
  file: File;
};

const UploadFileResponse = z.unknown();
//   id: z.string(),
// });
// const UploadFileResponse = z.object({
//   id: z.string(),
// });

type UploadFile = z.infer<typeof UploadFileResponse>;
type UploadUrl = z.infer<typeof UploadUrlResponse>;

export function useUploadFile() {
  const { token } = useAuth();
  const client = newFetchClient({ token });
  const fileClient = newFetchClient({ base: "" });
  const queryClient = useQueryClient();

  return {
    useUploadUrl(): RpcMutationNew<UploadUrlParam, UploadUrl> {
      const mutation = useMutation(
        (data: UploadUrlParam) =>
          client.POST<UploadUrl>("/v1/file/upload_url", {
            type: data.type,
            contentType: data.contentType,
          }),
        {},
      );

      return [
        (data, handlers?) => {
          mutation.mutate(data, buildCallbacksTyped(queryClient, UploadUrlResponse, handlers));
        },
        mutation,
      ];
    },
    useUploadSend(): RpcMutationNew<UploadFileParam, UploadFile> {
      const mutation = useMutation((data: UploadFileParam) => fileClient.PUT_FILE<UploadFile>(data.url, data.file), {});

      return [
        (data, handlers?) => {
          mutation.mutate(data, buildCallbacksTyped(queryClient, UploadFileResponse, handlers));
        },
        mutation,
      ];
    },
  };
}
