import type { AxiosInstance, AxiosRequestConfig } from "axios";
import { useQuery, useMutation, useQueryClient, type QueryClient, type UseMutationOptions, type UseQueryOptions, type MutationFunction, type UseMutationResult, type UseQueryResult } from "@tanstack/react-query";
export type Check = {
    id: string;
    pause?: boolean;
    hidden?: boolean;
    log: string;
    error: string;
    passed: boolean;
    round_id: string;
    host_service_id: string;
    team_id: string;
    rounds: Round;
    hostservice: HostService;
    team: Team;
};
export type CheckCreate = {
    id: string;
    pause?: boolean;
    hidden?: boolean;
    log: string;
    error: string;
    passed: boolean;
    round_id: string;
    host_service_id: string;
    team_id: string;
};
export type CheckList = {
    id: string;
    pause?: boolean;
    hidden?: boolean;
    log: string;
    error: string;
    passed: boolean;
    round_id: string;
    host_service_id: string;
    team_id: string;
};
export type CheckRead = {
    id: string;
    pause?: boolean;
    hidden?: boolean;
    log: string;
    error: string;
    passed: boolean;
    round_id: string;
    host_service_id: string;
    team_id: string;
};
export type CheckUpdate = {
    id: string;
    pause?: boolean;
    hidden?: boolean;
    log: string;
    error: string;
    passed: boolean;
    round_id: string;
    host_service_id: string;
    team_id: string;
};
export type CheckHostserviceRead = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    weight: number;
    point_boost: number;
    round_units: number;
    round_delay: number;
    host_id: string;
    team_id: string;
};
export type CheckRoundsRead = {
    id: string;
    round_number: number;
    note: string;
    err: string;
    started_at: string;
    finished_at: string;
    competition_id: string;
};
export type CheckTeamRead = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    number: number;
    competition_id: string;
};
export type Competition = {
    id: string;
    hidden?: boolean;
    pause?: boolean;
    name: string;
    display_name: string;
    viewable_to_public?: boolean;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
    teams?: Team[];
    services?: Service[];
    reports?: Report[];
    rounds?: Round[];
};
export type CompetitionCreate = {
    id: string;
    hidden?: boolean;
    pause?: boolean;
    name: string;
    display_name: string;
    viewable_to_public?: boolean;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type CompetitionList = {
    id: string;
    hidden?: boolean;
    pause?: boolean;
    name: string;
    display_name: string;
    viewable_to_public?: boolean;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type CompetitionRead = {
    id: string;
    hidden?: boolean;
    pause?: boolean;
    name: string;
    display_name: string;
    viewable_to_public?: boolean;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type CompetitionUpdate = {
    id: string;
    hidden?: boolean;
    pause?: boolean;
    name: string;
    display_name: string;
    viewable_to_public?: boolean;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type CompetitionReportsList = {
    id: number;
    log: string;
    error: string;
    competition_id: string;
};
export type CompetitionRoundsList = {
    id: string;
    round_number: number;
    note: string;
    err: string;
    started_at: string;
    finished_at: string;
    competition_id: string;
};
export type CompetitionServicesList = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    competition_id: string;
};
export type CompetitionTeamsList = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    number: number;
    competition_id: string;
};
export type Host = {
    id: string;
    pause?: boolean;
    hidden?: boolean;
    address: string;
    team_id: string;
    hostservices?: HostService[];
    team: Team;
};
export type HostCreate = {
    id: string;
    pause?: boolean;
    hidden?: boolean;
    address: string;
    team_id: string;
};
export type HostList = {
    id: string;
    pause?: boolean;
    hidden?: boolean;
    address: string;
    team_id: string;
};
export type HostRead = {
    id: string;
    pause?: boolean;
    hidden?: boolean;
    address: string;
    team_id: string;
};
export type HostService = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    weight: number;
    point_boost: number;
    round_units: number;
    round_delay: number;
    host_id: string;
    team_id: string;
    host: Host;
    checks?: Check[];
    properties?: Property[];
    team: Team;
};
export type HostServiceCreate = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    weight: number;
    point_boost: number;
    round_units: number;
    round_delay: number;
    host_id: string;
    team_id: string;
};
export type HostServiceList = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    weight: number;
    point_boost: number;
    round_units: number;
    round_delay: number;
    host_id: string;
    team_id: string;
};
export type HostServiceRead = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    weight: number;
    point_boost: number;
    round_units: number;
    round_delay: number;
    host_id: string;
    team_id: string;
};
export type HostServiceUpdate = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    weight: number;
    point_boost: number;
    round_units: number;
    round_delay: number;
    host_id: string;
    team_id: string;
};
export type HostServiceChecksList = {
    id: string;
    pause?: boolean;
    hidden?: boolean;
    log: string;
    error: string;
    passed: boolean;
    round_id: string;
    host_service_id: string;
    team_id: string;
};
export type HostServiceHostRead = {
    id: string;
    pause?: boolean;
    hidden?: boolean;
    address: string;
    team_id: string;
};
export type HostServicePropertiesList = {
    id: string;
    key: string;
    value: string;
    status: "view" | "edit" | "hide";
    host_service_id: string;
    team_id: string;
};
export type HostServiceTeamRead = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    number: number;
    competition_id: string;
};
export type HostUpdate = {
    id: string;
    pause?: boolean;
    hidden?: boolean;
    address: string;
    team_id: string;
};
export type HostHostservicesList = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    weight: number;
    point_boost: number;
    round_units: number;
    round_delay: number;
    host_id: string;
    team_id: string;
};
export type HostTeamRead = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    number: number;
    competition_id: string;
};
export type Property = {
    id: string;
    key: string;
    value: string;
    status: "view" | "edit" | "hide";
    host_service_id: string;
    team_id: string;
    hostservice: HostService;
    team: Team;
};
export type PropertyCreate = {
    id: string;
    key: string;
    value: string;
    status: "view" | "edit" | "hide";
    host_service_id: string;
    team_id: string;
};
export type PropertyList = {
    id: string;
    key: string;
    value: string;
    status: "view" | "edit" | "hide";
    host_service_id: string;
    team_id: string;
};
export type PropertyRead = {
    id: string;
    key: string;
    value: string;
    status: "view" | "edit" | "hide";
    host_service_id: string;
    team_id: string;
};
export type PropertyUpdate = {
    id: string;
    key: string;
    value: string;
    status: "view" | "edit" | "hide";
    host_service_id: string;
    team_id: string;
};
export type PropertyHostserviceRead = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    weight: number;
    point_boost: number;
    round_units: number;
    round_delay: number;
    host_id: string;
    team_id: string;
};
export type PropertyTeamRead = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    number: number;
    competition_id: string;
};
export type Report = {
    id: number;
    log: string;
    error: string;
    competition_id: string;
    competition: Competition;
};
export type ReportCreate = {
    id: number;
    log: string;
    error: string;
    competition_id: string;
};
export type ReportList = {
    id: number;
    log: string;
    error: string;
    competition_id: string;
};
export type ReportRead = {
    id: number;
    log: string;
    error: string;
    competition_id: string;
};
export type ReportUpdate = {
    id: number;
    log: string;
    error: string;
    competition_id: string;
};
export type ReportCompetitionRead = {
    id: string;
    hidden?: boolean;
    pause?: boolean;
    name: string;
    display_name: string;
    viewable_to_public?: boolean;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type Round = {
    id: string;
    round_number: number;
    note: string;
    err: string;
    started_at: string;
    finished_at: string;
    competition_id: string;
    checks?: Check[];
    competition: Competition;
};
export type RoundCreate = {
    id: string;
    round_number: number;
    note: string;
    err: string;
    started_at: string;
    finished_at: string;
    competition_id: string;
};
export type RoundList = {
    id: string;
    round_number: number;
    note: string;
    err: string;
    started_at: string;
    finished_at: string;
    competition_id: string;
};
export type RoundRead = {
    id: string;
    round_number: number;
    note: string;
    err: string;
    started_at: string;
    finished_at: string;
    competition_id: string;
};
export type RoundUpdate = {
    id: string;
    round_number: number;
    note: string;
    err: string;
    started_at: string;
    finished_at: string;
    competition_id: string;
};
export type RoundChecksList = {
    id: string;
    pause?: boolean;
    hidden?: boolean;
    log: string;
    error: string;
    passed: boolean;
    round_id: string;
    host_service_id: string;
    team_id: string;
};
export type RoundCompetitionRead = {
    id: string;
    hidden?: boolean;
    pause?: boolean;
    name: string;
    display_name: string;
    viewable_to_public?: boolean;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type Service = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    competition_id: string;
    competition: Competition;
};
export type ServiceCreate = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    competition_id: string;
};
export type ServiceList = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    competition_id: string;
};
export type ServiceRead = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    competition_id: string;
};
export type ServiceUpdate = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    competition_id: string;
};
export type ServiceCompetitionRead = {
    id: string;
    hidden?: boolean;
    pause?: boolean;
    name: string;
    display_name: string;
    viewable_to_public?: boolean;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type Team = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    number: number;
    competition_id: string;
    hosts?: Host[];
    hostservices?: HostService[];
    checks?: Check[];
    properties?: Property[];
    competition: Competition;
};
export type TeamCreate = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    number: number;
    competition_id: string;
};
export type TeamList = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    number: number;
    competition_id: string;
};
export type TeamRead = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    number: number;
    competition_id: string;
};
export type TeamUpdate = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    number: number;
    competition_id: string;
};
export type TeamChecksList = {
    id: string;
    pause?: boolean;
    hidden?: boolean;
    log: string;
    error: string;
    passed: boolean;
    round_id: string;
    host_service_id: string;
    team_id: string;
};
export type TeamCompetitionRead = {
    id: string;
    hidden?: boolean;
    pause?: boolean;
    name: string;
    display_name: string;
    viewable_to_public?: boolean;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type TeamHostsList = {
    id: string;
    pause?: boolean;
    hidden?: boolean;
    address: string;
    team_id: string;
};
export type TeamHostservicesList = {
    id: string;
    name: string;
    display_name: string;
    pause?: boolean;
    hidden?: boolean;
    weight: number;
    point_boost: number;
    round_units: number;
    round_delay: number;
    host_id: string;
    team_id: string;
};
export type TeamPropertiesList = {
    id: string;
    key: string;
    value: string;
    status: "view" | "edit" | "hide";
    host_service_id: string;
    team_id: string;
};
export type AxiosConfig = {
    paramsSerializer?: AxiosRequestConfig["paramsSerializer"];
};
export type Config = {
    mutations?: MutationConfigs;
    axios?: AxiosConfig;
};
export function initialize(axios: AxiosInstance, config?: Config) {
    const requests = makeRequests(axios, config?.axios);
    return {
        requests,
        queries: makeQueries(requests),
        mutations: makeMutations(requests, config?.mutations)
    };
}
function useRapiniMutation<TData = unknown, TError = unknown, TVariables = void, TContext = unknown>(mutationFn: MutationFunction<TData, TVariables>, config?: (queryClient: QueryClient) => Pick<UseMutationOptions<TData, TError, TVariables, TContext>, "onSuccess" | "onSettled" | "onError">, options?: Omit<UseMutationOptions<TData, TError, TVariables, TContext>, "mutationFn">): UseMutationResult<TData, TError, TVariables, TContext> {
    const { onSuccess, onError, onSettled, ...rest } = options ?? {};
    const queryClient = useQueryClient();
    const conf = config?.(queryClient);
    const mutationOptions: typeof options = {
        onSuccess: (data: TData, variables: TVariables, context?: TContext) => {
            conf?.onSuccess?.(data, variables, context);
            onSuccess?.(data, variables, context);
        },
        onError: (error: TError, variables: TVariables, context?: TContext) => {
            conf?.onError?.(error, variables, context);
            onError?.(error, variables, context);
        },
        onSettled: (data: TData | undefined, error: TError | null, variables: TVariables, context?: TContext) => {
            conf?.onSettled?.(data, error, variables, context);
            onSettled?.(data, error, variables, context);
        },
        ...rest
    };
    return useMutation({ mutationFn, ...mutationOptions });
}
function nullIfUndefined<T>(value: T): NonNullable<T> | null {
    return typeof value === "undefined" ? null : value as NonNullable<T> | null;
}
export const queryKeys = {
    listCheck: (page?: number, itemsPerPage?: number) => ["listCheck", nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readCheck: (id: string) => ["readCheck", id] as const,
    readCheckHostservice: (id: string) => ["readCheckHostservice", id] as const,
    readCheckRounds: (id: string) => ["readCheckRounds", id] as const,
    readCheckTeam: (id: string) => ["readCheckTeam", id] as const,
    listCompetition: (page?: number, itemsPerPage?: number) => ["listCompetition", nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readCompetition: (id: string) => ["readCompetition", id] as const,
    listCompetitionReports: (id: string, page?: number, itemsPerPage?: number) => ["listCompetitionReports", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    listCompetitionRounds: (id: string, page?: number, itemsPerPage?: number) => ["listCompetitionRounds", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    listCompetitionServices: (id: string, page?: number, itemsPerPage?: number) => ["listCompetitionServices", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    listCompetitionTeams: (id: string, page?: number, itemsPerPage?: number) => ["listCompetitionTeams", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    listHostService: (page?: number, itemsPerPage?: number) => ["listHostService", nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readHostService: (id: string) => ["readHostService", id] as const,
    listHostServiceChecks: (id: string, page?: number, itemsPerPage?: number) => ["listHostServiceChecks", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readHostServiceHost: (id: string) => ["readHostServiceHost", id] as const,
    listHostServiceProperties: (id: string, page?: number, itemsPerPage?: number) => ["listHostServiceProperties", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readHostServiceTeam: (id: string) => ["readHostServiceTeam", id] as const,
    listHost: (page?: number, itemsPerPage?: number) => ["listHost", nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readHost: (id: string) => ["readHost", id] as const,
    listHostHostservices: (id: string, page?: number, itemsPerPage?: number) => ["listHostHostservices", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readHostTeam: (id: string) => ["readHostTeam", id] as const,
    listProperty: (page?: number, itemsPerPage?: number) => ["listProperty", nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readProperty: (id: string) => ["readProperty", id] as const,
    readPropertyHostservice: (id: string) => ["readPropertyHostservice", id] as const,
    readPropertyTeam: (id: string) => ["readPropertyTeam", id] as const,
    listReport: (page?: number, itemsPerPage?: number) => ["listReport", nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readReport: (id: number) => ["readReport", id] as const,
    readReportCompetition: (id: number) => ["readReportCompetition", id] as const,
    listRound: (page?: number, itemsPerPage?: number) => ["listRound", nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readRound: (id: string) => ["readRound", id] as const,
    listRoundChecks: (id: string, page?: number, itemsPerPage?: number) => ["listRoundChecks", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readRoundCompetition: (id: string) => ["readRoundCompetition", id] as const,
    listService: (page?: number, itemsPerPage?: number) => ["listService", nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readService: (id: string) => ["readService", id] as const,
    readServiceCompetition: (id: string) => ["readServiceCompetition", id] as const,
    listTeam: (page?: number, itemsPerPage?: number) => ["listTeam", nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readTeam: (id: string) => ["readTeam", id] as const,
    listTeamChecks: (id: string, page?: number, itemsPerPage?: number) => ["listTeamChecks", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readTeamCompetition: (id: string) => ["readTeamCompetition", id] as const,
    listTeamHosts: (id: string, page?: number, itemsPerPage?: number) => ["listTeamHosts", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    listTeamHostservices: (id: string, page?: number, itemsPerPage?: number) => ["listTeamHostservices", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    listTeamProperties: (id: string, page?: number, itemsPerPage?: number) => ["listTeamProperties", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const
} as const;
export type QueryKeys = typeof queryKeys;
function makeRequests(axios: AxiosInstance, config?: AxiosConfig) {
    return {
        listCheck: (page?: number, itemsPerPage?: number) => axios.request<CheckList[]>({
            method: "get",
            url: `/checks`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        createCheck: (payload: {
            pause?: boolean;
            hidden?: boolean;
            log: string;
            error: string;
            passed: boolean;
            round_id: string;
            host_service_id: string;
            team_id: string;
            rounds: string;
            hostservice: string;
            team: string;
        }) => axios.request<CheckCreate>({
            method: "post",
            url: `/checks`,
            data: payload
        }).then(res => res.data),
        readCheck: (id: string) => axios.request<CheckRead>({
            method: "get",
            url: `/checks/${id}`
        }).then(res => res.data),
        deleteCheck: (id: string) => axios.request<unknown>({
            method: "delete",
            url: `/checks/${id}`
        }).then(res => res.data),
        updateCheck: (payload: {
            pause?: boolean;
            hidden?: boolean;
            log?: string;
            error?: string;
            passed?: boolean;
            round_id?: string;
            host_service_id?: string;
            rounds?: string;
            hostservice?: string;
            team?: string;
        }, id: string) => axios.request<CheckUpdate>({
            method: "patch",
            url: `/checks/${id}`,
            data: payload
        }).then(res => res.data),
        readCheckHostservice: (id: string) => axios.request<CheckHostserviceRead>({
            method: "get",
            url: `/checks/${id}/hostservice`
        }).then(res => res.data),
        readCheckRounds: (id: string) => axios.request<CheckRoundsRead>({
            method: "get",
            url: `/checks/${id}/rounds`
        }).then(res => res.data),
        readCheckTeam: (id: string) => axios.request<CheckTeamRead>({
            method: "get",
            url: `/checks/${id}/team`
        }).then(res => res.data),
        listCompetition: (page?: number, itemsPerPage?: number) => axios.request<CompetitionList[]>({
            method: "get",
            url: `/competitions`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        createCompetition: (payload: {
            hidden?: boolean;
            pause?: boolean;
            name: string;
            display_name: string;
            viewable_to_public?: boolean;
            to_be_started_at?: string;
            started_at?: string;
            finished_at?: string;
            teams?: string[];
            services?: string[];
            reports?: number[];
            rounds?: string[];
        }) => axios.request<CompetitionCreate>({
            method: "post",
            url: `/competitions`,
            data: payload
        }).then(res => res.data),
        readCompetition: (id: string) => axios.request<CompetitionRead>({
            method: "get",
            url: `/competitions/${id}`
        }).then(res => res.data),
        deleteCompetition: (id: string) => axios.request<unknown>({
            method: "delete",
            url: `/competitions/${id}`
        }).then(res => res.data),
        updateCompetition: (payload: {
            hidden?: boolean;
            pause?: boolean;
            name?: string;
            display_name?: string;
            viewable_to_public?: boolean;
            to_be_started_at?: string;
            started_at?: string;
            finished_at?: string;
            teams?: string[];
            services?: string[];
            reports?: number[];
            rounds?: string[];
        }, id: string) => axios.request<CompetitionUpdate>({
            method: "patch",
            url: `/competitions/${id}`,
            data: payload
        }).then(res => res.data),
        listCompetitionReports: (id: string, page?: number, itemsPerPage?: number) => axios.request<CompetitionReportsList[]>({
            method: "get",
            url: `/competitions/${id}/reports`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        listCompetitionRounds: (id: string, page?: number, itemsPerPage?: number) => axios.request<CompetitionRoundsList[]>({
            method: "get",
            url: `/competitions/${id}/rounds`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        listCompetitionServices: (id: string, page?: number, itemsPerPage?: number) => axios.request<CompetitionServicesList[]>({
            method: "get",
            url: `/competitions/${id}/services`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        listCompetitionTeams: (id: string, page?: number, itemsPerPage?: number) => axios.request<CompetitionTeamsList[]>({
            method: "get",
            url: `/competitions/${id}/teams`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        listHostService: (page?: number, itemsPerPage?: number) => axios.request<HostServiceList[]>({
            method: "get",
            url: `/host-services`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        createHostService: (payload: {
            name: string;
            display_name: string;
            pause?: boolean;
            hidden?: boolean;
            weight: number;
            point_boost: number;
            round_units: number;
            round_delay: number;
            host_id: string;
            team_id: string;
            host: string;
            checks?: string[];
            properties?: string[];
            team: string;
        }) => axios.request<HostServiceCreate>({
            method: "post",
            url: `/host-services`,
            data: payload
        }).then(res => res.data),
        readHostService: (id: string) => axios.request<HostServiceRead>({
            method: "get",
            url: `/host-services/${id}`
        }).then(res => res.data),
        deleteHostService: (id: string) => axios.request<unknown>({
            method: "delete",
            url: `/host-services/${id}`
        }).then(res => res.data),
        updateHostService: (payload: {
            name?: string;
            display_name?: string;
            pause?: boolean;
            hidden?: boolean;
            weight?: number;
            point_boost?: number;
            round_units?: number;
            round_delay?: number;
            host_id?: string;
            host?: string;
            checks?: string[];
            properties?: string[];
            team?: string;
        }, id: string) => axios.request<HostServiceUpdate>({
            method: "patch",
            url: `/host-services/${id}`,
            data: payload
        }).then(res => res.data),
        listHostServiceChecks: (id: string, page?: number, itemsPerPage?: number) => axios.request<HostServiceChecksList[]>({
            method: "get",
            url: `/host-services/${id}/checks`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        readHostServiceHost: (id: string) => axios.request<HostServiceHostRead>({
            method: "get",
            url: `/host-services/${id}/host`
        }).then(res => res.data),
        listHostServiceProperties: (id: string, page?: number, itemsPerPage?: number) => axios.request<HostServicePropertiesList[]>({
            method: "get",
            url: `/host-services/${id}/properties`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        readHostServiceTeam: (id: string) => axios.request<HostServiceTeamRead>({
            method: "get",
            url: `/host-services/${id}/team`
        }).then(res => res.data),
        listHost: (page?: number, itemsPerPage?: number) => axios.request<HostList[]>({
            method: "get",
            url: `/hosts`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        createHost: (payload: {
            pause?: boolean;
            hidden?: boolean;
            address: string;
            team_id: string;
            hostservices?: string[];
            team: string;
        }) => axios.request<HostCreate>({
            method: "post",
            url: `/hosts`,
            data: payload
        }).then(res => res.data),
        readHost: (id: string) => axios.request<HostRead>({
            method: "get",
            url: `/hosts/${id}`
        }).then(res => res.data),
        deleteHost: (id: string) => axios.request<unknown>({
            method: "delete",
            url: `/hosts/${id}`
        }).then(res => res.data),
        updateHost: (payload: {
            pause?: boolean;
            hidden?: boolean;
            address?: string;
            hostservices?: string[];
            team?: string;
        }, id: string) => axios.request<HostUpdate>({
            method: "patch",
            url: `/hosts/${id}`,
            data: payload
        }).then(res => res.data),
        listHostHostservices: (id: string, page?: number, itemsPerPage?: number) => axios.request<HostHostservicesList[]>({
            method: "get",
            url: `/hosts/${id}/hostservices`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        readHostTeam: (id: string) => axios.request<HostTeamRead>({
            method: "get",
            url: `/hosts/${id}/team`
        }).then(res => res.data),
        listProperty: (page?: number, itemsPerPage?: number) => axios.request<PropertyList[]>({
            method: "get",
            url: `/properties`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        createProperty: (payload: {
            key: string;
            value: string;
            status: "view" | "edit" | "hide";
            host_service_id: string;
            team_id: string;
            hostservice: string;
            team: string;
        }) => axios.request<PropertyCreate>({
            method: "post",
            url: `/properties`,
            data: payload
        }).then(res => res.data),
        readProperty: (id: string) => axios.request<PropertyRead>({
            method: "get",
            url: `/properties/${id}`
        }).then(res => res.data),
        deleteProperty: (id: string) => axios.request<unknown>({
            method: "delete",
            url: `/properties/${id}`
        }).then(res => res.data),
        updateProperty: (payload: {
            key?: string;
            value?: string;
            status?: "view" | "edit" | "hide";
            host_service_id?: string;
            hostservice?: string;
            team?: string;
        }, id: string) => axios.request<PropertyUpdate>({
            method: "patch",
            url: `/properties/${id}`,
            data: payload
        }).then(res => res.data),
        readPropertyHostservice: (id: string) => axios.request<PropertyHostserviceRead>({
            method: "get",
            url: `/properties/${id}/hostservice`
        }).then(res => res.data),
        readPropertyTeam: (id: string) => axios.request<PropertyTeamRead>({
            method: "get",
            url: `/properties/${id}/team`
        }).then(res => res.data),
        listReport: (page?: number, itemsPerPage?: number) => axios.request<ReportList[]>({
            method: "get",
            url: `/reports`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        createReport: (payload: {
            log: string;
            error: string;
            competition_id: string;
            competition: string;
        }) => axios.request<ReportCreate>({
            method: "post",
            url: `/reports`,
            data: payload
        }).then(res => res.data),
        readReport: (id: number) => axios.request<ReportRead>({
            method: "get",
            url: `/reports/${id}`
        }).then(res => res.data),
        deleteReport: (id: number) => axios.request<unknown>({
            method: "delete",
            url: `/reports/${id}`
        }).then(res => res.data),
        updateReport: (payload: {
            log?: string;
            error?: string;
            competition?: string;
        }, id: number) => axios.request<ReportUpdate>({
            method: "patch",
            url: `/reports/${id}`,
            data: payload
        }).then(res => res.data),
        readReportCompetition: (id: number) => axios.request<ReportCompetitionRead>({
            method: "get",
            url: `/reports/${id}/competition`
        }).then(res => res.data),
        listRound: (page?: number, itemsPerPage?: number) => axios.request<RoundList[]>({
            method: "get",
            url: `/rounds`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        createRound: (payload: {
            round_number: number;
            note: string;
            err: string;
            started_at: string;
            finished_at: string;
            competition_id: string;
            checks?: string[];
            competition: string;
        }) => axios.request<RoundCreate>({
            method: "post",
            url: `/rounds`,
            data: payload
        }).then(res => res.data),
        readRound: (id: string) => axios.request<RoundRead>({
            method: "get",
            url: `/rounds/${id}`
        }).then(res => res.data),
        deleteRound: (id: string) => axios.request<unknown>({
            method: "delete",
            url: `/rounds/${id}`
        }).then(res => res.data),
        updateRound: (payload: {
            round_number?: number;
            note?: string;
            err?: string;
            started_at?: string;
            finished_at?: string;
            checks?: string[];
            competition?: string;
        }, id: string) => axios.request<RoundUpdate>({
            method: "patch",
            url: `/rounds/${id}`,
            data: payload
        }).then(res => res.data),
        listRoundChecks: (id: string, page?: number, itemsPerPage?: number) => axios.request<RoundChecksList[]>({
            method: "get",
            url: `/rounds/${id}/checks`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        readRoundCompetition: (id: string) => axios.request<RoundCompetitionRead>({
            method: "get",
            url: `/rounds/${id}/competition`
        }).then(res => res.data),
        listService: (page?: number, itemsPerPage?: number) => axios.request<ServiceList[]>({
            method: "get",
            url: `/services`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        createService: (payload: {
            name: string;
            display_name: string;
            pause?: boolean;
            hidden?: boolean;
            competition_id: string;
            competition: string;
        }) => axios.request<ServiceCreate>({
            method: "post",
            url: `/services`,
            data: payload
        }).then(res => res.data),
        readService: (id: string) => axios.request<ServiceRead>({
            method: "get",
            url: `/services/${id}`
        }).then(res => res.data),
        deleteService: (id: string) => axios.request<unknown>({
            method: "delete",
            url: `/services/${id}`
        }).then(res => res.data),
        updateService: (payload: {
            name?: string;
            display_name?: string;
            pause?: boolean;
            hidden?: boolean;
            competition?: string;
        }, id: string) => axios.request<ServiceUpdate>({
            method: "patch",
            url: `/services/${id}`,
            data: payload
        }).then(res => res.data),
        readServiceCompetition: (id: string) => axios.request<ServiceCompetitionRead>({
            method: "get",
            url: `/services/${id}/competition`
        }).then(res => res.data),
        listTeam: (page?: number, itemsPerPage?: number) => axios.request<TeamList[]>({
            method: "get",
            url: `/teams`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        createTeam: (payload: {
            name: string;
            display_name: string;
            pause?: boolean;
            hidden?: boolean;
            number: number;
            competition_id: string;
            hosts?: string[];
            hostservices?: string[];
            checks?: string[];
            properties?: string[];
            competition: string;
        }) => axios.request<TeamCreate>({
            method: "post",
            url: `/teams`,
            data: payload
        }).then(res => res.data),
        readTeam: (id: string) => axios.request<TeamRead>({
            method: "get",
            url: `/teams/${id}`
        }).then(res => res.data),
        deleteTeam: (id: string) => axios.request<unknown>({
            method: "delete",
            url: `/teams/${id}`
        }).then(res => res.data),
        updateTeam: (payload: {
            name?: string;
            display_name?: string;
            pause?: boolean;
            hidden?: boolean;
            number?: number;
            hosts?: string[];
            hostservices?: string[];
            checks?: string[];
            properties?: string[];
            competition?: string;
        }, id: string) => axios.request<TeamUpdate>({
            method: "patch",
            url: `/teams/${id}`,
            data: payload
        }).then(res => res.data),
        listTeamChecks: (id: string, page?: number, itemsPerPage?: number) => axios.request<TeamChecksList[]>({
            method: "get",
            url: `/teams/${id}/checks`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        readTeamCompetition: (id: string) => axios.request<TeamCompetitionRead>({
            method: "get",
            url: `/teams/${id}/competition`
        }).then(res => res.data),
        listTeamHosts: (id: string, page?: number, itemsPerPage?: number) => axios.request<TeamHostsList[]>({
            method: "get",
            url: `/teams/${id}/hosts`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        listTeamHostservices: (id: string, page?: number, itemsPerPage?: number) => axios.request<TeamHostservicesList[]>({
            method: "get",
            url: `/teams/${id}/hostservices`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        listTeamProperties: (id: string, page?: number, itemsPerPage?: number) => axios.request<TeamPropertiesList[]>({
            method: "get",
            url: `/teams/${id}/properties`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data)
    } as const;
}
export type Requests = ReturnType<typeof makeRequests>;
export type Response<T extends keyof Requests> = Awaited<ReturnType<Requests[T]>>;
function makeQueries(requests: Requests) {
    return {
        useListCheck: (page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listCheck">, unknown, Response<"listCheck">, ReturnType<QueryKeys["listCheck"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listCheck">, unknown> => useQuery({ queryKey: queryKeys.listCheck(page, itemsPerPage), queryFn: () => requests.listCheck(page, itemsPerPage), ...options }),
        useReadCheck: (id: string, options?: Omit<UseQueryOptions<Response<"readCheck">, unknown, Response<"readCheck">, ReturnType<QueryKeys["readCheck"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readCheck">, unknown> => useQuery({ queryKey: queryKeys.readCheck(id), queryFn: () => requests.readCheck(id), ...options }),
        useReadCheckHostservice: (id: string, options?: Omit<UseQueryOptions<Response<"readCheckHostservice">, unknown, Response<"readCheckHostservice">, ReturnType<QueryKeys["readCheckHostservice"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readCheckHostservice">, unknown> => useQuery({ queryKey: queryKeys.readCheckHostservice(id), queryFn: () => requests.readCheckHostservice(id), ...options }),
        useReadCheckRounds: (id: string, options?: Omit<UseQueryOptions<Response<"readCheckRounds">, unknown, Response<"readCheckRounds">, ReturnType<QueryKeys["readCheckRounds"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readCheckRounds">, unknown> => useQuery({ queryKey: queryKeys.readCheckRounds(id), queryFn: () => requests.readCheckRounds(id), ...options }),
        useReadCheckTeam: (id: string, options?: Omit<UseQueryOptions<Response<"readCheckTeam">, unknown, Response<"readCheckTeam">, ReturnType<QueryKeys["readCheckTeam"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readCheckTeam">, unknown> => useQuery({ queryKey: queryKeys.readCheckTeam(id), queryFn: () => requests.readCheckTeam(id), ...options }),
        useListCompetition: (page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listCompetition">, unknown, Response<"listCompetition">, ReturnType<QueryKeys["listCompetition"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listCompetition">, unknown> => useQuery({ queryKey: queryKeys.listCompetition(page, itemsPerPage), queryFn: () => requests.listCompetition(page, itemsPerPage), ...options }),
        useReadCompetition: (id: string, options?: Omit<UseQueryOptions<Response<"readCompetition">, unknown, Response<"readCompetition">, ReturnType<QueryKeys["readCompetition"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readCompetition">, unknown> => useQuery({ queryKey: queryKeys.readCompetition(id), queryFn: () => requests.readCompetition(id), ...options }),
        useListCompetitionReports: (id: string, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listCompetitionReports">, unknown, Response<"listCompetitionReports">, ReturnType<QueryKeys["listCompetitionReports"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listCompetitionReports">, unknown> => useQuery({ queryKey: queryKeys.listCompetitionReports(id, page, itemsPerPage), queryFn: () => requests.listCompetitionReports(id, page, itemsPerPage), ...options }),
        useListCompetitionRounds: (id: string, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listCompetitionRounds">, unknown, Response<"listCompetitionRounds">, ReturnType<QueryKeys["listCompetitionRounds"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listCompetitionRounds">, unknown> => useQuery({ queryKey: queryKeys.listCompetitionRounds(id, page, itemsPerPage), queryFn: () => requests.listCompetitionRounds(id, page, itemsPerPage), ...options }),
        useListCompetitionServices: (id: string, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listCompetitionServices">, unknown, Response<"listCompetitionServices">, ReturnType<QueryKeys["listCompetitionServices"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listCompetitionServices">, unknown> => useQuery({ queryKey: queryKeys.listCompetitionServices(id, page, itemsPerPage), queryFn: () => requests.listCompetitionServices(id, page, itemsPerPage), ...options }),
        useListCompetitionTeams: (id: string, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listCompetitionTeams">, unknown, Response<"listCompetitionTeams">, ReturnType<QueryKeys["listCompetitionTeams"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listCompetitionTeams">, unknown> => useQuery({ queryKey: queryKeys.listCompetitionTeams(id, page, itemsPerPage), queryFn: () => requests.listCompetitionTeams(id, page, itemsPerPage), ...options }),
        useListHostService: (page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listHostService">, unknown, Response<"listHostService">, ReturnType<QueryKeys["listHostService"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listHostService">, unknown> => useQuery({ queryKey: queryKeys.listHostService(page, itemsPerPage), queryFn: () => requests.listHostService(page, itemsPerPage), ...options }),
        useReadHostService: (id: string, options?: Omit<UseQueryOptions<Response<"readHostService">, unknown, Response<"readHostService">, ReturnType<QueryKeys["readHostService"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readHostService">, unknown> => useQuery({ queryKey: queryKeys.readHostService(id), queryFn: () => requests.readHostService(id), ...options }),
        useListHostServiceChecks: (id: string, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listHostServiceChecks">, unknown, Response<"listHostServiceChecks">, ReturnType<QueryKeys["listHostServiceChecks"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listHostServiceChecks">, unknown> => useQuery({ queryKey: queryKeys.listHostServiceChecks(id, page, itemsPerPage), queryFn: () => requests.listHostServiceChecks(id, page, itemsPerPage), ...options }),
        useReadHostServiceHost: (id: string, options?: Omit<UseQueryOptions<Response<"readHostServiceHost">, unknown, Response<"readHostServiceHost">, ReturnType<QueryKeys["readHostServiceHost"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readHostServiceHost">, unknown> => useQuery({ queryKey: queryKeys.readHostServiceHost(id), queryFn: () => requests.readHostServiceHost(id), ...options }),
        useListHostServiceProperties: (id: string, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listHostServiceProperties">, unknown, Response<"listHostServiceProperties">, ReturnType<QueryKeys["listHostServiceProperties"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listHostServiceProperties">, unknown> => useQuery({ queryKey: queryKeys.listHostServiceProperties(id, page, itemsPerPage), queryFn: () => requests.listHostServiceProperties(id, page, itemsPerPage), ...options }),
        useReadHostServiceTeam: (id: string, options?: Omit<UseQueryOptions<Response<"readHostServiceTeam">, unknown, Response<"readHostServiceTeam">, ReturnType<QueryKeys["readHostServiceTeam"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readHostServiceTeam">, unknown> => useQuery({ queryKey: queryKeys.readHostServiceTeam(id), queryFn: () => requests.readHostServiceTeam(id), ...options }),
        useListHost: (page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listHost">, unknown, Response<"listHost">, ReturnType<QueryKeys["listHost"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listHost">, unknown> => useQuery({ queryKey: queryKeys.listHost(page, itemsPerPage), queryFn: () => requests.listHost(page, itemsPerPage), ...options }),
        useReadHost: (id: string, options?: Omit<UseQueryOptions<Response<"readHost">, unknown, Response<"readHost">, ReturnType<QueryKeys["readHost"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readHost">, unknown> => useQuery({ queryKey: queryKeys.readHost(id), queryFn: () => requests.readHost(id), ...options }),
        useListHostHostservices: (id: string, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listHostHostservices">, unknown, Response<"listHostHostservices">, ReturnType<QueryKeys["listHostHostservices"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listHostHostservices">, unknown> => useQuery({ queryKey: queryKeys.listHostHostservices(id, page, itemsPerPage), queryFn: () => requests.listHostHostservices(id, page, itemsPerPage), ...options }),
        useReadHostTeam: (id: string, options?: Omit<UseQueryOptions<Response<"readHostTeam">, unknown, Response<"readHostTeam">, ReturnType<QueryKeys["readHostTeam"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readHostTeam">, unknown> => useQuery({ queryKey: queryKeys.readHostTeam(id), queryFn: () => requests.readHostTeam(id), ...options }),
        useListProperty: (page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listProperty">, unknown, Response<"listProperty">, ReturnType<QueryKeys["listProperty"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listProperty">, unknown> => useQuery({ queryKey: queryKeys.listProperty(page, itemsPerPage), queryFn: () => requests.listProperty(page, itemsPerPage), ...options }),
        useReadProperty: (id: string, options?: Omit<UseQueryOptions<Response<"readProperty">, unknown, Response<"readProperty">, ReturnType<QueryKeys["readProperty"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readProperty">, unknown> => useQuery({ queryKey: queryKeys.readProperty(id), queryFn: () => requests.readProperty(id), ...options }),
        useReadPropertyHostservice: (id: string, options?: Omit<UseQueryOptions<Response<"readPropertyHostservice">, unknown, Response<"readPropertyHostservice">, ReturnType<QueryKeys["readPropertyHostservice"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readPropertyHostservice">, unknown> => useQuery({ queryKey: queryKeys.readPropertyHostservice(id), queryFn: () => requests.readPropertyHostservice(id), ...options }),
        useReadPropertyTeam: (id: string, options?: Omit<UseQueryOptions<Response<"readPropertyTeam">, unknown, Response<"readPropertyTeam">, ReturnType<QueryKeys["readPropertyTeam"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readPropertyTeam">, unknown> => useQuery({ queryKey: queryKeys.readPropertyTeam(id), queryFn: () => requests.readPropertyTeam(id), ...options }),
        useListReport: (page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listReport">, unknown, Response<"listReport">, ReturnType<QueryKeys["listReport"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listReport">, unknown> => useQuery({ queryKey: queryKeys.listReport(page, itemsPerPage), queryFn: () => requests.listReport(page, itemsPerPage), ...options }),
        useReadReport: (id: number, options?: Omit<UseQueryOptions<Response<"readReport">, unknown, Response<"readReport">, ReturnType<QueryKeys["readReport"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readReport">, unknown> => useQuery({ queryKey: queryKeys.readReport(id), queryFn: () => requests.readReport(id), ...options }),
        useReadReportCompetition: (id: number, options?: Omit<UseQueryOptions<Response<"readReportCompetition">, unknown, Response<"readReportCompetition">, ReturnType<QueryKeys["readReportCompetition"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readReportCompetition">, unknown> => useQuery({ queryKey: queryKeys.readReportCompetition(id), queryFn: () => requests.readReportCompetition(id), ...options }),
        useListRound: (page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listRound">, unknown, Response<"listRound">, ReturnType<QueryKeys["listRound"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listRound">, unknown> => useQuery({ queryKey: queryKeys.listRound(page, itemsPerPage), queryFn: () => requests.listRound(page, itemsPerPage), ...options }),
        useReadRound: (id: string, options?: Omit<UseQueryOptions<Response<"readRound">, unknown, Response<"readRound">, ReturnType<QueryKeys["readRound"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readRound">, unknown> => useQuery({ queryKey: queryKeys.readRound(id), queryFn: () => requests.readRound(id), ...options }),
        useListRoundChecks: (id: string, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listRoundChecks">, unknown, Response<"listRoundChecks">, ReturnType<QueryKeys["listRoundChecks"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listRoundChecks">, unknown> => useQuery({ queryKey: queryKeys.listRoundChecks(id, page, itemsPerPage), queryFn: () => requests.listRoundChecks(id, page, itemsPerPage), ...options }),
        useReadRoundCompetition: (id: string, options?: Omit<UseQueryOptions<Response<"readRoundCompetition">, unknown, Response<"readRoundCompetition">, ReturnType<QueryKeys["readRoundCompetition"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readRoundCompetition">, unknown> => useQuery({ queryKey: queryKeys.readRoundCompetition(id), queryFn: () => requests.readRoundCompetition(id), ...options }),
        useListService: (page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listService">, unknown, Response<"listService">, ReturnType<QueryKeys["listService"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listService">, unknown> => useQuery({ queryKey: queryKeys.listService(page, itemsPerPage), queryFn: () => requests.listService(page, itemsPerPage), ...options }),
        useReadService: (id: string, options?: Omit<UseQueryOptions<Response<"readService">, unknown, Response<"readService">, ReturnType<QueryKeys["readService"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readService">, unknown> => useQuery({ queryKey: queryKeys.readService(id), queryFn: () => requests.readService(id), ...options }),
        useReadServiceCompetition: (id: string, options?: Omit<UseQueryOptions<Response<"readServiceCompetition">, unknown, Response<"readServiceCompetition">, ReturnType<QueryKeys["readServiceCompetition"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readServiceCompetition">, unknown> => useQuery({ queryKey: queryKeys.readServiceCompetition(id), queryFn: () => requests.readServiceCompetition(id), ...options }),
        useListTeam: (page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listTeam">, unknown, Response<"listTeam">, ReturnType<QueryKeys["listTeam"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listTeam">, unknown> => useQuery({ queryKey: queryKeys.listTeam(page, itemsPerPage), queryFn: () => requests.listTeam(page, itemsPerPage), ...options }),
        useReadTeam: (id: string, options?: Omit<UseQueryOptions<Response<"readTeam">, unknown, Response<"readTeam">, ReturnType<QueryKeys["readTeam"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readTeam">, unknown> => useQuery({ queryKey: queryKeys.readTeam(id), queryFn: () => requests.readTeam(id), ...options }),
        useListTeamChecks: (id: string, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listTeamChecks">, unknown, Response<"listTeamChecks">, ReturnType<QueryKeys["listTeamChecks"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listTeamChecks">, unknown> => useQuery({ queryKey: queryKeys.listTeamChecks(id, page, itemsPerPage), queryFn: () => requests.listTeamChecks(id, page, itemsPerPage), ...options }),
        useReadTeamCompetition: (id: string, options?: Omit<UseQueryOptions<Response<"readTeamCompetition">, unknown, Response<"readTeamCompetition">, ReturnType<QueryKeys["readTeamCompetition"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readTeamCompetition">, unknown> => useQuery({ queryKey: queryKeys.readTeamCompetition(id), queryFn: () => requests.readTeamCompetition(id), ...options }),
        useListTeamHosts: (id: string, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listTeamHosts">, unknown, Response<"listTeamHosts">, ReturnType<QueryKeys["listTeamHosts"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listTeamHosts">, unknown> => useQuery({ queryKey: queryKeys.listTeamHosts(id, page, itemsPerPage), queryFn: () => requests.listTeamHosts(id, page, itemsPerPage), ...options }),
        useListTeamHostservices: (id: string, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listTeamHostservices">, unknown, Response<"listTeamHostservices">, ReturnType<QueryKeys["listTeamHostservices"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listTeamHostservices">, unknown> => useQuery({ queryKey: queryKeys.listTeamHostservices(id, page, itemsPerPage), queryFn: () => requests.listTeamHostservices(id, page, itemsPerPage), ...options }),
        useListTeamProperties: (id: string, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listTeamProperties">, unknown, Response<"listTeamProperties">, ReturnType<QueryKeys["listTeamProperties"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listTeamProperties">, unknown> => useQuery({ queryKey: queryKeys.listTeamProperties(id, page, itemsPerPage), queryFn: () => requests.listTeamProperties(id, page, itemsPerPage), ...options })
    } as const;
}
type MutationConfigs = {
    useCreateCheck?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"createCheck">, unknown, Parameters<Requests["createCheck"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useUpdateCheck?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"updateCheck">, unknown, Parameters<Requests["updateCheck"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useDeleteCheck?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"deleteCheck">, unknown, unknown, unknown>, "onSuccess" | "onSettled" | "onError">;
    useCreateCompetition?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"createCompetition">, unknown, Parameters<Requests["createCompetition"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useUpdateCompetition?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"updateCompetition">, unknown, Parameters<Requests["updateCompetition"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useDeleteCompetition?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"deleteCompetition">, unknown, unknown, unknown>, "onSuccess" | "onSettled" | "onError">;
    useCreateHostService?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"createHostService">, unknown, Parameters<Requests["createHostService"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useUpdateHostService?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"updateHostService">, unknown, Parameters<Requests["updateHostService"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useDeleteHostService?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"deleteHostService">, unknown, unknown, unknown>, "onSuccess" | "onSettled" | "onError">;
    useCreateHost?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"createHost">, unknown, Parameters<Requests["createHost"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useUpdateHost?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"updateHost">, unknown, Parameters<Requests["updateHost"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useDeleteHost?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"deleteHost">, unknown, unknown, unknown>, "onSuccess" | "onSettled" | "onError">;
    useCreateProperty?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"createProperty">, unknown, Parameters<Requests["createProperty"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useUpdateProperty?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"updateProperty">, unknown, Parameters<Requests["updateProperty"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useDeleteProperty?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"deleteProperty">, unknown, unknown, unknown>, "onSuccess" | "onSettled" | "onError">;
    useCreateReport?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"createReport">, unknown, Parameters<Requests["createReport"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useUpdateReport?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"updateReport">, unknown, Parameters<Requests["updateReport"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useDeleteReport?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"deleteReport">, unknown, unknown, unknown>, "onSuccess" | "onSettled" | "onError">;
    useCreateRound?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"createRound">, unknown, Parameters<Requests["createRound"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useUpdateRound?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"updateRound">, unknown, Parameters<Requests["updateRound"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useDeleteRound?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"deleteRound">, unknown, unknown, unknown>, "onSuccess" | "onSettled" | "onError">;
    useCreateService?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"createService">, unknown, Parameters<Requests["createService"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useUpdateService?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"updateService">, unknown, Parameters<Requests["updateService"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useDeleteService?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"deleteService">, unknown, unknown, unknown>, "onSuccess" | "onSettled" | "onError">;
    useCreateTeam?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"createTeam">, unknown, Parameters<Requests["createTeam"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useUpdateTeam?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"updateTeam">, unknown, Parameters<Requests["updateTeam"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useDeleteTeam?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"deleteTeam">, unknown, unknown, unknown>, "onSuccess" | "onSettled" | "onError">;
};
function makeMutations(requests: Requests, config?: Config["mutations"]) {
    return {
        useCreateCheck: (options?: Omit<UseMutationOptions<Response<"createCheck">, unknown, Parameters<Requests["createCheck"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"createCheck">, unknown, Parameters<Requests["createCheck"]>[0]>(payload => requests.createCheck(payload), config?.useCreateCheck, options),
        useUpdateCheck: (id: string, options?: Omit<UseMutationOptions<Response<"updateCheck">, unknown, Parameters<Requests["updateCheck"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"updateCheck">, unknown, Parameters<Requests["updateCheck"]>[0]>(payload => requests.updateCheck(payload, id), config?.useUpdateCheck, options),
        useDeleteCheck: (id: string, options?: Omit<UseMutationOptions<Response<"deleteCheck">, unknown, unknown, unknown>, "mutationFn">) => useRapiniMutation<Response<"deleteCheck">, unknown, unknown>(() => requests.deleteCheck(id), config?.useDeleteCheck, options),
        useCreateCompetition: (options?: Omit<UseMutationOptions<Response<"createCompetition">, unknown, Parameters<Requests["createCompetition"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"createCompetition">, unknown, Parameters<Requests["createCompetition"]>[0]>(payload => requests.createCompetition(payload), config?.useCreateCompetition, options),
        useUpdateCompetition: (id: string, options?: Omit<UseMutationOptions<Response<"updateCompetition">, unknown, Parameters<Requests["updateCompetition"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"updateCompetition">, unknown, Parameters<Requests["updateCompetition"]>[0]>(payload => requests.updateCompetition(payload, id), config?.useUpdateCompetition, options),
        useDeleteCompetition: (id: string, options?: Omit<UseMutationOptions<Response<"deleteCompetition">, unknown, unknown, unknown>, "mutationFn">) => useRapiniMutation<Response<"deleteCompetition">, unknown, unknown>(() => requests.deleteCompetition(id), config?.useDeleteCompetition, options),
        useCreateHostService: (options?: Omit<UseMutationOptions<Response<"createHostService">, unknown, Parameters<Requests["createHostService"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"createHostService">, unknown, Parameters<Requests["createHostService"]>[0]>(payload => requests.createHostService(payload), config?.useCreateHostService, options),
        useUpdateHostService: (id: string, options?: Omit<UseMutationOptions<Response<"updateHostService">, unknown, Parameters<Requests["updateHostService"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"updateHostService">, unknown, Parameters<Requests["updateHostService"]>[0]>(payload => requests.updateHostService(payload, id), config?.useUpdateHostService, options),
        useDeleteHostService: (id: string, options?: Omit<UseMutationOptions<Response<"deleteHostService">, unknown, unknown, unknown>, "mutationFn">) => useRapiniMutation<Response<"deleteHostService">, unknown, unknown>(() => requests.deleteHostService(id), config?.useDeleteHostService, options),
        useCreateHost: (options?: Omit<UseMutationOptions<Response<"createHost">, unknown, Parameters<Requests["createHost"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"createHost">, unknown, Parameters<Requests["createHost"]>[0]>(payload => requests.createHost(payload), config?.useCreateHost, options),
        useUpdateHost: (id: string, options?: Omit<UseMutationOptions<Response<"updateHost">, unknown, Parameters<Requests["updateHost"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"updateHost">, unknown, Parameters<Requests["updateHost"]>[0]>(payload => requests.updateHost(payload, id), config?.useUpdateHost, options),
        useDeleteHost: (id: string, options?: Omit<UseMutationOptions<Response<"deleteHost">, unknown, unknown, unknown>, "mutationFn">) => useRapiniMutation<Response<"deleteHost">, unknown, unknown>(() => requests.deleteHost(id), config?.useDeleteHost, options),
        useCreateProperty: (options?: Omit<UseMutationOptions<Response<"createProperty">, unknown, Parameters<Requests["createProperty"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"createProperty">, unknown, Parameters<Requests["createProperty"]>[0]>(payload => requests.createProperty(payload), config?.useCreateProperty, options),
        useUpdateProperty: (id: string, options?: Omit<UseMutationOptions<Response<"updateProperty">, unknown, Parameters<Requests["updateProperty"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"updateProperty">, unknown, Parameters<Requests["updateProperty"]>[0]>(payload => requests.updateProperty(payload, id), config?.useUpdateProperty, options),
        useDeleteProperty: (id: string, options?: Omit<UseMutationOptions<Response<"deleteProperty">, unknown, unknown, unknown>, "mutationFn">) => useRapiniMutation<Response<"deleteProperty">, unknown, unknown>(() => requests.deleteProperty(id), config?.useDeleteProperty, options),
        useCreateReport: (options?: Omit<UseMutationOptions<Response<"createReport">, unknown, Parameters<Requests["createReport"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"createReport">, unknown, Parameters<Requests["createReport"]>[0]>(payload => requests.createReport(payload), config?.useCreateReport, options),
        useUpdateReport: (id: number, options?: Omit<UseMutationOptions<Response<"updateReport">, unknown, Parameters<Requests["updateReport"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"updateReport">, unknown, Parameters<Requests["updateReport"]>[0]>(payload => requests.updateReport(payload, id), config?.useUpdateReport, options),
        useDeleteReport: (id: number, options?: Omit<UseMutationOptions<Response<"deleteReport">, unknown, unknown, unknown>, "mutationFn">) => useRapiniMutation<Response<"deleteReport">, unknown, unknown>(() => requests.deleteReport(id), config?.useDeleteReport, options),
        useCreateRound: (options?: Omit<UseMutationOptions<Response<"createRound">, unknown, Parameters<Requests["createRound"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"createRound">, unknown, Parameters<Requests["createRound"]>[0]>(payload => requests.createRound(payload), config?.useCreateRound, options),
        useUpdateRound: (id: string, options?: Omit<UseMutationOptions<Response<"updateRound">, unknown, Parameters<Requests["updateRound"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"updateRound">, unknown, Parameters<Requests["updateRound"]>[0]>(payload => requests.updateRound(payload, id), config?.useUpdateRound, options),
        useDeleteRound: (id: string, options?: Omit<UseMutationOptions<Response<"deleteRound">, unknown, unknown, unknown>, "mutationFn">) => useRapiniMutation<Response<"deleteRound">, unknown, unknown>(() => requests.deleteRound(id), config?.useDeleteRound, options),
        useCreateService: (options?: Omit<UseMutationOptions<Response<"createService">, unknown, Parameters<Requests["createService"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"createService">, unknown, Parameters<Requests["createService"]>[0]>(payload => requests.createService(payload), config?.useCreateService, options),
        useUpdateService: (id: string, options?: Omit<UseMutationOptions<Response<"updateService">, unknown, Parameters<Requests["updateService"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"updateService">, unknown, Parameters<Requests["updateService"]>[0]>(payload => requests.updateService(payload, id), config?.useUpdateService, options),
        useDeleteService: (id: string, options?: Omit<UseMutationOptions<Response<"deleteService">, unknown, unknown, unknown>, "mutationFn">) => useRapiniMutation<Response<"deleteService">, unknown, unknown>(() => requests.deleteService(id), config?.useDeleteService, options),
        useCreateTeam: (options?: Omit<UseMutationOptions<Response<"createTeam">, unknown, Parameters<Requests["createTeam"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"createTeam">, unknown, Parameters<Requests["createTeam"]>[0]>(payload => requests.createTeam(payload), config?.useCreateTeam, options),
        useUpdateTeam: (id: string, options?: Omit<UseMutationOptions<Response<"updateTeam">, unknown, Parameters<Requests["updateTeam"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"updateTeam">, unknown, Parameters<Requests["updateTeam"]>[0]>(payload => requests.updateTeam(payload, id), config?.useUpdateTeam, options),
        useDeleteTeam: (id: string, options?: Omit<UseMutationOptions<Response<"deleteTeam">, unknown, unknown, unknown>, "mutationFn">) => useRapiniMutation<Response<"deleteTeam">, unknown, unknown>(() => requests.deleteTeam(id), config?.useDeleteTeam, options)
    } as const;
}
