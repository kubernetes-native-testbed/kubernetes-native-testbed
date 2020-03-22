/**
 * @fileoverview gRPC-Web generated client stub for pointpb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.pointpb = require('./point_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.pointpb.PointAPIClient =
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
proto.pointpb.PointAPIPromiseClient =
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
 *   !proto.pointpb.GetRequest,
 *   !proto.pointpb.GetResponse>}
 */
const methodDescriptor_PointAPI_Get = new grpc.web.MethodDescriptor(
  '/pointpb.PointAPI/Get',
  grpc.web.MethodType.UNARY,
  proto.pointpb.GetRequest,
  proto.pointpb.GetResponse,
  /**
   * @param {!proto.pointpb.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pointpb.GetResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.pointpb.GetRequest,
 *   !proto.pointpb.GetResponse>}
 */
const methodInfo_PointAPI_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.pointpb.GetResponse,
  /**
   * @param {!proto.pointpb.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pointpb.GetResponse.deserializeBinary
);


/**
 * @param {!proto.pointpb.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.pointpb.GetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pointpb.GetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pointpb.PointAPIClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pointpb.PointAPI/Get',
      request,
      metadata || {},
      methodDescriptor_PointAPI_Get,
      callback);
};


/**
 * @param {!proto.pointpb.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pointpb.GetResponse>}
 *     A native promise that resolves to the response
 */
proto.pointpb.PointAPIPromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pointpb.PointAPI/Get',
      request,
      metadata || {},
      methodDescriptor_PointAPI_Get);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pointpb.SetRequest,
 *   !proto.pointpb.SetResponse>}
 */
const methodDescriptor_PointAPI_Set = new grpc.web.MethodDescriptor(
  '/pointpb.PointAPI/Set',
  grpc.web.MethodType.UNARY,
  proto.pointpb.SetRequest,
  proto.pointpb.SetResponse,
  /**
   * @param {!proto.pointpb.SetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pointpb.SetResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.pointpb.SetRequest,
 *   !proto.pointpb.SetResponse>}
 */
const methodInfo_PointAPI_Set = new grpc.web.AbstractClientBase.MethodInfo(
  proto.pointpb.SetResponse,
  /**
   * @param {!proto.pointpb.SetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pointpb.SetResponse.deserializeBinary
);


/**
 * @param {!proto.pointpb.SetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.pointpb.SetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pointpb.SetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pointpb.PointAPIClient.prototype.set =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pointpb.PointAPI/Set',
      request,
      metadata || {},
      methodDescriptor_PointAPI_Set,
      callback);
};


/**
 * @param {!proto.pointpb.SetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pointpb.SetResponse>}
 *     A native promise that resolves to the response
 */
proto.pointpb.PointAPIPromiseClient.prototype.set =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pointpb.PointAPI/Set',
      request,
      metadata || {},
      methodDescriptor_PointAPI_Set);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pointpb.UpdateRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_PointAPI_Update = new grpc.web.MethodDescriptor(
  '/pointpb.PointAPI/Update',
  grpc.web.MethodType.UNARY,
  proto.pointpb.UpdateRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.pointpb.UpdateRequest} request
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
 *   !proto.pointpb.UpdateRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_PointAPI_Update = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.pointpb.UpdateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.pointpb.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pointpb.PointAPIClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pointpb.PointAPI/Update',
      request,
      metadata || {},
      methodDescriptor_PointAPI_Update,
      callback);
};


/**
 * @param {!proto.pointpb.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.pointpb.PointAPIPromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pointpb.PointAPI/Update',
      request,
      metadata || {},
      methodDescriptor_PointAPI_Update);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pointpb.DeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_PointAPI_Delete = new grpc.web.MethodDescriptor(
  '/pointpb.PointAPI/Delete',
  grpc.web.MethodType.UNARY,
  proto.pointpb.DeleteRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.pointpb.DeleteRequest} request
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
 *   !proto.pointpb.DeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_PointAPI_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.pointpb.DeleteRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.pointpb.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pointpb.PointAPIClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pointpb.PointAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_PointAPI_Delete,
      callback);
};


/**
 * @param {!proto.pointpb.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.pointpb.PointAPIPromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pointpb.PointAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_PointAPI_Delete);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pointpb.GetTotalAmountRequest,
 *   !proto.pointpb.GetTotalAmountResponse>}
 */
const methodDescriptor_PointAPI_GetTotalAmount = new grpc.web.MethodDescriptor(
  '/pointpb.PointAPI/GetTotalAmount',
  grpc.web.MethodType.UNARY,
  proto.pointpb.GetTotalAmountRequest,
  proto.pointpb.GetTotalAmountResponse,
  /**
   * @param {!proto.pointpb.GetTotalAmountRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pointpb.GetTotalAmountResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.pointpb.GetTotalAmountRequest,
 *   !proto.pointpb.GetTotalAmountResponse>}
 */
const methodInfo_PointAPI_GetTotalAmount = new grpc.web.AbstractClientBase.MethodInfo(
  proto.pointpb.GetTotalAmountResponse,
  /**
   * @param {!proto.pointpb.GetTotalAmountRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pointpb.GetTotalAmountResponse.deserializeBinary
);


/**
 * @param {!proto.pointpb.GetTotalAmountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.pointpb.GetTotalAmountResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pointpb.GetTotalAmountResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pointpb.PointAPIClient.prototype.getTotalAmount =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pointpb.PointAPI/GetTotalAmount',
      request,
      metadata || {},
      methodDescriptor_PointAPI_GetTotalAmount,
      callback);
};


/**
 * @param {!proto.pointpb.GetTotalAmountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pointpb.GetTotalAmountResponse>}
 *     A native promise that resolves to the response
 */
proto.pointpb.PointAPIPromiseClient.prototype.getTotalAmount =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pointpb.PointAPI/GetTotalAmount',
      request,
      metadata || {},
      methodDescriptor_PointAPI_GetTotalAmount);
};


module.exports = proto.pointpb;

