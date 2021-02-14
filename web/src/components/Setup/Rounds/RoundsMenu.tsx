import React, {useEffect} from "react";
import {SetupProps} from "../util/util";
import Box from "@material-ui/core/Box";
import MaterialTable, {Column} from '@material-table/core'
import {Severity} from "../../../types/types";
import {CircularProgress} from "@material-ui/core";
import {GetAllRequest, Round} from "../../../grpc/pkg/round/roundpb/round_pb";

type roundColumns = {
    id: number
    start: Date | undefined
    finish: Date | undefined
    note: string
    err: string
}

function roundToRoundColumn(round: Round): roundColumns{
    return {
        id: round.getId(),
        start: round.getStart() ? new Date(round.getStart()?.getSeconds() as number * 1000) : undefined,
        finish: round.getFinish() ? new Date(round.getFinish()?.getSeconds() as number * 1000) : undefined,
        err: round.getErr(),
        note: round.getNote(),
    }
}

export default function RoundMenu(props: SetupProps) {
    const title = "Rounds"
    props.setTitle(title)
    const columns :Array<Column<roundColumns>>  =
        [
            { title: 'ID (optional)', field: 'id', editable: 'onAdd'},
            { title: 'Start Time', field: 'start', type: 'datetime' },
            { title: 'Note', field: 'note' },
            { title: 'Error', field: 'err'},
            { title: 'Finish Time', field: 'finish', type: 'datetime'},
        ]

    const [state, setState] = React.useState<{columns: any[], loader: boolean, data: roundColumns[]}>({
        columns,
        loader: true,
        data: []
    });

    function reloadSetter() {
        props.gRPCClients.roundClient.getAll(new GetAllRequest(), {}).then(roundsResponse => {
            setState(prevState => {return{...prevState, data: roundsResponse.getRoundsList().map((round): roundColumns => {
                    return roundToRoundColumn(round)}), loader: false}})}, (err: any) => {
            props.genericEnqueue(`Encountered an error while retrieving Rounds: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
    }
    useEffect(() => {
        reloadSetter()
    }, []);

    return (
        <React.Fragment>
            {!state.loader ?
                <Box height="100%" width="100%" >
                    <MaterialTable
                        title={title}
                        columns={state.columns}
                        data={state.data}
                        options={{pageSizeOptions: [5, 10, 20, 50, 100, 500, 1000], pageSize: 20, emptyRowsWhenPaging: false}}
                    />
                </Box>
                :
                <Box height="100%" width="100%" m="auto">
                    <CircularProgress  />
                </Box>
            }
        </React.Fragment>
    );
}