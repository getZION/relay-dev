// package: proto.zion.v1
// file: proto/zion/v1/contacts.proto

import * as jspb from "google-protobuf";
import * as proto_protoc_gen_gorm_options_gorm_pb from "../../../proto/protoc-gen-gorm/options/gorm_pb";

export class Contact extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Contact.AsObject;
  static toObject(includeInstance: boolean, msg: Contact): Contact.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Contact, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Contact;
  static deserializeBinaryFromReader(message: Contact, reader: jspb.BinaryReader): Contact;
}

export namespace Contact {
  export type AsObject = {
    id: number,
    name: string,
  }
}

export class AddByPubkeyRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddByPubkeyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddByPubkeyRequest): AddByPubkeyRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AddByPubkeyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddByPubkeyRequest;
  static deserializeBinaryFromReader(message: AddByPubkeyRequest, reader: jspb.BinaryReader): AddByPubkeyRequest;
}

export namespace AddByPubkeyRequest {
  export type AsObject = {
  }
}

export class AddByPubkeyResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddByPubkeyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddByPubkeyResponse): AddByPubkeyResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AddByPubkeyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddByPubkeyResponse;
  static deserializeBinaryFromReader(message: AddByPubkeyResponse, reader: jspb.BinaryReader): AddByPubkeyResponse;
}

export namespace AddByPubkeyResponse {
  export type AsObject = {
  }
}

export class AddByUsernameRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddByUsernameRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddByUsernameRequest): AddByUsernameRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AddByUsernameRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddByUsernameRequest;
  static deserializeBinaryFromReader(message: AddByUsernameRequest, reader: jspb.BinaryReader): AddByUsernameRequest;
}

export namespace AddByUsernameRequest {
  export type AsObject = {
  }
}

export class AddByUsernameResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddByUsernameResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddByUsernameResponse): AddByUsernameResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AddByUsernameResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddByUsernameResponse;
  static deserializeBinaryFromReader(message: AddByUsernameResponse, reader: jspb.BinaryReader): AddByUsernameResponse;
}

export namespace AddByUsernameResponse {
  export type AsObject = {
  }
}

export class DeleteContactRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteContactRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteContactRequest): DeleteContactRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteContactRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteContactRequest;
  static deserializeBinaryFromReader(message: DeleteContactRequest, reader: jspb.BinaryReader): DeleteContactRequest;
}

export namespace DeleteContactRequest {
  export type AsObject = {
    id: number,
  }
}

export class DeleteContactResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteContactResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteContactResponse): DeleteContactResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteContactResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteContactResponse;
  static deserializeBinaryFromReader(message: DeleteContactResponse, reader: jspb.BinaryReader): DeleteContactResponse;
}

export namespace DeleteContactResponse {
  export type AsObject = {
  }
}

export class GetMyContactsRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetMyContactsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetMyContactsRequest): GetMyContactsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetMyContactsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetMyContactsRequest;
  static deserializeBinaryFromReader(message: GetMyContactsRequest, reader: jspb.BinaryReader): GetMyContactsRequest;
}

export namespace GetMyContactsRequest {
  export type AsObject = {
  }
}

export class GetMyContactsResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetMyContactsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetMyContactsResponse): GetMyContactsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetMyContactsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetMyContactsResponse;
  static deserializeBinaryFromReader(message: GetMyContactsResponse, reader: jspb.BinaryReader): GetMyContactsResponse;
}

export namespace GetMyContactsResponse {
  export type AsObject = {
  }
}

