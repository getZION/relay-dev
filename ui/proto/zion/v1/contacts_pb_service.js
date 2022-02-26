// package: proto.zion.v1
// file: proto/zion/v1/contacts.proto

var proto_zion_v1_contacts_pb = require("../../../proto/zion/v1/contacts_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var ContactsService = (function () {
  function ContactsService() {}
  ContactsService.serviceName = "proto.zion.v1.ContactsService";
  return ContactsService;
}());

ContactsService.AddByPubkey = {
  methodName: "AddByPubkey",
  service: ContactsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_contacts_pb.AddByPubkeyRequest,
  responseType: proto_zion_v1_contacts_pb.AddByPubkeyResponse
};

ContactsService.AddByUsername = {
  methodName: "AddByUsername",
  service: ContactsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_contacts_pb.AddByUsernameRequest,
  responseType: proto_zion_v1_contacts_pb.AddByUsernameResponse
};

ContactsService.DeleteContact = {
  methodName: "DeleteContact",
  service: ContactsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_contacts_pb.DeleteContactRequest,
  responseType: proto_zion_v1_contacts_pb.DeleteContactResponse
};

ContactsService.GetMyContacts = {
  methodName: "GetMyContacts",
  service: ContactsService,
  requestStream: false,
  responseStream: false,
  requestType: proto_zion_v1_contacts_pb.GetMyContactsRequest,
  responseType: proto_zion_v1_contacts_pb.GetMyContactsResponse
};

exports.ContactsService = ContactsService;

function ContactsServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ContactsServiceClient.prototype.addByPubkey = function addByPubkey(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ContactsService.AddByPubkey, {
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

ContactsServiceClient.prototype.addByUsername = function addByUsername(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ContactsService.AddByUsername, {
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

ContactsServiceClient.prototype.deleteContact = function deleteContact(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ContactsService.DeleteContact, {
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

ContactsServiceClient.prototype.getMyContacts = function getMyContacts(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ContactsService.GetMyContacts, {
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

exports.ContactsServiceClient = ContactsServiceClient;

