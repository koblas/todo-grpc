import { useCallback, useEffect, useRef, useState } from "react";
import useWebSocket, { ReadyState } from "react-use-websocket";
import { v4 as uuidV4 } from "uuid";

import { create } from "zustand";
import { WebSocketHook } from "react-use-websocket/dist/lib/types";
import { useAuth } from "../hooks/auth";
import { Json } from "../types/json";

// function buildWebsocketUrl(): string {
//   const base = process.env.WS_URL ?? "";

//   if (base === "") {
//     return "";
//   }
//   if (base.startsWith("ws")) {
//     return base;
//   }

//   const loc = window.location;

//   return new URL((loc.protocol === "https:" ? "wss://" : "ws://") + loc.host + loc.pathname, base).toString();
// }

// const WS_URL = buildWebsocketUrl();

type ListenerFunc<E extends Json> = (event: E) => void;
type ListenerSelector<E extends Json> = { topic: string | null; handler: ListenerFunc<E> };

type BearState<T extends Json = Json> = {
  connectionId: string;
  socket: null | WebSocketHook;
  connected: boolean;
  listeners: ListenerSelector<T>[];
  setSocket: (socket: null | WebSocketHook) => void;
  setConnected: (connected: boolean) => void;
  addListener: (listener: ListenerSelector<T>) => void;
  removeListener: (listener: ListenerSelector<T>) => void;
};

// const { Provider, useStore } = createContext<BearState>();

const createStore = <T extends Json>() =>
  create<BearState<T>>((set) => ({
    connectionId: uuidV4(),
    socket: null,
    connected: false,
    listeners: [],
    setSocket: (s: null | WebSocketHook) => set({ socket: s }),
    setConnected: (connected: boolean) => set({ connected }),
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    addListener: (listener) => set((state) => ({ listeners: state.listeners.concat(listener) })),
    removeListener: (listener) =>
      set((state) => ({
        listeners: state.listeners.filter((item) => item.topic !== listener.topic || item.handler !== listener.handler),
      })),
  }));

const useStore = createStore();

export function WebsocketProvider({ children, url }: { url: string; children: JSX.Element }): JSX.Element | null {
  const { token } = useAuth();
  const [pingTimer, setPingTimer] = useState<null | NodeJS.Timer>(null);
  const [heartbeatTimer, setHeartbeatTimer] = useState<null | NodeJS.Timer>(null);
  const pos = useRef<{ value: string | null }>({ value: null });
  const store = useStore();

  const { lastJsonMessage, sendJsonMessage, readyState } = useWebSocket(
    url,
    {
      // queryParams: { t: token ?? "" },
      // queryParams: {},
      shouldReconnect: () => !!token,
      reconnectAttempts: 10,
      reconnectInterval: 10000,
    },
    !!token,
  );

  // console.log("HERE!", readyState);

  useEffect(() => {
    if (readyState === ReadyState.OPEN) {
      sendJsonMessage({ action: "authorization", token });
      sendJsonMessage({ action: "cursor", position: pos.current.value });
    }
  }, [pos.current.value, sendJsonMessage, readyState, token]);

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

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const { topic } = lastJsonMessage as any;

    // store.listeners.console.log(lastJsonMessage);
    store.listeners.forEach((item) => {
      if (!item.topic || item.topic === topic) {
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        item.handler(lastJsonMessage as any);
      }
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [lastJsonMessage]);

  // return <Provider createStore={createStore}>{children}</Provider>;
  return children;
}

export function useWebsocketUpdates() {
  const store = useStore();

  const addListener = <E extends Json>(topic: string | null, handler: ListenerFunc<E>) => {
    store.addListener({ topic, handler: handler as any });

    return () => {
      store.removeListener({ topic, handler: handler as any });
    };
  };

  return {
    socket: store.socket,
    connnectionId: store.connectionId,
    addListener,
  };
}
