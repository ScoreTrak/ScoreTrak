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