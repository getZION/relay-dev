// package: proto.identityhub.v1
// file: proto/identityhub/v1/responses.proto

import * as jspb from "google-protobuf";

export class Response extends jspb.Message {
  getRequestid(): string;
  setRequestid(value: string): void;

  hasStatus(): boolean;
  clearStatus(): void;
  getStatus(): Status | undefined;
  setStatus(value?: Status): void;

  clearRepliesList(): void;
  getRepliesList(): Array<Reply>;
  setRepliesList(value: Array<Reply>): void;
  addReplies(value?: Reply, index?: number): Reply;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Response.AsObject;
  static toObject(includeInstance: boolean, msg: Response): Response.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Response, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Response;
  static deserializeBinaryFromReader(message: Response, reader: jspb.BinaryReader): Response;
}

export namespace Response {
  export type AsObject = {
    requestid: string,
    status?: Status.AsObject,
    repliesList: Array<Reply.AsObject>,
  }
}

export class Status extends jspb.Message {
  getCode(): number;
  setCode(value: number): void;

  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Status.AsObject;
  static toObject(includeInstance: boolean, msg: Status): Status.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Status, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Status;
  static deserializeBinaryFromReader(message: Status, reader: jspb.BinaryReader): Status;
}

export namespace Status {
  export type AsObject = {
    code: number,
    message: string,
  }
}

export class Reply extends jspb.Message {
  getMessageid(): string;
  setMessageid(value: string): void;

  hasStatus(): boolean;
  clearStatus(): void;
  getStatus(): Status | undefined;
  setStatus(value?: Status): void;

  clearEntriesList(): void;
  getEntriesList(): Array<string>;
  setEntriesList(value: Array<string>): void;
  addEntries(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Reply.AsObject;
  static toObject(includeInstance: boolean, msg: Reply): Reply.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Reply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Reply;
  static deserializeBinaryFromReader(message: Reply, reader: jspb.BinaryReader): Reply;
}

export namespace Reply {
  export type AsObject = {
    messageid: string,
    status?: Status.AsObject,
    entriesList: Array<string>,
  }
}

