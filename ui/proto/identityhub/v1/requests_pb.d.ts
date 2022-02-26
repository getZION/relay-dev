// package: proto.identityhub.v1
// file: proto/identityhub/v1/requests.proto

import * as jspb from "google-protobuf";
import * as proto_identityhub_v1_messages_pb from "../../../proto/identityhub/v1/messages_pb";

export class Request extends jspb.Message {
  getRequestid(): string;
  setRequestid(value: string): void;

  getTarget(): string;
  setTarget(value: string): void;

  clearMessagesList(): void;
  getMessagesList(): Array<proto_identityhub_v1_messages_pb.Message>;
  setMessagesList(value: Array<proto_identityhub_v1_messages_pb.Message>): void;
  addMessages(value?: proto_identityhub_v1_messages_pb.Message, index?: number): proto_identityhub_v1_messages_pb.Message;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Request.AsObject;
  static toObject(includeInstance: boolean, msg: Request): Request.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Request, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Request;
  static deserializeBinaryFromReader(message: Request, reader: jspb.BinaryReader): Request;
}

export namespace Request {
  export type AsObject = {
    requestid: string,
    target: string,
    messagesList: Array<proto_identityhub_v1_messages_pb.Message.AsObject>,
  }
}

