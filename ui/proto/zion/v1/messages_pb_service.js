// package: proto.zion.v1
// file: proto/zion/v1/messages.proto

var proto_zion_v1_messages_pb = require("../../../proto/zion/v1/messages_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var MessagesService = (function () {
  function MessagesService() {}
  MessagesService.serviceName = "proto.zion.v1.MessagesService";
  return MessagesService;
}());

MessagesService.BoostMessage = {
  methodName: "BoostMessage",
  service: MessagesService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_messages_pb.BoostMessageRequest,
  responseType: proto_zion_v1_messages_pb.BoostMessageResponse
};

MessagesService.DeleteMessage = {
  methodName: "DeleteMessage",
  service: MessagesService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_messages_pb.DeleteMessageRequest,
  responseType: proto_zion_v1_messages_pb.DeleteMessageResponse
};

MessagesService.GetMessages = {
  methodName: "GetMessages",
  service: MessagesService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_messages_pb.GetMessagesRequest,
  responseType: proto_zion_v1_messages_pb.GetMessagesResponse
};

MessagesService.SendMessage = {
  methodName: "SendMessage",
  service: MessagesService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_messages_pb.SendMessageRequest,
  responseType: proto_zion_v1_messages_pb.SendMessageResponse
};

exports.MessagesService = MessagesService;

function MessagesServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

MessagesServiceClient.prototype.boostMessage = function boostMessage(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MessagesService.BoostMessage, {
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

MessagesServiceClient.prototype.deleteMessage = function deleteMessage(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MessagesService.DeleteMessage, {
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

MessagesServiceClient.prototype.getMessages = function getMessages(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MessagesService.GetMessages, {
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

MessagesServiceClient.prototype.sendMessage = function sendMessage(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MessagesService.SendMessage, {
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

exports.MessagesServiceClient = MessagesServiceClient;

