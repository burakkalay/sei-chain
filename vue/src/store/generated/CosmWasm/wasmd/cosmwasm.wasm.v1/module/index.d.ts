import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateAdmin } from "./types/cosmwasm/wasm/v1/tx";
import { MsgExecuteContract } from "./types/cosmwasm/wasm/v1/tx";
import { MsgMigrateContract } from "./types/cosmwasm/wasm/v1/tx";
import { MsgStoreCode } from "./types/cosmwasm/wasm/v1/tx";
import { MsgInstantiateContract } from "./types/cosmwasm/wasm/v1/tx";
import { MsgIBCCloseChannel } from "./types/cosmwasm/wasm/v1/ibc";
import { MsgIBCSend } from "./types/cosmwasm/wasm/v1/ibc";
import { MsgClearAdmin } from "./types/cosmwasm/wasm/v1/tx";
export declare const MissingWalletError: Error;
export declare const registry: Registry;
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }?: SignAndBroadcastOptions) => any;
    msgUpdateAdmin: (data: MsgUpdateAdmin) => EncodeObject;
    msgExecuteContract: (data: MsgExecuteContract) => EncodeObject;
    msgMigrateContract: (data: MsgMigrateContract) => EncodeObject;
    msgStoreCode: (data: MsgStoreCode) => EncodeObject;
    msgInstantiateContract: (data: MsgInstantiateContract) => EncodeObject;
    msgIBCCloseChannel: (data: MsgIBCCloseChannel) => EncodeObject;
    msgIBCSend: (data: MsgIBCSend) => EncodeObject;
    msgClearAdmin: (data: MsgClearAdmin) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
