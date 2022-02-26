// package: proto.zion.v1
// file: proto/zion/v1/nodeinfo.proto

import * as jspb from "google-protobuf";

export class GetNodeInfoRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetNodeInfoRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetNodeInfoRequest): GetNodeInfoRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetNodeInfoRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetNodeInfoRequest;
  static deserializeBinaryFromReader(message: GetNodeInfoRequest, reader: jspb.BinaryReader): GetNodeInfoRequest;
}

export namespace GetNodeInfoRequest {
  export type AsObject = {
  }
}

export class GetNodeInfoResponse extends jspb.Message {
  getPubkey(): string;
  setPubkey(value: string): void;

  getBalance(): number;
  setBalance(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetNodeInfoResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetNodeInfoResponse): GetNodeInfoResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetNodeInfoResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetNodeInfoResponse;
  static deserializeBinaryFromReader(message: GetNodeInfoResponse, reader: jspb.BinaryReader): GetNodeInfoResponse;
}

export namespace GetNodeInfoResponse {
  export type AsObject = {
    pubkey: string,
    balance: number,
  }
}

