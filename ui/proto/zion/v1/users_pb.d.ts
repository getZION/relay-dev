// package: proto.zion.v1
// file: proto/zion/v1/users.proto

import * as jspb from "google-protobuf";
import * as proto_protoc_gen_gorm_options_gorm_pb from "../../../proto/protoc-gen-gorm/options/gorm_pb";

export class User extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getDid(): string;
  setDid(value: string): void;

  getUsername(): string;
  setUsername(value: string): void;

  getName(): string;
  setName(value: string): void;

  getEmail(): string;
  setEmail(value: string): void;

  getBio(): string;
  setBio(value: string): void;

  getPicture(): string;
  setPicture(value: string): void;

  getPriceToMessage(): number;
  setPriceToMessage(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: number,
    did: string,
    username: string,
    name: string,
    email: string,
    bio: string,
    picture: string,
    priceToMessage: number,
  }
}

