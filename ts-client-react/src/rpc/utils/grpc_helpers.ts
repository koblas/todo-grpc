import * as grpcWeb from "grpc-web";
import { statusFromError } from "./grpc_status_details";
import * as ErrorDetailsPb from "../../genpb/google/rpc/error_details_pb";
import { RpcOptions } from "../errors";

export function handleGrpcError<T>(err: grpcWeb.RpcError, options: RpcOptions<T>) {
  options.onError?.(err);

  const [st, details] = statusFromError(err);
  if (details) {
    const fields: Record<string, string[]> = {};

    details.forEach((item) => {
      if (!(item instanceof ErrorDetailsPb.BadRequest.FieldViolation)) {
        return;
      }

      const { field, description } = item.toObject();

      const lcfield = field.toLowerCase();
      fields[lcfield] = [...(fields[lcfield] ?? []), description];
    });

    options.onErrorField?.(fields);
    return;
  }

  options.onErrorNetwork?.(st);
}
