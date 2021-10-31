import { Any } from "google-protobuf/google/protobuf/any_pb";
import { Status } from "../../genpb/google/rpc/status_pb";
import * as ErrorDetailsPb from "../../genpb/google/rpc/error_details_pb";

export type ErrorDetails =
  | ErrorDetailsPb.RetryInfo
  | ErrorDetailsPb.DebugInfo
  | ErrorDetailsPb.QuotaFailure
  | ErrorDetailsPb.ErrorInfo
  | ErrorDetailsPb.PreconditionFailure
  | ErrorDetailsPb.BadRequest
  | ErrorDetailsPb.BadRequest.FieldViolation
  | ErrorDetailsPb.RequestInfo
  | ErrorDetailsPb.ResourceInfo
  | ErrorDetailsPb.Help
  | ErrorDetailsPb.LocalizedMessage;

type Distribute<T> = T extends any
  ? {
      new (...args: any[]): T;
      deserializeBinary(bytes: Uint8Array): T;
    }
  : never;

const mapTypeUrlToErrorDetailClass = new Map<string, Distribute<ErrorDetails>>([
  ["type.googleapis.com/google.rpc.RetryInfo", ErrorDetailsPb.RetryInfo],
  ["type.googleapis.com/google.rpc.DebugInfo", ErrorDetailsPb.DebugInfo],
  ["type.googleapis.com/google.rpc.QuotaFailure", ErrorDetailsPb.QuotaFailure],
  ["type.googleapis.com/google.rpc.ErrorInfo", ErrorDetailsPb.ErrorInfo],
  ["type.googleapis.com/google.rpc.PreconditionFailure", ErrorDetailsPb.PreconditionFailure],
  ["type.googleapis.com/google.rpc.BadRequest", ErrorDetailsPb.BadRequest],
  ["type.googleapis.com/google.rpc.RequestInfo", ErrorDetailsPb.RequestInfo],
  ["type.googleapis.com/google.rpc.ResourceInfo", ErrorDetailsPb.ResourceInfo],
  ["type.googleapis.com/google.rpc.Help", ErrorDetailsPb.Help],
  ["type.googleapis.com/google.rpc.BadRequest.FieldViolation", ErrorDetailsPb.BadRequest.FieldViolation],
  ["type.googleapis.com/google.rpc.LocalizedMessage", ErrorDetailsPb.LocalizedMessage],
]);

function parseErrorDetails(details: Any): ErrorDetails | null {
  const typeUrl = details.getTypeUrl();
  const errorDetailsClass = mapTypeUrlToErrorDetailClass.get(typeUrl);
  if (!errorDetailsClass) {
    console.warn(`grpc-web-error-details: typeUrl "${typeUrl}" is not supported`);
    return null;
  }
  return errorDetailsClass.deserializeBinary(details.getValue_asU8());
}

export function statusFromError(
  err: any | { metadata: { "grpc-status-details-bin"?: string } },
): [Status, ErrorDetails[]] | [null, null] {
  // to get status, we requires err['metadata']['grpc-status-details-bin']
  const statusDetailsBinStr = err?.metadata?.["grpc-status-details-bin"];
  if (typeof statusDetailsBinStr !== "string") {
    // if the error does not contain status, return null
    return [null, null];
  }

  let bytes: Uint8Array;
  try {
    bytes = Buffer.from(statusDetailsBinStr, "base64");
  } catch {
    // `grpc-status-details-bin` has an invalid base64 string
    return [null, null];
  }

  const st = Status.deserializeBinary(bytes);
  const details = st
    .getDetailsList()
    .map((detail) => parseErrorDetails(detail))
    .filter((detail): detail is ErrorDetails => !!detail);

  return [st, details];
}
