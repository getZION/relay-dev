// package: proto.zion.v1
// file: proto/zion/v1/nodeinfo.proto

import * as proto_zion_v1_nodeinfo_pb from "../../../proto/zion/v1/nodeinfo_pb";
import {grpc} from "@improbable-eng/grpc-web";

type NodeInfoServiceGetNodeInfo = {
  readonly methodName: string;
  readonly service: typeof NodeInfoService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_nodeinfo_pb.GetNodeInfoRequest;
  readonly responseType: typeof proto_zion_v1_nodeinfo_pb.GetNodeInfoResponse;
};

export class NodeInfoService {
  static readonly serviceName: string;
  static readonly GetNodeInfo: NodeInfoServiceGetNodeInfo;
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

export class NodeInfoServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getNodeInfo(
    requestMessage: proto_zion_v1_nodeinfo_pb.GetNodeInfoRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_nodeinfo_pb.GetNodeInfoResponse|null) => void
  ): UnaryResponse;
  getNodeInfo(
    requestMessage: proto_zion_v1_nodeinfo_pb.GetNodeInfoRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_nodeinfo_pb.GetNodeInfoResponse|null) => void
  ): UnaryResponse;
}

