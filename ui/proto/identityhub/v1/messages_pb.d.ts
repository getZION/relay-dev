// package: proto.identityhub.v1
// file: proto/identityhub/v1/messages.proto

import * as jspb from "google-protobuf";

export class Message extends jspb.Message {
  getData(): string;
  setData(value: string): void;

  hasDescriptor(): boolean;
  clearDescriptor(): void;
  getDescriptor(): MessageDescriptor | undefined;
  setDescriptor(value?: MessageDescriptor): void;

  hasAttestation(): boolean;
  clearAttestation(): void;
  getAttestation(): Attestation | undefined;
  setAttestation(value?: Attestation): void;

  hasAuthorization(): boolean;
  clearAuthorization(): void;
  getAuthorization(): Authorization | undefined;
  setAuthorization(value?: Authorization): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Message.AsObject;
  static toObject(includeInstance: boolean, msg: Message): Message.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Message, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Message;
  static deserializeBinaryFromReader(message: Message, reader: jspb.BinaryReader): Message;
}

export namespace Message {
  export type AsObject = {
    data: string,
    descriptor?: MessageDescriptor.AsObject,
    attestation?: Attestation.AsObject,
    authorization?: Authorization.AsObject,
  }
}

export class MessageDescriptor extends jspb.Message {
  getMethod(): string;
  setMethod(value: string): void;

  getObjectid(): string;
  setObjectid(value: string): void;

  getSchema(): string;
  setSchema(value: string): void;

  getDataformat(): string;
  setDataformat(value: string): void;

  getDatecreated(): string;
  setDatecreated(value: string): void;

  getDatepublished(): string;
  setDatepublished(value: string): void;

  getDatesort(): string;
  setDatesort(value: string): void;

  getRoot(): string;
  setRoot(value: string): void;

  getParent(): string;
  setParent(value: string): void;

  getCid(): string;
  setCid(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MessageDescriptor.AsObject;
  static toObject(includeInstance: boolean, msg: MessageDescriptor): MessageDescriptor.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MessageDescriptor, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MessageDescriptor;
  static deserializeBinaryFromReader(message: MessageDescriptor, reader: jspb.BinaryReader): MessageDescriptor;
}

export namespace MessageDescriptor {
  export type AsObject = {
    method: string,
    objectid: string,
    schema: string,
    dataformat: string,
    datecreated: string,
    datepublished: string,
    datesort: string,
    root: string,
    parent: string,
    cid: string,
  }
}

export class Attestation extends jspb.Message {
  hasProtected(): boolean;
  clearProtected(): void;
  getProtected(): AttestationProtected | undefined;
  setProtected(value?: AttestationProtected): void;

  getPayload(): string;
  setPayload(value: string): void;

  getSignature(): string;
  setSignature(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Attestation.AsObject;
  static toObject(includeInstance: boolean, msg: Attestation): Attestation.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Attestation, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Attestation;
  static deserializeBinaryFromReader(message: Attestation, reader: jspb.BinaryReader): Attestation;
}

export namespace Attestation {
  export type AsObject = {
    pb_protected?: AttestationProtected.AsObject,
    payload: string,
    signature: string,
  }
}

export class AttestationProtected extends jspb.Message {
  getAlg(): string;
  setAlg(value: string): void;

  getKid(): string;
  setKid(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AttestationProtected.AsObject;
  static toObject(includeInstance: boolean, msg: AttestationProtected): AttestationProtected.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AttestationProtected, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AttestationProtected;
  static deserializeBinaryFromReader(message: AttestationProtected, reader: jspb.BinaryReader): AttestationProtected;
}

export namespace AttestationProtected {
  export type AsObject = {
    alg: string,
    kid: string,
  }
}

export class Authorization extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Authorization.AsObject;
  static toObject(includeInstance: boolean, msg: Authorization): Authorization.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Authorization, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Authorization;
  static deserializeBinaryFromReader(message: Authorization, reader: jspb.BinaryReader): Authorization;
}

export namespace Authorization {
  export type AsObject = {
  }
}

