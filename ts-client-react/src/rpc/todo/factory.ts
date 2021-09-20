
import { newTodoClient as newClientGrpc } from './grpc_web';
import { newTodoClient as newClientJson } from './json_web';

export function newTodoClient(type: "grpc"|"json") {
  if (type === 'grpc') {
    return newClientGrpc();
  }

  return newClientJson();
}
