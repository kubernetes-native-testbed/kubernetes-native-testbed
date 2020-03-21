/**
 * @fileoverview gRPC-Web generated client stub for productpb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.productpb = require('./product_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.productpb.ProductAPIClient =
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
proto.productpb.ProductAPIPromiseClient =
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
 *   !proto.productpb.GetRequest,
 *   !proto.productpb.GetResponse>}
 */
const methodDescriptor_ProductAPI_Get = new grpc.web.MethodDescriptor(
  '/productpb.ProductAPI/Get',
  grpc.web.MethodType.UNARY,
  proto.productpb.GetRequest,
  proto.productpb.GetResponse,
  /**
   * @param {!proto.productpb.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.productpb.GetResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.productpb.GetRequest,
 *   !proto.productpb.GetResponse>}
 */
const methodInfo_ProductAPI_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.productpb.GetResponse,
  /**
   * @param {!proto.productpb.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.productpb.GetResponse.deserializeBinary
);


/**
 * @param {!proto.productpb.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.productpb.GetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.productpb.GetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.productpb.ProductAPIClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/productpb.ProductAPI/Get',
      request,
      metadata || {},
      methodDescriptor_ProductAPI_Get,
      callback);
};


/**
 * @param {!proto.productpb.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.productpb.GetResponse>}
 *     A native promise that resolves to the response
 */
proto.productpb.ProductAPIPromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/productpb.ProductAPI/Get',
      request,
      metadata || {},
      methodDescriptor_ProductAPI_Get);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.productpb.SetRequest,
 *   !proto.productpb.SetResponse>}
 */
const methodDescriptor_ProductAPI_Set = new grpc.web.MethodDescriptor(
  '/productpb.ProductAPI/Set',
  grpc.web.MethodType.UNARY,
  proto.productpb.SetRequest,
  proto.productpb.SetResponse,
  /**
   * @param {!proto.productpb.SetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.productpb.SetResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.productpb.SetRequest,
 *   !proto.productpb.SetResponse>}
 */
const methodInfo_ProductAPI_Set = new grpc.web.AbstractClientBase.MethodInfo(
  proto.productpb.SetResponse,
  /**
   * @param {!proto.productpb.SetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.productpb.SetResponse.deserializeBinary
);


/**
 * @param {!proto.productpb.SetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.productpb.SetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.productpb.SetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.productpb.ProductAPIClient.prototype.set =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/productpb.ProductAPI/Set',
      request,
      metadata || {},
      methodDescriptor_ProductAPI_Set,
      callback);
};


/**
 * @param {!proto.productpb.SetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.productpb.SetResponse>}
 *     A native promise that resolves to the response
 */
proto.productpb.ProductAPIPromiseClient.prototype.set =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/productpb.ProductAPI/Set',
      request,
      metadata || {},
      methodDescriptor_ProductAPI_Set);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.productpb.UpdateRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_ProductAPI_Update = new grpc.web.MethodDescriptor(
  '/productpb.ProductAPI/Update',
  grpc.web.MethodType.UNARY,
  proto.productpb.UpdateRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.productpb.UpdateRequest} request
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
 *   !proto.productpb.UpdateRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_ProductAPI_Update = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.productpb.UpdateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.productpb.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.productpb.ProductAPIClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/productpb.ProductAPI/Update',
      request,
      metadata || {},
      methodDescriptor_ProductAPI_Update,
      callback);
};


/**
 * @param {!proto.productpb.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.productpb.ProductAPIPromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/productpb.ProductAPI/Update',
      request,
      metadata || {},
      methodDescriptor_ProductAPI_Update);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.productpb.DeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_ProductAPI_Delete = new grpc.web.MethodDescriptor(
  '/productpb.ProductAPI/Delete',
  grpc.web.MethodType.UNARY,
  proto.productpb.DeleteRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.productpb.DeleteRequest} request
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
 *   !proto.productpb.DeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_ProductAPI_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.productpb.DeleteRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.productpb.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.productpb.ProductAPIClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/productpb.ProductAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_ProductAPI_Delete,
      callback);
};


/**
 * @param {!proto.productpb.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.productpb.ProductAPIPromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/productpb.ProductAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_ProductAPI_Delete);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.productpb.IsExistsRequest,
 *   !proto.productpb.IsExistsResponse>}
 */
const methodDescriptor_ProductAPI_IsExists = new grpc.web.MethodDescriptor(
  '/productpb.ProductAPI/IsExists',
  grpc.web.MethodType.UNARY,
  proto.productpb.IsExistsRequest,
  proto.productpb.IsExistsResponse,
  /**
   * @param {!proto.productpb.IsExistsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.productpb.IsExistsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.productpb.IsExistsRequest,
 *   !proto.productpb.IsExistsResponse>}
 */
const methodInfo_ProductAPI_IsExists = new grpc.web.AbstractClientBase.MethodInfo(
  proto.productpb.IsExistsResponse,
  /**
   * @param {!proto.productpb.IsExistsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.productpb.IsExistsResponse.deserializeBinary
);


/**
 * @param {!proto.productpb.IsExistsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.productpb.IsExistsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.productpb.IsExistsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.productpb.ProductAPIClient.prototype.isExists =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/productpb.ProductAPI/IsExists',
      request,
      metadata || {},
      methodDescriptor_ProductAPI_IsExists,
      callback);
};


/**
 * @param {!proto.productpb.IsExistsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.productpb.IsExistsResponse>}
 *     A native promise that resolves to the response
 */
proto.productpb.ProductAPIPromiseClient.prototype.isExists =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/productpb.ProductAPI/IsExists',
      request,
      metadata || {},
      methodDescriptor_ProductAPI_IsExists);
};


module.exports = proto.productpb;

