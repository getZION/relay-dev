// package: proto.identityhub.v1
// file: proto/identityhub/v1/collections.proto

var proto_identityhub_v1_collections_pb = require("../../../proto/identityhub/v1/collections_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var CollectionsService = (function () {
  function CollectionsService() {}
  CollectionsService.serviceName = "proto.identityhub.v1.CollectionsService";
  return CollectionsService;
}());

CollectionsService.CollectionsQuery = {
  methodName: "CollectionsQuery",
  service: CollectionsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_identityhub_v1_collections_pb.CollectionsQueryRequest,
  responseType: proto_identityhub_v1_collections_pb.CollectionsQueryResponse
};

CollectionsService.CollectionsWrite = {
  methodName: "CollectionsWrite",
  service: CollectionsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_identityhub_v1_collections_pb.CollectionsWriteRequest,
  responseType: proto_identityhub_v1_collections_pb.CollectionsWriteResponse
};

CollectionsService.CollectionsCommit = {
  methodName: "CollectionsCommit",
  service: CollectionsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_identityhub_v1_collections_pb.CollectionsCommitRequest,
  responseType: proto_identityhub_v1_collections_pb.CollectionsCommitResponse
};

CollectionsService.CollectionsDelete = {
  methodName: "CollectionsDelete",
  service: CollectionsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_identityhub_v1_collections_pb.CollectionsDeleteRequest,
  responseType: proto_identityhub_v1_collections_pb.CollectionsDeleteResponse
};

exports.CollectionsService = CollectionsService;

function CollectionsServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

CollectionsServiceClient.prototype.collectionsQuery = function collectionsQuery(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(CollectionsService.CollectionsQuery, {
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

CollectionsServiceClient.prototype.collectionsWrite = function collectionsWrite(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(CollectionsService.CollectionsWrite, {
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

CollectionsServiceClient.prototype.collectionsCommit = function collectionsCommit(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(CollectionsService.CollectionsCommit, {
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

CollectionsServiceClient.prototype.collectionsDelete = function collectionsDelete(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(CollectionsService.CollectionsDelete, {
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

exports.CollectionsServiceClient = CollectionsServiceClient;

