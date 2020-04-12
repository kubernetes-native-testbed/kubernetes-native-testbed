/**
 * @fileoverview gRPC-Web generated client stub for commentpb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.commentpb = require('./comment_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.commentpb.CommentAPIClient =
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
proto.commentpb.CommentAPIPromiseClient =
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
 *   !proto.commentpb.GetRequest,
 *   !proto.commentpb.GetResponse>}
 */
const methodDescriptor_CommentAPI_Get = new grpc.web.MethodDescriptor(
  '/commentpb.CommentAPI/Get',
  grpc.web.MethodType.UNARY,
  proto.commentpb.GetRequest,
  proto.commentpb.GetResponse,
  /**
   * @param {!proto.commentpb.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.commentpb.GetResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.commentpb.GetRequest,
 *   !proto.commentpb.GetResponse>}
 */
const methodInfo_CommentAPI_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.commentpb.GetResponse,
  /**
   * @param {!proto.commentpb.GetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.commentpb.GetResponse.deserializeBinary
);


/**
 * @param {!proto.commentpb.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.commentpb.GetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.commentpb.GetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.commentpb.CommentAPIClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/commentpb.CommentAPI/Get',
      request,
      metadata || {},
      methodDescriptor_CommentAPI_Get,
      callback);
};


/**
 * @param {!proto.commentpb.GetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.commentpb.GetResponse>}
 *     A native promise that resolves to the response
 */
proto.commentpb.CommentAPIPromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/commentpb.CommentAPI/Get',
      request,
      metadata || {},
      methodDescriptor_CommentAPI_Get);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.commentpb.SetRequest,
 *   !proto.commentpb.SetResponse>}
 */
const methodDescriptor_CommentAPI_Set = new grpc.web.MethodDescriptor(
  '/commentpb.CommentAPI/Set',
  grpc.web.MethodType.UNARY,
  proto.commentpb.SetRequest,
  proto.commentpb.SetResponse,
  /**
   * @param {!proto.commentpb.SetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.commentpb.SetResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.commentpb.SetRequest,
 *   !proto.commentpb.SetResponse>}
 */
const methodInfo_CommentAPI_Set = new grpc.web.AbstractClientBase.MethodInfo(
  proto.commentpb.SetResponse,
  /**
   * @param {!proto.commentpb.SetRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.commentpb.SetResponse.deserializeBinary
);


/**
 * @param {!proto.commentpb.SetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.commentpb.SetResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.commentpb.SetResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.commentpb.CommentAPIClient.prototype.set =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/commentpb.CommentAPI/Set',
      request,
      metadata || {},
      methodDescriptor_CommentAPI_Set,
      callback);
};


/**
 * @param {!proto.commentpb.SetRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.commentpb.SetResponse>}
 *     A native promise that resolves to the response
 */
proto.commentpb.CommentAPIPromiseClient.prototype.set =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/commentpb.CommentAPI/Set',
      request,
      metadata || {},
      methodDescriptor_CommentAPI_Set);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.commentpb.UpdateRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_CommentAPI_Update = new grpc.web.MethodDescriptor(
  '/commentpb.CommentAPI/Update',
  grpc.web.MethodType.UNARY,
  proto.commentpb.UpdateRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.commentpb.UpdateRequest} request
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
 *   !proto.commentpb.UpdateRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_CommentAPI_Update = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.commentpb.UpdateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.commentpb.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.commentpb.CommentAPIClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/commentpb.CommentAPI/Update',
      request,
      metadata || {},
      methodDescriptor_CommentAPI_Update,
      callback);
};


/**
 * @param {!proto.commentpb.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.commentpb.CommentAPIPromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/commentpb.CommentAPI/Update',
      request,
      metadata || {},
      methodDescriptor_CommentAPI_Update);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.commentpb.DeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_CommentAPI_Delete = new grpc.web.MethodDescriptor(
  '/commentpb.CommentAPI/Delete',
  grpc.web.MethodType.UNARY,
  proto.commentpb.DeleteRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.commentpb.DeleteRequest} request
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
 *   !proto.commentpb.DeleteRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_CommentAPI_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.commentpb.DeleteRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.commentpb.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.commentpb.CommentAPIClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/commentpb.CommentAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_CommentAPI_Delete,
      callback);
};


/**
 * @param {!proto.commentpb.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     A native promise that resolves to the response
 */
proto.commentpb.CommentAPIPromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/commentpb.CommentAPI/Delete',
      request,
      metadata || {},
      methodDescriptor_CommentAPI_Delete);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.commentpb.IsExistsRequest,
 *   !proto.commentpb.IsExistsResponse>}
 */
const methodDescriptor_CommentAPI_IsExists = new grpc.web.MethodDescriptor(
  '/commentpb.CommentAPI/IsExists',
  grpc.web.MethodType.UNARY,
  proto.commentpb.IsExistsRequest,
  proto.commentpb.IsExistsResponse,
  /**
   * @param {!proto.commentpb.IsExistsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.commentpb.IsExistsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.commentpb.IsExistsRequest,
 *   !proto.commentpb.IsExistsResponse>}
 */
const methodInfo_CommentAPI_IsExists = new grpc.web.AbstractClientBase.MethodInfo(
  proto.commentpb.IsExistsResponse,
  /**
   * @param {!proto.commentpb.IsExistsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.commentpb.IsExistsResponse.deserializeBinary
);


/**
 * @param {!proto.commentpb.IsExistsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.commentpb.IsExistsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.commentpb.IsExistsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.commentpb.CommentAPIClient.prototype.isExists =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/commentpb.CommentAPI/IsExists',
      request,
      metadata || {},
      methodDescriptor_CommentAPI_IsExists,
      callback);
};


/**
 * @param {!proto.commentpb.IsExistsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.commentpb.IsExistsResponse>}
 *     A native promise that resolves to the response
 */
proto.commentpb.CommentAPIPromiseClient.prototype.isExists =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/commentpb.CommentAPI/IsExists',
      request,
      metadata || {},
      methodDescriptor_CommentAPI_IsExists);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.commentpb.ChildCommentsRequest,
 *   !proto.commentpb.ChildCommentsResponse>}
 */
const methodDescriptor_CommentAPI_ChildComments = new grpc.web.MethodDescriptor(
  '/commentpb.CommentAPI/ChildComments',
  grpc.web.MethodType.UNARY,
  proto.commentpb.ChildCommentsRequest,
  proto.commentpb.ChildCommentsResponse,
  /**
   * @param {!proto.commentpb.ChildCommentsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.commentpb.ChildCommentsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.commentpb.ChildCommentsRequest,
 *   !proto.commentpb.ChildCommentsResponse>}
 */
const methodInfo_CommentAPI_ChildComments = new grpc.web.AbstractClientBase.MethodInfo(
  proto.commentpb.ChildCommentsResponse,
  /**
   * @param {!proto.commentpb.ChildCommentsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.commentpb.ChildCommentsResponse.deserializeBinary
);


/**
 * @param {!proto.commentpb.ChildCommentsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.commentpb.ChildCommentsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.commentpb.ChildCommentsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.commentpb.CommentAPIClient.prototype.childComments =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/commentpb.CommentAPI/ChildComments',
      request,
      metadata || {},
      methodDescriptor_CommentAPI_ChildComments,
      callback);
};


/**
 * @param {!proto.commentpb.ChildCommentsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.commentpb.ChildCommentsResponse>}
 *     A native promise that resolves to the response
 */
proto.commentpb.CommentAPIPromiseClient.prototype.childComments =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/commentpb.CommentAPI/ChildComments',
      request,
      metadata || {},
      methodDescriptor_CommentAPI_ChildComments);
};


module.exports = proto.commentpb;

