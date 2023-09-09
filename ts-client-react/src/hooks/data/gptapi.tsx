import { useMutation, useQueryClient } from "@tanstack/react-query";
import { z } from "zod";
import { useAuth } from "../auth";
import { newFetchClient } from "../../rpc/utils";
import { buildCallbacksTyped } from "../../rpc/utils/helper";
import { RpcMutationNew } from "../../rpc/errors";

const GptApiRequest = z.object({
  text: z.string(),
});
const GptApiResponse = z.object({
  text: z.string(),
});

type GptApiRequestType = z.infer<typeof GptApiRequest>;
type GptApiResponseType = z.infer<typeof GptApiResponse>;

export function useGptApi() {
  const { token } = useAuth();
  const client = newFetchClient({ token });
  const queryClient = useQueryClient();

  return {
    useTextApi(): RpcMutationNew<GptApiRequestType, GptApiResponseType> {
      const mutation = useMutation(
        (data: GptApiRequestType) =>
          client.POST<GptApiResponseType>("/v1/gpt/create", {
            prompt: data.text,
          }),
        {},
      );

      return [
        (data, handlers?) => {
          mutation.mutate(data, buildCallbacksTyped(queryClient, GptApiResponse, handlers));
        },
        mutation,
      ];
    },
  };
}
