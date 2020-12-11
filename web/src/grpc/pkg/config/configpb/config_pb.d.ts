import * as jspb from 'google-protobuf'

import * as google_protobuf_wrappers_pb from 'google-protobuf/google/protobuf/wrappers_pb';


export class DynamicConfig extends jspb.Message {
  getRoundDuration(): number;
  setRoundDuration(value: number): DynamicConfig;

  getEnabled(): google_protobuf_wrappers_pb.BoolValue | undefined;
  setEnabled(value?: google_protobuf_wrappers_pb.BoolValue): DynamicConfig;
  hasEnabled(): boolean;
  clearEnabled(): DynamicConfig;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DynamicConfig.AsObject;
  static toObject(includeInstance: boolean, msg: DynamicConfig): DynamicConfig.AsObject;
  static serializeBinaryToWriter(message: DynamicConfig, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DynamicConfig;
  static deserializeBinaryFromReader(message: DynamicConfig, reader: jspb.BinaryReader): DynamicConfig;
}

export namespace DynamicConfig {
  export type AsObject = {
    roundDuration: number,
    enabled?: google_protobuf_wrappers_pb.BoolValue.AsObject,
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
  getDynamicConfig(): DynamicConfig | undefined;
  setDynamicConfig(value?: DynamicConfig): GetResponse;
  hasDynamicConfig(): boolean;
  clearDynamicConfig(): GetResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetResponse): GetResponse.AsObject;
  static serializeBinaryToWriter(message: GetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetResponse;
  static deserializeBinaryFromReader(message: GetResponse, reader: jspb.BinaryReader): GetResponse;
}

export namespace GetResponse {
  export type AsObject = {
    dynamicConfig?: DynamicConfig.AsObject,
  }
}

export class UpdateRequest extends jspb.Message {
  getDynamicConfig(): DynamicConfig | undefined;
  setDynamicConfig(value?: DynamicConfig): UpdateRequest;
  hasDynamicConfig(): boolean;
  clearDynamicConfig(): UpdateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRequest): UpdateRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRequest;
  static deserializeBinaryFromReader(message: UpdateRequest, reader: jspb.BinaryReader): UpdateRequest;
}

export namespace UpdateRequest {
  export type AsObject = {
    dynamicConfig?: DynamicConfig.AsObject,
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

export class GetStaticConfigRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetStaticConfigRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetStaticConfigRequest): GetStaticConfigRequest.AsObject;
  static serializeBinaryToWriter(message: GetStaticConfigRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetStaticConfigRequest;
  static deserializeBinaryFromReader(message: GetStaticConfigRequest, reader: jspb.BinaryReader): GetStaticConfigRequest;
}

export namespace GetStaticConfigRequest {
  export type AsObject = {
  }
}

export class GetStaticConfigResponse extends jspb.Message {
  getStaticConfig(): string;
  setStaticConfig(value: string): GetStaticConfigResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetStaticConfigResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetStaticConfigResponse): GetStaticConfigResponse.AsObject;
  static serializeBinaryToWriter(message: GetStaticConfigResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetStaticConfigResponse;
  static deserializeBinaryFromReader(message: GetStaticConfigResponse, reader: jspb.BinaryReader): GetStaticConfigResponse;
}

export namespace GetStaticConfigResponse {
  export type AsObject = {
    staticConfig: string,
  }
}

