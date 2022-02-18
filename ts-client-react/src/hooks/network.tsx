import React, { createContext, useContext, PropsWithChildren } from "react";
import { useImmerReducer } from "use-immer";
import { Draft } from "immer";
import { RpcOptions, RpcError } from "../rpc/errors";
import { assert } from "../util/assert";

type Handlers = Pick<
  RpcOptions<unknown>,
  "onError" | "onErrorAuthentication" | "onErrorNetwork" | "onErrorField" | "onFinished" | "onFinished" | "onBegin"
>;

enum ActionType {
  SET,
}
///
type DispatchAction = {
  type: ActionType.SET;
  value: Handlers;
};

interface NetworkState extends Handlers {}

const defaultState: NetworkState = {
  onBegin: undefined,
  onFinished: undefined,
  onErrorField: undefined,
  onError: undefined,
  onErrorAuthentication: undefined,
  onErrorNetwork: undefined,
};

const NetworkContext = createContext<{
  state: NetworkState;
  dispatch: React.Dispatch<DispatchAction>;
}>({
  state: defaultState,
  dispatch() {
    throw new Error("not initialized");
  },
});

function authReducer(draft: Draft<NetworkState>, action: DispatchAction) {
  assert(action.type === ActionType.SET);

  draft.onError = action.value.onError;
  draft.onErrorAuthentication = action.value.onErrorAuthentication;
  draft.onErrorNetwork = action.value.onErrorNetwork;
}

export function NetworkContextProvider({ children }: PropsWithChildren<unknown>): JSX.Element {
  // The reason we need to use a reducer here rather that a setState is
  // that when you are doing optimistic updates your not able to get a handle
  // on the "current" set of items in your list.  You might have done 2..3 actions
  // but your local closure will only have the state at the time you dispatched
  // your asyncronist event.  Thus you need to bump the world into a reducer...
  const [state, dispatch] = useImmerReducer(authReducer, defaultState);

  return <NetworkContext.Provider value={{ state, dispatch }}>{children}</NetworkContext.Provider>;
}

export function useNetworkContext() {
  const { state, dispatch } = useContext(NetworkContext);

  return {
    networkErrors: {
      onBegin: state.onBegin,
      onFinished: state.onFinished,
      onError: state.onError,
      onErrorField: state.onErrorField,
      onErrorNetwork: state.onErrorNetwork,
      onErrorAuthentication: state.onErrorAuthentication,
    },
    setHandlers(handlers: Handlers) {
      dispatch({ type: ActionType.SET, value: handlers });
    },
  };
}

export type ErrorHandler<T> = (options: RpcOptions<T>) => RpcOptions<T>;

export function useNetworkContextErrors<T>(parent?: RpcOptions<T>) {
  const { networkErrors } = useNetworkContext();

  function handle<V, U = V>(base: RpcOptions<V>, ...options: (RpcOptions<U> | undefined)[]) {
    return {
      onCompleted(data: V) {
        [base, ...options].reverse().forEach((option) => {
          // U and V should be the same but I can't get the types right
          // eslint-disable-next-line @typescript-eslint/no-explicit-any
          option?.onCompleted?.(data as any);
        });
      },
      onBegin() {
        [networkErrors, parent, base, ...options].reverse().forEach((option) => {
          option?.onBegin?.();
        });
      },
      onFinished() {
        [networkErrors, parent, base, ...options].reverse().forEach((option) => {
          option?.onFinished?.();
        });
      },
      onError(error: RpcError) {
        [networkErrors, parent, base, ...options].reverse().forEach((option) => {
          option?.onError?.(error);
        });
      },
      onErrorField(fields: Record<string, string[]>) {
        [networkErrors, parent, base, ...options].reverse().forEach((option) => {
          option?.onErrorField?.(fields);
        });
      },
      onErrorNetwork(error: RpcError) {
        [networkErrors, parent, base, ...options].reverse().forEach((option) => {
          option?.onErrorNetwork?.(error);
        });
      },
      onErrorAuthentication(error: unknown) {
        [networkErrors, parent, base, ...options].reverse().forEach((option) => {
          option?.onErrorAuthentication?.(error);
        });
      },
    };
  }

  return handle;
}
