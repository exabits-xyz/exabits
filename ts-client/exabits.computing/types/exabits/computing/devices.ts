/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "exabits.computing";

export interface Devices {
  id: number;
  address: string;
  machineId: string;
  machineToken: string;
  coords: string;
  ip: string;
  creator: string;
}

function createBaseDevices(): Devices {
  return { id: 0, address: "", machineId: "", machineToken: "", coords: "", ip: "", creator: "" };
}

export const Devices = {
  encode(message: Devices, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
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
    if (message.creator !== "") {
      writer.uint32(58).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Devices {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDevices();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
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
        case 7:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Devices {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      address: isSet(object.address) ? String(object.address) : "",
      machineId: isSet(object.machineId) ? String(object.machineId) : "",
      machineToken: isSet(object.machineToken) ? String(object.machineToken) : "",
      coords: isSet(object.coords) ? String(object.coords) : "",
      ip: isSet(object.ip) ? String(object.ip) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
    };
  },

  toJSON(message: Devices): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.address !== undefined && (obj.address = message.address);
    message.machineId !== undefined && (obj.machineId = message.machineId);
    message.machineToken !== undefined && (obj.machineToken = message.machineToken);
    message.coords !== undefined && (obj.coords = message.coords);
    message.ip !== undefined && (obj.ip = message.ip);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Devices>, I>>(object: I): Devices {
    const message = createBaseDevices();
    message.id = object.id ?? 0;
    message.address = object.address ?? "";
    message.machineId = object.machineId ?? "";
    message.machineToken = object.machineToken ?? "";
    message.coords = object.coords ?? "";
    message.ip = object.ip ?? "";
    message.creator = object.creator ?? "";
    return message;
  },
};

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
