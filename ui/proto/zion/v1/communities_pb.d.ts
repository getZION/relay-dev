// package: proto.zion.v1
// file: proto/zion/v1/communities.proto

import * as jspb from "google-protobuf";
import * as proto_protoc_gen_gorm_options_gorm_pb from "../../../proto/protoc-gen-gorm/options/gorm_pb";

export class Community extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getZid(): string;
  setZid(value: string): void;

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
    zid: string,
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

