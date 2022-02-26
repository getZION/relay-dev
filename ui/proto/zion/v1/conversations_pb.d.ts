// package: proto.zion.v1
// file: proto/zion/v1/conversations.proto

import * as jspb from "google-protobuf";
import * as proto_protoc_gen_gorm_options_gorm_pb from "../../../proto/protoc-gen-gorm/options/gorm_pb";

export class Conversation extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getText(): string;
  setText(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Conversation.AsObject;
  static toObject(includeInstance: boolean, msg: Conversation): Conversation.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Conversation, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Conversation;
  static deserializeBinaryFromReader(message: Conversation, reader: jspb.BinaryReader): Conversation;
}

export namespace Conversation {
  export type AsObject = {
    id: number,
    text: string,
  }
}

export class Comment extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getText(): string;
  setText(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Comment.AsObject;
  static toObject(includeInstance: boolean, msg: Comment): Comment.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Comment, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Comment;
  static deserializeBinaryFromReader(message: Comment, reader: jspb.BinaryReader): Comment;
}

export namespace Comment {
  export type AsObject = {
    id: number,
    text: string,
  }
}

export class CreateConversationRequest extends jspb.Message {
  hasPayload(): boolean;
  clearPayload(): void;
  getPayload(): Conversation | undefined;
  setPayload(value?: Conversation): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateConversationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateConversationRequest): CreateConversationRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateConversationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateConversationRequest;
  static deserializeBinaryFromReader(message: CreateConversationRequest, reader: jspb.BinaryReader): CreateConversationRequest;
}

export namespace CreateConversationRequest {
  export type AsObject = {
    payload?: Conversation.AsObject,
  }
}

export class CreateConversationResponse extends jspb.Message {
  hasResult(): boolean;
  clearResult(): void;
  getResult(): Conversation | undefined;
  setResult(value?: Conversation): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateConversationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateConversationResponse): CreateConversationResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateConversationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateConversationResponse;
  static deserializeBinaryFromReader(message: CreateConversationResponse, reader: jspb.BinaryReader): CreateConversationResponse;
}

export namespace CreateConversationResponse {
  export type AsObject = {
    result?: Conversation.AsObject,
  }
}

export class DeleteCommentRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteCommentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteCommentRequest): DeleteCommentRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteCommentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteCommentRequest;
  static deserializeBinaryFromReader(message: DeleteCommentRequest, reader: jspb.BinaryReader): DeleteCommentRequest;
}

export namespace DeleteCommentRequest {
  export type AsObject = {
    id: number,
  }
}

export class DeleteCommentResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteCommentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteCommentResponse): DeleteCommentResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteCommentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteCommentResponse;
  static deserializeBinaryFromReader(message: DeleteCommentResponse, reader: jspb.BinaryReader): DeleteCommentResponse;
}

export namespace DeleteCommentResponse {
  export type AsObject = {
  }
}

export class DeleteConversationRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteConversationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteConversationRequest): DeleteConversationRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteConversationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteConversationRequest;
  static deserializeBinaryFromReader(message: DeleteConversationRequest, reader: jspb.BinaryReader): DeleteConversationRequest;
}

export namespace DeleteConversationRequest {
  export type AsObject = {
    id: number,
  }
}

export class DeleteConversationResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteConversationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteConversationResponse): DeleteConversationResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteConversationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteConversationResponse;
  static deserializeBinaryFromReader(message: DeleteConversationResponse, reader: jspb.BinaryReader): DeleteConversationResponse;
}

export namespace DeleteConversationResponse {
  export type AsObject = {
  }
}

export class GetActivityRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetActivityRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetActivityRequest): GetActivityRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetActivityRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetActivityRequest;
  static deserializeBinaryFromReader(message: GetActivityRequest, reader: jspb.BinaryReader): GetActivityRequest;
}

export namespace GetActivityRequest {
  export type AsObject = {
  }
}

export class GetActivityResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetActivityResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetActivityResponse): GetActivityResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetActivityResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetActivityResponse;
  static deserializeBinaryFromReader(message: GetActivityResponse, reader: jspb.BinaryReader): GetActivityResponse;
}

export namespace GetActivityResponse {
  export type AsObject = {
  }
}

export class GetConversationRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetConversationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetConversationRequest): GetConversationRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetConversationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetConversationRequest;
  static deserializeBinaryFromReader(message: GetConversationRequest, reader: jspb.BinaryReader): GetConversationRequest;
}

export namespace GetConversationRequest {
  export type AsObject = {
  }
}

export class GetConversationResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetConversationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetConversationResponse): GetConversationResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetConversationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetConversationResponse;
  static deserializeBinaryFromReader(message: GetConversationResponse, reader: jspb.BinaryReader): GetConversationResponse;
}

export namespace GetConversationResponse {
  export type AsObject = {
  }
}

export class GetFeedRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetFeedRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetFeedRequest): GetFeedRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetFeedRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetFeedRequest;
  static deserializeBinaryFromReader(message: GetFeedRequest, reader: jspb.BinaryReader): GetFeedRequest;
}

export namespace GetFeedRequest {
  export type AsObject = {
  }
}

export class GetFeedResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetFeedResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetFeedResponse): GetFeedResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetFeedResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetFeedResponse;
  static deserializeBinaryFromReader(message: GetFeedResponse, reader: jspb.BinaryReader): GetFeedResponse;
}

export namespace GetFeedResponse {
  export type AsObject = {
  }
}

