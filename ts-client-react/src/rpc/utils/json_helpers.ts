import { RpcOptions } from "../errors";

export function handleJsonError<T>(err: unknown, options: RpcOptions<T>) {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  options.onError?.(err as any);
  if (typeof err !== "object" || !(err instanceof Error)) {
    return;
  }
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const { code, body } = err as any;
  if (typeof code === "number" && typeof body === "object") {
    const fields: Record<string, string[]> = {};

    if (code === 400 && typeof body?.meta === "object") {
      // This is the twirp error response
      Object.entries(body.meta).forEach(([key, value]) => {
        if (typeof key === "string" && typeof value === "string") {
          const msg = (body?.msg ?? value ?? "") as string;

          fields[key.toLowerCase()] = [msg.replace(key, "").trim()];
        }
      });
    } else if (code === 400 && Array.isArray(body?.details)) {
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

    if (Object.keys(fields).length !== 0) {
      options.onErrorField?.(fields);
    }
    return;
  }
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  options.onErrorNetwork?.(err as any);
}
