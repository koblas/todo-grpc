import React from "react";

import { ChakraProvider, CSSReset, Flex } from "@chakra-ui/react";

import { BrowserRouter, Routes, Route, Outlet } from "react-router-dom";
import { TodoContextProvider } from "./hooks/todo";
import { AuthPages } from "./pages/auth";
import { SettingsPage } from "./pages/settings";
import { TodoPage } from "./pages/TodoPage";
import { HomePage } from "./pages/HomePage";
import { NetworkContextProvider } from "./hooks/network";
import Sidebar from "./components/Sidebar";

// const theme = createTheme();

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
        <TodoContextProvider>
          <BrowserRouter>
            <Routes>
              <Route path="/auth/*" element={<AuthPages />} />
              <Route path="*" element={<Site />} />
            </Routes>
          </BrowserRouter>
        </TodoContextProvider>
      </NetworkContextProvider>
    </ChakraProvider>
  );
}
