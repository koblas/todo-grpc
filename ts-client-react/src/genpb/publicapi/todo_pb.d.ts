import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from '../google/api/annotations_pb';


export class GetTodoParams extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTodoParams.AsObject;
  static toObject(includeInstance: boolean, msg: GetTodoParams): GetTodoParams.AsObject;
  static serializeBinaryToWriter(message: GetTodoParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTodoParams;
  static deserializeBinaryFromReader(message: GetTodoParams, reader: jspb.BinaryReader): GetTodoParams;
}

export namespace GetTodoParams {
  export type AsObject = {
  }
}

export class AddTodoParams extends jspb.Message {
  getTask(): string;
  setTask(value: string): AddTodoParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddTodoParams.AsObject;
  static toObject(includeInstance: boolean, msg: AddTodoParams): AddTodoParams.AsObject;
  static serializeBinaryToWriter(message: AddTodoParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddTodoParams;
  static deserializeBinaryFromReader(message: AddTodoParams, reader: jspb.BinaryReader): AddTodoParams;
}

export namespace AddTodoParams {
  export type AsObject = {
    task: string,
  }
}

export class DeleteTodoParams extends jspb.Message {
  getId(): string;
  setId(value: string): DeleteTodoParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteTodoParams.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteTodoParams): DeleteTodoParams.AsObject;
  static serializeBinaryToWriter(message: DeleteTodoParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteTodoParams;
  static deserializeBinaryFromReader(message: DeleteTodoParams, reader: jspb.BinaryReader): DeleteTodoParams;
}

export namespace DeleteTodoParams {
  export type AsObject = {
    id: string,
  }
}

export class TodoObject extends jspb.Message {
  getId(): string;
  setId(value: string): TodoObject;

  getTask(): string;
  setTask(value: string): TodoObject;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TodoObject.AsObject;
  static toObject(includeInstance: boolean, msg: TodoObject): TodoObject.AsObject;
  static serializeBinaryToWriter(message: TodoObject, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TodoObject;
  static deserializeBinaryFromReader(message: TodoObject, reader: jspb.BinaryReader): TodoObject;
}

export namespace TodoObject {
  export type AsObject = {
    id: string,
    task: string,
  }
}

export class TodoResponse extends jspb.Message {
  getTodosList(): Array<TodoObject>;
  setTodosList(value: Array<TodoObject>): TodoResponse;
  clearTodosList(): TodoResponse;
  addTodos(value?: TodoObject, index?: number): TodoObject;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TodoResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TodoResponse): TodoResponse.AsObject;
  static serializeBinaryToWriter(message: TodoResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TodoResponse;
  static deserializeBinaryFromReader(message: TodoResponse, reader: jspb.BinaryReader): TodoResponse;
}

export namespace TodoResponse {
  export type AsObject = {
    todosList: Array<TodoObject.AsObject>,
  }
}

export class DeleteResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): DeleteResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteResponse): DeleteResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteResponse;
  static deserializeBinaryFromReader(message: DeleteResponse, reader: jspb.BinaryReader): DeleteResponse;
}

export namespace DeleteResponse {
  export type AsObject = {
    message: string,
  }
}

