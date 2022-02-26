// package: proto.zion.v1
// file: proto/zion/v1/nodeinfo.proto

var proto_zion_v1_nodeinfo_pb = require("../../../proto/zion/v1/nodeinfo_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var NodeInfoService = (function () {
  function NodeInfoService() {}
  NodeInfoService.serviceName = "proto.zion.v1.NodeInfoService";
  return NodeInfoService;
}());

NodeInfoService.GetNodeInfo = {
  methodName: "GetNodeInfo",
  service: NodeInfoService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_nodeinfo_pb.GetNodeInfoRequest,
  responseType: proto_zion_v1_nodeinfo_pb.GetNodeInfoResponse
};

exports.NodeInfoService = NodeInfoService;

function NodeInfoServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

NodeInfoServiceClient.prototype.getNodeInfo = function getNodeInfo(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(NodeInfoService.GetNodeInfo, {
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

exports.NodeInfoServiceClient = NodeInfoServiceClient;

