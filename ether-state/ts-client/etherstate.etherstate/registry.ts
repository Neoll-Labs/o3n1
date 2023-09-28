import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgDisableEthAddress } from "./types/etherstate/etherstate/tx";
import { MsgEnableEthAddress } from "./types/etherstate/etherstate/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/etherstate.etherstate.MsgDisableEthAddress", MsgDisableEthAddress],
    ["/etherstate.etherstate.MsgEnableEthAddress", MsgEnableEthAddress],
    
];

export { msgTypes }