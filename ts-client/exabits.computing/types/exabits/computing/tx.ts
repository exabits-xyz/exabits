/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "exabits.computing";

export interface MsgCreateDevices {
  creator: string;
  address: string;
  machineId: string;
  machineToken: string;
  coords: string;
  ip: string;
}

export interface MsgCreateDevicesResponse {
  id: number;
}

export interface MsgUpdateDevices {
  creator: string;
  id: number;
  address: string;
  machineId: string;
  machineToken: string;
  coords: string;
  ip: string;
}

export interface MsgUpdateDevicesResponse {
}

export interface MsgDeleteDevices {
  creator: string;
  id: number;
}

export interface MsgDeleteDevicesResponse {
}

function createBaseMsgCreateDevices(): MsgCreateDevices {
  return { creator: "", address: "", machineId: "", machineToken: "", coords: "", ip: "" };
}

export const MsgCreateDevices = {
  encode(message: MsgCreateDevices, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.machineId !== "") {
      writer.uint32(26).string(message.machineId);
    }
    if (message.machineToken !== "") {
      writer.uint32(34).string(message.machineToken);
    }
    if (message.coords !== "") {
      writer.uint32(42).string(message.coords);
    }
    if (message.ip !== "") {
      writer.uint32(50).string(message.ip);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateDevices {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateDevices();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.machineId = reader.string();
          break;
        case 4:
          message.machineToken = reader.string();
          break;
        case 5:
          message.coords = reader.string();
          break;
        case 6:
          message.ip = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateDevices {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      address: isSet(object.address) ? String(object.address) : "",
      machineId: isSet(object.machineId) ? String(object.machineId) : "",
      machineToken: isSet(object.machineToken) ? String(object.machineToken) : "",
      coords: isSet(object.coords) ? String(object.coords) : "",
      ip: isSet(object.ip) ? String(object.ip) : "",
    };
  },

  toJSON(message: MsgCreateDevices): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    message.machineId !== undefined && (obj.machineId = message.machineId);
    message.machineToken !== undefined && (obj.machineToken = message.machineToken);
    message.coords !== undefined && (obj.coords = message.coords);
    message.ip !== undefined && (obj.ip = message.ip);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateDevices>, I>>(object: I): MsgCreateDevices {
    const message = createBaseMsgCreateDevices();
    message.creator = object.creator ?? "";
    message.address = object.address ?? "";
    message.machineId = object.machineId ?? "";
    message.machineToken = object.machineToken ?? "";
    message.coords = object.coords ?? "";
    message.ip = object.ip ?? "";
    return message;
  },
};

function createBaseMsgCreateDevicesResponse(): MsgCreateDevicesResponse {
  return { id: 0 };
}

export const MsgCreateDevicesResponse = {
  encode(message: MsgCreateDevicesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateDevicesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateDevicesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateDevicesResponse {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: MsgCreateDevicesResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateDevicesResponse>, I>>(object: I): MsgCreateDevicesResponse {
    const message = createBaseMsgCreateDevicesResponse();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgUpdateDevices(): MsgUpdateDevices {
  return { creator: "", id: 0, address: "", machineId: "", machineToken: "", coords: "", ip: "" };
}

export const MsgUpdateDevices = {
  encode(message: MsgUpdateDevices, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    if (message.machineId !== "") {
      writer.uint32(34).string(message.machineId);
    }
    if (message.machineToken !== "") {
      writer.uint32(42).string(message.machineToken);
    }
    if (message.coords !== "") {
      writer.uint32(50).string(message.coords);
    }
    if (message.ip !== "") {
      writer.uint32(58).string(message.ip);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateDevices {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateDevices();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.address = reader.string();
          break;
        case 4:
          message.machineId = reader.string();
          break;
        case 5:
          message.machineToken = reader.string();
          break;
        case 6:
          message.coords = reader.string();
          break;
        case 7:
          message.ip = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateDevices {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
      address: isSet(object.address) ? String(object.address) : "",
      machineId: isSet(object.machineId) ? String(object.machineId) : "",
      machineToken: isSet(object.machineToken) ? String(object.machineToken) : "",
      coords: isSet(object.coords) ? String(object.coords) : "",
      ip: isSet(object.ip) ? String(object.ip) : "",
    };
  },

  toJSON(message: MsgUpdateDevices): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.address !== undefined && (obj.address = message.address);
    message.machineId !== undefined && (obj.machineId = message.machineId);
    message.machineToken !== undefined && (obj.machineToken = message.machineToken);
    message.coords !== undefined && (obj.coords = message.coords);
    message.ip !== undefined && (obj.ip = message.ip);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateDevices>, I>>(object: I): MsgUpdateDevices {
    const message = createBaseMsgUpdateDevices();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    message.address = object.address ?? "";
    message.machineId = object.machineId ?? "";
    message.machineToken = object.machineToken ?? "";
    message.coords = object.coords ?? "";
    message.ip = object.ip ?? "";
    return message;
  },
};

function createBaseMsgUpdateDevicesResponse(): MsgUpdateDevicesResponse {
  return {};
}

export const MsgUpdateDevicesResponse = {
  encode(_: MsgUpdateDevicesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateDevicesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateDevicesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateDevicesResponse {
    return {};
  },

  toJSON(_: MsgUpdateDevicesResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateDevicesResponse>, I>>(_: I): MsgUpdateDevicesResponse {
    const message = createBaseMsgUpdateDevicesResponse();
    return message;
  },
};

function createBaseMsgDeleteDevices(): MsgDeleteDevices {
  return { creator: "", id: 0 };
}

export const MsgDeleteDevices = {
  encode(message: MsgDeleteDevices, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteDevices {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteDevices();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteDevices {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgDeleteDevices): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteDevices>, I>>(object: I): MsgDeleteDevices {
    const message = createBaseMsgDeleteDevices();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgDeleteDevicesResponse(): MsgDeleteDevicesResponse {
  return {};
}

export const MsgDeleteDevicesResponse = {
  encode(_: MsgDeleteDevicesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteDevicesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteDevicesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteDevicesResponse {
    return {};
  },

  toJSON(_: MsgDeleteDevicesResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteDevicesResponse>, I>>(_: I): MsgDeleteDevicesResponse {
    const message = createBaseMsgDeleteDevicesResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateDevices(request: MsgCreateDevices): Promise<MsgCreateDevicesResponse>;
  UpdateDevices(request: MsgUpdateDevices): Promise<MsgUpdateDevicesResponse>;
  DeleteDevices(request: MsgDeleteDevices): Promise<MsgDeleteDevicesResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateDevices = this.CreateDevices.bind(this);
    this.UpdateDevices = this.UpdateDevices.bind(this);
    this.DeleteDevices = this.DeleteDevices.bind(this);
  }
  CreateDevices(request: MsgCreateDevices): Promise<MsgCreateDevicesResponse> {
    const data = MsgCreateDevices.encode(request).finish();
    const promise = this.rpc.request("exabits.computing.Msg", "CreateDevices", data);
    return promise.then((data) => MsgCreateDevicesResponse.decode(new _m0.Reader(data)));
  }

  UpdateDevices(request: MsgUpdateDevices): Promise<MsgUpdateDevicesResponse> {
    const data = MsgUpdateDevices.encode(request).finish();
    const promise = this.rpc.request("exabits.computing.Msg", "UpdateDevices", data);
    return promise.then((data) => MsgUpdateDevicesResponse.decode(new _m0.Reader(data)));
  }

  DeleteDevices(request: MsgDeleteDevices): Promise<MsgDeleteDevicesResponse> {
    const data = MsgDeleteDevices.encode(request).finish();
    const promise = this.rpc.request("exabits.computing.Msg", "DeleteDevices", data);
    return promise.then((data) => MsgDeleteDevicesResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
