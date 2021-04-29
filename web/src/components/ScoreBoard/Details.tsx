import React, {useEffect} from "react";
import SingleTeamDetails from "./SingleTeamDetails";
import TableContainer from "@material-ui/core/TableContainer";
import Paper from "@material-ui/core/Paper";
import TableHead from "@material-ui/core/TableHead";
import Table from "@material-ui/core/Table";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";
import TableBody from "@material-ui/core/TableBody";
import IconButton from "@material-ui/core/IconButton";
import { makeStyles } from '@material-ui/core/styles';
import Box from '@material-ui/core/Box';
import Collapse from '@material-ui/core/Collapse';
import KeyboardArrowDownIcon from '@material-ui/icons/KeyboardArrowDown';
import KeyboardArrowUpIcon from '@material-ui/icons/KeyboardArrowUp';
import {SimpleReport} from "../../types/types";
import {Role, token} from "../../grpc/token/token";
import {GRPCClients} from "../../grpc/gRPCClients";

type DetailsProps = {
    report: SimpleReport
    isDarkTheme: boolean
    gRPCClients: GRPCClients,
    genericEnqueue: Function
}

type Row = {team_id: string, team_name: string}

export default function Details(props: DetailsProps) {
    useEffect(() => {
        document.title = "Details"
    }, []);
    const report = props.report
    function BlackTeamPanel() {
        const data: Row [] = []
        Object.keys(report.Teams).forEach(team_id => {
            data.push({
                team_id,
                team_name: report.Teams[team_id].Name,
            })
        })
        data.sort((a, b) => (a.team_name > b.team_name) ? 1 : -1)
        return (
            <TableContainer component={Paper}>
                <Table aria-label="collapsible table">
                    <TableHead>
                        <TableRow>
                            <TableCell />
                            <TableCell align="right">Team Name</TableCell>
                            <TableCell align="right">Team ID</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {data.map((row) => (
                            <CustomRow key={row.team_id} {...props} row={row} gRPCClients={props.gRPCClients}/>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        )
    }
    return (
        <div>
            <link
                rel="stylesheet"
                href="https://fonts.googleapis.com/icon?family=Material+Icons"
            />
            {
                (token.getCurrentRole() === Role.Blue || token.getCurrentRole() === Role.Red) ?
                    <SingleTeamDetails {...props} teamID={token.getCurrentTeamID() as string} gRPCClients={props.gRPCClients}/>
                    :
                BlackTeamPanel()
            }
        </div>
    );
}

const useRowStyles = makeStyles({
    root: {
        '& > *': {
            borderBottom: 'unset',
        },
    },
});

type CustomRowProps = {
    key: string,
    row: Row,
    report: SimpleReport,
    gRPCClients: GRPCClients
    genericEnqueue: Function
}

function CustomRow(props: CustomRowProps) {
    const { row } = props;
    const [open, setOpen] = React.useState(false);
    const classes = useRowStyles();
    return (
        <React.Fragment>
            <TableRow className={classes.root}>
                <TableCell>
                    <IconButton aria-label="expand row" size="small" onClick={() => setOpen(!open)}>
                        {open ? <KeyboardArrowUpIcon /> : <KeyboardArrowDownIcon />}
                    </IconButton>
                </TableCell>
                <TableCell align="right">{row.team_name}</TableCell>
                <TableCell align="right">{row.team_id}</TableCell>
            </TableRow>
            <TableRow>
                <TableCell style={{ paddingBottom: 0, paddingTop: 0 }} colSpan={6}>
                    <Collapse in={open} timeout="auto">
                        <Box margin={1}>
                            <SingleTeamDetails {...props} teamID={row.team_id} gRPCClients={props.gRPCClients}/>
                        </Box>
                    </Collapse>
                </TableCell>
            </TableRow>
        </React.Fragment>
    );
}