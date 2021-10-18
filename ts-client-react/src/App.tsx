import React from "react";

import { ChakraProvider, CSSReset } from "@chakra-ui/react";

import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import { TodoContextProvider } from "./hooks/todo";
import { AuthLoginPage } from "./pages/AuthLoginPage";
import { AuthLogoutPage } from "./pages/AuthLogoutPage";
import { TodoPage } from "./pages/TodoPage";
import { HomePage } from "./pages/HomePage";
import { AuthContextProvider } from "./hooks/auth";

// const theme = createTheme();

export default function App() {
  return (
    <ChakraProvider>
      <CSSReset />
      <AuthContextProvider>
        <TodoContextProvider>
          <Router>
            <Switch>
              <Route path="/auth/login">
                <AuthLoginPage />
              </Route>
              <Route path="/auth/logout">
                <AuthLogoutPage />
              </Route>
              <Route path="/todo">
                <TodoPage />
              </Route>
              <Route path="*">
                <HomePage />
              </Route>
            </Switch>
          </Router>
        </TodoContextProvider>
      </AuthContextProvider>
    </ChakraProvider>
  );
}
