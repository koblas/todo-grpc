import { Json } from "../../types/json";
import { RpcOptions } from "../errors";
import { handleJsonError } from "./json_helpers";

// https://fny9c2xm3b.execute-api.us-east-1.amazonaws.com/v1/auth.AuthenticationService/Authenticate
export const BASE_URL = process.env.BASE_URL ?? "http://localhost:8080";

export class FetchError extends Error {
  public body: Json;
  public code: number;

  constructor(code: number, body: Json) {
    super(`HTTP Error code=${code}`);
    this.code = code;
    this.body = body;
  }

  toString(): string {
    return `super().toString() ${JSON.stringify(this.body)}`;
  }
}

export type FetchHandlers = Record<
  string,
  (response: Response, input: RequestInfo, init: RequestInit) => Promise<Response>
>;

const fetchHandlers: FetchHandlers = {
  "2xx": async (response) => response,
  "3xx": async (response) => {
    const body = await response.text();
    let error;
    try {
      error = new FetchError(response.status, JSON.parse(body));
    } catch {
      error = new FetchError(response.status, body);
    }
    throw error;
  },
  "4xx": async (response) => {
    const body = await response.text();
    let error;
    try {
      error = new FetchError(response.status, JSON.parse(body));
    } catch {
      error = new FetchError(response.status, body);
    }
    throw error;
  },
  "5xx": async (response) => {
    const body = await response.text();
    let error;
    try {
      error = new FetchError(response.status, JSON.parse(body));
    } catch {
      error = new FetchError(response.status, body);
    }
    throw error;
  },
  "401": async () => {
    throw new Error("Need authentication");
  },
};

const baseHeaders = {
  Accept: "application/json",
  "Content-Type": "application/json",
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
      response = await fetch(`${config?.base ?? BASE_URL}${input}`, init);
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
          ...baseHeaders,
          ...(config?.token ? { Authorization: `Bearer ${config?.token}` } : {}),
        },
      });

      return response.json();
    },
    async GET<T = Json>(url: string): Promise<T> {
      const response = await fetchCommon(url, {
        method: "GET",
        headers: {
          ...baseHeaders,
          ...(config?.token ? { Authorization: `Bearer ${config?.token}` } : {}),
        },
      });

      return response.json();
    },
    async DELETE<T = Json>(url: string): Promise<T> {
      const response = await fetchCommon(url, {
        method: "DELETE",
        headers: {
          ...baseHeaders,
          ...(config?.token ? { Authorization: `Bearer ${config?.token}` } : {}),
        },
      });

      return response.json();
    },
  };
}

/**
 * newFetchPOST is a standard wrapper around newFetchClient that
 * handles the error cases without a lot of DRY violations.
 */
export function newFetchPOST(opts?: Parameters<typeof newFetchClient>[0]) {
  const client = newFetchClient(opts);

  return {
    call<T, R = T>(url: string, params: Json, options: RpcOptions<R>, xform?: (input: T) => R) {
      client
        .POST<T>(url, params)
        .then((data) => {
          options.onCompleted?.(xform ? xform(data) : (data as unknown as R));
        })
        .catch((err) => {
          handleJsonError(err, options);
        })
        .finally(() => {
          options.onFinished?.();
        });
    },
  };
}
