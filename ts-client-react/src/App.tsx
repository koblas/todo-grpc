import React, { useEffect, useMemo } from "react";

import { ChakraProvider, CSSReset, Flex, Spinner, useToast } from "@chakra-ui/react";
import { BrowserRouter, Routes, Route, Outlet } from "react-router-dom";
import { QueryClient, QueryClientProvider } from "react-query";

import { AuthPages } from "./pages/auth";
import { SettingsPage } from "./pages/settings";
import { TodoPage } from "./pages/TodoPage";
import { HomePage } from "./pages/HomePage";
import { NetworkContextProvider } from "./hooks/network";
import Sidebar from "./components/Sidebar";
import { useAuth } from "./hooks/auth";
import { WebsocketProvider } from "./rpc/websocket";
import { FetchError } from "./rpc/utils";
import { useTodoListener } from "./hooks/data/todo";

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

function ClearOnLogout({ queryClient }: { queryClient: QueryClient }) {
  const { token } = useAuth();

  useEffect(() => {
    if (!token) {
      queryClient.clear();
    }
  }, [token, queryClient]);

  useTodoListener(queryClient);

  return null;
}

export default function App() {
  const toast = useToast();

  const queryClient = useMemo(() => {
    function onError(error: unknown) {
      if (error instanceof FetchError && error.getInfo().code !== "invalid_argument") {
        toast({
          title: "Network error",
          status: "error",
          isClosable: true,
        });
      }
    }

    return new QueryClient({
      defaultOptions: {
        queries: {
          suspense: true,
          onError,
        },
        mutations: {
          onError,
        },
      },
    });
  }, [toast]);

  return (
    <ChakraProvider>
      <CSSReset />
      <QueryClientProvider client={queryClient}>
        <NetworkContextProvider>
          <WebsocketProvider url={WS_URL}>
            <React.Suspense fallback={<Spinner />}>
              <ClearOnLogout queryClient={queryClient} />
              <BrowserRouter>
                <Routes>
                  <Route path="/auth/*" element={<AuthPages />} />
                  <Route path="*" element={<Site />} />
                </Routes>
              </BrowserRouter>
            </React.Suspense>
          </WebsocketProvider>
        </NetworkContextProvider>
      </QueryClientProvider>
    </ChakraProvider>
  );
}
