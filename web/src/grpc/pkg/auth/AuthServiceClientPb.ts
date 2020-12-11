/**
 * @fileoverview gRPC-Web generated client stub for pkg.auth
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as pkg_auth_auth_pb from '../../pkg/auth/auth_pb';


export class AuthServiceClient {
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

  methodInfoLogin = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_auth_auth_pb.LoginResponse,
    (request: pkg_auth_auth_pb.LoginRequest) => {
      return request.serializeBinary();
    },
    pkg_auth_auth_pb.LoginResponse.deserializeBinary
  );

  login(
    request: pkg_auth_auth_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_auth_auth_pb.LoginResponse>;

  login(
    request: pkg_auth_auth_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_auth_auth_pb.LoginResponse) => void): grpcWeb.ClientReadableStream<pkg_auth_auth_pb.LoginResponse>;

  login(
    request: pkg_auth_auth_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_auth_auth_pb.LoginResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.auth.AuthService/Login',
        request,
        metadata || {},
        this.methodInfoLogin,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.auth.AuthService/Login',
    request,
    metadata || {},
    this.methodInfoLogin);
  }

}

