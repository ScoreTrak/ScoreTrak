import React, {useEffect} from 'react';
import {forwardRef} from 'react';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import TextField from "@material-ui/core/TextField";
import Box from "@material-ui/core/Box";
import CircularProgress from "@material-ui/core/CircularProgress";
import Button from "@material-ui/core/Button";
import {SetupProps} from "../util/util";
import {Team} from "../../../grpc/pkg/team/teampb/team_pb";
import {HostGroup} from "../../../grpc/pkg/host_group/host_grouppb/host_group_pb";


import {GetAllRequest as GetAllRequestHostGroup} from "../../../grpc/pkg/host_group/host_grouppb/host_group_pb";
import {GetAllRequest as GetAllRequestTeam} from "../../../grpc/pkg/team/teampb/team_pb";
import {Severity} from "../../../types/types";
import {StoreRequest} from "../../../grpc/pkg/host/hostpb/host_pb";
import {hostColumnsToHost} from "./HostMenu";



const HostCreate = forwardRef((props: SetupProps, ref) => {

    const [dt, setData] = React.useState<{loaderTeam: boolean, loaderHostGroup: boolean, teams: Team[], hostGroups: HostGroup[]}>({loaderTeam: true, loaderHostGroup: true, teams:[], hostGroups:[]})

    useEffect(() => {
        props.gRPCClients.teamClient.getAll(new GetAllRequestTeam(), {}).then(respTeam =>{
            setData(prevState => {return {...prevState, loaderTeam: false, teams: respTeam.getTeamsList().sort((a,b) => {
                    const aidx = a.getIndex()?.getValue()
                    const bidx = b.getIndex()?.getValue()
                    if (!aidx){
                        return -1
                    } else if (!bidx){
                        return 1
                    } else {
                        return (aidx > bidx) ? 1 : -1
                    }
                })}})
        }, (err: any) => {
            props.genericEnqueue(`Encountered an error while retrieving Teams: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })

        props.gRPCClients.hostGroupClient.getAll(new GetAllRequestHostGroup(), {}).then(respHostGroup =>{
            setData(prevState => {return {...prevState, hostGroups: respHostGroup.getHostGroupsList(), loaderHostGroup: false}})
        }, (err: any) => {
            props.genericEnqueue(`Encountered an error while retrieving Host Groups: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
    }, []);

    const [rowsData, setRowData] = React.useState<Record<string, string>>({});

    function modifyTeamRows(hostGroupId: string, templateValue: string){
        let nextRowData:Record<string, string> = {}
        if (templateValue.includes('X')){
            for (let i = 0; i < dt.teams.length; i++){
                if (dt.teams[i].getIndex()?.getValue()) {
                    nextRowData[`${dt.teams[i].getId()?.getValue()}_${hostGroupId}`] = templateValue.replace("X", (dt.teams[i].getIndex()?.getValue() as number).toString())
                }
            }

            setRowData(prevState => {return{...prevState, ...nextRowData}})
        }
    }
    function submit() {
        const storeRequest = new StoreRequest()
        Object.keys(rowsData).forEach(teamHostGrpId => {
            let teamId, hostGroupId
            [teamId, hostGroupId] = teamHostGrpId.split("_")
            if (rowsData[teamHostGrpId]){
                storeRequest.addHosts(hostColumnsToHost({
                    address: rowsData[teamHostGrpId],
                    editHost: false,
                    enabled: true,
                    teamId: teamId,
                    hostGroupId: hostGroupId,
                    id: undefined
                }))
            }
        })
        props.gRPCClients.hostClient.store(storeRequest, {}).then(r => {
            props.genericEnqueue("Success!", Severity.Success, 3000)
        }, (err:any) =>{
            props.genericEnqueue(`Encountered an error while Storing Hosts: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
    }

    return (
        <React.Fragment>
            <div>
                {!dt.loaderTeam && !dt.loaderHostGroup ?
                    <Table stickyHeader aria-label="sticky table">
                        <TableHead>
                            <TableRow>
                                <TableCell />
                                {dt.hostGroups.map((column) => (
                                    <TableCell>
                                        {column.getName()}
                                    </TableCell>
                                ))}
                            </TableRow>
                        </TableHead>
                        <TableHead>
                            <TableRow>
                                <TableCell />
                                {dt.hostGroups.map((column) => (
                                    <TableCell>
                                        <TextField label="Template" id={`id_${column.getId()?.getValue()}`} helperText="Ex. 10.1.X.1" onChange={event => {modifyTeamRows(column.getId()?.getValue() as string, event.target.value)}}/>
                                    </TableCell>
                                ))}
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {dt.teams.map((row) => {
                                if (row.getIndex()?.getValue()){
                                    return (
                                        <TableRow hover role="checkbox" tabIndex={-1}>
                                            <TableCell key={row.getName()}>
                                                {row.getName()}
                                            </TableCell>

                                            {dt.hostGroups.map((column) => {
                                                return (
                                                    <TableCell>
                                                        <TextField id={`${row.getId()?.getValue()}_${column.getId()?.getValue()}`} value={rowsData[`${row.getId()?.getValue()}_${column.getId()?.getValue()}`]} onChange={(event => {
                                                            const val = event.target.value
                                                            setRowData(prevState => {
                                                                return {...prevState, [`${row.getId()?.getValue()}_${column.getId()?.getValue()}`]: val}
                                                            })
                                                        })}
                                                        />
                                                    </TableCell>
                                                );
                                            })}
                                        </TableRow>
                                    );
                                } else {
                                    return
                                }
                            })}
                        </TableBody>
                    </Table>
                    :
                    <Box height="100%" width="100%" m="auto">
                        <CircularProgress  />
                    </Box>
                }
            </div>
            <div style={{display: 'flex',  justifyContent: 'flex-end'}}>
                <Button onClick={submit} variant="contained" style={{ marginRight: '8px', marginTop: '8px'}}>Submit</Button>
            </div>
        </React.Fragment>
    );
})

export default HostCreate;