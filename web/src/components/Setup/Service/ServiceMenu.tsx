import React, {useEffect} from "react";
import Box from "@material-ui/core/Box";
import Stepper from "@material-ui/core/Stepper";
import Step from "@material-ui/core/Step";
import StepButton from "@material-ui/core/StepButton";
import {SetupProps} from "../util/util";
import {Severity} from "../../../types/types";
import MaterialTable, {Column} from '@material-table/core'
import {UUID} from "../../../grpc/pkg/proto/utilpb/uuid_pb";
import {CircularProgress} from "@material-ui/core";
import ServiceCreate from "./ServiceCreate";
import {GetAllRequest as GetAllRequestHost} from "../../../grpc/pkg/host/hostpb/host_pb";
import {
    DeleteRequest,
    GetAllRequest,
    Service,
    StoreRequest,
    TestServiceRequest,
    UpdateRequest
} from "../../../grpc/pkg/service/servicepb/service_pb";
import {BoolValue, UInt64Value} from "google-protobuf/google/protobuf/wrappers_pb";
import {
    GetAllRequest as GetAllRequestServiceGroup,
} from "../../../grpc/pkg/service_group/service_grouppb/service_group_pb";


function getSteps() {
    return ['Regular View', 'Quick Create'];
}




export default function ServiceMenu(props: SetupProps) {
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
                    activeStep === 0 &&
                    <ServiceMenuTable {...props}/>
                }
                {
                    activeStep === 1 &&
                    <ServiceCreate {...props} />
                }
            </div>
        </Box>
    );
}

export type serviceColumns = {
    id: string | undefined,
    name: string,
    displayName: string,
    roundUnits: number,
    roundDelay: number | undefined,
    pointsBoost: number | undefined
    serviceGroupId: string | undefined
    hostId: string | undefined
    weight: number | undefined
    pause: boolean | undefined
    hide: boolean | undefined
}

export function serviceToServiceColumn(service: Service): serviceColumns{
    return {
        displayName: service.getDisplayName(),
        pause: service.getPause()?.getValue(),
        hide: service.getHide()?.getValue(),
        hostId: service.getHostId()?.getValue(),
        id: service.getId()?.getValue(),
        name: service.getName(),
        pointsBoost: service.getPointsBoost()?.getValue(),
        roundDelay: service.getRoundDelay()?.getValue(),
        roundUnits: service.getRoundUnits(),
        serviceGroupId: service.getServiceGroupId()?.getValue(),
        weight: service.getWeight()?.getValue()
    }
}

export function serviceColumnsToService(serviceC: serviceColumns): Service{
    const u = new Service()
    if (serviceC.id && serviceC.id !== "") u.setId((new UUID().setValue(serviceC.id)))
    u.setDisplayName(serviceC.displayName)
    if (serviceC.hostId && serviceC.hostId !== "") u.setHostId((new UUID().setValue(serviceC.hostId)))
    if (serviceC.serviceGroupId && serviceC.serviceGroupId !== "") u.setServiceGroupId((new UUID().setValue(serviceC.serviceGroupId)))
    if (serviceC.pause !== undefined) u.setPause(new BoolValue().setValue(serviceC.pause))
    if (serviceC.hide !== undefined) u.setHide(new BoolValue().setValue(serviceC.hide))
    u.setName(serviceC.name)
    if (serviceC.weight !== undefined) u.setWeight(new UInt64Value().setValue(serviceC.weight))
    u.setRoundUnits(serviceC.roundUnits)
    if (serviceC.roundDelay !== undefined) u.setRoundDelay(new UInt64Value().setValue(serviceC.roundDelay))
    if (serviceC.pointsBoost !== undefined) u.setPointsBoost(new UInt64Value().setValue(serviceC.pointsBoost))
    return u
}

function ServiceMenuTable(props: SetupProps) {
    const title = "Services"
    props.setTitle(title)
    const columns : Array<Column<serviceColumns>> =
        [
            { title: 'ID (optional)', field: 'id', editable: 'onAdd'},
            { title: 'Name', field: "name", lookup: {
                    'PING': 'PING', 'DNS': 'DNS', 'FTP': 'FTP', 'LDAP': 'LDAP',
                    'HTTP': 'HTTP', 'IMAP': 'IMAP', 'SMB': 'SMB', 'SSH': 'SSH',
                    'WINRM': 'WINRM', "SQL": "SQL", "CalDav": "CalDav"
                }},
            { title: 'Display Name(Columns on Status page)', field: 'displayName' },
            { title: 'Weight(Points per successful check)', field: 'weight', type: 'numeric', },
            { title: 'Points Boost', field: 'pointsBoost', type: 'numeric', initialEditValue: 0},
            { title: 'Hide from Scoreboard', field: 'hide', type: 'boolean', initialEditValue: false},
            { title: 'Pause Scoring', field: 'pause', type: 'boolean', initialEditValue: false},
            { title: 'Service Group ID', field: 'serviceGroupId' },
            { title: 'Host ID', field: 'hostId' },
            { title: 'Round Units(Frequency)', field: 'roundUnits', type: 'numeric', initialEditValue: 1},
            { title: 'Round Delay(Shift in frequency)', field: 'roundDelay', type: 'numeric', initialEditValue: 0 },
        ]
    const [state, setState] = React.useState<{columns: any[], loaderServiceGroup: boolean, loaderService: boolean, loaderHost: boolean, data: serviceColumns[]}>({
        columns,
        loaderService: true,
        loaderServiceGroup: true,
        loaderHost: true,
        data: []
    });

    function reloadSetter() {

        props.gRPCClients.hostClient.getAll(new GetAllRequestHost(), {}).then(hostsResponse => {
            const lookup: Record<string, string> = {}
            for (let i = 0; i < hostsResponse.getHostsList().length; i++){
                lookup[hostsResponse.getHostsList()[i].getId()?.getValue() as string] = `${hostsResponse.getHostsList()[i].getAddress()} (ID:${hostsResponse.getHostsList()[i].getId()?.getValue() as string})`
            }
            setState(prevState => {
                const columns = prevState.columns
                for (let i = 0; i < columns.length; i++){
                    if (columns[i].field === "hostId"){
                        columns[i].lookup = lookup
                    }
                }
                return{...prevState, columns, loaderHost: false
                }})
        }, (err: any) => {
            props.genericEnqueue(`Encountered an error while retrieving parent Teams: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })

        props.gRPCClients.serviceGroupClient.getAll(new GetAllRequestServiceGroup(), {}).then(serviceGroupRequest => {
            const lookup: Record<string, string> = {}
            for (let i = 0; i < serviceGroupRequest.getServiceGroupsList().length; i++){
                lookup[serviceGroupRequest.getServiceGroupsList()[i].getId()?.getValue() as string] = `${serviceGroupRequest.getServiceGroupsList()[i].getName()} (ID:${serviceGroupRequest.getServiceGroupsList()[i].getId()?.getValue() as string})`
            }
            setState(prevState => {
                const columns = prevState.columns
                for (let i = 0; i < columns.length; i++){
                    if (columns[i].field === "serviceGroupId"){
                        columns[i].lookup = lookup
                    }
                }
                return{...prevState, columns, loaderServiceGroup: false}})
        }, (err: any) => {
            props.genericEnqueue(`Encountered an error while retrieving parent Teams: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
        props.gRPCClients.serviceClient.getAll(new GetAllRequest(), {}).then(servicesResponse => {
            setState(prevState => {return{...prevState, data: servicesResponse.getServicesList().map((service): serviceColumns => {
                    return serviceToServiceColumn(service)}), loaderService: false}})}, (err: any) => {
            props.genericEnqueue(`Encountered an error while retrieving Services: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
    }
    useEffect(() => {
        reloadSetter()
    }, []);

    return (
        <React.Fragment>
            {!state.loaderHost && !state.loaderService && !state.loaderServiceGroup  ?
                <Box height="100%" width="100%" >
                    <MaterialTable
                        title={title}
                        actions={[
                            {icon: "flash_on", tooltip: 'test service', onClick: (event, rowData) => {
                                return props.gRPCClients.serviceClient.testService(new TestServiceRequest().setId(new UUID().setValue((rowData as serviceColumns).id as string)), {}).then((response) => { // ToDo: Implement Deadline
                                    if (response.getCheck()?.getPassed()?.getValue()){
                                        props.genericEnqueue(`Check Passed. Log: ${response.getCheck()?.getLog()}.`, Severity.Success)
                                    } else {
                                        props.genericEnqueue(`Check Failed. Log: ${response.getCheck()?.getLog()}. Err: ${response.getCheck()?.getErr()}.`, Severity.Warning)
                                    }
                                }, (err: any) => {
                                    props.genericEnqueue(`Failed to dispatch a check: ${err.message}. Error code: ${err.code}`, Severity.Error)
                                })
                            }}
                        ]}
                        columns={state.columns}
                        data={state.data}
                        options={{pageSizeOptions: [5, 10, 20, 50, 100, 500, 1000], pageSize: 20, emptyRowsWhenPaging: false}}
                        editable={{
                            onRowAdd: (newData) =>
                                new Promise((resolve, reject) => {
                                    setTimeout(() => {
                                        const storeRequest = new StoreRequest()
                                        const u = serviceColumnsToService(newData)

                                        storeRequest.addServices(u, 0)
                                        props.gRPCClients.serviceClient.store(storeRequest, {}).then(result => {
                                            u.setId(result.getIdsList()[0])
                                            setState((prevState) => {
                                                const data = [...prevState.data];
                                                data.push(serviceToServiceColumn(u));
                                                return { ...prevState, data };
                                            });
                                            resolve();
                                        }, (err: any) => {
                                            props.genericEnqueue(`Unable to store service: ${err.message}. Error code: ${err.code}`, Severity.Error)
                                            reject()
                                        })
                                    }, 600);
                                }),
                            onRowUpdate: (newData, oldData) =>
                                new Promise((resolve, reject) => {
                                    setTimeout(() => {
                                        if (oldData){
                                            const updateRequest = new UpdateRequest()
                                            const u = serviceColumnsToService(newData)
                                            updateRequest.setService(u)
                                            props.gRPCClients.serviceClient.update(updateRequest, {}).then(result => {
                                                setState((prevState) => {
                                                    const data = [...prevState.data];
                                                    data[data.indexOf(oldData)] = newData;
                                                    return { ...prevState, data };
                                                });
                                                resolve();
                                            }, (err: any) => {
                                                props.genericEnqueue(`Unable to update service: ${err.message}. Error code: ${err.code}`, Severity.Error)
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
                                        props.gRPCClients.serviceClient.delete(deleteRequest, {}).then(result => {
                                            setState((prevState) => {
                                                const data = [...prevState.data];
                                                data.splice(data.indexOf(oldData), 1);
                                                return { ...prevState, data };
                                            });
                                            resolve();
                                        }, (err: any) => {
                                            props.genericEnqueue(`Unable to delete service: ${err.message}. Error code: ${err.code}`, Severity.Error)
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