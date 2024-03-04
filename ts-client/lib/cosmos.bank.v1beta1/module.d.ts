import { DeliverTxResponse, StdFee } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { IgniteClient } from "../client";
import { Api } from "./rest";
import { MsgMultiSendResponse } from "./types/cosmos/bank/v1beta1/tx";
import { MsgUpdateParamsResponse } from "./types/cosmos/bank/v1beta1/tx";
import { Params } from "./types/cosmos/bank/v1beta1/bank";
import { QueryBalanceRequest } from "./types/cosmos/bank/v1beta1/query";
import { MsgSend } from "./types/cosmos/bank/v1beta1/tx";
import { QueryDenomOwnersResponse } from "./types/cosmos/bank/v1beta1/query";
import { QueryDenomMetadataByQueryStringRequest } from "./types/cosmos/bank/v1beta1/query";
import { QueryDenomOwnersByQueryResponse } from "./types/cosmos/bank/v1beta1/query";
import { GenesisState } from "./types/cosmos/bank/v1beta1/genesis";
import { MsgMultiSend } from "./types/cosmos/bank/v1beta1/tx";
import { QuerySpendableBalancesResponse } from "./types/cosmos/bank/v1beta1/query";
import { QuerySpendableBalanceByDenomRequest } from "./types/cosmos/bank/v1beta1/query";
import { QuerySendEnabledRequest } from "./types/cosmos/bank/v1beta1/query";
import { MsgSetSendEnabled } from "./types/cosmos/bank/v1beta1/tx";
import { Output } from "./types/cosmos/bank/v1beta1/bank";
import { QueryDenomOwnersRequest } from "./types/cosmos/bank/v1beta1/query";
import { QuerySupplyOfRequest } from "./types/cosmos/bank/v1beta1/query";
import { QueryDenomsMetadataResponse } from "./types/cosmos/bank/v1beta1/query";
import { MsgUpdateParams } from "./types/cosmos/bank/v1beta1/tx";
import { MsgSetSendEnabledResponse } from "./types/cosmos/bank/v1beta1/tx";
import { Supply } from "./types/cosmos/bank/v1beta1/bank";
import { SendAuthorization } from "./types/cosmos/bank/v1beta1/authz";
import { SendEnabled } from "./types/cosmos/bank/v1beta1/bank";
import { DenomUnit } from "./types/cosmos/bank/v1beta1/bank";
import { QueryBalanceResponse } from "./types/cosmos/bank/v1beta1/query";
import { QuerySpendableBalancesRequest } from "./types/cosmos/bank/v1beta1/query";
import { QuerySpendableBalanceByDenomResponse } from "./types/cosmos/bank/v1beta1/query";
import { QueryDenomMetadataResponse } from "./types/cosmos/bank/v1beta1/query";
import { DenomOwner } from "./types/cosmos/bank/v1beta1/query";
import { QueryParamsRequest } from "./types/cosmos/bank/v1beta1/query";
import { QueryDenomMetadataByQueryStringResponse } from "./types/cosmos/bank/v1beta1/query";
import { QuerySupplyOfResponse } from "./types/cosmos/bank/v1beta1/query";
import { Metadata } from "./types/cosmos/bank/v1beta1/bank";
import { MsgSendResponse } from "./types/cosmos/bank/v1beta1/tx";
import { QueryParamsResponse } from "./types/cosmos/bank/v1beta1/query";
import { QuerySendEnabledResponse } from "./types/cosmos/bank/v1beta1/query";
import { QueryDenomOwnersByQueryRequest } from "./types/cosmos/bank/v1beta1/query";
import { Input } from "./types/cosmos/bank/v1beta1/bank";
import { QueryAllBalancesResponse } from "./types/cosmos/bank/v1beta1/query";
import { QueryDenomsMetadataRequest } from "./types/cosmos/bank/v1beta1/query";
import { QueryDenomMetadataRequest } from "./types/cosmos/bank/v1beta1/query";
import { QueryAllBalancesRequest } from "./types/cosmos/bank/v1beta1/query";
import { QueryTotalSupplyRequest } from "./types/cosmos/bank/v1beta1/query";
import { QueryTotalSupplyResponse } from "./types/cosmos/bank/v1beta1/query";
import { Balance } from "./types/cosmos/bank/v1beta1/genesis";
export { MsgMultiSendResponse, MsgUpdateParamsResponse, Params, QueryBalanceRequest, MsgSend, QueryDenomOwnersResponse, QueryDenomMetadataByQueryStringRequest, QueryDenomOwnersByQueryResponse, GenesisState, MsgMultiSend, QuerySpendableBalancesResponse, QuerySpendableBalanceByDenomRequest, QuerySendEnabledRequest, MsgSetSendEnabled, Output, QueryDenomOwnersRequest, QuerySupplyOfRequest, QueryDenomsMetadataResponse, MsgUpdateParams, MsgSetSendEnabledResponse, Supply, SendAuthorization, SendEnabled, DenomUnit, QueryBalanceResponse, QuerySpendableBalancesRequest, QuerySpendableBalanceByDenomResponse, QueryDenomMetadataResponse, DenomOwner, QueryParamsRequest, QueryDenomMetadataByQueryStringResponse, QuerySupplyOfResponse, Metadata, MsgSendResponse, QueryParamsResponse, QuerySendEnabledResponse, QueryDenomOwnersByQueryRequest, Input, QueryAllBalancesResponse, QueryDenomsMetadataRequest, QueryDenomMetadataRequest, QueryAllBalancesRequest, QueryTotalSupplyRequest, QueryTotalSupplyResponse, Balance };
type sendMsgMultiSendResponseParams = {
    value: MsgMultiSendResponse;
    fee?: StdFee;
    memo?: string;
};
type sendMsgUpdateParamsResponseParams = {
    value: MsgUpdateParamsResponse;
    fee?: StdFee;
    memo?: string;
};
type sendParamsParams = {
    value: Params;
    fee?: StdFee;
    memo?: string;
};
type sendQueryBalanceRequestParams = {
    value: QueryBalanceRequest;
    fee?: StdFee;
    memo?: string;
};
type sendMsgSendParams = {
    value: MsgSend;
    fee?: StdFee;
    memo?: string;
};
type sendQueryDenomOwnersResponseParams = {
    value: QueryDenomOwnersResponse;
    fee?: StdFee;
    memo?: string;
};
type sendQueryDenomMetadataByQueryStringRequestParams = {
    value: QueryDenomMetadataByQueryStringRequest;
    fee?: StdFee;
    memo?: string;
};
type sendQueryDenomOwnersByQueryResponseParams = {
    value: QueryDenomOwnersByQueryResponse;
    fee?: StdFee;
    memo?: string;
};
type sendGenesisStateParams = {
    value: GenesisState;
    fee?: StdFee;
    memo?: string;
};
type sendMsgMultiSendParams = {
    value: MsgMultiSend;
    fee?: StdFee;
    memo?: string;
};
type sendQuerySpendableBalancesResponseParams = {
    value: QuerySpendableBalancesResponse;
    fee?: StdFee;
    memo?: string;
};
type sendQuerySpendableBalanceByDenomRequestParams = {
    value: QuerySpendableBalanceByDenomRequest;
    fee?: StdFee;
    memo?: string;
};
type sendQuerySendEnabledRequestParams = {
    value: QuerySendEnabledRequest;
    fee?: StdFee;
    memo?: string;
};
type sendMsgSetSendEnabledParams = {
    value: MsgSetSendEnabled;
    fee?: StdFee;
    memo?: string;
};
type sendOutputParams = {
    value: Output;
    fee?: StdFee;
    memo?: string;
};
type sendQueryDenomOwnersRequestParams = {
    value: QueryDenomOwnersRequest;
    fee?: StdFee;
    memo?: string;
};
type sendQuerySupplyOfRequestParams = {
    value: QuerySupplyOfRequest;
    fee?: StdFee;
    memo?: string;
};
type sendQueryDenomsMetadataResponseParams = {
    value: QueryDenomsMetadataResponse;
    fee?: StdFee;
    memo?: string;
};
type sendMsgUpdateParamsParams = {
    value: MsgUpdateParams;
    fee?: StdFee;
    memo?: string;
};
type sendMsgSetSendEnabledResponseParams = {
    value: MsgSetSendEnabledResponse;
    fee?: StdFee;
    memo?: string;
};
type sendSupplyParams = {
    value: Supply;
    fee?: StdFee;
    memo?: string;
};
type sendSendAuthorizationParams = {
    value: SendAuthorization;
    fee?: StdFee;
    memo?: string;
};
type sendSendEnabledParams = {
    value: SendEnabled;
    fee?: StdFee;
    memo?: string;
};
type sendDenomUnitParams = {
    value: DenomUnit;
    fee?: StdFee;
    memo?: string;
};
type sendQueryBalanceResponseParams = {
    value: QueryBalanceResponse;
    fee?: StdFee;
    memo?: string;
};
type sendQuerySpendableBalancesRequestParams = {
    value: QuerySpendableBalancesRequest;
    fee?: StdFee;
    memo?: string;
};
type sendQuerySpendableBalanceByDenomResponseParams = {
    value: QuerySpendableBalanceByDenomResponse;
    fee?: StdFee;
    memo?: string;
};
type sendQueryDenomMetadataResponseParams = {
    value: QueryDenomMetadataResponse;
    fee?: StdFee;
    memo?: string;
};
type sendDenomOwnerParams = {
    value: DenomOwner;
    fee?: StdFee;
    memo?: string;
};
type sendQueryParamsRequestParams = {
    value: QueryParamsRequest;
    fee?: StdFee;
    memo?: string;
};
type sendQueryDenomMetadataByQueryStringResponseParams = {
    value: QueryDenomMetadataByQueryStringResponse;
    fee?: StdFee;
    memo?: string;
};
type sendQuerySupplyOfResponseParams = {
    value: QuerySupplyOfResponse;
    fee?: StdFee;
    memo?: string;
};
type sendMetadataParams = {
    value: Metadata;
    fee?: StdFee;
    memo?: string;
};
type sendMsgSendResponseParams = {
    value: MsgSendResponse;
    fee?: StdFee;
    memo?: string;
};
type sendQueryParamsResponseParams = {
    value: QueryParamsResponse;
    fee?: StdFee;
    memo?: string;
};
type sendQuerySendEnabledResponseParams = {
    value: QuerySendEnabledResponse;
    fee?: StdFee;
    memo?: string;
};
type sendQueryDenomOwnersByQueryRequestParams = {
    value: QueryDenomOwnersByQueryRequest;
    fee?: StdFee;
    memo?: string;
};
type sendInputParams = {
    value: Input;
    fee?: StdFee;
    memo?: string;
};
type sendQueryAllBalancesResponseParams = {
    value: QueryAllBalancesResponse;
    fee?: StdFee;
    memo?: string;
};
type sendQueryDenomsMetadataRequestParams = {
    value: QueryDenomsMetadataRequest;
    fee?: StdFee;
    memo?: string;
};
type sendQueryDenomMetadataRequestParams = {
    value: QueryDenomMetadataRequest;
    fee?: StdFee;
    memo?: string;
};
type sendQueryAllBalancesRequestParams = {
    value: QueryAllBalancesRequest;
    fee?: StdFee;
    memo?: string;
};
type sendQueryTotalSupplyRequestParams = {
    value: QueryTotalSupplyRequest;
    fee?: StdFee;
    memo?: string;
};
type sendQueryTotalSupplyResponseParams = {
    value: QueryTotalSupplyResponse;
    fee?: StdFee;
    memo?: string;
};
type sendBalanceParams = {
    value: Balance;
    fee?: StdFee;
    memo?: string;
};
type msgMultiSendResponseParams = {
    value: MsgMultiSendResponse;
};
type msgUpdateParamsResponseParams = {
    value: MsgUpdateParamsResponse;
};
type paramsParams = {
    value: Params;
};
type queryBalanceRequestParams = {
    value: QueryBalanceRequest;
};
type msgSendParams = {
    value: MsgSend;
};
type queryDenomOwnersResponseParams = {
    value: QueryDenomOwnersResponse;
};
type queryDenomMetadataByQueryStringRequestParams = {
    value: QueryDenomMetadataByQueryStringRequest;
};
type queryDenomOwnersByQueryResponseParams = {
    value: QueryDenomOwnersByQueryResponse;
};
type genesisStateParams = {
    value: GenesisState;
};
type msgMultiSendParams = {
    value: MsgMultiSend;
};
type querySpendableBalancesResponseParams = {
    value: QuerySpendableBalancesResponse;
};
type querySpendableBalanceByDenomRequestParams = {
    value: QuerySpendableBalanceByDenomRequest;
};
type querySendEnabledRequestParams = {
    value: QuerySendEnabledRequest;
};
type msgSetSendEnabledParams = {
    value: MsgSetSendEnabled;
};
type outputParams = {
    value: Output;
};
type queryDenomOwnersRequestParams = {
    value: QueryDenomOwnersRequest;
};
type querySupplyOfRequestParams = {
    value: QuerySupplyOfRequest;
};
type queryDenomsMetadataResponseParams = {
    value: QueryDenomsMetadataResponse;
};
type msgUpdateParamsParams = {
    value: MsgUpdateParams;
};
type msgSetSendEnabledResponseParams = {
    value: MsgSetSendEnabledResponse;
};
type supplyParams = {
    value: Supply;
};
type sendAuthorizationParams = {
    value: SendAuthorization;
};
type sendEnabledParams = {
    value: SendEnabled;
};
type denomUnitParams = {
    value: DenomUnit;
};
type queryBalanceResponseParams = {
    value: QueryBalanceResponse;
};
type querySpendableBalancesRequestParams = {
    value: QuerySpendableBalancesRequest;
};
type querySpendableBalanceByDenomResponseParams = {
    value: QuerySpendableBalanceByDenomResponse;
};
type queryDenomMetadataResponseParams = {
    value: QueryDenomMetadataResponse;
};
type denomOwnerParams = {
    value: DenomOwner;
};
type queryParamsRequestParams = {
    value: QueryParamsRequest;
};
type queryDenomMetadataByQueryStringResponseParams = {
    value: QueryDenomMetadataByQueryStringResponse;
};
type querySupplyOfResponseParams = {
    value: QuerySupplyOfResponse;
};
type metadataParams = {
    value: Metadata;
};
type msgSendResponseParams = {
    value: MsgSendResponse;
};
type queryParamsResponseParams = {
    value: QueryParamsResponse;
};
type querySendEnabledResponseParams = {
    value: QuerySendEnabledResponse;
};
type queryDenomOwnersByQueryRequestParams = {
    value: QueryDenomOwnersByQueryRequest;
};
type inputParams = {
    value: Input;
};
type queryAllBalancesResponseParams = {
    value: QueryAllBalancesResponse;
};
type queryDenomsMetadataRequestParams = {
    value: QueryDenomsMetadataRequest;
};
type queryDenomMetadataRequestParams = {
    value: QueryDenomMetadataRequest;
};
type queryAllBalancesRequestParams = {
    value: QueryAllBalancesRequest;
};
type queryTotalSupplyRequestParams = {
    value: QueryTotalSupplyRequest;
};
type queryTotalSupplyResponseParams = {
    value: QueryTotalSupplyResponse;
};
type balanceParams = {
    value: Balance;
};
export declare const registry: Registry;
interface TxClientOptions {
    addr: string;
    prefix: string;
    signer?: OfflineSigner;
}
export declare const txClient: ({ signer, prefix, addr }?: TxClientOptions) => {
    sendMsgMultiSendResponse({ value, fee, memo }: sendMsgMultiSendResponseParams): Promise<DeliverTxResponse>;
    sendMsgUpdateParamsResponse({ value, fee, memo }: sendMsgUpdateParamsResponseParams): Promise<DeliverTxResponse>;
    sendParams({ value, fee, memo }: sendParamsParams): Promise<DeliverTxResponse>;
    sendQueryBalanceRequest({ value, fee, memo }: sendQueryBalanceRequestParams): Promise<DeliverTxResponse>;
    sendMsgSend({ value, fee, memo }: sendMsgSendParams): Promise<DeliverTxResponse>;
    sendQueryDenomOwnersResponse({ value, fee, memo }: sendQueryDenomOwnersResponseParams): Promise<DeliverTxResponse>;
    sendQueryDenomMetadataByQueryStringRequest({ value, fee, memo }: sendQueryDenomMetadataByQueryStringRequestParams): Promise<DeliverTxResponse>;
    sendQueryDenomOwnersByQueryResponse({ value, fee, memo }: sendQueryDenomOwnersByQueryResponseParams): Promise<DeliverTxResponse>;
    sendGenesisState({ value, fee, memo }: sendGenesisStateParams): Promise<DeliverTxResponse>;
    sendMsgMultiSend({ value, fee, memo }: sendMsgMultiSendParams): Promise<DeliverTxResponse>;
    sendQuerySpendableBalancesResponse({ value, fee, memo }: sendQuerySpendableBalancesResponseParams): Promise<DeliverTxResponse>;
    sendQuerySpendableBalanceByDenomRequest({ value, fee, memo }: sendQuerySpendableBalanceByDenomRequestParams): Promise<DeliverTxResponse>;
    sendQuerySendEnabledRequest({ value, fee, memo }: sendQuerySendEnabledRequestParams): Promise<DeliverTxResponse>;
    sendMsgSetSendEnabled({ value, fee, memo }: sendMsgSetSendEnabledParams): Promise<DeliverTxResponse>;
    sendOutput({ value, fee, memo }: sendOutputParams): Promise<DeliverTxResponse>;
    sendQueryDenomOwnersRequest({ value, fee, memo }: sendQueryDenomOwnersRequestParams): Promise<DeliverTxResponse>;
    sendQuerySupplyOfRequest({ value, fee, memo }: sendQuerySupplyOfRequestParams): Promise<DeliverTxResponse>;
    sendQueryDenomsMetadataResponse({ value, fee, memo }: sendQueryDenomsMetadataResponseParams): Promise<DeliverTxResponse>;
    sendMsgUpdateParams({ value, fee, memo }: sendMsgUpdateParamsParams): Promise<DeliverTxResponse>;
    sendMsgSetSendEnabledResponse({ value, fee, memo }: sendMsgSetSendEnabledResponseParams): Promise<DeliverTxResponse>;
    sendSupply({ value, fee, memo }: sendSupplyParams): Promise<DeliverTxResponse>;
    sendSendAuthorization({ value, fee, memo }: sendSendAuthorizationParams): Promise<DeliverTxResponse>;
    sendSendEnabled({ value, fee, memo }: sendSendEnabledParams): Promise<DeliverTxResponse>;
    sendDenomUnit({ value, fee, memo }: sendDenomUnitParams): Promise<DeliverTxResponse>;
    sendQueryBalanceResponse({ value, fee, memo }: sendQueryBalanceResponseParams): Promise<DeliverTxResponse>;
    sendQuerySpendableBalancesRequest({ value, fee, memo }: sendQuerySpendableBalancesRequestParams): Promise<DeliverTxResponse>;
    sendQuerySpendableBalanceByDenomResponse({ value, fee, memo }: sendQuerySpendableBalanceByDenomResponseParams): Promise<DeliverTxResponse>;
    sendQueryDenomMetadataResponse({ value, fee, memo }: sendQueryDenomMetadataResponseParams): Promise<DeliverTxResponse>;
    sendDenomOwner({ value, fee, memo }: sendDenomOwnerParams): Promise<DeliverTxResponse>;
    sendQueryParamsRequest({ value, fee, memo }: sendQueryParamsRequestParams): Promise<DeliverTxResponse>;
    sendQueryDenomMetadataByQueryStringResponse({ value, fee, memo }: sendQueryDenomMetadataByQueryStringResponseParams): Promise<DeliverTxResponse>;
    sendQuerySupplyOfResponse({ value, fee, memo }: sendQuerySupplyOfResponseParams): Promise<DeliverTxResponse>;
    sendMetadata({ value, fee, memo }: sendMetadataParams): Promise<DeliverTxResponse>;
    sendMsgSendResponse({ value, fee, memo }: sendMsgSendResponseParams): Promise<DeliverTxResponse>;
    sendQueryParamsResponse({ value, fee, memo }: sendQueryParamsResponseParams): Promise<DeliverTxResponse>;
    sendQuerySendEnabledResponse({ value, fee, memo }: sendQuerySendEnabledResponseParams): Promise<DeliverTxResponse>;
    sendQueryDenomOwnersByQueryRequest({ value, fee, memo }: sendQueryDenomOwnersByQueryRequestParams): Promise<DeliverTxResponse>;
    sendInput({ value, fee, memo }: sendInputParams): Promise<DeliverTxResponse>;
    sendQueryAllBalancesResponse({ value, fee, memo }: sendQueryAllBalancesResponseParams): Promise<DeliverTxResponse>;
    sendQueryDenomsMetadataRequest({ value, fee, memo }: sendQueryDenomsMetadataRequestParams): Promise<DeliverTxResponse>;
    sendQueryDenomMetadataRequest({ value, fee, memo }: sendQueryDenomMetadataRequestParams): Promise<DeliverTxResponse>;
    sendQueryAllBalancesRequest({ value, fee, memo }: sendQueryAllBalancesRequestParams): Promise<DeliverTxResponse>;
    sendQueryTotalSupplyRequest({ value, fee, memo }: sendQueryTotalSupplyRequestParams): Promise<DeliverTxResponse>;
    sendQueryTotalSupplyResponse({ value, fee, memo }: sendQueryTotalSupplyResponseParams): Promise<DeliverTxResponse>;
    sendBalance({ value, fee, memo }: sendBalanceParams): Promise<DeliverTxResponse>;
    msgMultiSendResponse({ value }: msgMultiSendResponseParams): EncodeObject;
    msgUpdateParamsResponse({ value }: msgUpdateParamsResponseParams): EncodeObject;
    params({ value }: paramsParams): EncodeObject;
    queryBalanceRequest({ value }: queryBalanceRequestParams): EncodeObject;
    msgSend({ value }: msgSendParams): EncodeObject;
    queryDenomOwnersResponse({ value }: queryDenomOwnersResponseParams): EncodeObject;
    queryDenomMetadataByQueryStringRequest({ value }: queryDenomMetadataByQueryStringRequestParams): EncodeObject;
    queryDenomOwnersByQueryResponse({ value }: queryDenomOwnersByQueryResponseParams): EncodeObject;
    genesisState({ value }: genesisStateParams): EncodeObject;
    msgMultiSend({ value }: msgMultiSendParams): EncodeObject;
    querySpendableBalancesResponse({ value }: querySpendableBalancesResponseParams): EncodeObject;
    querySpendableBalanceByDenomRequest({ value }: querySpendableBalanceByDenomRequestParams): EncodeObject;
    querySendEnabledRequest({ value }: querySendEnabledRequestParams): EncodeObject;
    msgSetSendEnabled({ value }: msgSetSendEnabledParams): EncodeObject;
    output({ value }: outputParams): EncodeObject;
    queryDenomOwnersRequest({ value }: queryDenomOwnersRequestParams): EncodeObject;
    querySupplyOfRequest({ value }: querySupplyOfRequestParams): EncodeObject;
    queryDenomsMetadataResponse({ value }: queryDenomsMetadataResponseParams): EncodeObject;
    msgUpdateParams({ value }: msgUpdateParamsParams): EncodeObject;
    msgSetSendEnabledResponse({ value }: msgSetSendEnabledResponseParams): EncodeObject;
    supply({ value }: supplyParams): EncodeObject;
    sendAuthorization({ value }: sendAuthorizationParams): EncodeObject;
    sendEnabled({ value }: sendEnabledParams): EncodeObject;
    denomUnit({ value }: denomUnitParams): EncodeObject;
    queryBalanceResponse({ value }: queryBalanceResponseParams): EncodeObject;
    querySpendableBalancesRequest({ value }: querySpendableBalancesRequestParams): EncodeObject;
    querySpendableBalanceByDenomResponse({ value }: querySpendableBalanceByDenomResponseParams): EncodeObject;
    queryDenomMetadataResponse({ value }: queryDenomMetadataResponseParams): EncodeObject;
    denomOwner({ value }: denomOwnerParams): EncodeObject;
    queryParamsRequest({ value }: queryParamsRequestParams): EncodeObject;
    queryDenomMetadataByQueryStringResponse({ value }: queryDenomMetadataByQueryStringResponseParams): EncodeObject;
    querySupplyOfResponse({ value }: querySupplyOfResponseParams): EncodeObject;
    metadata({ value }: metadataParams): EncodeObject;
    msgSendResponse({ value }: msgSendResponseParams): EncodeObject;
    queryParamsResponse({ value }: queryParamsResponseParams): EncodeObject;
    querySendEnabledResponse({ value }: querySendEnabledResponseParams): EncodeObject;
    queryDenomOwnersByQueryRequest({ value }: queryDenomOwnersByQueryRequestParams): EncodeObject;
    input({ value }: inputParams): EncodeObject;
    queryAllBalancesResponse({ value }: queryAllBalancesResponseParams): EncodeObject;
    queryDenomsMetadataRequest({ value }: queryDenomsMetadataRequestParams): EncodeObject;
    queryDenomMetadataRequest({ value }: queryDenomMetadataRequestParams): EncodeObject;
    queryAllBalancesRequest({ value }: queryAllBalancesRequestParams): EncodeObject;
    queryTotalSupplyRequest({ value }: queryTotalSupplyRequestParams): EncodeObject;
    queryTotalSupplyResponse({ value }: queryTotalSupplyResponseParams): EncodeObject;
    balance({ value }: balanceParams): EncodeObject;
};
interface QueryClientOptions {
    addr: string;
}
export declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Api<unknown>;
declare class SDKModule {
    query: ReturnType<typeof queryClient>;
    tx: ReturnType<typeof txClient>;
    structure: Record<string, unknown>;
    registry: Array<[string, GeneratedType]>;
    constructor(client: IgniteClient);
    updateTX(client: IgniteClient): void;
}
declare const IgntModule: (test: IgniteClient) => {
    module: {
        CosmosBankV1Beta1: SDKModule;
    };
    registry: [string, GeneratedType][];
};
export default IgntModule;
