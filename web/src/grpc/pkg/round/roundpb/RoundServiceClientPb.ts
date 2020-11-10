/**
 * @fileoverview gRPC-Web generated client stub for pkg.round.roundpb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as pkg_round_roundpb_round_pb from '../../../pkg/round/roundpb/round_pb';


export class RoundServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoGetLastNonElapsingRound = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_round_roundpb_round_pb.GetLastNonElapsingRoundResponse,
    (request: pkg_round_roundpb_round_pb.GetLastNonElapsingRoundRequest) => {
      return request.serializeBinary();
    },
    pkg_round_roundpb_round_pb.GetLastNonElapsingRoundResponse.deserializeBinary
  );

  getLastNonElapsingRound(
    request: pkg_round_roundpb_round_pb.GetLastNonElapsingRoundRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_round_roundpb_round_pb.GetLastNonElapsingRoundResponse>;

  getLastNonElapsingRound(
    request: pkg_round_roundpb_round_pb.GetLastNonElapsingRoundRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_round_roundpb_round_pb.GetLastNonElapsingRoundResponse) => void): grpcWeb.ClientReadableStream<pkg_round_roundpb_round_pb.GetLastNonElapsingRoundResponse>;

  getLastNonElapsingRound(
    request: pkg_round_roundpb_round_pb.GetLastNonElapsingRoundRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_round_roundpb_round_pb.GetLastNonElapsingRoundResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.round.roundpb.RoundService/GetLastNonElapsingRound',
        request,
        metadata || {},
        this.methodInfoGetLastNonElapsingRound,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.round.roundpb.RoundService/GetLastNonElapsingRound',
    request,
    metadata || {},
    this.methodInfoGetLastNonElapsingRound);
  }

  methodInfoGetAll = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_round_roundpb_round_pb.GetAllResponse,
    (request: pkg_round_roundpb_round_pb.GetAllRequest) => {
      return request.serializeBinary();
    },
    pkg_round_roundpb_round_pb.GetAllResponse.deserializeBinary
  );

  getAll(
    request: pkg_round_roundpb_round_pb.GetAllRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_round_roundpb_round_pb.GetAllResponse>;

  getAll(
    request: pkg_round_roundpb_round_pb.GetAllRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_round_roundpb_round_pb.GetAllResponse) => void): grpcWeb.ClientReadableStream<pkg_round_roundpb_round_pb.GetAllResponse>;

  getAll(
    request: pkg_round_roundpb_round_pb.GetAllRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_round_roundpb_round_pb.GetAllResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.round.roundpb.RoundService/GetAll',
        request,
        metadata || {},
        this.methodInfoGetAll,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.round.roundpb.RoundService/GetAll',
    request,
    metadata || {},
    this.methodInfoGetAll);
  }

  methodInfoGetByID = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_round_roundpb_round_pb.GetByIDResponse,
    (request: pkg_round_roundpb_round_pb.GetByIDRequest) => {
      return request.serializeBinary();
    },
    pkg_round_roundpb_round_pb.GetByIDResponse.deserializeBinary
  );

  getByID(
    request: pkg_round_roundpb_round_pb.GetByIDRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_round_roundpb_round_pb.GetByIDResponse>;

  getByID(
    request: pkg_round_roundpb_round_pb.GetByIDRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_round_roundpb_round_pb.GetByIDResponse) => void): grpcWeb.ClientReadableStream<pkg_round_roundpb_round_pb.GetByIDResponse>;

  getByID(
    request: pkg_round_roundpb_round_pb.GetByIDRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_round_roundpb_round_pb.GetByIDResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.round.roundpb.RoundService/GetByID',
        request,
        metadata || {},
        this.methodInfoGetByID,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.round.roundpb.RoundService/GetByID',
    request,
    metadata || {},
    this.methodInfoGetByID);
  }

  methodInfoGetLastRound = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_round_roundpb_round_pb.GetLastRoundResponse,
    (request: pkg_round_roundpb_round_pb.GetLastRoundRequest) => {
      return request.serializeBinary();
    },
    pkg_round_roundpb_round_pb.GetLastRoundResponse.deserializeBinary
  );

  getLastRound(
    request: pkg_round_roundpb_round_pb.GetLastRoundRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_round_roundpb_round_pb.GetLastRoundResponse>;

  getLastRound(
    request: pkg_round_roundpb_round_pb.GetLastRoundRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_round_roundpb_round_pb.GetLastRoundResponse) => void): grpcWeb.ClientReadableStream<pkg_round_roundpb_round_pb.GetLastRoundResponse>;

  getLastRound(
    request: pkg_round_roundpb_round_pb.GetLastRoundRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_round_roundpb_round_pb.GetLastRoundResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.round.roundpb.RoundService/GetLastRound',
        request,
        metadata || {},
        this.methodInfoGetLastRound,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.round.roundpb.RoundService/GetLastRound',
    request,
    metadata || {},
    this.methodInfoGetLastRound);
  }

}

