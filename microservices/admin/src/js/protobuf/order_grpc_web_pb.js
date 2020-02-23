/**
 * @fileoverview gRPC-Web generated client stub for orderpb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.orderpb = require('./order_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.orderpb.OrderAPIClient =
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
proto.orderpb.OrderAPIPromiseClient =
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
 *   !proto.orderpb.GetRequest,
 *   !proto.orderpb.GetResponse>}
 */
const methodDescriptor_OrderAPI_Get = new grpc.web.MethodDescriptor(
  '/orderpb.OrderAPI/Get',
  grpc.web.MethodType.UNARY,
  proto.orderpb.GetRequest,
  proto.orderpb.GetResponse,
  /**
   * @param {!proto.orderpb.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.orderpb.GetResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.orderpb.GetRequest,
 *   !proto.orderpb.GetResponse>}
 */
const methodInfo_OrderAPI_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.orderpb.GetResponse,
  /**
   * @param {!proto.orderpb.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.orderpb.GetResponse.deserializeBinary
);


/**
 * @param {!proto.orderpb.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.orderpb.GetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.orderpb.GetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.orderpb.OrderAPIClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/orderpb.OrderAPI/Get',
      request,
      metadata || {},
      methodDescriptor_OrderAPI_Get,
      callback);
};


/**
 * @param {!proto.orderpb.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.orderpb.GetResponse>}
 *     A native promise that resolves to the response
 */
proto.orderpb.OrderAPIPromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/orderpb.OrderAPI/Get',
      request,
      metadata || {},
      methodDescriptor_OrderAPI_Get);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.orderpb.SetRequest,
 *   !proto.orderpb.SetResponse>}
 */
const methodDescriptor_OrderAPI_Set = new grpc.web.MethodDescriptor(
  '/orderpb.OrderAPI/Set',
  grpc.web.MethodType.UNARY,
  proto.orderpb.SetRequest,
  proto.orderpb.SetResponse,
  /**
   * @param {!proto.orderpb.SetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.orderpb.SetResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.orderpb.SetRequest,
 *   !proto.orderpb.SetResponse>}
 */
const methodInfo_OrderAPI_Set = new grpc.web.AbstractClientBase.MethodInfo(
  proto.orderpb.SetResponse,
  /**
   * @param {!proto.orderpb.SetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.orderpb.SetResponse.deserializeBinary
);


/**
 * @param {!proto.orderpb.SetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.orderpb.SetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.orderpb.SetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.orderpb.OrderAPIClient.prototype.set =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/orderpb.OrderAPI/Set',
      request,
      metadata || {},
      methodDescriptor_OrderAPI_Set,
      callback);
};


/**
 * @param {!proto.orderpb.SetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.orderpb.SetResponse>}
 *     A native promise that resolves to the response
 */
proto.orderpb.OrderAPIPromiseClient.prototype.set =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/orderpb.OrderAPI/Set',
      request,
      metadata || {},
      methodDescriptor_OrderAPI_Set);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.orderpb.UpdateRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_OrderAPI_Update = new grpc.web.MethodDescriptor(
  '/orderpb.OrderAPI/Update',
  grpc.web.MethodType.UNARY,
  proto.orderpb.UpdateRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.orderpb.UpdateRequest} request
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
 *   !proto.orderpb.UpdateRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_OrderAPI_Update = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.orderpb.UpdateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.orderpb.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.orderpb.OrderAPIClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/orderpb.OrderAPI/Update',
      request,
      metadata || {},
      methodDescriptor_OrderAPI_Update,
      callback);
};


/**
 * @param {!proto.orderpb.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.orderpb.OrderAPIPromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/orderpb.OrderAPI/Update',
      request,
      metadata || {},
      methodDescriptor_OrderAPI_Update);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.orderpb.DeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_OrderAPI_Delete = new grpc.web.MethodDescriptor(
  '/orderpb.OrderAPI/Delete',
  grpc.web.MethodType.UNARY,
  proto.orderpb.DeleteRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.orderpb.DeleteRequest} request
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
 *   !proto.orderpb.DeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_OrderAPI_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.orderpb.DeleteRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.orderpb.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.orderpb.OrderAPIClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/orderpb.OrderAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_OrderAPI_Delete,
      callback);
};


/**
 * @param {!proto.orderpb.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.orderpb.OrderAPIPromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/orderpb.OrderAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_OrderAPI_Delete);
};


module.exports = proto.orderpb;

