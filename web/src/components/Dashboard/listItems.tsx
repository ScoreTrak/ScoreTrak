import React from "react";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import SettingsIcon from '@material-ui/icons/Settings';
import PersonIcon from '@material-ui/icons/Person';
import GroupIcon from '@material-ui/icons/Group';
import { Link } from 'react-router-dom'
import DesktopWindowsIcon from '@material-ui/icons/DesktopWindows';
import SettingsApplicationsIcon from '@material-ui/icons/SettingsApplications';
import TextFieldsIcon from '@material-ui/icons/TextFields';
import FlipCameraAndroidIcon from '@material-ui/icons/FlipCameraAndroid';

export const adminListItems = (
    <div>
        <ListItem button component={Link} to="/settings">
            <ListItemIcon>
                <SettingsIcon/>
            </ListItemIcon>
            <ListItemText primary="Global Settings" />
        </ListItem>
        <ListItem button component={Link} to="/setup/teams">
            <ListItemIcon>
                <GroupIcon/>
            </ListItemIcon>
            <ListItemText primary="Manage Teams" />
        </ListItem>
        <ListItem button component={Link} to="/setup/users">
            <ListItemIcon>
                <PersonIcon/>
            </ListItemIcon>
            <ListItemText primary="Manage Users" />
        </ListItem>
        <ListItem button component={Link} to="/setup/host_groups">
            <ListItemIcon>
                <DesktopWindowsIcon/>
            </ListItemIcon>
            <ListItemText primary="Manage Host Groups" />
        </ListItem>
        <ListItem button component={Link} to="/setup/hosts">
            <ListItemIcon>
                <DesktopWindowsIcon/>
            </ListItemIcon>
            <ListItemText primary="Manage Hosts" />
        </ListItem>
        <ListItem button component={Link} to="/setup/service_groups">
            <ListItemIcon>
                <SettingsApplicationsIcon/>
            </ListItemIcon>
            <ListItemText primary="Manage Service Groups" />
        </ListItem>

        <ListItem button component={Link} to="/setup/services">
            <ListItemIcon>
                <SettingsApplicationsIcon/>
            </ListItemIcon>
            <ListItemText primary="Manage Services" />
        </ListItem>

        <ListItem button component={Link} to="/setup/properties">
            <ListItemIcon>
                <TextFieldsIcon/>
            </ListItemIcon>
            <ListItemText primary="Manage Properties" />
        </ListItem>
        <ListItem button component={Link} to="/setup/rounds">
            <ListItemIcon>
                <FlipCameraAndroidIcon/>
            </ListItemIcon>
            <ListItemText primary="Show Rounds Logs" />
        </ListItem>
    </div>
);