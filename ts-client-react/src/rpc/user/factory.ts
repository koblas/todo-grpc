// import { newUserClient as newClientGrpc } from "./grpc_web";
import { newUserClient as newClientJson } from "./json_web";

export function newUserClient(token: string | null, type: "grpc" | "json") {
  if (type === "grpc") {
    // return newClientGrpc(token);
    throw new Error("Not implemented")
  }

  return newClientJson(token);
}
