/**
 * @fileoverview gRPC-Web generated client stub for cartpb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')
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
 *   !proto.cartpb.ShowRequest,
 *   !proto.cartpb.ShowResponse>}
 */
const methodDescriptor_CartAPI_Show = new grpc.web.MethodDescriptor(
  '/cartpb.CartAPI/Show',
  grpc.web.MethodType.UNARY,
  proto.cartpb.ShowRequest,
  proto.cartpb.ShowResponse,
  /**
   * @param {!proto.cartpb.ShowRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cartpb.ShowResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.cartpb.ShowRequest,
 *   !proto.cartpb.ShowResponse>}
 */
const methodInfo_CartAPI_Show = new grpc.web.AbstractClientBase.MethodInfo(
  proto.cartpb.ShowResponse,
  /**
   * @param {!proto.cartpb.ShowRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cartpb.ShowResponse.deserializeBinary
);


/**
 * @param {!proto.cartpb.ShowRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.cartpb.ShowResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.cartpb.ShowResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cartpb.CartAPIClient.prototype.show =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cartpb.CartAPI/Show',
      request,
      metadata || {},
      methodDescriptor_CartAPI_Show,
      callback);
};


/**
 * @param {!proto.cartpb.ShowRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.cartpb.ShowResponse>}
 *     A native promise that resolves to the response
 */
proto.cartpb.CartAPIPromiseClient.prototype.show =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cartpb.CartAPI/Show',
      request,
      metadata || {},
      methodDescriptor_CartAPI_Show);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cartpb.AddRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_CartAPI_Add = new grpc.web.MethodDescriptor(
  '/cartpb.CartAPI/Add',
  grpc.web.MethodType.UNARY,
  proto.cartpb.AddRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.cartpb.AddRequest} request
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
 *   !proto.cartpb.AddRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_CartAPI_Add = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.cartpb.AddRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.cartpb.AddRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cartpb.CartAPIClient.prototype.add =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cartpb.CartAPI/Add',
      request,
      metadata || {},
      methodDescriptor_CartAPI_Add,
      callback);
};


/**
 * @param {!proto.cartpb.AddRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.cartpb.CartAPIPromiseClient.prototype.add =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cartpb.CartAPI/Add',
      request,
      metadata || {},
      methodDescriptor_CartAPI_Add);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cartpb.RemoveRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_CartAPI_Remove = new grpc.web.MethodDescriptor(
  '/cartpb.CartAPI/Remove',
  grpc.web.MethodType.UNARY,
  proto.cartpb.RemoveRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.cartpb.RemoveRequest} request
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
 *   !proto.cartpb.RemoveRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_CartAPI_Remove = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.cartpb.RemoveRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.cartpb.RemoveRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cartpb.CartAPIClient.prototype.remove =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cartpb.CartAPI/Remove',
      request,
      metadata || {},
      methodDescriptor_CartAPI_Remove,
      callback);
};


/**
 * @param {!proto.cartpb.RemoveRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.cartpb.CartAPIPromiseClient.prototype.remove =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cartpb.CartAPI/Remove',
      request,
      metadata || {},
      methodDescriptor_CartAPI_Remove);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cartpb.CommitRequest,
 *   !proto.cartpb.CommitResponse>}
 */
const methodDescriptor_CartAPI_Commit = new grpc.web.MethodDescriptor(
  '/cartpb.CartAPI/Commit',
  grpc.web.MethodType.UNARY,
  proto.cartpb.CommitRequest,
  proto.cartpb.CommitResponse,
  /**
   * @param {!proto.cartpb.CommitRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cartpb.CommitResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.cartpb.CommitRequest,
 *   !proto.cartpb.CommitResponse>}
 */
const methodInfo_CartAPI_Commit = new grpc.web.AbstractClientBase.MethodInfo(
  proto.cartpb.CommitResponse,
  /**
   * @param {!proto.cartpb.CommitRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cartpb.CommitResponse.deserializeBinary
);


/**
 * @param {!proto.cartpb.CommitRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.cartpb.CommitResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.cartpb.CommitResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cartpb.CartAPIClient.prototype.commit =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cartpb.CartAPI/Commit',
      request,
      metadata || {},
      methodDescriptor_CartAPI_Commit,
      callback);
};


/**
 * @param {!proto.cartpb.CommitRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.cartpb.CommitResponse>}
 *     A native promise that resolves to the response
 */
proto.cartpb.CartAPIPromiseClient.prototype.commit =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cartpb.CartAPI/Commit',
      request,
      metadata || {},
      methodDescriptor_CartAPI_Commit);
};


module.exports = proto.cartpb;

