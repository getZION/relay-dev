// package: proto.zion.v1
// file: proto/zion/v1/messages.proto

import * as jspb from "google-protobuf";
import * as proto_protoc_gen_gorm_options_gorm_pb from "../../../proto/protoc-gen-gorm/options/gorm_pb";

export class Message extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getZid(): string;
  setZid(value: string): void;

  getUserDid(): number;
  setUserDid(value: number): void;

  getReceivingUserDid(): number;
  setReceivingUserDid(value: number): void;

  getConversationZid(): number;
  setConversationZid(value: number): void;

  getReplyToMessageZid(): number;
  setReplyToMessageZid(value: number): void;

  getText(): string;
  setText(value: string): void;

  getLink(): string;
  setLink(value: string): void;

  getImg(): string;
  setImg(value: string): void;

  getVideo(): string;
  setVideo(value: string): void;

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
    id: number,
    zid: string,
    userDid: number,
    receivingUserDid: number,
    conversationZid: number,
    replyToMessageZid: number,
    text: string,
    link: string,
    img: string,
    video: string,
  }
}

