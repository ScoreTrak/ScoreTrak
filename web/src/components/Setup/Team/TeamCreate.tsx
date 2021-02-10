import React from 'react';
import {forwardRef} from 'react';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";
import {SetupProps, parse_index} from "../util/util";
import {StoreRequest, Team} from "../../../grpc/pkg/team/teampb/team_pb";
import {Severity} from "../../../types/types";
import {UInt64Value} from "google-protobuf/google/protobuf/wrappers_pb";


const TeamCreate = forwardRef((props: SetupProps, ref) => {
    const [rows, setRows] = React.useState<{name: JSX.Element, index: number}[]>([]);

    const columns = [
        { id: 'name', label: 'Team Name'},
        { id: 'index', label: 'Index', field: <TextField onChange={
        ((event: React.ChangeEvent<HTMLInputElement>)  => {
            setRows(parse_index(event.target.value).map((idx: number) => {
                return {
                    name: <TextField required label="Team Name" id={`name_${idx}`}/> ,
                    index: idx
                }
            }
            ))
        }
        )} id="filled-helperText" label="Index" helperText="This field is used to create host addresses. Ex: 1,2,4-15"/>

        },
    ];

    function submit() {
            const teams = rows.map(row => {
                return {index: row.index, name: (document.getElementById(`name_${row.index}`) as HTMLInputElement).value }
            })
            const storeRequest = new StoreRequest()
            teams.forEach(team => {
                storeRequest.addTeams(new Team().setIndex(new UInt64Value().setValue(team.index)).setName(team.name), 0)
            })
            props.gRPCClients.teamClient.store(storeRequest, {}).then(result => {
                props.genericEnqueue(`Teams Created!`, Severity.Success)
            }, (err: any) => {
                props.genericEnqueue(`Unable to store teams: ${err.message}. Error code: ${err.code}`, Severity.Error)
          })
        }


    return (
        <React.Fragment>
            <div>
                <Table stickyHeader aria-label="sticky table">
                    <TableHead>
                        <TableRow>
                            {columns.map((column) => (
                                <TableCell key={column.id} >
                                    {column.label}
                                </TableCell>
                            ))}
                        </TableRow>
                    </TableHead>
                    <TableHead>
                        <TableRow>
                            {columns.map((column) => (
                                <TableCell style={{minWidth: "300px"}}
                                    key={column.id}
                                >
                                    {column.field}
                                </TableCell>
                            ))}
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {rows.map((row) => {
                            return (
                                <TableRow hover role="checkbox" tabIndex={-1}>
                                    {columns.map((column) => {
                                        let value;
                                        if (column.id === 'name'){
                                            value  = row.name;
                                        }
                                        if (column.id === 'index'){
                                            value = row.index;
                                        }
                                        return (
                                            <TableCell key={column.id} >
                                                {value}
                                            </TableCell>
                                        );
                                    })}
                                </TableRow>
                            );
                        })}
                    </TableBody>
                </Table>
            </div>
            <div style={{display: 'flex',  justifyContent: 'flex-end'}}>
                 <Button onClick={submit} variant="contained" style={{ marginRight: '8px', marginTop: '8px'}}>Submit</Button>
            </div>
    </React.Fragment>
    );
})

export default TeamCreate;