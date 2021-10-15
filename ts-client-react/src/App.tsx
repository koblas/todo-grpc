import React, { useEffect, useState } from "react";
import { BrowserRouter as Router, Switch, Route, useHistory } from "react-router-dom";
import { useTodos, TodoContextProvider } from "./hooks/todo";
import { TodoItem } from "./rpc/todo";
import { LoginPage } from "./auth/Login";

import "tailwindcss/dist/tailwind.css";
import { AuthContextProvider } from "./hooks/auth";

function Item({ todo }: { todo: TodoItem }) {
  const { deleteTodo } = useTodos();
  const { task, id } = todo;

  return (
    <li className="bg-white text-center p-2 m-1 w-2/6 border border-color-gray-500">
      <div className="flex justify-between">
        <h5 className="font-light">{task}</h5>
        <button
          onClick={() => deleteTodo(id)}
          className="text-red-500 pl-2 pr-2 font-medium rounded-xl hover:bg-red-600 hover:text-white"
        >
          X
        </button>
      </div>
    </li>
  );
}

function List({ todos }: { todos: TodoItem[] }) {
  return (
    <ul className="grid justify-items-center">
      {todos.map((todo) => (
        <Item key={todo.id} todo={todo} />
      ))}
    </ul>
  );
}

function TodoPage() {
  const { addTodo, todos } = useTodos();
  const [addText, setAddText] = useState("");

  function handleAdd() {
    addTodo(addText);
    setAddText("");
  }

  return (
    <div id="app" className="font-sans">
      <section>
        <h3 className="text-gray-800 text-center max-w-2xl text-2xl font-light m-auto mt-5 mb-5 p-5">
          TODO gRPC Client
        </h3>
        <form
          className="flex justify-center"
          onSubmit={(e) => {
            e.preventDefault();
            handleAdd();
          }}
        >
          <div className="space-x-5" role="add">
            <input
              placeholder="Walk my dog"
              className="pl-2 p-2 border-b border-gray-300 rounded-md w-64"
              type="text"
              value={addText}
              onChange={(e) => setAddText(e.target.value)}
            ></input>
            <button
              className="hover:bg-blue-800 bg-blue-600 rounded-md text-white px-5 text-base py-1"
              type="submit"
              name="action"
              onClick={(e) => {
                e.preventDefault();
                handleAdd();
              }}
            >
              Add Todo
            </button>
          </div>
        </form>
      </section>
      <div className="mt-6">
        <List todos={todos} />
      </div>
    </div>
  );
}

function NotFoundPage() {
  const history = useHistory();

  useEffect(() => {
    history.push("/todo");
  });

  return null;
}

export default function App() {
  return (
    <AuthContextProvider>
      <TodoContextProvider>
        <Router>
          <Switch>
            <Route path="/auth/login">
              <LoginPage />
            </Route>
            <Route path="/todo">
              <TodoPage />
            </Route>
            <Route path="*">
              <NotFoundPage />
            </Route>
          </Switch>
        </Router>
      </TodoContextProvider>
    </AuthContextProvider>
  );
}
