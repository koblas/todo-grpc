import React, { createContext, useContext, PropsWithChildren, useCallback, useState } from "react";
import { useImmerReducer } from "use-immer";
import { Draft } from "immer";
import { RpcOptions } from "../rpc/errors";

type Handlers = Pick<RpcOptions<unknown>, "onError" | "onErrorAuthentication" | "onErrorNetwork" | "onErrorField">;

///
type DispatchAction = {
  type: "set";
  value: Handlers;
};

interface NetworkState extends Handlers {}

const defaultState: NetworkState = {
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
  switch (action.type) {
    case "set":
      draft.onError = action.value.onError;
      draft.onErrorAuthentication = action.value.onErrorAuthentication;
      draft.onErrorNetwork = action.value.onErrorNetwork;
      break;
  }
}

export function NetworkContextProvider({ children }: PropsWithChildren<unknown>) {
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
      onError: state.onError,
      onErrorField: state.onErrorField,
      onErrorNetwork: state.onErrorNetwork,
      onErrorAuthentication: state.onErrorAuthentication,
    },
    setHandlers(handlers: Handlers) {
      dispatch({ type: "set", value: handlers });
    },
  };
}

export type ErrorHandler<T> = (options: RpcOptions<T>) => RpcOptions<T>;

export function useNetworkContextErrors<T>(parent?: RpcOptions<T>) {
  const { networkErrors } = useNetworkContext();

  function handle<V, U>(base: RpcOptions<V>, ...options: (RpcOptions<U> | undefined)[]) {
    return {
      onCompleted(data: V) {
        base?.onCompleted?.(data);
      },
      onError(error: unknown) {
        [networkErrors, parent, ...options].forEach((option) => {
          option?.onError?.(error);
        });
      },
      onErrorField(fields: Record<string, string[]>) {
        [networkErrors, parent, ...options].forEach((option) => {
          option?.onErrorField?.(fields);
        });
      },
      onErrorNetwork(error: unknown) {
        [networkErrors, parent, ...options].forEach((option) => {
          option?.onErrorNetwork?.(error);
        });
      },
      onErrorAuthentication(error: unknown) {
        [networkErrors, parent, ...options].forEach((option) => {
          option?.onErrorAuthentication?.(error);
        });
      },
    };
  }

  return handle;
}
