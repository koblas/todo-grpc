// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var protos_todo_pb = require('../protos/todo_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');

function serialize_todo_addTodoParams(arg) {
  if (!(arg instanceof protos_todo_pb.addTodoParams)) {
    throw new Error('Expected argument of type todo.addTodoParams');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_todo_addTodoParams(buffer_arg) {
  return protos_todo_pb.addTodoParams.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todo_deleteResponse(arg) {
  if (!(arg instanceof protos_todo_pb.deleteResponse)) {
    throw new Error('Expected argument of type todo.deleteResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_todo_deleteResponse(buffer_arg) {
  return protos_todo_pb.deleteResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todo_deleteTodoParams(arg) {
  if (!(arg instanceof protos_todo_pb.deleteTodoParams)) {
    throw new Error('Expected argument of type todo.deleteTodoParams');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_todo_deleteTodoParams(buffer_arg) {
  return protos_todo_pb.deleteTodoParams.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todo_getTodoParams(arg) {
  if (!(arg instanceof protos_todo_pb.getTodoParams)) {
    throw new Error('Expected argument of type todo.getTodoParams');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_todo_getTodoParams(buffer_arg) {
  return protos_todo_pb.getTodoParams.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todo_todoObject(arg) {
  if (!(arg instanceof protos_todo_pb.todoObject)) {
    throw new Error('Expected argument of type todo.todoObject');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_todo_todoObject(buffer_arg) {
  return protos_todo_pb.todoObject.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todo_todoResponse(arg) {
  if (!(arg instanceof protos_todo_pb.todoResponse)) {
    throw new Error('Expected argument of type todo.todoResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_todo_todoResponse(buffer_arg) {
  return protos_todo_pb.todoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var todoServiceService = exports.todoServiceService = {
  addTodo: {
    path: '/todo.todoService/addTodo',
    requestStream: false,
    responseStream: false,
    requestType: protos_todo_pb.addTodoParams,
    responseType: protos_todo_pb.todoObject,
    requestSerialize: serialize_todo_addTodoParams,
    requestDeserialize: deserialize_todo_addTodoParams,
    responseSerialize: serialize_todo_todoObject,
    responseDeserialize: deserialize_todo_todoObject,
  },
  deleteTodo: {
    path: '/todo.todoService/deleteTodo',
    requestStream: false,
    responseStream: false,
    requestType: protos_todo_pb.deleteTodoParams,
    responseType: protos_todo_pb.deleteResponse,
    requestSerialize: serialize_todo_deleteTodoParams,
    requestDeserialize: deserialize_todo_deleteTodoParams,
    responseSerialize: serialize_todo_deleteResponse,
    responseDeserialize: deserialize_todo_deleteResponse,
  },
  getTodos: {
    path: '/todo.todoService/getTodos',
    requestStream: false,
    responseStream: false,
    requestType: protos_todo_pb.getTodoParams,
    responseType: protos_todo_pb.todoResponse,
    requestSerialize: serialize_todo_getTodoParams,
    requestDeserialize: deserialize_todo_getTodoParams,
    responseSerialize: serialize_todo_todoResponse,
    responseDeserialize: deserialize_todo_todoResponse,
  },
};

exports.todoServiceClient = grpc.makeGenericClientConstructor(todoServiceService);
