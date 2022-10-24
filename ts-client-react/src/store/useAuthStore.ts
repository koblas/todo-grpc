import create from "zustand";
import { storageFactory } from "../util/storageFactory";

function newTokenStore() {
  const TOKEN = "auth-token";

  const tokenStore = storageFactory(() => localStorage);

  function get(): string | null {
    return tokenStore.getItem(TOKEN) ?? null;
  }

  return {
    get,
    clear(): void {
      tokenStore.clear();
    },
    set(value?: string | null): boolean {
      if (value === get()) {
        return false;
      }
      if (value === undefined || value === null) {
        tokenStore.removeItem(TOKEN);
      } else {
        tokenStore.setItem(TOKEN, value);
      }
      return true;
    },
  };
}

const tokenStore = newTokenStore();

export interface AuthState {
  readonly token: string | null;

  setToken(token: string | null): void;
}

export const useAuthStore = create<AuthState>((set) => ({
  token: tokenStore.get(),
  setToken: (token: string | null) => {
    if (tokenStore.set(token)) {
      set((state) => ({ ...state, token }));
    }
  },
}));
