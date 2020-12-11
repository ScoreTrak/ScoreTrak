/**
 * @fileoverview gRPC-Web generated client stub for pkg.user.userpb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as pkg_user_userpb_user_pb from '../../../pkg/user/userpb/user_pb';


export class UserServiceClient {
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
    pkg_user_userpb_user_pb.GetAllResponse,
    (request: pkg_user_userpb_user_pb.GetAllRequest) => {
      return request.serializeBinary();
    },
    pkg_user_userpb_user_pb.GetAllResponse.deserializeBinary
  );

  getAll(
    request: pkg_user_userpb_user_pb.GetAllRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_user_userpb_user_pb.GetAllResponse>;

  getAll(
    request: pkg_user_userpb_user_pb.GetAllRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_user_userpb_user_pb.GetAllResponse) => void): grpcWeb.ClientReadableStream<pkg_user_userpb_user_pb.GetAllResponse>;

  getAll(
    request: pkg_user_userpb_user_pb.GetAllRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_user_userpb_user_pb.GetAllResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.user.userpb.UserService/GetAll',
        request,
        metadata || {},
        this.methodInfoGetAll,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.user.userpb.UserService/GetAll',
    request,
    metadata || {},
    this.methodInfoGetAll);
  }

  methodInfoGetByID = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_user_userpb_user_pb.GetByIDResponse,
    (request: pkg_user_userpb_user_pb.GetByIDRequest) => {
      return request.serializeBinary();
    },
    pkg_user_userpb_user_pb.GetByIDResponse.deserializeBinary
  );

  getByID(
    request: pkg_user_userpb_user_pb.GetByIDRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_user_userpb_user_pb.GetByIDResponse>;

  getByID(
    request: pkg_user_userpb_user_pb.GetByIDRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_user_userpb_user_pb.GetByIDResponse) => void): grpcWeb.ClientReadableStream<pkg_user_userpb_user_pb.GetByIDResponse>;

  getByID(
    request: pkg_user_userpb_user_pb.GetByIDRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_user_userpb_user_pb.GetByIDResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.user.userpb.UserService/GetByID',
        request,
        metadata || {},
        this.methodInfoGetByID,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.user.userpb.UserService/GetByID',
    request,
    metadata || {},
    this.methodInfoGetByID);
  }

  methodInfoDelete = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_user_userpb_user_pb.DeleteResponse,
    (request: pkg_user_userpb_user_pb.DeleteRequest) => {
      return request.serializeBinary();
    },
    pkg_user_userpb_user_pb.DeleteResponse.deserializeBinary
  );

  delete(
    request: pkg_user_userpb_user_pb.DeleteRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_user_userpb_user_pb.DeleteResponse>;

  delete(
    request: pkg_user_userpb_user_pb.DeleteRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_user_userpb_user_pb.DeleteResponse) => void): grpcWeb.ClientReadableStream<pkg_user_userpb_user_pb.DeleteResponse>;

  delete(
    request: pkg_user_userpb_user_pb.DeleteRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_user_userpb_user_pb.DeleteResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.user.userpb.UserService/Delete',
        request,
        metadata || {},
        this.methodInfoDelete,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.user.userpb.UserService/Delete',
    request,
    metadata || {},
    this.methodInfoDelete);
  }

  methodInfoStore = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_user_userpb_user_pb.StoreResponse,
    (request: pkg_user_userpb_user_pb.StoreRequest) => {
      return request.serializeBinary();
    },
    pkg_user_userpb_user_pb.StoreResponse.deserializeBinary
  );

  store(
    request: pkg_user_userpb_user_pb.StoreRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_user_userpb_user_pb.StoreResponse>;

  store(
    request: pkg_user_userpb_user_pb.StoreRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_user_userpb_user_pb.StoreResponse) => void): grpcWeb.ClientReadableStream<pkg_user_userpb_user_pb.StoreResponse>;

  store(
    request: pkg_user_userpb_user_pb.StoreRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_user_userpb_user_pb.StoreResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.user.userpb.UserService/Store',
        request,
        metadata || {},
        this.methodInfoStore,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.user.userpb.UserService/Store',
    request,
    metadata || {},
    this.methodInfoStore);
  }

  methodInfoUpdate = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_user_userpb_user_pb.UpdateResponse,
    (request: pkg_user_userpb_user_pb.UpdateRequest) => {
      return request.serializeBinary();
    },
    pkg_user_userpb_user_pb.UpdateResponse.deserializeBinary
  );

  update(
    request: pkg_user_userpb_user_pb.UpdateRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_user_userpb_user_pb.UpdateResponse>;

  update(
    request: pkg_user_userpb_user_pb.UpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_user_userpb_user_pb.UpdateResponse) => void): grpcWeb.ClientReadableStream<pkg_user_userpb_user_pb.UpdateResponse>;

  update(
    request: pkg_user_userpb_user_pb.UpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_user_userpb_user_pb.UpdateResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.user.userpb.UserService/Update',
        request,
        metadata || {},
        this.methodInfoUpdate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.user.userpb.UserService/Update',
    request,
    metadata || {},
    this.methodInfoUpdate);
  }

  methodInfoGetByUsername = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_user_userpb_user_pb.GetByUsernameResponse,
    (request: pkg_user_userpb_user_pb.GetByUsernameRequest) => {
      return request.serializeBinary();
    },
    pkg_user_userpb_user_pb.GetByUsernameResponse.deserializeBinary
  );

  getByUsername(
    request: pkg_user_userpb_user_pb.GetByUsernameRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_user_userpb_user_pb.GetByUsernameResponse>;

  getByUsername(
    request: pkg_user_userpb_user_pb.GetByUsernameRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_user_userpb_user_pb.GetByUsernameResponse) => void): grpcWeb.ClientReadableStream<pkg_user_userpb_user_pb.GetByUsernameResponse>;

  getByUsername(
    request: pkg_user_userpb_user_pb.GetByUsernameRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_user_userpb_user_pb.GetByUsernameResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.user.userpb.UserService/GetByUsername',
        request,
        metadata || {},
        this.methodInfoGetByUsername,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.user.userpb.UserService/GetByUsername',
    request,
    metadata || {},
    this.methodInfoGetByUsername);
  }

}

