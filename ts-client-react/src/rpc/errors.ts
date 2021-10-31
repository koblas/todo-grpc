export class ErrorUnauthenticated extends Error {}

export interface RpcOptions<T> {
  onCompleted?: (data: T) => void;
  onError?: (error: unknown) => void;
  onErrorField?: (error: Record<string, string[]>) => void;
  onErrorAuthentication?: (error: unknown) => void;
  onErrorNetwork?: (error: unknown) => void;
}

export interface RpcResult<T> {
  loading: boolean;
  data?: T;
  error?: {
    fields?: Record<string, string[]>;
    network?: unknown;
    authentication?: unknown;
  };
}

export type RpcMutation<P, R> = [mutator: (params: P, options?: RpcOptions<R>) => void, state: RpcResult<R>];
