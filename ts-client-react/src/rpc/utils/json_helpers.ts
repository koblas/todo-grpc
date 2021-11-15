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
    if (code === 400 && Array.isArray(body?.details)) {
      const fields: Record<string, string[]> = {};
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
      options.onErrorField?.(fields);
      return;
    }
  }
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  options.onErrorNetwork?.(err as any);
}
