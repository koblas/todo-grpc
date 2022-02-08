import React, { useCallback, useEffect, useMemo, useRef } from "react";
import useWebSocket, { ReadyState } from "react-use-websocket";
import { v4 as uuidV4 } from "uuid";

import create from "zustand";
import createContext from "zustand/context";
import { WebSocketHook } from "react-use-websocket/dist/lib/types";
import { useAuth } from "../hooks/auth";
import { Json } from "../types/json";

const WS_URL = process.env.WS_URL ?? "";

type ListenerFunc = (event: Json) => void;
type ListenerSelector = { topic: string | null; handler: <E = Json>(event: E) => void };

type BearState = {
  connectionId: string;
  socket: null | WebSocketHook;
  connected: boolean;
  listeners: ListenerSelector[];
  setSocket: (socket: null | WebSocketHook) => void;
  setConnected: (connected: boolean) => void;
  addListener: (listener: ListenerSelector) => void;
};

// const { Provider, useStore } = createContext<BearState>();

const createStore = () =>
  create<BearState>((set) => ({
    connectionId: uuidV4(),
    socket: null,
    connected: false,
    listeners: [],
    setSocket: (s: null | WebSocketHook) => set({ socket: s }),
    setConnected: (connected: boolean) => set({ connected }),
    addListener: (listener: ListenerSelector) => set((state) => ({ listeners: state.listeners.concat(listener) })),
  }));

const useStore = createStore();

export function WebsocketProvider({ children }: React.PropsWithChildren<unknown>) {
  const { token } = useAuth();
  const pos = useRef<{ value: string | null }>({ value: null });
  const store = useStore();

  const { lastJsonMessage, sendJsonMessage, readyState } = useWebSocket(
    WS_URL,
    {
      queryParams: { t: token ?? "" },
      // shouldReconnect: (e) => {
      //   console.log("SHOULD RECONNECT", token, pos.current, e);
      //   /* some comment */
      //   if (!token) {
      //     return false;
      //   }
      //   return true;
      // },
      reconnectAttempts: 10,
      reconnectInterval: 10000,
    },
    !!token,
  );

  useEffect(() => {
    sendJsonMessage({ cursor: pos.current.value });
  }, [pos.current.value, sendJsonMessage]);

  useEffect(() => {
    if (!lastJsonMessage) {
      return;
    }
    const { topic } = lastJsonMessage;

    // store.listeners.console.log(lastJsonMessage);
    store.listeners.forEach((item) => {
      if (!item.topic || item.topic === topic) {
        item.handler(lastJsonMessage);
      }
    });
  }, [lastJsonMessage]);

  useEffect(() => {
    store.setConnected(readyState === ReadyState.OPEN);
  }, [readyState]);

  // return <Provider createStore={createStore}>{children}</Provider>;
  return children;
}

export function useWebsocketUpdates() {
  const store = useStore();

  // console.log("STORE = ", store);

  const addListener = useCallback(
    (topic: string | null, handler: ListenerFunc) => {
      store.addListener({ topic, handler });
    },
    [store.addListener],
  );

  return {
    socket: store.socket,
    connnectionId: store.connectionId,
    addListener,
  };
}
