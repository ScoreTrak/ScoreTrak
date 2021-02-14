import React, {useEffect} from "react";
import {SetupProps} from "../util/util";
import {Role} from "../../../grpc/token/token";
import Box from "@material-ui/core/Box";
import MaterialTable, {Column} from '@material-table/core'
import {GetAllRequest as GetAllRequestTeam} from "../../../grpc/pkg/team/teampb/team_pb";
import {
    DeleteRequest,
    GetAllRequest as GetAllRequestUser,
    Role as ProtoRole,
    StoreRequest, UpdateRequest,
    User
} from "../../../grpc/pkg/user/userpb/user_pb";
import {Severity} from "../../../types/types";
import {UUID} from "../../../grpc/pkg/proto/utilpb/uuid_pb";
import {CircularProgress} from "@material-ui/core";

type userColumns = {
    id: string | undefined
    username: string,
    password: string,
    passwordHash: string | undefined,
    teamId: string | undefined
    role: Role | undefined
}

function userToUserColumn(user: User): userColumns{
    return {
        id: user.getId()?.getValue(),
        password: user.getPassword(),
        passwordHash: user.getPasswordHash(),
        role: ProtoRoleToRole(user.getRole()),
        teamId: user.getTeamId()?.getValue() ,
        username: user.getUsername()
    }
}

function userColumnsToUser(userC: userColumns): User{
    const u = new User()
    if (userC.id && userC.id !== "") u.setId((new UUID().setValue(userC.id)))
    u.setPassword(userC.password)
    u.setUsername(userC.username)
    u.setRole(RoleToProtoRole(userC.role))
    if (userC.teamId && userC.teamId !== "") u.setTeamId((new UUID().setValue(userC.teamId)))
    return u
}

function ProtoRoleToRole (eRole : ProtoRole): Role | undefined{
    if (eRole === ProtoRole.BLUE) return Role.Blue
    if (eRole === ProtoRole.BLACK) return Role.Black
    if (eRole === ProtoRole.RED) return Role.Red
    return undefined
}

function RoleToProtoRole (role : Role | undefined): ProtoRole {
    if (role === Role.Blue) return ProtoRole.BLUE
    if (role === Role.Black) return ProtoRole.BLACK
    if (role === Role.Red) return ProtoRole.RED
    return ProtoRole.ROLE_NOT_SET
}


export default function UserMenu(props: SetupProps) {
    const title = "Users"
    props.setTitle(title)
    const columns :Array<Column<userColumns>> =
        [
            { title: 'ID (optional)', field: 'id', editable: 'onAdd' as const},
            { title: 'Username', field: 'username' },
            { title: 'Password', field: 'password', render: (rowData: any) => <React.Fragment/> },
            { title: 'Password Hash', field: 'passwordHash', editable: 'never' as const},
            { title: 'Team ID', field: 'teamId' },
            { title: 'Role', field: 'role', lookup: { [Role.Black]: Role.Black, [Role.Blue]: Role.Blue, [Role.Red]: Role.Red }},
        ]
    const [state, setState] = React.useState<{columns: any[], loaderHost: boolean, loaderUser: boolean, data: userColumns[]}>({
        columns,
        loaderHost: true,
        loaderUser: true,
        data: []
    });

    function reloadSetter() {
        const lookup: Record<string, string> = {}
        props.gRPCClients.teamClient.getAll(new GetAllRequestTeam(), {}).then(teamsResponse => {
            for (let i = 0; i < teamsResponse.getTeamsList().length; i++){
                lookup[teamsResponse.getTeamsList()[i].getId()?.getValue() as string] = `${teamsResponse.getTeamsList()[i].getName()} (ID:${teamsResponse.getTeamsList()[i].getId()?.getValue() as string})`
            }
            setState(prevState => {
                const columns = prevState.columns
                for (let i = 0; i < columns.length; i++){
                    if (columns[i].field === "teamId"){
                        columns[i].lookup = lookup
                    }
                }
                return{...prevState, columns, loaderHost: false
            }})
        }, (err: any) => {
            props.genericEnqueue(`Encountered an error while retrieving parent Teams: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
        props.gRPCClients.userClient.getAll(new GetAllRequestUser(), {}).then(usersResponse => {
            setState(prevState => {return{...prevState, data: usersResponse.getUsersList().map((user): userColumns => {
                return userToUserColumn(user)}), loaderUser: false}})}, (err: any) => {
            props.genericEnqueue(`Encountered an error while retrieving Users: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
    }
    useEffect(() => {
        reloadSetter()
    }, []);

    return (
        <React.Fragment>
            {!state.loaderUser && !state.loaderHost ?
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
                                        // https://github.com/protocolbuffers/protobuf/issues/1591
                                        const u = userColumnsToUser(newData)
                                        storeRequest.addUsers(u, 0)
                                        props.gRPCClients.userClient.store(storeRequest, {}).then(result => {
                                            u.setId(result.getIdsList()[0])
                                            setState((prevState) => {
                                                const data = [...prevState.data];
                                                data.push(userToUserColumn(u));
                                                return { ...prevState, data };
                                            });
                                            resolve();
                                        }, (err: any) => {
                                            props.genericEnqueue(`Unable to store user: ${err.message}. Error code: ${err.code}`, Severity.Error)
                                            reject()
                                        })
                                    }, 600);
                                }),
                            onRowUpdate: (newData, oldData) =>
                                new Promise((resolve, reject) => {
                                    setTimeout(() => {
                                        if (oldData){
                                            const updateRequest = new UpdateRequest()
                                            const u = userColumnsToUser(newData)
                                            updateRequest.setUser(u)
                                            props.gRPCClients.userClient.update(updateRequest, {}).then(result => {
                                                setState((prevState) => {
                                                    const data = [...prevState.data];
                                                    data[data.indexOf(oldData)] = newData;
                                                    return { ...prevState, data };
                                                });
                                                resolve();
                                            }, (err: any) => {
                                                props.genericEnqueue(`Unable to update user: ${err.message}. Error code: ${err.code}`, Severity.Error)
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
                                        props.gRPCClients.userClient.delete(deleteRequest, {}).then(result => {
                                            setState((prevState) => {
                                                const data = [...prevState.data];
                                                data.splice(data.indexOf(oldData), 1);
                                                return { ...prevState, data };
                                            });
                                            resolve();
                                        }, (err: any) => {
                                            props.genericEnqueue(`Unable to delete user: ${err.message}. Error code: ${err.code}`, Severity.Error)
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