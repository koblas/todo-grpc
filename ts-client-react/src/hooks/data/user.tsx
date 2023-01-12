import { QueryClient, useMutation, useQuery, useQueryClient } from "react-query";
import { useEffect } from "react";
import { z } from "zod";
import { User, UserType } from "../../rpc/user";
import { useAuth } from "../auth";
import { newFetchClient } from "../../rpc/utils";
import { Json } from "../../types/json";
import { useWebsocketUpdates } from "../../rpc/websocket";
import { buildCallbacksTyped, buildCallbacksTypedQuery } from "../../rpc/utils/helper";

const UserNetwork = z.object({
  user: User,
});

type UpdateUserParam = Pick<UserType, "email" | "name"> & {
  password?: string;
  passwordNew: string;
};

export function useUser() {
  const { token } = useAuth();
  const client = newFetchClient({ token });
  const queryClient = useQueryClient();

  const result = useQuery("user", () => client.POST("/v1/user/get_user", {}), {
    staleTime: 300_000,
    ...buildCallbacksTypedQuery(queryClient, UserNetwork, {}),
  });

  const updateUser = useMutation<Json, unknown, UpdateUserParam, unknown>(
    "user",
    (data) => client.POST("/v1/user/update_user", data as unknown as Json),
    buildCallbacksTyped(queryClient, UserNetwork, {
      onCompleted(data) {
        queryClient.setQueryData("user", data);
      },
    }),
  );

  const parsed = UserNetwork.safeParse(result.data);

  return {
    user: parsed.success ? parsed.data.user : null,
    isLoading: result.isLoading,
    isError: result.isError,
    mutations: {
      updateUser,
    },
  };
}

const UserEvent = z.object({
  object_id: z.string(),
  action: z.enum(["delete", "create", "update"]),
  topic: z.literal("user"),
  body: z.nullable(User),
});

export function useUserListener(queryClient: QueryClient) {
  const { addListener } = useWebsocketUpdates();

  useEffect(() => {
    addListener("user", (event: z.infer<typeof UserEvent>) => {
      queryClient.setQueriesData("user", { user: event.body });
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
}
