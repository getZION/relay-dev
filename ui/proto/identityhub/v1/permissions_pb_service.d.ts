// package: proto.identityhub.v1
// file: proto/identityhub/v1/permissions.proto

import * as proto_identityhub_v1_permissions_pb from "../../../proto/identityhub/v1/permissions_pb";
import {grpc} from "@improbable-eng/grpc-web";

type PermissionsServicePermissionsRequest = {
  readonly methodName: string;
  readonly service: typeof PermissionsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_identityhub_v1_permissions_pb.PermissionsRequestRequest;
  readonly responseType: typeof proto_identityhub_v1_permissions_pb.PermissionsRequestResponse;
};

type PermissionsServicePermissionsQuery = {
  readonly methodName: string;
  readonly service: typeof PermissionsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_identityhub_v1_permissions_pb.PermissionsQueryRequest;
  readonly responseType: typeof proto_identityhub_v1_permissions_pb.PermissionsQueryResponse;
};

type PermissionsServicePermissionsGrant = {
  readonly methodName: string;
  readonly service: typeof PermissionsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_identityhub_v1_permissions_pb.PermissionsGrantRequest;
  readonly responseType: typeof proto_identityhub_v1_permissions_pb.PermissionsGrantResponse;
};

type PermissionsServicePermissionsRevoke = {
  readonly methodName: string;
  readonly service: typeof PermissionsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_identityhub_v1_permissions_pb.PermissionsRevokeRequest;
  readonly responseType: typeof proto_identityhub_v1_permissions_pb.PermissionsRevokeResponse;
};

export class PermissionsService {
  static readonly serviceName: string;
  static readonly PermissionsRequest: PermissionsServicePermissionsRequest;
  static readonly PermissionsQuery: PermissionsServicePermissionsQuery;
  static readonly PermissionsGrant: PermissionsServicePermissionsGrant;
  static readonly PermissionsRevoke: PermissionsServicePermissionsRevoke;
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

export class PermissionsServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  permissionsRequest(
    requestMessage: proto_identityhub_v1_permissions_pb.PermissionsRequestRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_permissions_pb.PermissionsRequestResponse|null) => void
  ): UnaryResponse;
  permissionsRequest(
    requestMessage: proto_identityhub_v1_permissions_pb.PermissionsRequestRequest,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_permissions_pb.PermissionsRequestResponse|null) => void
  ): UnaryResponse;
  permissionsQuery(
    requestMessage: proto_identityhub_v1_permissions_pb.PermissionsQueryRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_permissions_pb.PermissionsQueryResponse|null) => void
  ): UnaryResponse;
  permissionsQuery(
    requestMessage: proto_identityhub_v1_permissions_pb.PermissionsQueryRequest,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_permissions_pb.PermissionsQueryResponse|null) => void
  ): UnaryResponse;
  permissionsGrant(
    requestMessage: proto_identityhub_v1_permissions_pb.PermissionsGrantRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_permissions_pb.PermissionsGrantResponse|null) => void
  ): UnaryResponse;
  permissionsGrant(
    requestMessage: proto_identityhub_v1_permissions_pb.PermissionsGrantRequest,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_permissions_pb.PermissionsGrantResponse|null) => void
  ): UnaryResponse;
  permissionsRevoke(
    requestMessage: proto_identityhub_v1_permissions_pb.PermissionsRevokeRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_permissions_pb.PermissionsRevokeResponse|null) => void
  ): UnaryResponse;
  permissionsRevoke(
    requestMessage: proto_identityhub_v1_permissions_pb.PermissionsRevokeRequest,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_permissions_pb.PermissionsRevokeResponse|null) => void
  ): UnaryResponse;
}

