import { z } from "zod";
import { MutateOptions, QueryClient } from "react-query";
import { RpcOptions } from "../errors";
import { handleJsonError } from "./json_helpers";

export function buildCallbacksTyped<
  TZod extends z.ZodType,
  TData = z.infer<TZod>,
  TError = unknown,
  TVariables = void,
  TContext = unknown,
>(
  queryClient: QueryClient,
  tzod: TZod,
  ...handlers: (RpcOptions<TData, TVariables> | undefined)[]
): Pick<MutateOptions<TData, TError, TVariables, TContext>, "onSettled" | "onError" | "onSuccess"> {
  const defaultOptions = queryClient.getDefaultOptions();

  const options = {
    onSuccess(result: TData, variables: TVariables, context: TContext) {
      const parsed = tzod.safeParse(result);

      const found = handlers?.map((handler) => {
        if (!parsed.success) {
          if (handler) {
            handleJsonError(parsed.error, handler);

            return true;
          }
        } else if (handler?.onCompleted) {
          handler.onCompleted(parsed.data, variables);
          return true;
        }
        return false;
      });

      if (found.some((x) => x)) {
        if (!parsed.success) {
          defaultOptions.mutations?.onError?.(parsed.error, variables, context);
        } else {
          defaultOptions.mutations?.onSuccess?.(result, variables, context);
        }
      }
    },
    onSettled(data: TData | undefined, error: TError | null, variables: TVariables, context: TContext | undefined) {
      const found = handlers?.map((handler) => {
        handler?.onFinished?.();

        return Boolean(handler?.onFinished);
      });

      if (found.some((x) => x)) {
        defaultOptions.mutations?.onSettled?.(data, error, variables, context);
      }
    },
    // Handle an error, determine if it's a HTTP error or something else
    onError(error: TError, variables: TVariables, context: TContext | undefined) {
      // onError(error: TError) {
      const found = handlers?.map((handler) => handleJsonError(error, handler));

      if (found.some((x) => x)) {
        defaultOptions.mutations?.onError?.(error, variables, context);
      }
    },
  };

  return options;
}
