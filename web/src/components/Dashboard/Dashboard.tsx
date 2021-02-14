import React, {useEffect, useState} from "react";
import clsx from "clsx";
import {makeStyles} from "@material-ui/core/styles";
import CssBaseline from "@material-ui/core/CssBaseline";
import Switch from "@material-ui/core/Switch";
import Drawer from "@material-ui/core/Drawer";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import List from "@material-ui/core/List";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import IconButton from "@material-ui/core/IconButton";
import Container from "@material-ui/core/Container";
import MenuIcon from "@material-ui/icons/Menu";
import ChevronLeftIcon from "@material-ui/icons/ChevronLeft";
import {adminListItems} from "./listItems";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import {Link, Route, Switch as RouterSwitch, useHistory} from "react-router-dom";
import AssignmentIndIcon from '@material-ui/icons/AssignmentInd';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import {FullScreen, useFullScreenHandle} from 'react-full-screen';
import BarChartIcon from "@material-ui/icons/BarChart";
import CheckCircleIcon from "@material-ui/icons/CheckCircle";
import DetailsIcon from "@material-ui/icons/Details";
import Button from "@material-ui/core/Button";
import {Severity, ThemeState} from "../../types/types";
import {Role, token} from "../../grpc/token/token";
import {GRPCClients} from "../../grpc/gRPCClients";
import {GetRequest, GetResponse, Policy} from "../../grpc/pkg/policy/policypb/policy_pb";
import Login from "../Login/Login";
import ScoreBoard from "../ScoreBoard/ScoreBoard";
import grey from "@material-ui/core/colors/grey";
import {useSnackbar} from 'notistack';
import Setup from "../Setup/Setup";
import Settings from "../Settings/Settings";


const drawerWidth = 260;

const useStyles = makeStyles(theme => ({
  root: {
    display: "flex"
  },
  toolbar: {
    position: "relative",
    paddingRight: 24
  },
  toolbarIcon: {
    display: "flex",
    alignItems: "center",
    justifyContent: "flex-end",
    padding: "0 8px",
    ...theme.mixins.toolbar
  },
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
    transition: theme.transitions.create(["width", "margin"], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen
    })
  },
  appBarShift: {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(["width", "margin"], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen
    })
  },
  menuButton: {
    marginRight: 36
  },
  menuButtonHidden: {
    display: "none"
  },
  title: {
    flexGrow: 1
  },
  drawerPaper: {
    position: "relative",
    whiteSpace: "nowrap",
    width: drawerWidth,
    transition: theme.transitions.create("width", {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen
    })
  },
  drawerPaperClose: {
    overflowX: "hidden",
    transition: theme.transitions.create("width", {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen
    }),
    width: theme.spacing(7),
    [theme.breakpoints.up("sm")]: {
      width: theme.spacing(9)
    }
  },
  appBarSpacer: theme.mixins.toolbar,
  content: {
    flexGrow: 1,
    height: "100vh",
    overflow: "auto"
  },
  container: {
    paddingTop: theme.spacing(4),
    paddingBottom: theme.spacing(4),
  },
  paper: {
    padding: theme.spacing(2),
    display: "flex",
    overflow: "auto",
    flexDirection: "column"
  },
  fixedHeight: {
    height: '85vh'
  },
  fullSizeHeight: {
    height: '100vh'
  }
}));

interface DashboardProps{
  theme: ThemeState,
  gRPCClients: GRPCClients
}

export default function Dashboard(props: DashboardProps) {
  const [open, setOpen] = useState<boolean>(false);
  const [Title, setTitle] = useState<string>("ScoreBoard")
  const setIsDarkTheme = props.theme.setIsDarkTheme
  const isDarkTheme = props.theme.isDarkTheme
  const classes = useStyles();

  const { enqueueSnackbar, closeSnackbar } = useSnackbar();


  const handleThemeChange = () => {
    setIsDarkTheme(prevState => {
      if (!prevState){
        localStorage.setItem("theme", "dark")
      } else{
        localStorage.setItem("theme", "light")
      }
      return !prevState
    });
  };


  const action = (key: string) => (
      <React.Fragment>
        <Button variant="outlined" onClick={() => { closeSnackbar(key) }}>
          Dismiss
        </Button>
      </React.Fragment>
  );

  const genericEnqueue = (message: string, severity: Severity, autoHideDuration: number | null | undefined = null, uniqueID?: string) => {
    enqueueSnackbar(message, {
        variant: severity,
        autoHideDuration,
        key: uniqueID,
        action,
    })
  }

  const handleDrawerOpen = () => {
    setOpen(true);
  };
  const handleDrawerClose = () => {
    setOpen(false);
  };
  const logout = () => {
    token.logout()
    window.location.reload()
  }
  const handleFullScreen = useFullScreenHandle()
  const [currentPolicy, setPolicy] = useState<Policy.AsObject | undefined>(undefined);
  const history = useHistory();
  useEffect(() => {
    const streamRequest = new GetRequest();
    const stream = props.gRPCClients.policyClient.get(streamRequest, {})
    stream.on('error', (err: any) => {
      if (err.code === 7 || err.code === 16){
        genericEnqueue(`You are not authorized to perform this action. Please Log in`, Severity.Error)
        token.logout()
        history.push("/login");
      } else{
        genericEnqueue(`Encountered an error while fetching policy: ${err.message}. Error code: ${err.code}`, Severity.Error)
      }
    });
    stream.on("data", (response) => {
      setPolicy((response as GetResponse).getPolicy()?.toObject())
    });
    return () => stream.cancel();
  }, []);


  return (
      <div className={classes.root}>
        <CssBaseline />
        <RouterSwitch>
          <Route exact path="/login" render={() => (
              <Login authClient={props.gRPCClients.authClient}/>
          )} />
          {
            currentPolicy && <Route path="/" render={() => (
                <React.Fragment>
                  <AppBar
                      position="absolute"
                      className={clsx(classes.appBar, open && classes.appBarShift)}
                  >
                    <Toolbar className={classes.toolbar}>
                      <IconButton
                          edge="start"
                          color="inherit"
                          aria-label="open drawer"
                          onClick={handleDrawerOpen}
                          className={clsx(
                              classes.menuButton,
                              open && classes.menuButtonHidden
                          )}
                      >
                        <MenuIcon />
                      </IconButton>
                      <Typography
                          component="h1"
                          variant="h6"
                          color="inherit"
                          noWrap
                          className={classes.title}
                      >{Title}
                      </Typography>
                    </Toolbar>
                  </AppBar>
                  <Drawer
                      variant="permanent"
                      classes={{
                        paper: clsx(classes.drawerPaper, !open && classes.drawerPaperClose)
                      }}
                      open={open}>
                    <div className={classes.toolbarIcon}>
                      <IconButton onClick={handleDrawerClose}>
                        <ChevronLeftIcon />
                      </IconButton>
                    </div>
                    <Divider/>
                    <Switch checked={isDarkTheme} onChange={handleThemeChange} />
                    <Divider />
                    <List>
                      <div>
                        { (currentPolicy.showPoints?.value || token.getCurrentRole() === Role.Black) &&
                        <ListItem button component={Link} to="/ranks">
                          <ListItemIcon>
                            <BarChartIcon/>
                          </ListItemIcon>
                          <ListItemText primary="Ranks" />
                        </ListItem>
                        }
                        <ListItem button component={Link} to="/">
                          <ListItemIcon>
                            <CheckCircleIcon />
                          </ListItemIcon>
                          <ListItemText primary="Status" />
                        </ListItem>
                        { (token.getCurrentRole() === Role.Red || token.getCurrentRole() === Role.Blue || token.getCurrentRole() === Role.Black) &&
                        <ListItem button component={Link} to="/details">
                          <ListItemIcon>
                            <DetailsIcon />
                          </ListItemIcon>
                          <ListItemText primary="Details" />
                        </ListItem>
                        }
                      </div>
                    </List>
                    {
                      token.getCurrentRole() === Role.Black  &&
                      <List>
                        <Divider/>
                        {adminListItems}
                      </List>
                    }
                    <Divider/>
                    {
                      !token.isAValidToken() ?
                          <ListItem button component={Link} to="/login">
                            <ListItemIcon>
                              <AssignmentIndIcon />
                            </ListItemIcon>
                            <ListItemText primary="Sign In" />
                          </ListItem>
                          :
                          <ListItem button onClick={logout}>
                            <ListItemIcon>
                              <ExitToAppIcon />
                            </ListItemIcon>
                            <ListItemText primary="Sign Out" />
                          </ListItem>
                    }
                  </Drawer>
                  <main className={classes.content}>
                    <div className={classes.appBarSpacer} />
                    <Container maxWidth="xl" className={classes.container}>
                          <Route exact path={["/", "/ranks", "/details"]} render={() => (
                              <FullScreen handle={handleFullScreen}>
                                <div style={(handleFullScreen.active && ((!isDarkTheme && { background: grey[50]}) || { background: grey.A400})) || undefined}>
                                  <ScoreBoard {...props} genericEnqueue={genericEnqueue} currentPolicy={currentPolicy} setTitle={setTitle} handleFullScreen={handleFullScreen}/>
                                </div>
                              </FullScreen>
                          )} />
                          <Route exact path="/settings" render={() => (
                              <Settings isDarkTheme={props.theme.isDarkTheme} genericEnqueue={genericEnqueue} setTitle={setTitle}  gRPCClients={props.gRPCClients} currentPolicy={currentPolicy}/>
                          )} />
                          <Route path="/setup" render={() => (
                              <Setup isDarkTheme={props.theme.isDarkTheme} genericEnqueue={genericEnqueue} setTitle={setTitle}  gRPCClients={props.gRPCClients}  />
                          )} />
                    </Container>
                  </main>
                </React.Fragment>
            )} />
          }
        </RouterSwitch>
      </div>

  );
}