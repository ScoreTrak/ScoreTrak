import {GRPCClients} from "../../../grpc/gRPCClients";
import React from "react";

function numberRange (start: number, end: number) {
    if (start > end){
        const tmp = start
        start = end
        end = tmp
    }
    end = end + 1
    return new Array(end - start).fill(undefined).map((d, i) => i + start);
}

function parse_index(rng: string){
    const stripped = rng.replace(/\s+/g, '')
    const ranges = stripped.split(',');
    const ret: number[] = []

    for (let i = 0; i < ranges.length; i++){
        if (ranges[i].split('-').length === 1){
            const num = parseInt(ranges[i])
            if (!isNaN(num)){
                ret.push(num)
            } else {
                return []
            }
        }
        else if (ranges[i].split('-').length === 2){
            const start = parseInt(ranges[i].split('-')[0])
            const end = parseInt(ranges[i].split('-')[1])
            if (!isNaN(start) && !isNaN(end)){
                ret.push(...numberRange(start, end))
            } else{
                return []
            }
        }
        else {
            return []
        }
    }
    return Array.from(new Set(ret).values()).sort(function(a, b) {
        return a - b;
    })
}

type SetupProps = {
    gRPCClients: GRPCClients
    isDarkTheme: boolean
    setTitle: React.Dispatch<React.SetStateAction<string>>;
    genericEnqueue: Function
}


export {parse_index, numberRange};
export type { SetupProps };
