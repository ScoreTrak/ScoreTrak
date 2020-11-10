/**
 * @fileoverview gRPC-Web generated client stub for pkg.competition.competitionpb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as pkg_competition_competitionpb_competition_pb from '../../../pkg/competition/competitionpb/competition_pb';


export class CompetitionServiceClient {
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

  methodInfoLoadCompetition = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_competition_competitionpb_competition_pb.LoadCompetitionResponse,
    (request: pkg_competition_competitionpb_competition_pb.LoadCompetitionRequest) => {
      return request.serializeBinary();
    },
    pkg_competition_competitionpb_competition_pb.LoadCompetitionResponse.deserializeBinary
  );

  loadCompetition(
    request: pkg_competition_competitionpb_competition_pb.LoadCompetitionRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_competition_competitionpb_competition_pb.LoadCompetitionResponse>;

  loadCompetition(
    request: pkg_competition_competitionpb_competition_pb.LoadCompetitionRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_competition_competitionpb_competition_pb.LoadCompetitionResponse) => void): grpcWeb.ClientReadableStream<pkg_competition_competitionpb_competition_pb.LoadCompetitionResponse>;

  loadCompetition(
    request: pkg_competition_competitionpb_competition_pb.LoadCompetitionRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_competition_competitionpb_competition_pb.LoadCompetitionResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.competition.competitionpb.CompetitionService/LoadCompetition',
        request,
        metadata || {},
        this.methodInfoLoadCompetition,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.competition.competitionpb.CompetitionService/LoadCompetition',
    request,
    metadata || {},
    this.methodInfoLoadCompetition);
  }

  methodInfoFetchCoreCompetition = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_competition_competitionpb_competition_pb.FetchCoreCompetitionResponse,
    (request: pkg_competition_competitionpb_competition_pb.FetchCoreCompetitionRequest) => {
      return request.serializeBinary();
    },
    pkg_competition_competitionpb_competition_pb.FetchCoreCompetitionResponse.deserializeBinary
  );

  fetchCoreCompetition(
    request: pkg_competition_competitionpb_competition_pb.FetchCoreCompetitionRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_competition_competitionpb_competition_pb.FetchCoreCompetitionResponse>;

  fetchCoreCompetition(
    request: pkg_competition_competitionpb_competition_pb.FetchCoreCompetitionRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_competition_competitionpb_competition_pb.FetchCoreCompetitionResponse) => void): grpcWeb.ClientReadableStream<pkg_competition_competitionpb_competition_pb.FetchCoreCompetitionResponse>;

  fetchCoreCompetition(
    request: pkg_competition_competitionpb_competition_pb.FetchCoreCompetitionRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_competition_competitionpb_competition_pb.FetchCoreCompetitionResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.competition.competitionpb.CompetitionService/FetchCoreCompetition',
        request,
        metadata || {},
        this.methodInfoFetchCoreCompetition,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.competition.competitionpb.CompetitionService/FetchCoreCompetition',
    request,
    metadata || {},
    this.methodInfoFetchCoreCompetition);
  }

  methodInfoFetchEntireCompetition = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_competition_competitionpb_competition_pb.FetchEntireCompetitionResponse,
    (request: pkg_competition_competitionpb_competition_pb.FetchEntireCompetitionRequest) => {
      return request.serializeBinary();
    },
    pkg_competition_competitionpb_competition_pb.FetchEntireCompetitionResponse.deserializeBinary
  );

  fetchEntireCompetition(
    request: pkg_competition_competitionpb_competition_pb.FetchEntireCompetitionRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_competition_competitionpb_competition_pb.FetchEntireCompetitionResponse>;

  fetchEntireCompetition(
    request: pkg_competition_competitionpb_competition_pb.FetchEntireCompetitionRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_competition_competitionpb_competition_pb.FetchEntireCompetitionResponse) => void): grpcWeb.ClientReadableStream<pkg_competition_competitionpb_competition_pb.FetchEntireCompetitionResponse>;

  fetchEntireCompetition(
    request: pkg_competition_competitionpb_competition_pb.FetchEntireCompetitionRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_competition_competitionpb_competition_pb.FetchEntireCompetitionResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.competition.competitionpb.CompetitionService/FetchEntireCompetition',
        request,
        metadata || {},
        this.methodInfoFetchEntireCompetition,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.competition.competitionpb.CompetitionService/FetchEntireCompetition',
    request,
    metadata || {},
    this.methodInfoFetchEntireCompetition);
  }

  methodInfoResetScores = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_competition_competitionpb_competition_pb.ResetScoresResponse,
    (request: pkg_competition_competitionpb_competition_pb.ResetScoresRequest) => {
      return request.serializeBinary();
    },
    pkg_competition_competitionpb_competition_pb.ResetScoresResponse.deserializeBinary
  );

  resetScores(
    request: pkg_competition_competitionpb_competition_pb.ResetScoresRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_competition_competitionpb_competition_pb.ResetScoresResponse>;

  resetScores(
    request: pkg_competition_competitionpb_competition_pb.ResetScoresRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_competition_competitionpb_competition_pb.ResetScoresResponse) => void): grpcWeb.ClientReadableStream<pkg_competition_competitionpb_competition_pb.ResetScoresResponse>;

  resetScores(
    request: pkg_competition_competitionpb_competition_pb.ResetScoresRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_competition_competitionpb_competition_pb.ResetScoresResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.competition.competitionpb.CompetitionService/ResetScores',
        request,
        metadata || {},
        this.methodInfoResetScores,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.competition.competitionpb.CompetitionService/ResetScores',
    request,
    metadata || {},
    this.methodInfoResetScores);
  }

  methodInfoDeleteCompetition = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_competition_competitionpb_competition_pb.DeleteCompetitionResponse,
    (request: pkg_competition_competitionpb_competition_pb.DeleteCompetitionRequest) => {
      return request.serializeBinary();
    },
    pkg_competition_competitionpb_competition_pb.DeleteCompetitionResponse.deserializeBinary
  );

  deleteCompetition(
    request: pkg_competition_competitionpb_competition_pb.DeleteCompetitionRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_competition_competitionpb_competition_pb.DeleteCompetitionResponse>;

  deleteCompetition(
    request: pkg_competition_competitionpb_competition_pb.DeleteCompetitionRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_competition_competitionpb_competition_pb.DeleteCompetitionResponse) => void): grpcWeb.ClientReadableStream<pkg_competition_competitionpb_competition_pb.DeleteCompetitionResponse>;

  deleteCompetition(
    request: pkg_competition_competitionpb_competition_pb.DeleteCompetitionRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_competition_competitionpb_competition_pb.DeleteCompetitionResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.competition.competitionpb.CompetitionService/DeleteCompetition',
        request,
        metadata || {},
        this.methodInfoDeleteCompetition,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.competition.competitionpb.CompetitionService/DeleteCompetition',
    request,
    metadata || {},
    this.methodInfoDeleteCompetition);
  }

}

