import { RpcOptions } from "../errors";

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
        // this is the gRPC status response
        body.details
          .filter(
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            (item: any) =>
              typeof item.field === "string" &&
              typeof item.description === "string" &&
              item["@type"] === "type.googleapis.com/google.rpc.BadRequest.FieldViolation",
          )
          .forEach(({ field, description }: { field: string; description: string }) => {
            const lcField = field.toLowerCase();
            fields[lcField] = (fields[lcField] ?? []).concat(description);
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
