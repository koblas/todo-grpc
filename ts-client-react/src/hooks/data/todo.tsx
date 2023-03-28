import { QueryClient, useMutation, useQuery, useQueryClient } from "react-query";
import { useEffect } from "react";
import { z } from "zod";
import * as rpcTodo from "../../rpc/todo";
import { useAuth } from "../auth";
import { newFetchClient } from "../../rpc/utils";
import { useWebsocketUpdates } from "../../rpc/websocket";
import { buildCallbacksTyped, buildCallbacksTypedQuery } from "../../rpc/utils/helper";

function cacheAddTodo(queryClient: QueryClient, item: rpcTodo.TodoObjectT) {
  queryClient.setQueriesData<{ todos: rpcTodo.TodoObjectT[] }>("todos", (old) => {
    if (!old) {
      return { todos: [item] };
    }

    if (old.todos.some(({ id }) => id === item.id)) {
      return old;
    }

    return { todos: old.todos.concat(item) };
  });
}

function cacheDeleteTodo(queryClient: QueryClient, id: rpcTodo.TodoObjectT["id"]) {
  queryClient.setQueriesData<{ todos: rpcTodo.TodoObjectT[] }>("todos", (old) => {
    if (!old) {
      return { todos: [] };
    }

    return { todos: old.todos.filter((item) => item.id !== id) };
  });
}

export function useTodos() {
  const { token } = useAuth();
  const client = newFetchClient({ token });
  const queryClient = useQueryClient();

  const result = useQuery("todos", () => client.POST<rpcTodo.TodoListResponseT>("/v1/todo/todo_list", {}), {
    staleTime: 300_000,
    enabled: !!token,
    ...buildCallbacksTypedQuery(queryClient, rpcTodo.TodoListRequest, {}),
  });

  const addTodo = useMutation(
    "todos",
    (data: rpcTodo.TodoAddRequestT) => client.POST<rpcTodo.TodoAddResponseT>("/v1/todo/todo_add", data),
    buildCallbacksTyped(queryClient, rpcTodo.TodoAddResponse, {
      onCompleted(data) {
        cacheAddTodo(queryClient, data.todo);
      },
    }),
  );

  const deleteTodo = useMutation(
    "todos",
    (data: rpcTodo.TodoDeleteRequestT) => client.POST<rpcTodo.TodoDeleteResponseT>("/v1/todo/todo_delete", data),
    buildCallbacksTyped(queryClient, rpcTodo.TodoDeleteResponse, {
      onCompleted(data, variables) {
        cacheDeleteTodo(queryClient, variables.id);
      },
    }),
  );

  const parsed = rpcTodo.TodoListResponse.safeParse(result.data);

  return {
    todos: parsed.success ? parsed.data.todos : null,
    isLoading: result.isLoading,
    isError: result.isError,
    mutations: {
      addTodo,
      deleteTodo,
    },
  };
}

const TodoEvent = z.object({
  object_id: z.string(),
  action: z.enum(["delete", "create"]),
  topic: z.literal("todo"),
  body: z.nullable(
    z.object({
      id: z.string(),
      task: z.string(),
    }),
  ),
});

export function useTodoListener(queryClient: QueryClient) {
  const { addListener } = useWebsocketUpdates();

  useEffect(
    () =>
      addListener("todo", (event: z.infer<typeof TodoEvent>) => {
        if (event.action === "delete") {
          cacheDeleteTodo(queryClient, event.object_id);
        } else if (event.action === "create" && event.body !== null) {
          cacheAddTodo(queryClient, event.body);
        } else {
          console.log("UNKNOWN TODO EVENT", event);
        }
      }),
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [queryClient],
  );
}
