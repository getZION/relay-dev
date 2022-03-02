// package: proto.zion.v1
// file: proto/zion/v1/communities.proto

var proto_zion_v1_communities_pb = require("../../../proto/zion/v1/communities_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var CommunitiesService = (function () {
  function CommunitiesService() {}
  CommunitiesService.serviceName = "proto.zion.v1.CommunitiesService";
  return CommunitiesService;
}());

CommunitiesService.CreateCommunity = {
  methodName: "CreateCommunity",
  service: CommunitiesService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_communities_pb.CreateCommunityRequest,
  responseType: proto_zion_v1_communities_pb.CreateCommunityResponse
};

CommunitiesService.DeleteCommunity = {
  methodName: "DeleteCommunity",
  service: CommunitiesService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_communities_pb.DeleteCommunityRequest,
  responseType: proto_zion_v1_communities_pb.DeleteCommunityResponse
};

CommunitiesService.EditCommunity = {
  methodName: "EditCommunity",
  service: CommunitiesService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_communities_pb.EditCommunityRequest,
  responseType: proto_zion_v1_communities_pb.EditCommunityResponse
};

CommunitiesService.JoinCommunity = {
  methodName: "JoinCommunity",
  service: CommunitiesService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_communities_pb.JoinCommunityRequest,
  responseType: proto_zion_v1_communities_pb.JoinCommunityResponse
};

CommunitiesService.ListCommunity = {
  methodName: "ListCommunity",
  service: CommunitiesService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_communities_pb.ListCommunityRequest,
  responseType: proto_zion_v1_communities_pb.ListCommunityResponse
};

CommunitiesService.ReadCommunity = {
  methodName: "ReadCommunity",
  service: CommunitiesService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_communities_pb.ReadCommunityRequest,
  responseType: proto_zion_v1_communities_pb.ReadCommunityResponse
};

exports.CommunitiesService = CommunitiesService;

function CommunitiesServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

CommunitiesServiceClient.prototype.createCommunity = function createCommunity(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(CommunitiesService.CreateCommunity, {
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

CommunitiesServiceClient.prototype.deleteCommunity = function deleteCommunity(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(CommunitiesService.DeleteCommunity, {
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

CommunitiesServiceClient.prototype.editCommunity = function editCommunity(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(CommunitiesService.EditCommunity, {
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

CommunitiesServiceClient.prototype.joinCommunity = function joinCommunity(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(CommunitiesService.JoinCommunity, {
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

CommunitiesServiceClient.prototype.listCommunity = function listCommunity(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(CommunitiesService.ListCommunity, {
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

CommunitiesServiceClient.prototype.readCommunity = function readCommunity(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(CommunitiesService.ReadCommunity, {
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

exports.CommunitiesServiceClient = CommunitiesServiceClient;

