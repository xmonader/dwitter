import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../dwitter/params";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Tweet } from "../dwitter/tweet";
export declare const protobufPackage = "xmonader.dwitter.dwitter";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QueryTweetsRequest {
    pagination: PageRequest | undefined;
}
export interface QueryTweetsResponse {
    Tweet: Tweet[];
    pagination: PageResponse | undefined;
}
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
};
export declare const QueryTweetsRequest: {
    encode(message: QueryTweetsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryTweetsRequest;
    fromJSON(object: any): QueryTweetsRequest;
    toJSON(message: QueryTweetsRequest): unknown;
    fromPartial(object: DeepPartial<QueryTweetsRequest>): QueryTweetsRequest;
};
export declare const QueryTweetsResponse: {
    encode(message: QueryTweetsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryTweetsResponse;
    fromJSON(object: any): QueryTweetsResponse;
    toJSON(message: QueryTweetsResponse): unknown;
    fromPartial(object: DeepPartial<QueryTweetsResponse>): QueryTweetsResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a list of Tweets items. */
    Tweets(request: QueryTweetsRequest): Promise<QueryTweetsResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    Tweets(request: QueryTweetsRequest): Promise<QueryTweetsResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
