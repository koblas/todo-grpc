import { Json } from "../types/json";

const BASE = "http://localhost:8080";

export type FetchHandlers = Record<
  string,
  (response: Response, input: RequestInfo, init: RequestInit) => Promise<Response>
>;

const fetchHandlers: FetchHandlers = {
  "2xx": async (response) => response,
  "3xx": async (response) => {
    throw new Error(`Unexpected response.status=${response.status}`);
  },
  "4xx": async (response) => {
    throw new Error(`Unexpected response.status=${response.status}`);
  },
  "5xx": async (response) => {
    throw new Error(`Unexpected response.status=${response.status}`);
  },
  "401": async () => {
    throw new Error("Need authentication");
  },
};

export function newFetchClient(config?: { token?: string | null; base?: string | null; handlers?: FetchHandlers }): {
  fetch: typeof fetch;
  POST<T = Json>(url: string, body: Json): Promise<T>;
  GET<T = Json>(url: string): Promise<T>;
  DELETE<T = Json>(url: string): Promise<T>;
} {
  const hcombined = {
    ...fetchHandlers,
    ...(config?.handlers ?? {}),
  };

  async function fetchCommon(input: RequestInfo, init?: RequestInit): Promise<Response> {
    let response: Response;
    if (typeof input === "string") {
      response = await fetch(`${config?.base ?? BASE}${input}`, init);
    } else {
      response = await fetch(input, init);
    }

    const statusString = String(response.status);
    const statusGroup = String(response.status - (response.status % 100)).replace(/0/g, "x");

    if (statusString in hcombined) {
      return await hcombined[statusString]?.(response, input, init ?? {});
    }
    if (statusGroup in hcombined) {
      return await hcombined[statusGroup]?.(response, input, init ?? {});
    }

    throw new Error("Unhandled response");
  }

  return {
    fetch,
    async POST<T = Json>(url: string, body: Json): Promise<T> {
      const response = await fetchCommon(url, {
        method: "POST",
        body: JSON.stringify(body),
        headers: {
          "Content-Type": "application/json",
          ...(config?.token ? { Authorization: `Bearer ${config?.token}` } : {}),
        },
      });

      return response.json();
    },
    async GET<T = Json>(url: string): Promise<T> {
      const response = await fetchCommon(url, {
        method: "GET",
        headers: {
          ...(config?.token ? { Authorization: `Bearer ${config?.token}` } : {}),
        },
      });

      return response.json();
    },
    async DELETE<T = Json>(url: string): Promise<T> {
      const response = await fetchCommon(url, {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
          ...(config?.token ? { Authorization: `Bearer ${config?.token}` } : {}),
        },
      });

      return response.json();
    },
  };
}
