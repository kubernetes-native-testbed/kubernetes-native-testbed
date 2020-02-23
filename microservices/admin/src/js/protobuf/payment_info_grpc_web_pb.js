/**
 * @fileoverview gRPC-Web generated client stub for paymentinfopb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.paymentinfopb = require('./payment_info_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.paymentinfopb.PaymentInfoAPIClient =
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
proto.paymentinfopb.PaymentInfoAPIPromiseClient =
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
 *   !proto.paymentinfopb.GetRequest,
 *   !proto.paymentinfopb.GetResponse>}
 */
const methodDescriptor_PaymentInfoAPI_Get = new grpc.web.MethodDescriptor(
  '/paymentinfopb.PaymentInfoAPI/Get',
  grpc.web.MethodType.UNARY,
  proto.paymentinfopb.GetRequest,
  proto.paymentinfopb.GetResponse,
  /**
   * @param {!proto.paymentinfopb.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.paymentinfopb.GetResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.paymentinfopb.GetRequest,
 *   !proto.paymentinfopb.GetResponse>}
 */
const methodInfo_PaymentInfoAPI_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.paymentinfopb.GetResponse,
  /**
   * @param {!proto.paymentinfopb.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.paymentinfopb.GetResponse.deserializeBinary
);


/**
 * @param {!proto.paymentinfopb.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.paymentinfopb.GetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.paymentinfopb.GetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.paymentinfopb.PaymentInfoAPIClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/paymentinfopb.PaymentInfoAPI/Get',
      request,
      metadata || {},
      methodDescriptor_PaymentInfoAPI_Get,
      callback);
};


/**
 * @param {!proto.paymentinfopb.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.paymentinfopb.GetResponse>}
 *     A native promise that resolves to the response
 */
proto.paymentinfopb.PaymentInfoAPIPromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/paymentinfopb.PaymentInfoAPI/Get',
      request,
      metadata || {},
      methodDescriptor_PaymentInfoAPI_Get);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.paymentinfopb.SetRequest,
 *   !proto.paymentinfopb.SetResponse>}
 */
const methodDescriptor_PaymentInfoAPI_Set = new grpc.web.MethodDescriptor(
  '/paymentinfopb.PaymentInfoAPI/Set',
  grpc.web.MethodType.UNARY,
  proto.paymentinfopb.SetRequest,
  proto.paymentinfopb.SetResponse,
  /**
   * @param {!proto.paymentinfopb.SetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.paymentinfopb.SetResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.paymentinfopb.SetRequest,
 *   !proto.paymentinfopb.SetResponse>}
 */
const methodInfo_PaymentInfoAPI_Set = new grpc.web.AbstractClientBase.MethodInfo(
  proto.paymentinfopb.SetResponse,
  /**
   * @param {!proto.paymentinfopb.SetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.paymentinfopb.SetResponse.deserializeBinary
);


/**
 * @param {!proto.paymentinfopb.SetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.paymentinfopb.SetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.paymentinfopb.SetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.paymentinfopb.PaymentInfoAPIClient.prototype.set =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/paymentinfopb.PaymentInfoAPI/Set',
      request,
      metadata || {},
      methodDescriptor_PaymentInfoAPI_Set,
      callback);
};


/**
 * @param {!proto.paymentinfopb.SetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.paymentinfopb.SetResponse>}
 *     A native promise that resolves to the response
 */
proto.paymentinfopb.PaymentInfoAPIPromiseClient.prototype.set =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/paymentinfopb.PaymentInfoAPI/Set',
      request,
      metadata || {},
      methodDescriptor_PaymentInfoAPI_Set);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.paymentinfopb.UpdateRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_PaymentInfoAPI_Update = new grpc.web.MethodDescriptor(
  '/paymentinfopb.PaymentInfoAPI/Update',
  grpc.web.MethodType.UNARY,
  proto.paymentinfopb.UpdateRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.paymentinfopb.UpdateRequest} request
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
 *   !proto.paymentinfopb.UpdateRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_PaymentInfoAPI_Update = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.paymentinfopb.UpdateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.paymentinfopb.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.paymentinfopb.PaymentInfoAPIClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/paymentinfopb.PaymentInfoAPI/Update',
      request,
      metadata || {},
      methodDescriptor_PaymentInfoAPI_Update,
      callback);
};


/**
 * @param {!proto.paymentinfopb.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.paymentinfopb.PaymentInfoAPIPromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/paymentinfopb.PaymentInfoAPI/Update',
      request,
      metadata || {},
      methodDescriptor_PaymentInfoAPI_Update);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.paymentinfopb.DeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_PaymentInfoAPI_Delete = new grpc.web.MethodDescriptor(
  '/paymentinfopb.PaymentInfoAPI/Delete',
  grpc.web.MethodType.UNARY,
  proto.paymentinfopb.DeleteRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.paymentinfopb.DeleteRequest} request
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
 *   !proto.paymentinfopb.DeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_PaymentInfoAPI_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.paymentinfopb.DeleteRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.paymentinfopb.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.paymentinfopb.PaymentInfoAPIClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/paymentinfopb.PaymentInfoAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_PaymentInfoAPI_Delete,
      callback);
};


/**
 * @param {!proto.paymentinfopb.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.paymentinfopb.PaymentInfoAPIPromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/paymentinfopb.PaymentInfoAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_PaymentInfoAPI_Delete);
};


module.exports = proto.paymentinfopb;

