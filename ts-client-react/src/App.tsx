import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import { TodoContextProvider } from "./hooks/todo";
import { AuthLoginPage } from "./pages/AuthLoginPage";
import { TodoPage } from "./pages/TodoPage";
import { HomePage } from "./pages/HomePage";

import "tailwindcss/dist/tailwind.css";
import { AuthContextProvider } from "./hooks/auth";

export default function App() {
  return (
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
  );
}
