// package: proto.identityhub.v1
// file: proto/identityhub/v1/permissions.proto

import * as jspb from "google-protobuf";
import * as proto_identityhub_v1_requests_pb from "../../../proto/identityhub/v1/requests_pb";
import * as proto_identityhub_v1_responses_pb from "../../../proto/identityhub/v1/responses_pb";

export class PermissionsRequestRequest extends jspb.Message {
  hasRequest(): boolean;
  clearRequest(): void;
  getRequest(): proto_identityhub_v1_requests_pb.Request | undefined;
  setRequest(value?: proto_identityhub_v1_requests_pb.Request): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PermissionsRequestRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PermissionsRequestRequest): PermissionsRequestRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PermissionsRequestRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PermissionsRequestRequest;
  static deserializeBinaryFromReader(message: PermissionsRequestRequest, reader: jspb.BinaryReader): PermissionsRequestRequest;
}

export namespace PermissionsRequestRequest {
  export type AsObject = {
    request?: proto_identityhub_v1_requests_pb.Request.AsObject,
  }
}

export class PermissionsRequestResponse extends jspb.Message {
  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): proto_identityhub_v1_responses_pb.Response | undefined;
  setResponse(value?: proto_identityhub_v1_responses_pb.Response): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PermissionsRequestResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PermissionsRequestResponse): PermissionsRequestResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PermissionsRequestResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PermissionsRequestResponse;
  static deserializeBinaryFromReader(message: PermissionsRequestResponse, reader: jspb.BinaryReader): PermissionsRequestResponse;
}

export namespace PermissionsRequestResponse {
  export type AsObject = {
    response?: proto_identityhub_v1_responses_pb.Response.AsObject,
  }
}

export class PermissionsQueryRequest extends jspb.Message {
  hasRequest(): boolean;
  clearRequest(): void;
  getRequest(): proto_identityhub_v1_requests_pb.Request | undefined;
  setRequest(value?: proto_identityhub_v1_requests_pb.Request): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PermissionsQueryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PermissionsQueryRequest): PermissionsQueryRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PermissionsQueryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PermissionsQueryRequest;
  static deserializeBinaryFromReader(message: PermissionsQueryRequest, reader: jspb.BinaryReader): PermissionsQueryRequest;
}

export namespace PermissionsQueryRequest {
  export type AsObject = {
    request?: proto_identityhub_v1_requests_pb.Request.AsObject,
  }
}

export class PermissionsQueryResponse extends jspb.Message {
  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): proto_identityhub_v1_responses_pb.Response | undefined;
  setResponse(value?: proto_identityhub_v1_responses_pb.Response): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PermissionsQueryResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PermissionsQueryResponse): PermissionsQueryResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PermissionsQueryResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PermissionsQueryResponse;
  static deserializeBinaryFromReader(message: PermissionsQueryResponse, reader: jspb.BinaryReader): PermissionsQueryResponse;
}

export namespace PermissionsQueryResponse {
  export type AsObject = {
    response?: proto_identityhub_v1_responses_pb.Response.AsObject,
  }
}

export class PermissionsGrantRequest extends jspb.Message {
  hasRequest(): boolean;
  clearRequest(): void;
  getRequest(): proto_identityhub_v1_requests_pb.Request | undefined;
  setRequest(value?: proto_identityhub_v1_requests_pb.Request): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PermissionsGrantRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PermissionsGrantRequest): PermissionsGrantRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PermissionsGrantRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PermissionsGrantRequest;
  static deserializeBinaryFromReader(message: PermissionsGrantRequest, reader: jspb.BinaryReader): PermissionsGrantRequest;
}

export namespace PermissionsGrantRequest {
  export type AsObject = {
    request?: proto_identityhub_v1_requests_pb.Request.AsObject,
  }
}

export class PermissionsGrantResponse extends jspb.Message {
  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): proto_identityhub_v1_responses_pb.Response | undefined;
  setResponse(value?: proto_identityhub_v1_responses_pb.Response): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PermissionsGrantResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PermissionsGrantResponse): PermissionsGrantResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PermissionsGrantResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PermissionsGrantResponse;
  static deserializeBinaryFromReader(message: PermissionsGrantResponse, reader: jspb.BinaryReader): PermissionsGrantResponse;
}

export namespace PermissionsGrantResponse {
  export type AsObject = {
    response?: proto_identityhub_v1_responses_pb.Response.AsObject,
  }
}

export class PermissionsRevokeRequest extends jspb.Message {
  hasRequest(): boolean;
  clearRequest(): void;
  getRequest(): proto_identityhub_v1_requests_pb.Request | undefined;
  setRequest(value?: proto_identityhub_v1_requests_pb.Request): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PermissionsRevokeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PermissionsRevokeRequest): PermissionsRevokeRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PermissionsRevokeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PermissionsRevokeRequest;
  static deserializeBinaryFromReader(message: PermissionsRevokeRequest, reader: jspb.BinaryReader): PermissionsRevokeRequest;
}

export namespace PermissionsRevokeRequest {
  export type AsObject = {
    request?: proto_identityhub_v1_requests_pb.Request.AsObject,
  }
}

export class PermissionsRevokeResponse extends jspb.Message {
  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): proto_identityhub_v1_responses_pb.Response | undefined;
  setResponse(value?: proto_identityhub_v1_responses_pb.Response): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PermissionsRevokeResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PermissionsRevokeResponse): PermissionsRevokeResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PermissionsRevokeResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PermissionsRevokeResponse;
  static deserializeBinaryFromReader(message: PermissionsRevokeResponse, reader: jspb.BinaryReader): PermissionsRevokeResponse;
}

export namespace PermissionsRevokeResponse {
  export type AsObject = {
    response?: proto_identityhub_v1_responses_pb.Response.AsObject,
  }
}

