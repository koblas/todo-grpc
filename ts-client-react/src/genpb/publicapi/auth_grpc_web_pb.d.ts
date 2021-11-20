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
               response: publicapi_auth_pb.TokenRegister) => void
  ): grpcWeb.ClientReadableStream<publicapi_auth_pb.TokenRegister>;

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
               response: publicapi_auth_pb.Success) => void
  ): grpcWeb.ClientReadableStream<publicapi_auth_pb.Success>;

  recover_send(
    request: publicapi_auth_pb.RecoverySendParams,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: publicapi_auth_pb.Success) => void
  ): grpcWeb.ClientReadableStream<publicapi_auth_pb.Success>;

  recover_verify(
    request: publicapi_auth_pb.RecoveryUpdateParams,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: publicapi_auth_pb.Success) => void
  ): grpcWeb.ClientReadableStream<publicapi_auth_pb.Success>;

  recover_update(
    request: publicapi_auth_pb.RecoveryUpdateParams,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: publicapi_auth_pb.Token) => void
  ): grpcWeb.ClientReadableStream<publicapi_auth_pb.Token>;

  oauth_login(
    request: publicapi_auth_pb.OauthAssociateParams,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: publicapi_auth_pb.TokenRegister) => void
  ): grpcWeb.ClientReadableStream<publicapi_auth_pb.TokenRegister>;

  oauth_url(
    request: publicapi_auth_pb.OauthUrlParams,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: publicapi_auth_pb.OauthUrlResult) => void
  ): grpcWeb.ClientReadableStream<publicapi_auth_pb.OauthUrlResult>;

}

export class AuthenticationServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  register(
    request: publicapi_auth_pb.RegisterParams,
    metadata?: grpcWeb.Metadata
  ): Promise<publicapi_auth_pb.TokenRegister>;

  authenticate(
    request: publicapi_auth_pb.LoginParams,
    metadata?: grpcWeb.Metadata
  ): Promise<publicapi_auth_pb.Token>;

  verify_email(
    request: publicapi_auth_pb.ConfirmParams,
    metadata?: grpcWeb.Metadata
  ): Promise<publicapi_auth_pb.Success>;

  recover_send(
    request: publicapi_auth_pb.RecoverySendParams,
    metadata?: grpcWeb.Metadata
  ): Promise<publicapi_auth_pb.Success>;

  recover_verify(
    request: publicapi_auth_pb.RecoveryUpdateParams,
    metadata?: grpcWeb.Metadata
  ): Promise<publicapi_auth_pb.Success>;

  recover_update(
    request: publicapi_auth_pb.RecoveryUpdateParams,
    metadata?: grpcWeb.Metadata
  ): Promise<publicapi_auth_pb.Token>;

  oauth_login(
    request: publicapi_auth_pb.OauthAssociateParams,
    metadata?: grpcWeb.Metadata
  ): Promise<publicapi_auth_pb.TokenRegister>;

  oauth_url(
    request: publicapi_auth_pb.OauthUrlParams,
    metadata?: grpcWeb.Metadata
  ): Promise<publicapi_auth_pb.OauthUrlResult>;

}

