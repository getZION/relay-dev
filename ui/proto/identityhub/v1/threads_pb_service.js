// package: proto.identityhub.v1
// file: proto/identityhub/v1/threads.proto

var proto_identityhub_v1_threads_pb = require("../../../proto/identityhub/v1/threads_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var ThreadsService = (function () {
  function ThreadsService() {}
  ThreadsService.serviceName = "proto.identityhub.v1.ThreadsService";
  return ThreadsService;
}());

ThreadsService.ThreadsQuery = {
  methodName: "ThreadsQuery",
  service: ThreadsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_identityhub_v1_threads_pb.ThreadsQueryRequest,
  responseType: proto_identityhub_v1_threads_pb.ThreadsQueryResponse
};

ThreadsService.ThreadsCreate = {
  methodName: "ThreadsCreate",
  service: ThreadsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_identityhub_v1_threads_pb.ThreadsCreateRequest,
  responseType: proto_identityhub_v1_threads_pb.ThreadsCreateResponse
};

ThreadsService.ThreadsReply = {
  methodName: "ThreadsReply",
  service: ThreadsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_identityhub_v1_threads_pb.ThreadsReplyRequest,
  responseType: proto_identityhub_v1_threads_pb.ThreadsReplyResponse
};

ThreadsService.ThreadsClose = {
  methodName: "ThreadsClose",
  service: ThreadsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_identityhub_v1_threads_pb.ThreadsCloseRequest,
  responseType: proto_identityhub_v1_threads_pb.ThreadsCloseResponse
};

ThreadsService.ThreadsDelete = {
  methodName: "ThreadsDelete",
  service: ThreadsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_identityhub_v1_threads_pb.ThreadsDeleteRequest,
  responseType: proto_identityhub_v1_threads_pb.ThreadsDeleteResponse
};

exports.ThreadsService = ThreadsService;

function ThreadsServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ThreadsServiceClient.prototype.threadsQuery = function threadsQuery(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ThreadsService.ThreadsQuery, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ThreadsServiceClient.prototype.threadsCreate = function threadsCreate(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ThreadsService.ThreadsCreate, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ThreadsServiceClient.prototype.threadsReply = function threadsReply(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ThreadsService.ThreadsReply, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ThreadsServiceClient.prototype.threadsClose = function threadsClose(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ThreadsService.ThreadsClose, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ThreadsServiceClient.prototype.threadsDelete = function threadsDelete(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ThreadsService.ThreadsDelete, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.ThreadsServiceClient = ThreadsServiceClient;

