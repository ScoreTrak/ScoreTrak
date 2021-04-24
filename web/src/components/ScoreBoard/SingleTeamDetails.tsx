import React, {useEffect, useRef, useState} from "react";
import Accordion from "@material-ui/core/Accordion";
import AccordionSummary from "@material-ui/core/AccordionSummary";
import ExpandMoreIcon from "@material-ui/icons/ExpandMore";
import Typography from "@material-ui/core/Typography";
import AccordionDetails from "@material-ui/core/AccordionDetails";
import Box from "@material-ui/core/Box";
import {makeStyles} from "@material-ui/core/styles";
import CheckCircleOutlineIcon from '@material-ui/icons/CheckCircleOutline';
import ErrorIcon from '@material-ui/icons/Error';
import Alert from "@material-ui/lab/Alert";
import AlertTitle from "@material-ui/lab/AlertTitle";
import Grid from "@material-ui/core/Grid";
import MaterialTable, {Column} from '@material-table/core'
import FormControl from "@material-ui/core/FormControl";
import InputLabel from "@material-ui/core/InputLabel";
import Input from "@material-ui/core/Input";
import FormHelperText from "@material-ui/core/FormHelperText";
import Button from "@material-ui/core/Button";
import {Severity, SimpleReport, SimpleService} from "../../types/types";
import {GRPCClients} from "../../grpc/gRPCClients";
import PauseCircleOutlineIcon from '@material-ui/icons/PauseCircleOutline';
import {
    GetAllByServiceIDRequest as GetAllByServiceIDRequestCheck,
    GetAllByServiceIDResponse as GetAllByServiceIDResponseCheck,
} from "../../grpc/pkg/check/checkpb/check_pb";
import {
    GetAllByServiceIDRequest as GetAllByServiceIDRequestProperty,
    GetAllByServiceIDResponse as GetAllByServiceIDResponseProperty,
    Property,
    Status,
    UpdateRequest as UpdateRequestProperty
} from "../../grpc/pkg/property/propertypb/property_pb";
import {
    GetByIDRequest as GetByIDRequestHost,
    Host,
    UpdateRequest as UpdateRequestHost
} from "../../grpc/pkg/host/hostpb/host_pb";


import {UUID} from "../../grpc/pkg/proto/utilpb/uuid_pb";
import {StringValue} from "google-protobuf/google/protobuf/wrappers_pb";

const useStyles = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
        width: '100%',

    },

    customAccordionSuccessHeader: {
        borderRight: `1px solid ${theme.palette.success.main}`,
        borderLeft: `1px solid ${theme.palette.success.main}`,
    },
    customAccordionErrorHeader: {
        borderRight: `1px solid ${theme.palette.error.main}`,
        borderLeft: `1px solid ${theme.palette.error.main}`,
    },
    customAccordionWarningHeader: {
        borderRight: `1px solid ${theme.palette.warning.main}`,
        borderLeft: `1px solid ${theme.palette.warning.main}`,
    },

    paper: {
        padding: theme.spacing(1),
        textAlign: 'center',
        color: theme.palette.text.secondary,
    },

    button: {
        margin: theme.spacing(1),
    },

    input: {
        display: 'none',
    },

    heading: {
        fontSize: theme.typography.pxToRem(15),
        flexBasis: '33.33%',
        flexShrink: 0,
    },
    secondaryHeading: {
        fontSize: theme.typography.pxToRem(15),
        color: theme.palette.text.secondary,
    },
    iconSuccess: {
        color: theme.palette.success.main,
        marginRight: 12,
        fontSize: 22,
        opacity: 0.9,
    },

    iconWarning: {
        color: theme.palette.warning.main,
        marginRight: 12,
        fontSize: 22,
        opacity: 0.9,
    },

    iconError: {
        color: theme.palette.error.main,
        marginRight: 12,
        fontSize: 22,
        opacity: 0.9,
    },

}));

type CustomSingleTeamDetailsProps = {
    teamID: string,
    report: SimpleReport,
    gRPCClients: GRPCClients,
    genericEnqueue: Function
}

type SingleCheckDetails = {
    service_id: string
    host_id: string
    passed: boolean | null | undefined
    err: string | null | undefined
    log: string | null | undefined
    round_id: number
    pause: boolean
}

export default function SingleTeamDetails(props: CustomSingleTeamDetailsProps) {
    const classes = useStyles();
    const [PropertiesData, setPropertiesData] = useState<PropertiesData[]>([]);
    const [hostData, setHostData] = useState<HostData | undefined>(undefined);
    const [expanded, setExpanded] = useState<boolean | string>(false);
    const [history, setHistory] = useState<Record<string, SingleCheckDetails[]>>({});

    const handleChange = (panel: string) => (event: React.FormEvent<EventTarget>, isExpanded: boolean) => {
        setPropertiesData([])
        setHostData(undefined)
        if (!isExpanded){
            setExpanded(false)
        } else {
            setExpanded(panel);
        }
    };
    function usePreviousDT(value: SimpleReport) {
        const ref = useRef<SimpleReport>();
        useEffect(() => {
            ref.current = value;
        });
        return ref.current;
    }
    const prevDT = usePreviousDT({...props.report});

    useEffect(() => {
        if (prevDT){
            setHistory(prevState => {
                const nextState: Record<string, SingleCheckDetails[]> = {}
                Object.keys(prevState).forEach(cached_service_id => {
                    if (prevState[cached_service_id].length !== 0) {
                        const prevService = prevDT.Teams[teamID].Hosts[prevState[cached_service_id][prevState[cached_service_id].length - 1].host_id].Services
                        const check = prevService[cached_service_id].Check
                        if (!prevService[cached_service_id].Pause){
                            nextState[cached_service_id] = [...prevState[cached_service_id],
                                {
                                    service_id: prevState[cached_service_id][prevState[cached_service_id].length - 1].service_id,
                                    host_id: prevState[cached_service_id][prevState[cached_service_id].length - 1].host_id,
                                    pause: prevService[cached_service_id].Pause,
                                    passed: check !== null ? check?.Passed : null, // Causes a bug if cached service id doesnt exist anymore
                                    err: check !== null ? check?.Err : null,
                                    log: check !== null ? check?.Log : null,
                                    round_id: prevDT.Round,
                                }
                            ]
                        }
                    } else {
                        Object.keys(props.report.Teams[teamID].Hosts).forEach((host) => {
                            const currentHost = props.report.Teams[teamID].Hosts[host]
                            Object.keys(currentHost.Services).forEach((service_id) => {
                                if (cached_service_id === service_id){
                                    nextState[service_id] = [
                                        {
                                        service_id,
                                        host_id: host,
                                        pause: prevDT.Teams[teamID].Hosts[host].Services[service_id].Pause,
                                        passed: prevDT.Teams[teamID].Hosts[host].Services[service_id].Check?.Passed,
                                        err: prevDT.Teams[teamID].Hosts[host].Services[service_id].Check?.Err,
                                        log: prevDT.Teams[teamID].Hosts[host].Services[service_id].Check?.Log,
                                        round_id: prevDT.Round,
                                        }
                                    ]
                                }
                            })
                        })
                    }
                })
                return nextState
            })
        }
    }, [props.report]);
// TODO: REFACTOR the above

    const teamID = props.teamID
    return (
            <Box height="100%" width="100%" >
                {Object.keys(props.report.Teams[teamID].Hosts).map((host) => {
                    const currentHost = props.report.Teams[teamID].Hosts[host]
                    return Object.keys(currentHost.Services).map((service_id) => {
                        const simpleService = currentHost.Services[service_id]
                        let keyName
                        if (simpleService.DisplayName){
                            keyName = simpleService.DisplayName
                        } else {
                            if (currentHost.HostGroup){
                                keyName = currentHost.HostGroup.Name + "-" + simpleService.Name
                            } else{
                                keyName = simpleService.Name
                            }
                        }
                        simpleService.Pause = props.report.Teams[teamID].Pause || (currentHost.HostGroup !== undefined ? currentHost.HostGroup.Pause : false) || currentHost.Pause || currentHost.Services[service_id].Pause
                        return (
                            <Accordion expanded={expanded === keyName} onChange={handleChange(keyName)} className={!simpleService.Pause ? (simpleService.Check?.Passed ? classes.customAccordionSuccessHeader : classes.customAccordionErrorHeader) : classes.customAccordionWarningHeader}>
                                <AccordionSummary
                                    expandIcon={<ExpandMoreIcon />}
                                    aria-controls={`${keyName}bh-content`}
                                    id={`${keyName}bh-header`}>
                                    {!simpleService.Pause ? (simpleService.Check?.Passed ? <CheckCircleOutlineIcon className={classes.iconSuccess}  />  : <ErrorIcon className={classes.iconError}/>) : <PauseCircleOutlineIcon className={classes.iconWarning} />}
                                    <Typography className={classes.heading}>{keyName}</Typography>
                                    <Typography className={classes.secondaryHeading}>Host used for last round: {currentHost.Address}. Service Weight: {simpleService.Weight}</Typography>

                                </AccordionSummary>
                                <AccordionDetails>
                                    {expanded === keyName &&
                                        <SingleTeamDetailsAccordionDetailsBox {...props} history={history} prevDT={prevDT} setHistory={setHistory} host_id={host} service_id={service_id} simpleService={simpleService} setHostData={setHostData} hostData={hostData} PropertiesData={PropertiesData} setPropertiesData={setPropertiesData} />
                                    }
                                </AccordionDetails>
                            </Accordion>
                        );
                    })
                })}
            </Box>
    );
}

type SingleTeamDetailsAccordionDetailsBoxProps = {
    simpleService: SimpleService
    PropertiesData: PropertiesData[]
    setPropertiesData: React.Dispatch<React.SetStateAction<PropertiesData[]>>;
    gRPCClients: GRPCClients,
    service_id: string
    history: Record<string, SingleCheckDetails[]>
    setHistory: React.Dispatch<React.SetStateAction<Record<string, SingleCheckDetails[]>>>;
    host_id: string
    prevDT: SimpleReport | undefined
    report: SimpleReport
    setHostData: React.Dispatch<React.SetStateAction<HostData | undefined>>
    hostData: HostData | undefined
    genericEnqueue: Function
}

type HostData = {
    address: string
    edit_host: boolean
}


type PropertiesData = {
    key: string,
    value_used: string,
    service_id: string,
    value: undefined | string
    editable_value: boolean
}

function SingleTeamDetailsAccordionDetailsBox(props: SingleTeamDetailsAccordionDetailsBoxProps) {
    const simpleService = props.simpleService
    const PropertiesData = props.PropertiesData
    const classes = useStyles();
    const setPropertiesData = props.setPropertiesData
    const service_id = props.service_id
    const history = props.history
    const setHistory = props.setHistory
    const host_id = props.host_id


    const columns: Array<Column<PropertiesData>> = [
        { title: 'Key', field: 'key', editable: "never" as const},
        { title: 'Value Used', field: 'value_used', editable: "never" as const},
        { title: 'Current Value', field: 'value', emptyValue: '(empty)' },
        { title: 'editable_value', field: 'status', hidden: true },
    ]

    const columnsPreviousRounds = [
        { width: 5,  render: (rowData: any) => <div> {rowData.passed ? <CheckCircleOutlineIcon className={classes.iconSuccess}  />  : <ErrorIcon className={classes.iconError}/> }  </div>},
        { title: 'Round', field: 'round_id', defaultSort: "desc" as const, type: 'numeric' as const, width: 5},
        { title: 'Passed', field: 'passed', hidden: true, type: 'boolean' as const},
        { title: 'Parent Host ID', field: 'host_id', hidden: true},
        { title: 'Service ID', field: 'service_id', hidden: true},
        { title: 'Response', field: 'log' },
        { title: 'Error Details', field: 'err'},
        { title: 'Pause', field: 'pause', hidden: true},
    ]

    async function reloadPreviousChecks(service: string): Promise<GetAllByServiceIDResponseCheck> {
        const checksRequest =  new GetAllByServiceIDRequestCheck()
        const uuid = new UUID()
        uuid.setValue(service)
        checksRequest.setServiceId(uuid)
        return props.gRPCClients.checkClient.getAllByServiceID(checksRequest, {})
    }


    async function reloadProperties(service: string): Promise<GetAllByServiceIDResponseProperty> {
        const propertiesRequest = new GetAllByServiceIDRequestProperty()
        const uuid = new UUID()
        uuid.setValue(service)
        propertiesRequest.setServiceId(uuid)
        return props.gRPCClients.propertyClient.getAllByServiceID(propertiesRequest, {})
    }

    async function reloadHost(hostID: string) {
        const hostsRequest =  new GetByIDRequestHost()
        const uuid = new UUID()
        uuid.setValue(hostID)
        hostsRequest.setId(uuid)
        return props.gRPCClients.hostClient.getByID(hostsRequest, {})
    }

    function reloadPropertiesSetter(service_id: string, simpleService: SimpleService) {
        reloadProperties(service_id).then(results => {
            const d: PropertiesData[] = []
            for (const [key, property] of Object.entries(simpleService.Properties)) {
                const obj: PropertiesData = {key, value_used: property.Value, service_id: key, value: "", editable_value: false}
                results.getPropertiesList().forEach(res => {
                    if (key === res.getKey() && res.getStatus() === Status.EDIT){
                        obj.value = res.getValue()?.getValue() ? res.getValue()?.getValue() : undefined
                        obj.editable_value = true
                    }
                })
                d.push(obj)
            }
            setPropertiesData(d)
        }, (err: any) => {
            props.genericEnqueue(`Encountered an error while loading properties for service ${service_id}: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
    }

    const handleSetHostAddress = (e: React.FormEvent<EventTarget>, hstID: string) => {
        e.preventDefault()
        const address = (document.getElementById(`host_address_${hstID}`) as HTMLInputElement).value
        const hostsRequest = new UpdateRequestHost()
        const uuid = new UUID()
        uuid.setValue(hstID)
        const host = new Host()
        host.setId(uuid)
        host.setAddress(address)
        hostsRequest.setHost(host)
        props.gRPCClients.hostClient.update(hostsRequest, {}).then(r => {
            props.setHostData({edit_host: true, address})
            reloadHostSetter(hstID)
        }, (err: any) => {
            props.genericEnqueue(`Failed to update host details ${hstID}: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
    }

    function reloadHostSetter(host_id: string) {
        reloadHost(host_id).then(results => {
            if (results.getHost() === undefined){
                props.setHostData(undefined)
            }
            props.setHostData({address: results.getHost()?.getAddress() as string, edit_host: results.getHost()?.getEditHost()?.getValue() as boolean})
        }, (err: any) => {
            props.genericEnqueue(`Encountered an error while loading host ${host_id}: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
    }

    useEffect(() => {
        if (!history[service_id]){
            reloadPreviousChecks(service_id).then(results => {
                const d: SingleCheckDetails[] = []
                results.getChecksList().forEach(res => {
                    if (res.getRoundId().valueOf() < props.report.Round){
                        d.push({service_id, round_id: res.getRoundId().valueOf(), passed: res.getPassed()?.getValue() as boolean, log: res.getLog(), err: res.getErr(), host_id, pause: false})
                    }
                })
                setHistory(prevState => {return {...prevState, [service_id]: d }})
            }, (err: any) => {
                    props.genericEnqueue(`Encountered an error while loading previous checks: ${err.message}. Error code: ${err.code}`, Severity.Error)
                }
            )
        }
        reloadHostSetter(host_id)
        reloadPropertiesSetter(service_id, simpleService)
    }, []);

    return (
        <Box width="100%" bgcolor="background.paper" textAlign="left">
            <Grid container spacing={3}>
                <Grid item xs={12}>
                    {simpleService.Check?.Log &&
                    <Alert severity={simpleService.Check?.Passed ? "info" : "warning"}>
                        <AlertTitle >Response</AlertTitle>
                        {simpleService.Check?.Log}
                    </Alert>
                    }
                    <br/>
                    {simpleService.Check?.Err &&
                    <Alert severity="error">
                        <AlertTitle>Error Details</AlertTitle>
                        {simpleService.Check?.Err}
                    </Alert>
                    }
                </Grid>
            </Grid>
            <Grid container spacing={3}>
                <Grid item xs={6}>
                    <MaterialTable
                        options={{pageSizeOptions: [5, 10, 20, 50, 100], pageSize: 5}}
                        title="Previous Rounds"
                        columns={columnsPreviousRounds}
                        data={history[service_id]}
                    />
                </Grid>
                <Grid item xs={6}>
                    {   PropertiesData.length !== 0 &&
                    <MaterialTable
                        options={{pageSizeOptions: [3, 5, 10, 20, 50, 100], pageSize: PropertiesData.length, emptyRowsWhenPaging: false}} //
                        title="Properties"
                        columns={columns}
                        data={PropertiesData}
                        cellEditable={{
                            isCellEditable: (rowData, columnDef ) => {return rowData.editable_value},
                            onCellEditApproved: (newValue, oldValue, rowData, columnDef) => {
                                return new Promise((resolve, reject) => {
                                    setTimeout(() => {
                                        if (newValue !== "" && newValue.trim() === ""){
                                            reject("Only empty strings are not allowed")
                                        }
                                        const property = new Property()
                                        property.setKey(rowData.key)
                                        const uuid = new UUID()
                                        uuid.setValue(service_id)
                                        property.setServiceId(uuid)
                                        const stringValue = new StringValue()
                                        stringValue.setValue(newValue)
                                        property.setValue(stringValue)
                                        const updatedProperty = new UpdateRequestProperty()
                                        updatedProperty.setProperty(property)
                                        props.gRPCClients.propertyClient.update(updatedProperty, {}).then(r => {
                                            setPropertiesData((prevState) => {
                                                return prevState.map(property => {
                                                    if (property.key === rowData.key) {
                                                        return {...rowData, value: newValue ? newValue : undefined }
                                                    }
                                                    return {...property}
                                                })
                                            });
                                            resolve();
                                        }, (err: any) => {
                                            props.genericEnqueue(`Encountered an error while loading previous checks: ${err.message}. Error code: ${err.code}`, Severity.Error)
                                            reject()
                                        })
                                    }, 600);
                                })
                            }
                        }}
                    />

                    }
                    {
                        props.hostData && props.hostData.edit_host &&
                        <form style={{width: "100%", marginTop: "1vh"}} onSubmit={e => {handleSetHostAddress(e, host_id) }}>
                            <FormControl style={{ display: 'flex', flexDirection: 'row', width: "100%"}}>
                                <div>
                                    <InputLabel htmlFor="host_address">Host (Current: {props.hostData.address})</InputLabel>
                                    <Input id={`host_address_${host_id}`} aria-describedby="my-helper-text" />
                                    <FormHelperText id="my-helper-text">Set the address of the remote machine</FormHelperText>
                                </div>
                                <Button type="submit" variant="outlined" color="primary" style={{width: "10vh", height: "3vh", marginLeft: "3vh", marginTop: "auto", marginBottom: "auto"}}>
                                    Set
                                </Button >
                            </FormControl>
                        </form>
                    }
                </Grid>

            </Grid>
        </Box>
    )
}