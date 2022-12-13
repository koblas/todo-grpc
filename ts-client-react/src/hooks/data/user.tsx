import { useMutation, useQuery, useQueryClient } from "react-query";
import { z } from "zod";
import { User, UserType } from "../../rpc/user";
import { useAuth } from "../auth";
import { newFetchClient } from "../../rpc/utils";
import { Json } from "../../types/json";
import { buildCallbacksTyped } from "../../rpc/utils/helper";

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
