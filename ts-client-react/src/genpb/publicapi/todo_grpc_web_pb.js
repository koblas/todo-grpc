/**
 * @fileoverview gRPC-Web generated client stub for todo
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_api_annotations_pb = require('../google/api/annotations_pb.js')
const proto = {};
proto.todo = require('./todo_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.todo.TodoServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.todo.TodoServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.todo.AddTodoParams,
 *   !proto.todo.TodoObject>}
 */
const methodDescriptor_TodoService_addTodo = new grpc.web.MethodDescriptor(
  '/todo.TodoService/addTodo',
  grpc.web.MethodType.UNARY,
  proto.todo.AddTodoParams,
  proto.todo.TodoObject,
  /**
   * @param {!proto.todo.AddTodoParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.TodoObject.deserializeBinary
);


/**
 * @param {!proto.todo.AddTodoParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.todo.TodoObject)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.todo.TodoObject>|undefined}
 *     The XHR Node Readable Stream
 */
proto.todo.TodoServiceClient.prototype.addTodo =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/todo.TodoService/addTodo',
      request,
      metadata || {},
      methodDescriptor_TodoService_addTodo,
      callback);
};


/**
 * @param {!proto.todo.AddTodoParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.todo.TodoObject>}
 *     Promise that resolves to the response
 */
proto.todo.TodoServicePromiseClient.prototype.addTodo =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/todo.TodoService/addTodo',
      request,
      metadata || {},
      methodDescriptor_TodoService_addTodo);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.todo.DeleteTodoParams,
 *   !proto.todo.DeleteResponse>}
 */
const methodDescriptor_TodoService_deleteTodo = new grpc.web.MethodDescriptor(
  '/todo.TodoService/deleteTodo',
  grpc.web.MethodType.UNARY,
  proto.todo.DeleteTodoParams,
  proto.todo.DeleteResponse,
  /**
   * @param {!proto.todo.DeleteTodoParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.DeleteResponse.deserializeBinary
);


/**
 * @param {!proto.todo.DeleteTodoParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.todo.DeleteResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.todo.DeleteResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.todo.TodoServiceClient.prototype.deleteTodo =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/todo.TodoService/deleteTodo',
      request,
      metadata || {},
      methodDescriptor_TodoService_deleteTodo,
      callback);
};


/**
 * @param {!proto.todo.DeleteTodoParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.todo.DeleteResponse>}
 *     Promise that resolves to the response
 */
proto.todo.TodoServicePromiseClient.prototype.deleteTodo =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/todo.TodoService/deleteTodo',
      request,
      metadata || {},
      methodDescriptor_TodoService_deleteTodo);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.todo.GetTodoParams,
 *   !proto.todo.TodoResponse>}
 */
const methodDescriptor_TodoService_getTodos = new grpc.web.MethodDescriptor(
  '/todo.TodoService/getTodos',
  grpc.web.MethodType.UNARY,
  proto.todo.GetTodoParams,
  proto.todo.TodoResponse,
  /**
   * @param {!proto.todo.GetTodoParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.todo.TodoResponse.deserializeBinary
);


/**
 * @param {!proto.todo.GetTodoParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.todo.TodoResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.todo.TodoResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.todo.TodoServiceClient.prototype.getTodos =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/todo.TodoService/getTodos',
      request,
      metadata || {},
      methodDescriptor_TodoService_getTodos,
      callback);
};


/**
 * @param {!proto.todo.GetTodoParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.todo.TodoResponse>}
 *     Promise that resolves to the response
 */
proto.todo.TodoServicePromiseClient.prototype.getTodos =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/todo.TodoService/getTodos',
      request,
      metadata || {},
      methodDescriptor_TodoService_getTodos);
};


module.exports = proto.todo;

