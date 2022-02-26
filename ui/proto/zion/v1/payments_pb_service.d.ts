// package: proto.zion.v1
// file: proto/zion/v1/payments.proto

import * as proto_zion_v1_payments_pb from "../../../proto/zion/v1/payments_pb";
import {grpc} from "@improbable-eng/grpc-web";

type PaymentsServiceCreateInvoice = {
  readonly methodName: string;
  readonly service: typeof PaymentsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_payments_pb.CreateInvoiceRequest;
  readonly responseType: typeof proto_zion_v1_payments_pb.CreateInvoiceResponse;
};

type PaymentsServiceGetPaymentsHistory = {
  readonly methodName: string;
  readonly service: typeof PaymentsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_payments_pb.GetPaymentsHistoryRequest;
  readonly responseType: typeof proto_zion_v1_payments_pb.GetPaymentsHistoryResponse;
};

type PaymentsServicePayInvoice = {
  readonly methodName: string;
  readonly service: typeof PaymentsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_payments_pb.PayInvoiceRequest;
  readonly responseType: typeof proto_zion_v1_payments_pb.PayInvoiceResponse;
};

type PaymentsServicePayUser = {
  readonly methodName: string;
  readonly service: typeof PaymentsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_zion_v1_payments_pb.PayUserRequest;
  readonly responseType: typeof proto_zion_v1_payments_pb.PayUserResponse;
};

export class PaymentsService {
  static readonly serviceName: string;
  static readonly CreateInvoice: PaymentsServiceCreateInvoice;
  static readonly GetPaymentsHistory: PaymentsServiceGetPaymentsHistory;
  static readonly PayInvoice: PaymentsServicePayInvoice;
  static readonly PayUser: PaymentsServicePayUser;
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

export class PaymentsServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  createInvoice(
    requestMessage: proto_zion_v1_payments_pb.CreateInvoiceRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_payments_pb.CreateInvoiceResponse|null) => void
  ): UnaryResponse;
  createInvoice(
    requestMessage: proto_zion_v1_payments_pb.CreateInvoiceRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_payments_pb.CreateInvoiceResponse|null) => void
  ): UnaryResponse;
  getPaymentsHistory(
    requestMessage: proto_zion_v1_payments_pb.GetPaymentsHistoryRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_payments_pb.GetPaymentsHistoryResponse|null) => void
  ): UnaryResponse;
  getPaymentsHistory(
    requestMessage: proto_zion_v1_payments_pb.GetPaymentsHistoryRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_payments_pb.GetPaymentsHistoryResponse|null) => void
  ): UnaryResponse;
  payInvoice(
    requestMessage: proto_zion_v1_payments_pb.PayInvoiceRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_payments_pb.PayInvoiceResponse|null) => void
  ): UnaryResponse;
  payInvoice(
    requestMessage: proto_zion_v1_payments_pb.PayInvoiceRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_payments_pb.PayInvoiceResponse|null) => void
  ): UnaryResponse;
  payUser(
    requestMessage: proto_zion_v1_payments_pb.PayUserRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_payments_pb.PayUserResponse|null) => void
  ): UnaryResponse;
  payUser(
    requestMessage: proto_zion_v1_payments_pb.PayUserRequest,
    callback: (error: ServiceError|null, responseMessage: proto_zion_v1_payments_pb.PayUserResponse|null) => void
  ): UnaryResponse;
}

