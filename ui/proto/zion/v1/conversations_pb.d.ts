// package: proto.zion.v1
// file: proto/zion/v1/conversations.proto

import * as jspb from "google-protobuf";
import * as proto_protoc_gen_gorm_options_gorm_pb from "../../../proto/protoc-gen-gorm/options/gorm_pb";
import * as proto_zion_v1_messages_pb from "../../../proto/zion/v1/messages_pb";

export class Conversation extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getZid(): string;
  setZid(value: string): void;

  getCommunityZid(): string;
  setCommunityZid(value: string): void;

  hasMessage(): boolean;
  clearMessage(): void;
  getMessage(): proto_zion_v1_messages_pb.Message | undefined;
  setMessage(value?: proto_zion_v1_messages_pb.Message): void;

  getPublic(): boolean;
  setPublic(value: boolean): void;

  getPublicPrice(): number;
  setPublicPrice(value: number): void;

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
    zid: string,
    communityZid: string,
    message?: proto_zion_v1_messages_pb.Message.AsObject,
    pb_public: boolean,
    publicPrice: number,
  }
}

