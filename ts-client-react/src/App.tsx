import React from "react";

import { ChakraProvider, CSSReset, Flex } from "@chakra-ui/react";

import { BrowserRouter, Routes, Route } from "react-router-dom";
import { TodoContextProvider } from "./hooks/todo";
import { AuthPages } from "./pages/auth";
import { SettingsPage } from "./pages/settings";
import { TodoPage } from "./pages/TodoPage";
import { HomePage } from "./pages/HomePage";
import { AuthContextProvider } from "./hooks/auth";
import { NetworkContextProvider } from "./hooks/network";
import Sidebar from "./components/Sidebar";

// const theme = createTheme();

function SiteLayout() {
  return (
    <Flex w="100%">
      <Sidebar />
    </Flex>
  );
}
function Site() {
  return (
    <Routes>
      <Route path="/" element={<SiteLayout />}>
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
        <AuthContextProvider>
          <TodoContextProvider>
            <BrowserRouter>
              <Routes>
                <Route path="/auth/*" element={<AuthPages />} />
                <Route path="*" element={<Site />} />
              </Routes>
            </BrowserRouter>
          </TodoContextProvider>
        </AuthContextProvider>
      </NetworkContextProvider>
    </ChakraProvider>
  );
}
