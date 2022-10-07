import React, { useEffect } from "react";

import { ChakraProvider, CSSReset, Flex } from "@chakra-ui/react";

import { BrowserRouter, Routes, Route, Outlet } from "react-router-dom";
import { AuthPages } from "./pages/auth";
import { SettingsPage } from "./pages/settings";
import { TodoPage } from "./pages/TodoPage";
import { HomePage } from "./pages/HomePage";
import { NetworkContextProvider } from "./hooks/network";
import Sidebar from "./components/Sidebar";
import { useTodos } from "./hooks/todo";
import { useAuth } from "./hooks/auth";
import { WebsocketProvider } from "./rpc/websocket";

function buildWebsocketUrl(): string {
  const base = process.env.WS_URL ?? "/wsapi";

  if (base === "") {
    return "";
  }
  if (base.startsWith("ws")) {
    return base;
  }

  const loc = window.location;

  return new URL(base, (loc.protocol === "https:" ? "wss://" : "ws://") + loc.host + loc.pathname).toString();
}

const WS_URL = buildWebsocketUrl();

function ReloadState() {
  const { token } = useAuth();
  // const pos = useRef<{ value: string | null }>({ value: null });
  const { mutations } = useTodos();
  const [loadTodos] = mutations.useLoadTodos();
  const [clearTodos] = mutations.useClearTodos();

  useEffect(() => {
    if (token) {
      loadTodos({});
    } else {
      clearTodos();
    }
  }, [token]);

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
        <WebsocketProvider url={WS_URL}>
          <>
            <ReloadState />
            <BrowserRouter>
              <Routes>
                <Route path="/auth/*" element={<AuthPages />} />
                <Route path="*" element={<Site />} />
              </Routes>
            </BrowserRouter>
          </>
        </WebsocketProvider>
      </NetworkContextProvider>
    </ChakraProvider>
  );
}
