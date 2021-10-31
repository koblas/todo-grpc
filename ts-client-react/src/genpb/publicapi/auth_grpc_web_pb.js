/**
 * @fileoverview gRPC-Web generated client stub for auth
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
proto.auth = require('./auth_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.auth.AuthenticationServiceClient =
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
proto.auth.AuthenticationServicePromiseClient =
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
 *   !proto.auth.RegisterParams,
 *   !proto.auth.Token>}
 */
const methodDescriptor_AuthenticationService_register = new grpc.web.MethodDescriptor(
  '/auth.AuthenticationService/register',
  grpc.web.MethodType.UNARY,
  proto.auth.RegisterParams,
  proto.auth.Token,
  /**
   * @param {!proto.auth.RegisterParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.Token.deserializeBinary
);


/**
 * @param {!proto.auth.RegisterParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.auth.Token)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.auth.Token>|undefined}
 *     The XHR Node Readable Stream
 */
proto.auth.AuthenticationServiceClient.prototype.register =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/auth.AuthenticationService/register',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_register,
      callback);
};


/**
 * @param {!proto.auth.RegisterParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.auth.Token>}
 *     Promise that resolves to the response
 */
proto.auth.AuthenticationServicePromiseClient.prototype.register =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/auth.AuthenticationService/register',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_register);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.auth.LoginParams,
 *   !proto.auth.Token>}
 */
const methodDescriptor_AuthenticationService_authenticate = new grpc.web.MethodDescriptor(
  '/auth.AuthenticationService/authenticate',
  grpc.web.MethodType.UNARY,
  proto.auth.LoginParams,
  proto.auth.Token,
  /**
   * @param {!proto.auth.LoginParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.Token.deserializeBinary
);


/**
 * @param {!proto.auth.LoginParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.auth.Token)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.auth.Token>|undefined}
 *     The XHR Node Readable Stream
 */
proto.auth.AuthenticationServiceClient.prototype.authenticate =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/auth.AuthenticationService/authenticate',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_authenticate,
      callback);
};


/**
 * @param {!proto.auth.LoginParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.auth.Token>}
 *     Promise that resolves to the response
 */
proto.auth.AuthenticationServicePromiseClient.prototype.authenticate =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/auth.AuthenticationService/authenticate',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_authenticate);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.auth.ConfirmParams,
 *   !proto.auth.TokenEither>}
 */
const methodDescriptor_AuthenticationService_verify_email = new grpc.web.MethodDescriptor(
  '/auth.AuthenticationService/verify_email',
  grpc.web.MethodType.UNARY,
  proto.auth.ConfirmParams,
  proto.auth.TokenEither,
  /**
   * @param {!proto.auth.ConfirmParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.TokenEither.deserializeBinary
);


/**
 * @param {!proto.auth.ConfirmParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.auth.TokenEither)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.auth.TokenEither>|undefined}
 *     The XHR Node Readable Stream
 */
proto.auth.AuthenticationServiceClient.prototype.verify_email =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/auth.AuthenticationService/verify_email',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_verify_email,
      callback);
};


/**
 * @param {!proto.auth.ConfirmParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.auth.TokenEither>}
 *     Promise that resolves to the response
 */
proto.auth.AuthenticationServicePromiseClient.prototype.verify_email =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/auth.AuthenticationService/verify_email',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_verify_email);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.auth.RecoveryParams,
 *   !proto.auth.SuccessEither>}
 */
const methodDescriptor_AuthenticationService_recover_send = new grpc.web.MethodDescriptor(
  '/auth.AuthenticationService/recover_send',
  grpc.web.MethodType.UNARY,
  proto.auth.RecoveryParams,
  proto.auth.SuccessEither,
  /**
   * @param {!proto.auth.RecoveryParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.SuccessEither.deserializeBinary
);


/**
 * @param {!proto.auth.RecoveryParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.auth.SuccessEither)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.auth.SuccessEither>|undefined}
 *     The XHR Node Readable Stream
 */
proto.auth.AuthenticationServiceClient.prototype.recover_send =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/auth.AuthenticationService/recover_send',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_recover_send,
      callback);
};


/**
 * @param {!proto.auth.RecoveryParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.auth.SuccessEither>}
 *     Promise that resolves to the response
 */
proto.auth.AuthenticationServicePromiseClient.prototype.recover_send =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/auth.AuthenticationService/recover_send',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_recover_send);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.auth.RecoveryParams,
 *   !proto.auth.SuccessEither>}
 */
const methodDescriptor_AuthenticationService_recover_verify = new grpc.web.MethodDescriptor(
  '/auth.AuthenticationService/recover_verify',
  grpc.web.MethodType.UNARY,
  proto.auth.RecoveryParams,
  proto.auth.SuccessEither,
  /**
   * @param {!proto.auth.RecoveryParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.SuccessEither.deserializeBinary
);


/**
 * @param {!proto.auth.RecoveryParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.auth.SuccessEither)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.auth.SuccessEither>|undefined}
 *     The XHR Node Readable Stream
 */
proto.auth.AuthenticationServiceClient.prototype.recover_verify =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/auth.AuthenticationService/recover_verify',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_recover_verify,
      callback);
};


/**
 * @param {!proto.auth.RecoveryParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.auth.SuccessEither>}
 *     Promise that resolves to the response
 */
proto.auth.AuthenticationServicePromiseClient.prototype.recover_verify =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/auth.AuthenticationService/recover_verify',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_recover_verify);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.auth.RecoveryParams,
 *   !proto.auth.TokenEither>}
 */
const methodDescriptor_AuthenticationService_recover_update = new grpc.web.MethodDescriptor(
  '/auth.AuthenticationService/recover_update',
  grpc.web.MethodType.UNARY,
  proto.auth.RecoveryParams,
  proto.auth.TokenEither,
  /**
   * @param {!proto.auth.RecoveryParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.TokenEither.deserializeBinary
);


/**
 * @param {!proto.auth.RecoveryParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.auth.TokenEither)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.auth.TokenEither>|undefined}
 *     The XHR Node Readable Stream
 */
proto.auth.AuthenticationServiceClient.prototype.recover_update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/auth.AuthenticationService/recover_update',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_recover_update,
      callback);
};


/**
 * @param {!proto.auth.RecoveryParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.auth.TokenEither>}
 *     Promise that resolves to the response
 */
proto.auth.AuthenticationServicePromiseClient.prototype.recover_update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/auth.AuthenticationService/recover_update',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_recover_update);
};


module.exports = proto.auth;

