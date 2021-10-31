/* eslint-disable */
import Long from "long";
import { grpc } from "@improbable-eng/grpc-web";
import _m0 from "protobufjs/minimal";
import { BrowserHeaders } from "browser-headers";

export const protobufPackage = "todo";

export interface GetTodoParams {}

export interface AddTodoParams {
  task: string;
}

export interface DeleteTodoParams {
  id: string;
}

export interface TodoObject {
  id: string;
  task: string;
}

export interface TodoResponse {
  todos: TodoObject[];
}

export interface DeleteResponse {
  message: string;
}

const baseGetTodoParams: object = {};

export const GetTodoParams = {
  encode(
    _: GetTodoParams,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetTodoParams {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGetTodoParams } as GetTodoParams;
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

  fromJSON(_: any): GetTodoParams {
    const message = { ...baseGetTodoParams } as GetTodoParams;
    return message;
  },

  toJSON(_: GetTodoParams): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<GetTodoParams>): GetTodoParams {
    const message = { ...baseGetTodoParams } as GetTodoParams;
    return message;
  },
};

const baseAddTodoParams: object = { task: "" };

export const AddTodoParams = {
  encode(
    message: AddTodoParams,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.task !== "") {
      writer.uint32(10).string(message.task);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AddTodoParams {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAddTodoParams } as AddTodoParams;
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

  fromJSON(object: any): AddTodoParams {
    const message = { ...baseAddTodoParams } as AddTodoParams;
    if (object.task !== undefined && object.task !== null) {
      message.task = String(object.task);
    } else {
      message.task = "";
    }
    return message;
  },

  toJSON(message: AddTodoParams): unknown {
    const obj: any = {};
    message.task !== undefined && (obj.task = message.task);
    return obj;
  },

  fromPartial(object: DeepPartial<AddTodoParams>): AddTodoParams {
    const message = { ...baseAddTodoParams } as AddTodoParams;
    if (object.task !== undefined && object.task !== null) {
      message.task = object.task;
    } else {
      message.task = "";
    }
    return message;
  },
};

const baseDeleteTodoParams: object = { id: "" };

export const DeleteTodoParams = {
  encode(
    message: DeleteTodoParams,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteTodoParams {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDeleteTodoParams } as DeleteTodoParams;
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

  fromJSON(object: any): DeleteTodoParams {
    const message = { ...baseDeleteTodoParams } as DeleteTodoParams;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    return message;
  },

  toJSON(message: DeleteTodoParams): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<DeleteTodoParams>): DeleteTodoParams {
    const message = { ...baseDeleteTodoParams } as DeleteTodoParams;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    return message;
  },
};

const baseTodoObject: object = { id: "", task: "" };

export const TodoObject = {
  encode(
    message: TodoObject,
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

  decode(input: _m0.Reader | Uint8Array, length?: number): TodoObject {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTodoObject } as TodoObject;
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

  fromJSON(object: any): TodoObject {
    const message = { ...baseTodoObject } as TodoObject;
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

  toJSON(message: TodoObject): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.task !== undefined && (obj.task = message.task);
    return obj;
  },

  fromPartial(object: DeepPartial<TodoObject>): TodoObject {
    const message = { ...baseTodoObject } as TodoObject;
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

const baseTodoResponse: object = {};

export const TodoResponse = {
  encode(
    message: TodoResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.todos) {
      TodoObject.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TodoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTodoResponse } as TodoResponse;
    message.todos = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.todos.push(TodoObject.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TodoResponse {
    const message = { ...baseTodoResponse } as TodoResponse;
    message.todos = [];
    if (object.todos !== undefined && object.todos !== null) {
      for (const e of object.todos) {
        message.todos.push(TodoObject.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: TodoResponse): unknown {
    const obj: any = {};
    if (message.todos) {
      obj.todos = message.todos.map((e) =>
        e ? TodoObject.toJSON(e) : undefined
      );
    } else {
      obj.todos = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<TodoResponse>): TodoResponse {
    const message = { ...baseTodoResponse } as TodoResponse;
    message.todos = [];
    if (object.todos !== undefined && object.todos !== null) {
      for (const e of object.todos) {
        message.todos.push(TodoObject.fromPartial(e));
      }
    }
    return message;
  },
};

const baseDeleteResponse: object = { message: "" };

export const DeleteResponse = {
  encode(
    message: DeleteResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.message !== "") {
      writer.uint32(10).string(message.message);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDeleteResponse } as DeleteResponse;
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

  fromJSON(object: any): DeleteResponse {
    const message = { ...baseDeleteResponse } as DeleteResponse;
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    return message;
  },

  toJSON(message: DeleteResponse): unknown {
    const obj: any = {};
    message.message !== undefined && (obj.message = message.message);
    return obj;
  },

  fromPartial(object: DeepPartial<DeleteResponse>): DeleteResponse {
    const message = { ...baseDeleteResponse } as DeleteResponse;
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    return message;
  },
};

export interface TodoService {
  addTodo(
    request: DeepPartial<AddTodoParams>,
    metadata?: grpc.Metadata
  ): Promise<TodoObject>;
  deleteTodo(
    request: DeepPartial<DeleteTodoParams>,
    metadata?: grpc.Metadata
  ): Promise<DeleteResponse>;
  getTodos(
    request: DeepPartial<GetTodoParams>,
    metadata?: grpc.Metadata
  ): Promise<TodoResponse>;
}

export class TodoServiceClientImpl implements TodoService {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.addTodo = this.addTodo.bind(this);
    this.deleteTodo = this.deleteTodo.bind(this);
    this.getTodos = this.getTodos.bind(this);
  }

  addTodo(
    request: DeepPartial<AddTodoParams>,
    metadata?: grpc.Metadata
  ): Promise<TodoObject> {
    return this.rpc.unary(
      TodoServiceaddTodoDesc,
      AddTodoParams.fromPartial(request),
      metadata
    );
  }

  deleteTodo(
    request: DeepPartial<DeleteTodoParams>,
    metadata?: grpc.Metadata
  ): Promise<DeleteResponse> {
    return this.rpc.unary(
      TodoServicedeleteTodoDesc,
      DeleteTodoParams.fromPartial(request),
      metadata
    );
  }

  getTodos(
    request: DeepPartial<GetTodoParams>,
    metadata?: grpc.Metadata
  ): Promise<TodoResponse> {
    return this.rpc.unary(
      TodoServicegetTodosDesc,
      GetTodoParams.fromPartial(request),
      metadata
    );
  }
}

export const TodoServiceDesc = {
  serviceName: "todo.TodoService",
};

export const TodoServiceaddTodoDesc: UnaryMethodDefinitionish = {
  methodName: "addTodo",
  service: TodoServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return AddTodoParams.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...TodoObject.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const TodoServicedeleteTodoDesc: UnaryMethodDefinitionish = {
  methodName: "deleteTodo",
  service: TodoServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return DeleteTodoParams.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...DeleteResponse.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const TodoServicegetTodosDesc: UnaryMethodDefinitionish = {
  methodName: "getTodos",
  service: TodoServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return GetTodoParams.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...TodoResponse.decode(data),
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
