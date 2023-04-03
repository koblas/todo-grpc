import React, { useEffect, useMemo } from "react";
import { ChakraProvider, CSSReset, Spinner, useToast } from "@chakra-ui/react";
import { BrowserRouter } from "react-router-dom";
import { QueryClient, QueryClientProvider } from "react-query";
import { ErrorBoundary, FallbackProps } from "react-error-boundary";
import * as Sentry from "@sentry/react";

import { useAuth } from "./hooks/auth";
import { WebsocketProvider } from "./rpc/websocket";
import { FetchError } from "./rpc/utils";
import { useTodoListener } from "./hooks/data/todo";
import { useUserListener } from "./hooks/data/user";
import { AppRoutes } from "./AppRoutes";
import { useMessageListener } from "./hooks/data/message";
import { useAuthStore } from "./store/useAuthStore";

Sentry.init({
  debug: true,
});

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

function AuthErrorBoundary({ error }: FallbackProps) {
  // const navigate = useNavigate();
  // const { mutations } = useAuth();
  // const logout = mutations.useLogout();
  const { setToken } = useAuthStore((s) => s);

  if (error instanceof FetchError && error.code === 401) {
    // logout({
    //   onCompleted() {
    //     navigate("/auth/login");
    //   },
    // });
    setToken(null);
  } else {
    throw error;
  }

  // toast({
  //   status: "error",
  //   title: "Authentication failed",
  //   isClosable: true,
  //   onCloseComplete() {
  //     logout({
  //       onCompleted() {
  //         navigate("/auth/login");
  //       },
  //     });
  //   },
  // });

  return null;
}

function ClearOnLogout({ queryClient }: { queryClient: QueryClient }) {
  const { token } = useAuth();

  useEffect(() => {
    if (!token) {
      queryClient.clear();
    }
  }, [token, queryClient]);

  useTodoListener(queryClient);
  useUserListener(queryClient);
  useMessageListener(queryClient);

  return null;
}

export default function App(): React.ReactElement {
  const toast = useToast();

  const queryClient = useMemo(() => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    function onError(error: any) {
      Sentry.captureException(error);
      console.log("IN TOP LEVEL ERROR", error);
      if (error instanceof FetchError) {
        const { code } = error.getInfo();

        if (code === "internal") {
          toast({
            title: "Internal error",
            status: "error",
            isClosable: true,
          });
        } else if (code === "unauthenticated") {
          throw error;
        } else if (code !== "invalid_argument" && code !== "unauthenticated") {
          toast({
            title: "Network error",
            status: "error",
            isClosable: true,
          });
        }
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
      } else if (error.code !== 400) {
        // We ignore 400 on the assumption that it shouldn't happen

        // This is react-query throwing a CancelError
        const isCancelledError = error && Object.hasOwn(error, "silent");
        if (isCancelledError) {
          return;
        }

        toast({
          status: "error",
          title: `An unexpected error occured ${error.code}`,
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
      <ErrorBoundary FallbackComponent={AuthErrorBoundary}>
        <QueryClientProvider client={queryClient}>
          <WebsocketProvider url={WS_URL}>
            <React.Suspense fallback={<Spinner />}>
              <ClearOnLogout queryClient={queryClient} />
              <BrowserRouter>
                <AppRoutes />
              </BrowserRouter>
            </React.Suspense>
          </WebsocketProvider>
        </QueryClientProvider>
      </ErrorBoundary>
    </ChakraProvider>
  );
}
