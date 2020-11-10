/**
 * @fileoverview gRPC-Web generated client stub for pkg.team.teampb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as pkg_team_teampb_team_pb from '../../../pkg/team/teampb/team_pb';


export class TeamServiceClient {
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

  methodInfoGetAll = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_team_teampb_team_pb.GetAllResponse,
    (request: pkg_team_teampb_team_pb.GetAllRequest) => {
      return request.serializeBinary();
    },
    pkg_team_teampb_team_pb.GetAllResponse.deserializeBinary
  );

  getAll(
    request: pkg_team_teampb_team_pb.GetAllRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_team_teampb_team_pb.GetAllResponse>;

  getAll(
    request: pkg_team_teampb_team_pb.GetAllRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_team_teampb_team_pb.GetAllResponse) => void): grpcWeb.ClientReadableStream<pkg_team_teampb_team_pb.GetAllResponse>;

  getAll(
    request: pkg_team_teampb_team_pb.GetAllRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_team_teampb_team_pb.GetAllResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.team.teampb.TeamService/GetAll',
        request,
        metadata || {},
        this.methodInfoGetAll,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.team.teampb.TeamService/GetAll',
    request,
    metadata || {},
    this.methodInfoGetAll);
  }

  methodInfoGetByID = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_team_teampb_team_pb.GetByIDResponse,
    (request: pkg_team_teampb_team_pb.GetByIDRequest) => {
      return request.serializeBinary();
    },
    pkg_team_teampb_team_pb.GetByIDResponse.deserializeBinary
  );

  getByID(
    request: pkg_team_teampb_team_pb.GetByIDRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_team_teampb_team_pb.GetByIDResponse>;

  getByID(
    request: pkg_team_teampb_team_pb.GetByIDRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_team_teampb_team_pb.GetByIDResponse) => void): grpcWeb.ClientReadableStream<pkg_team_teampb_team_pb.GetByIDResponse>;

  getByID(
    request: pkg_team_teampb_team_pb.GetByIDRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_team_teampb_team_pb.GetByIDResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.team.teampb.TeamService/GetByID',
        request,
        metadata || {},
        this.methodInfoGetByID,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.team.teampb.TeamService/GetByID',
    request,
    metadata || {},
    this.methodInfoGetByID);
  }

  methodInfoDelete = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_team_teampb_team_pb.DeleteResponse,
    (request: pkg_team_teampb_team_pb.DeleteRequest) => {
      return request.serializeBinary();
    },
    pkg_team_teampb_team_pb.DeleteResponse.deserializeBinary
  );

  delete(
    request: pkg_team_teampb_team_pb.DeleteRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_team_teampb_team_pb.DeleteResponse>;

  delete(
    request: pkg_team_teampb_team_pb.DeleteRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_team_teampb_team_pb.DeleteResponse) => void): grpcWeb.ClientReadableStream<pkg_team_teampb_team_pb.DeleteResponse>;

  delete(
    request: pkg_team_teampb_team_pb.DeleteRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_team_teampb_team_pb.DeleteResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.team.teampb.TeamService/Delete',
        request,
        metadata || {},
        this.methodInfoDelete,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.team.teampb.TeamService/Delete',
    request,
    metadata || {},
    this.methodInfoDelete);
  }

  methodInfoStore = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_team_teampb_team_pb.StoreResponse,
    (request: pkg_team_teampb_team_pb.StoreRequest) => {
      return request.serializeBinary();
    },
    pkg_team_teampb_team_pb.StoreResponse.deserializeBinary
  );

  store(
    request: pkg_team_teampb_team_pb.StoreRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_team_teampb_team_pb.StoreResponse>;

  store(
    request: pkg_team_teampb_team_pb.StoreRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_team_teampb_team_pb.StoreResponse) => void): grpcWeb.ClientReadableStream<pkg_team_teampb_team_pb.StoreResponse>;

  store(
    request: pkg_team_teampb_team_pb.StoreRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_team_teampb_team_pb.StoreResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.team.teampb.TeamService/Store',
        request,
        metadata || {},
        this.methodInfoStore,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.team.teampb.TeamService/Store',
    request,
    metadata || {},
    this.methodInfoStore);
  }

  methodInfoUpdate = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_team_teampb_team_pb.UpdateResponse,
    (request: pkg_team_teampb_team_pb.UpdateRequest) => {
      return request.serializeBinary();
    },
    pkg_team_teampb_team_pb.UpdateResponse.deserializeBinary
  );

  update(
    request: pkg_team_teampb_team_pb.UpdateRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_team_teampb_team_pb.UpdateResponse>;

  update(
    request: pkg_team_teampb_team_pb.UpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_team_teampb_team_pb.UpdateResponse) => void): grpcWeb.ClientReadableStream<pkg_team_teampb_team_pb.UpdateResponse>;

  update(
    request: pkg_team_teampb_team_pb.UpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_team_teampb_team_pb.UpdateResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.team.teampb.TeamService/Update',
        request,
        metadata || {},
        this.methodInfoUpdate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.team.teampb.TeamService/Update',
    request,
    metadata || {},
    this.methodInfoUpdate);
  }

}

