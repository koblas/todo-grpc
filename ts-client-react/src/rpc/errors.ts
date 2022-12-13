/* eslint-disable max-classes-per-file */
export class ErrorUnauthenticated extends Error {}

export interface RpcError {
  fields?: Record<string, string[]>;
  network?: unknown;
  authentication?: unknown;
}

export const RpcStop: unique symbol = Symbol("rpc_stop");

export interface RpcOptions<T, TVar = unknown> {
  /// Called when a request starts
  onBegin?: () => void | typeof RpcStop;
  /// Always called regardless of error or completed
  onFinished?: () => void | typeof RpcStop;
  /// Called whith the successful response payload
  onCompleted?: (data: T, variables: TVar) => void | typeof RpcStop;
  /// Called for any error
  onError?: (error: RpcError) => void | typeof RpcStop;
  /// Called for a field validation error
  onErrorField?: (error: Record<string, string[]>) => void | typeof RpcStop;
  /// Called when authentication returns failure -- in addition to onError
  onErrorAuthentication?: (error: unknown) => void | typeof RpcStop;
  /// Called when there is a network error -- in addition to onError
  onErrorNetwork?: (error: RpcError) => void | typeof RpcStop;
}

export interface RpcResult<T> {
  loading: boolean;
  data?: T;
  error?: RpcError;
}

export type RpcMutation<P, R> = [mutator: (params: P, options?: RpcOptions<R>) => void, state: RpcResult<R>];
