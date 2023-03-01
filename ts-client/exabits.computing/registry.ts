import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateDevices } from "./types/exabits/computing/tx";
import { MsgDeleteDevices } from "./types/exabits/computing/tx";
import { MsgCreateDevices } from "./types/exabits/computing/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/exabits.computing.MsgUpdateDevices", MsgUpdateDevices],
    ["/exabits.computing.MsgDeleteDevices", MsgDeleteDevices],
    ["/exabits.computing.MsgCreateDevices", MsgCreateDevices],
    
];

export { msgTypes }