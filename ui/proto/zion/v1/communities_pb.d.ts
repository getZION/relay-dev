// package: proto.zion.v1
// file: proto/zion/v1/communities.proto

import * as jspb from "google-protobuf";
import * as proto_protoc_gen_gorm_options_gorm_pb from "../../../proto/protoc-gen-gorm/options/gorm_pb";

export class Community extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getZionId(): string;
  setZionId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getOwnerDid(): string;
  setOwnerDid(value: string): void;

  getOwnerUsername(): string;
  setOwnerUsername(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  getImg(): string;
  setImg(value: string): void;

  clearTagsList(): void;
  getTagsList(): Array<string>;
  setTagsList(value: Array<string>): void;
  addTags(value: string, index?: number): string;

  getPriceToJoin(): number;
  setPriceToJoin(value: number): void;

  getPricePerMessage(): number;
  setPricePerMessage(value: number): void;

  getEscrowAmount(): number;
  setEscrowAmount(value: number): void;

  getLastActive(): number;
  setLastActive(value: number): void;

  getPublic(): boolean;
  setPublic(value: boolean): void;

  getDeleted(): boolean;
  setDeleted(value: boolean): void;

  getCreated(): number;
  setCreated(value: number): void;

  getUpdated(): number;
  setUpdated(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Community.AsObject;
  static toObject(includeInstance: boolean, msg: Community): Community.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Community, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Community;
  static deserializeBinaryFromReader(message: Community, reader: jspb.BinaryReader): Community;
}

export namespace Community {
  export type AsObject = {
    id: number,
    zionId: string,
    name: string,
    ownerDid: string,
    ownerUsername: string,
    description: string,
    img: string,
    tagsList: Array<string>,
    priceToJoin: number,
    pricePerMessage: number,
    escrowAmount: number,
    lastActive: number,
    pb_public: boolean,
    deleted: boolean,
    created: number,
    updated: number,
  }
}

export class CreateCommunityRequest extends jspb.Message {
  hasPayload(): boolean;
  clearPayload(): void;
  getPayload(): Community | undefined;
  setPayload(value?: Community): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateCommunityRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateCommunityRequest): CreateCommunityRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateCommunityRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateCommunityRequest;
  static deserializeBinaryFromReader(message: CreateCommunityRequest, reader: jspb.BinaryReader): CreateCommunityRequest;
}

export namespace CreateCommunityRequest {
  export type AsObject = {
    payload?: Community.AsObject,
  }
}

export class CreateCommunityResponse extends jspb.Message {
  hasResult(): boolean;
  clearResult(): void;
  getResult(): Community | undefined;
  setResult(value?: Community): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateCommunityResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateCommunityResponse): CreateCommunityResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateCommunityResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateCommunityResponse;
  static deserializeBinaryFromReader(message: CreateCommunityResponse, reader: jspb.BinaryReader): CreateCommunityResponse;
}

export namespace CreateCommunityResponse {
  export type AsObject = {
    result?: Community.AsObject,
  }
}

export class DeleteCommunityRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteCommunityRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteCommunityRequest): DeleteCommunityRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteCommunityRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteCommunityRequest;
  static deserializeBinaryFromReader(message: DeleteCommunityRequest, reader: jspb.BinaryReader): DeleteCommunityRequest;
}

export namespace DeleteCommunityRequest {
  export type AsObject = {
    id: number,
  }
}

export class DeleteCommunityResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteCommunityResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteCommunityResponse): DeleteCommunityResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteCommunityResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteCommunityResponse;
  static deserializeBinaryFromReader(message: DeleteCommunityResponse, reader: jspb.BinaryReader): DeleteCommunityResponse;
}

export namespace DeleteCommunityResponse {
  export type AsObject = {
    success: boolean,
    message: string,
  }
}

export class EditCommunityRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EditCommunityRequest.AsObject;
  static toObject(includeInstance: boolean, msg: EditCommunityRequest): EditCommunityRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: EditCommunityRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EditCommunityRequest;
  static deserializeBinaryFromReader(message: EditCommunityRequest, reader: jspb.BinaryReader): EditCommunityRequest;
}

export namespace EditCommunityRequest {
  export type AsObject = {
    id: number,
  }
}

export class EditCommunityResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EditCommunityResponse.AsObject;
  static toObject(includeInstance: boolean, msg: EditCommunityResponse): EditCommunityResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: EditCommunityResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EditCommunityResponse;
  static deserializeBinaryFromReader(message: EditCommunityResponse, reader: jspb.BinaryReader): EditCommunityResponse;
}

export namespace EditCommunityResponse {
  export type AsObject = {
    success: boolean,
    message: string,
  }
}

export class HideCommunityRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HideCommunityRequest.AsObject;
  static toObject(includeInstance: boolean, msg: HideCommunityRequest): HideCommunityRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HideCommunityRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HideCommunityRequest;
  static deserializeBinaryFromReader(message: HideCommunityRequest, reader: jspb.BinaryReader): HideCommunityRequest;
}

export namespace HideCommunityRequest {
  export type AsObject = {
    id: number,
  }
}

export class HideCommunityResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HideCommunityResponse.AsObject;
  static toObject(includeInstance: boolean, msg: HideCommunityResponse): HideCommunityResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HideCommunityResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HideCommunityResponse;
  static deserializeBinaryFromReader(message: HideCommunityResponse, reader: jspb.BinaryReader): HideCommunityResponse;
}

export namespace HideCommunityResponse {
  export type AsObject = {
    success: boolean,
    message: string,
  }
}

export class JoinCommunityRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JoinCommunityRequest.AsObject;
  static toObject(includeInstance: boolean, msg: JoinCommunityRequest): JoinCommunityRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: JoinCommunityRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JoinCommunityRequest;
  static deserializeBinaryFromReader(message: JoinCommunityRequest, reader: jspb.BinaryReader): JoinCommunityRequest;
}

export namespace JoinCommunityRequest {
  export type AsObject = {
    id: number,
  }
}

export class JoinCommunityResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JoinCommunityResponse.AsObject;
  static toObject(includeInstance: boolean, msg: JoinCommunityResponse): JoinCommunityResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: JoinCommunityResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JoinCommunityResponse;
  static deserializeBinaryFromReader(message: JoinCommunityResponse, reader: jspb.BinaryReader): JoinCommunityResponse;
}

export namespace JoinCommunityResponse {
  export type AsObject = {
    success: boolean,
    message: string,
  }
}

export class ListCommunityRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCommunityRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListCommunityRequest): ListCommunityRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListCommunityRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCommunityRequest;
  static deserializeBinaryFromReader(message: ListCommunityRequest, reader: jspb.BinaryReader): ListCommunityRequest;
}

export namespace ListCommunityRequest {
  export type AsObject = {
    id: number,
  }
}

export class ListCommunityResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  getMessage(): string;
  setMessage(value: string): void;

  clearResultsList(): void;
  getResultsList(): Array<Community>;
  setResultsList(value: Array<Community>): void;
  addResults(value?: Community, index?: number): Community;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCommunityResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListCommunityResponse): ListCommunityResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListCommunityResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCommunityResponse;
  static deserializeBinaryFromReader(message: ListCommunityResponse, reader: jspb.BinaryReader): ListCommunityResponse;
}

export namespace ListCommunityResponse {
  export type AsObject = {
    success: boolean,
    message: string,
    resultsList: Array<Community.AsObject>,
  }
}

