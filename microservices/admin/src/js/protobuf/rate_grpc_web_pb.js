/**
 * @fileoverview gRPC-Web generated client stub for ratepb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.ratepb = require('./rate_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.ratepb.RateAPIClient =
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
proto.ratepb.RateAPIPromiseClient =
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
 *   !proto.ratepb.GetRequest,
 *   !proto.ratepb.GetResponse>}
 */
const methodDescriptor_RateAPI_Get = new grpc.web.MethodDescriptor(
  '/ratepb.RateAPI/Get',
  grpc.web.MethodType.UNARY,
  proto.ratepb.GetRequest,
  proto.ratepb.GetResponse,
  /**
   * @param {!proto.ratepb.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ratepb.GetResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ratepb.GetRequest,
 *   !proto.ratepb.GetResponse>}
 */
const methodInfo_RateAPI_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.ratepb.GetResponse,
  /**
   * @param {!proto.ratepb.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ratepb.GetResponse.deserializeBinary
);


/**
 * @param {!proto.ratepb.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.ratepb.GetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.ratepb.GetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ratepb.RateAPIClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ratepb.RateAPI/Get',
      request,
      metadata || {},
      methodDescriptor_RateAPI_Get,
      callback);
};


/**
 * @param {!proto.ratepb.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.ratepb.GetResponse>}
 *     A native promise that resolves to the response
 */
proto.ratepb.RateAPIPromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ratepb.RateAPI/Get',
      request,
      metadata || {},
      methodDescriptor_RateAPI_Get);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.ratepb.SetRequest,
 *   !proto.ratepb.SetResponse>}
 */
const methodDescriptor_RateAPI_Set = new grpc.web.MethodDescriptor(
  '/ratepb.RateAPI/Set',
  grpc.web.MethodType.UNARY,
  proto.ratepb.SetRequest,
  proto.ratepb.SetResponse,
  /**
   * @param {!proto.ratepb.SetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ratepb.SetResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ratepb.SetRequest,
 *   !proto.ratepb.SetResponse>}
 */
const methodInfo_RateAPI_Set = new grpc.web.AbstractClientBase.MethodInfo(
  proto.ratepb.SetResponse,
  /**
   * @param {!proto.ratepb.SetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ratepb.SetResponse.deserializeBinary
);


/**
 * @param {!proto.ratepb.SetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.ratepb.SetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.ratepb.SetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ratepb.RateAPIClient.prototype.set =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ratepb.RateAPI/Set',
      request,
      metadata || {},
      methodDescriptor_RateAPI_Set,
      callback);
};


/**
 * @param {!proto.ratepb.SetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.ratepb.SetResponse>}
 *     A native promise that resolves to the response
 */
proto.ratepb.RateAPIPromiseClient.prototype.set =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ratepb.RateAPI/Set',
      request,
      metadata || {},
      methodDescriptor_RateAPI_Set);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.ratepb.UpdateRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_RateAPI_Update = new grpc.web.MethodDescriptor(
  '/ratepb.RateAPI/Update',
  grpc.web.MethodType.UNARY,
  proto.ratepb.UpdateRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.ratepb.UpdateRequest} request
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
 *   !proto.ratepb.UpdateRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_RateAPI_Update = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.ratepb.UpdateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.ratepb.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ratepb.RateAPIClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ratepb.RateAPI/Update',
      request,
      metadata || {},
      methodDescriptor_RateAPI_Update,
      callback);
};


/**
 * @param {!proto.ratepb.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.ratepb.RateAPIPromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ratepb.RateAPI/Update',
      request,
      metadata || {},
      methodDescriptor_RateAPI_Update);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.ratepb.DeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_RateAPI_Delete = new grpc.web.MethodDescriptor(
  '/ratepb.RateAPI/Delete',
  grpc.web.MethodType.UNARY,
  proto.ratepb.DeleteRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.ratepb.DeleteRequest} request
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
 *   !proto.ratepb.DeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_RateAPI_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.ratepb.DeleteRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.ratepb.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ratepb.RateAPIClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ratepb.RateAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_RateAPI_Delete,
      callback);
};


/**
 * @param {!proto.ratepb.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.ratepb.RateAPIPromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ratepb.RateAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_RateAPI_Delete);
};


module.exports = proto.ratepb;

