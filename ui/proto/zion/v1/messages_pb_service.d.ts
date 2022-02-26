// package: proto.zion.v1
// file: proto/zion/v1/messages.proto

import * as proto_zion_v1_messages_pb from "../../../proto/zion/v1/messages_pb";
import {grpc} from "@improbable-eng/grpc-web";

type MessagesServiceBoostMessage = {
  readonly methodName: string;
  readonly service: typeof MessagesService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_messages_pb.BoostMessageRequest;
  readonly responseType: typeof proto_zion_v1_messages_pb.BoostMessageResponse;
};

type MessagesServiceDeleteMessage = {
  readonly methodName: string;
  readonly service: typeof MessagesService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_messages_pb.DeleteMessageRequest;
  readonly responseType: typeof proto_zion_v1_messages_pb.DeleteMessageResponse;
};

type MessagesServiceGetMessages = {
  readonly methodName: string;
  readonly service: typeof MessagesService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_messages_pb.GetMessagesRequest;
  readonly responseType: typeof proto_zion_v1_messages_pb.GetMessagesResponse;
};

type MessagesServiceSendMessage = {
  readonly methodName: string;
  readonly service: typeof MessagesService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_messages_pb.SendMessageRequest;
  readonly responseType: typeof proto_zion_v1_messages_pb.SendMessageResponse;
};

export class MessagesService {
  static readonly serviceName: string;
  static readonly BoostMessage: MessagesServiceBoostMessage;
  static readonly DeleteMessage: MessagesServiceDeleteMessage;
  static readonly GetMessages: MessagesServiceGetMessages;
  static readonly SendMessage: MessagesServiceSendMessage;
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

export class MessagesServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  boostMessage(
    requestMessage: proto_zion_v1_messages_pb.BoostMessageRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_messages_pb.BoostMessageResponse|null) => void
  ): UnaryResponse;
  boostMessage(
    requestMessage: proto_zion_v1_messages_pb.BoostMessageRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_messages_pb.BoostMessageResponse|null) => void
  ): UnaryResponse;
  deleteMessage(
    requestMessage: proto_zion_v1_messages_pb.DeleteMessageRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_messages_pb.DeleteMessageResponse|null) => void
  ): UnaryResponse;
  deleteMessage(
    requestMessage: proto_zion_v1_messages_pb.DeleteMessageRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_messages_pb.DeleteMessageResponse|null) => void
  ): UnaryResponse;
  getMessages(
    requestMessage: proto_zion_v1_messages_pb.GetMessagesRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_messages_pb.GetMessagesResponse|null) => void
  ): UnaryResponse;
  getMessages(
    requestMessage: proto_zion_v1_messages_pb.GetMessagesRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_messages_pb.GetMessagesResponse|null) => void
  ): UnaryResponse;
  sendMessage(
    requestMessage: proto_zion_v1_messages_pb.SendMessageRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_messages_pb.SendMessageResponse|null) => void
  ): UnaryResponse;
  sendMessage(
    requestMessage: proto_zion_v1_messages_pb.SendMessageRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_messages_pb.SendMessageResponse|null) => void
  ): UnaryResponse;
}

