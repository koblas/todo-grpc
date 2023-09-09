import { QueryClient, useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useEffect } from "react";
import * as rpcTodo from "../../rpc/todo";
import { useAuth } from "../auth";
import { newFetchClient } from "../../rpc/utils";
import { useWebsocketUpdates } from "../../rpc/websocket";
import { buildCallbacksTyped, buildCallbacksTypedQuery } from "../../rpc/utils/helper";
import { sharedEventCrud } from "./shared";

const CACHE_KEY = ["todos"];

function handleEvent(queryClient: QueryClient, event: Omit<rpcTodo.TodoEventT, "topic">) {
  queryClient.setQueriesData<rpcTodo.TodoListResponseT>(CACHE_KEY, (old) => ({
    todos: sharedEventCrud(old?.todos, event),
  }));
}

export function useTodos(fetch: boolean = false) {
  const { token } = useAuth();
  const client = newFetchClient({ token });
  const queryClient = useQueryClient();

  const result = useQuery(CACHE_KEY, () => client.POST<rpcTodo.TodoListResponseT>("/v1/todo/todo_list", {}), {
    staleTime: Infinity,
    enabled: !!token && fetch,
    suspense: fetch,
    ...buildCallbacksTypedQuery(queryClient, rpcTodo.TodoListResponse, {}),
  });

  const addTodo = useMutation(
    ["todos"],
    (data: rpcTodo.TodoAddRequestT) => client.POST<rpcTodo.TodoAddResponseT>("/v1/todo/todo_add", data),
    buildCallbacksTyped(queryClient, rpcTodo.TodoAddResponse, {
      onCompleted(data) {
        handleEvent(queryClient, {
          action: "create",
          object_id: data.todo.id,
          body: data.todo,
        });
      },
    }),
  );

  const deleteTodo = useMutation(
    ["todos"],
    (data: rpcTodo.TodoDeleteRequestT) => client.POST<rpcTodo.TodoDeleteResponseT>("/v1/todo/todo_delete", data),
    buildCallbacksTyped(queryClient, rpcTodo.TodoDeleteResponse, {
      onCompleted(data, variables) {
        handleEvent(queryClient, {
          action: "delete",
          object_id: variables.id,
          body: null,
        });
      },
    }),
  );

  const parsed = rpcTodo.TodoListResponse.safeParse(result.data);

  return {
    todos: parsed.success ? parsed.data.todos : [],
    mutations: {
      addTodo,
      deleteTodo,
    },
  };
}

export function useTodoListener() {
  const { addListener } = useWebsocketUpdates();
  const queryClient = useQueryClient();

  useEffect(
    () =>
      addListener("todo", (event: unknown) => {
        const data = rpcTodo.TodoEvent.parse(event);

        handleEvent(queryClient, data);
      }),
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [queryClient],
  );
}
