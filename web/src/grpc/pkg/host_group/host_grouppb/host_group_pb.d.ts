import * as jspb from 'google-protobuf'

import * as google_protobuf_wrappers_pb from 'google-protobuf/google/protobuf/wrappers_pb';
import * as pkg_proto_utilpb_uuid_pb from '../../../pkg/proto/utilpb/uuid_pb';
import * as pkg_host_hostpb_host_pb from '../../../pkg/host/hostpb/host_pb';


export class HostGroup extends jspb.Message {
  getId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setId(value?: pkg_proto_utilpb_uuid_pb.UUID): HostGroup;
  hasId(): boolean;
  clearId(): HostGroup;

  getName(): string;
  setName(value: string): HostGroup;

  getHide(): google_protobuf_wrappers_pb.BoolValue | undefined;
  setHide(value?: google_protobuf_wrappers_pb.BoolValue): HostGroup;
  hasHide(): boolean;
  clearHide(): HostGroup;

  getPause(): google_protobuf_wrappers_pb.BoolValue | undefined;
  setPause(value?: google_protobuf_wrappers_pb.BoolValue): HostGroup;
  hasPause(): boolean;
  clearPause(): HostGroup;

  getHosts(): pkg_host_hostpb_host_pb.Host | undefined;
  setHosts(value?: pkg_host_hostpb_host_pb.Host): HostGroup;
  hasHosts(): boolean;
  clearHosts(): HostGroup;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HostGroup.AsObject;
  static toObject(includeInstance: boolean, msg: HostGroup): HostGroup.AsObject;
  static serializeBinaryToWriter(message: HostGroup, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HostGroup;
  static deserializeBinaryFromReader(message: HostGroup, reader: jspb.BinaryReader): HostGroup;
}

export namespace HostGroup {
  export type AsObject = {
    id?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
    name: string,
    hide?: google_protobuf_wrappers_pb.BoolValue.AsObject,
    pause?: google_protobuf_wrappers_pb.BoolValue.AsObject,
    hosts?: pkg_host_hostpb_host_pb.Host.AsObject,
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
  getHostGroupsList(): Array<HostGroup>;
  setHostGroupsList(value: Array<HostGroup>): GetAllResponse;
  clearHostGroupsList(): GetAllResponse;
  addHostGroups(value?: HostGroup, index?: number): HostGroup;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllResponse): GetAllResponse.AsObject;
  static serializeBinaryToWriter(message: GetAllResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllResponse;
  static deserializeBinaryFromReader(message: GetAllResponse, reader: jspb.BinaryReader): GetAllResponse;
}

export namespace GetAllResponse {
  export type AsObject = {
    hostGroupsList: Array<HostGroup.AsObject>,
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
  getHostGroup(): HostGroup | undefined;
  setHostGroup(value?: HostGroup): GetByIDResponse;
  hasHostGroup(): boolean;
  clearHostGroup(): GetByIDResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetByIDResponse): GetByIDResponse.AsObject;
  static serializeBinaryToWriter(message: GetByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetByIDResponse;
  static deserializeBinaryFromReader(message: GetByIDResponse, reader: jspb.BinaryReader): GetByIDResponse;
}

export namespace GetByIDResponse {
  export type AsObject = {
    hostGroup?: HostGroup.AsObject,
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
  getHostGroupsList(): Array<HostGroup>;
  setHostGroupsList(value: Array<HostGroup>): StoreRequest;
  clearHostGroupsList(): StoreRequest;
  addHostGroups(value?: HostGroup, index?: number): HostGroup;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StoreRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StoreRequest): StoreRequest.AsObject;
  static serializeBinaryToWriter(message: StoreRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StoreRequest;
  static deserializeBinaryFromReader(message: StoreRequest, reader: jspb.BinaryReader): StoreRequest;
}

export namespace StoreRequest {
  export type AsObject = {
    hostGroupsList: Array<HostGroup.AsObject>,
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
  getHostGroup(): HostGroup | undefined;
  setHostGroup(value?: HostGroup): UpdateRequest;
  hasHostGroup(): boolean;
  clearHostGroup(): UpdateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRequest): UpdateRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRequest;
  static deserializeBinaryFromReader(message: UpdateRequest, reader: jspb.BinaryReader): UpdateRequest;
}

export namespace UpdateRequest {
  export type AsObject = {
    hostGroup?: HostGroup.AsObject,
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

