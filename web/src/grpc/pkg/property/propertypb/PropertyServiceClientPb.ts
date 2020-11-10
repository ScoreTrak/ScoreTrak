/**
 * @fileoverview gRPC-Web generated client stub for pkg.property.propertypb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as pkg_property_propertypb_property_pb from '../../../pkg/property/propertypb/property_pb';


export class PropertyServiceClient {
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
    pkg_property_propertypb_property_pb.GetAllResponse,
    (request: pkg_property_propertypb_property_pb.GetAllRequest) => {
      return request.serializeBinary();
    },
    pkg_property_propertypb_property_pb.GetAllResponse.deserializeBinary
  );

  getAll(
    request: pkg_property_propertypb_property_pb.GetAllRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_property_propertypb_property_pb.GetAllResponse>;

  getAll(
    request: pkg_property_propertypb_property_pb.GetAllRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_property_propertypb_property_pb.GetAllResponse) => void): grpcWeb.ClientReadableStream<pkg_property_propertypb_property_pb.GetAllResponse>;

  getAll(
    request: pkg_property_propertypb_property_pb.GetAllRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_property_propertypb_property_pb.GetAllResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.property.propertypb.PropertyService/GetAll',
        request,
        metadata || {},
        this.methodInfoGetAll,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.property.propertypb.PropertyService/GetAll',
    request,
    metadata || {},
    this.methodInfoGetAll);
  }

  methodInfoDelete = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_property_propertypb_property_pb.DeleteResponse,
    (request: pkg_property_propertypb_property_pb.DeleteRequest) => {
      return request.serializeBinary();
    },
    pkg_property_propertypb_property_pb.DeleteResponse.deserializeBinary
  );

  delete(
    request: pkg_property_propertypb_property_pb.DeleteRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_property_propertypb_property_pb.DeleteResponse>;

  delete(
    request: pkg_property_propertypb_property_pb.DeleteRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_property_propertypb_property_pb.DeleteResponse) => void): grpcWeb.ClientReadableStream<pkg_property_propertypb_property_pb.DeleteResponse>;

  delete(
    request: pkg_property_propertypb_property_pb.DeleteRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_property_propertypb_property_pb.DeleteResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.property.propertypb.PropertyService/Delete',
        request,
        metadata || {},
        this.methodInfoDelete,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.property.propertypb.PropertyService/Delete',
    request,
    metadata || {},
    this.methodInfoDelete);
  }

  methodInfoStore = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_property_propertypb_property_pb.StoreResponse,
    (request: pkg_property_propertypb_property_pb.StoreRequest) => {
      return request.serializeBinary();
    },
    pkg_property_propertypb_property_pb.StoreResponse.deserializeBinary
  );

  store(
    request: pkg_property_propertypb_property_pb.StoreRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_property_propertypb_property_pb.StoreResponse>;

  store(
    request: pkg_property_propertypb_property_pb.StoreRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_property_propertypb_property_pb.StoreResponse) => void): grpcWeb.ClientReadableStream<pkg_property_propertypb_property_pb.StoreResponse>;

  store(
    request: pkg_property_propertypb_property_pb.StoreRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_property_propertypb_property_pb.StoreResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.property.propertypb.PropertyService/Store',
        request,
        metadata || {},
        this.methodInfoStore,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.property.propertypb.PropertyService/Store',
    request,
    metadata || {},
    this.methodInfoStore);
  }

  methodInfoUpdate = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_property_propertypb_property_pb.UpdateResponse,
    (request: pkg_property_propertypb_property_pb.UpdateRequest) => {
      return request.serializeBinary();
    },
    pkg_property_propertypb_property_pb.UpdateResponse.deserializeBinary
  );

  update(
    request: pkg_property_propertypb_property_pb.UpdateRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_property_propertypb_property_pb.UpdateResponse>;

  update(
    request: pkg_property_propertypb_property_pb.UpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_property_propertypb_property_pb.UpdateResponse) => void): grpcWeb.ClientReadableStream<pkg_property_propertypb_property_pb.UpdateResponse>;

  update(
    request: pkg_property_propertypb_property_pb.UpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_property_propertypb_property_pb.UpdateResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.property.propertypb.PropertyService/Update',
        request,
        metadata || {},
        this.methodInfoUpdate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.property.propertypb.PropertyService/Update',
    request,
    metadata || {},
    this.methodInfoUpdate);
  }

  methodInfoGetByServiceIDKey = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_property_propertypb_property_pb.GetByServiceIDKeyResponse,
    (request: pkg_property_propertypb_property_pb.GetByServiceIDKeyRequest) => {
      return request.serializeBinary();
    },
    pkg_property_propertypb_property_pb.GetByServiceIDKeyResponse.deserializeBinary
  );

  getByServiceIDKey(
    request: pkg_property_propertypb_property_pb.GetByServiceIDKeyRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_property_propertypb_property_pb.GetByServiceIDKeyResponse>;

  getByServiceIDKey(
    request: pkg_property_propertypb_property_pb.GetByServiceIDKeyRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_property_propertypb_property_pb.GetByServiceIDKeyResponse) => void): grpcWeb.ClientReadableStream<pkg_property_propertypb_property_pb.GetByServiceIDKeyResponse>;

  getByServiceIDKey(
    request: pkg_property_propertypb_property_pb.GetByServiceIDKeyRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_property_propertypb_property_pb.GetByServiceIDKeyResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.property.propertypb.PropertyService/GetByServiceIDKey',
        request,
        metadata || {},
        this.methodInfoGetByServiceIDKey,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.property.propertypb.PropertyService/GetByServiceIDKey',
    request,
    metadata || {},
    this.methodInfoGetByServiceIDKey);
  }

  methodInfoGetAllByServiceID = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_property_propertypb_property_pb.GetAllByServiceIDResponse,
    (request: pkg_property_propertypb_property_pb.GetAllByServiceIDRequest) => {
      return request.serializeBinary();
    },
    pkg_property_propertypb_property_pb.GetAllByServiceIDResponse.deserializeBinary
  );

  getAllByServiceID(
    request: pkg_property_propertypb_property_pb.GetAllByServiceIDRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_property_propertypb_property_pb.GetAllByServiceIDResponse>;

  getAllByServiceID(
    request: pkg_property_propertypb_property_pb.GetAllByServiceIDRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_property_propertypb_property_pb.GetAllByServiceIDResponse) => void): grpcWeb.ClientReadableStream<pkg_property_propertypb_property_pb.GetAllByServiceIDResponse>;

  getAllByServiceID(
    request: pkg_property_propertypb_property_pb.GetAllByServiceIDRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_property_propertypb_property_pb.GetAllByServiceIDResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.property.propertypb.PropertyService/GetAllByServiceID',
        request,
        metadata || {},
        this.methodInfoGetAllByServiceID,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.property.propertypb.PropertyService/GetAllByServiceID',
    request,
    metadata || {},
    this.methodInfoGetAllByServiceID);
  }

}

