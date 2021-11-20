import React from "react";

import { ChakraProvider, CSSReset, Flex } from "@chakra-ui/react";

import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import { TodoContextProvider } from "./hooks/todo";
import { AuthPages } from "./pages/auth";
import { SettingsPage } from "./pages/settings";
import { TodoPage } from "./pages/TodoPage";
import { HomePage } from "./pages/HomePage";
import { AuthContextProvider } from "./hooks/auth";
import { NetworkContextProvider } from "./hooks/network";
import Sidebar from "./components/Sidebar";

// const theme = createTheme();

function Site() {
  return (
    <Flex w="100%">
      <Sidebar />
      <Switch>
        <Route path="/settings">
          <SettingsPage />
        </Route>
        <Route path="/todo">
          <TodoPage />
        </Route>
        <Route path="*">
          <HomePage />
        </Route>
      </Switch>
    </Flex>
  );
}

export default function App() {
  return (
    <ChakraProvider>
      <CSSReset />
      <NetworkContextProvider>
        <AuthContextProvider>
          <TodoContextProvider>
            <Router>
              <Switch>
                <Route path="/auth">
                  <AuthPages />
                </Route>
                <Route path="*">
                  <Site />
                </Route>
              </Switch>
            </Router>
          </TodoContextProvider>
        </AuthContextProvider>
      </NetworkContextProvider>
    </ChakraProvider>
  );
}
