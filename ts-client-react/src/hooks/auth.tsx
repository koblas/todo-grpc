import React, { createContext, useContext, PropsWithChildren, useCallback, useState } from "react";
import { useImmerReducer } from "use-immer";
import { Draft } from "immer";
import { newAuthClient } from "../rpc/auth/factory";
import { storageFactory } from "../util/storageFactory";
import { useNetworkContext, useNetworkContextErrors } from "./network";
import { RpcMutation, RpcOptions, RpcResult } from "../rpc/errors";

/**
 * Construct an accessor for a persistent token store
 */
function newTokenStore() {
  const TOKEN = "auth-token";

  const tokenStore = storageFactory(() => localStorage);

  return {
    get(): string | null {
      return tokenStore.getItem(TOKEN) ?? null;
    },
    clear(): void {
      tokenStore.clear();
    },
    set(value?: string | null): void {
      if (value === undefined || value === null) {
        tokenStore.removeItem(TOKEN);
      } else {
        tokenStore.setItem(TOKEN, value);
      }
    },
  };
}

const tokenStore = newTokenStore();

///
type DispatchAction = {
  type: "set";
  token: string | null;
};

interface AuthState {
  readonly token: string | null;
}

const defaultState: AuthState = {
  token: tokenStore.get(),
};

const AuthContext = createContext<{
  state: AuthState;
  dispatch: React.Dispatch<DispatchAction>;
}>({
  state: defaultState,
  dispatch() {
    throw new Error("not initialized");
  },
});

function authReducer(draft: Draft<AuthState>, action: DispatchAction) {
  switch (action.type) {
    case "set":
      draft.token = action.token;
      tokenStore.set(action.token);
      break;
  }
}

export function AuthContextProvider({ children }: PropsWithChildren<unknown>) {
  // The reason we need to use a reducer here rather that a setState is
  // that when you are doing optimistic updates your not able to get a handle
  // on the "current" set of items in your list.  You might have done 2..3 actions
  // but your local closure will only have the state at the time you dispatched
  // your asyncronist event.  Thus you need to bump the world into a reducer...
  const [state, dispatch] = useImmerReducer(authReducer, defaultState);

  return <AuthContext.Provider value={{ state, dispatch }}>{children}</AuthContext.Provider>;
}

const authClient = newAuthClient("json");

export type LoginParams = {
  email: string;
  password: string;
};
export type RegisterParams = {
  name: string;
  email: string;
  password: string;
};

export function useAuth() {
  const { state, dispatch } = useContext(AuthContext);
  const addHandler = useNetworkContextErrors();

  return {
    token: state.token,
    isAuthenticated: !!state.token,
    mutations: {
      useRegister(): RpcMutation<RegisterParams, string> {
        const [data, setData] = useState<string | undefined>("");
        const [loading, setLoading] = useState(true);
        const [error, setError] = useState<any>(undefined);

        const func = (params: RegisterParams, options?: RpcOptions<string>) => {
          authClient.register(
            params,
            addHandler(
              {
                onCompleted: ({ token }) => {
                  setLoading(false);
                  setData(token);
                  dispatch({ type: "set", token });

                  options?.onCompleted?.(token);
                },
                onError(error: unknown) {
                  setError(error);
                  options?.onError?.(error);
                },
              },
              options,
            ),
          );
        };

        return [func, { data, loading, error }];
      },
      useLogin(): RpcMutation<LoginParams, string> {
        const [data, setData] = useState<string>("");
        const [loading, setLoading] = useState(true);
        const [error, setError] = useState<any>(null);

        const func = (params: LoginParams, options?: RpcOptions<string>) => {
          authClient.authenticate(
            params,
            addHandler(
              {
                onCompleted: ({ token }) => {
                  setLoading(false);
                  setData(token);
                  dispatch({ type: "set", token });

                  options?.onCompleted?.(token);
                },
                onError(error: unknown) {
                  setError(error);
                  options?.onError?.(error);
                },
              },
              options,
            ),
          );
        };

        return [func, { data, loading, error }];
      },
      useLogout() {
        return (options?: RpcOptions<void>) => {
          dispatch({ type: "set", token: null });

          options?.onCompleted?.();
        };
      },
    },
  };
}
