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

  getUrlbase(): string;
  setUrlbase(value: string): RegisterParams;

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
    urlbase: string,
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
    token: string,
  }
}

export class RecoveryParams extends jspb.Message {
  getUrlbase(): string;
  setUrlbase(value: string): RecoveryParams;

  getEmail(): string;
  setEmail(value: string): RecoveryParams;

  getToken(): string;
  setToken(value: string): RecoveryParams;

  getPassword(): string;
  setPassword(value: string): RecoveryParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RecoveryParams.AsObject;
  static toObject(includeInstance: boolean, msg: RecoveryParams): RecoveryParams.AsObject;
  static serializeBinaryToWriter(message: RecoveryParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RecoveryParams;
  static deserializeBinaryFromReader(message: RecoveryParams, reader: jspb.BinaryReader): RecoveryParams;
}

export namespace RecoveryParams {
  export type AsObject = {
    urlbase: string,
    email: string,
    token: string,
    password: string,
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

export class TokenEither extends jspb.Message {
  getErrorsList(): Array<ValidationError>;
  setErrorsList(value: Array<ValidationError>): TokenEither;
  clearErrorsList(): TokenEither;
  addErrors(value?: ValidationError, index?: number): ValidationError;

  getToken(): Token | undefined;
  setToken(value?: Token): TokenEither;
  hasToken(): boolean;
  clearToken(): TokenEither;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TokenEither.AsObject;
  static toObject(includeInstance: boolean, msg: TokenEither): TokenEither.AsObject;
  static serializeBinaryToWriter(message: TokenEither, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TokenEither;
  static deserializeBinaryFromReader(message: TokenEither, reader: jspb.BinaryReader): TokenEither;
}

export namespace TokenEither {
  export type AsObject = {
    errorsList: Array<ValidationError.AsObject>,
    token?: Token.AsObject,
  }
}

export class SuccessEither extends jspb.Message {
  getErrorsList(): Array<ValidationError>;
  setErrorsList(value: Array<ValidationError>): SuccessEither;
  clearErrorsList(): SuccessEither;
  addErrors(value?: ValidationError, index?: number): ValidationError;

  getSuccess(): boolean;
  setSuccess(value: boolean): SuccessEither;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SuccessEither.AsObject;
  static toObject(includeInstance: boolean, msg: SuccessEither): SuccessEither.AsObject;
  static serializeBinaryToWriter(message: SuccessEither, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SuccessEither;
  static deserializeBinaryFromReader(message: SuccessEither, reader: jspb.BinaryReader): SuccessEither;
}

export namespace SuccessEither {
  export type AsObject = {
    errorsList: Array<ValidationError.AsObject>,
    success: boolean,
  }
}

