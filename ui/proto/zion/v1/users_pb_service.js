// package: proto.zion.v1
// file: proto/zion/v1/users.proto

var proto_zion_v1_users_pb = require("../../../proto/zion/v1/users_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var UsersService = (function () {
  function UsersService() {}
  UsersService.serviceName = "proto.zion.v1.UsersService";
  return UsersService;
}());

UsersService.CreateUser = {
  methodName: "CreateUser",
  service: UsersService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_users_pb.CreateUserRequest,
  responseType: proto_zion_v1_users_pb.CreateUserResponse
};

UsersService.EditUser = {
  methodName: "EditUser",
  service: UsersService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_users_pb.EditUserRequest,
  responseType: proto_zion_v1_users_pb.EditUserResponse
};

UsersService.FindById = {
  methodName: "FindById",
  service: UsersService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_users_pb.FindByIdRequest,
  responseType: proto_zion_v1_users_pb.FindByIdResponse
};

exports.UsersService = UsersService;

function UsersServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

UsersServiceClient.prototype.createUser = function createUser(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(UsersService.CreateUser, {
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

UsersServiceClient.prototype.editUser = function editUser(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(UsersService.EditUser, {
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

UsersServiceClient.prototype.findById = function findById(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(UsersService.FindById, {
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

exports.UsersServiceClient = UsersServiceClient;

