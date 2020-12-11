/**
 * @fileoverview gRPC-Web generated client stub for pkg.config.config
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as pkg_config_configpb_config_pb from '../../../pkg/config/configpb/config_pb';


export class DynamicConfigServiceClient {
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

  methodInfoGet = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_config_configpb_config_pb.GetResponse,
    (request: pkg_config_configpb_config_pb.GetRequest) => {
      return request.serializeBinary();
    },
    pkg_config_configpb_config_pb.GetResponse.deserializeBinary
  );

  get(
    request: pkg_config_configpb_config_pb.GetRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_config_configpb_config_pb.GetResponse>;

  get(
    request: pkg_config_configpb_config_pb.GetRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_config_configpb_config_pb.GetResponse) => void): grpcWeb.ClientReadableStream<pkg_config_configpb_config_pb.GetResponse>;

  get(
    request: pkg_config_configpb_config_pb.GetRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_config_configpb_config_pb.GetResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.config.config.DynamicConfigService/Get',
        request,
        metadata || {},
        this.methodInfoGet,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.config.config.DynamicConfigService/Get',
    request,
    metadata || {},
    this.methodInfoGet);
  }

  methodInfoUpdate = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_config_configpb_config_pb.UpdateResponse,
    (request: pkg_config_configpb_config_pb.UpdateRequest) => {
      return request.serializeBinary();
    },
    pkg_config_configpb_config_pb.UpdateResponse.deserializeBinary
  );

  update(
    request: pkg_config_configpb_config_pb.UpdateRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_config_configpb_config_pb.UpdateResponse>;

  update(
    request: pkg_config_configpb_config_pb.UpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_config_configpb_config_pb.UpdateResponse) => void): grpcWeb.ClientReadableStream<pkg_config_configpb_config_pb.UpdateResponse>;

  update(
    request: pkg_config_configpb_config_pb.UpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_config_configpb_config_pb.UpdateResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.config.config.DynamicConfigService/Update',
        request,
        metadata || {},
        this.methodInfoUpdate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.config.config.DynamicConfigService/Update',
    request,
    metadata || {},
    this.methodInfoUpdate);
  }

}

export class StaticConfigServiceClient {
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

  methodInfoGet = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_config_configpb_config_pb.GetStaticConfigResponse,
    (request: pkg_config_configpb_config_pb.GetStaticConfigRequest) => {
      return request.serializeBinary();
    },
    pkg_config_configpb_config_pb.GetStaticConfigResponse.deserializeBinary
  );

  get(
    request: pkg_config_configpb_config_pb.GetStaticConfigRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_config_configpb_config_pb.GetStaticConfigResponse>;

  get(
    request: pkg_config_configpb_config_pb.GetStaticConfigRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_config_configpb_config_pb.GetStaticConfigResponse) => void): grpcWeb.ClientReadableStream<pkg_config_configpb_config_pb.GetStaticConfigResponse>;

  get(
    request: pkg_config_configpb_config_pb.GetStaticConfigRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_config_configpb_config_pb.GetStaticConfigResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.config.config.StaticConfigService/Get',
        request,
        metadata || {},
        this.methodInfoGet,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.config.config.StaticConfigService/Get',
    request,
    metadata || {},
    this.methodInfoGet);
  }

}

