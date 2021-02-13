import React, {useEffect} from 'react';
import {forwardRef, } from 'react';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import TextField from "@material-ui/core/TextField";
import Box from "@material-ui/core/Box";
import CircularProgress from "@material-ui/core/CircularProgress";
import {Typography} from "@material-ui/core";
import Select from "@material-ui/core/Select";
import MenuItem from "@material-ui/core/MenuItem";
import {availableChecks, Checks, } from "../Service/serviceDefaultProperties"
import Switch from "@material-ui/core/Switch";
import Button from "@material-ui/core/Button";
import {SetupProps} from "../util/util";
import {GetAllRequest} from "../../../grpc/pkg/service/servicepb/service_pb";
import {Severity} from "../../../types/types";
import {serviceColumns, serviceToServiceColumn} from "../Service/ServiceMenu";
import {propertyColumns, propertyColumnsToProperty, Status} from "./PropertiesMenu";
import {StoreRequest} from "../../../grpc/pkg/property/propertypb/property_pb";
//Todo: If Display name missing, replace with HostGroup.Name + Service.Name equivalent
const PropertiesCreate = forwardRef((props: SetupProps, ref) => {
    const [dt, setData] = React.useState<{loader: boolean, services: serviceColumns[]}>({loader: true, services: []})
    type RowType = Record<string, any> // Todo: Implement more specific types
    const [rowsData, setRowData] = React.useState<RowType>({});


    useEffect(() => {
        props.gRPCClients.serviceClient.getAll(new GetAllRequest(), {}).then(respService => {
            const rowdt: RowType = {}
            const displayNames = new Set<string>()
            respService.getServicesList().forEach(serv => {
                if (serv.getDisplayName()){
                    displayNames.add(serv.getDisplayName())
                    if (!(serv.getDisplayName() in rowdt)){
                        rowdt[serv.getDisplayName()] = {enableProcessingProperty: false}
                        Object.keys(Checks[serv.getName() as availableChecks]).forEach(key => {
                             rowdt[serv.getDisplayName()][key] = {
                                ...Checks[serv.getName() as availableChecks][key],
                                value: 'defaultValue' in Checks[serv.getName() as availableChecks][key] ? Checks[serv.getName() as availableChecks][key].defaultValue as string : '',
                                status: Checks[serv.getName() as availableChecks][key].defaultStatus ? Checks[serv.getName() as availableChecks][key].defaultStatus : Status.View,
                            }
                        })
                    }
                }
            })

            const serviceCols: serviceColumns[] = []

            respService.getServicesList().forEach(serv => {
                serviceCols.push(serviceToServiceColumn(serv))
            })

            setData({loader: false, services: serviceCols})
            setRowData({...rowdt})
        }, (err: any) => {
            props.genericEnqueue(`Encountered an error while retrieving parent Service: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
    }, []);


    const processProperties = (displayName: string, enable: boolean) => {
        setRowData(prevState => {return {...prevState, [displayName]: {...prevState[displayName], enableProcessingProperty: enable}}})
    }

    function submit() {
            const properties: propertyColumns[] = []
            Object.keys(rowsData).forEach((DisplayName) => {
                if (rowsData[DisplayName].enableProcessingProperty){
                    dt.services.forEach(service => {
                        if (service.displayName === DisplayName){
                            Object.keys(rowsData[DisplayName]).forEach(propertyKey => {
                                if (propertyKey !== "enableProcessingProperty"){
                                    properties.push({
                                        serviceId: service.id,
                                        key: propertyKey,
                                        status: rowsData[DisplayName][propertyKey].status,
                                        value: rowsData[DisplayName][propertyKey].value
                                    })
                                }
                            })
                        }
                    })
                }
            })

            const storeRequest = new StoreRequest()
            properties.forEach(property => {
                storeRequest.addProperties(propertyColumnsToProperty(property))
            })
            props.gRPCClients.propertyClient.store(storeRequest, {}).then(() => {
                    props.genericEnqueue("Success!", Severity.Success, 3000)
                },
                (err: any) => {
                    props.genericEnqueue(`Failed to save properties: ${err.message}. Error code: ${err.code}`, Severity.Error)
                })
    }

    const columns = [
        { id: 'key', label: 'Key'},
        { id: 'value', label: 'Value'},
        { id: 'status', label: 'Status'},
    ];
    return (
        <React.Fragment>
        <div>
            {!dt.loader ?
                Object.keys(rowsData).map((table) => (
                    <React.Fragment>
                        <Typography>Properties for: {table}</Typography>
                        <Table stickyHeader aria-label="sticky table" style={{marginBottom: '4vh'}}>
                            <TableHead>
                                <TableRow>
                                    {columns.map((column) => (
                                        <TableCell key={column.id} >
                                            {column.label}
                                        </TableCell>
                                    ))}
                                    <TableCell>
                                        <Switch
                                            checked={rowsData[table].enableProcessingProperty}
                                            onChange={(event) => {processProperties(table, event.target.checked)}}
                                            inputProps={{ 'aria-label': 'secondary checkbox' }}
                                        />
                                    </TableCell>
                                </TableRow>
                            </TableHead>

                            <TableBody>
                                { rowsData[table].enableProcessingProperty &&
                                    Object.keys(rowsData[table]).filter(property => {
                                        return property !== 'enableProcessingProperty';

                                    }).map(property => {
                                        return <TableRow hover role="checkbox" tabIndex={-1} >
                                            {columns.map(column => (
                                                <TableCell>
                                                    {
                                                        column.id === 'key' && rowsData[table][property].name
                                                    }

                                                    {
                                                        column.id === 'value' && rowsData[table][property].type === 'field' &&
                                                        <TextField value={rowsData[table][property].value}
                                                                   onChange={(event => {
                                                                       const val = event.target.value
                                                                       setRowData(prevState => {
                                                                           return {...prevState, [table]: {  ...prevState[table], [property]: {...prevState[table][property], value: val}}}
                                                                       })
                                                                   })}
                                                        />
                                                    }
                                                    {   column.id === 'value' && rowsData[table][property].type === 'select' &&
                                                            <Select
                                                                value={rowsData[table][property].value}
                                                                onChange={(event => {
                                                                    const val = event.target.value
                                                                    setRowData(prevState => {
                                                                        return {...prevState, [table]: {  ...prevState[table], [property]: {...prevState[table][property], value: val}}}
                                                                    })
                                                                })}
                                                            >
                                                                {
                                                                    rowsData[table][property].options.map((stat: string) => {
                                                                        return <MenuItem value={stat}>{stat}</MenuItem>
                                                                    })
                                                                }
                                                            </Select>
                                                    }
                                                    {
                                                        column.id === 'status' &&
                                                        <Select
                                                            value={rowsData[table][property].status}
                                                            onChange={(event => {
                                                                const val = event.target.value
                                                                setRowData(prevState => {
                                                                    return {...prevState, [table]: {  ...prevState[table], [property]: {...prevState[table][property], status: val}}}
                                                                })
                                                            })}
                                                        >
                                                            {
                                                                ["View", "Edit", "Hide"].map(stat => {
                                                                    return <MenuItem value={stat}>{stat}</MenuItem>
                                                                })
                                                            }
                                                        </Select>
                                                    }
                                                </TableCell>
                                            ))}

                                            <TableCell />
                                        </TableRow>
                                    })
                                }
                            </TableBody>
                        </Table>
                    </React.Fragment>
                ))
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

export default PropertiesCreate;