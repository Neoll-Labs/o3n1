/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "etherstate.etherstate";

export interface EthereumAddressState {
  index: string;
  state: number;
  block: number;
  nonce: number;
}

function createBaseEthereumAddressState(): EthereumAddressState {
  return { index: "", state: 0, block: 0, nonce: 0 };
}

export const EthereumAddressState = {
  encode(message: EthereumAddressState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.state !== 0) {
      writer.uint32(16).uint64(message.state);
    }
    if (message.block !== 0) {
      writer.uint32(24).uint64(message.block);
    }
    if (message.nonce !== 0) {
      writer.uint32(32).uint64(message.nonce);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EthereumAddressState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEthereumAddressState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.state = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.block = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.nonce = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EthereumAddressState {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      state: isSet(object.state) ? Number(object.state) : 0,
      block: isSet(object.block) ? Number(object.block) : 0,
      nonce: isSet(object.nonce) ? Number(object.nonce) : 0,
    };
  },

  toJSON(message: EthereumAddressState): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.state !== undefined && (obj.state = Math.round(message.state));
    message.block !== undefined && (obj.block = Math.round(message.block));
    message.nonce !== undefined && (obj.nonce = Math.round(message.nonce));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EthereumAddressState>, I>>(object: I): EthereumAddressState {
    const message = createBaseEthereumAddressState();
    message.index = object.index ?? "";
    message.state = object.state ?? 0;
    message.block = object.block ?? 0;
    message.nonce = object.nonce ?? 0;
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
