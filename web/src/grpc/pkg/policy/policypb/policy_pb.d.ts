import * as jspb from 'google-protobuf'

import * as google_protobuf_wrappers_pb from 'google-protobuf/google/protobuf/wrappers_pb';


export class Policy extends jspb.Message {
  getAllowUnauthenticatedUsers(): google_protobuf_wrappers_pb.BoolValue | undefined;
  setAllowUnauthenticatedUsers(value?: google_protobuf_wrappers_pb.BoolValue): Policy;
  hasAllowUnauthenticatedUsers(): boolean;
  clearAllowUnauthenticatedUsers(): Policy;

  getAllowChangingUsernamesAndPasswords(): google_protobuf_wrappers_pb.BoolValue | undefined;
  setAllowChangingUsernamesAndPasswords(value?: google_protobuf_wrappers_pb.BoolValue): Policy;
  hasAllowChangingUsernamesAndPasswords(): boolean;
  clearAllowChangingUsernamesAndPasswords(): Policy;

  getShowPoints(): google_protobuf_wrappers_pb.BoolValue | undefined;
  setShowPoints(value?: google_protobuf_wrappers_pb.BoolValue): Policy;
  hasShowPoints(): boolean;
  clearShowPoints(): Policy;

  getShowAddresses(): google_protobuf_wrappers_pb.BoolValue | undefined;
  setShowAddresses(value?: google_protobuf_wrappers_pb.BoolValue): Policy;
  hasShowAddresses(): boolean;
  clearShowAddresses(): Policy;

  getAllowRedTeamLaunchingServiceTestsManually(): google_protobuf_wrappers_pb.BoolValue | undefined;
  setAllowRedTeamLaunchingServiceTestsManually(value?: google_protobuf_wrappers_pb.BoolValue): Policy;
  hasAllowRedTeamLaunchingServiceTestsManually(): boolean;
  clearAllowRedTeamLaunchingServiceTestsManually(): Policy;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Policy.AsObject;
  static toObject(includeInstance: boolean, msg: Policy): Policy.AsObject;
  static serializeBinaryToWriter(message: Policy, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Policy;
  static deserializeBinaryFromReader(message: Policy, reader: jspb.BinaryReader): Policy;
}

export namespace Policy {
  export type AsObject = {
    allowUnauthenticatedUsers?: google_protobuf_wrappers_pb.BoolValue.AsObject,
    allowChangingUsernamesAndPasswords?: google_protobuf_wrappers_pb.BoolValue.AsObject,
    showPoints?: google_protobuf_wrappers_pb.BoolValue.AsObject,
    showAddresses?: google_protobuf_wrappers_pb.BoolValue.AsObject,
    allowRedTeamLaunchingServiceTestsManually?: google_protobuf_wrappers_pb.BoolValue.AsObject,
  }
}

export class GetRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetRequest): GetRequest.AsObject;
  static serializeBinaryToWriter(message: GetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRequest;
  static deserializeBinaryFromReader(message: GetRequest, reader: jspb.BinaryReader): GetRequest;
}

export namespace GetRequest {
  export type AsObject = {
  }
}

export class GetResponse extends jspb.Message {
  getPolicy(): Policy | undefined;
  setPolicy(value?: Policy): GetResponse;
  hasPolicy(): boolean;
  clearPolicy(): GetResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetResponse): GetResponse.AsObject;
  static serializeBinaryToWriter(message: GetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetResponse;
  static deserializeBinaryFromReader(message: GetResponse, reader: jspb.BinaryReader): GetResponse;
}

export namespace GetResponse {
  export type AsObject = {
    policy?: Policy.AsObject,
  }
}

export class UpdateRequest extends jspb.Message {
  getPolicy(): Policy | undefined;
  setPolicy(value?: Policy): UpdateRequest;
  hasPolicy(): boolean;
  clearPolicy(): UpdateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRequest): UpdateRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRequest;
  static deserializeBinaryFromReader(message: UpdateRequest, reader: jspb.BinaryReader): UpdateRequest;
}

export namespace UpdateRequest {
  export type AsObject = {
    policy?: Policy.AsObject,
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

