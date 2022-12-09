import { QueryClient, useMutation, useQuery, useQueryClient } from "react-query";
import { useEffect } from "react";
import { z } from "zod";
import { TodoItem, TodoItemType, TodoList, TodoListType } from "../../rpc/todo";
import { useAuth } from "../auth";
import { newFetchClient } from "../../rpc/utils";
import { Json } from "../../types/json";
import { useWebsocketUpdates } from "../../rpc/websocket";

function cacheAddTodo(queryClient: QueryClient, item: TodoItemType) {
  queryClient.setQueriesData("todos", (old: TodoListType | undefined) => {
    if (old?.todos?.find(({ id }) => id === item.id)) {
      return old;
    }

    const updated = (old?.todos ?? []).concat([item]);

    return { todos: updated };
  });
}

function cacheDeleteTodo(queryClient: QueryClient, id: TodoItemType["id"]) {
  queryClient.setQueriesData("todos", (old: TodoListType | undefined) => {
    const updated = (old?.todos ?? []).filter(({ id: value }) => value !== id);

    return { todos: updated };
  });
}

export function useTodos() {
  const { token } = useAuth();
  const client = newFetchClient({ token });
  const queryClient = useQueryClient();

  const result = useQuery(
    "todos",
    () =>
      client.POST("/v1/todo/get_todos", {
        enabled: !!token,
      }),
    {
      staleTime: 300_000,
    },
  );

  const addTodo = useMutation<
    Json,
    unknown,
    {
      task?: string;
    },
    unknown
  >("todos", (data) => client.POST("/v1/todo/add_todo", data as unknown as Json), {
    onSuccess(data) {
      const parsed = TodoItem.safeParse(data);

      if (parsed.success) {
        cacheAddTodo(queryClient, parsed.data);
      }
    },
  });

  const deleteTodo = useMutation<
    Json,
    unknown,
    {
      id: string;
    },
    unknown
  >("todos", (data) => client.POST("/v1/todo/delete_todo", data as unknown as Json), {
    onSuccess(data, variables) {
      const parsed = z.object({}).safeParse(data);

      if (parsed.success) {
        cacheDeleteTodo(queryClient, variables.id);
      }
    },
  });

  const parsed = TodoList.safeParse(result.data);

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

  useEffect(() => {
    addListener("todo", (event: z.infer<typeof TodoEvent>) => {
      // console.log("GOT TODO EVENT", event);

      if (event.action === "delete") {
        cacheDeleteTodo(queryClient, event.object_id);
      } else if (event.action === "create" && event.body !== null) {
        cacheAddTodo(queryClient, event.body);
      } else {
        console.log("UNKNOWN TODO EVENT", event);
      }
    });
  }, []);
}