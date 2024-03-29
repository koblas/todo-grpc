import { Json } from "../../types/json";

// https://fny9c2xm3b.execute-api.us-east-1.amazonaws.com/v1/auth.AuthenticationService/Authenticate
export const BASE_URL = process.env.BASE_URL ?? "/api";

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

  getInfo(): { code?: string; argument?: string; msg?: string } {
    if (!this.body) {
      return {
        code: "unknown",
      };
    }

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const { code, msg, meta } = this.body as any;

    if (!code) {
      return {
        code: "unknown",
      };
    }

    let argument: string | undefined;

    if (typeof meta === "object" && meta !== null && typeof meta.argument === "string") {
      argument = meta.argument;
    }

    return { code, argument, msg };
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
  "401": async (response) => {
    // In theory we should capture this specifically
    const body = await response.text();
    let error;
    try {
      error = new FetchError(response.status, JSON.parse(body));
    } catch {
      error = new FetchError(response.status, body);
    }
    throw error;
  },
};

const BASE_HEADERS = new Headers({
  Accept: "application/json",
  "Content-Type": "application/json",
  "X-sleep-delay": "1",
});

function buildHeaders(hdrs?: Record<string, string>) {
  const h = new Headers(BASE_HEADERS);

  if (!hdrs) {
    return h;
  }

  Object.entries(hdrs).forEach(([k, v]) => h.set(k, v));

  return h;
}

export function newFetchClient(config?: { token?: string | null; base?: string | null; handlers?: FetchHandlers }): {
  fetch: typeof fetch;
  POST<T = unknown>(url: string, body: Json): Promise<T>;
  PUT_FILE<T = unknown>(url: string, file: File): Promise<T>;
  POST_FILE<T = unknown>(url: string, file: File): Promise<T>;
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
    async POST<T>(url: string, body: Json): Promise<T> {
      const response = await fetchCommon(url, {
        method: "POST",
        body: JSON.stringify(body),
        headers: buildHeaders(config?.token ? { Authorization: `Bearer ${config?.token}` } : {}),
      });

      return response.json();
    },
    async POST_FILE<T>(url: string, file: File): Promise<T> {
      const fileHdrs = {
        "Content-Type": file.type,
        "Content-Length": String(file.size),
      };
      const response = await fetchCommon(url, {
        method: "POST",
        body: file,
        headers: buildHeaders(config?.token ? { Authorization: `Bearer ${config?.token}`, ...fileHdrs } : fileHdrs),
      });

      return response.json();
    },
    async PUT_FILE<T>(url: string, file: File): Promise<T> {
      const fileHdrs = {
        // M
        // "Content-Type": "application/octet-stream",
        "Content-Type": file.type,
        "Content-Length": String(file.size),
      };
      const response = await fetchCommon(url, {
        method: "PUT",
        body: file,
        headers: buildHeaders(config?.token ? { Authorization: `Bearer ${config?.token}`, ...fileHdrs } : fileHdrs),
      });

      if (response.headers.get("content-type")?.includes("json")) {
        return response.json();
      }
      return response.text() as T;
    },
    async GET<T>(url: string): Promise<T> {
      const response = await fetchCommon(url, {
        method: "GET",
        headers: buildHeaders(config?.token ? { Authorization: `Bearer ${config?.token}` } : {}),
      });

      return response.json();
    },
    async DELETE<T>(url: string): Promise<T> {
      const response = await fetchCommon(url, {
        method: "DELETE",
        headers: buildHeaders(config?.token ? { Authorization: `Bearer ${config?.token}` } : {}),
      });

      return response.json();
    },
  };
}
