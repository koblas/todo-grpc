import { useMutation, useQuery, useQueryClient } from "react-query";
import { User } from "../../rpc/user";
import { useAuth } from "../auth";
import { newFetchClient } from "../../rpc/utils";
import { Json } from "../../types/json";

export function useUser() {
  const { token } = useAuth();
  const client = newFetchClient({ token });
  const queryClient = useQueryClient();

  const result = useQuery("user", () => client.POST("/v1/user/get_user", {}), {
    staleTime: 300_000,
  });

  const updateUser = useMutation<
    Json,
    unknown,
    {
      email?: string;
      name?: string;
    },
    unknown
  >("user", (data) => client.POST("/v1/user/update_user", data as unknown as Json), {
    onSuccess(data) {
      queryClient.setQueryData("user", data);
    },
  });

  return {
    user: (result?.data?.user as User) ?? null,
    isLoading: result.isLoading,
    isError: result.isError,
    mutations: {
      updateUser,
    },
  };
}
