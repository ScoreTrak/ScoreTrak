
import './App.css';
import React, {useEffect, useState} from 'react';
import './App.css';
import {
  BrowserRouter as Router
} from 'react-router-dom';
import {createMuiTheme, ThemeProvider} from "@material-ui/core/styles";
import {deepOrange, deepPurple, lightBlue, orange} from "@material-ui/core/colors";
import CssBaseline from "@material-ui/core/CssBaseline";
import {gRPCClients} from "./grpc/gRPCClients";
import Dashboard from "./components/Dashboard/Dashboard";
import {SnackbarProvider} from "notistack";

function App() {
  useEffect(() => {
    document.title = "ScoreTrak"
  }, []);
  if (localStorage.getItem("theme") !== "light"){
    localStorage.setItem("theme", "dark")
  }
  const [isDarkTheme, setIsDarkTheme] = useState<boolean>(localStorage.getItem("theme") === "dark");
  const palletType = isDarkTheme ? "dark" : "light";
  const mainPrimaryColor = isDarkTheme ? orange[500] : lightBlue[500];
  const mainSecondaryColor = isDarkTheme ? deepOrange[900] : deepPurple[500];
  const darkTheme = createMuiTheme({
    palette: {
      type: palletType,
      primary: {
        main: mainPrimaryColor
      },
      secondary: {
        main: mainSecondaryColor
      }
    }
  });

  return (
    <div className="App">
      <ThemeProvider theme={darkTheme}>
        <SnackbarProvider maxSnack={3} anchorOrigin={{
          vertical: 'bottom',
          horizontal: 'right',
        }} dense preventDuplicate>
        <CssBaseline />
          <Router>
            <Dashboard theme={{isDarkTheme, setIsDarkTheme}}  gRPCClients={gRPCClients} />
          </Router>
        </SnackbarProvider>
      </ThemeProvider>
    </div>
  );
}

export default App;
