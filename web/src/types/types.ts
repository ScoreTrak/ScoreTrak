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

export type SimpleCheck = {
    Passed:             boolean
    Log:                string
    Err:                string
}

export type SimpleService = {
    Check:              SimpleCheck | undefined
    Pause:              boolean
    Hide:               boolean,
    Name:               string,
    DisplayName:        string,
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
    HostGroup: SimpleHostGroup | undefined,
    Address:   string,
    Services:  Record<string, SimpleService>,
    Pause:     boolean,
    Hide:      boolean,
}

export type SimpleTeam = {
    Hosts:   Record<string, SimpleHost>,
    Name:    string,
    Pause:   boolean,
    Hide:    boolean,
    TotalPoints: number
}

export type SimpleHostGroup = {
    ID:      string,
    Name:    string,
    Pause:   boolean,
    Hide:    boolean,
}

export type SimpleReport = {
    Round: number,
    Teams: Record<string, SimpleTeam> ,
}