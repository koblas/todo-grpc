import { newAuthClient as newClientGrpc } from "./grpc_web";
import { newAuthClient as newClientJson } from "./json_web";

export function newAuthClient(type: "grpc" | "json") {
  if (type === "grpc") {
    return newClientGrpc();
  }

  return newClientJson();
}
