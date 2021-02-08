import * as jspb from 'google-protobuf'

import * as google_protobuf_wrappers_pb from 'google-protobuf/google/protobuf/wrappers_pb';
import * as pkg_proto_utilpb_uuid_pb from '../../../pkg/proto/utilpb/uuid_pb';


export class Property extends jspb.Message {
  getServiceId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setServiceId(value?: pkg_proto_utilpb_uuid_pb.UUID): Property;
  hasServiceId(): boolean;
  clearServiceId(): Property;

  getKey(): string;
  setKey(value: string): Property;

  getValue(): google_protobuf_wrappers_pb.StringValue | undefined;
  setValue(value?: google_protobuf_wrappers_pb.StringValue): Property;
  hasValue(): boolean;
  clearValue(): Property;

  getStatus(): Status;
  setStatus(value: Status): Property;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Property.AsObject;
  static toObject(includeInstance: boolean, msg: Property): Property.AsObject;
  static serializeBinaryToWriter(message: Property, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Property;
  static deserializeBinaryFromReader(message: Property, reader: jspb.BinaryReader): Property;
}

export namespace Property {
  export type AsObject = {
    serviceId?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
    key: string,
    value?: google_protobuf_wrappers_pb.StringValue.AsObject,
    status: Status,
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
  getPropertiesList(): Array<Property>;
  setPropertiesList(value: Array<Property>): GetAllResponse;
  clearPropertiesList(): GetAllResponse;
  addProperties(value?: Property, index?: number): Property;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllResponse): GetAllResponse.AsObject;
  static serializeBinaryToWriter(message: GetAllResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllResponse;
  static deserializeBinaryFromReader(message: GetAllResponse, reader: jspb.BinaryReader): GetAllResponse;
}

export namespace GetAllResponse {
  export type AsObject = {
    propertiesList: Array<Property.AsObject>,
  }
}

export class GetByServiceIDKeyRequest extends jspb.Message {
  getServiceId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setServiceId(value?: pkg_proto_utilpb_uuid_pb.UUID): GetByServiceIDKeyRequest;
  hasServiceId(): boolean;
  clearServiceId(): GetByServiceIDKeyRequest;

  getKey(): string;
  setKey(value: string): GetByServiceIDKeyRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetByServiceIDKeyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetByServiceIDKeyRequest): GetByServiceIDKeyRequest.AsObject;
  static serializeBinaryToWriter(message: GetByServiceIDKeyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetByServiceIDKeyRequest;
  static deserializeBinaryFromReader(message: GetByServiceIDKeyRequest, reader: jspb.BinaryReader): GetByServiceIDKeyRequest;
}

export namespace GetByServiceIDKeyRequest {
  export type AsObject = {
    serviceId?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
    key: string,
  }
}

export class GetByServiceIDKeyResponse extends jspb.Message {
  getProperty(): Property | undefined;
  setProperty(value?: Property): GetByServiceIDKeyResponse;
  hasProperty(): boolean;
  clearProperty(): GetByServiceIDKeyResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetByServiceIDKeyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetByServiceIDKeyResponse): GetByServiceIDKeyResponse.AsObject;
  static serializeBinaryToWriter(message: GetByServiceIDKeyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetByServiceIDKeyResponse;
  static deserializeBinaryFromReader(message: GetByServiceIDKeyResponse, reader: jspb.BinaryReader): GetByServiceIDKeyResponse;
}

export namespace GetByServiceIDKeyResponse {
  export type AsObject = {
    property?: Property.AsObject,
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
  getPropertiesList(): Array<Property>;
  setPropertiesList(value: Array<Property>): GetAllByServiceIDResponse;
  clearPropertiesList(): GetAllByServiceIDResponse;
  addProperties(value?: Property, index?: number): Property;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllByServiceIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllByServiceIDResponse): GetAllByServiceIDResponse.AsObject;
  static serializeBinaryToWriter(message: GetAllByServiceIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllByServiceIDResponse;
  static deserializeBinaryFromReader(message: GetAllByServiceIDResponse, reader: jspb.BinaryReader): GetAllByServiceIDResponse;
}

export namespace GetAllByServiceIDResponse {
  export type AsObject = {
    propertiesList: Array<Property.AsObject>,
  }
}

export class DeleteRequest extends jspb.Message {
  getServiceId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setServiceId(value?: pkg_proto_utilpb_uuid_pb.UUID): DeleteRequest;
  hasServiceId(): boolean;
  clearServiceId(): DeleteRequest;

  getKey(): string;
  setKey(value: string): DeleteRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRequest): DeleteRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRequest;
  static deserializeBinaryFromReader(message: DeleteRequest, reader: jspb.BinaryReader): DeleteRequest;
}

export namespace DeleteRequest {
  export type AsObject = {
    serviceId?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
    key: string,
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
  getPropertiesList(): Array<Property>;
  setPropertiesList(value: Array<Property>): StoreRequest;
  clearPropertiesList(): StoreRequest;
  addProperties(value?: Property, index?: number): Property;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StoreRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StoreRequest): StoreRequest.AsObject;
  static serializeBinaryToWriter(message: StoreRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StoreRequest;
  static deserializeBinaryFromReader(message: StoreRequest, reader: jspb.BinaryReader): StoreRequest;
}

export namespace StoreRequest {
  export type AsObject = {
    propertiesList: Array<Property.AsObject>,
  }
}

export class StoreResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StoreResponse.AsObject;
  static toObject(includeInstance: boolean, msg: StoreResponse): StoreResponse.AsObject;
  static serializeBinaryToWriter(message: StoreResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StoreResponse;
  static deserializeBinaryFromReader(message: StoreResponse, reader: jspb.BinaryReader): StoreResponse;
}

export namespace StoreResponse {
  export type AsObject = {
  }
}

export class UpdateRequest extends jspb.Message {
  getProperty(): Property | undefined;
  setProperty(value?: Property): UpdateRequest;
  hasProperty(): boolean;
  clearProperty(): UpdateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRequest): UpdateRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRequest;
  static deserializeBinaryFromReader(message: UpdateRequest, reader: jspb.BinaryReader): UpdateRequest;
}

export namespace UpdateRequest {
  export type AsObject = {
    property?: Property.AsObject,
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

export enum Status { 
  STATUS_NOT_SET = 0,
  VIEW = 1,
  EDIT = 2,
  HIDE = 3,
}
