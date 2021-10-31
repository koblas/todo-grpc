import * as grpcWeb from 'grpc-web';

import * as publicapi_auth_pb from '../publicapi/auth_pb';


export class AuthenticationServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  register(
    request: publicapi_auth_pb.RegisterParams,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: publicapi_auth_pb.Token) => void
  ): grpcWeb.ClientReadableStream<publicapi_auth_pb.Token>;

  authenticate(
    request: publicapi_auth_pb.LoginParams,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: publicapi_auth_pb.Token) => void
  ): grpcWeb.ClientReadableStream<publicapi_auth_pb.Token>;

  verify_email(
    request: publicapi_auth_pb.ConfirmParams,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: publicapi_auth_pb.TokenEither) => void
  ): grpcWeb.ClientReadableStream<publicapi_auth_pb.TokenEither>;

  recover_send(
    request: publicapi_auth_pb.RecoveryParams,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: publicapi_auth_pb.SuccessEither) => void
  ): grpcWeb.ClientReadableStream<publicapi_auth_pb.SuccessEither>;

  recover_verify(
    request: publicapi_auth_pb.RecoveryParams,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: publicapi_auth_pb.SuccessEither) => void
  ): grpcWeb.ClientReadableStream<publicapi_auth_pb.SuccessEither>;

  recover_update(
    request: publicapi_auth_pb.RecoveryParams,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: publicapi_auth_pb.TokenEither) => void
  ): grpcWeb.ClientReadableStream<publicapi_auth_pb.TokenEither>;

}

export class AuthenticationServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  register(
    request: publicapi_auth_pb.RegisterParams,
    metadata?: grpcWeb.Metadata
  ): Promise<publicapi_auth_pb.Token>;

  authenticate(
    request: publicapi_auth_pb.LoginParams,
    metadata?: grpcWeb.Metadata
  ): Promise<publicapi_auth_pb.Token>;

  verify_email(
    request: publicapi_auth_pb.ConfirmParams,
    metadata?: grpcWeb.Metadata
  ): Promise<publicapi_auth_pb.TokenEither>;

  recover_send(
    request: publicapi_auth_pb.RecoveryParams,
    metadata?: grpcWeb.Metadata
  ): Promise<publicapi_auth_pb.SuccessEither>;

  recover_verify(
    request: publicapi_auth_pb.RecoveryParams,
    metadata?: grpcWeb.Metadata
  ): Promise<publicapi_auth_pb.SuccessEither>;

  recover_update(
    request: publicapi_auth_pb.RecoveryParams,
    metadata?: grpcWeb.Metadata
  ): Promise<publicapi_auth_pb.TokenEither>;

}

