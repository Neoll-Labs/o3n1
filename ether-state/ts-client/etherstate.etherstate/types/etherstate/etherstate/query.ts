/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { EthereumAddress } from "./ethereum_address";
import { EthereumAddressState } from "./ethereum_address_state";
import { Params } from "./params";

export const protobufPackage = "etherstate.etherstate";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetEthereumAddressRequest {
  index: string;
}

export interface QueryGetEthereumAddressResponse {
  ethereumAddress: EthereumAddress | undefined;
}

export interface QueryAllEthereumAddressRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllEthereumAddressResponse {
  ethereumAddress: EthereumAddress[];
  pagination: PageResponse | undefined;
}

export interface QueryGetEthereumAddressStateRequest {
  index: string;
}

export interface QueryGetEthereumAddressStateResponse {
  ethereumAddressState: EthereumAddressState | undefined;
}

export interface QueryAllEthereumAddressStateRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllEthereumAddressStateResponse {
  ethereumAddressState: EthereumAddressState[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
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

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetEthereumAddressRequest(): QueryGetEthereumAddressRequest {
  return { index: "" };
}

export const QueryGetEthereumAddressRequest = {
  encode(message: QueryGetEthereumAddressRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetEthereumAddressRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetEthereumAddressRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEthereumAddressRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetEthereumAddressRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetEthereumAddressRequest>, I>>(
    object: I,
  ): QueryGetEthereumAddressRequest {
    const message = createBaseQueryGetEthereumAddressRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetEthereumAddressResponse(): QueryGetEthereumAddressResponse {
  return { ethereumAddress: undefined };
}

export const QueryGetEthereumAddressResponse = {
  encode(message: QueryGetEthereumAddressResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ethereumAddress !== undefined) {
      EthereumAddress.encode(message.ethereumAddress, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetEthereumAddressResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetEthereumAddressResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ethereumAddress = EthereumAddress.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEthereumAddressResponse {
    return {
      ethereumAddress: isSet(object.ethereumAddress) ? EthereumAddress.fromJSON(object.ethereumAddress) : undefined,
    };
  },

  toJSON(message: QueryGetEthereumAddressResponse): unknown {
    const obj: any = {};
    message.ethereumAddress !== undefined
      && (obj.ethereumAddress = message.ethereumAddress ? EthereumAddress.toJSON(message.ethereumAddress) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetEthereumAddressResponse>, I>>(
    object: I,
  ): QueryGetEthereumAddressResponse {
    const message = createBaseQueryGetEthereumAddressResponse();
    message.ethereumAddress = (object.ethereumAddress !== undefined && object.ethereumAddress !== null)
      ? EthereumAddress.fromPartial(object.ethereumAddress)
      : undefined;
    return message;
  },
};

function createBaseQueryAllEthereumAddressRequest(): QueryAllEthereumAddressRequest {
  return { pagination: undefined };
}

export const QueryAllEthereumAddressRequest = {
  encode(message: QueryAllEthereumAddressRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllEthereumAddressRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllEthereumAddressRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllEthereumAddressRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllEthereumAddressRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllEthereumAddressRequest>, I>>(
    object: I,
  ): QueryAllEthereumAddressRequest {
    const message = createBaseQueryAllEthereumAddressRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllEthereumAddressResponse(): QueryAllEthereumAddressResponse {
  return { ethereumAddress: [], pagination: undefined };
}

export const QueryAllEthereumAddressResponse = {
  encode(message: QueryAllEthereumAddressResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.ethereumAddress) {
      EthereumAddress.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllEthereumAddressResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllEthereumAddressResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ethereumAddress.push(EthereumAddress.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllEthereumAddressResponse {
    return {
      ethereumAddress: Array.isArray(object?.ethereumAddress)
        ? object.ethereumAddress.map((e: any) => EthereumAddress.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllEthereumAddressResponse): unknown {
    const obj: any = {};
    if (message.ethereumAddress) {
      obj.ethereumAddress = message.ethereumAddress.map((e) => e ? EthereumAddress.toJSON(e) : undefined);
    } else {
      obj.ethereumAddress = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllEthereumAddressResponse>, I>>(
    object: I,
  ): QueryAllEthereumAddressResponse {
    const message = createBaseQueryAllEthereumAddressResponse();
    message.ethereumAddress = object.ethereumAddress?.map((e) => EthereumAddress.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetEthereumAddressStateRequest(): QueryGetEthereumAddressStateRequest {
  return { index: "" };
}

export const QueryGetEthereumAddressStateRequest = {
  encode(message: QueryGetEthereumAddressStateRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetEthereumAddressStateRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetEthereumAddressStateRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEthereumAddressStateRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetEthereumAddressStateRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetEthereumAddressStateRequest>, I>>(
    object: I,
  ): QueryGetEthereumAddressStateRequest {
    const message = createBaseQueryGetEthereumAddressStateRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetEthereumAddressStateResponse(): QueryGetEthereumAddressStateResponse {
  return { ethereumAddressState: undefined };
}

export const QueryGetEthereumAddressStateResponse = {
  encode(message: QueryGetEthereumAddressStateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ethereumAddressState !== undefined) {
      EthereumAddressState.encode(message.ethereumAddressState, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetEthereumAddressStateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetEthereumAddressStateResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ethereumAddressState = EthereumAddressState.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEthereumAddressStateResponse {
    return {
      ethereumAddressState: isSet(object.ethereumAddressState)
        ? EthereumAddressState.fromJSON(object.ethereumAddressState)
        : undefined,
    };
  },

  toJSON(message: QueryGetEthereumAddressStateResponse): unknown {
    const obj: any = {};
    message.ethereumAddressState !== undefined && (obj.ethereumAddressState = message.ethereumAddressState
      ? EthereumAddressState.toJSON(message.ethereumAddressState)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetEthereumAddressStateResponse>, I>>(
    object: I,
  ): QueryGetEthereumAddressStateResponse {
    const message = createBaseQueryGetEthereumAddressStateResponse();
    message.ethereumAddressState = (object.ethereumAddressState !== undefined && object.ethereumAddressState !== null)
      ? EthereumAddressState.fromPartial(object.ethereumAddressState)
      : undefined;
    return message;
  },
};

function createBaseQueryAllEthereumAddressStateRequest(): QueryAllEthereumAddressStateRequest {
  return { pagination: undefined };
}

export const QueryAllEthereumAddressStateRequest = {
  encode(message: QueryAllEthereumAddressStateRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllEthereumAddressStateRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllEthereumAddressStateRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllEthereumAddressStateRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllEthereumAddressStateRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllEthereumAddressStateRequest>, I>>(
    object: I,
  ): QueryAllEthereumAddressStateRequest {
    const message = createBaseQueryAllEthereumAddressStateRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllEthereumAddressStateResponse(): QueryAllEthereumAddressStateResponse {
  return { ethereumAddressState: [], pagination: undefined };
}

export const QueryAllEthereumAddressStateResponse = {
  encode(message: QueryAllEthereumAddressStateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.ethereumAddressState) {
      EthereumAddressState.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllEthereumAddressStateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllEthereumAddressStateResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ethereumAddressState.push(EthereumAddressState.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllEthereumAddressStateResponse {
    return {
      ethereumAddressState: Array.isArray(object?.ethereumAddressState)
        ? object.ethereumAddressState.map((e: any) => EthereumAddressState.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllEthereumAddressStateResponse): unknown {
    const obj: any = {};
    if (message.ethereumAddressState) {
      obj.ethereumAddressState = message.ethereumAddressState.map((e) =>
        e ? EthereumAddressState.toJSON(e) : undefined
      );
    } else {
      obj.ethereumAddressState = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllEthereumAddressStateResponse>, I>>(
    object: I,
  ): QueryAllEthereumAddressStateResponse {
    const message = createBaseQueryAllEthereumAddressStateResponse();
    message.ethereumAddressState = object.ethereumAddressState?.map((e) => EthereumAddressState.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of EthereumAddress items. */
  EthereumAddress(request: QueryGetEthereumAddressRequest): Promise<QueryGetEthereumAddressResponse>;
  EthereumAddressAll(request: QueryAllEthereumAddressRequest): Promise<QueryAllEthereumAddressResponse>;
  /** Queries a list of EthereumAddressState items. */
  EthereumAddressState(request: QueryGetEthereumAddressStateRequest): Promise<QueryGetEthereumAddressStateResponse>;
  EthereumAddressStateAll(request: QueryAllEthereumAddressStateRequest): Promise<QueryAllEthereumAddressStateResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.EthereumAddress = this.EthereumAddress.bind(this);
    this.EthereumAddressAll = this.EthereumAddressAll.bind(this);
    this.EthereumAddressState = this.EthereumAddressState.bind(this);
    this.EthereumAddressStateAll = this.EthereumAddressStateAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("etherstate.etherstate.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  EthereumAddress(request: QueryGetEthereumAddressRequest): Promise<QueryGetEthereumAddressResponse> {
    const data = QueryGetEthereumAddressRequest.encode(request).finish();
    const promise = this.rpc.request("etherstate.etherstate.Query", "EthereumAddress", data);
    return promise.then((data) => QueryGetEthereumAddressResponse.decode(new _m0.Reader(data)));
  }

  EthereumAddressAll(request: QueryAllEthereumAddressRequest): Promise<QueryAllEthereumAddressResponse> {
    const data = QueryAllEthereumAddressRequest.encode(request).finish();
    const promise = this.rpc.request("etherstate.etherstate.Query", "EthereumAddressAll", data);
    return promise.then((data) => QueryAllEthereumAddressResponse.decode(new _m0.Reader(data)));
  }

  EthereumAddressState(request: QueryGetEthereumAddressStateRequest): Promise<QueryGetEthereumAddressStateResponse> {
    const data = QueryGetEthereumAddressStateRequest.encode(request).finish();
    const promise = this.rpc.request("etherstate.etherstate.Query", "EthereumAddressState", data);
    return promise.then((data) => QueryGetEthereumAddressStateResponse.decode(new _m0.Reader(data)));
  }

  EthereumAddressStateAll(request: QueryAllEthereumAddressStateRequest): Promise<QueryAllEthereumAddressStateResponse> {
    const data = QueryAllEthereumAddressStateRequest.encode(request).finish();
    const promise = this.rpc.request("etherstate.etherstate.Query", "EthereumAddressStateAll", data);
    return promise.then((data) => QueryAllEthereumAddressStateResponse.decode(new _m0.Reader(data)));
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
