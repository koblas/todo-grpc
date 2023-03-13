import React from "react";
import { act } from "react-dom/test-utils";
import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import "@testing-library/jest-dom";
import App from "./App";

// jest.mock("./rpc/todo/factory.ts");

function waitForNextTick() {
  return new Promise<void>((resolve) => {
    setTimeout(resolve);
  });
}

async function renderWait(ui: React.ReactElement) {
  render(ui);

  // Need to wait for the promises to finish so that it thinks we're done
  await act(waitForNextTick);
}

test("smoke", async () => {
  await renderWait(<App />);

  const linkElement = screen.getByText(/Sign in to your account/i);
  expect(linkElement).toBeInTheDocument();
});

test("add", async () => {
  await renderWait(<App />);

  const input = await waitFor(() => screen.getByRole("textbox", { name: "Email address" }));
  const button = screen.getByRole("button", { name: "Sign in" });

  fireEvent.change(input, { target: { value: "user@example.com" } });
  fireEvent.click(button);

  // const items = await waitFor(() => screen.findAllByRole("error"));
  const items = await waitFor(() => screen.findAllByText("Password is required"));
  expect(items).toHaveLength(1);
});
