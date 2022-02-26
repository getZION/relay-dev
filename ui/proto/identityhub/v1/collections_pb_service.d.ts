// package: proto.identityhub.v1
// file: proto/identityhub/v1/collections.proto

import * as proto_identityhub_v1_collections_pb from "../../../proto/identityhub/v1/collections_pb";
import {grpc} from "@improbable-eng/grpc-web";

type CollectionsServiceCollectionsQuery = {
  readonly methodName: string;
  readonly service: typeof CollectionsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_identityhub_v1_collections_pb.CollectionsQueryRequest;
  readonly responseType: typeof proto_identityhub_v1_collections_pb.CollectionsQueryResponse;
};

type CollectionsServiceCollectionsWrite = {
  readonly methodName: string;
  readonly service: typeof CollectionsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_identityhub_v1_collections_pb.CollectionsWriteRequest;
  readonly responseType: typeof proto_identityhub_v1_collections_pb.CollectionsWriteResponse;
};

type CollectionsServiceCollectionsCommit = {
  readonly methodName: string;
  readonly service: typeof CollectionsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_identityhub_v1_collections_pb.CollectionsCommitRequest;
  readonly responseType: typeof proto_identityhub_v1_collections_pb.CollectionsCommitResponse;
};

type CollectionsServiceCollectionsDelete = {
  readonly methodName: string;
  readonly service: typeof CollectionsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_identityhub_v1_collections_pb.CollectionsDeleteRequest;
  readonly responseType: typeof proto_identityhub_v1_collections_pb.CollectionsDeleteResponse;
};

export class CollectionsService {
  static readonly serviceName: string;
  static readonly CollectionsQuery: CollectionsServiceCollectionsQuery;
  static readonly CollectionsWrite: CollectionsServiceCollectionsWrite;
  static readonly CollectionsCommit: CollectionsServiceCollectionsCommit;
  static readonly CollectionsDelete: CollectionsServiceCollectionsDelete;
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

export class CollectionsServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  collectionsQuery(
    requestMessage: proto_identityhub_v1_collections_pb.CollectionsQueryRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_collections_pb.CollectionsQueryResponse|null) => void
  ): UnaryResponse;
  collectionsQuery(
    requestMessage: proto_identityhub_v1_collections_pb.CollectionsQueryRequest,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_collections_pb.CollectionsQueryResponse|null) => void
  ): UnaryResponse;
  collectionsWrite(
    requestMessage: proto_identityhub_v1_collections_pb.CollectionsWriteRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_collections_pb.CollectionsWriteResponse|null) => void
  ): UnaryResponse;
  collectionsWrite(
    requestMessage: proto_identityhub_v1_collections_pb.CollectionsWriteRequest,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_collections_pb.CollectionsWriteResponse|null) => void
  ): UnaryResponse;
  collectionsCommit(
    requestMessage: proto_identityhub_v1_collections_pb.CollectionsCommitRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_collections_pb.CollectionsCommitResponse|null) => void
  ): UnaryResponse;
  collectionsCommit(
    requestMessage: proto_identityhub_v1_collections_pb.CollectionsCommitRequest,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_collections_pb.CollectionsCommitResponse|null) => void
  ): UnaryResponse;
  collectionsDelete(
    requestMessage: proto_identityhub_v1_collections_pb.CollectionsDeleteRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_collections_pb.CollectionsDeleteResponse|null) => void
  ): UnaryResponse;
  collectionsDelete(
    requestMessage: proto_identityhub_v1_collections_pb.CollectionsDeleteRequest,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_collections_pb.CollectionsDeleteResponse|null) => void
  ): UnaryResponse;
}

