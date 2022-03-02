// package: proto.zion.v1
// file: proto/zion/v1/communities.proto

import * as proto_zion_v1_communities_pb from "../../../proto/zion/v1/communities_pb";
import {grpc} from "@improbable-eng/grpc-web";

type CommunitiesServiceCreateCommunity = {
  readonly methodName: string;
  readonly service: typeof CommunitiesService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_communities_pb.CreateCommunityRequest;
  readonly responseType: typeof proto_zion_v1_communities_pb.CreateCommunityResponse;
};

type CommunitiesServiceDeleteCommunity = {
  readonly methodName: string;
  readonly service: typeof CommunitiesService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_communities_pb.DeleteCommunityRequest;
  readonly responseType: typeof proto_zion_v1_communities_pb.DeleteCommunityResponse;
};

type CommunitiesServiceEditCommunity = {
  readonly methodName: string;
  readonly service: typeof CommunitiesService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_communities_pb.EditCommunityRequest;
  readonly responseType: typeof proto_zion_v1_communities_pb.EditCommunityResponse;
};

type CommunitiesServiceJoinCommunity = {
  readonly methodName: string;
  readonly service: typeof CommunitiesService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_communities_pb.JoinCommunityRequest;
  readonly responseType: typeof proto_zion_v1_communities_pb.JoinCommunityResponse;
};

type CommunitiesServiceListCommunity = {
  readonly methodName: string;
  readonly service: typeof CommunitiesService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_communities_pb.ListCommunityRequest;
  readonly responseType: typeof proto_zion_v1_communities_pb.ListCommunityResponse;
};

type CommunitiesServiceReadCommunity = {
  readonly methodName: string;
  readonly service: typeof CommunitiesService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_communities_pb.ReadCommunityRequest;
  readonly responseType: typeof proto_zion_v1_communities_pb.ReadCommunityResponse;
};

export class CommunitiesService {
  static readonly serviceName: string;
  static readonly CreateCommunity: CommunitiesServiceCreateCommunity;
  static readonly DeleteCommunity: CommunitiesServiceDeleteCommunity;
  static readonly EditCommunity: CommunitiesServiceEditCommunity;
  static readonly JoinCommunity: CommunitiesServiceJoinCommunity;
  static readonly ListCommunity: CommunitiesServiceListCommunity;
  static readonly ReadCommunity: CommunitiesServiceReadCommunity;
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

export class CommunitiesServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  createCommunity(
    requestMessage: proto_zion_v1_communities_pb.CreateCommunityRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_communities_pb.CreateCommunityResponse|null) => void
  ): UnaryResponse;
  createCommunity(
    requestMessage: proto_zion_v1_communities_pb.CreateCommunityRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_communities_pb.CreateCommunityResponse|null) => void
  ): UnaryResponse;
  deleteCommunity(
    requestMessage: proto_zion_v1_communities_pb.DeleteCommunityRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_communities_pb.DeleteCommunityResponse|null) => void
  ): UnaryResponse;
  deleteCommunity(
    requestMessage: proto_zion_v1_communities_pb.DeleteCommunityRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_communities_pb.DeleteCommunityResponse|null) => void
  ): UnaryResponse;
  editCommunity(
    requestMessage: proto_zion_v1_communities_pb.EditCommunityRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_communities_pb.EditCommunityResponse|null) => void
  ): UnaryResponse;
  editCommunity(
    requestMessage: proto_zion_v1_communities_pb.EditCommunityRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_communities_pb.EditCommunityResponse|null) => void
  ): UnaryResponse;
  joinCommunity(
    requestMessage: proto_zion_v1_communities_pb.JoinCommunityRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_communities_pb.JoinCommunityResponse|null) => void
  ): UnaryResponse;
  joinCommunity(
    requestMessage: proto_zion_v1_communities_pb.JoinCommunityRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_communities_pb.JoinCommunityResponse|null) => void
  ): UnaryResponse;
  listCommunity(
    requestMessage: proto_zion_v1_communities_pb.ListCommunityRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_communities_pb.ListCommunityResponse|null) => void
  ): UnaryResponse;
  listCommunity(
    requestMessage: proto_zion_v1_communities_pb.ListCommunityRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_communities_pb.ListCommunityResponse|null) => void
  ): UnaryResponse;
  readCommunity(
    requestMessage: proto_zion_v1_communities_pb.ReadCommunityRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_communities_pb.ReadCommunityResponse|null) => void
  ): UnaryResponse;
  readCommunity(
    requestMessage: proto_zion_v1_communities_pb.ReadCommunityRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_communities_pb.ReadCommunityResponse|null) => void
  ): UnaryResponse;
}

