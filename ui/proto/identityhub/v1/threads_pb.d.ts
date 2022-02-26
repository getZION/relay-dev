// package: proto.identityhub.v1
// file: proto/identityhub/v1/threads.proto

import * as jspb from "google-protobuf";
import * as proto_identityhub_v1_requests_pb from "../../../proto/identityhub/v1/requests_pb";
import * as proto_identityhub_v1_responses_pb from "../../../proto/identityhub/v1/responses_pb";

export class ThreadsQueryRequest extends jspb.Message {
  hasRequest(): boolean;
  clearRequest(): void;
  getRequest(): proto_identityhub_v1_requests_pb.Request | undefined;
  setRequest(value?: proto_identityhub_v1_requests_pb.Request): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ThreadsQueryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ThreadsQueryRequest): ThreadsQueryRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ThreadsQueryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ThreadsQueryRequest;
  static deserializeBinaryFromReader(message: ThreadsQueryRequest, reader: jspb.BinaryReader): ThreadsQueryRequest;
}

export namespace ThreadsQueryRequest {
  export type AsObject = {
    request?: proto_identityhub_v1_requests_pb.Request.AsObject,
  }
}

export class ThreadsQueryResponse extends jspb.Message {
  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): proto_identityhub_v1_responses_pb.Response | undefined;
  setResponse(value?: proto_identityhub_v1_responses_pb.Response): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ThreadsQueryResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ThreadsQueryResponse): ThreadsQueryResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ThreadsQueryResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ThreadsQueryResponse;
  static deserializeBinaryFromReader(message: ThreadsQueryResponse, reader: jspb.BinaryReader): ThreadsQueryResponse;
}

export namespace ThreadsQueryResponse {
  export type AsObject = {
    response?: proto_identityhub_v1_responses_pb.Response.AsObject,
  }
}

export class ThreadsCreateRequest extends jspb.Message {
  hasRequest(): boolean;
  clearRequest(): void;
  getRequest(): proto_identityhub_v1_requests_pb.Request | undefined;
  setRequest(value?: proto_identityhub_v1_requests_pb.Request): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ThreadsCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ThreadsCreateRequest): ThreadsCreateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ThreadsCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ThreadsCreateRequest;
  static deserializeBinaryFromReader(message: ThreadsCreateRequest, reader: jspb.BinaryReader): ThreadsCreateRequest;
}

export namespace ThreadsCreateRequest {
  export type AsObject = {
    request?: proto_identityhub_v1_requests_pb.Request.AsObject,
  }
}

export class ThreadsCreateResponse extends jspb.Message {
  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): proto_identityhub_v1_responses_pb.Response | undefined;
  setResponse(value?: proto_identityhub_v1_responses_pb.Response): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ThreadsCreateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ThreadsCreateResponse): ThreadsCreateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ThreadsCreateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ThreadsCreateResponse;
  static deserializeBinaryFromReader(message: ThreadsCreateResponse, reader: jspb.BinaryReader): ThreadsCreateResponse;
}

export namespace ThreadsCreateResponse {
  export type AsObject = {
    response?: proto_identityhub_v1_responses_pb.Response.AsObject,
  }
}

export class ThreadsReplyRequest extends jspb.Message {
  hasRequest(): boolean;
  clearRequest(): void;
  getRequest(): proto_identityhub_v1_requests_pb.Request | undefined;
  setRequest(value?: proto_identityhub_v1_requests_pb.Request): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ThreadsReplyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ThreadsReplyRequest): ThreadsReplyRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ThreadsReplyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ThreadsReplyRequest;
  static deserializeBinaryFromReader(message: ThreadsReplyRequest, reader: jspb.BinaryReader): ThreadsReplyRequest;
}

export namespace ThreadsReplyRequest {
  export type AsObject = {
    request?: proto_identityhub_v1_requests_pb.Request.AsObject,
  }
}

export class ThreadsReplyResponse extends jspb.Message {
  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): proto_identityhub_v1_responses_pb.Response | undefined;
  setResponse(value?: proto_identityhub_v1_responses_pb.Response): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ThreadsReplyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ThreadsReplyResponse): ThreadsReplyResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ThreadsReplyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ThreadsReplyResponse;
  static deserializeBinaryFromReader(message: ThreadsReplyResponse, reader: jspb.BinaryReader): ThreadsReplyResponse;
}

export namespace ThreadsReplyResponse {
  export type AsObject = {
    response?: proto_identityhub_v1_responses_pb.Response.AsObject,
  }
}

export class ThreadsCloseRequest extends jspb.Message {
  hasRequest(): boolean;
  clearRequest(): void;
  getRequest(): proto_identityhub_v1_requests_pb.Request | undefined;
  setRequest(value?: proto_identityhub_v1_requests_pb.Request): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ThreadsCloseRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ThreadsCloseRequest): ThreadsCloseRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ThreadsCloseRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ThreadsCloseRequest;
  static deserializeBinaryFromReader(message: ThreadsCloseRequest, reader: jspb.BinaryReader): ThreadsCloseRequest;
}

export namespace ThreadsCloseRequest {
  export type AsObject = {
    request?: proto_identityhub_v1_requests_pb.Request.AsObject,
  }
}

export class ThreadsCloseResponse extends jspb.Message {
  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): proto_identityhub_v1_responses_pb.Response | undefined;
  setResponse(value?: proto_identityhub_v1_responses_pb.Response): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ThreadsCloseResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ThreadsCloseResponse): ThreadsCloseResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ThreadsCloseResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ThreadsCloseResponse;
  static deserializeBinaryFromReader(message: ThreadsCloseResponse, reader: jspb.BinaryReader): ThreadsCloseResponse;
}

export namespace ThreadsCloseResponse {
  export type AsObject = {
    response?: proto_identityhub_v1_responses_pb.Response.AsObject,
  }
}

export class ThreadsDeleteRequest extends jspb.Message {
  hasRequest(): boolean;
  clearRequest(): void;
  getRequest(): proto_identityhub_v1_requests_pb.Request | undefined;
  setRequest(value?: proto_identityhub_v1_requests_pb.Request): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ThreadsDeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ThreadsDeleteRequest): ThreadsDeleteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ThreadsDeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ThreadsDeleteRequest;
  static deserializeBinaryFromReader(message: ThreadsDeleteRequest, reader: jspb.BinaryReader): ThreadsDeleteRequest;
}

export namespace ThreadsDeleteRequest {
  export type AsObject = {
    request?: proto_identityhub_v1_requests_pb.Request.AsObject,
  }
}

export class ThreadsDeleteResponse extends jspb.Message {
  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): proto_identityhub_v1_responses_pb.Response | undefined;
  setResponse(value?: proto_identityhub_v1_responses_pb.Response): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ThreadsDeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ThreadsDeleteResponse): ThreadsDeleteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ThreadsDeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ThreadsDeleteResponse;
  static deserializeBinaryFromReader(message: ThreadsDeleteResponse, reader: jspb.BinaryReader): ThreadsDeleteResponse;
}

export namespace ThreadsDeleteResponse {
  export type AsObject = {
    response?: proto_identityhub_v1_responses_pb.Response.AsObject,
  }
}

