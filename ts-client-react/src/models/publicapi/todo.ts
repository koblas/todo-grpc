/* eslint-disable */
import Long from "long";
import { grpc } from "@improbable-eng/grpc-web";
import _m0 from "protobufjs/minimal";
import { BrowserHeaders } from "browser-headers";

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

export interface todoService {
  addTodo(
    request: DeepPartial<addTodoParams>,
    metadata?: grpc.Metadata
  ): Promise<todoObject>;
  deleteTodo(
    request: DeepPartial<deleteTodoParams>,
    metadata?: grpc.Metadata
  ): Promise<deleteResponse>;
  getTodos(
    request: DeepPartial<getTodoParams>,
    metadata?: grpc.Metadata
  ): Promise<todoResponse>;
}

export class todoServiceClientImpl implements todoService {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.addTodo = this.addTodo.bind(this);
    this.deleteTodo = this.deleteTodo.bind(this);
    this.getTodos = this.getTodos.bind(this);
  }

  addTodo(
    request: DeepPartial<addTodoParams>,
    metadata?: grpc.Metadata
  ): Promise<todoObject> {
    return this.rpc.unary(
      todoServiceaddTodoDesc,
      addTodoParams.fromPartial(request),
      metadata
    );
  }

  deleteTodo(
    request: DeepPartial<deleteTodoParams>,
    metadata?: grpc.Metadata
  ): Promise<deleteResponse> {
    return this.rpc.unary(
      todoServicedeleteTodoDesc,
      deleteTodoParams.fromPartial(request),
      metadata
    );
  }

  getTodos(
    request: DeepPartial<getTodoParams>,
    metadata?: grpc.Metadata
  ): Promise<todoResponse> {
    return this.rpc.unary(
      todoServicegetTodosDesc,
      getTodoParams.fromPartial(request),
      metadata
    );
  }
}

export const todoServiceDesc = {
  serviceName: "todo.todoService",
};

export const todoServiceaddTodoDesc: UnaryMethodDefinitionish = {
  methodName: "addTodo",
  service: todoServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return addTodoParams.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...todoObject.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const todoServicedeleteTodoDesc: UnaryMethodDefinitionish = {
  methodName: "deleteTodo",
  service: todoServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return deleteTodoParams.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...deleteResponse.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const todoServicegetTodosDesc: UnaryMethodDefinitionish = {
  methodName: "getTodos",
  service: todoServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return getTodoParams.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...todoResponse.decode(data),
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
