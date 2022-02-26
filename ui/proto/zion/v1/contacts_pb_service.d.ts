// package: proto.zion.v1
// file: proto/zion/v1/contacts.proto

import * as proto_zion_v1_contacts_pb from "../../../proto/zion/v1/contacts_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ContactsServiceAddByPubkey = {
  readonly methodName: string;
  readonly service: typeof ContactsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_contacts_pb.AddByPubkeyRequest;
  readonly responseType: typeof proto_zion_v1_contacts_pb.AddByPubkeyResponse;
};

type ContactsServiceAddByUsername = {
  readonly methodName: string;
  readonly service: typeof ContactsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_contacts_pb.AddByUsernameRequest;
  readonly responseType: typeof proto_zion_v1_contacts_pb.AddByUsernameResponse;
};

type ContactsServiceDeleteContact = {
  readonly methodName: string;
  readonly service: typeof ContactsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_contacts_pb.DeleteContactRequest;
  readonly responseType: typeof proto_zion_v1_contacts_pb.DeleteContactResponse;
};

type ContactsServiceGetMyContacts = {
  readonly methodName: string;
  readonly service: typeof ContactsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_contacts_pb.GetMyContactsRequest;
  readonly responseType: typeof proto_zion_v1_contacts_pb.GetMyContactsResponse;
};

export class ContactsService {
  static readonly serviceName: string;
  static readonly AddByPubkey: ContactsServiceAddByPubkey;
  static readonly AddByUsername: ContactsServiceAddByUsername;
  static readonly DeleteContact: ContactsServiceDeleteContact;
  static readonly GetMyContacts: ContactsServiceGetMyContacts;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class ContactsServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  addByPubkey(
    requestMessage: proto_zion_v1_contacts_pb.AddByPubkeyRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_contacts_pb.AddByPubkeyResponse|null) => void
  ): UnaryResponse;
  addByPubkey(
    requestMessage: proto_zion_v1_contacts_pb.AddByPubkeyRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_contacts_pb.AddByPubkeyResponse|null) => void
  ): UnaryResponse;
  addByUsername(
    requestMessage: proto_zion_v1_contacts_pb.AddByUsernameRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_contacts_pb.AddByUsernameResponse|null) => void
  ): UnaryResponse;
  addByUsername(
    requestMessage: proto_zion_v1_contacts_pb.AddByUsernameRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_contacts_pb.AddByUsernameResponse|null) => void
  ): UnaryResponse;
  deleteContact(
    requestMessage: proto_zion_v1_contacts_pb.DeleteContactRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_contacts_pb.DeleteContactResponse|null) => void
  ): UnaryResponse;
  deleteContact(
    requestMessage: proto_zion_v1_contacts_pb.DeleteContactRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_contacts_pb.DeleteContactResponse|null) => void
  ): UnaryResponse;
  getMyContacts(
    requestMessage: proto_zion_v1_contacts_pb.GetMyContactsRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_contacts_pb.GetMyContactsResponse|null) => void
  ): UnaryResponse;
  getMyContacts(
    requestMessage: proto_zion_v1_contacts_pb.GetMyContactsRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_contacts_pb.GetMyContactsResponse|null) => void
  ): UnaryResponse;
}

