import * as jspb from 'google-protobuf'

import * as pkg_proto_utilpb_uuid_pb from '../../../pkg/proto/utilpb/uuid_pb';


export class User extends jspb.Message {
  getId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setId(value?: pkg_proto_utilpb_uuid_pb.UUID): User;
  hasId(): boolean;
  clearId(): User;

  getUsername(): string;
  setUsername(value: string): User;

  getTeamId(): pkg_proto_utilpb_uuid_pb.UUID | undefined;
  setTeamId(value?: pkg_proto_utilpb_uuid_pb.UUID): User;
  hasTeamId(): boolean;
  clearTeamId(): User;

  getPassword(): string;
  setPassword(value: string): User;

  getRole(): Role;
  setRole(value: Role): User;

  getPasswordHash(): string;
  setPasswordHash(value: string): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
    username: string,
    teamId?: pkg_proto_utilpb_uuid_pb.UUID.AsObject,
    password: string,
    role: Role,
    passwordHash: string,
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
  getUsersList(): Array<User>;
  setUsersList(value: Array<User>): GetAllResponse;
  clearUsersList(): GetAllResponse;
  addUsers(value?: User, index?: number): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllResponse): GetAllResponse.AsObject;
  static serializeBinaryToWriter(message: GetAllResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllResponse;
  static deserializeBinaryFromReader(message: GetAllResponse, reader: jspb.BinaryReader): GetAllResponse;
}

export namespace GetAllResponse {
  export type AsObject = {
    usersList: Array<User.AsObject>,
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
  getUser(): User | undefined;
  setUser(value?: User): GetByIDResponse;
  hasUser(): boolean;
  clearUser(): GetByIDResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetByIDResponse): GetByIDResponse.AsObject;
  static serializeBinaryToWriter(message: GetByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetByIDResponse;
  static deserializeBinaryFromReader(message: GetByIDResponse, reader: jspb.BinaryReader): GetByIDResponse;
}

export namespace GetByIDResponse {
  export type AsObject = {
    user?: User.AsObject,
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
  getUsersList(): Array<User>;
  setUsersList(value: Array<User>): StoreRequest;
  clearUsersList(): StoreRequest;
  addUsers(value?: User, index?: number): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StoreRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StoreRequest): StoreRequest.AsObject;
  static serializeBinaryToWriter(message: StoreRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StoreRequest;
  static deserializeBinaryFromReader(message: StoreRequest, reader: jspb.BinaryReader): StoreRequest;
}

export namespace StoreRequest {
  export type AsObject = {
    usersList: Array<User.AsObject>,
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
  getUser(): User | undefined;
  setUser(value?: User): UpdateRequest;
  hasUser(): boolean;
  clearUser(): UpdateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRequest): UpdateRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRequest;
  static deserializeBinaryFromReader(message: UpdateRequest, reader: jspb.BinaryReader): UpdateRequest;
}

export namespace UpdateRequest {
  export type AsObject = {
    user?: User.AsObject,
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

export class GetByUsernameRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): GetByUsernameRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetByUsernameRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetByUsernameRequest): GetByUsernameRequest.AsObject;
  static serializeBinaryToWriter(message: GetByUsernameRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetByUsernameRequest;
  static deserializeBinaryFromReader(message: GetByUsernameRequest, reader: jspb.BinaryReader): GetByUsernameRequest;
}

export namespace GetByUsernameRequest {
  export type AsObject = {
    username: string,
  }
}

export class GetByUsernameResponse extends jspb.Message {
  getUser(): User | undefined;
  setUser(value?: User): GetByUsernameResponse;
  hasUser(): boolean;
  clearUser(): GetByUsernameResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetByUsernameResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetByUsernameResponse): GetByUsernameResponse.AsObject;
  static serializeBinaryToWriter(message: GetByUsernameResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetByUsernameResponse;
  static deserializeBinaryFromReader(message: GetByUsernameResponse, reader: jspb.BinaryReader): GetByUsernameResponse;
}

export namespace GetByUsernameResponse {
  export type AsObject = {
    user?: User.AsObject,
  }
}

export enum Role { 
  ROLE_NOT_SET = 0,
  BLUE = 1,
  RED = 2,
  BLACK = 3,
}
