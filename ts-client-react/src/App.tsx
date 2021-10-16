import React from "react";

import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import { TodoContextProvider } from "./hooks/todo";
import { AuthLoginPage } from "./pages/AuthLoginPage";
import { AuthLoginPage as AuthLoginPage3 } from "./pages/AuthLoginPage3";
import { TodoPage } from "./pages/TodoPage";
import { HomePage } from "./pages/HomePage";
import { AuthContextProvider } from "./hooks/auth";

export default function App() {
  return (
    <AuthContextProvider>
      <TodoContextProvider>
        <Router>
          <Switch>
            <Route path="/auth/login3">
              <AuthLoginPage3 />
            </Route>
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
  );
}
