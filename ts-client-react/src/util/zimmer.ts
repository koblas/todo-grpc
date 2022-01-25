import { produce, Draft } from "immer";
import { GetState, SetState, State, StateCreator, StoreApi } from "zustand";

/// Immer reducer for zustand
export const zimmer =
  <
    T extends State,
    CustomSetState extends SetState<T> = SetState<T>,
    CustomGetState extends GetState<T> = GetState<T>,
    CustomStoreApi extends StoreApi<T> = StoreApi<T>,
  >(
    config: StateCreator<
      T,
      (partial: ((draft: Draft<T>) => void) | T, replace?: boolean) => void,
      CustomGetState,
      CustomStoreApi
    >,
  ): StateCreator<T, CustomSetState, CustomGetState, CustomStoreApi> =>
  (set, get, api) =>
    config(
      (partial, replace) => {
        const nextState = typeof partial === "function" ? produce(partial as (state: Draft<T>) => T) : (partial as T);
        return set(nextState, replace);
      },
      get,
      api,
    );
