import React, {useEffect} from "react";
import Stepper from "@material-ui/core/Stepper";
import Step from "@material-ui/core/Step";
import StepButton from "@material-ui/core/StepButton";
import {SetupProps} from "../util/util";
import TeamCreate from "./TeamCreate";
import Box from "@material-ui/core/Box";
import MaterialTable, {Column} from '@material-table/core'
import {
    DeleteRequest,
    GetAllRequest,
    StoreRequest, UpdateRequest,
    Team
} from "../../../grpc/pkg/team/teampb/team_pb";
import {Severity} from "../../../types/types";
import {UUID} from "../../../grpc/pkg/proto/utilpb/uuid_pb";
import {CircularProgress} from "@material-ui/core";
import {BoolValue, UInt64Value} from "google-protobuf/google/protobuf/wrappers_pb";

function getSteps() {
    return ['Regular View', 'Quick Create'];
}

export default function TeamMenu(props: SetupProps) {
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
                    <TeamMenuTable {...props}/>
                }
                {
                    activeStep === 1 &&
                    <TeamCreate {...props} />
                }
            </div>
        </Box>
    );
}


export type teamColumns = {
    id: string | undefined
    name: string,
    index: number | undefined
    pause: boolean | undefined
    hide: boolean | undefined
}

export function teamToTeamColumn(team: Team): teamColumns{
    return {
        pause: team.getPause()?.getValue(),
        id: team.getId()?.getValue(),
        name: team.getName(),
        index: team.getIndex()?.getValue(),
        hide: team.getHide()?.getValue()
    }
}

export function teamColumnsToTeam(teamC: teamColumns): Team{
    const t = new Team()
    if (teamC.pause !== undefined) t.setPause(new BoolValue().setValue(teamC.pause))
    if (teamC.hide !== undefined) t.setHide(new BoolValue().setValue(teamC.hide))
    if (teamC.id && teamC.id !== "") t.setId((new UUID().setValue(teamC.id)))
    t.setName(teamC.name)
    if (teamC.index !== undefined) t.setIndex(new UInt64Value().setValue(teamC.index))
    return t
}


function TeamMenuTable(props: SetupProps) {
    const title = "Teams"
    props.setTitle(title)
    const columns :Array<Column<teamColumns>> =
        [
            { title: 'ID (optional)', field: 'id', editable: 'onAdd'},
            { title: 'Team Name', field: 'name' },
            { title: 'Index', field: 'index', type: 'numeric' },
            { title: 'Hide from Scoreboard', field: 'hide', type: 'boolean', initialEditValue: false},
            { title: 'Pause Scoring', field: 'pause', type: 'boolean', initialEditValue: false},
        ]

    const [state, setState] = React.useState<{columns: any[], loader: boolean, data: teamColumns[]}>({
        columns,
        loader: true,
        data: []
    });

    function reloadSetter() {
        props.gRPCClients.teamClient.getAll(new GetAllRequest(), {}).then(teamsResponse => {
            setState(prevState => {return{...prevState, data: teamsResponse.getTeamsList().map((team): teamColumns => {
                return teamToTeamColumn(team)}), loader: false}})}, (err: any) => {
            props.genericEnqueue(`Encountered an error while retrieving Teams: ${err.message}. Error code: ${err.code}`, Severity.Error)
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
                        editable={{
                            onRowAdd: (newData) =>
                                new Promise((resolve, reject) => {
                                    setTimeout(() => {
                                        const storeRequest = new StoreRequest()
                                        const u = teamColumnsToTeam(newData)
                                        storeRequest.addTeams(u, 0)
                                        props.gRPCClients.teamClient.store(storeRequest, {}).then(result => {
                                            u.setId(result.getIdsList()[0])
                                            setState((prevState) => {
                                                const data = [...prevState.data];
                                                data.push(teamToTeamColumn(u));
                                                return { ...prevState, data };
                                            });
                                            resolve();
                                        }, (err: any) => {
                                            props.genericEnqueue(`Unable to store team: ${err.message}. Error code: ${err.code}`, Severity.Error)
                                            reject()
                                        })
                                    }, 600);
                                }),
                            onRowUpdate: (newData, oldData) =>
                                new Promise((resolve, reject) => {
                                    setTimeout(() => {
                                        if (oldData){
                                            const updateRequest = new UpdateRequest()
                                            const u = teamColumnsToTeam(newData)
                                            updateRequest.setTeam(u)
                                            props.gRPCClients.teamClient.update(updateRequest, {}).then(result => {
                                                setState((prevState) => {
                                                    const data = [...prevState.data];
                                                    data[data.indexOf(oldData)] = newData;
                                                    return { ...prevState, data };
                                                });
                                                resolve();
                                            }, (err: any) => {
                                                props.genericEnqueue(`Unable to update team: ${err.message}. Error code: ${err.code}`, Severity.Error)
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
                                        props.gRPCClients.teamClient.delete(deleteRequest, {}).then(result => {
                                            setState((prevState) => {
                                                const data = [...prevState.data];
                                                data.splice(data.indexOf(oldData), 1);
                                                return { ...prevState, data };
                                            });
                                            resolve();
                                        }, (err: any) => {
                                            props.genericEnqueue(`Unable to delete team: ${err.message}. Error code: ${err.code}`, Severity.Error)
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