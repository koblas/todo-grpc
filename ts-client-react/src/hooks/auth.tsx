import React, { createContext, useContext, PropsWithChildren } from "react";
import { useImmerReducer } from "use-immer";
import { Draft } from "immer";
import { newAuthClient } from "../rpc/auth/factory";
import { storageFactory } from "../util/storageFactory";

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

export function useAuth() {
  const { state, dispatch } = useContext(AuthContext);

  return {
    token: state.token,
    isAuthenticated: !!state.token,
    async login(username: string, password: string) {
      const { token } = await authClient.login(username, password);
      dispatch({ type: "set", token });
    },
    async logout() {
      dispatch({ type: "set", token: null });
    },
  };
}
