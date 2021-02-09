import * as jspb from 'google-protobuf'

import * as google_protobuf_wrappers_pb from 'google-protobuf/google/protobuf/wrappers_pb';
import * as pkg_proto_utilpb_uuid_pb from '../../../pkg/proto/utilpb/uuid_pb';


export class Check extends jspb.Message {
  getServiceId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setServiceId(value?: pkg_proto_utilpb_uuid_pb.UUID): Check;
  hasServiceId(): boolean;
  clearServiceId(): Check;

  getRoundId(): number;
  setRoundId(value: number): Check;

  getLog(): string;
  setLog(value: string): Check;

  getErr(): string;
  setErr(value: string): Check;

  getPassed(): google_protobuf_wrappers_pb.BoolValue | undefined;
  setPassed(value?: google_protobuf_wrappers_pb.BoolValue): Check;
  hasPassed(): boolean;
  clearPassed(): Check;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Check.AsObject;
  static toObject(includeInstance: boolean, msg: Check): Check.AsObject;
  static serializeBinaryToWriter(message: Check, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Check;
  static deserializeBinaryFromReader(message: Check, reader: jspb.BinaryReader): Check;
}

export namespace Check {
  export type AsObject = {
    serviceId?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
    roundId: number,
    log: string,
    err: string,
    passed?: google_protobuf_wrappers_pb.BoolValue.AsObject,
  }
}

export class GetAllByRoundIDRequest extends jspb.Message {
  getRoundId(): number;
  setRoundId(value: number): GetAllByRoundIDRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllByRoundIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllByRoundIDRequest): GetAllByRoundIDRequest.AsObject;
  static serializeBinaryToWriter(message: GetAllByRoundIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllByRoundIDRequest;
  static deserializeBinaryFromReader(message: GetAllByRoundIDRequest, reader: jspb.BinaryReader): GetAllByRoundIDRequest;
}

export namespace GetAllByRoundIDRequest {
  export type AsObject = {
    roundId: number,
  }
}

export class GetAllByRoundIDResponse extends jspb.Message {
  getChecksList(): Array<Check>;
  setChecksList(value: Array<Check>): GetAllByRoundIDResponse;
  clearChecksList(): GetAllByRoundIDResponse;
  addChecks(value?: Check, index?: number): Check;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllByRoundIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllByRoundIDResponse): GetAllByRoundIDResponse.AsObject;
  static serializeBinaryToWriter(message: GetAllByRoundIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllByRoundIDResponse;
  static deserializeBinaryFromReader(message: GetAllByRoundIDResponse, reader: jspb.BinaryReader): GetAllByRoundIDResponse;
}

export namespace GetAllByRoundIDResponse {
  export type AsObject = {
    checksList: Array<Check.AsObject>,
  }
}

export class GetByRoundServiceIDRequest extends jspb.Message {
  getServiceId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setServiceId(value?: pkg_proto_utilpb_uuid_pb.UUID): GetByRoundServiceIDRequest;
  hasServiceId(): boolean;
  clearServiceId(): GetByRoundServiceIDRequest;

  getRoundId(): number;
  setRoundId(value: number): GetByRoundServiceIDRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetByRoundServiceIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetByRoundServiceIDRequest): GetByRoundServiceIDRequest.AsObject;
  static serializeBinaryToWriter(message: GetByRoundServiceIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetByRoundServiceIDRequest;
  static deserializeBinaryFromReader(message: GetByRoundServiceIDRequest, reader: jspb.BinaryReader): GetByRoundServiceIDRequest;
}

export namespace GetByRoundServiceIDRequest {
  export type AsObject = {
    serviceId?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
    roundId: number,
  }
}

export class GetByRoundServiceIDResponse extends jspb.Message {
  getCheck(): Check | undefined;
  setCheck(value?: Check): GetByRoundServiceIDResponse;
  hasCheck(): boolean;
  clearCheck(): GetByRoundServiceIDResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetByRoundServiceIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetByRoundServiceIDResponse): GetByRoundServiceIDResponse.AsObject;
  static serializeBinaryToWriter(message: GetByRoundServiceIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetByRoundServiceIDResponse;
  static deserializeBinaryFromReader(message: GetByRoundServiceIDResponse, reader: jspb.BinaryReader): GetByRoundServiceIDResponse;
}

export namespace GetByRoundServiceIDResponse {
  export type AsObject = {
    check?: Check.AsObject,
  }
}

export class GetAllByServiceIDRequest extends jspb.Message {
  getServiceId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setServiceId(value?: pkg_proto_utilpb_uuid_pb.UUID): GetAllByServiceIDRequest;
  hasServiceId(): boolean;
  clearServiceId(): GetAllByServiceIDRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllByServiceIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllByServiceIDRequest): GetAllByServiceIDRequest.AsObject;
  static serializeBinaryToWriter(message: GetAllByServiceIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllByServiceIDRequest;
  static deserializeBinaryFromReader(message: GetAllByServiceIDRequest, reader: jspb.BinaryReader): GetAllByServiceIDRequest;
}

export namespace GetAllByServiceIDRequest {
  export type AsObject = {
    serviceId?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
  }
}

export class GetAllByServiceIDResponse extends jspb.Message {
  getChecksList(): Array<Check>;
  setChecksList(value: Array<Check>): GetAllByServiceIDResponse;
  clearChecksList(): GetAllByServiceIDResponse;
  addChecks(value?: Check, index?: number): Check;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllByServiceIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllByServiceIDResponse): GetAllByServiceIDResponse.AsObject;
  static serializeBinaryToWriter(message: GetAllByServiceIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllByServiceIDResponse;
  static deserializeBinaryFromReader(message: GetAllByServiceIDResponse, reader: jspb.BinaryReader): GetAllByServiceIDResponse;
}

export namespace GetAllByServiceIDResponse {
  export type AsObject = {
    checksList: Array<Check.AsObject>,
  }
}

