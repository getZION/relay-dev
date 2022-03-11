// package: proto.zion.v1
// file: proto/zion/v1/payments.proto

import * as jspb from "google-protobuf";
import * as proto_protoc_gen_gorm_options_gorm_pb from "../../../proto/protoc-gen-gorm/options/gorm_pb";

export class Payment extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getZid(): string;
  setZid(value: string): void;

  getSenderDid(): string;
  setSenderDid(value: string): void;

  getRecipientDid(): string;
  setRecipientDid(value: string): void;

  getRecipientNodePubkey(): string;
  setRecipientNodePubkey(value: string): void;

  getRecipientRelayUrl(): string;
  setRecipientRelayUrl(value: string): void;

  getStatus(): string;
  setStatus(value: string): void;

  getAmount(): number;
  setAmount(value: number): void;

  getType(): number;
  setType(value: number): void;

  getMemo(): string;
  setMemo(value: string): void;

  getMessageZid(): string;
  setMessageZid(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Payment.AsObject;
  static toObject(includeInstance: boolean, msg: Payment): Payment.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Payment, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Payment;
  static deserializeBinaryFromReader(message: Payment, reader: jspb.BinaryReader): Payment;
}

export namespace Payment {
  export type AsObject = {
    id: number,
    zid: string,
    senderDid: string,
    recipientDid: string,
    recipientNodePubkey: string,
    recipientRelayUrl: string,
    status: string,
    amount: number,
    type: number,
    memo: string,
    messageZid: string,
  }
}

