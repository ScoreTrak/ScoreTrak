import * as jspb from 'google-protobuf'

import * as google_protobuf_wrappers_pb from 'google-protobuf/google/protobuf/wrappers_pb';
import * as pkg_proto_utilpb_uuid_pb from '../../../pkg/proto/utilpb/uuid_pb';
import * as pkg_property_propertypb_property_pb from '../../../pkg/property/propertypb/property_pb';
import * as pkg_check_checkpb_check_pb from '../../../pkg/check/checkpb/check_pb';


export class Service extends jspb.Message {
  getId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setId(value?: pkg_proto_utilpb_uuid_pb.UUID): Service;
  hasId(): boolean;
  clearId(): Service;

  getName(): string;
  setName(value: string): Service;

  getDisplayName(): string;
  setDisplayName(value: string): Service;

  getWeight(): google_protobuf_wrappers_pb.UInt64Value | undefined;
  setWeight(value?: google_protobuf_wrappers_pb.UInt64Value): Service;
  hasWeight(): boolean;
  clearWeight(): Service;

  getPointsBoost(): google_protobuf_wrappers_pb.UInt64Value | undefined;
  setPointsBoost(value?: google_protobuf_wrappers_pb.UInt64Value): Service;
  hasPointsBoost(): boolean;
  clearPointsBoost(): Service;

  getRoundUnits(): number;
  setRoundUnits(value: number): Service;

  getRoundDelay(): google_protobuf_wrappers_pb.UInt64Value | undefined;
  setRoundDelay(value?: google_protobuf_wrappers_pb.UInt64Value): Service;
  hasRoundDelay(): boolean;
  clearRoundDelay(): Service;

  getServiceGroupId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setServiceGroupId(value?: pkg_proto_utilpb_uuid_pb.UUID): Service;
  hasServiceGroupId(): boolean;
  clearServiceGroupId(): Service;

  getHostId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setHostId(value?: pkg_proto_utilpb_uuid_pb.UUID): Service;
  hasHostId(): boolean;
  clearHostId(): Service;

  getHide(): google_protobuf_wrappers_pb.BoolValue | undefined;
  setHide(value?: google_protobuf_wrappers_pb.BoolValue): Service;
  hasHide(): boolean;
  clearHide(): Service;

  getPause(): google_protobuf_wrappers_pb.BoolValue | undefined;
  setPause(value?: google_protobuf_wrappers_pb.BoolValue): Service;
  hasPause(): boolean;
  clearPause(): Service;

  getPropertiesList(): Array<pkg_property_propertypb_property_pb.Property>;
  setPropertiesList(value: Array<pkg_property_propertypb_property_pb.Property>): Service;
  clearPropertiesList(): Service;
  addProperties(value?: pkg_property_propertypb_property_pb.Property, index?: number): pkg_property_propertypb_property_pb.Property;

  getChecksList(): Array<pkg_check_checkpb_check_pb.Check>;
  setChecksList(value: Array<pkg_check_checkpb_check_pb.Check>): Service;
  clearChecksList(): Service;
  addChecks(value?: pkg_check_checkpb_check_pb.Check, index?: number): pkg_check_checkpb_check_pb.Check;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Service.AsObject;
  static toObject(includeInstance: boolean, msg: Service): Service.AsObject;
  static serializeBinaryToWriter(message: Service, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Service;
  static deserializeBinaryFromReader(message: Service, reader: jspb.BinaryReader): Service;
}

export namespace Service {
  export type AsObject = {
    id?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
    name: string,
    displayName: string,
    weight?: google_protobuf_wrappers_pb.UInt64Value.AsObject,
    pointsBoost?: google_protobuf_wrappers_pb.UInt64Value.AsObject,
    roundUnits: number,
    roundDelay?: google_protobuf_wrappers_pb.UInt64Value.AsObject,
    serviceGroupId?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
    hostId?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
    hide?: google_protobuf_wrappers_pb.BoolValue.AsObject,
    pause?: google_protobuf_wrappers_pb.BoolValue.AsObject,
    propertiesList: Array<pkg_property_propertypb_property_pb.Property.AsObject>,
    checksList: Array<pkg_check_checkpb_check_pb.Check.AsObject>,
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
  getServicesList(): Array<Service>;
  setServicesList(value: Array<Service>): GetAllResponse;
  clearServicesList(): GetAllResponse;
  addServices(value?: Service, index?: number): Service;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllResponse): GetAllResponse.AsObject;
  static serializeBinaryToWriter(message: GetAllResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllResponse;
  static deserializeBinaryFromReader(message: GetAllResponse, reader: jspb.BinaryReader): GetAllResponse;
}

export namespace GetAllResponse {
  export type AsObject = {
    servicesList: Array<Service.AsObject>,
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
  getService(): Service | undefined;
  setService(value?: Service): GetByIDResponse;
  hasService(): boolean;
  clearService(): GetByIDResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetByIDResponse): GetByIDResponse.AsObject;
  static serializeBinaryToWriter(message: GetByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetByIDResponse;
  static deserializeBinaryFromReader(message: GetByIDResponse, reader: jspb.BinaryReader): GetByIDResponse;
}

export namespace GetByIDResponse {
  export type AsObject = {
    service?: Service.AsObject,
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
  getServicesList(): Array<Service>;
  setServicesList(value: Array<Service>): StoreRequest;
  clearServicesList(): StoreRequest;
  addServices(value?: Service, index?: number): Service;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StoreRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StoreRequest): StoreRequest.AsObject;
  static serializeBinaryToWriter(message: StoreRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StoreRequest;
  static deserializeBinaryFromReader(message: StoreRequest, reader: jspb.BinaryReader): StoreRequest;
}

export namespace StoreRequest {
  export type AsObject = {
    servicesList: Array<Service.AsObject>,
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
  getService(): Service | undefined;
  setService(value?: Service): UpdateRequest;
  hasService(): boolean;
  clearService(): UpdateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRequest): UpdateRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRequest;
  static deserializeBinaryFromReader(message: UpdateRequest, reader: jspb.BinaryReader): UpdateRequest;
}

export namespace UpdateRequest {
  export type AsObject = {
    service?: Service.AsObject,
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

export class TestServiceRequest extends jspb.Message {
  getId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setId(value?: pkg_proto_utilpb_uuid_pb.UUID): TestServiceRequest;
  hasId(): boolean;
  clearId(): TestServiceRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TestServiceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TestServiceRequest): TestServiceRequest.AsObject;
  static serializeBinaryToWriter(message: TestServiceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TestServiceRequest;
  static deserializeBinaryFromReader(message: TestServiceRequest, reader: jspb.BinaryReader): TestServiceRequest;
}

export namespace TestServiceRequest {
  export type AsObject = {
    id?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
  }
}

export class TestServiceResponse extends jspb.Message {
  getCheck(): pkg_check_checkpb_check_pb.Check | undefined;
  setCheck(value?: pkg_check_checkpb_check_pb.Check): TestServiceResponse;
  hasCheck(): boolean;
  clearCheck(): TestServiceResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TestServiceResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TestServiceResponse): TestServiceResponse.AsObject;
  static serializeBinaryToWriter(message: TestServiceResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TestServiceResponse;
  static deserializeBinaryFromReader(message: TestServiceResponse, reader: jspb.BinaryReader): TestServiceResponse;
}

export namespace TestServiceResponse {
  export type AsObject = {
    check?: pkg_check_checkpb_check_pb.Check.AsObject,
  }
}

