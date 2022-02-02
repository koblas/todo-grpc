import React, { useEffect, useRef } from "react";

import { ChakraProvider, CSSReset, Flex } from "@chakra-ui/react";
import useWebSocket, { ReadyState } from "react-use-websocket";

import { BrowserRouter, Routes, Route, Outlet } from "react-router-dom";
import { AuthPages } from "./pages/auth";
import { SettingsPage } from "./pages/settings";
import { TodoPage } from "./pages/TodoPage";
import { HomePage } from "./pages/HomePage";
import { NetworkContextProvider } from "./hooks/network";
import Sidebar from "./components/Sidebar";
import { useTodos } from "./hooks/todo";
import { useAuth } from "./hooks/auth";

export const WS_URL = process.env.WS_URL ?? "";

function ReloadState() {
  const { token } = useAuth();
  const pos = useRef<{ value: string | null }>({ value: null });
  const { mutations } = useTodos();
  const [loadTodos] = mutations.useLoadTodos();
  const [clearTodos] = mutations.useClearTodos();
  const { lastJsonMessage, readyState, sendJsonMessage } = useWebSocket(
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
    if (token) {
      loadTodos({});
    } else {
      clearTodos();
    }
  }, [token]);

  useEffect(() => {
    // null is an 'object'
    if (typeof lastJsonMessage === "object" && typeof lastJsonMessage?.pos === "string") {
      pos.current.value = lastJsonMessage.pos;
    }
  }, [lastJsonMessage]);

  useEffect(() => {
    // const connectionStatus = {
    //   [ReadyState.CONNECTING]: "Connecting",
    //   [ReadyState.OPEN]: "Open",
    //   [ReadyState.CLOSING]: "Closing",
    //   [ReadyState.CLOSED]: "Closed",
    //   [ReadyState.UNINSTANTIATED]: "Uninstantiated",
    // }[readyState];

    // console.log("READY STATE", connectionStatus, readyState);

    if (readyState === ReadyState.OPEN) {
      // console.log("SENDING", pos.current.value);
      sendJsonMessage({ cursor: pos.current.value });
    }

    // if (readyState === ReadyState.OPEN) {
    //   setTimeout(() => {
    //     getWebSocket()?.close();
    //   }, 5000);
    // }
  }, [readyState]);

  // useEffect(() => {
  //   const handler = setInterval(() => {
  //     const s = getWebSocket();
  //     if (readyState === ReadyState.OPEN && s) {
  //       s.dispatchEvent
  //       // sendJsonMessage({ action: "ping" });
  //     }
  //   }, 30_000);

  //   return () => clearInterval(handler);
  // }, [readyState]);

  return null;
}

function SiteLayout() {
  return (
    <Flex w="100%">
      <Sidebar />
      <Outlet />
    </Flex>
  );
}

function Site() {
  return (
    <Routes>
      <Route path="/" element={<SiteLayout />}>
        <Route index element={<HomePage />} />
        <Route path="settings/*" element={<SettingsPage />} />
        <Route path="todo/*" element={<TodoPage />} />
        <Route path="*" element={<HomePage />} />
      </Route>
    </Routes>
  );
}

export default function App() {
  return (
    <ChakraProvider>
      <CSSReset />
      <NetworkContextProvider>
        <ReloadState />
        <BrowserRouter>
          <Routes>
            <Route path="/auth/*" element={<AuthPages />} />
            <Route path="*" element={<Site />} />
          </Routes>
        </BrowserRouter>
      </NetworkContextProvider>
    </ChakraProvider>
  );
}
