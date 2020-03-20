/**
 * @fileoverview gRPC-Web generated client stub for productpb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')
const proto = {};
proto.productpb = require('./image_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.productpb.ImageAPIClient =
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
proto.productpb.ImageAPIPromiseClient =
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
 *   !proto.productpb.ImageUploadRequest,
 *   !proto.productpb.ImageUploadResponse>}
 */
const methodDescriptor_ImageAPI_Upload = new grpc.web.MethodDescriptor(
  '/productpb.ImageAPI/Upload',
  grpc.web.MethodType.UNARY,
  proto.productpb.ImageUploadRequest,
  proto.productpb.ImageUploadResponse,
  /**
   * @param {!proto.productpb.ImageUploadRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.productpb.ImageUploadResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.productpb.ImageUploadRequest,
 *   !proto.productpb.ImageUploadResponse>}
 */
const methodInfo_ImageAPI_Upload = new grpc.web.AbstractClientBase.MethodInfo(
  proto.productpb.ImageUploadResponse,
  /**
   * @param {!proto.productpb.ImageUploadRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.productpb.ImageUploadResponse.deserializeBinary
);


/**
 * @param {!proto.productpb.ImageUploadRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.productpb.ImageUploadResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.productpb.ImageUploadResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.productpb.ImageAPIClient.prototype.upload =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/productpb.ImageAPI/Upload',
      request,
      metadata || {},
      methodDescriptor_ImageAPI_Upload,
      callback);
};


/**
 * @param {!proto.productpb.ImageUploadRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.productpb.ImageUploadResponse>}
 *     A native promise that resolves to the response
 */
proto.productpb.ImageAPIPromiseClient.prototype.upload =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/productpb.ImageAPI/Upload',
      request,
      metadata || {},
      methodDescriptor_ImageAPI_Upload);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.productpb.ImageDeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_ImageAPI_Delete = new grpc.web.MethodDescriptor(
  '/productpb.ImageAPI/Delete',
  grpc.web.MethodType.UNARY,
  proto.productpb.ImageDeleteRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.productpb.ImageDeleteRequest} request
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
 *   !proto.productpb.ImageDeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_ImageAPI_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.productpb.ImageDeleteRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.productpb.ImageDeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.productpb.ImageAPIClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/productpb.ImageAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_ImageAPI_Delete,
      callback);
};


/**
 * @param {!proto.productpb.ImageDeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.productpb.ImageAPIPromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/productpb.ImageAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_ImageAPI_Delete);
};


module.exports = proto.productpb;

