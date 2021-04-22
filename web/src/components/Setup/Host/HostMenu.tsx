import React, {useEffect} from "react";
import Box from "@material-ui/core/Box";
import Stepper from "@material-ui/core/Stepper";
import Step from "@material-ui/core/Step";
import StepButton from "@material-ui/core/StepButton";
import {SetupProps} from "../util/util";
import HostCreate from "./HostCreate";
import {UUID} from "../../../grpc/pkg/proto/utilpb/uuid_pb";
import {BoolValue, StringValue} from "google-protobuf/google/protobuf/wrappers_pb";
import {
    DeleteRequest,
    GetAllRequest,
    Host,
    StoreRequest, UpdateRequest
} from "../../../grpc/pkg/host/hostpb/host_pb";
import {Severity} from "../../../types/types";
import MaterialTable, {Column} from '@material-table/core'
import {CircularProgress} from "@material-ui/core";
import {GetAllRequest as GetAllRequestHostGroup} from "../../../grpc/pkg/host_group/host_grouppb/host_group_pb";
import {GetAllRequest as GetAllRequestTeam} from "../../../grpc/pkg/team/teampb/team_pb";


function getSteps() {
    return ['Regular View', 'Quick Create'];
}

export default function HostMenu(props: SetupProps) {

    const [activeStep, setActiveStep] = React.useState<number>(0);
    const steps = getSteps();
    const handleStep = (step: number) => () => {
        setActiveStep(step);
    };
    return (
        <Box height="100%" width="100%" >
            <Stepper nonLinear activeStep={activeStep}>
                {steps.map((label, index) => (
                    <Step key={label}>
                        <StepButton onClick={handleStep(index)}>
                            {label}
                        </StepButton>
                    </Step>
                ))}
            </Stepper>
            <div>
                {
                    activeStep === 0 && <HostMenuTable {...props} />
                }
                {
                    activeStep === 1 && <HostCreate {...props} />
                }
            </div>
        </Box>
    );
}



export type hostColumns = {
    id: string | undefined
    pause: boolean | undefined
    hide: boolean | undefined
    editHost: boolean | undefined
    address: string
    addressListRange: string | undefined
    hostGroupId: string | undefined
    teamId: string | undefined
}


export function defaultHostColumns(): hostColumns {
    return {
        address: "", addressListRange: "", editHost: false, pause: false, hide: false, hostGroupId: undefined, id: undefined, teamId: undefined
    }
}



export function hostToHostColumn(host: Host): hostColumns{
    return {
        id: host.getId()?.getValue(),
        pause: host.getPause()?.getValue(),
        hide: host.getHide()?.getValue(),
        address: host.getAddress(),
        editHost: host.getEditHost()?.getValue(),
        hostGroupId: host.getHostGroupId()?.getValue(),
        teamId: host.getTeamId()?.getValue(),
        addressListRange: host.getAddressListRange()?.getValue()
    }
}

export function hostColumnsToHost(hostC: hostColumns): Host{
    const u = new Host()
    if (hostC.id && hostC.id !== "") u.setId((new UUID().setValue(hostC.id)))
    if (hostC.hostGroupId && hostC.hostGroupId !== "") u.setHostGroupId((new UUID().setValue(hostC.hostGroupId)))
    if (hostC.teamId && hostC.teamId !== "") u.setTeamId((new UUID().setValue(hostC.teamId)))
    if (hostC.hide !== undefined) u.setHide(new BoolValue().setValue(hostC.hide))
    if (hostC.pause !== undefined) u.setPause(new BoolValue().setValue(hostC.pause))
    if (hostC.editHost !== undefined) u.setEditHost(new BoolValue().setValue(hostC.editHost))
    if (hostC.addressListRange !== undefined) u.setAddressListRange(new StringValue().setValue(hostC.addressListRange))
    u.setAddress(hostC.address)
    return u
}

function HostMenuTable(props: SetupProps) {
    const title =  "Hosts"
    props.setTitle(title)
    const columns:Array<Column<hostColumns>> =
        [
            { title: 'ID (optional)', field: 'id', editable: 'onAdd'},
            { title: 'Address', field: 'address' },
            { title: 'Host Group ID', field: 'hostGroupId' },
            { title: 'Team ID', field: 'teamId' },
            { title: 'Hide from Scoreboard', field: 'hide', type: 'boolean', initialEditValue: false},
            { title: 'Pause Scoring', field: 'pause', type: 'boolean', initialEditValue: false},
            { title: 'Edit Host(Allow users to change Addresses)', field: 'editHost', type: 'boolean' },
            { title: "Address Range(comma separated list of allowed CIDR ranges and hostnames)", field: 'addressListRange'}
        ]

    const [state, setState] = React.useState<{columns: any[], loaderTeam: boolean, loaderHost: boolean, loaderHostGroup: boolean, data: hostColumns[]}>({
        columns,
        loaderTeam: true,
        loaderHost: true,
        loaderHostGroup: true,
        data: []
    });

    function reloadSetter() {

        props.gRPCClients.hostGroupClient.getAll(new GetAllRequestHostGroup(), {}).then(hostsGroupResponse => {
            const lookup: Record<string, string> = {}
            for (let i = 0; i < hostsGroupResponse.getHostGroupsList().length; i++){
                lookup[hostsGroupResponse.getHostGroupsList()[i].getId()?.getValue() as string] = `${hostsGroupResponse.getHostGroupsList()[i].getName()} (ID:${hostsGroupResponse.getHostGroupsList()[i].getId()?.getValue() as string})`
            }
            setState(prevState => {
                const columns = prevState.columns
                for (let i = 0; i < columns.length; i++){
                    if (columns[i].field === "hostGroupId"){
                        columns[i].lookup = lookup
                    }
                }
                return{...prevState, columns, loaderHostGroup: false
                }})
        }, (err: any) => {
            props.genericEnqueue(`Encountered an error while retrieving parent Host Groups: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })

        props.gRPCClients.teamClient.getAll(new GetAllRequestTeam(), {}).then(teamResponse => {
            const lookup: Record<string, string> = {}
            for (let i = 0; i < teamResponse.getTeamsList().length; i++){
                lookup[teamResponse.getTeamsList()[i].getId()?.getValue() as string] = `${teamResponse.getTeamsList()[i].getName()} (ID:${teamResponse.getTeamsList()[i].getId()?.getValue() as string})`
            }
            setState(prevState => {
                const columns = prevState.columns
                for (let i = 0; i < columns.length; i++){
                    if (columns[i].field === "teamId"){
                        columns[i].lookup = lookup
                    }
                }
                return{...prevState, columns, loaderTeam: false}})
        }, (err: any) => {
            props.genericEnqueue(`Encountered an error while retrieving parent Teams: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
        props.gRPCClients.hostClient.getAll(new GetAllRequest(), {}).then(hostsResponse => {
            setState(prevState => {return{...prevState, data: hostsResponse.getHostsList().map((host): hostColumns => {
                    return hostToHostColumn(host)}), loader: false, loaderHost: false}})}, (err: any) => {
            props.genericEnqueue(`Encountered an error while retrieving Hosts: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
    }
    useEffect(() => {
        reloadSetter()
    }, []);

    return (
        <React.Fragment>
            {!state.loaderHost && !state.loaderTeam && !state.loaderHostGroup ?
                <Box height="100%" width="100%" >
                    <MaterialTable
                        title={title}
                        columns={state.columns}
                        data={state.data}
                        options={{pageSizeOptions: [5, 10, 20, 50, 100, 500, 1000], pageSize: 20, emptyRowsWhenPaging: false}}
                        editable={{
                            onRowAdd: (newData) =>
                                new Promise((resolve, reject) => {
                                    setTimeout(() => {
                                        const storeRequest = new StoreRequest()
                                        const u = hostColumnsToHost(newData)
                                        storeRequest.addHosts(u, 0)
                                        props.gRPCClients.hostClient.store(storeRequest, {}).then(result => {
                                            u.setId(result.getIdsList()[0])
                                            setState((prevState) => {
                                                const data = [...prevState.data];
                                                data.push(hostToHostColumn(u));
                                                return { ...prevState, data };
                                            });
                                            resolve();
                                        }, (err: any) => {
                                            props.genericEnqueue(`Unable to store Host: ${err.message}. Error code: ${err.code}`, Severity.Error)
                                            reject()
                                        })
                                    }, 600);
                                }),
                            onRowUpdate: (newData, oldData) =>
                                new Promise((resolve, reject) => {
                                    setTimeout(() => {
                                        if (oldData){
                                            const updateRequest = new UpdateRequest()
                                            const u = hostColumnsToHost(newData)
                                            updateRequest.setHost(u)
                                            props.gRPCClients.hostClient.update(updateRequest, {}).then(result => {
                                                setState((prevState) => {
                                                    const data = [...prevState.data];
                                                    data[data.indexOf(oldData)] = newData;
                                                    return { ...prevState, data };
                                                });
                                                resolve();
                                            }, (err: any) => {
                                                props.genericEnqueue(`Unable to update hosts: ${err.message}. Error code: ${err.code}`, Severity.Error)
                                                reject()
                                            })
                                        }
                                    }, 600);
                                }),
                            onRowDelete: (oldData) =>
                                new Promise((resolve, reject) => {
                                    setTimeout(() => {
                                        const deleteRequest = new DeleteRequest()
                                        deleteRequest.setId((new UUID().setValue(oldData.id as string)))
                                        props.gRPCClients.hostClient.delete(deleteRequest, {}).then(result => {
                                            setState((prevState) => {
                                                const data = [...prevState.data];
                                                data.splice(data.indexOf(oldData), 1);
                                                return { ...prevState, data };
                                            });
                                            resolve();
                                        }, (err: any) => {
                                            props.genericEnqueue(`Unable to delete host: ${err.message}. Error code: ${err.code}`, Severity.Error)
                                            reject()
                                        })
                                    }, 600);
                                }),
                        }}
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