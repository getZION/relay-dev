// package: proto.identityhub.v1
// file: proto/identityhub/v1/collections.proto

import * as jspb from "google-protobuf";
import * as proto_identityhub_v1_requests_pb from "../../../proto/identityhub/v1/requests_pb";
import * as proto_identityhub_v1_responses_pb from "../../../proto/identityhub/v1/responses_pb";

export class CollectionsQueryRequest extends jspb.Message {
  hasRequest(): boolean;
  clearRequest(): void;
  getRequest(): proto_identityhub_v1_requests_pb.Request | undefined;
  setRequest(value?: proto_identityhub_v1_requests_pb.Request): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CollectionsQueryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CollectionsQueryRequest): CollectionsQueryRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CollectionsQueryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CollectionsQueryRequest;
  static deserializeBinaryFromReader(message: CollectionsQueryRequest, reader: jspb.BinaryReader): CollectionsQueryRequest;
}

export namespace CollectionsQueryRequest {
  export type AsObject = {
    request?: proto_identityhub_v1_requests_pb.Request.AsObject,
  }
}

export class CollectionsQueryResponse extends jspb.Message {
  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): proto_identityhub_v1_responses_pb.Response | undefined;
  setResponse(value?: proto_identityhub_v1_responses_pb.Response): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CollectionsQueryResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CollectionsQueryResponse): CollectionsQueryResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CollectionsQueryResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CollectionsQueryResponse;
  static deserializeBinaryFromReader(message: CollectionsQueryResponse, reader: jspb.BinaryReader): CollectionsQueryResponse;
}

export namespace CollectionsQueryResponse {
  export type AsObject = {
    response?: proto_identityhub_v1_responses_pb.Response.AsObject,
  }
}

export class CollectionsWriteRequest extends jspb.Message {
  hasRequest(): boolean;
  clearRequest(): void;
  getRequest(): proto_identityhub_v1_requests_pb.Request | undefined;
  setRequest(value?: proto_identityhub_v1_requests_pb.Request): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CollectionsWriteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CollectionsWriteRequest): CollectionsWriteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CollectionsWriteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CollectionsWriteRequest;
  static deserializeBinaryFromReader(message: CollectionsWriteRequest, reader: jspb.BinaryReader): CollectionsWriteRequest;
}

export namespace CollectionsWriteRequest {
  export type AsObject = {
    request?: proto_identityhub_v1_requests_pb.Request.AsObject,
  }
}

export class CollectionsWriteResponse extends jspb.Message {
  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): proto_identityhub_v1_responses_pb.Response | undefined;
  setResponse(value?: proto_identityhub_v1_responses_pb.Response): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CollectionsWriteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CollectionsWriteResponse): CollectionsWriteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CollectionsWriteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CollectionsWriteResponse;
  static deserializeBinaryFromReader(message: CollectionsWriteResponse, reader: jspb.BinaryReader): CollectionsWriteResponse;
}

export namespace CollectionsWriteResponse {
  export type AsObject = {
    response?: proto_identityhub_v1_responses_pb.Response.AsObject,
  }
}

export class CollectionsCommitRequest extends jspb.Message {
  hasRequest(): boolean;
  clearRequest(): void;
  getRequest(): proto_identityhub_v1_requests_pb.Request | undefined;
  setRequest(value?: proto_identityhub_v1_requests_pb.Request): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CollectionsCommitRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CollectionsCommitRequest): CollectionsCommitRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CollectionsCommitRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CollectionsCommitRequest;
  static deserializeBinaryFromReader(message: CollectionsCommitRequest, reader: jspb.BinaryReader): CollectionsCommitRequest;
}

export namespace CollectionsCommitRequest {
  export type AsObject = {
    request?: proto_identityhub_v1_requests_pb.Request.AsObject,
  }
}

export class CollectionsCommitResponse extends jspb.Message {
  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): proto_identityhub_v1_responses_pb.Response | undefined;
  setResponse(value?: proto_identityhub_v1_responses_pb.Response): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CollectionsCommitResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CollectionsCommitResponse): CollectionsCommitResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CollectionsCommitResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CollectionsCommitResponse;
  static deserializeBinaryFromReader(message: CollectionsCommitResponse, reader: jspb.BinaryReader): CollectionsCommitResponse;
}

export namespace CollectionsCommitResponse {
  export type AsObject = {
    response?: proto_identityhub_v1_responses_pb.Response.AsObject,
  }
}

export class CollectionsDeleteRequest extends jspb.Message {
  hasRequest(): boolean;
  clearRequest(): void;
  getRequest(): proto_identityhub_v1_requests_pb.Request | undefined;
  setRequest(value?: proto_identityhub_v1_requests_pb.Request): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CollectionsDeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CollectionsDeleteRequest): CollectionsDeleteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CollectionsDeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CollectionsDeleteRequest;
  static deserializeBinaryFromReader(message: CollectionsDeleteRequest, reader: jspb.BinaryReader): CollectionsDeleteRequest;
}

export namespace CollectionsDeleteRequest {
  export type AsObject = {
    request?: proto_identityhub_v1_requests_pb.Request.AsObject,
  }
}

export class CollectionsDeleteResponse extends jspb.Message {
  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): proto_identityhub_v1_responses_pb.Response | undefined;
  setResponse(value?: proto_identityhub_v1_responses_pb.Response): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CollectionsDeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CollectionsDeleteResponse): CollectionsDeleteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CollectionsDeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CollectionsDeleteResponse;
  static deserializeBinaryFromReader(message: CollectionsDeleteResponse, reader: jspb.BinaryReader): CollectionsDeleteResponse;
}

export namespace CollectionsDeleteResponse {
  export type AsObject = {
    response?: proto_identityhub_v1_responses_pb.Response.AsObject,
  }
}

