import * as jspb from 'google-protobuf'

import * as google_protobuf_wrappers_pb from 'google-protobuf/google/protobuf/wrappers_pb';
import * as pkg_proto_utilpb_uuid_pb from '../../../pkg/proto/utilpb/uuid_pb';
import * as pkg_service_servicepb_service_pb from '../../../pkg/service/servicepb/service_pb';


export class Host extends jspb.Message {
  getId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setId(value?: pkg_proto_utilpb_uuid_pb.UUID): Host;
  hasId(): boolean;
  clearId(): Host;

  getAddress(): string;
  setAddress(value: string): Host;

  getHostGroupId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setHostGroupId(value?: pkg_proto_utilpb_uuid_pb.UUID): Host;
  hasHostGroupId(): boolean;
  clearHostGroupId(): Host;

  getTeamId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setTeamId(value?: pkg_proto_utilpb_uuid_pb.UUID): Host;
  hasTeamId(): boolean;
  clearTeamId(): Host;

  getPause(): google_protobuf_wrappers_pb.BoolValue | undefined;
  setPause(value?: google_protobuf_wrappers_pb.BoolValue): Host;
  hasPause(): boolean;
  clearPause(): Host;

  getHide(): google_protobuf_wrappers_pb.BoolValue | undefined;
  setHide(value?: google_protobuf_wrappers_pb.BoolValue): Host;
  hasHide(): boolean;
  clearHide(): Host;

  getEditHost(): google_protobuf_wrappers_pb.BoolValue | undefined;
  setEditHost(value?: google_protobuf_wrappers_pb.BoolValue): Host;
  hasEditHost(): boolean;
  clearEditHost(): Host;

  getServicesList(): Array<pkg_service_servicepb_service_pb.Service>;
  setServicesList(value: Array<pkg_service_servicepb_service_pb.Service>): Host;
  clearServicesList(): Host;
  addServices(value?: pkg_service_servicepb_service_pb.Service, index?: number): pkg_service_servicepb_service_pb.Service;

  getAddressListRange(): google_protobuf_wrappers_pb.StringValue | undefined;
  setAddressListRange(value?: google_protobuf_wrappers_pb.StringValue): Host;
  hasAddressListRange(): boolean;
  clearAddressListRange(): Host;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Host.AsObject;
  static toObject(includeInstance: boolean, msg: Host): Host.AsObject;
  static serializeBinaryToWriter(message: Host, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Host;
  static deserializeBinaryFromReader(message: Host, reader: jspb.BinaryReader): Host;
}

export namespace Host {
  export type AsObject = {
    id?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
    address: string,
    hostGroupId?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
    teamId?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
    pause?: google_protobuf_wrappers_pb.BoolValue.AsObject,
    hide?: google_protobuf_wrappers_pb.BoolValue.AsObject,
    editHost?: google_protobuf_wrappers_pb.BoolValue.AsObject,
    servicesList: Array<pkg_service_servicepb_service_pb.Service.AsObject>,
    addressListRange?: google_protobuf_wrappers_pb.StringValue.AsObject,
  }
}

export class GetAllRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllRequest): GetAllRequest.AsObject;
  static serializeBinaryToWriter(message: GetAllRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllRequest;
  static deserializeBinaryFromReader(message: GetAllRequest, reader: jspb.BinaryReader): GetAllRequest;
}

export namespace GetAllRequest {
  export type AsObject = {
  }
}

export class GetAllResponse extends jspb.Message {
  getHostsList(): Array<Host>;
  setHostsList(value: Array<Host>): GetAllResponse;
  clearHostsList(): GetAllResponse;
  addHosts(value?: Host, index?: number): Host;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllResponse): GetAllResponse.AsObject;
  static serializeBinaryToWriter(message: GetAllResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllResponse;
  static deserializeBinaryFromReader(message: GetAllResponse, reader: jspb.BinaryReader): GetAllResponse;
}

export namespace GetAllResponse {
  export type AsObject = {
    hostsList: Array<Host.AsObject>,
  }
}

export class GetByIDRequest extends jspb.Message {
  getId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setId(value?: pkg_proto_utilpb_uuid_pb.UUID): GetByIDRequest;
  hasId(): boolean;
  clearId(): GetByIDRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetByIDRequest): GetByIDRequest.AsObject;
  static serializeBinaryToWriter(message: GetByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetByIDRequest;
  static deserializeBinaryFromReader(message: GetByIDRequest, reader: jspb.BinaryReader): GetByIDRequest;
}

export namespace GetByIDRequest {
  export type AsObject = {
    id?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
  }
}

export class GetByIDResponse extends jspb.Message {
  getHost(): Host | undefined;
  setHost(value?: Host): GetByIDResponse;
  hasHost(): boolean;
  clearHost(): GetByIDResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetByIDResponse): GetByIDResponse.AsObject;
  static serializeBinaryToWriter(message: GetByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetByIDResponse;
  static deserializeBinaryFromReader(message: GetByIDResponse, reader: jspb.BinaryReader): GetByIDResponse;
}

export namespace GetByIDResponse {
  export type AsObject = {
    host?: Host.AsObject,
  }
}

export class DeleteRequest extends jspb.Message {
  getId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setId(value?: pkg_proto_utilpb_uuid_pb.UUID): DeleteRequest;
  hasId(): boolean;
  clearId(): DeleteRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRequest): DeleteRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRequest;
  static deserializeBinaryFromReader(message: DeleteRequest, reader: jspb.BinaryReader): DeleteRequest;
}

export namespace DeleteRequest {
  export type AsObject = {
    id?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
  }
}

export class DeleteResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteResponse): DeleteResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteResponse;
  static deserializeBinaryFromReader(message: DeleteResponse, reader: jspb.BinaryReader): DeleteResponse;
}

export namespace DeleteResponse {
  export type AsObject = {
  }
}

export class StoreRequest extends jspb.Message {
  getHostsList(): Array<Host>;
  setHostsList(value: Array<Host>): StoreRequest;
  clearHostsList(): StoreRequest;
  addHosts(value?: Host, index?: number): Host;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StoreRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StoreRequest): StoreRequest.AsObject;
  static serializeBinaryToWriter(message: StoreRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StoreRequest;
  static deserializeBinaryFromReader(message: StoreRequest, reader: jspb.BinaryReader): StoreRequest;
}

export namespace StoreRequest {
  export type AsObject = {
    hostsList: Array<Host.AsObject>,
  }
}

export class StoreResponse extends jspb.Message {
  getIdsList(): Array<pkg_proto_utilpb_uuid_pb.UUID>;
  setIdsList(value: Array<pkg_proto_utilpb_uuid_pb.UUID>): StoreResponse;
  clearIdsList(): StoreResponse;
  addIds(value?: pkg_proto_utilpb_uuid_pb.UUID, index?: number): pkg_proto_utilpb_uuid_pb.UUID;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StoreResponse.AsObject;
  static toObject(includeInstance: boolean, msg: StoreResponse): StoreResponse.AsObject;
  static serializeBinaryToWriter(message: StoreResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StoreResponse;
  static deserializeBinaryFromReader(message: StoreResponse, reader: jspb.BinaryReader): StoreResponse;
}

export namespace StoreResponse {
  export type AsObject = {
    idsList: Array<pkg_proto_utilpb_uuid_pb.UUID.AsObject>,
  }
}

export class UpdateRequest extends jspb.Message {
  getHost(): Host | undefined;
  setHost(value?: Host): UpdateRequest;
  hasHost(): boolean;
  clearHost(): UpdateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRequest): UpdateRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRequest;
  static deserializeBinaryFromReader(message: UpdateRequest, reader: jspb.BinaryReader): UpdateRequest;
}

export namespace UpdateRequest {
  export type AsObject = {
    host?: Host.AsObject,
  }
}

export class UpdateResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateResponse): UpdateResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateResponse;
  static deserializeBinaryFromReader(message: UpdateResponse, reader: jspb.BinaryReader): UpdateResponse;
}

export namespace UpdateResponse {
  export type AsObject = {
  }
}

