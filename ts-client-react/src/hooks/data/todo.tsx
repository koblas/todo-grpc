import { QueryClient, useMutation, useQuery, useQueryClient } from "react-query";
import { useEffect } from "react";
import { z } from "zod";
import { TodoAdd, TodoAddType, TodoItemType, TodoList, TodoListType } from "../../rpc/todo";
import { useAuth } from "../auth";
import { newFetchClient } from "../../rpc/utils";
import { Json } from "../../types/json";
import { useWebsocketUpdates } from "../../rpc/websocket";
import { buildCallbacksTyped } from "../../rpc/utils/helper";

type AddTodoParam = Pick<TodoItemType, "task">;
type DeleteTodoParam = Pick<TodoItemType, "id">;

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

  const result = useQuery("todos", () => client.POST("/v1/todo/todo_list", {}), {
    staleTime: 300_000,
    enabled: !!token,
  });

  const addTodo = useMutation<TodoAddType, unknown, AddTodoParam, unknown>(
    "todos",
    (data) => client.POST<TodoAddType>("/v1/todo/todo_add", data as unknown as Json),
    buildCallbacksTyped(queryClient, TodoAdd, {
      onCompleted(data) {
        cacheAddTodo(queryClient, data.todo);
      },
    }),
  );

  const deleteTodo = useMutation<Json, unknown, DeleteTodoParam, unknown>(
    "todos",
    (data) => client.POST<Json>("/v1/todo/todo_delete", data),
    buildCallbacksTyped<z.ZodUnknown, unknown, unknown, DeleteTodoParam>(queryClient, z.unknown(), {
      onCompleted(data, variables) {
        cacheDeleteTodo(queryClient, variables.id);
      },
    }),
  );

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
      if (event.action === "delete") {
        cacheDeleteTodo(queryClient, event.object_id);
      } else if (event.action === "create" && event.body !== null) {
        cacheAddTodo(queryClient, event.body);
      } else {
        console.log("UNKNOWN TODO EVENT", event);
      }
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
}
