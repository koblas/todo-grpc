import React from "react";
import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import App from "./App";

test("smoke", () => {
  const {} = render(<App />);
  const linkElement = screen.getByText(/Add Todo/i);
  expect(linkElement).toBeInTheDocument();
});
