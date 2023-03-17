import { RpcOptions } from "../errors";

const VALID_TYPES = ["google.rpc.BadRequest", "type.googleapis.com/google.rpc.BadRequest"];

export function handleJsonError<TData, TVar>(err: unknown, options?: RpcOptions<TData, TVar>): boolean {
  if (!options) {
    return false;
  }

  let found = Boolean(options.onError);

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  options.onError?.(err as any);
  if (typeof err !== "object" || !(err instanceof Error)) {
    return found;
  }
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const { code, body } = err as any;
  if (typeof code === "number" && typeof body === "object") {
    const fields: Record<string, string[]> = {};

    if (code === 400) {
      if (typeof body?.meta === "object") {
        // This is the twirp error response
        Object.entries(body.meta).forEach(([key, value]) => {
          if (typeof key === "string" && typeof value === "string") {
            const msg = (body?.msg ?? value ?? "") as string;

            fields[key.toLowerCase()] = [msg.replace(key, "").trim()];
          }
        });
      } else if (Array.isArray(body?.details)) {
        // Extract grpc status code from either grpc-web or buf-connect
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        const items = body.details.filter((item: any) => {
          if (typeof item !== "object" || item === null) {
            return false;
          }
          if (!Object.prototype.hasOwnProperty.call(item, "type") || !VALID_TYPES.includes(item.type)) {
            return false;
          }
          return true;
        });

        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        items.forEach((item: any) => {
          const fieldList = item.fieldViolations ?? item.debug?.fieldViolations;

          if (!Array.isArray(fieldList)) {
            return;
          }
          fieldList.forEach((entry) => {
            if (typeof entry !== "object" || entry === null) {
              return;
            }
            const { field, description } = entry;
            if (typeof description !== "string" || typeof field !== "string") {
              return;
            }
            const lcField = field.toLowerCase();
            fields[lcField] = (fields[lcField] ?? []).concat(description);
          });
        });
      }
    }

    if (Object.keys(fields).length !== 0) {
      options.onErrorField?.(fields);

      found = found || Boolean(options.onErrorField);
    }
  }

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  options.onErrorNetwork?.(err as any);

  return found || Boolean(options.onErrorNetwork);
}
