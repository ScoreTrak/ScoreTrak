import * as React from "react";

export enum Severity {
    Error = "error",
    Warning = "warning",
    Info = "info",
    Success = "success"
}

export interface ThemeState {
    isDarkTheme: boolean;
    setIsDarkTheme: React.Dispatch<React.SetStateAction<boolean>>;
}

export type SimpleProperty = {
    Value: string,
    Status: string
}

export type SimpleService = {
    Enabled:            boolean,
    Name:               string,
    DisplayName:        string,
    Passed:             boolean,
    Log:                string,
    Err:                string,
    Weight:             number,
    Points:             number,
    PointsBoost:        number,
    Properties:         Record<string, SimpleProperty>,
    SimpleServiceGroup: SimpleServiceGroup
}

export type SimpleServiceGroup = {
    ID:      string,
    Name:    string,
    Enabled: boolean
}

export type SimpleHost = {
    HostGroup: SimpleHostGroup,
    Address:   string,
    Services:  Record<string, SimpleService>,
    Enabled:   boolean
}

export type SimpleTeam = {
    Hosts:   Record<string, SimpleHost>,
    Name:    string,
    Enabled: boolean,
    Hidden:  boolean,
    TotalPoints: number
}

export type SimpleHostGroup = {
    ID:      string,
    Name:    string,
    Enabled: boolean
}

export type SimpleReport = {
    Round: number,
    Teams: Record<string, SimpleTeam> ,
}