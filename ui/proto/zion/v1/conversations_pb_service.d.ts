// package: proto.zion.v1
// file: proto/zion/v1/conversations.proto

import * as proto_zion_v1_conversations_pb from "../../../proto/zion/v1/conversations_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ConversationServiceCreateConversation = {
  readonly methodName: string;
  readonly service: typeof ConversationService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_conversations_pb.CreateConversationRequest;
  readonly responseType: typeof proto_zion_v1_conversations_pb.CreateConversationResponse;
};

type ConversationServiceDeleteComment = {
  readonly methodName: string;
  readonly service: typeof ConversationService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_conversations_pb.DeleteCommentRequest;
  readonly responseType: typeof proto_zion_v1_conversations_pb.DeleteCommentResponse;
};

type ConversationServiceDeleteConversation = {
  readonly methodName: string;
  readonly service: typeof ConversationService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_conversations_pb.DeleteConversationRequest;
  readonly responseType: typeof proto_zion_v1_conversations_pb.DeleteConversationResponse;
};

type ConversationServiceGetActivity = {
  readonly methodName: string;
  readonly service: typeof ConversationService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_conversations_pb.GetActivityRequest;
  readonly responseType: typeof proto_zion_v1_conversations_pb.GetActivityResponse;
};

type ConversationServiceGetConversation = {
  readonly methodName: string;
  readonly service: typeof ConversationService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_conversations_pb.GetConversationRequest;
  readonly responseType: typeof proto_zion_v1_conversations_pb.GetConversationResponse;
};

type ConversationServiceGetFeed = {
  readonly methodName: string;
  readonly service: typeof ConversationService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_conversations_pb.GetFeedRequest;
  readonly responseType: typeof proto_zion_v1_conversations_pb.GetFeedResponse;
};

export class ConversationService {
  static readonly serviceName: string;
  static readonly CreateConversation: ConversationServiceCreateConversation;
  static readonly DeleteComment: ConversationServiceDeleteComment;
  static readonly DeleteConversation: ConversationServiceDeleteConversation;
  static readonly GetActivity: ConversationServiceGetActivity;
  static readonly GetConversation: ConversationServiceGetConversation;
  static readonly GetFeed: ConversationServiceGetFeed;
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

export class ConversationServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  createConversation(
    requestMessage: proto_zion_v1_conversations_pb.CreateConversationRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_conversations_pb.CreateConversationResponse|null) => void
  ): UnaryResponse;
  createConversation(
    requestMessage: proto_zion_v1_conversations_pb.CreateConversationRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_conversations_pb.CreateConversationResponse|null) => void
  ): UnaryResponse;
  deleteComment(
    requestMessage: proto_zion_v1_conversations_pb.DeleteCommentRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_conversations_pb.DeleteCommentResponse|null) => void
  ): UnaryResponse;
  deleteComment(
    requestMessage: proto_zion_v1_conversations_pb.DeleteCommentRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_conversations_pb.DeleteCommentResponse|null) => void
  ): UnaryResponse;
  deleteConversation(
    requestMessage: proto_zion_v1_conversations_pb.DeleteConversationRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_conversations_pb.DeleteConversationResponse|null) => void
  ): UnaryResponse;
  deleteConversation(
    requestMessage: proto_zion_v1_conversations_pb.DeleteConversationRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_conversations_pb.DeleteConversationResponse|null) => void
  ): UnaryResponse;
  getActivity(
    requestMessage: proto_zion_v1_conversations_pb.GetActivityRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_conversations_pb.GetActivityResponse|null) => void
  ): UnaryResponse;
  getActivity(
    requestMessage: proto_zion_v1_conversations_pb.GetActivityRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_conversations_pb.GetActivityResponse|null) => void
  ): UnaryResponse;
  getConversation(
    requestMessage: proto_zion_v1_conversations_pb.GetConversationRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_conversations_pb.GetConversationResponse|null) => void
  ): UnaryResponse;
  getConversation(
    requestMessage: proto_zion_v1_conversations_pb.GetConversationRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_conversations_pb.GetConversationResponse|null) => void
  ): UnaryResponse;
  getFeed(
    requestMessage: proto_zion_v1_conversations_pb.GetFeedRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_conversations_pb.GetFeedResponse|null) => void
  ): UnaryResponse;
  getFeed(
    requestMessage: proto_zion_v1_conversations_pb.GetFeedRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_conversations_pb.GetFeedResponse|null) => void
  ): UnaryResponse;
}

