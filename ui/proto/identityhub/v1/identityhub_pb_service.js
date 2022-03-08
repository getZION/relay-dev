// package: proto.identityhub.v1
// file: proto/identityhub/v1/identityhub.proto

var proto_identityhub_v1_identityhub_pb = require("../../../proto/identityhub/v1/identityhub_pb");
var proto_identityhub_v1_requests_pb = require("../../../proto/identityhub/v1/requests_pb");
var proto_identityhub_v1_responses_pb = require("../../../proto/identityhub/v1/responses_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var HubRequestService = (function () {
  function HubRequestService() {}
  HubRequestService.serviceName = "proto.identityhub.v1.HubRequestService";
  return HubRequestService;
}());

HubRequestService.Process = {
  methodName: "Process",
  service: HubRequestService,
  requestStream: false,
  responseStream: false,
  requestType: proto_identityhub_v1_requests_pb.Request,
  responseType: proto_identityhub_v1_responses_pb.Response
};

exports.HubRequestService = HubRequestService;

function HubRequestServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

HubRequestServiceClient.prototype.process = function process(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(HubRequestService.Process, {
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

exports.HubRequestServiceClient = HubRequestServiceClient;

