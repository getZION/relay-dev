// package: proto.identityhub.v1
// file: proto/identityhub/v1/permissions.proto

var proto_identityhub_v1_permissions_pb = require("../../../proto/identityhub/v1/permissions_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var PermissionsService = (function () {
  function PermissionsService() {}
  PermissionsService.serviceName = "proto.identityhub.v1.PermissionsService";
  return PermissionsService;
}());

PermissionsService.PermissionsRequest = {
  methodName: "PermissionsRequest",
  service: PermissionsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_identityhub_v1_permissions_pb.PermissionsRequestRequest,
  responseType: proto_identityhub_v1_permissions_pb.PermissionsRequestResponse
};

PermissionsService.PermissionsQuery = {
  methodName: "PermissionsQuery",
  service: PermissionsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_identityhub_v1_permissions_pb.PermissionsQueryRequest,
  responseType: proto_identityhub_v1_permissions_pb.PermissionsQueryResponse
};

PermissionsService.PermissionsGrant = {
  methodName: "PermissionsGrant",
  service: PermissionsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_identityhub_v1_permissions_pb.PermissionsGrantRequest,
  responseType: proto_identityhub_v1_permissions_pb.PermissionsGrantResponse
};

PermissionsService.PermissionsRevoke = {
  methodName: "PermissionsRevoke",
  service: PermissionsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_identityhub_v1_permissions_pb.PermissionsRevokeRequest,
  responseType: proto_identityhub_v1_permissions_pb.PermissionsRevokeResponse
};

exports.PermissionsService = PermissionsService;

function PermissionsServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

PermissionsServiceClient.prototype.permissionsRequest = function permissionsRequest(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PermissionsService.PermissionsRequest, {
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

PermissionsServiceClient.prototype.permissionsQuery = function permissionsQuery(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PermissionsService.PermissionsQuery, {
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

PermissionsServiceClient.prototype.permissionsGrant = function permissionsGrant(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PermissionsService.PermissionsGrant, {
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

PermissionsServiceClient.prototype.permissionsRevoke = function permissionsRevoke(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PermissionsService.PermissionsRevoke, {
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

exports.PermissionsServiceClient = PermissionsServiceClient;

