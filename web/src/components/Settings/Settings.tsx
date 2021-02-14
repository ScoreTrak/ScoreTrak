import React, {useEffect, useState} from "react";
import Typography from "@material-ui/core/Typography";
import Box from '@material-ui/core/Box';
import ReactJson from 'react-json-view'
import CircularProgress from "@material-ui/core/CircularProgress";
import { makeStyles } from '@material-ui/core/styles';
import SaveIcon from '@material-ui/icons/Save';
import Button from '@material-ui/core/Button';
import CloudUploadIcon from '@material-ui/icons/CloudUpload';
import Slide from "@material-ui/core/Slide";
import Dialog from "@material-ui/core/Dialog";
import DialogContent from "@material-ui/core/DialogContent";
import DialogContentText from "@material-ui/core/DialogContentText";
import DialogActions from "@material-ui/core/DialogActions";
import Accordion from '@material-ui/core/Accordion';
import AccordionDetails from '@material-ui/core/AccordionDetails';
import AccordionSummary from '@material-ui/core/AccordionSummary';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import Switch from "@material-ui/core/Switch";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import FormControl from "@material-ui/core/FormControl";
import InputLabel from "@material-ui/core/InputLabel";
import Input from "@material-ui/core/Input";
import FormHelperText from "@material-ui/core/FormHelperText";
import {SetupProps} from "../Setup/util/util";
import {Policy, UpdateRequest as UpdateRequestPolicy} from "../../grpc/pkg/policy/policypb/policy_pb";
import { saveAs } from 'file-saver';
import {
    DynamicConfig,
    GetRequest,
    GetStaticConfigRequest,
    UpdateRequest as UpdateRequestDynamicConfig
} from "../../grpc/pkg/config/configpb/config_pb";
import {Severity} from "../../types/types";
import {BoolValue, StringValue, UInt64Value} from "google-protobuf/google/protobuf/wrappers_pb";
import {
    Competition,
    DeleteCompetitionRequest,
    FetchCoreCompetitionRequest, FetchEntireCompetitionRequest, LoadCompetitionRequest,
    ResetScoresRequest
} from "../../grpc/pkg/competition/competitionpb/competition_pb";
import { TransitionProps } from '@material-ui/core/transitions';
import {Report} from "../../grpc/pkg/report/reportpb/report_pb";
import {Timestamp} from "google-protobuf/google/protobuf/timestamp_pb";
import {Check} from "../../grpc/pkg/check/checkpb/check_pb";
import {UUID} from "../../grpc/pkg/proto/utilpb/uuid_pb";
import {Round} from "../../grpc/pkg/round/roundpb/round_pb";
import {Team} from "../../grpc/pkg/team/teampb/team_pb";
import {User} from "../../grpc/pkg/user/userpb/user_pb";
import {Host} from "../../grpc/pkg/host/hostpb/host_pb";
import {HostGroup} from "../../grpc/pkg/host_group/host_grouppb/host_group_pb";
import {ServiceGroup} from "../../grpc/pkg/service_group/service_grouppb/service_group_pb";
import {Service} from "../../grpc/pkg/service/servicepb/service_pb";
import {Property} from "../../grpc/pkg/property/propertypb/property_pb";

const useStyles = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
        width: '100%',
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

}));


const Transition = React.forwardRef(function Transition(props: TransitionProps & { children?: React.ReactElement<any, any> }, ref: React.Ref<unknown>) {
    return <Slide direction="up" ref={ref} {...props} />;
});


export default function Settings(props: SetupProps & {currentPolicy: Policy.AsObject}) {
    const classes = useStyles();
    const setTitle = props.setTitle
    setTitle("Settings")
    const [dt, setData] = useState<{loader: boolean, dynamicConfig: undefined | DynamicConfig, staticConfig: object}>({
        loader: true, dynamicConfig: undefined,
        staticConfig: {}
    });



    const [open, setOpen] = React.useState<string>("");
    const [fileSelected, setFileSelected] = React.useState({selected: false, name: ""});

    const handleClickOpen = (panel: string) => () => {
        setOpen(panel);
    };

    const handleClose = () => {
        setOpen("");
    };

    const [expanded, setExpanded] = React.useState('panelConfig');

    const handleChange = (panel: string) => (event: React.ChangeEvent<{}>, isExpanded: boolean) => {
        setExpanded(isExpanded ? panel : "");
    };


    const handleSetFileSelected = () => {
        const file = document.getElementById('file') as HTMLInputElement
        setFileSelected({selected: true, name: file.files ? file.files[0].name : "" })
    }

    function loadAll(){
        props.gRPCClients.staticConfigClient.get(new GetStaticConfigRequest(), {}).then(response => {
            setData(prevState => {return{ ...prevState, staticConfig: JSON.parse(response.getStaticConfig())}})
        }, (err: any) => {
            props.genericEnqueue(`Failed to fetch static config: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })

        props.gRPCClients.dynamicConfigClient.get(new GetRequest(), {}).then(response => {
            setData(prevState => {return{ ...prevState, dynamicConfig: response.getDynamicConfig(), loader: false }})
        }, (err: any) => {
            props.genericEnqueue(`Failed to fetch dynamic config: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
    }
    useEffect(() => {
        loadAll()
    }, []);

    const handleSetEnabled = (e: React.ChangeEvent<HTMLInputElement>) => {
        const boolVal = new BoolValue().setValue(e.target.checked)
        props.gRPCClients.dynamicConfigClient.update(new UpdateRequestDynamicConfig().setDynamicConfig(new DynamicConfig().setEnabled(boolVal)), {}).then(resp => {
            setData(prevState => {
                const newDN = prevState.dynamicConfig
                newDN?.setEnabled(boolVal)
                return{...prevState, dynamicConfig: newDN}})
        }, (err: any) => {
            props.genericEnqueue(`Failed to enable competition: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
    }

    const handleSetRoundDuration = (e: React.SyntheticEvent) => {
        e.preventDefault()
        const val = Number((document.getElementById("round_duration") as HTMLInputElement).value)
        props.gRPCClients.dynamicConfigClient.update(new UpdateRequestDynamicConfig().setDynamicConfig(new DynamicConfig().setRoundDuration(val)), {}).then(resp => {
            setData(prevState => {
                const newDN = prevState.dynamicConfig
                newDN?.setRoundDuration(val)
                return{...prevState, dynamicConfig: newDN}})
        }, (err: any) => {
            props.genericEnqueue(`Failed to enable competition: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
    }

    const handleSetPolicy = (policy: Policy) => {
        props.gRPCClients.policyClient.update(new UpdateRequestPolicy().setPolicy(policy), {}).then(resp => {}, (err: any) => {
            props.genericEnqueue(`Failed to update policy: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
    }

    const handleUpload = () => {
        const file = document.getElementById('file') as HTMLInputElement
        if (file.files && file.files[0]){
            const reader = new FileReader();
            reader.readAsText(file.files[0], 'UTF-8')
            reader.onload = function (evt){
                const obj = JSON.parse(evt.target?.result as string) as Competition.AsObject
                const comp = new Competition()
                if (obj.report){
                    comp.setReport(new Report().setCache(obj.report.cache).setUpdatedAt(new Timestamp().setSeconds(obj.report.updatedAt?.seconds as number)))
                }
                if (obj.checksList && obj.checksList.length !== 0){
                    const checks: Check[] = []
                    obj.checksList.forEach(chck => {
                        const ch = new Check()
                        ch.setErr(chck.err)
                        ch.setLog(chck.log)
                        ch.setPassed(new BoolValue().setValue(chck.passed?.value as boolean))
                        ch.setRoundId(chck.roundId)
                        ch.setServiceId(new UUID().setValue(chck.serviceId?.value as string))
                        checks.push(ch)
                    })
                    comp.setChecksList(checks)
                }
                if (obj.roundsList && obj.roundsList.length !== 0){
                    const rounds: Round[] = []
                    obj.roundsList.forEach(rnd => {
                        const rn = new Round()
                        rn.setStart(new Timestamp().setSeconds(rnd.start?.seconds as number))
                        if (rnd.finish){
                            rn.setFinish(new Timestamp().setSeconds(rnd.finish.seconds))
                        }
                        rn.setNote(rnd.note)
                        rn.setId(rnd.id)
                        rn.setErr(rnd.err)
                        rounds.push(rn)
                    })
                    comp.setRoundsList(rounds)
                }

                comp.setDynamicConfig(new DynamicConfig().setRoundDuration(obj.dynamicConfig?.roundDuration as number).setEnabled(new BoolValue().setValue(obj.dynamicConfig?.enabled?.value as boolean)))
                comp.setPolicy(new Policy()
                    .setAllowChangingUsernamesAndPasswords(new BoolValue().setValue(obj.policy?.allowChangingUsernamesAndPasswords?.value as boolean))
                    .setAllowRedTeamLaunchingServiceTestsManually(new BoolValue().setValue(obj.policy?.allowRedTeamLaunchingServiceTestsManually?.value as boolean))
                    .setShowAddresses(new BoolValue().setValue(obj.policy?.showAddresses?.value as boolean))
                    .setShowPoints(new BoolValue().setValue(obj.policy?.showPoints?.value as boolean))
                    .setAllowUnauthenticatedUsers(new BoolValue().setValue(obj.policy?.allowUnauthenticatedUsers?.value as boolean))
                )

                if (obj.teamsList && obj.teamsList.length !== 0){
                    const teams: Team[] = []
                    obj.teamsList.forEach(tm => {
                        const t = new Team()
                        t.setName(tm.name)
                        t.setIndex(new UInt64Value().setValue(tm.index?.value as number))
                        t.setId(new UUID().setValue(tm.id?.value as string))
                        t.setHide(new BoolValue().setValue(tm.hide?.value as boolean))
                        t.setPause(new BoolValue().setValue(tm.pause?.value as boolean))
                        teams.push(t)
                    })
                    comp.setTeamsList(teams)
                }

                if (obj.hostGroupsList && obj.hostGroupsList.length !== 0){
                    const hostGroups: HostGroup[] = []
                    obj.hostGroupsList.forEach(hostGrp => {
                        const h = new HostGroup()
                        h.setName(hostGrp.name)
                        h.setId(new UUID().setValue(hostGrp.id?.value as string))
                        h.setPause(new BoolValue().setValue(hostGrp.pause?.value as boolean))
                        h.setHide(new BoolValue().setValue(hostGrp.hide?.value as boolean))
                        hostGroups.push(h)
                    })
                    comp.setHostGroupsList(hostGroups)
                }

                if (obj.serviceGroupsList && obj.serviceGroupsList.length !== 0){
                    const serviceGroup: ServiceGroup[] = []
                    obj.serviceGroupsList.forEach(servGrp => {
                        const sg = new ServiceGroup()
                        sg.setSkipHelper(servGrp.skipHelper)
                        sg.setName(servGrp.name)
                        sg.setLabel(servGrp.label)
                        sg.setDisplayName(servGrp.displayName)
                        sg.setEnabled(new BoolValue().setValue(servGrp.enabled?.value as boolean))
                        sg.setId(new UUID().setValue(servGrp.id?.value as string))
                        serviceGroup.push(sg)
                    })
                    comp.setServiceGroupsList(serviceGroup)
                }

                if (obj.usersList && obj.usersList.length !== 0){
                    const users: User[] = []
                    obj.usersList.forEach(usr => {
                        const u = new User()
                        u.setId(new UUID().setValue(usr.id?.value as string))
                        u.setTeamId(new UUID().setValue(usr.teamId?.value as string))
                        u.setRole(usr.role)
                        u.setPassword(usr.password)
                        u.setUsername(usr.username)
                        u.setPasswordHash(usr.passwordHash)
                        users.push(u)
                    })
                    comp.setUsersList(users)
                }

                if (obj.hostsList && obj.hostsList.length !== 0){
                    const hosts: Host [] = []
                    obj.hostsList.forEach(hst => {
                        const h = new Host()
                        h.setPause(new BoolValue().setValue(hst.pause?.value as boolean))
                        h.setHide(new BoolValue().setValue(hst.hide?.value as boolean))
                        h.setId(new UUID().setValue(hst.id?.value as string))
                        h.setHostGroupId(new UUID().setValue(hst.hostGroupId?.value as string))
                        h.setEditHost(new BoolValue().setValue(hst.editHost?.value as boolean))
                        h.setAddress(hst.address)
                        h.setTeamId(new UUID().setValue(hst.teamId?.value as string))
                        hosts.push(h)
                    })
                    comp.setHostsList(hosts)
                }

                if (obj.servicesList && obj.servicesList.length !== 0){
                    const services: Service [] = []
                    obj.servicesList.forEach(serv => {
                        const s = new Service()
                        s.setPointsBoost(new UInt64Value().setValue(serv.pointsBoost?.value as number))
                        s.setRoundDelay(new UInt64Value().setValue(serv.roundDelay?.value as number))
                        s.setServiceGroupId(new UUID().setValue(serv.serviceGroupId?.value as string))
                        s.setRoundUnits(serv.roundUnits)
                        s.setWeight(new UInt64Value().setValue(serv.weight?.value as number))
                        s.setName(serv.name)
                        s.setPause(new BoolValue().setValue(serv.pause?.value as boolean))
                        s.setHide(new BoolValue().setValue(serv.hide?.value as boolean))
                        s.setHostId(new UUID().setValue(serv.hostId?.value as string))
                        s.setDisplayName(serv.displayName)
                        s.setId(new UUID().setValue(serv.id?.value as string))
                        services.push(s)
                    })
                    comp.setServicesList(services)
                }

                if (obj.propertiesList && obj.propertiesList.length !== 0){
                    const properties: Property[] = []
                    obj.propertiesList.forEach(prop => {
                        const p = new Property()
                        p.setStatus(prop.status)
                        p.setValue(new StringValue().setValue(prop.value?.value as string))
                        p.setKey(prop.key)
                        p.setServiceId(new UUID().setValue(prop.serviceId?.value as string))
                        properties.push(p)
                    })
                    comp.setPropertiesList(properties)
                }
                props.gRPCClients.competitionClient.loadCompetition(new LoadCompetitionRequest().setCompetition(comp), {}).then(resp => {
                    props.genericEnqueue("Success!", Severity.Success)
                }, (err: any) => {
                    props.genericEnqueue(`Failed to upload competition: ${err.message}. Error code: ${err.code}`, Severity.Error)
                })


            }
            reader.onerror = function (evt) {
                props.genericEnqueue(`Failed to open the file`, Severity.Error)
            }
        }
        handleClose()
    }

    const handleResetCompetition = () => {
        props.gRPCClients.competitionClient.resetScores(new ResetScoresRequest(), {}).then(resp => {
            props.genericEnqueue("Successfully reset all of the scores!", Severity.Success)
        }, (err: any) => {
            props.genericEnqueue(`Failed to reset scores: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
        handleClose()
    }

    const handleDeleteCompetition = () => {
        props.gRPCClients.competitionClient.deleteCompetition(new DeleteCompetitionRequest(), {}).then(resp => {
            props.genericEnqueue("Successfully deleted all competition data!", Severity.Success)
        }, (err: any) => {
            props.genericEnqueue(`Failed to delete competition data: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
        handleClose()
    }

    function saveJSONtoFile(obj: object | undefined, filename: string) {
        const fileToSave = new Blob([JSON.stringify(obj)], {
            type: 'application/json'
        });
        saveAs(fileToSave, filename);
    }

    return (

            <Box height='85vh' width="100%" textAlign="left">
            {!dt.loader ?
                <React.Fragment>
                    <Accordion expanded={expanded === 'panelConfig'} onChange={handleChange('panelConfig')}>
                        <AccordionSummary
                            expandIcon={<ExpandMoreIcon />}
                            aria-controls="panelConfigbh-content"
                            id="panelConfigbh-header"
                        >
                            <Typography className={classes.heading}>ScoreTrak Settings</Typography>
                            <Typography className={classes.secondaryHeading}>Following are the dynamically configurable settings for scoring masters</Typography>

                        </AccordionSummary>
                        <AccordionDetails>
                            <Box p={1} m={1} bgcolor="background.paper">

                            <FormControlLabel
                                control={
                            <Switch checked={dt.dynamicConfig?.getEnabled()?.getValue()} onChange={handleSetEnabled} />
                                }
                                label="Enable Competition?"
                            />

                            <br/>

                                <form onSubmit={handleSetRoundDuration}>
                                    <FormControl>
                                        <InputLabel htmlFor="round_duration">Round Duration (Current: {dt.dynamicConfig?.getRoundDuration()})</InputLabel>
                                        <Input id="round_duration" aria-describedby="my-helper-text" type="number" inputProps={{ min: "20"}} />
                                        <FormHelperText id="my-helper-text">Number of seconds it takes for one round to elapse.</FormHelperText>
                                        <Button type="submit" variant="outlined" color="primary" >
                                            Set
                                        </Button>
                                    </FormControl>
                                </form>

                            </Box>
                        </AccordionDetails>
                    </Accordion>

                    <Accordion expanded={expanded === 'panelPolicy'} onChange={handleChange('panelPolicy')}>
                        <AccordionSummary
                            expandIcon={<ExpandMoreIcon />}
                            aria-controls="panelPolicybh-content"
                            id="panelPolicybh-header"
                        >
                            <Typography className={classes.heading}>Policy</Typography>
                            <Typography className={classes.secondaryHeading}>Configure policies for allowing/disallowing resources</Typography>

                        </AccordionSummary>
                        <AccordionDetails>
                            <Box p={1} m={1} bgcolor="background.paper">
                            <FormControlLabel
                                control={
                                    <Switch checked={props.currentPolicy.allowUnauthenticatedUsers?.value} onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
                                        handleSetPolicy(new Policy().setAllowUnauthenticatedUsers(new BoolValue().setValue(e.target.checked)))
                                    }} value="allow_unauthenticated_users" />
                                }
                                label="Allow unauthenticated users to see scoreboard"
                            />
                            <br />
                            <FormControlLabel
                                control={
                                    <Switch checked={props.currentPolicy.allowChangingUsernamesAndPasswords?.value} onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
                                        handleSetPolicy(new Policy().setAllowChangingUsernamesAndPasswords(new BoolValue().setValue(e.target.checked)))
                                    }} value="allow_changing_usernames_and_passwords" />
                                }
                                label="Allow users to change usernames and passwords"
                            />
                                <br />
                            <FormControlLabel
                                control={
                                    <Switch checked={props.currentPolicy.showPoints?.value} onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
                                        handleSetPolicy(new Policy().setShowPoints(new BoolValue().setValue(e.target.checked)))
                                    }} value="allow_to_see_points" />
                                }
                                label="Allow users to see other teams' points"
                            />
                                <br />
                            <FormControlLabel
                                control={
                                    <Switch checked={props.currentPolicy.showAddresses?.value} onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
                                        handleSetPolicy(new Policy().setShowAddresses(new BoolValue().setValue(e.target.checked)))
                                    }} value="show_addresses" />
                                }
                                label="Allow users to see other teams' addresses"
                            />
                                <br />
                            <FormControlLabel
                                control={
                                    <Switch checked={props.currentPolicy.allowRedTeamLaunchingServiceTestsManually?.value} onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
                                        handleSetPolicy(new Policy().setAllowRedTeamLaunchingServiceTestsManually(new BoolValue().setValue(e.target.checked)))
                                    }} value="allow_red_team_launching_service_tests_manually" />
                                }
                                label="Allow Red Team to manually launch service tests(Only applies to a parrent team)"
                            />
                            </Box>
                        </AccordionDetails>
                    </Accordion>

                    <Accordion expanded={expanded === 'panelStaticConfig'} onChange={handleChange('panelStaticConfig')}>
                        <AccordionSummary
                            expandIcon={<ExpandMoreIcon />}
                            aria-controls="panelStaticConfigbh-content"
                            id="panelStaticConfigbh-header"
                        >
                            <Typography className={classes.heading}>Static ScoreTrak Config</Typography>
                            <Typography className={classes.secondaryHeading}>
                                This is a JSON representation of the Static Config for Scoring Master
                            </Typography>
                        </AccordionSummary>
                        <AccordionDetails>
                            <Box component="span" display="block" p={1} m={1} bgcolor="background.paper">
                                <ReactJson src={dt.staticConfig} style={{backgroundColor: "inherit"}} onDelete={false} onEdit={false} displayDataTypes={false} displayObjectSize={false} theme={props.isDarkTheme ? "monokai" : "bright:inverted"}/>
                            </Box>
                        </AccordionDetails>
                    </Accordion>
                    <Accordion expanded={expanded === 'panelImportExport'} onChange={handleChange('panelImportExport')}>
                        <AccordionSummary
                            expandIcon={<ExpandMoreIcon />}
                            aria-controls="panelImportExportbh-content"
                            id="panelImportExportbh-header"
                        >
                            <Typography className={classes.heading}>Export/Import Competition</Typography>
                            <Typography className={classes.secondaryHeading}>
                                Export or Import the competition using a JSON representation of the competition
                            </Typography>
                        </AccordionSummary>
                        <AccordionDetails>
                            <Box width="100%" p={1} m={1} bgcolor="background.paper" textAlign="center">
                                <Button
                                    variant="contained"
                                    color="primary"
                                    className={classes.button}
                                    startIcon={<SaveIcon />}
                                    onClick={() => {
                                        props.gRPCClients.competitionClient.fetchCoreCompetition(new FetchCoreCompetitionRequest(), {}).then(resp => {
                                        saveJSONtoFile(resp.getCompetition()?.toObject(), "core-competition.json")
                                    }, (err: any) => {
                                        props.genericEnqueue(`Failed to fetch competition: ${err.message}. Error code: ${err.code}`, Severity.Error)
                                    })}}
                                >
                                    Export Core Competition
                                </Button>
                                <Button variant="outlined" color="primary" onClick={handleClickOpen("upload")}>
                                    Upload Competition
                                </Button>
                                <Dialog
                                    open={open === 'upload'}
                                    TransitionComponent={Transition}
                                    keepMounted
                                    onClose={handleClose}
                                    aria-labelledby="alert-dialog-slide-title"
                                    aria-describedby="alert-dialog-slide-description"
                                >
                                    <DialogContent>
                                        <DialogContentText id="alert-dialog-slide-description" align="center">
                                            <input
                                                className={classes.input}
                                                id="file"
                                                type="file"
                                                onChange={handleSetFileSelected}
                                            />
                                            <label htmlFor="file">
                                                <Button
                                                    variant="contained"
                                                    color="primary"
                                                    className={classes.button}
                                                    component="span"
                                                    startIcon={<CloudUploadIcon />}
                                                >
                                                    Load Competition
                                                </Button>
                                            </label>
                                        </DialogContentText>
                                        { fileSelected.selected &&
                                        <Typography align="center">{fileSelected.name}</Typography>
                                        }
                                    </DialogContent>
                                    {fileSelected &&
                                    <DialogActions>
                                        <Button onClick={handleClose} color="primary">
                                            Cancel
                                        </Button>
                                        <Button onClick={handleUpload} color="primary">
                                            Upload
                                        </Button>
                                    </DialogActions>
                                    }
                                </Dialog>
                                <Button
                                    variant="contained"
                                    color="primary"
                                    className={classes.button}
                                    startIcon={<SaveIcon />}
                                    onClick={() => {props.gRPCClients.competitionClient.fetchEntireCompetition(new FetchEntireCompetitionRequest(), {}).then(resp => {
                                        saveJSONtoFile(resp.getCompetition()?.toObject(), "entire-competition.json")
                                    }, (err: any) => {
                                        props.genericEnqueue(`Failed to fetch competition: ${err.message}. Error code: ${err.code}`, Severity.Error)
                                    })}}
                                >
                                    Export Entire Competition
                                </Button>
                            </Box>
                        </AccordionDetails>
                    </Accordion>

                    <Accordion expanded={expanded === 'panelDeleteReset'} onChange={handleChange('panelDeleteReset')}>
                        <AccordionSummary
                            expandIcon={<ExpandMoreIcon />}
                            aria-controls="panelDeleteReseth-content"
                            id="panelDeleteResetbh-header"
                        >
                            <Typography className={classes.heading}>Reset/Delete Competition</Typography>
                            <Typography className={classes.secondaryHeading}>
                                Reset: Resets Scores, and Rounds. Delete: Removes everything
                            </Typography>
                        </AccordionSummary>
                        <AccordionDetails>
                            <Box width="100%" p={1} m={1} bgcolor="background.paper" textAlign="center">
                                <Button variant="outlined" style={{color: "red", border: '1px solid red' }} onClick={handleClickOpen("reset")} className={classes.button}>
                                    Reset Competition
                                </Button>
                                <Dialog
                                    open={open === "reset"}
                                    TransitionComponent={Transition}
                                    keepMounted
                                    onClose={handleClose}
                                    aria-labelledby="alert-dialog-slide-title"
                                    aria-describedby="alert-dialog-slide-description"
                                >
                                    <DialogContent>
                                        <DialogContentText id="alert-dialog-slide-description" align="center">
                                            Are you sure?
                                        </DialogContentText>
                                    </DialogContent>

                                    <DialogActions>
                                        <Button onClick={handleClose} color="primary">
                                            Cancel
                                        </Button>
                                        <Button onClick={handleResetCompetition} color="primary">
                                            Reset Competition
                                        </Button>
                                    </DialogActions>

                                </Dialog>


                                <Button variant="outlined" style={{color: "red", border: '1px solid red' }} onClick={handleClickOpen("delete")} className={classes.button}>
                                    Delete Competition
                                </Button>
                                <Dialog
                                    open={open === "delete"}
                                    TransitionComponent={Transition}
                                    keepMounted
                                    onClose={handleClose}
                                    aria-labelledby="alert-dialog-slide-title"
                                    aria-describedby="alert-dialog-slide-description"
                                >
                                    <DialogContent>
                                        <DialogContentText id="alert-dialog-slide-description" align="center">
                                           Are you sure?
                                        </DialogContentText>
                                    </DialogContent>

                                    <DialogActions>
                                        <Button onClick={handleClose} color="primary">
                                            Cancel
                                        </Button>
                                        <Button onClick={handleDeleteCompetition} color="primary">
                                            Delete Competition
                                        </Button>
                                    </DialogActions>

                                </Dialog>
                            </Box>
                        </AccordionDetails>
                    </Accordion>

                </React.Fragment>
                :
                <Box height="100%" width="100%" m="auto">
                    <CircularProgress  />
                </Box>
            }

            </Box>
    );
}