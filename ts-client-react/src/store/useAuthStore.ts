import create from "zustand";
import { produce } from "immer";
import { storageFactory } from "../util/storageFactory";
import { zimmer } from "../util/zimmer";

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

export interface AuthState {
  readonly token: string | null;

  setToken(token: string | null): void;
}

export const useAuthStore = create(
  zimmer<AuthState>((set) => ({
    token: tokenStore.get(),
    setToken: (token: string | null) => {
      set(
        produce((draft) => {
          tokenStore.set(token);
          draft.token = token;
        }),
      );
    },
  })),
);
