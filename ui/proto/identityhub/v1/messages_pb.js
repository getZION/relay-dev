// source: proto/identityhub/v1/messages.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = (function() {
  if (this) { return this; }
  if (typeof window !== 'undefined') { return window; }
  if (typeof global !== 'undefined') { return global; }
  if (typeof self !== 'undefined') { return self; }
  return Function('return this')();
}.call(null));

goog.exportSymbol('proto.proto.identityhub.v1.Attestation', null, global);
goog.exportSymbol('proto.proto.identityhub.v1.AttestationProtected', null, global);
goog.exportSymbol('proto.proto.identityhub.v1.Authorization', null, global);
goog.exportSymbol('proto.proto.identityhub.v1.Message', null, global);
goog.exportSymbol('proto.proto.identityhub.v1.MessageDescriptor', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.proto.identityhub.v1.Message = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.identityhub.v1.Message, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.identityhub.v1.Message.displayName = 'proto.proto.identityhub.v1.Message';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.proto.identityhub.v1.MessageDescriptor = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.identityhub.v1.MessageDescriptor, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.identityhub.v1.MessageDescriptor.displayName = 'proto.proto.identityhub.v1.MessageDescriptor';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.proto.identityhub.v1.Attestation = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.identityhub.v1.Attestation, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.identityhub.v1.Attestation.displayName = 'proto.proto.identityhub.v1.Attestation';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.proto.identityhub.v1.AttestationProtected = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.identityhub.v1.AttestationProtected, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.identityhub.v1.AttestationProtected.displayName = 'proto.proto.identityhub.v1.AttestationProtected';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.proto.identityhub.v1.Authorization = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.identityhub.v1.Authorization, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.identityhub.v1.Authorization.displayName = 'proto.proto.identityhub.v1.Authorization';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.proto.identityhub.v1.Message.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.identityhub.v1.Message.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.identityhub.v1.Message} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.identityhub.v1.Message.toObject = function(includeInstance, msg) {
  var f, obj = {
    data: jspb.Message.getFieldWithDefault(msg, 1, ""),
    descriptor: (f = msg.getDescriptor()) && proto.proto.identityhub.v1.MessageDescriptor.toObject(includeInstance, f),
    attestation: (f = msg.getAttestation()) && proto.proto.identityhub.v1.Attestation.toObject(includeInstance, f),
    authorization: (f = msg.getAuthorization()) && proto.proto.identityhub.v1.Authorization.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.proto.identityhub.v1.Message}
 */
proto.proto.identityhub.v1.Message.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.identityhub.v1.Message;
  return proto.proto.identityhub.v1.Message.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.identityhub.v1.Message} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.identityhub.v1.Message}
 */
proto.proto.identityhub.v1.Message.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setData(value);
      break;
    case 2:
      var value = new proto.proto.identityhub.v1.MessageDescriptor;
      reader.readMessage(value,proto.proto.identityhub.v1.MessageDescriptor.deserializeBinaryFromReader);
      msg.setDescriptor(value);
      break;
    case 3:
      var value = new proto.proto.identityhub.v1.Attestation;
      reader.readMessage(value,proto.proto.identityhub.v1.Attestation.deserializeBinaryFromReader);
      msg.setAttestation(value);
      break;
    case 4:
      var value = new proto.proto.identityhub.v1.Authorization;
      reader.readMessage(value,proto.proto.identityhub.v1.Authorization.deserializeBinaryFromReader);
      msg.setAuthorization(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.proto.identityhub.v1.Message.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.identityhub.v1.Message.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.identityhub.v1.Message} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.identityhub.v1.Message.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getData();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getDescriptor();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.proto.identityhub.v1.MessageDescriptor.serializeBinaryToWriter
    );
  }
  f = message.getAttestation();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.proto.identityhub.v1.Attestation.serializeBinaryToWriter
    );
  }
  f = message.getAuthorization();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.proto.identityhub.v1.Authorization.serializeBinaryToWriter
    );
  }
};


/**
 * optional string data = 1;
 * @return {string}
 */
proto.proto.identityhub.v1.Message.prototype.getData = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.identityhub.v1.Message} returns this
 */
proto.proto.identityhub.v1.Message.prototype.setData = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional MessageDescriptor descriptor = 2;
 * @return {?proto.proto.identityhub.v1.MessageDescriptor}
 */
proto.proto.identityhub.v1.Message.prototype.getDescriptor = function() {
  return /** @type{?proto.proto.identityhub.v1.MessageDescriptor} */ (
    jspb.Message.getWrapperField(this, proto.proto.identityhub.v1.MessageDescriptor, 2));
};


/**
 * @param {?proto.proto.identityhub.v1.MessageDescriptor|undefined} value
 * @return {!proto.proto.identityhub.v1.Message} returns this
*/
proto.proto.identityhub.v1.Message.prototype.setDescriptor = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.proto.identityhub.v1.Message} returns this
 */
proto.proto.identityhub.v1.Message.prototype.clearDescriptor = function() {
  return this.setDescriptor(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.proto.identityhub.v1.Message.prototype.hasDescriptor = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional Attestation attestation = 3;
 * @return {?proto.proto.identityhub.v1.Attestation}
 */
proto.proto.identityhub.v1.Message.prototype.getAttestation = function() {
  return /** @type{?proto.proto.identityhub.v1.Attestation} */ (
    jspb.Message.getWrapperField(this, proto.proto.identityhub.v1.Attestation, 3));
};


/**
 * @param {?proto.proto.identityhub.v1.Attestation|undefined} value
 * @return {!proto.proto.identityhub.v1.Message} returns this
*/
proto.proto.identityhub.v1.Message.prototype.setAttestation = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.proto.identityhub.v1.Message} returns this
 */
proto.proto.identityhub.v1.Message.prototype.clearAttestation = function() {
  return this.setAttestation(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.proto.identityhub.v1.Message.prototype.hasAttestation = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional Authorization authorization = 4;
 * @return {?proto.proto.identityhub.v1.Authorization}
 */
proto.proto.identityhub.v1.Message.prototype.getAuthorization = function() {
  return /** @type{?proto.proto.identityhub.v1.Authorization} */ (
    jspb.Message.getWrapperField(this, proto.proto.identityhub.v1.Authorization, 4));
};


/**
 * @param {?proto.proto.identityhub.v1.Authorization|undefined} value
 * @return {!proto.proto.identityhub.v1.Message} returns this
*/
proto.proto.identityhub.v1.Message.prototype.setAuthorization = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.proto.identityhub.v1.Message} returns this
 */
proto.proto.identityhub.v1.Message.prototype.clearAuthorization = function() {
  return this.setAuthorization(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.proto.identityhub.v1.Message.prototype.hasAuthorization = function() {
  return jspb.Message.getField(this, 4) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.identityhub.v1.MessageDescriptor.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.identityhub.v1.MessageDescriptor} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.identityhub.v1.MessageDescriptor.toObject = function(includeInstance, msg) {
  var f, obj = {
    method: jspb.Message.getFieldWithDefault(msg, 1, ""),
    objectid: jspb.Message.getFieldWithDefault(msg, 2, ""),
    schema: jspb.Message.getFieldWithDefault(msg, 3, ""),
    dataformat: jspb.Message.getFieldWithDefault(msg, 4, ""),
    datecreated: jspb.Message.getFieldWithDefault(msg, 5, ""),
    datepublished: jspb.Message.getFieldWithDefault(msg, 6, ""),
    datesort: jspb.Message.getFieldWithDefault(msg, 7, ""),
    root: jspb.Message.getFieldWithDefault(msg, 8, ""),
    parent: jspb.Message.getFieldWithDefault(msg, 9, ""),
    cid: jspb.Message.getFieldWithDefault(msg, 10, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.proto.identityhub.v1.MessageDescriptor}
 */
proto.proto.identityhub.v1.MessageDescriptor.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.identityhub.v1.MessageDescriptor;
  return proto.proto.identityhub.v1.MessageDescriptor.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.identityhub.v1.MessageDescriptor} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.identityhub.v1.MessageDescriptor}
 */
proto.proto.identityhub.v1.MessageDescriptor.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setMethod(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setObjectid(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setSchema(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setDataformat(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setDatecreated(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setDatepublished(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setDatesort(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setRoot(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setParent(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.setCid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.identityhub.v1.MessageDescriptor.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.identityhub.v1.MessageDescriptor} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.identityhub.v1.MessageDescriptor.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMethod();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getObjectid();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getSchema();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getDataformat();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getDatecreated();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getDatepublished();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getDatesort();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getRoot();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getParent();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
  f = message.getCid();
  if (f.length > 0) {
    writer.writeString(
      10,
      f
    );
  }
};


/**
 * optional string method = 1;
 * @return {string}
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.getMethod = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.identityhub.v1.MessageDescriptor} returns this
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.setMethod = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string objectId = 2;
 * @return {string}
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.getObjectid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.identityhub.v1.MessageDescriptor} returns this
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.setObjectid = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string schema = 3;
 * @return {string}
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.getSchema = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.identityhub.v1.MessageDescriptor} returns this
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.setSchema = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string dataFormat = 4;
 * @return {string}
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.getDataformat = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.identityhub.v1.MessageDescriptor} returns this
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.setDataformat = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string dateCreated = 5;
 * @return {string}
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.getDatecreated = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.identityhub.v1.MessageDescriptor} returns this
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.setDatecreated = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string datePublished = 6;
 * @return {string}
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.getDatepublished = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.identityhub.v1.MessageDescriptor} returns this
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.setDatepublished = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string dateSort = 7;
 * @return {string}
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.getDatesort = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.identityhub.v1.MessageDescriptor} returns this
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.setDatesort = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional string root = 8;
 * @return {string}
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.getRoot = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.identityhub.v1.MessageDescriptor} returns this
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.setRoot = function(value) {
  return jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional string parent = 9;
 * @return {string}
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.getParent = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.identityhub.v1.MessageDescriptor} returns this
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.setParent = function(value) {
  return jspb.Message.setProto3StringField(this, 9, value);
};


/**
 * optional string cid = 10;
 * @return {string}
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.getCid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 10, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.identityhub.v1.MessageDescriptor} returns this
 */
proto.proto.identityhub.v1.MessageDescriptor.prototype.setCid = function(value) {
  return jspb.Message.setProto3StringField(this, 10, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.proto.identityhub.v1.Attestation.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.identityhub.v1.Attestation.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.identityhub.v1.Attestation} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.identityhub.v1.Attestation.toObject = function(includeInstance, msg) {
  var f, obj = {
    pb_protected: (f = msg.getProtected()) && proto.proto.identityhub.v1.AttestationProtected.toObject(includeInstance, f),
    payload: jspb.Message.getFieldWithDefault(msg, 2, ""),
    signature: jspb.Message.getFieldWithDefault(msg, 3, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.proto.identityhub.v1.Attestation}
 */
proto.proto.identityhub.v1.Attestation.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.identityhub.v1.Attestation;
  return proto.proto.identityhub.v1.Attestation.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.identityhub.v1.Attestation} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.identityhub.v1.Attestation}
 */
proto.proto.identityhub.v1.Attestation.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.proto.identityhub.v1.AttestationProtected;
      reader.readMessage(value,proto.proto.identityhub.v1.AttestationProtected.deserializeBinaryFromReader);
      msg.setProtected(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPayload(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setSignature(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.proto.identityhub.v1.Attestation.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.identityhub.v1.Attestation.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.identityhub.v1.Attestation} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.identityhub.v1.Attestation.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProtected();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.proto.identityhub.v1.AttestationProtected.serializeBinaryToWriter
    );
  }
  f = message.getPayload();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getSignature();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional AttestationProtected protected = 1;
 * @return {?proto.proto.identityhub.v1.AttestationProtected}
 */
proto.proto.identityhub.v1.Attestation.prototype.getProtected = function() {
  return /** @type{?proto.proto.identityhub.v1.AttestationProtected} */ (
    jspb.Message.getWrapperField(this, proto.proto.identityhub.v1.AttestationProtected, 1));
};


/**
 * @param {?proto.proto.identityhub.v1.AttestationProtected|undefined} value
 * @return {!proto.proto.identityhub.v1.Attestation} returns this
*/
proto.proto.identityhub.v1.Attestation.prototype.setProtected = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.proto.identityhub.v1.Attestation} returns this
 */
proto.proto.identityhub.v1.Attestation.prototype.clearProtected = function() {
  return this.setProtected(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.proto.identityhub.v1.Attestation.prototype.hasProtected = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string payload = 2;
 * @return {string}
 */
proto.proto.identityhub.v1.Attestation.prototype.getPayload = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.identityhub.v1.Attestation} returns this
 */
proto.proto.identityhub.v1.Attestation.prototype.setPayload = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string signature = 3;
 * @return {string}
 */
proto.proto.identityhub.v1.Attestation.prototype.getSignature = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.identityhub.v1.Attestation} returns this
 */
proto.proto.identityhub.v1.Attestation.prototype.setSignature = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.proto.identityhub.v1.AttestationProtected.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.identityhub.v1.AttestationProtected.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.identityhub.v1.AttestationProtected} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.identityhub.v1.AttestationProtected.toObject = function(includeInstance, msg) {
  var f, obj = {
    alg: jspb.Message.getFieldWithDefault(msg, 1, ""),
    kid: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.proto.identityhub.v1.AttestationProtected}
 */
proto.proto.identityhub.v1.AttestationProtected.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.identityhub.v1.AttestationProtected;
  return proto.proto.identityhub.v1.AttestationProtected.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.identityhub.v1.AttestationProtected} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.identityhub.v1.AttestationProtected}
 */
proto.proto.identityhub.v1.AttestationProtected.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setAlg(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setKid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.proto.identityhub.v1.AttestationProtected.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.identityhub.v1.AttestationProtected.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.identityhub.v1.AttestationProtected} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.identityhub.v1.AttestationProtected.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAlg();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getKid();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string alg = 1;
 * @return {string}
 */
proto.proto.identityhub.v1.AttestationProtected.prototype.getAlg = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.identityhub.v1.AttestationProtected} returns this
 */
proto.proto.identityhub.v1.AttestationProtected.prototype.setAlg = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string kid = 2;
 * @return {string}
 */
proto.proto.identityhub.v1.AttestationProtected.prototype.getKid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.identityhub.v1.AttestationProtected} returns this
 */
proto.proto.identityhub.v1.AttestationProtected.prototype.setKid = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.proto.identityhub.v1.Authorization.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.identityhub.v1.Authorization.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.identityhub.v1.Authorization} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.identityhub.v1.Authorization.toObject = function(includeInstance, msg) {
  var f, obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.proto.identityhub.v1.Authorization}
 */
proto.proto.identityhub.v1.Authorization.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.identityhub.v1.Authorization;
  return proto.proto.identityhub.v1.Authorization.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.identityhub.v1.Authorization} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.identityhub.v1.Authorization}
 */
proto.proto.identityhub.v1.Authorization.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.proto.identityhub.v1.Authorization.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.identityhub.v1.Authorization.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.identityhub.v1.Authorization} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.identityhub.v1.Authorization.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};


goog.object.extend(exports, proto.proto.identityhub.v1);
