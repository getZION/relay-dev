// package: proto.zion.v1
// file: proto/zion/v1/payments.proto

import * as jspb from "google-protobuf";
import * as proto_protoc_gen_gorm_options_gorm_pb from "../../../proto/protoc-gen-gorm/options/gorm_pb";

export class Invoice extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getText(): string;
  setText(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Invoice.AsObject;
  static toObject(includeInstance: boolean, msg: Invoice): Invoice.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Invoice, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Invoice;
  static deserializeBinaryFromReader(message: Invoice, reader: jspb.BinaryReader): Invoice;
}

export namespace Invoice {
  export type AsObject = {
    id: number,
    text: string,
  }
}

export class CreateInvoiceRequest extends jspb.Message {
  hasPayload(): boolean;
  clearPayload(): void;
  getPayload(): Invoice | undefined;
  setPayload(value?: Invoice): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateInvoiceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateInvoiceRequest): CreateInvoiceRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateInvoiceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateInvoiceRequest;
  static deserializeBinaryFromReader(message: CreateInvoiceRequest, reader: jspb.BinaryReader): CreateInvoiceRequest;
}

export namespace CreateInvoiceRequest {
  export type AsObject = {
    payload?: Invoice.AsObject,
  }
}

export class CreateInvoiceResponse extends jspb.Message {
  hasResult(): boolean;
  clearResult(): void;
  getResult(): Invoice | undefined;
  setResult(value?: Invoice): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateInvoiceResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateInvoiceResponse): CreateInvoiceResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateInvoiceResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateInvoiceResponse;
  static deserializeBinaryFromReader(message: CreateInvoiceResponse, reader: jspb.BinaryReader): CreateInvoiceResponse;
}

export namespace CreateInvoiceResponse {
  export type AsObject = {
    result?: Invoice.AsObject,
  }
}

export class GetPaymentsHistoryRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPaymentsHistoryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetPaymentsHistoryRequest): GetPaymentsHistoryRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetPaymentsHistoryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPaymentsHistoryRequest;
  static deserializeBinaryFromReader(message: GetPaymentsHistoryRequest, reader: jspb.BinaryReader): GetPaymentsHistoryRequest;
}

export namespace GetPaymentsHistoryRequest {
  export type AsObject = {
  }
}

export class GetPaymentsHistoryResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPaymentsHistoryResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetPaymentsHistoryResponse): GetPaymentsHistoryResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetPaymentsHistoryResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPaymentsHistoryResponse;
  static deserializeBinaryFromReader(message: GetPaymentsHistoryResponse, reader: jspb.BinaryReader): GetPaymentsHistoryResponse;
}

export namespace GetPaymentsHistoryResponse {
  export type AsObject = {
  }
}

export class PayInvoiceRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PayInvoiceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PayInvoiceRequest): PayInvoiceRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PayInvoiceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PayInvoiceRequest;
  static deserializeBinaryFromReader(message: PayInvoiceRequest, reader: jspb.BinaryReader): PayInvoiceRequest;
}

export namespace PayInvoiceRequest {
  export type AsObject = {
  }
}

export class PayInvoiceResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PayInvoiceResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PayInvoiceResponse): PayInvoiceResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PayInvoiceResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PayInvoiceResponse;
  static deserializeBinaryFromReader(message: PayInvoiceResponse, reader: jspb.BinaryReader): PayInvoiceResponse;
}

export namespace PayInvoiceResponse {
  export type AsObject = {
  }
}

export class PayUserRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PayUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PayUserRequest): PayUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PayUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PayUserRequest;
  static deserializeBinaryFromReader(message: PayUserRequest, reader: jspb.BinaryReader): PayUserRequest;
}

export namespace PayUserRequest {
  export type AsObject = {
  }
}

export class PayUserResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PayUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PayUserResponse): PayUserResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PayUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PayUserResponse;
  static deserializeBinaryFromReader(message: PayUserResponse, reader: jspb.BinaryReader): PayUserResponse;
}

export namespace PayUserResponse {
  export type AsObject = {
  }
}

