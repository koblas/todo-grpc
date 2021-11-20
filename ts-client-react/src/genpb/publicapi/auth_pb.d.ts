import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from '../google/api/annotations_pb';


export class RegisterParams extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): RegisterParams;

  getPassword(): string;
  setPassword(value: string): RegisterParams;

  getName(): string;
  setName(value: string): RegisterParams;

  getInvite(): string;
  setInvite(value: string): RegisterParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterParams.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterParams): RegisterParams.AsObject;
  static serializeBinaryToWriter(message: RegisterParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterParams;
  static deserializeBinaryFromReader(message: RegisterParams, reader: jspb.BinaryReader): RegisterParams;
}

export namespace RegisterParams {
  export type AsObject = {
    email: string,
    password: string,
    name: string,
    invite: string,
  }
}

export class LoginParams extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): LoginParams;

  getPassword(): string;
  setPassword(value: string): LoginParams;

  getTfaOtp(): string;
  setTfaOtp(value: string): LoginParams;

  getTfaType(): string;
  setTfaType(value: string): LoginParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginParams.AsObject;
  static toObject(includeInstance: boolean, msg: LoginParams): LoginParams.AsObject;
  static serializeBinaryToWriter(message: LoginParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginParams;
  static deserializeBinaryFromReader(message: LoginParams, reader: jspb.BinaryReader): LoginParams;
}

export namespace LoginParams {
  export type AsObject = {
    email: string,
    password: string,
    tfaOtp: string,
    tfaType: string,
  }
}

export class ConfirmParams extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): ConfirmParams;

  getToken(): string;
  setToken(value: string): ConfirmParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ConfirmParams.AsObject;
  static toObject(includeInstance: boolean, msg: ConfirmParams): ConfirmParams.AsObject;
  static serializeBinaryToWriter(message: ConfirmParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ConfirmParams;
  static deserializeBinaryFromReader(message: ConfirmParams, reader: jspb.BinaryReader): ConfirmParams;
}

export namespace ConfirmParams {
  export type AsObject = {
    userId: string,
    token: string,
  }
}

export class RecoverySendParams extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): RecoverySendParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RecoverySendParams.AsObject;
  static toObject(includeInstance: boolean, msg: RecoverySendParams): RecoverySendParams.AsObject;
  static serializeBinaryToWriter(message: RecoverySendParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RecoverySendParams;
  static deserializeBinaryFromReader(message: RecoverySendParams, reader: jspb.BinaryReader): RecoverySendParams;
}

export namespace RecoverySendParams {
  export type AsObject = {
    email: string,
  }
}

export class RecoveryUpdateParams extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): RecoveryUpdateParams;

  getToken(): string;
  setToken(value: string): RecoveryUpdateParams;

  getPassword(): string;
  setPassword(value: string): RecoveryUpdateParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RecoveryUpdateParams.AsObject;
  static toObject(includeInstance: boolean, msg: RecoveryUpdateParams): RecoveryUpdateParams.AsObject;
  static serializeBinaryToWriter(message: RecoveryUpdateParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RecoveryUpdateParams;
  static deserializeBinaryFromReader(message: RecoveryUpdateParams, reader: jspb.BinaryReader): RecoveryUpdateParams;
}

export namespace RecoveryUpdateParams {
  export type AsObject = {
    userId: string,
    token: string,
    password: string,
  }
}

export class OauthUrlParams extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): OauthUrlParams;

  getRedirectUrl(): string;
  setRedirectUrl(value: string): OauthUrlParams;

  getState(): string;
  setState(value: string): OauthUrlParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OauthUrlParams.AsObject;
  static toObject(includeInstance: boolean, msg: OauthUrlParams): OauthUrlParams.AsObject;
  static serializeBinaryToWriter(message: OauthUrlParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OauthUrlParams;
  static deserializeBinaryFromReader(message: OauthUrlParams, reader: jspb.BinaryReader): OauthUrlParams;
}

export namespace OauthUrlParams {
  export type AsObject = {
    provider: string,
    redirectUrl: string,
    state: string,
  }
}

export class OauthUrlResult extends jspb.Message {
  getUrl(): string;
  setUrl(value: string): OauthUrlResult;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OauthUrlResult.AsObject;
  static toObject(includeInstance: boolean, msg: OauthUrlResult): OauthUrlResult.AsObject;
  static serializeBinaryToWriter(message: OauthUrlResult, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OauthUrlResult;
  static deserializeBinaryFromReader(message: OauthUrlResult, reader: jspb.BinaryReader): OauthUrlResult;
}

export namespace OauthUrlResult {
  export type AsObject = {
    url: string,
  }
}

export class OauthAssociateParams extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): OauthAssociateParams;

  getRedirectUrl(): string;
  setRedirectUrl(value: string): OauthAssociateParams;

  getCode(): string;
  setCode(value: string): OauthAssociateParams;

  getState(): string;
  setState(value: string): OauthAssociateParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OauthAssociateParams.AsObject;
  static toObject(includeInstance: boolean, msg: OauthAssociateParams): OauthAssociateParams.AsObject;
  static serializeBinaryToWriter(message: OauthAssociateParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OauthAssociateParams;
  static deserializeBinaryFromReader(message: OauthAssociateParams, reader: jspb.BinaryReader): OauthAssociateParams;
}

export namespace OauthAssociateParams {
  export type AsObject = {
    provider: string,
    redirectUrl: string,
    code: string,
    state: string,
  }
}

export class ValidationError extends jspb.Message {
  getField(): string;
  setField(value: string): ValidationError;

  getMessage(): string;
  setMessage(value: string): ValidationError;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ValidationError.AsObject;
  static toObject(includeInstance: boolean, msg: ValidationError): ValidationError.AsObject;
  static serializeBinaryToWriter(message: ValidationError, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ValidationError;
  static deserializeBinaryFromReader(message: ValidationError, reader: jspb.BinaryReader): ValidationError;
}

export namespace ValidationError {
  export type AsObject = {
    field: string,
    message: string,
  }
}

export class Success extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): Success;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Success.AsObject;
  static toObject(includeInstance: boolean, msg: Success): Success.AsObject;
  static serializeBinaryToWriter(message: Success, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Success;
  static deserializeBinaryFromReader(message: Success, reader: jspb.BinaryReader): Success;
}

export namespace Success {
  export type AsObject = {
    success: boolean,
  }
}

export class Token extends jspb.Message {
  getAccessToken(): string;
  setAccessToken(value: string): Token;

  getTokenType(): string;
  setTokenType(value: string): Token;

  getExpiresIn(): number;
  setExpiresIn(value: number): Token;

  getRefreshToken(): string;
  setRefreshToken(value: string): Token;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Token.AsObject;
  static toObject(includeInstance: boolean, msg: Token): Token.AsObject;
  static serializeBinaryToWriter(message: Token, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Token;
  static deserializeBinaryFromReader(message: Token, reader: jspb.BinaryReader): Token;
}

export namespace Token {
  export type AsObject = {
    accessToken: string,
    tokenType: string,
    expiresIn: number,
    refreshToken: string,
  }
}

export class TokenRegister extends jspb.Message {
  getToken(): Token | undefined;
  setToken(value?: Token): TokenRegister;
  hasToken(): boolean;
  clearToken(): TokenRegister;

  getCreated(): boolean;
  setCreated(value: boolean): TokenRegister;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TokenRegister.AsObject;
  static toObject(includeInstance: boolean, msg: TokenRegister): TokenRegister.AsObject;
  static serializeBinaryToWriter(message: TokenRegister, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TokenRegister;
  static deserializeBinaryFromReader(message: TokenRegister, reader: jspb.BinaryReader): TokenRegister;
}

export namespace TokenRegister {
  export type AsObject = {
    token?: Token.AsObject,
    created: boolean,
  }
}

