import create from "zustand";

import { User } from "../rpc/user";

export interface UserState {
  readonly user: User | null;

  // Actions
  updateUser(user: Partial<User>): void;
}

export const getUserStore = create<UserState>((set) => ({
  user: null,

  updateUser(user: User): void {
    return set((state) => ({ ...state, user }));
  },
}));
