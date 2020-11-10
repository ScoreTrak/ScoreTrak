import * as jspb from 'google-protobuf'

import * as pkg_check_checkpb_check_pb from '../../../pkg/check/checkpb/check_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class Round extends jspb.Message {
  getId(): number;
  setId(value: number): Round;

  getStart(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setStart(value?: google_protobuf_timestamp_pb.Timestamp): Round;
  hasStart(): boolean;
  clearStart(): Round;

  getNote(): string;
  setNote(value: string): Round;

  getErr(): string;
  setErr(value: string): Round;

  getFinish(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setFinish(value?: google_protobuf_timestamp_pb.Timestamp): Round;
  hasFinish(): boolean;
  clearFinish(): Round;

  getChecksList(): Array<pkg_check_checkpb_check_pb.Check>;
  setChecksList(value: Array<pkg_check_checkpb_check_pb.Check>): Round;
  clearChecksList(): Round;
  addChecks(value?: pkg_check_checkpb_check_pb.Check, index?: number): pkg_check_checkpb_check_pb.Check;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Round.AsObject;
  static toObject(includeInstance: boolean, msg: Round): Round.AsObject;
  static serializeBinaryToWriter(message: Round, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Round;
  static deserializeBinaryFromReader(message: Round, reader: jspb.BinaryReader): Round;
}

export namespace Round {
  export type AsObject = {
    id: number,
    start?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    note: string,
    err: string,
    finish?: google_protobuf_timestamp_pb.Timestamp.AsObject,
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
  getRoundsList(): Array<Round>;
  setRoundsList(value: Array<Round>): GetAllResponse;
  clearRoundsList(): GetAllResponse;
  addRounds(value?: Round, index?: number): Round;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllResponse): GetAllResponse.AsObject;
  static serializeBinaryToWriter(message: GetAllResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllResponse;
  static deserializeBinaryFromReader(message: GetAllResponse, reader: jspb.BinaryReader): GetAllResponse;
}

export namespace GetAllResponse {
  export type AsObject = {
    roundsList: Array<Round.AsObject>,
  }
}

export class GetByIDRequest extends jspb.Message {
  getId(): number;
  setId(value: number): GetByIDRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetByIDRequest): GetByIDRequest.AsObject;
  static serializeBinaryToWriter(message: GetByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetByIDRequest;
  static deserializeBinaryFromReader(message: GetByIDRequest, reader: jspb.BinaryReader): GetByIDRequest;
}

export namespace GetByIDRequest {
  export type AsObject = {
    id: number,
  }
}

export class GetByIDResponse extends jspb.Message {
  getRound(): Round | undefined;
  setRound(value?: Round): GetByIDResponse;
  hasRound(): boolean;
  clearRound(): GetByIDResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetByIDResponse): GetByIDResponse.AsObject;
  static serializeBinaryToWriter(message: GetByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetByIDResponse;
  static deserializeBinaryFromReader(message: GetByIDResponse, reader: jspb.BinaryReader): GetByIDResponse;
}

export namespace GetByIDResponse {
  export type AsObject = {
    round?: Round.AsObject,
  }
}

export class GetLastRoundRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetLastRoundRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetLastRoundRequest): GetLastRoundRequest.AsObject;
  static serializeBinaryToWriter(message: GetLastRoundRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetLastRoundRequest;
  static deserializeBinaryFromReader(message: GetLastRoundRequest, reader: jspb.BinaryReader): GetLastRoundRequest;
}

export namespace GetLastRoundRequest {
  export type AsObject = {
  }
}

export class GetLastRoundResponse extends jspb.Message {
  getRound(): Round | undefined;
  setRound(value?: Round): GetLastRoundResponse;
  hasRound(): boolean;
  clearRound(): GetLastRoundResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetLastRoundResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetLastRoundResponse): GetLastRoundResponse.AsObject;
  static serializeBinaryToWriter(message: GetLastRoundResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetLastRoundResponse;
  static deserializeBinaryFromReader(message: GetLastRoundResponse, reader: jspb.BinaryReader): GetLastRoundResponse;
}

export namespace GetLastRoundResponse {
  export type AsObject = {
    round?: Round.AsObject,
  }
}

export class GetLastNonElapsingRoundRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetLastNonElapsingRoundRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetLastNonElapsingRoundRequest): GetLastNonElapsingRoundRequest.AsObject;
  static serializeBinaryToWriter(message: GetLastNonElapsingRoundRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetLastNonElapsingRoundRequest;
  static deserializeBinaryFromReader(message: GetLastNonElapsingRoundRequest, reader: jspb.BinaryReader): GetLastNonElapsingRoundRequest;
}

export namespace GetLastNonElapsingRoundRequest {
  export type AsObject = {
  }
}

export class GetLastNonElapsingRoundResponse extends jspb.Message {
  getRound(): Round | undefined;
  setRound(value?: Round): GetLastNonElapsingRoundResponse;
  hasRound(): boolean;
  clearRound(): GetLastNonElapsingRoundResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetLastNonElapsingRoundResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetLastNonElapsingRoundResponse): GetLastNonElapsingRoundResponse.AsObject;
  static serializeBinaryToWriter(message: GetLastNonElapsingRoundResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetLastNonElapsingRoundResponse;
  static deserializeBinaryFromReader(message: GetLastNonElapsingRoundResponse, reader: jspb.BinaryReader): GetLastNonElapsingRoundResponse;
}

export namespace GetLastNonElapsingRoundResponse {
  export type AsObject = {
    round?: Round.AsObject,
  }
}

