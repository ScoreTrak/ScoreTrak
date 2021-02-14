import React, {useEffect} from "react";
import {SetupProps} from "../util/util";
import {Severity} from "../../../types/types";
import {
    DeleteRequest,
    GetAllRequest, RedeployRequest,
    ServiceGroup,
    StoreRequest, UpdateRequest
} from "../../../grpc/pkg/service_group/service_grouppb/service_group_pb";
import {BoolValue} from "google-protobuf/google/protobuf/wrappers_pb";
import {UUID} from "../../../grpc/pkg/proto/utilpb/uuid_pb";
import MaterialTable, {Column} from '@material-table/core'
import {Box, CircularProgress} from "@material-ui/core";


type serviceGroupColumns = {
    id: string | undefined
    name: string,
    displayName: string,
    label: string,
    enabled: boolean | undefined
    skipHelper: boolean
}

function serviceGroupToServiceGroupColumn(serviceGroup: ServiceGroup): serviceGroupColumns{
    return {
        enabled: serviceGroup.getEnabled()?.getValue(),
        id: serviceGroup.getId()?.getValue(),
        name: serviceGroup.getName(),
        displayName: serviceGroup.getDisplayName(),
        label: serviceGroup.getLabel(),
        skipHelper: serviceGroup.getSkipHelper()
    }
}

function serviceGroupColumnsToServiceGroup(serviceGroupC: serviceGroupColumns): ServiceGroup{
    const t = new ServiceGroup()
    if (serviceGroupC.enabled !== undefined) t.setEnabled(new BoolValue().setValue(serviceGroupC.enabled))
    if (serviceGroupC.id && serviceGroupC.id !== "") t.setId((new UUID().setValue(serviceGroupC.id)))
    t.setDisplayName(serviceGroupC.displayName)
    t.setLabel(serviceGroupC.label)
    t.setName(serviceGroupC.name)
    t.setSkipHelper(serviceGroupC.skipHelper)
    return t
}


export function ServiceGroupsMenu(props: SetupProps) {
    const title = "Service Group"
    props.setTitle(title)
    const columns: Array<Column<serviceGroupColumns>> =
        [
            { title: 'ID (optional)', field: 'id', editable: 'onAdd'},
            { title: 'Service Group Name', field: 'name', editable: 'onAdd' },
            { title: 'Enabled', field: 'enabled', type: 'boolean'},
            { title: 'Display Name', field: 'displayName' },
            { title: 'Skip Helper(Skips autodeploy of workers)', field: 'skipHelper', type: 'boolean' },
            { title: 'Label(Workers would be deployed on nodes with the following label)', field: 'label', editable: 'onAdd'},
        ]

    const [state, setState] = React.useState<{columns: any[], loader: boolean, data: serviceGroupColumns[]}>({
        columns,
        loader: true,
        data: []
    });

    function reloadSetter() {
        props.gRPCClients.serviceGroupClient.getAll(new GetAllRequest(), {}).then(serviceGroupResponse => {
            setState(prevState => {return{...prevState, data: serviceGroupResponse.getServiceGroupsList().map((service): serviceGroupColumns => {
                    return serviceGroupToServiceGroupColumn(service)}), loader: false}})}, (err: any) => {
            props.genericEnqueue(`Encountered an error while retrieving Service Groups: ${err.message}. Error code: ${err.code}`, Severity.Error)
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
                        actions={[
                            {
                                icon: "replay", tooltip: 'redeploy workers',
                                onClick: (event, rowData) => {
                                    return props.gRPCClients.serviceGroupClient.redeploy(new RedeployRequest().setId(new UUID().setValue((rowData as serviceGroupColumns).id as string)), {}).then(() => {
                                        props.genericEnqueue("Workers were deployed! Please make sure they are in a healthy state before enabling the service group.", Severity.Success)
                                    }, (err: any) => {
                                        props.genericEnqueue(`Unable to redeploy service group workers: ${err.message}. Error code: ${err.code}`, Severity.Error)
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
                                        const u = serviceGroupColumnsToServiceGroup(newData)
                                        storeRequest.setServiceGroup(u)
                                        props.gRPCClients.serviceGroupClient.store(storeRequest, {}).then(result => {
                                            u.setId(result.getId())
                                            setState((prevState) => {
                                                const data = [...prevState.data];
                                                data.push(serviceGroupToServiceGroupColumn(u));
                                                return { ...prevState, data };
                                            });
                                            resolve();
                                        }, (err: any) => {
                                            props.genericEnqueue(`Unable to store service group: ${err.message}. Error code: ${err.code}`, Severity.Error)
                                            reject()
                                        })
                                    }, 600);
                                }),
                            onRowUpdate: (newData, oldData) =>
                                new Promise((resolve, reject) => {
                                    setTimeout(() => {
                                        if (oldData){
                                            const updateRequest = new UpdateRequest()
                                            const u = serviceGroupColumnsToServiceGroup(newData)
                                            updateRequest.setServiceGroup(u)
                                            props.gRPCClients.serviceGroupClient.update(updateRequest, {}).then(result => {
                                                setState((prevState) => {
                                                    const data = [...prevState.data];
                                                    data[data.indexOf(oldData)] = newData;
                                                    return { ...prevState, data };
                                                });
                                                resolve();
                                            }, (err: any) => {
                                                props.genericEnqueue(`Unable to update service group: ${err.message}. Error code: ${err.code}`, Severity.Error)
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
                                        props.gRPCClients.serviceGroupClient.delete(deleteRequest, {}).then(result => {
                                            setState((prevState) => {
                                                const data = [...prevState.data];
                                                data.splice(data.indexOf(oldData), 1);
                                                return { ...prevState, data };
                                            });
                                            resolve();
                                        }, (err: any) => {
                                            props.genericEnqueue(`Unable to delete Service Group: ${err.message}. Error code: ${err.code}`, Severity.Error)
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