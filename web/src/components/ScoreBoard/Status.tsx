import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableFooter from '@material-ui/core/TableFooter';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TablePagination from '@material-ui/core/TablePagination';
import TableRow from '@material-ui/core/TableRow';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Switch from '@material-ui/core/Switch';
import {SimpleReport, SimpleService} from "../../types/types";
import {Policy} from "../../grpc/pkg/policy/policypb/policy_pb";
import {token} from "../../grpc/token/token";


const useStyles = makeStyles({
    root: {
        width: '100%',
        height: '100%'
    },
    tableNavigator:{
        marginRight: "5vh",
        marginLeft: "5vh"
    }
});

type RanksProps = {
    report: SimpleReport
    isDarkTheme: boolean
    currentPolicy: Policy.AsObject,
}

export default function Status(props: RanksProps) {
    const classes = useStyles();
    const [rowPage, setRowPage] = React.useState<number>(0);
    const [rowsPerPage, setRowsPerPage] = React.useState<number>(25);
    const [dense, setDense] = React.useState<boolean>(false);
    const [hideAddresses, setHideAddresses] = React.useState<boolean>(false);

    const toggleHideAddresses = () => {
        setHideAddresses(prevState => !prevState)
    };

    const toggleChangeDense = () => {
        setDense(prevState => !prevState);
    };

    const handleRowChangePage = (event: React.MouseEvent<HTMLButtonElement> | null, page: number) => {
        setRowPage(page);
    };

    const handleChangeRowsPerPage = (event: { target: { value: React.ReactText; }; }) => {
        setRowsPerPage(Number(event.target.value));
        setRowPage(0);
    };
    
    const [columnPage, setColumnPage] = React.useState(0);
    const [columnsPerPage, setColumnsPerPage] = React.useState(25);
    const handleColumnChangePage = (event: React.MouseEvent<HTMLButtonElement> | null, page: number) => {
        setColumnPage(page);
    };
    const handleChangeColumnsPerPage = (event: { target: { value: React.ReactText; }; }) => {
        setColumnsPerPage(Number(event.target.value));
        setColumnPage(0);
    };

    const report = props.report
    let teamNamesSet = new Set<string>();
    let data: Record<string, Record<string, SimpleService & {Address: string}>> = {}
    let dataKeys = new Set<string>();
    if ("Teams" in report){
        for (let team in report["Teams"]) {
            if (report["Teams"].hasOwnProperty(team)) {
                data[report["Teams"][team]["Name"]] = {}
                for (let host in report.Teams[team]["Hosts"]){
                    if (report.Teams[team]["Hosts"].hasOwnProperty(host)) {
                        if (Object.keys(report.Teams[team]["Hosts"][host]["Services"]).length !== 0){
                            for (let service in report.Teams[team]["Hosts"][host]["Services"]) {
                                if (report.Teams[team]["Hosts"][host]["Services"].hasOwnProperty(service)) {
                                    let sr = report.Teams[team]["Hosts"][host]["Services"][service]
                                    let keyName = ""
                                    if (sr["DisplayName"]){
                                        keyName = sr["DisplayName"]
                                    } else {
                                        if (report.Teams[team]["Hosts"][host]["HostGroup"]){
                                            keyName = report.Teams[team]["Hosts"][host]["HostGroup"]["Name"] + "-" + sr["Name"]
                                        } else{
                                            keyName = sr["Name"]
                                        }
                                    }
                                    data[report["Teams"][team]["Name"]][keyName] = {...sr, Address: report["Teams"][team]["Hosts"][host].Address}
                                    dataKeys.add(keyName)
                                    teamNamesSet.add(report["Teams"][team]["Name"])
                                }
                            }
                        }
                    }
                }
            }
        }
    }
    const dataKeysArray = Array.from(dataKeys)
    const teamNames = Array.from(teamNamesSet)
    dataKeysArray.sort()
    teamNames.sort()
    return (
        <Paper className={classes.root}>
            <TableContainer>
                <Table stickyHeader aria-label="sticky table" size={dense ? 'small' : 'medium'}>
                    <TableHead>
                        <TableRow>
                            <TableCell
                                key="name"
                            >
                                Team Name
                            </TableCell>

                            {dataKeysArray.slice(columnPage * columnsPerPage, columnPage * columnsPerPage + columnsPerPage).map((column) => (
                                <TableCell align="center"
                                    key={column}
                                >
                                    {column}
                                </TableCell>
                            ))}
                        </TableRow>
                    </TableHead>


                    <TableBody>
                        {teamNames.slice(rowPage * rowsPerPage, rowPage * rowsPerPage + rowsPerPage).map((name) => {
                            return (
                                <TableRow hover tabIndex={-1} key={name}>
                                    <TableCell key={name}>
                                        {name}
                                    </TableCell>
                                    {dataKeysArray.slice(columnPage * columnsPerPage, columnPage * columnsPerPage + columnsPerPage).map((column) => (
                                        <TableCell key={name+column} style={(() => {
                                            if (data[name][column]) {
                                                if (data[name][column]["Passed"]){
                                                    return {backgroundColor: "green"}
                                                }
                                                return {backgroundColor: "red", color: "white"}
                                            }
                                        })()} align="center"
                                        >
                                            {!hideAddresses && data[name][column] && (() => {
                                                let msg = ""
                                                if (data[name][column]["Address"]) {
                                                    msg += data[name][column]["Address"]
                                                    if (column in data[name] && "Properties" in data[name][column]){
                                                        Object.keys(data[name][column]["Properties"]).forEach(key =>{
                                                            if (key === "Port"){
                                                                msg += ":" + data[name][column]["Properties"][key].Value
                                                            }
                                                        })
                                                    }
                                                }
                                                return msg
                                            })()}
                                        </TableCell>
                                    ))}
                                </TableRow>
                            );
                        })}
                        
                    </TableBody>
                    <TableFooter>
                        <TableRow>
                            <TableCell colSpan={dataKeysArray.length+1}>
                                <div style={{display:"flex", flexDirection: "row", justifyContent: "center", alignItems: "center"}}>
                                <TablePagination className={classes.tableNavigator}
                                                 rowsPerPageOptions={[1, 5, 10, 25, 100]}
                                                 component="div"
                                                 count={teamNames.length}
                                                 rowsPerPage={rowsPerPage}
                                                 page={rowPage}
                                                 onChangePage={handleRowChangePage}
                                                 onChangeRowsPerPage={handleChangeRowsPerPage}
                                />
                                <FormControlLabel className={classes.tableNavigator}
                                                  control={<Switch checked={dense} onChange={toggleChangeDense} />}
                                                  label="Dense padding"
                                />
                                { (token.isAValidToken() || props.currentPolicy.showAddresses?.value) &&
                                <FormControlLabel className={classes.tableNavigator}
                                                  control={<Switch checked={hideAddresses} onChange={toggleHideAddresses} />}
                                                  label={"Hide Addresses"}
                                />
                                }
                                <TablePagination
                                    labelRowsPerPage="Columns per page"
                                    rowsPerPageOptions={[1, 5, 10, 25, 100]}
                                    component="div"
                                    count={dataKeysArray.length}
                                    rowsPerPage={columnsPerPage}
                                    page={columnPage}
                                    className={classes.tableNavigator}
                                    onChangePage={handleColumnChangePage}
                                    onChangeRowsPerPage={handleChangeColumnsPerPage}
                                />
                                </div>
                            </TableCell>
                        </TableRow>
                    </TableFooter>
                </Table>
            </TableContainer>
        </Paper>
    );
}