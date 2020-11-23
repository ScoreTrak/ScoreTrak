import React from "react";
import {Route} from "react-router-dom";
import Box from "@material-ui/core/Box";
import RoundsMenu from "./Rounds/RoundsMenu";
import UserMenu from "./User/UserMenu";
import {SetupProps} from "./util/util";
import {centerStyle} from "../../styles/styles";
import TeamMenuTable from "./Team/TeamMenu";
import {ServiceGroupsMenu} from "./ServiceGroup/ServiceGroupMenu";
import ServiceMenu from "./Service/ServiceMenu";
import HostGroupsMenu from "./HostGroup/HostGroupMenu";
import HostMenu from "./Host/HostMenu";
import PropertiesMenu from "./Property/PropertiesMenu";

export default function Setup(props: SetupProps) {

    const classes = centerStyle();

    return (
        <Box m="auto" className={classes.alignItemsAndJustifyContent} height='85vh' width="100%" >
            <link
                rel="stylesheet"
                href="https://fonts.googleapis.com/icon?family=Material+Icons"
            />
           <Route exact path="/setup/teams" render={() => (
               <TeamMenuTable {...props} />
           )} />
           <Route exact path="/setup/host_groups" render={() => (
               <HostGroupsMenu {...props} />
           )} />
           <Route exact path="/setup/users" render={() => (
               <UserMenu {...props} />
           )} />
           <Route exact path="/setup/hosts" render={() => (
               <HostMenu {...props} />
           )} />
           <Route exact path="/setup/services" render={() => (
               <ServiceMenu {...props} />
           )} />
           <Route exact path="/setup/properties" render={() => (
               <PropertiesMenu {...props}/>
           )} />
           <Route exact path="/setup/rounds" render={() => (
               <RoundsMenu {...props}/>
           )} />
           <Route exact path="/setup/service_groups" render={() => (
               <ServiceGroupsMenu {...props}/>
           )} />
        </Box>
    );
}