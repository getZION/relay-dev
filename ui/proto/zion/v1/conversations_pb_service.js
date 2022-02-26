// package: proto.zion.v1
// file: proto/zion/v1/conversations.proto

var proto_zion_v1_conversations_pb = require("../../../proto/zion/v1/conversations_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var ConversationService = (function () {
  function ConversationService() {}
  ConversationService.serviceName = "proto.zion.v1.ConversationService";
  return ConversationService;
}());

ConversationService.CreateConversation = {
  methodName: "CreateConversation",
  service: ConversationService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_conversations_pb.CreateConversationRequest,
  responseType: proto_zion_v1_conversations_pb.CreateConversationResponse
};

ConversationService.DeleteComment = {
  methodName: "DeleteComment",
  service: ConversationService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_conversations_pb.DeleteCommentRequest,
  responseType: proto_zion_v1_conversations_pb.DeleteCommentResponse
};

ConversationService.DeleteConversation = {
  methodName: "DeleteConversation",
  service: ConversationService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_conversations_pb.DeleteConversationRequest,
  responseType: proto_zion_v1_conversations_pb.DeleteConversationResponse
};

ConversationService.GetActivity = {
  methodName: "GetActivity",
  service: ConversationService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_conversations_pb.GetActivityRequest,
  responseType: proto_zion_v1_conversations_pb.GetActivityResponse
};

ConversationService.GetConversation = {
  methodName: "GetConversation",
  service: ConversationService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_conversations_pb.GetConversationRequest,
  responseType: proto_zion_v1_conversations_pb.GetConversationResponse
};

ConversationService.GetFeed = {
  methodName: "GetFeed",
  service: ConversationService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_conversations_pb.GetFeedRequest,
  responseType: proto_zion_v1_conversations_pb.GetFeedResponse
};

exports.ConversationService = ConversationService;

function ConversationServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ConversationServiceClient.prototype.createConversation = function createConversation(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ConversationService.CreateConversation, {
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

ConversationServiceClient.prototype.deleteComment = function deleteComment(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ConversationService.DeleteComment, {
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

ConversationServiceClient.prototype.deleteConversation = function deleteConversation(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ConversationService.DeleteConversation, {
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

ConversationServiceClient.prototype.getActivity = function getActivity(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ConversationService.GetActivity, {
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

ConversationServiceClient.prototype.getConversation = function getConversation(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ConversationService.GetConversation, {
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

ConversationServiceClient.prototype.getFeed = function getFeed(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ConversationService.GetFeed, {
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

exports.ConversationServiceClient = ConversationServiceClient;

