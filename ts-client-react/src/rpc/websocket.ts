import React, { useCallback, useEffect, useRef, useState } from "react";
import useWebSocket, { ReadyState } from "react-use-websocket";
import { v4 as uuidV4 } from "uuid";

import create from "zustand";
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

export function WebsocketProvider({ children }: { children: JSX.Element }): JSX.Element | null {
  const { token } = useAuth();
  const [pingTimer, setPingTimer] = useState<null | NodeJS.Timer>(null);
  const [heartbeatTimer, setHeartbeatTimer] = useState<null | NodeJS.Timer>(null);
  const pos = useRef<{ value: string | null }>({ value: null });
  const store = useStore();

  const { lastJsonMessage, sendJsonMessage, readyState } = useWebSocket(
    WS_URL,
    {
      queryParams: { t: token ?? "" },
      shouldReconnect: () => !!token,
      reconnectAttempts: 10,
      reconnectInterval: 10000,
    },
    !!token,
  );

  // console.log("HERE!", readyState);

  useEffect(() => {
    if (readyState === ReadyState.OPEN) {
      sendJsonMessage({ action: "cursor", position: pos.current.value });
    }
  }, [pos.current.value, sendJsonMessage, readyState]);

  useEffect(() => {
    function clearTimers() {
      if (pingTimer !== null) {
        clearInterval(pingTimer);
        setPingTimer(null);
      }
      if (heartbeatTimer !== null) {
        clearInterval(heartbeatTimer);
        setHeartbeatTimer(null);
      }
    }

    if (readyState === ReadyState.OPEN) {
      // The ping timer fires to keep the socket open, to prevent
      //  VPN gateways from dropping the connection
      // -- Every 45seconds
      if (pingTimer === null) {
        setPingTimer(
          setInterval(() => {
            sendJsonMessage({ action: "ping" });
          }, 45_000),
        );
      }
      // The heartbeat lets AWS know that the socket is still
      //  connected and to refresh the connection expiration
      // -- Every 45minutes
      if (heartbeatTimer === null) {
        setHeartbeatTimer(
          setInterval(() => {
            sendJsonMessage({ action: "heartbeat" });
          }, 60 * 45_000),
        );
      }
    } else {
      clearTimers();
    }

    return clearTimers;
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [readyState]);

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
