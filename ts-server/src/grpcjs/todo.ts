/* eslint-disable */
import Long from "long";
import {
  makeGenericClientConstructor,
  ChannelCredentials,
  ChannelOptions,
  UntypedServiceImplementation,
  handleUnaryCall,
  Client,
  ClientUnaryCall,
  Metadata,
  CallOptions,
  ServiceError,
} from "@grpc/grpc-js";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "todo";

export interface getTodoParams {}

export interface addTodoParams {
  task: string;
}

export interface deleteTodoParams {
  id: string;
}

export interface todoObject {
  id: string;
  task: string;
}

export interface todoResponse {
  todos: todoObject[];
}

export interface deleteResponse {
  message: string;
}

const basegetTodoParams: object = {};

export const getTodoParams = {
  encode(
    _: getTodoParams,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): getTodoParams {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basegetTodoParams } as getTodoParams;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): getTodoParams {
    const message = { ...basegetTodoParams } as getTodoParams;
    return message;
  },

  toJSON(_: getTodoParams): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<getTodoParams>): getTodoParams {
    const message = { ...basegetTodoParams } as getTodoParams;
    return message;
  },
};

const baseaddTodoParams: object = { task: "" };

export const addTodoParams = {
  encode(
    message: addTodoParams,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.task !== "") {
      writer.uint32(10).string(message.task);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): addTodoParams {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseaddTodoParams } as addTodoParams;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.task = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): addTodoParams {
    const message = { ...baseaddTodoParams } as addTodoParams;
    if (object.task !== undefined && object.task !== null) {
      message.task = String(object.task);
    } else {
      message.task = "";
    }
    return message;
  },

  toJSON(message: addTodoParams): unknown {
    const obj: any = {};
    message.task !== undefined && (obj.task = message.task);
    return obj;
  },

  fromPartial(object: DeepPartial<addTodoParams>): addTodoParams {
    const message = { ...baseaddTodoParams } as addTodoParams;
    if (object.task !== undefined && object.task !== null) {
      message.task = object.task;
    } else {
      message.task = "";
    }
    return message;
  },
};

const basedeleteTodoParams: object = { id: "" };

export const deleteTodoParams = {
  encode(
    message: deleteTodoParams,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): deleteTodoParams {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basedeleteTodoParams } as deleteTodoParams;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): deleteTodoParams {
    const message = { ...basedeleteTodoParams } as deleteTodoParams;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    return message;
  },

  toJSON(message: deleteTodoParams): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<deleteTodoParams>): deleteTodoParams {
    const message = { ...basedeleteTodoParams } as deleteTodoParams;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    return message;
  },
};

const basetodoObject: object = { id: "", task: "" };

export const todoObject = {
  encode(
    message: todoObject,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.task !== "") {
      writer.uint32(18).string(message.task);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): todoObject {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basetodoObject } as todoObject;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.task = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): todoObject {
    const message = { ...basetodoObject } as todoObject;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.task !== undefined && object.task !== null) {
      message.task = String(object.task);
    } else {
      message.task = "";
    }
    return message;
  },

  toJSON(message: todoObject): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.task !== undefined && (obj.task = message.task);
    return obj;
  },

  fromPartial(object: DeepPartial<todoObject>): todoObject {
    const message = { ...basetodoObject } as todoObject;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.task !== undefined && object.task !== null) {
      message.task = object.task;
    } else {
      message.task = "";
    }
    return message;
  },
};

const basetodoResponse: object = {};

export const todoResponse = {
  encode(
    message: todoResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.todos) {
      todoObject.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): todoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basetodoResponse } as todoResponse;
    message.todos = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.todos.push(todoObject.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): todoResponse {
    const message = { ...basetodoResponse } as todoResponse;
    message.todos = [];
    if (object.todos !== undefined && object.todos !== null) {
      for (const e of object.todos) {
        message.todos.push(todoObject.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: todoResponse): unknown {
    const obj: any = {};
    if (message.todos) {
      obj.todos = message.todos.map((e) =>
        e ? todoObject.toJSON(e) : undefined
      );
    } else {
      obj.todos = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<todoResponse>): todoResponse {
    const message = { ...basetodoResponse } as todoResponse;
    message.todos = [];
    if (object.todos !== undefined && object.todos !== null) {
      for (const e of object.todos) {
        message.todos.push(todoObject.fromPartial(e));
      }
    }
    return message;
  },
};

const basedeleteResponse: object = { message: "" };

export const deleteResponse = {
  encode(
    message: deleteResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.message !== "") {
      writer.uint32(10).string(message.message);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): deleteResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basedeleteResponse } as deleteResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.message = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): deleteResponse {
    const message = { ...basedeleteResponse } as deleteResponse;
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    return message;
  },

  toJSON(message: deleteResponse): unknown {
    const obj: any = {};
    message.message !== undefined && (obj.message = message.message);
    return obj;
  },

  fromPartial(object: DeepPartial<deleteResponse>): deleteResponse {
    const message = { ...basedeleteResponse } as deleteResponse;
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    return message;
  },
};

export const todoServiceService = {
  addTodo: {
    path: "/todo.todoService/addTodo",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: addTodoParams) =>
      Buffer.from(addTodoParams.encode(value).finish()),
    requestDeserialize: (value: Buffer) => addTodoParams.decode(value),
    responseSerialize: (value: todoObject) =>
      Buffer.from(todoObject.encode(value).finish()),
    responseDeserialize: (value: Buffer) => todoObject.decode(value),
  },
  deleteTodo: {
    path: "/todo.todoService/deleteTodo",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: deleteTodoParams) =>
      Buffer.from(deleteTodoParams.encode(value).finish()),
    requestDeserialize: (value: Buffer) => deleteTodoParams.decode(value),
    responseSerialize: (value: deleteResponse) =>
      Buffer.from(deleteResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => deleteResponse.decode(value),
  },
  getTodos: {
    path: "/todo.todoService/getTodos",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: getTodoParams) =>
      Buffer.from(getTodoParams.encode(value).finish()),
    requestDeserialize: (value: Buffer) => getTodoParams.decode(value),
    responseSerialize: (value: todoResponse) =>
      Buffer.from(todoResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => todoResponse.decode(value),
  },
} as const;

export interface todoServiceServer extends UntypedServiceImplementation {
  addTodo: handleUnaryCall<addTodoParams, todoObject>;
  deleteTodo: handleUnaryCall<deleteTodoParams, deleteResponse>;
  getTodos: handleUnaryCall<getTodoParams, todoResponse>;
}

export interface todoServiceClient extends Client {
  addTodo(
    request: addTodoParams,
    callback: (error: ServiceError | null, response: todoObject) => void
  ): ClientUnaryCall;
  addTodo(
    request: addTodoParams,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: todoObject) => void
  ): ClientUnaryCall;
  addTodo(
    request: addTodoParams,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: todoObject) => void
  ): ClientUnaryCall;
  deleteTodo(
    request: deleteTodoParams,
    callback: (error: ServiceError | null, response: deleteResponse) => void
  ): ClientUnaryCall;
  deleteTodo(
    request: deleteTodoParams,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: deleteResponse) => void
  ): ClientUnaryCall;
  deleteTodo(
    request: deleteTodoParams,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: deleteResponse) => void
  ): ClientUnaryCall;
  getTodos(
    request: getTodoParams,
    callback: (error: ServiceError | null, response: todoResponse) => void
  ): ClientUnaryCall;
  getTodos(
    request: getTodoParams,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: todoResponse) => void
  ): ClientUnaryCall;
  getTodos(
    request: getTodoParams,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: todoResponse) => void
  ): ClientUnaryCall;
}

export const todoServiceClient = makeGenericClientConstructor(
  todoServiceService,
  "todo.todoService"
) as unknown as {
  new (
    address: string,
    credentials: ChannelCredentials,
    options?: Partial<ChannelOptions>
  ): todoServiceClient;
};

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
