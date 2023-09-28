/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "etherstate.etherstate";

export interface MsgEnableEthAddress {
  creator: string;
  address: string;
}

export interface MsgEnableEthAddressResponse {
  success: boolean;
}

export interface MsgDisableEthAddress {
  creator: string;
  address: string;
}

export interface MsgDisableEthAddressResponse {
  success: boolean;
}

function createBaseMsgEnableEthAddress(): MsgEnableEthAddress {
  return { creator: "", address: "" };
}

export const MsgEnableEthAddress = {
  encode(message: MsgEnableEthAddress, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgEnableEthAddress {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgEnableEthAddress();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEnableEthAddress {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      address: isSet(object.address) ? String(object.address) : "",
    };
  },

  toJSON(message: MsgEnableEthAddress): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgEnableEthAddress>, I>>(object: I): MsgEnableEthAddress {
    const message = createBaseMsgEnableEthAddress();
    message.creator = object.creator ?? "";
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseMsgEnableEthAddressResponse(): MsgEnableEthAddressResponse {
  return { success: false };
}

export const MsgEnableEthAddressResponse = {
  encode(message: MsgEnableEthAddressResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgEnableEthAddressResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgEnableEthAddressResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEnableEthAddressResponse {
    return { success: isSet(object.success) ? Boolean(object.success) : false };
  },

  toJSON(message: MsgEnableEthAddressResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgEnableEthAddressResponse>, I>>(object: I): MsgEnableEthAddressResponse {
    const message = createBaseMsgEnableEthAddressResponse();
    message.success = object.success ?? false;
    return message;
  },
};

function createBaseMsgDisableEthAddress(): MsgDisableEthAddress {
  return { creator: "", address: "" };
}

export const MsgDisableEthAddress = {
  encode(message: MsgDisableEthAddress, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDisableEthAddress {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDisableEthAddress();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDisableEthAddress {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      address: isSet(object.address) ? String(object.address) : "",
    };
  },

  toJSON(message: MsgDisableEthAddress): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDisableEthAddress>, I>>(object: I): MsgDisableEthAddress {
    const message = createBaseMsgDisableEthAddress();
    message.creator = object.creator ?? "";
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseMsgDisableEthAddressResponse(): MsgDisableEthAddressResponse {
  return { success: false };
}

export const MsgDisableEthAddressResponse = {
  encode(message: MsgDisableEthAddressResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDisableEthAddressResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDisableEthAddressResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDisableEthAddressResponse {
    return { success: isSet(object.success) ? Boolean(object.success) : false };
  },

  toJSON(message: MsgDisableEthAddressResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDisableEthAddressResponse>, I>>(object: I): MsgDisableEthAddressResponse {
    const message = createBaseMsgDisableEthAddressResponse();
    message.success = object.success ?? false;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  EnableEthAddress(request: MsgEnableEthAddress): Promise<MsgEnableEthAddressResponse>;
  DisableEthAddress(request: MsgDisableEthAddress): Promise<MsgDisableEthAddressResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.EnableEthAddress = this.EnableEthAddress.bind(this);
    this.DisableEthAddress = this.DisableEthAddress.bind(this);
  }
  EnableEthAddress(request: MsgEnableEthAddress): Promise<MsgEnableEthAddressResponse> {
    const data = MsgEnableEthAddress.encode(request).finish();
    const promise = this.rpc.request("etherstate.etherstate.Msg", "EnableEthAddress", data);
    return promise.then((data) => MsgEnableEthAddressResponse.decode(new _m0.Reader(data)));
  }

  DisableEthAddress(request: MsgDisableEthAddress): Promise<MsgDisableEthAddressResponse> {
    const data = MsgDisableEthAddress.encode(request).finish();
    const promise = this.rpc.request("etherstate.etherstate.Msg", "DisableEthAddress", data);
    return promise.then((data) => MsgDisableEthAddressResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
