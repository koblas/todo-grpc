// package: todo
// file: todo.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as todo_pb from "./todo_pb";

interface ItodoServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    addTodo: ItodoServiceService_IaddTodo;
    deleteTodo: ItodoServiceService_IdeleteTodo;
    getTodos: ItodoServiceService_IgetTodos;
}

interface ItodoServiceService_IaddTodo extends grpc.MethodDefinition<todo_pb.addTodoParams, todo_pb.todoObject> {
    path: "/todo.todoService/addTodo";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<todo_pb.addTodoParams>;
    requestDeserialize: grpc.deserialize<todo_pb.addTodoParams>;
    responseSerialize: grpc.serialize<todo_pb.todoObject>;
    responseDeserialize: grpc.deserialize<todo_pb.todoObject>;
}
interface ItodoServiceService_IdeleteTodo extends grpc.MethodDefinition<todo_pb.deleteTodoParams, todo_pb.deleteResponse> {
    path: "/todo.todoService/deleteTodo";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<todo_pb.deleteTodoParams>;
    requestDeserialize: grpc.deserialize<todo_pb.deleteTodoParams>;
    responseSerialize: grpc.serialize<todo_pb.deleteResponse>;
    responseDeserialize: grpc.deserialize<todo_pb.deleteResponse>;
}
interface ItodoServiceService_IgetTodos extends grpc.MethodDefinition<todo_pb.getTodoParams, todo_pb.todoResponse> {
    path: "/todo.todoService/getTodos";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<todo_pb.getTodoParams>;
    requestDeserialize: grpc.deserialize<todo_pb.getTodoParams>;
    responseSerialize: grpc.serialize<todo_pb.todoResponse>;
    responseDeserialize: grpc.deserialize<todo_pb.todoResponse>;
}

export const todoServiceService: ItodoServiceService;

export interface ItodoServiceServer extends grpc.UntypedServiceImplementation {
    addTodo: grpc.handleUnaryCall<todo_pb.addTodoParams, todo_pb.todoObject>;
    deleteTodo: grpc.handleUnaryCall<todo_pb.deleteTodoParams, todo_pb.deleteResponse>;
    getTodos: grpc.handleUnaryCall<todo_pb.getTodoParams, todo_pb.todoResponse>;
}

export interface ItodoServiceClient {
    addTodo(request: todo_pb.addTodoParams, callback: (error: grpc.ServiceError | null, response: todo_pb.todoObject) => void): grpc.ClientUnaryCall;
    addTodo(request: todo_pb.addTodoParams, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: todo_pb.todoObject) => void): grpc.ClientUnaryCall;
    addTodo(request: todo_pb.addTodoParams, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: todo_pb.todoObject) => void): grpc.ClientUnaryCall;
    deleteTodo(request: todo_pb.deleteTodoParams, callback: (error: grpc.ServiceError | null, response: todo_pb.deleteResponse) => void): grpc.ClientUnaryCall;
    deleteTodo(request: todo_pb.deleteTodoParams, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: todo_pb.deleteResponse) => void): grpc.ClientUnaryCall;
    deleteTodo(request: todo_pb.deleteTodoParams, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: todo_pb.deleteResponse) => void): grpc.ClientUnaryCall;
    getTodos(request: todo_pb.getTodoParams, callback: (error: grpc.ServiceError | null, response: todo_pb.todoResponse) => void): grpc.ClientUnaryCall;
    getTodos(request: todo_pb.getTodoParams, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: todo_pb.todoResponse) => void): grpc.ClientUnaryCall;
    getTodos(request: todo_pb.getTodoParams, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: todo_pb.todoResponse) => void): grpc.ClientUnaryCall;
}

export class todoServiceClient extends grpc.Client implements ItodoServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public addTodo(request: todo_pb.addTodoParams, callback: (error: grpc.ServiceError | null, response: todo_pb.todoObject) => void): grpc.ClientUnaryCall;
    public addTodo(request: todo_pb.addTodoParams, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: todo_pb.todoObject) => void): grpc.ClientUnaryCall;
    public addTodo(request: todo_pb.addTodoParams, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: todo_pb.todoObject) => void): grpc.ClientUnaryCall;
    public deleteTodo(request: todo_pb.deleteTodoParams, callback: (error: grpc.ServiceError | null, response: todo_pb.deleteResponse) => void): grpc.ClientUnaryCall;
    public deleteTodo(request: todo_pb.deleteTodoParams, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: todo_pb.deleteResponse) => void): grpc.ClientUnaryCall;
    public deleteTodo(request: todo_pb.deleteTodoParams, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: todo_pb.deleteResponse) => void): grpc.ClientUnaryCall;
    public getTodos(request: todo_pb.getTodoParams, callback: (error: grpc.ServiceError | null, response: todo_pb.todoResponse) => void): grpc.ClientUnaryCall;
    public getTodos(request: todo_pb.getTodoParams, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: todo_pb.todoResponse) => void): grpc.ClientUnaryCall;
    public getTodos(request: todo_pb.getTodoParams, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: todo_pb.todoResponse) => void): grpc.ClientUnaryCall;
}
