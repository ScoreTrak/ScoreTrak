import React, {useEffect} from 'react';
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
import {SimpleCheck, SimpleReport, SimpleService} from "../../types/types";
import {Policy} from "../../grpc/pkg/policy/policypb/policy_pb";
import {token} from "../../grpc/token/token";


const useStyles = makeStyles({
    root: {
        width: '100%',
        height: '100%'
    },
    tableNavigator: {
        marginRight: "3vh",
        marginLeft: "3vh"
    }
});

type RanksProps = {
    report: SimpleReport
    isDarkTheme: boolean
    currentPolicy: Policy.AsObject,
}

export default function Status(props: RanksProps) {
    useEffect(() => {
        document.title = "Status"
    }, []);
    const classes = useStyles();
    const [rowPage, setRowPage] = React.useState<number>(0);
    const [rowsPerPage, setRowsPerPage] = React.useState<number>(25);
    const [dense, setDense] = React.useState<boolean>(true);
    const [hideAddresses, setHideAddresses] = React.useState<boolean>(true);
    const [highlightParentTeam, setHighlightParentTeam] = React.useState<boolean>(true);

    const toggleHideAddresses = () => {
        setHideAddresses(prevState => !prevState)
    };

    const toggleHighlightParentTeam = () => {
        setHighlightParentTeam(prevState => !prevState)
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
    const teamNamesSet = new Set<string>();
    const data: Record<string, Record<string, SimpleService & {Address: string}>> = {}
    const dataKeys = new Set<string>();
    if ("Teams" in report){
        for (const team in report.Teams) {
            if (report.Teams.hasOwnProperty(team)) {
                data[report.Teams[team].Name] = {}
                for (const host in report.Teams[team].Hosts){
                    if (report.Teams[team].Hosts.hasOwnProperty(host)) {
                        if (Object.keys(report.Teams[team].Hosts[host].Services).length !== 0){
                            for (const service in report.Teams[team].Hosts[host].Services) {
                                if (report.Teams[team].Hosts[host].Services.hasOwnProperty(service)) {
                                    const hst = report.Teams[team].Hosts[host]
                                    const sr = hst.Services[service]
                                    let keyName = ""
                                    if (sr.DisplayName){
                                        keyName = sr.DisplayName
                                    } else {
                                        if (hst.HostGroup !== undefined){
                                            keyName = hst.HostGroup.Name + "-" + sr.Name
                                        } else{
                                            keyName = sr.Name
                                        }
                                    }

                                    data[report.Teams[team].Name][keyName] = {...sr, Address: report.Teams[team].Hosts[host].Address,
                                        Pause: report.Teams[team].Pause || (hst.HostGroup !== undefined ? hst.HostGroup.Pause : false) || hst.Pause || sr.Pause
                                    }



                                    dataKeys.add(keyName)
                                    teamNamesSet.add(report.Teams[team].Name)
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
            <TableContainer>
                <Table stickyHeader aria-label="sticky table" size={dense ? 'small' : 'medium'}>
                    <TableHead>
                        <TableRow>
                            <TableCell>
                                Team Name
                            </TableCell>

                            {dataKeysArray.slice(columnPage * columnsPerPage, columnPage * columnsPerPage + columnsPerPage).map((column) => (
                                <TableCell width={`${100/(dataKeysArray.slice(columnPage * columnsPerPage, columnPage * columnsPerPage + columnsPerPage).length)}%`}
                                           align="center"
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
                                    <TableCell style={{
                                        'whiteSpace': 'nowrap',
                                        'overflow': 'hidden',
                                        'textOverflow': 'ellipsis'}}>
                                        {name}
                                    </TableCell>
                                    {dataKeysArray.slice(columnPage * columnsPerPage, columnPage * columnsPerPage + columnsPerPage).map((column) => (
                                        <TableCell key={name + column} width={`${100/(dataKeysArray.slice(columnPage * columnsPerPage, columnPage * columnsPerPage + columnsPerPage).length)}%`} style={(() => {
                                            if (data[name][column]) {
                                                let style = {}
                                                if (props.isDarkTheme){
                                                    if (data[name][column].Pause){
                                                        style = {backgroundColor: "#000000"}
                                                    } else if (data[name][column].Check !== undefined && data[name][column].Check?.Passed){
                                                        style =  {backgroundColor: "#259B0B"}
                                                    } else{
                                                        style = {backgroundColor: "#d20c23", color: "white"}
                                                    }
                                                } else{
                                                    if (data[name][column].Pause){
                                                        style = {backgroundColor: "#000000"}
                                                    } else if (data[name][column].Check !== undefined && data[name][column].Check?.Passed){
                                                         style =  {backgroundColor: "green"}
                                                    } else{
                                                        style = {backgroundColor: "red", color: "white"}
                                                    }
                                                }
                                                const teamId = token.getCurrentTeamID()

                                                if (token.isAValidToken() && teamId !== undefined && teamId in report.Teams && report.Teams[teamId].Name === name && highlightParentTeam) {
                                                    style = {
                                                        ...style,
                                                        borderTop: '2px solid rgba(0, 0, 0, 1)',
                                                        borderBottom: '2px solid rgba(0, 0, 0, 1)',
                                                        borderLeft: '1px solid rgba(0, 0, 0, 0.5)',
                                                        borderRight: '1px solid rgba(0, 0, 0, 0.5)',

                                                    }
                                                } else{
                                                    style = {...style, border: '1px solid rgba(0, 0, 0, 0.5)'}
                                                }

                                                return style
                                            }
                                        })()} align="center" padding="none"
                                        >
                                            {!hideAddresses && data[name][column] && (() => {
                                                let msg = ""
                                                if (data[name][column].Address) {
                                                    msg += data[name][column].Address
                                                    if (column in data[name] && "Properties" in data[name][column]){
                                                        Object.keys(data[name][column].Properties).forEach(key => {
                                                            if (key === "Port"){
                                                                msg += ":" + data[name][column].Properties[key].Value
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
                            <TableCell colSpan={dataKeysArray.length + 1}>
                                <div style={{display: "flex", flexDirection: "row", justifyContent: "center", alignItems: "center"}}>
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
                                { (token.isAValidToken()) &&
                                <FormControlLabel className={classes.tableNavigator}
                                                  control={<Switch checked={highlightParentTeam} onChange={toggleHighlightParentTeam} />}
                                                  label={"Highlight Team Cells"}
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
    );
}