// package: proto.identityhub.v1
// file: proto/identityhub/v1/identityhub.proto

import * as proto_identityhub_v1_identityhub_pb from "../../../proto/identityhub/v1/identityhub_pb";
import * as proto_identityhub_v1_requests_pb from "../../../proto/identityhub/v1/requests_pb";
import * as proto_identityhub_v1_responses_pb from "../../../proto/identityhub/v1/responses_pb";
import {grpc} from "@improbable-eng/grpc-web";

type HubRequestServiceProcess = {
  readonly methodName: string;
  readonly service: typeof HubRequestService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_identityhub_v1_requests_pb.Request;
  readonly responseType: typeof proto_identityhub_v1_responses_pb.Response;
};

export class HubRequestService {
  static readonly serviceName: string;
  static readonly Process: HubRequestServiceProcess;
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

export class HubRequestServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  process(
    requestMessage: proto_identityhub_v1_requests_pb.Request,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_responses_pb.Response|null) => void
  ): UnaryResponse;
  process(
    requestMessage: proto_identityhub_v1_requests_pb.Request,
    callback: (error: ServiceError|null, responseMessage: proto_identityhub_v1_responses_pb.Response|null) => void
  ): UnaryResponse;
}

