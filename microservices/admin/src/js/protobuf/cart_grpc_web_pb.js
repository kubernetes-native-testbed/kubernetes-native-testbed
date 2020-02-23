/**
 * @fileoverview gRPC-Web generated client stub for cartpb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.cartpb = require('./cart_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.cartpb.CartAPIClient =
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
proto.cartpb.CartAPIPromiseClient =
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
 *   !proto.cartpb.GetRequest,
 *   !proto.cartpb.GetResponse>}
 */
const methodDescriptor_CartAPI_Get = new grpc.web.MethodDescriptor(
  '/cartpb.CartAPI/Get',
  grpc.web.MethodType.UNARY,
  proto.cartpb.GetRequest,
  proto.cartpb.GetResponse,
  /**
   * @param {!proto.cartpb.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cartpb.GetResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.cartpb.GetRequest,
 *   !proto.cartpb.GetResponse>}
 */
const methodInfo_CartAPI_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.cartpb.GetResponse,
  /**
   * @param {!proto.cartpb.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cartpb.GetResponse.deserializeBinary
);


/**
 * @param {!proto.cartpb.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.cartpb.GetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.cartpb.GetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cartpb.CartAPIClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cartpb.CartAPI/Get',
      request,
      metadata || {},
      methodDescriptor_CartAPI_Get,
      callback);
};


/**
 * @param {!proto.cartpb.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.cartpb.GetResponse>}
 *     A native promise that resolves to the response
 */
proto.cartpb.CartAPIPromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cartpb.CartAPI/Get',
      request,
      metadata || {},
      methodDescriptor_CartAPI_Get);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cartpb.SetRequest,
 *   !proto.cartpb.SetResponse>}
 */
const methodDescriptor_CartAPI_Set = new grpc.web.MethodDescriptor(
  '/cartpb.CartAPI/Set',
  grpc.web.MethodType.UNARY,
  proto.cartpb.SetRequest,
  proto.cartpb.SetResponse,
  /**
   * @param {!proto.cartpb.SetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cartpb.SetResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.cartpb.SetRequest,
 *   !proto.cartpb.SetResponse>}
 */
const methodInfo_CartAPI_Set = new grpc.web.AbstractClientBase.MethodInfo(
  proto.cartpb.SetResponse,
  /**
   * @param {!proto.cartpb.SetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cartpb.SetResponse.deserializeBinary
);


/**
 * @param {!proto.cartpb.SetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.cartpb.SetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.cartpb.SetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cartpb.CartAPIClient.prototype.set =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cartpb.CartAPI/Set',
      request,
      metadata || {},
      methodDescriptor_CartAPI_Set,
      callback);
};


/**
 * @param {!proto.cartpb.SetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.cartpb.SetResponse>}
 *     A native promise that resolves to the response
 */
proto.cartpb.CartAPIPromiseClient.prototype.set =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cartpb.CartAPI/Set',
      request,
      metadata || {},
      methodDescriptor_CartAPI_Set);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cartpb.UpdateRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_CartAPI_Update = new grpc.web.MethodDescriptor(
  '/cartpb.CartAPI/Update',
  grpc.web.MethodType.UNARY,
  proto.cartpb.UpdateRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.cartpb.UpdateRequest} request
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
 *   !proto.cartpb.UpdateRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_CartAPI_Update = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.cartpb.UpdateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.cartpb.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cartpb.CartAPIClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cartpb.CartAPI/Update',
      request,
      metadata || {},
      methodDescriptor_CartAPI_Update,
      callback);
};


/**
 * @param {!proto.cartpb.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.cartpb.CartAPIPromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cartpb.CartAPI/Update',
      request,
      metadata || {},
      methodDescriptor_CartAPI_Update);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cartpb.DeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_CartAPI_Delete = new grpc.web.MethodDescriptor(
  '/cartpb.CartAPI/Delete',
  grpc.web.MethodType.UNARY,
  proto.cartpb.DeleteRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.cartpb.DeleteRequest} request
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
 *   !proto.cartpb.DeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_CartAPI_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.cartpb.DeleteRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.cartpb.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cartpb.CartAPIClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cartpb.CartAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_CartAPI_Delete,
      callback);
};


/**
 * @param {!proto.cartpb.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.cartpb.CartAPIPromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cartpb.CartAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_CartAPI_Delete);
};


module.exports = proto.cartpb;

