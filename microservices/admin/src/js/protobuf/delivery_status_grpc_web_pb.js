/**
 * @fileoverview gRPC-Web generated client stub for deliverystatuspb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.deliverystatuspb = require('./delivery_status_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.deliverystatuspb.DeliveryStatusAPIClient =
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
proto.deliverystatuspb.DeliveryStatusAPIPromiseClient =
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
 *   !proto.deliverystatuspb.GetRequest,
 *   !proto.deliverystatuspb.GetResponse>}
 */
const methodDescriptor_DeliveryStatusAPI_Get = new grpc.web.MethodDescriptor(
  '/deliverystatuspb.DeliveryStatusAPI/Get',
  grpc.web.MethodType.UNARY,
  proto.deliverystatuspb.GetRequest,
  proto.deliverystatuspb.GetResponse,
  /**
   * @param {!proto.deliverystatuspb.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.deliverystatuspb.GetResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.deliverystatuspb.GetRequest,
 *   !proto.deliverystatuspb.GetResponse>}
 */
const methodInfo_DeliveryStatusAPI_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.deliverystatuspb.GetResponse,
  /**
   * @param {!proto.deliverystatuspb.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.deliverystatuspb.GetResponse.deserializeBinary
);


/**
 * @param {!proto.deliverystatuspb.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.deliverystatuspb.GetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.deliverystatuspb.GetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.deliverystatuspb.DeliveryStatusAPIClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/deliverystatuspb.DeliveryStatusAPI/Get',
      request,
      metadata || {},
      methodDescriptor_DeliveryStatusAPI_Get,
      callback);
};


/**
 * @param {!proto.deliverystatuspb.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.deliverystatuspb.GetResponse>}
 *     A native promise that resolves to the response
 */
proto.deliverystatuspb.DeliveryStatusAPIPromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/deliverystatuspb.DeliveryStatusAPI/Get',
      request,
      metadata || {},
      methodDescriptor_DeliveryStatusAPI_Get);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.deliverystatuspb.SetRequest,
 *   !proto.deliverystatuspb.SetResponse>}
 */
const methodDescriptor_DeliveryStatusAPI_Set = new grpc.web.MethodDescriptor(
  '/deliverystatuspb.DeliveryStatusAPI/Set',
  grpc.web.MethodType.UNARY,
  proto.deliverystatuspb.SetRequest,
  proto.deliverystatuspb.SetResponse,
  /**
   * @param {!proto.deliverystatuspb.SetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.deliverystatuspb.SetResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.deliverystatuspb.SetRequest,
 *   !proto.deliverystatuspb.SetResponse>}
 */
const methodInfo_DeliveryStatusAPI_Set = new grpc.web.AbstractClientBase.MethodInfo(
  proto.deliverystatuspb.SetResponse,
  /**
   * @param {!proto.deliverystatuspb.SetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.deliverystatuspb.SetResponse.deserializeBinary
);


/**
 * @param {!proto.deliverystatuspb.SetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.deliverystatuspb.SetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.deliverystatuspb.SetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.deliverystatuspb.DeliveryStatusAPIClient.prototype.set =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/deliverystatuspb.DeliveryStatusAPI/Set',
      request,
      metadata || {},
      methodDescriptor_DeliveryStatusAPI_Set,
      callback);
};


/**
 * @param {!proto.deliverystatuspb.SetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.deliverystatuspb.SetResponse>}
 *     A native promise that resolves to the response
 */
proto.deliverystatuspb.DeliveryStatusAPIPromiseClient.prototype.set =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/deliverystatuspb.DeliveryStatusAPI/Set',
      request,
      metadata || {},
      methodDescriptor_DeliveryStatusAPI_Set);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.deliverystatuspb.UpdateRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_DeliveryStatusAPI_Update = new grpc.web.MethodDescriptor(
  '/deliverystatuspb.DeliveryStatusAPI/Update',
  grpc.web.MethodType.UNARY,
  proto.deliverystatuspb.UpdateRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.deliverystatuspb.UpdateRequest} request
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
 *   !proto.deliverystatuspb.UpdateRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_DeliveryStatusAPI_Update = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.deliverystatuspb.UpdateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.deliverystatuspb.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.deliverystatuspb.DeliveryStatusAPIClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/deliverystatuspb.DeliveryStatusAPI/Update',
      request,
      metadata || {},
      methodDescriptor_DeliveryStatusAPI_Update,
      callback);
};


/**
 * @param {!proto.deliverystatuspb.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.deliverystatuspb.DeliveryStatusAPIPromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/deliverystatuspb.DeliveryStatusAPI/Update',
      request,
      metadata || {},
      methodDescriptor_DeliveryStatusAPI_Update);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.deliverystatuspb.DeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_DeliveryStatusAPI_Delete = new grpc.web.MethodDescriptor(
  '/deliverystatuspb.DeliveryStatusAPI/Delete',
  grpc.web.MethodType.UNARY,
  proto.deliverystatuspb.DeleteRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.deliverystatuspb.DeleteRequest} request
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
 *   !proto.deliverystatuspb.DeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_DeliveryStatusAPI_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.deliverystatuspb.DeleteRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.deliverystatuspb.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.deliverystatuspb.DeliveryStatusAPIClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/deliverystatuspb.DeliveryStatusAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_DeliveryStatusAPI_Delete,
      callback);
};


/**
 * @param {!proto.deliverystatuspb.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.deliverystatuspb.DeliveryStatusAPIPromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/deliverystatuspb.DeliveryStatusAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_DeliveryStatusAPI_Delete);
};


module.exports = proto.deliverystatuspb;

