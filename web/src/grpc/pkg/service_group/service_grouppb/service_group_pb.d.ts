import * as jspb from 'google-protobuf'

import * as google_protobuf_wrappers_pb from 'google-protobuf/google/protobuf/wrappers_pb';
import * as pkg_proto_utilpb_uuid_pb from '../../../pkg/proto/utilpb/uuid_pb';
import * as pkg_service_servicepb_service_pb from '../../../pkg/service/servicepb/service_pb';


export class ServiceGroup extends jspb.Message {
  getId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setId(value?: pkg_proto_utilpb_uuid_pb.UUID): ServiceGroup;
  hasId(): boolean;
  clearId(): ServiceGroup;

  getName(): string;
  setName(value: string): ServiceGroup;

  getDisplayName(): string;
  setDisplayName(value: string): ServiceGroup;

  getEnabled(): google_protobuf_wrappers_pb.BoolValue | undefined;
  setEnabled(value?: google_protobuf_wrappers_pb.BoolValue): ServiceGroup;
  hasEnabled(): boolean;
  clearEnabled(): ServiceGroup;

  getSkipHelper(): boolean;
  setSkipHelper(value: boolean): ServiceGroup;

  getLabel(): string;
  setLabel(value: string): ServiceGroup;

  getServicesList(): Array<pkg_service_servicepb_service_pb.Service>;
  setServicesList(value: Array<pkg_service_servicepb_service_pb.Service>): ServiceGroup;
  clearServicesList(): ServiceGroup;
  addServices(value?: pkg_service_servicepb_service_pb.Service, index?: number): pkg_service_servicepb_service_pb.Service;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ServiceGroup.AsObject;
  static toObject(includeInstance: boolean, msg: ServiceGroup): ServiceGroup.AsObject;
  static serializeBinaryToWriter(message: ServiceGroup, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ServiceGroup;
  static deserializeBinaryFromReader(message: ServiceGroup, reader: jspb.BinaryReader): ServiceGroup;
}

export namespace ServiceGroup {
  export type AsObject = {
    id?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
    name: string,
    displayName: string,
    enabled?: google_protobuf_wrappers_pb.BoolValue.AsObject,
    skipHelper: boolean,
    label: string,
    servicesList: Array<pkg_service_servicepb_service_pb.Service.AsObject>,
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
  getServiceGroupsList(): Array<ServiceGroup>;
  setServiceGroupsList(value: Array<ServiceGroup>): GetAllResponse;
  clearServiceGroupsList(): GetAllResponse;
  addServiceGroups(value?: ServiceGroup, index?: number): ServiceGroup;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllResponse): GetAllResponse.AsObject;
  static serializeBinaryToWriter(message: GetAllResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllResponse;
  static deserializeBinaryFromReader(message: GetAllResponse, reader: jspb.BinaryReader): GetAllResponse;
}

export namespace GetAllResponse {
  export type AsObject = {
    serviceGroupsList: Array<ServiceGroup.AsObject>,
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
  getServiceGroup(): ServiceGroup | undefined;
  setServiceGroup(value?: ServiceGroup): GetByIDResponse;
  hasServiceGroup(): boolean;
  clearServiceGroup(): GetByIDResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetByIDResponse): GetByIDResponse.AsObject;
  static serializeBinaryToWriter(message: GetByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetByIDResponse;
  static deserializeBinaryFromReader(message: GetByIDResponse, reader: jspb.BinaryReader): GetByIDResponse;
}

export namespace GetByIDResponse {
  export type AsObject = {
    serviceGroup?: ServiceGroup.AsObject,
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
  getServiceGroup(): ServiceGroup | undefined;
  setServiceGroup(value?: ServiceGroup): StoreRequest;
  hasServiceGroup(): boolean;
  clearServiceGroup(): StoreRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StoreRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StoreRequest): StoreRequest.AsObject;
  static serializeBinaryToWriter(message: StoreRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StoreRequest;
  static deserializeBinaryFromReader(message: StoreRequest, reader: jspb.BinaryReader): StoreRequest;
}

export namespace StoreRequest {
  export type AsObject = {
    serviceGroup?: ServiceGroup.AsObject,
  }
}

export class StoreResponse extends jspb.Message {
  getId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setId(value?: pkg_proto_utilpb_uuid_pb.UUID): StoreResponse;
  hasId(): boolean;
  clearId(): StoreResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StoreResponse.AsObject;
  static toObject(includeInstance: boolean, msg: StoreResponse): StoreResponse.AsObject;
  static serializeBinaryToWriter(message: StoreResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StoreResponse;
  static deserializeBinaryFromReader(message: StoreResponse, reader: jspb.BinaryReader): StoreResponse;
}

export namespace StoreResponse {
  export type AsObject = {
    id?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
  }
}

export class UpdateRequest extends jspb.Message {
  getServiceGroup(): ServiceGroup | undefined;
  setServiceGroup(value?: ServiceGroup): UpdateRequest;
  hasServiceGroup(): boolean;
  clearServiceGroup(): UpdateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRequest): UpdateRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRequest;
  static deserializeBinaryFromReader(message: UpdateRequest, reader: jspb.BinaryReader): UpdateRequest;
}

export namespace UpdateRequest {
  export type AsObject = {
    serviceGroup?: ServiceGroup.AsObject,
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

export class RedeployRequest extends jspb.Message {
  getId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setId(value?: pkg_proto_utilpb_uuid_pb.UUID): RedeployRequest;
  hasId(): boolean;
  clearId(): RedeployRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RedeployRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RedeployRequest): RedeployRequest.AsObject;
  static serializeBinaryToWriter(message: RedeployRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RedeployRequest;
  static deserializeBinaryFromReader(message: RedeployRequest, reader: jspb.BinaryReader): RedeployRequest;
}

export namespace RedeployRequest {
  export type AsObject = {
    id?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
  }
}

export class RedeployResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RedeployResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RedeployResponse): RedeployResponse.AsObject;
  static serializeBinaryToWriter(message: RedeployResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RedeployResponse;
  static deserializeBinaryFromReader(message: RedeployResponse, reader: jspb.BinaryReader): RedeployResponse;
}

export namespace RedeployResponse {
  export type AsObject = {
  }
}

