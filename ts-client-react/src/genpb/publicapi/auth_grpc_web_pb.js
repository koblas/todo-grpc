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
 *   !proto.auth.TokenRegister>}
 */
const methodDescriptor_AuthenticationService_register = new grpc.web.MethodDescriptor(
  '/auth.AuthenticationService/register',
  grpc.web.MethodType.UNARY,
  proto.auth.RegisterParams,
  proto.auth.TokenRegister,
  /**
   * @param {!proto.auth.RegisterParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.TokenRegister.deserializeBinary
);


/**
 * @param {!proto.auth.RegisterParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.auth.TokenRegister)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.auth.TokenRegister>|undefined}
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
 * @return {!Promise<!proto.auth.TokenRegister>}
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
 *   !proto.auth.Success>}
 */
const methodDescriptor_AuthenticationService_verify_email = new grpc.web.MethodDescriptor(
  '/auth.AuthenticationService/verify_email',
  grpc.web.MethodType.UNARY,
  proto.auth.ConfirmParams,
  proto.auth.Success,
  /**
   * @param {!proto.auth.ConfirmParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.Success.deserializeBinary
);


/**
 * @param {!proto.auth.ConfirmParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.auth.Success)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.auth.Success>|undefined}
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
 * @return {!Promise<!proto.auth.Success>}
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
 *   !proto.auth.RecoverySendParams,
 *   !proto.auth.Success>}
 */
const methodDescriptor_AuthenticationService_recover_send = new grpc.web.MethodDescriptor(
  '/auth.AuthenticationService/recover_send',
  grpc.web.MethodType.UNARY,
  proto.auth.RecoverySendParams,
  proto.auth.Success,
  /**
   * @param {!proto.auth.RecoverySendParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.Success.deserializeBinary
);


/**
 * @param {!proto.auth.RecoverySendParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.auth.Success)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.auth.Success>|undefined}
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
 * @param {!proto.auth.RecoverySendParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.auth.Success>}
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
 *   !proto.auth.RecoveryUpdateParams,
 *   !proto.auth.Success>}
 */
const methodDescriptor_AuthenticationService_recover_verify = new grpc.web.MethodDescriptor(
  '/auth.AuthenticationService/recover_verify',
  grpc.web.MethodType.UNARY,
  proto.auth.RecoveryUpdateParams,
  proto.auth.Success,
  /**
   * @param {!proto.auth.RecoveryUpdateParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.Success.deserializeBinary
);


/**
 * @param {!proto.auth.RecoveryUpdateParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.auth.Success)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.auth.Success>|undefined}
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
 * @param {!proto.auth.RecoveryUpdateParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.auth.Success>}
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
 *   !proto.auth.RecoveryUpdateParams,
 *   !proto.auth.Token>}
 */
const methodDescriptor_AuthenticationService_recover_update = new grpc.web.MethodDescriptor(
  '/auth.AuthenticationService/recover_update',
  grpc.web.MethodType.UNARY,
  proto.auth.RecoveryUpdateParams,
  proto.auth.Token,
  /**
   * @param {!proto.auth.RecoveryUpdateParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.Token.deserializeBinary
);


/**
 * @param {!proto.auth.RecoveryUpdateParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.auth.Token)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.auth.Token>|undefined}
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
 * @param {!proto.auth.RecoveryUpdateParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.auth.Token>}
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


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.auth.OauthAssociateParams,
 *   !proto.auth.TokenRegister>}
 */
const methodDescriptor_AuthenticationService_oauth_login = new grpc.web.MethodDescriptor(
  '/auth.AuthenticationService/oauth_login',
  grpc.web.MethodType.UNARY,
  proto.auth.OauthAssociateParams,
  proto.auth.TokenRegister,
  /**
   * @param {!proto.auth.OauthAssociateParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.TokenRegister.deserializeBinary
);


/**
 * @param {!proto.auth.OauthAssociateParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.auth.TokenRegister)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.auth.TokenRegister>|undefined}
 *     The XHR Node Readable Stream
 */
proto.auth.AuthenticationServiceClient.prototype.oauth_login =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/auth.AuthenticationService/oauth_login',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_oauth_login,
      callback);
};


/**
 * @param {!proto.auth.OauthAssociateParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.auth.TokenRegister>}
 *     Promise that resolves to the response
 */
proto.auth.AuthenticationServicePromiseClient.prototype.oauth_login =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/auth.AuthenticationService/oauth_login',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_oauth_login);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.auth.OauthUrlParams,
 *   !proto.auth.OauthUrlResult>}
 */
const methodDescriptor_AuthenticationService_oauth_url = new grpc.web.MethodDescriptor(
  '/auth.AuthenticationService/oauth_url',
  grpc.web.MethodType.UNARY,
  proto.auth.OauthUrlParams,
  proto.auth.OauthUrlResult,
  /**
   * @param {!proto.auth.OauthUrlParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.OauthUrlResult.deserializeBinary
);


/**
 * @param {!proto.auth.OauthUrlParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.auth.OauthUrlResult)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.auth.OauthUrlResult>|undefined}
 *     The XHR Node Readable Stream
 */
proto.auth.AuthenticationServiceClient.prototype.oauth_url =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/auth.AuthenticationService/oauth_url',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_oauth_url,
      callback);
};


/**
 * @param {!proto.auth.OauthUrlParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.auth.OauthUrlResult>}
 *     Promise that resolves to the response
 */
proto.auth.AuthenticationServicePromiseClient.prototype.oauth_url =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/auth.AuthenticationService/oauth_url',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_oauth_url);
};


module.exports = proto.auth;

