/* eslint-disable */
import Long from "long";
import { grpc } from "@improbable-eng/grpc-web";
import _m0 from "protobufjs/minimal";
import { BrowserHeaders } from "browser-headers";

export const protobufPackage = "auth";

export interface RegisterParams {
  /** Required */
  email: string;
  /** Required */
  password: string;
  /** Username */
  name: string;
  /** invite code */
  invite: string;
  /** URL base for the UI */
  urlbase: string;
}

export interface LoginParams {
  /** Required */
  email: string;
  /** Required */
  password: string;
  /** TFA One Time Token */
  tfaOtp: string;
  /** TFA Type */
  tfaType: string;
}

export interface ConfirmParams {
  /** Email confirmation token */
  token: string;
}

export interface RecoveryParams {
  /** URL base for the UI */
  urlbase: string;
  /** Required: Send */
  email: string;
  /** Required: Verify and Update */
  token: string;
  /** Required: Update */
  password: string;
}

export interface ValidationError {
  /** Field name */
  field: string;
  /** Human readable message */
  message: string;
}

export interface Success {
  success: boolean;
}

export interface Token {
  accessToken: string;
  tokenType: string;
  expiresIn: number;
  refreshToken: string;
}

export interface TokenEither {
  /** If there are errors */
  errors: ValidationError[];
  /** Present if errors is empty */
  token: Token | undefined;
}

export interface SuccessEither {
  /** If there are errors */
  errors: ValidationError[];
  /** Present if errors is empty */
  success: boolean;
}

const baseRegisterParams: object = {
  email: "",
  password: "",
  name: "",
  invite: "",
  urlbase: "",
};

export const RegisterParams = {
  encode(
    message: RegisterParams,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.email !== "") {
      writer.uint32(10).string(message.email);
    }
    if (message.password !== "") {
      writer.uint32(18).string(message.password);
    }
    if (message.name !== "") {
      writer.uint32(34).string(message.name);
    }
    if (message.invite !== "") {
      writer.uint32(42).string(message.invite);
    }
    if (message.urlbase !== "") {
      writer.uint32(50).string(message.urlbase);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RegisterParams {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseRegisterParams } as RegisterParams;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.email = reader.string();
          break;
        case 2:
          message.password = reader.string();
          break;
        case 4:
          message.name = reader.string();
          break;
        case 5:
          message.invite = reader.string();
          break;
        case 6:
          message.urlbase = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RegisterParams {
    const message = { ...baseRegisterParams } as RegisterParams;
    if (object.email !== undefined && object.email !== null) {
      message.email = String(object.email);
    } else {
      message.email = "";
    }
    if (object.password !== undefined && object.password !== null) {
      message.password = String(object.password);
    } else {
      message.password = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.invite !== undefined && object.invite !== null) {
      message.invite = String(object.invite);
    } else {
      message.invite = "";
    }
    if (object.urlbase !== undefined && object.urlbase !== null) {
      message.urlbase = String(object.urlbase);
    } else {
      message.urlbase = "";
    }
    return message;
  },

  toJSON(message: RegisterParams): unknown {
    const obj: any = {};
    message.email !== undefined && (obj.email = message.email);
    message.password !== undefined && (obj.password = message.password);
    message.name !== undefined && (obj.name = message.name);
    message.invite !== undefined && (obj.invite = message.invite);
    message.urlbase !== undefined && (obj.urlbase = message.urlbase);
    return obj;
  },

  fromPartial(object: DeepPartial<RegisterParams>): RegisterParams {
    const message = { ...baseRegisterParams } as RegisterParams;
    if (object.email !== undefined && object.email !== null) {
      message.email = object.email;
    } else {
      message.email = "";
    }
    if (object.password !== undefined && object.password !== null) {
      message.password = object.password;
    } else {
      message.password = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.invite !== undefined && object.invite !== null) {
      message.invite = object.invite;
    } else {
      message.invite = "";
    }
    if (object.urlbase !== undefined && object.urlbase !== null) {
      message.urlbase = object.urlbase;
    } else {
      message.urlbase = "";
    }
    return message;
  },
};

const baseLoginParams: object = {
  email: "",
  password: "",
  tfaOtp: "",
  tfaType: "",
};

export const LoginParams = {
  encode(
    message: LoginParams,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.email !== "") {
      writer.uint32(10).string(message.email);
    }
    if (message.password !== "") {
      writer.uint32(18).string(message.password);
    }
    if (message.tfaOtp !== "") {
      writer.uint32(34).string(message.tfaOtp);
    }
    if (message.tfaType !== "") {
      writer.uint32(42).string(message.tfaType);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LoginParams {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLoginParams } as LoginParams;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.email = reader.string();
          break;
        case 2:
          message.password = reader.string();
          break;
        case 4:
          message.tfaOtp = reader.string();
          break;
        case 5:
          message.tfaType = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LoginParams {
    const message = { ...baseLoginParams } as LoginParams;
    if (object.email !== undefined && object.email !== null) {
      message.email = String(object.email);
    } else {
      message.email = "";
    }
    if (object.password !== undefined && object.password !== null) {
      message.password = String(object.password);
    } else {
      message.password = "";
    }
    if (object.tfaOtp !== undefined && object.tfaOtp !== null) {
      message.tfaOtp = String(object.tfaOtp);
    } else {
      message.tfaOtp = "";
    }
    if (object.tfaType !== undefined && object.tfaType !== null) {
      message.tfaType = String(object.tfaType);
    } else {
      message.tfaType = "";
    }
    return message;
  },

  toJSON(message: LoginParams): unknown {
    const obj: any = {};
    message.email !== undefined && (obj.email = message.email);
    message.password !== undefined && (obj.password = message.password);
    message.tfaOtp !== undefined && (obj.tfaOtp = message.tfaOtp);
    message.tfaType !== undefined && (obj.tfaType = message.tfaType);
    return obj;
  },

  fromPartial(object: DeepPartial<LoginParams>): LoginParams {
    const message = { ...baseLoginParams } as LoginParams;
    if (object.email !== undefined && object.email !== null) {
      message.email = object.email;
    } else {
      message.email = "";
    }
    if (object.password !== undefined && object.password !== null) {
      message.password = object.password;
    } else {
      message.password = "";
    }
    if (object.tfaOtp !== undefined && object.tfaOtp !== null) {
      message.tfaOtp = object.tfaOtp;
    } else {
      message.tfaOtp = "";
    }
    if (object.tfaType !== undefined && object.tfaType !== null) {
      message.tfaType = object.tfaType;
    } else {
      message.tfaType = "";
    }
    return message;
  },
};

const baseConfirmParams: object = { token: "" };

export const ConfirmParams = {
  encode(
    message: ConfirmParams,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ConfirmParams {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseConfirmParams } as ConfirmParams;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ConfirmParams {
    const message = { ...baseConfirmParams } as ConfirmParams;
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    return message;
  },

  toJSON(message: ConfirmParams): unknown {
    const obj: any = {};
    message.token !== undefined && (obj.token = message.token);
    return obj;
  },

  fromPartial(object: DeepPartial<ConfirmParams>): ConfirmParams {
    const message = { ...baseConfirmParams } as ConfirmParams;
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    return message;
  },
};

const baseRecoveryParams: object = {
  urlbase: "",
  email: "",
  token: "",
  password: "",
};

export const RecoveryParams = {
  encode(
    message: RecoveryParams,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.urlbase !== "") {
      writer.uint32(10).string(message.urlbase);
    }
    if (message.email !== "") {
      writer.uint32(18).string(message.email);
    }
    if (message.token !== "") {
      writer.uint32(26).string(message.token);
    }
    if (message.password !== "") {
      writer.uint32(34).string(message.password);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RecoveryParams {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseRecoveryParams } as RecoveryParams;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.urlbase = reader.string();
          break;
        case 2:
          message.email = reader.string();
          break;
        case 3:
          message.token = reader.string();
          break;
        case 4:
          message.password = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RecoveryParams {
    const message = { ...baseRecoveryParams } as RecoveryParams;
    if (object.urlbase !== undefined && object.urlbase !== null) {
      message.urlbase = String(object.urlbase);
    } else {
      message.urlbase = "";
    }
    if (object.email !== undefined && object.email !== null) {
      message.email = String(object.email);
    } else {
      message.email = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = String(object.token);
    } else {
      message.token = "";
    }
    if (object.password !== undefined && object.password !== null) {
      message.password = String(object.password);
    } else {
      message.password = "";
    }
    return message;
  },

  toJSON(message: RecoveryParams): unknown {
    const obj: any = {};
    message.urlbase !== undefined && (obj.urlbase = message.urlbase);
    message.email !== undefined && (obj.email = message.email);
    message.token !== undefined && (obj.token = message.token);
    message.password !== undefined && (obj.password = message.password);
    return obj;
  },

  fromPartial(object: DeepPartial<RecoveryParams>): RecoveryParams {
    const message = { ...baseRecoveryParams } as RecoveryParams;
    if (object.urlbase !== undefined && object.urlbase !== null) {
      message.urlbase = object.urlbase;
    } else {
      message.urlbase = "";
    }
    if (object.email !== undefined && object.email !== null) {
      message.email = object.email;
    } else {
      message.email = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = object.token;
    } else {
      message.token = "";
    }
    if (object.password !== undefined && object.password !== null) {
      message.password = object.password;
    } else {
      message.password = "";
    }
    return message;
  },
};

const baseValidationError: object = { field: "", message: "" };

export const ValidationError = {
  encode(
    message: ValidationError,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.field !== "") {
      writer.uint32(10).string(message.field);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ValidationError {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseValidationError } as ValidationError;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.field = reader.string();
          break;
        case 2:
          message.message = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ValidationError {
    const message = { ...baseValidationError } as ValidationError;
    if (object.field !== undefined && object.field !== null) {
      message.field = String(object.field);
    } else {
      message.field = "";
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    return message;
  },

  toJSON(message: ValidationError): unknown {
    const obj: any = {};
    message.field !== undefined && (obj.field = message.field);
    message.message !== undefined && (obj.message = message.message);
    return obj;
  },

  fromPartial(object: DeepPartial<ValidationError>): ValidationError {
    const message = { ...baseValidationError } as ValidationError;
    if (object.field !== undefined && object.field !== null) {
      message.field = object.field;
    } else {
      message.field = "";
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    return message;
  },
};

const baseSuccess: object = { success: false };

export const Success = {
  encode(
    message: Success,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Success {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSuccess } as Success;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Success {
    const message = { ...baseSuccess } as Success;
    if (object.success !== undefined && object.success !== null) {
      message.success = Boolean(object.success);
    } else {
      message.success = false;
    }
    return message;
  },

  toJSON(message: Success): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial(object: DeepPartial<Success>): Success {
    const message = { ...baseSuccess } as Success;
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = false;
    }
    return message;
  },
};

const baseToken: object = {
  accessToken: "",
  tokenType: "",
  expiresIn: 0,
  refreshToken: "",
};

export const Token = {
  encode(message: Token, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.accessToken !== "") {
      writer.uint32(10).string(message.accessToken);
    }
    if (message.tokenType !== "") {
      writer.uint32(18).string(message.tokenType);
    }
    if (message.expiresIn !== 0) {
      writer.uint32(24).int32(message.expiresIn);
    }
    if (message.refreshToken !== "") {
      writer.uint32(34).string(message.refreshToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Token {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseToken } as Token;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.accessToken = reader.string();
          break;
        case 2:
          message.tokenType = reader.string();
          break;
        case 3:
          message.expiresIn = reader.int32();
          break;
        case 4:
          message.refreshToken = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Token {
    const message = { ...baseToken } as Token;
    if (object.accessToken !== undefined && object.accessToken !== null) {
      message.accessToken = String(object.accessToken);
    } else {
      message.accessToken = "";
    }
    if (object.tokenType !== undefined && object.tokenType !== null) {
      message.tokenType = String(object.tokenType);
    } else {
      message.tokenType = "";
    }
    if (object.expiresIn !== undefined && object.expiresIn !== null) {
      message.expiresIn = Number(object.expiresIn);
    } else {
      message.expiresIn = 0;
    }
    if (object.refreshToken !== undefined && object.refreshToken !== null) {
      message.refreshToken = String(object.refreshToken);
    } else {
      message.refreshToken = "";
    }
    return message;
  },

  toJSON(message: Token): unknown {
    const obj: any = {};
    message.accessToken !== undefined &&
      (obj.accessToken = message.accessToken);
    message.tokenType !== undefined && (obj.tokenType = message.tokenType);
    message.expiresIn !== undefined && (obj.expiresIn = message.expiresIn);
    message.refreshToken !== undefined &&
      (obj.refreshToken = message.refreshToken);
    return obj;
  },

  fromPartial(object: DeepPartial<Token>): Token {
    const message = { ...baseToken } as Token;
    if (object.accessToken !== undefined && object.accessToken !== null) {
      message.accessToken = object.accessToken;
    } else {
      message.accessToken = "";
    }
    if (object.tokenType !== undefined && object.tokenType !== null) {
      message.tokenType = object.tokenType;
    } else {
      message.tokenType = "";
    }
    if (object.expiresIn !== undefined && object.expiresIn !== null) {
      message.expiresIn = object.expiresIn;
    } else {
      message.expiresIn = 0;
    }
    if (object.refreshToken !== undefined && object.refreshToken !== null) {
      message.refreshToken = object.refreshToken;
    } else {
      message.refreshToken = "";
    }
    return message;
  },
};

const baseTokenEither: object = {};

export const TokenEither = {
  encode(
    message: TokenEither,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.errors) {
      ValidationError.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.token !== undefined) {
      Token.encode(message.token, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TokenEither {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTokenEither } as TokenEither;
    message.errors = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.errors.push(ValidationError.decode(reader, reader.uint32()));
          break;
        case 2:
          message.token = Token.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TokenEither {
    const message = { ...baseTokenEither } as TokenEither;
    message.errors = [];
    if (object.errors !== undefined && object.errors !== null) {
      for (const e of object.errors) {
        message.errors.push(ValidationError.fromJSON(e));
      }
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = Token.fromJSON(object.token);
    } else {
      message.token = undefined;
    }
    return message;
  },

  toJSON(message: TokenEither): unknown {
    const obj: any = {};
    if (message.errors) {
      obj.errors = message.errors.map((e) =>
        e ? ValidationError.toJSON(e) : undefined
      );
    } else {
      obj.errors = [];
    }
    message.token !== undefined &&
      (obj.token = message.token ? Token.toJSON(message.token) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<TokenEither>): TokenEither {
    const message = { ...baseTokenEither } as TokenEither;
    message.errors = [];
    if (object.errors !== undefined && object.errors !== null) {
      for (const e of object.errors) {
        message.errors.push(ValidationError.fromPartial(e));
      }
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = Token.fromPartial(object.token);
    } else {
      message.token = undefined;
    }
    return message;
  },
};

const baseSuccessEither: object = { success: false };

export const SuccessEither = {
  encode(
    message: SuccessEither,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.errors) {
      ValidationError.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.success === true) {
      writer.uint32(16).bool(message.success);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SuccessEither {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSuccessEither } as SuccessEither;
    message.errors = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.errors.push(ValidationError.decode(reader, reader.uint32()));
          break;
        case 2:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SuccessEither {
    const message = { ...baseSuccessEither } as SuccessEither;
    message.errors = [];
    if (object.errors !== undefined && object.errors !== null) {
      for (const e of object.errors) {
        message.errors.push(ValidationError.fromJSON(e));
      }
    }
    if (object.success !== undefined && object.success !== null) {
      message.success = Boolean(object.success);
    } else {
      message.success = false;
    }
    return message;
  },

  toJSON(message: SuccessEither): unknown {
    const obj: any = {};
    if (message.errors) {
      obj.errors = message.errors.map((e) =>
        e ? ValidationError.toJSON(e) : undefined
      );
    } else {
      obj.errors = [];
    }
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial(object: DeepPartial<SuccessEither>): SuccessEither {
    const message = { ...baseSuccessEither } as SuccessEither;
    message.errors = [];
    if (object.errors !== undefined && object.errors !== null) {
      for (const e of object.errors) {
        message.errors.push(ValidationError.fromPartial(e));
      }
    }
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = false;
    }
    return message;
  },
};

export interface AuthenticationService {
  register(
    request: DeepPartial<RegisterParams>,
    metadata?: grpc.Metadata
  ): Promise<TokenEither>;
  authenticate(
    request: DeepPartial<LoginParams>,
    metadata?: grpc.Metadata
  ): Promise<TokenEither>;
  verify_email(
    request: DeepPartial<ConfirmParams>,
    metadata?: grpc.Metadata
  ): Promise<TokenEither>;
  recover_send(
    request: DeepPartial<RecoveryParams>,
    metadata?: grpc.Metadata
  ): Promise<SuccessEither>;
  recover_verify(
    request: DeepPartial<RecoveryParams>,
    metadata?: grpc.Metadata
  ): Promise<SuccessEither>;
  recover_update(
    request: DeepPartial<RecoveryParams>,
    metadata?: grpc.Metadata
  ): Promise<TokenEither>;
}

export class AuthenticationServiceClientImpl implements AuthenticationService {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.register = this.register.bind(this);
    this.authenticate = this.authenticate.bind(this);
    this.verify_email = this.verify_email.bind(this);
    this.recover_send = this.recover_send.bind(this);
    this.recover_verify = this.recover_verify.bind(this);
    this.recover_update = this.recover_update.bind(this);
  }

  register(
    request: DeepPartial<RegisterParams>,
    metadata?: grpc.Metadata
  ): Promise<TokenEither> {
    return this.rpc.unary(
      AuthenticationServiceregisterDesc,
      RegisterParams.fromPartial(request),
      metadata
    );
  }

  authenticate(
    request: DeepPartial<LoginParams>,
    metadata?: grpc.Metadata
  ): Promise<TokenEither> {
    return this.rpc.unary(
      AuthenticationServiceauthenticateDesc,
      LoginParams.fromPartial(request),
      metadata
    );
  }

  verify_email(
    request: DeepPartial<ConfirmParams>,
    metadata?: grpc.Metadata
  ): Promise<TokenEither> {
    return this.rpc.unary(
      AuthenticationServiceverify_emailDesc,
      ConfirmParams.fromPartial(request),
      metadata
    );
  }

  recover_send(
    request: DeepPartial<RecoveryParams>,
    metadata?: grpc.Metadata
  ): Promise<SuccessEither> {
    return this.rpc.unary(
      AuthenticationServicerecover_sendDesc,
      RecoveryParams.fromPartial(request),
      metadata
    );
  }

  recover_verify(
    request: DeepPartial<RecoveryParams>,
    metadata?: grpc.Metadata
  ): Promise<SuccessEither> {
    return this.rpc.unary(
      AuthenticationServicerecover_verifyDesc,
      RecoveryParams.fromPartial(request),
      metadata
    );
  }

  recover_update(
    request: DeepPartial<RecoveryParams>,
    metadata?: grpc.Metadata
  ): Promise<TokenEither> {
    return this.rpc.unary(
      AuthenticationServicerecover_updateDesc,
      RecoveryParams.fromPartial(request),
      metadata
    );
  }
}

export const AuthenticationServiceDesc = {
  serviceName: "auth.AuthenticationService",
};

export const AuthenticationServiceregisterDesc: UnaryMethodDefinitionish = {
  methodName: "register",
  service: AuthenticationServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return RegisterParams.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...TokenEither.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const AuthenticationServiceauthenticateDesc: UnaryMethodDefinitionish = {
  methodName: "authenticate",
  service: AuthenticationServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return LoginParams.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...TokenEither.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const AuthenticationServiceverify_emailDesc: UnaryMethodDefinitionish = {
  methodName: "verify_email",
  service: AuthenticationServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return ConfirmParams.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...TokenEither.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const AuthenticationServicerecover_sendDesc: UnaryMethodDefinitionish = {
  methodName: "recover_send",
  service: AuthenticationServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return RecoveryParams.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...SuccessEither.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const AuthenticationServicerecover_verifyDesc: UnaryMethodDefinitionish =
  {
    methodName: "recover_verify",
    service: AuthenticationServiceDesc,
    requestStream: false,
    responseStream: false,
    requestType: {
      serializeBinary() {
        return RecoveryParams.encode(this).finish();
      },
    } as any,
    responseType: {
      deserializeBinary(data: Uint8Array) {
        return {
          ...SuccessEither.decode(data),
          toObject() {
            return this;
          },
        };
      },
    } as any,
  };

export const AuthenticationServicerecover_updateDesc: UnaryMethodDefinitionish =
  {
    methodName: "recover_update",
    service: AuthenticationServiceDesc,
    requestStream: false,
    responseStream: false,
    requestType: {
      serializeBinary() {
        return RecoveryParams.encode(this).finish();
      },
    } as any,
    responseType: {
      deserializeBinary(data: Uint8Array) {
        return {
          ...TokenEither.decode(data),
          toObject() {
            return this;
          },
        };
      },
    } as any,
  };

interface UnaryMethodDefinitionishR
  extends grpc.UnaryMethodDefinition<any, any> {
  requestStream: any;
  responseStream: any;
}

type UnaryMethodDefinitionish = UnaryMethodDefinitionishR;

interface Rpc {
  unary<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    request: any,
    metadata: grpc.Metadata | undefined
  ): Promise<any>;
}

export class GrpcWebImpl {
  private host: string;
  private options: {
    transport?: grpc.TransportFactory;

    debug?: boolean;
    metadata?: grpc.Metadata;
  };

  constructor(
    host: string,
    options: {
      transport?: grpc.TransportFactory;

      debug?: boolean;
      metadata?: grpc.Metadata;
    }
  ) {
    this.host = host;
    this.options = options;
  }

  unary<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    _request: any,
    metadata: grpc.Metadata | undefined
  ): Promise<any> {
    const request = { ..._request, ...methodDesc.requestType };
    const maybeCombinedMetadata =
      metadata && this.options.metadata
        ? new BrowserHeaders({
            ...this.options?.metadata.headersMap,
            ...metadata?.headersMap,
          })
        : metadata || this.options.metadata;
    return new Promise((resolve, reject) => {
      grpc.unary(methodDesc, {
        request,
        host: this.host,
        metadata: maybeCombinedMetadata,
        transport: this.options.transport,
        debug: this.options.debug,
        onEnd: function (response) {
          if (response.status === grpc.Code.OK) {
            resolve(response.message);
          } else {
            const err = new Error(response.statusMessage) as any;
            err.code = response.status;
            err.metadata = response.trailers;
            reject(err);
          }
        },
      });
    });
  }
}

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}
