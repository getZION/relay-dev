// package: proto.identityhub.v1
// file: proto/identityhub/v1/threads.proto

import * as proto_identityhub_v1_threads_pb from "../../../proto/identityhub/v1/threads_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ThreadsServiceThreadsQuery = {
  readonly methodName: string;
  readonly service: typeof ThreadsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_identityhub_v1_threads_pb.ThreadsQueryRequest;
  readonly responseType: typeof proto_identityhub_v1_threads_pb.ThreadsQueryResponse;
};

type ThreadsServiceThreadsCreate = {
  readonly methodName: string;
  readonly service: typeof ThreadsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_identityhub_v1_threads_pb.ThreadsCreateRequest;
  readonly responseType: typeof proto_identityhub_v1_threads_pb.ThreadsCreateResponse;
};

type ThreadsServiceThreadsReply = {
  readonly methodName: string;
  readonly service: typeof ThreadsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_identityhub_v1_threads_pb.ThreadsReplyRequest;
  readonly responseType: typeof proto_identityhub_v1_threads_pb.ThreadsReplyResponse;
};

type ThreadsServiceThreadsClose = {
  readonly methodName: string;
  readonly service: typeof ThreadsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_identityhub_v1_threads_pb.ThreadsCloseRequest;
  readonly responseType: typeof proto_identityhub_v1_threads_pb.ThreadsCloseResponse;
};

type ThreadsServiceThreadsDelete = {
  readonly methodName: string;
  readonly service: typeof ThreadsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_identityhub_v1_threads_pb.ThreadsDeleteRequest;
  readonly responseType: typeof proto_identityhub_v1_threads_pb.ThreadsDeleteResponse;
};

export class ThreadsService {
  static readonly serviceName: string;
  static readonly ThreadsQuery: ThreadsServiceThreadsQuery;
  static readonly ThreadsCreate: ThreadsServiceThreadsCreate;
  static readonly ThreadsReply: ThreadsServiceThreadsReply;
  static readonly ThreadsClose: ThreadsServiceThreadsClose;
  static readonly ThreadsDelete: ThreadsServiceThreadsDelete;
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

export class ThreadsServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  threadsQuery(
    requestMessage: proto_identityhub_v1_threads_pb.ThreadsQueryRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_threads_pb.ThreadsQueryResponse|null) => void
  ): UnaryResponse;
  threadsQuery(
    requestMessage: proto_identityhub_v1_threads_pb.ThreadsQueryRequest,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_threads_pb.ThreadsQueryResponse|null) => void
  ): UnaryResponse;
  threadsCreate(
    requestMessage: proto_identityhub_v1_threads_pb.ThreadsCreateRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_threads_pb.ThreadsCreateResponse|null) => void
  ): UnaryResponse;
  threadsCreate(
    requestMessage: proto_identityhub_v1_threads_pb.ThreadsCreateRequest,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_threads_pb.ThreadsCreateResponse|null) => void
  ): UnaryResponse;
  threadsReply(
    requestMessage: proto_identityhub_v1_threads_pb.ThreadsReplyRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_threads_pb.ThreadsReplyResponse|null) => void
  ): UnaryResponse;
  threadsReply(
    requestMessage: proto_identityhub_v1_threads_pb.ThreadsReplyRequest,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_threads_pb.ThreadsReplyResponse|null) => void
  ): UnaryResponse;
  threadsClose(
    requestMessage: proto_identityhub_v1_threads_pb.ThreadsCloseRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_threads_pb.ThreadsCloseResponse|null) => void
  ): UnaryResponse;
  threadsClose(
    requestMessage: proto_identityhub_v1_threads_pb.ThreadsCloseRequest,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_threads_pb.ThreadsCloseResponse|null) => void
  ): UnaryResponse;
  threadsDelete(
    requestMessage: proto_identityhub_v1_threads_pb.ThreadsDeleteRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_threads_pb.ThreadsDeleteResponse|null) => void
  ): UnaryResponse;
  threadsDelete(
    requestMessage: proto_identityhub_v1_threads_pb.ThreadsDeleteRequest,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_threads_pb.ThreadsDeleteResponse|null) => void
  ): UnaryResponse;
}

