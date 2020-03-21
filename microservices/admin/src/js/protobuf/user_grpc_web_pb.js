/**
 * @fileoverview gRPC-Web generated client stub for userpb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.userpb = require('./user_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.userpb.UserAPIClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

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
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.userpb.UserAPIPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

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
 *   !proto.userpb.GetRequest,
 *   !proto.userpb.GetResponse>}
 */
const methodDescriptor_UserAPI_Get = new grpc.web.MethodDescriptor(
  '/userpb.UserAPI/Get',
  grpc.web.MethodType.UNARY,
  proto.userpb.GetRequest,
  proto.userpb.GetResponse,
  /**
   * @param {!proto.userpb.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.userpb.GetResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.userpb.GetRequest,
 *   !proto.userpb.GetResponse>}
 */
const methodInfo_UserAPI_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.userpb.GetResponse,
  /**
   * @param {!proto.userpb.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.userpb.GetResponse.deserializeBinary
);


/**
 * @param {!proto.userpb.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.userpb.GetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.userpb.GetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.userpb.UserAPIClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/userpb.UserAPI/Get',
      request,
      metadata || {},
      methodDescriptor_UserAPI_Get,
      callback);
};


/**
 * @param {!proto.userpb.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.userpb.GetResponse>}
 *     A native promise that resolves to the response
 */
proto.userpb.UserAPIPromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/userpb.UserAPI/Get',
      request,
      metadata || {},
      methodDescriptor_UserAPI_Get);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.userpb.SetRequest,
 *   !proto.userpb.SetResponse>}
 */
const methodDescriptor_UserAPI_Set = new grpc.web.MethodDescriptor(
  '/userpb.UserAPI/Set',
  grpc.web.MethodType.UNARY,
  proto.userpb.SetRequest,
  proto.userpb.SetResponse,
  /**
   * @param {!proto.userpb.SetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.userpb.SetResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.userpb.SetRequest,
 *   !proto.userpb.SetResponse>}
 */
const methodInfo_UserAPI_Set = new grpc.web.AbstractClientBase.MethodInfo(
  proto.userpb.SetResponse,
  /**
   * @param {!proto.userpb.SetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.userpb.SetResponse.deserializeBinary
);


/**
 * @param {!proto.userpb.SetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.userpb.SetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.userpb.SetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.userpb.UserAPIClient.prototype.set =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/userpb.UserAPI/Set',
      request,
      metadata || {},
      methodDescriptor_UserAPI_Set,
      callback);
};


/**
 * @param {!proto.userpb.SetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.userpb.SetResponse>}
 *     A native promise that resolves to the response
 */
proto.userpb.UserAPIPromiseClient.prototype.set =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/userpb.UserAPI/Set',
      request,
      metadata || {},
      methodDescriptor_UserAPI_Set);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.userpb.UpdateRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_UserAPI_Update = new grpc.web.MethodDescriptor(
  '/userpb.UserAPI/Update',
  grpc.web.MethodType.UNARY,
  proto.userpb.UpdateRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.userpb.UpdateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.userpb.UpdateRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_UserAPI_Update = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.userpb.UpdateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.userpb.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.userpb.UserAPIClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/userpb.UserAPI/Update',
      request,
      metadata || {},
      methodDescriptor_UserAPI_Update,
      callback);
};


/**
 * @param {!proto.userpb.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.userpb.UserAPIPromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/userpb.UserAPI/Update',
      request,
      metadata || {},
      methodDescriptor_UserAPI_Update);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.userpb.DeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_UserAPI_Delete = new grpc.web.MethodDescriptor(
  '/userpb.UserAPI/Delete',
  grpc.web.MethodType.UNARY,
  proto.userpb.DeleteRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.userpb.DeleteRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.userpb.DeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_UserAPI_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.userpb.DeleteRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.userpb.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.userpb.UserAPIClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/userpb.UserAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_UserAPI_Delete,
      callback);
};


/**
 * @param {!proto.userpb.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.userpb.UserAPIPromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/userpb.UserAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_UserAPI_Delete);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.userpb.IsExistsRequest,
 *   !proto.userpb.IsExistsResponse>}
 */
const methodDescriptor_UserAPI_IsExists = new grpc.web.MethodDescriptor(
  '/userpb.UserAPI/IsExists',
  grpc.web.MethodType.UNARY,
  proto.userpb.IsExistsRequest,
  proto.userpb.IsExistsResponse,
  /**
   * @param {!proto.userpb.IsExistsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.userpb.IsExistsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.userpb.IsExistsRequest,
 *   !proto.userpb.IsExistsResponse>}
 */
const methodInfo_UserAPI_IsExists = new grpc.web.AbstractClientBase.MethodInfo(
  proto.userpb.IsExistsResponse,
  /**
   * @param {!proto.userpb.IsExistsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.userpb.IsExistsResponse.deserializeBinary
);


/**
 * @param {!proto.userpb.IsExistsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.userpb.IsExistsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.userpb.IsExistsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.userpb.UserAPIClient.prototype.isExists =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/userpb.UserAPI/IsExists',
      request,
      metadata || {},
      methodDescriptor_UserAPI_IsExists,
      callback);
};


/**
 * @param {!proto.userpb.IsExistsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.userpb.IsExistsResponse>}
 *     A native promise that resolves to the response
 */
proto.userpb.UserAPIPromiseClient.prototype.isExists =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/userpb.UserAPI/IsExists',
      request,
      metadata || {},
      methodDescriptor_UserAPI_IsExists);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.userpb.AuthenticationRequest,
 *   !proto.userpb.AuthenticationResponse>}
 */
const methodDescriptor_UserAPI_Authentication = new grpc.web.MethodDescriptor(
  '/userpb.UserAPI/Authentication',
  grpc.web.MethodType.UNARY,
  proto.userpb.AuthenticationRequest,
  proto.userpb.AuthenticationResponse,
  /**
   * @param {!proto.userpb.AuthenticationRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.userpb.AuthenticationResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.userpb.AuthenticationRequest,
 *   !proto.userpb.AuthenticationResponse>}
 */
const methodInfo_UserAPI_Authentication = new grpc.web.AbstractClientBase.MethodInfo(
  proto.userpb.AuthenticationResponse,
  /**
   * @param {!proto.userpb.AuthenticationRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.userpb.AuthenticationResponse.deserializeBinary
);


/**
 * @param {!proto.userpb.AuthenticationRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.userpb.AuthenticationResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.userpb.AuthenticationResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.userpb.UserAPIClient.prototype.authentication =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/userpb.UserAPI/Authentication',
      request,
      metadata || {},
      methodDescriptor_UserAPI_Authentication,
      callback);
};


/**
 * @param {!proto.userpb.AuthenticationRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.userpb.AuthenticationResponse>}
 *     A native promise that resolves to the response
 */
proto.userpb.UserAPIPromiseClient.prototype.authentication =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/userpb.UserAPI/Authentication',
      request,
      metadata || {},
      methodDescriptor_UserAPI_Authentication);
};


module.exports = proto.userpb;

