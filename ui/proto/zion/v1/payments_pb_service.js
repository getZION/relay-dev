// package: proto.zion.v1
// file: proto/zion/v1/payments.proto

var proto_zion_v1_payments_pb = require("../../../proto/zion/v1/payments_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var PaymentsService = (function () {
  function PaymentsService() {}
  PaymentsService.serviceName = "proto.zion.v1.PaymentsService";
  return PaymentsService;
}());

PaymentsService.CreateInvoice = {
  methodName: "CreateInvoice",
  service: PaymentsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_payments_pb.CreateInvoiceRequest,
  responseType: proto_zion_v1_payments_pb.CreateInvoiceResponse
};

PaymentsService.GetPaymentsHistory = {
  methodName: "GetPaymentsHistory",
  service: PaymentsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_payments_pb.GetPaymentsHistoryRequest,
  responseType: proto_zion_v1_payments_pb.GetPaymentsHistoryResponse
};

PaymentsService.PayInvoice = {
  methodName: "PayInvoice",
  service: PaymentsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_payments_pb.PayInvoiceRequest,
  responseType: proto_zion_v1_payments_pb.PayInvoiceResponse
};

PaymentsService.PayUser = {
  methodName: "PayUser",
  service: PaymentsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_payments_pb.PayUserRequest,
  responseType: proto_zion_v1_payments_pb.PayUserResponse
};

exports.PaymentsService = PaymentsService;

function PaymentsServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

PaymentsServiceClient.prototype.createInvoice = function createInvoice(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PaymentsService.CreateInvoice, {
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

PaymentsServiceClient.prototype.getPaymentsHistory = function getPaymentsHistory(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PaymentsService.GetPaymentsHistory, {
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

PaymentsServiceClient.prototype.payInvoice = function payInvoice(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PaymentsService.PayInvoice, {
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

PaymentsServiceClient.prototype.payUser = function payUser(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PaymentsService.PayUser, {
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

exports.PaymentsServiceClient = PaymentsServiceClient;

