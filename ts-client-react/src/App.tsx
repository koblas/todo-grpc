import React from "react";

import CssBaseline from "@mui/material/CssBaseline";
import { createTheme, ThemeProvider } from "@mui/material/styles";

import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import { TodoContextProvider } from "./hooks/todo";
import { AuthLoginPage } from "./pages/AuthLoginPage";
import { TodoPage } from "./pages/TodoPage";
import { HomePage } from "./pages/HomePage";
import { AuthContextProvider } from "./hooks/auth";

const theme = createTheme();

export default function App() {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />

      <AuthContextProvider>
        <TodoContextProvider>
          <Router>
            <Switch>
              <Route path="/auth/login">
                <AuthLoginPage />
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
    </ThemeProvider>
  );
}
