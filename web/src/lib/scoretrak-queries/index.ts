import type { AxiosInstance, AxiosRequestConfig } from "axios";
import { useQuery, useMutation, useQueryClient, type QueryClient, type UseMutationOptions, type UseQueryOptions, type MutationFunction, type UseMutationResult, type UseQueryResult } from "@tanstack/react-query";
export type Check = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    log: string;
    error: string;
    passed: boolean;
    competition: Competition;
    rounds: Round;
    services: Service;
};
export type CheckCreate = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    log: string;
    error: string;
    passed: boolean;
};
export type CheckList = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    log: string;
    error: string;
    passed: boolean;
};
export type CheckRead = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    log: string;
    error: string;
    passed: boolean;
};
export type CheckUpdate = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    log: string;
    error: string;
    passed: boolean;
};
export type CheckCompetitionRead = {
    id: number;
    create_time: string;
    update_time: string;
    hidden: boolean;
    pause: boolean;
    name: string;
    display_name: string;
    round_duration: number;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type CheckRoundsRead = {
    id: number;
    create_time: string;
    update_time: string;
    competition_id: number;
    round_number: number;
    note: string;
    err: string;
    started_at: string;
    finished_at: string;
};
export type CheckServicesRead = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    name: string;
    display_name: string;
    weight: number;
    point_boost: number;
    round_units: number;
    round_delay: number;
};
export type Competition = {
    id: number;
    create_time: string;
    update_time: string;
    hidden: boolean;
    pause: boolean;
    name: string;
    display_name: string;
    round_duration: number;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
    teams?: Team[];
    users?: User[];
};
export type CompetitionCreate = {
    id: number;
    create_time: string;
    update_time: string;
    hidden: boolean;
    pause: boolean;
    name: string;
    display_name: string;
    round_duration: number;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type CompetitionList = {
    id: number;
    create_time: string;
    update_time: string;
    hidden: boolean;
    pause: boolean;
    name: string;
    display_name: string;
    round_duration: number;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type CompetitionRead = {
    id: number;
    create_time: string;
    update_time: string;
    hidden: boolean;
    pause: boolean;
    name: string;
    display_name: string;
    round_duration: number;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type CompetitionUpdate = {
    id: number;
    create_time: string;
    update_time: string;
    hidden: boolean;
    pause: boolean;
    name: string;
    display_name: string;
    round_duration: number;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type CompetitionTeamsList = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    name: string;
    index: number;
};
export type CompetitionUsersList = {
    id: number;
    create_time: string;
    update_time: string;
    username: string;
};
export type Host = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    address: string;
    address_list_range: string;
    editable: boolean;
    competition: Competition;
    team: Team;
    services?: Service[];
    host_group: HostGroup;
};
export type HostCreate = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    address: string;
    address_list_range: string;
    editable: boolean;
};
export type HostGroup = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    name: string;
    competition: Competition;
    team: Team;
    hosts?: Host[];
};
export type HostGroupCreate = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    name: string;
};
export type HostGroupList = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    name: string;
};
export type HostGroupRead = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    name: string;
};
export type HostGroupUpdate = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    name: string;
};
export type HostGroupCompetitionRead = {
    id: number;
    create_time: string;
    update_time: string;
    hidden: boolean;
    pause: boolean;
    name: string;
    display_name: string;
    round_duration: number;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type HostGroupHostsList = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    address: string;
    address_list_range: string;
    editable: boolean;
};
export type HostGroupTeamRead = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    name: string;
    index: number;
};
export type HostList = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    address: string;
    address_list_range: string;
    editable: boolean;
};
export type HostRead = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    address: string;
    address_list_range: string;
    editable: boolean;
};
export type HostUpdate = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    address: string;
    address_list_range: string;
    editable: boolean;
};
export type HostCompetitionRead = {
    id: number;
    create_time: string;
    update_time: string;
    hidden: boolean;
    pause: boolean;
    name: string;
    display_name: string;
    round_duration: number;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type HostHostGroupRead = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    name: string;
};
export type HostServicesList = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    name: string;
    display_name: string;
    weight: number;
    point_boost: number;
    round_units: number;
    round_delay: number;
};
export type HostTeamRead = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    name: string;
    index: number;
};
export type Property = {
    id: number;
    create_time: string;
    update_time: string;
    competition_id: number;
    team_id: number;
    key: string;
    value: string;
    status: "view" | "edit" | "hide";
    competition: Competition;
    team: Team;
    services: Service;
};
export type PropertyCreate = {
    id: number;
    create_time: string;
    update_time: string;
    competition_id: number;
    team_id: number;
    key: string;
    value: string;
    status: "view" | "edit" | "hide";
};
export type PropertyList = {
    id: number;
    create_time: string;
    update_time: string;
    competition_id: number;
    team_id: number;
    key: string;
    value: string;
    status: "view" | "edit" | "hide";
};
export type PropertyRead = {
    id: number;
    create_time: string;
    update_time: string;
    competition_id: number;
    team_id: number;
    key: string;
    value: string;
    status: "view" | "edit" | "hide";
};
export type PropertyUpdate = {
    id: number;
    create_time: string;
    update_time: string;
    competition_id: number;
    team_id: number;
    key: string;
    value: string;
    status: "view" | "edit" | "hide";
};
export type PropertyCompetitionRead = {
    id: number;
    create_time: string;
    update_time: string;
    hidden: boolean;
    pause: boolean;
    name: string;
    display_name: string;
    round_duration: number;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type PropertyServicesRead = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    name: string;
    display_name: string;
    weight: number;
    point_boost: number;
    round_units: number;
    round_delay: number;
};
export type PropertyTeamRead = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    name: string;
    index: number;
};
export type Round = {
    id: number;
    create_time: string;
    update_time: string;
    competition_id: number;
    round_number: number;
    note: string;
    err: string;
    started_at: string;
    finished_at: string;
    competition: Competition;
    checks?: Check[];
};
export type RoundCreate = {
    id: number;
    create_time: string;
    update_time: string;
    competition_id: number;
    round_number: number;
    note: string;
    err: string;
    started_at: string;
    finished_at: string;
};
export type RoundList = {
    id: number;
    create_time: string;
    update_time: string;
    competition_id: number;
    round_number: number;
    note: string;
    err: string;
    started_at: string;
    finished_at: string;
};
export type RoundRead = {
    id: number;
    create_time: string;
    update_time: string;
    competition_id: number;
    round_number: number;
    note: string;
    err: string;
    started_at: string;
    finished_at: string;
};
export type RoundUpdate = {
    id: number;
    create_time: string;
    update_time: string;
    competition_id: number;
    round_number: number;
    note: string;
    err: string;
    started_at: string;
    finished_at: string;
};
export type RoundChecksList = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    log: string;
    error: string;
    passed: boolean;
};
export type RoundCompetitionRead = {
    id: number;
    create_time: string;
    update_time: string;
    hidden: boolean;
    pause: boolean;
    name: string;
    display_name: string;
    round_duration: number;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type Service = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    name: string;
    display_name: string;
    weight: number;
    point_boost: number;
    round_units: number;
    round_delay: number;
    competition: Competition;
    team: Team;
    hosts?: Host;
    checks?: Check[];
    properties?: Property[];
};
export type ServiceCreate = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    name: string;
    display_name: string;
    weight: number;
    point_boost: number;
    round_units: number;
    round_delay: number;
};
export type ServiceList = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    name: string;
    display_name: string;
    weight: number;
    point_boost: number;
    round_units: number;
    round_delay: number;
};
export type ServiceRead = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    name: string;
    display_name: string;
    weight: number;
    point_boost: number;
    round_units: number;
    round_delay: number;
};
export type ServiceUpdate = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    name: string;
    display_name: string;
    weight: number;
    point_boost: number;
    round_units: number;
    round_delay: number;
};
export type ServiceChecksList = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    log: string;
    error: string;
    passed: boolean;
};
export type ServiceCompetitionRead = {
    id: number;
    create_time: string;
    update_time: string;
    hidden: boolean;
    pause: boolean;
    name: string;
    display_name: string;
    round_duration: number;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type ServiceHostsRead = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    address: string;
    address_list_range: string;
    editable: boolean;
};
export type ServicePropertiesList = {
    id: number;
    create_time: string;
    update_time: string;
    competition_id: number;
    team_id: number;
    key: string;
    value: string;
    status: "view" | "edit" | "hide";
};
export type ServiceTeamRead = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    name: string;
    index: number;
};
export type Team = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    name: string;
    index: number;
    competition: Competition;
    users?: User[];
    hosts?: Host[];
};
export type TeamCreate = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    name: string;
    index: number;
};
export type TeamList = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    name: string;
    index: number;
};
export type TeamRead = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    name: string;
    index: number;
};
export type TeamUpdate = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    name: string;
    index: number;
};
export type TeamCompetitionRead = {
    id: number;
    create_time: string;
    update_time: string;
    hidden: boolean;
    pause: boolean;
    name: string;
    display_name: string;
    round_duration: number;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type TeamHostsList = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    team_id: number;
    address: string;
    address_list_range: string;
    editable: boolean;
};
export type TeamUsersList = {
    id: number;
    create_time: string;
    update_time: string;
    username: string;
};
export type User = {
    id: number;
    create_time: string;
    update_time: string;
    username: string;
    teams?: Team[];
    competitions?: Competition[];
};
export type UserCreate = {
    id: number;
    create_time: string;
    update_time: string;
    username: string;
};
export type UserList = {
    id: number;
    create_time: string;
    update_time: string;
    username: string;
};
export type UserRead = {
    id: number;
    create_time: string;
    update_time: string;
    username: string;
};
export type UserUpdate = {
    id: number;
    create_time: string;
    update_time: string;
    username: string;
};
export type UserCompetitionsList = {
    id: number;
    create_time: string;
    update_time: string;
    hidden: boolean;
    pause: boolean;
    name: string;
    display_name: string;
    round_duration: number;
    to_be_started_at?: string;
    started_at?: string;
    finished_at?: string;
};
export type UserTeamsList = {
    id: number;
    create_time: string;
    update_time: string;
    pause: boolean;
    hidden: boolean;
    competition_id: number;
    name: string;
    index: number;
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
    readCheck: (id: number) => ["readCheck", id] as const,
    readCheckCompetition: (id: number) => ["readCheckCompetition", id] as const,
    readCheckRounds: (id: number) => ["readCheckRounds", id] as const,
    readCheckServices: (id: number) => ["readCheckServices", id] as const,
    listCompetition: (page?: number, itemsPerPage?: number) => ["listCompetition", nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readCompetition: (id: number) => ["readCompetition", id] as const,
    listCompetitionTeams: (id: number, page?: number, itemsPerPage?: number) => ["listCompetitionTeams", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    listCompetitionUsers: (id: number, page?: number, itemsPerPage?: number) => ["listCompetitionUsers", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    listHostGroup: (page?: number, itemsPerPage?: number) => ["listHostGroup", nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readHostGroup: (id: number) => ["readHostGroup", id] as const,
    readHostGroupCompetition: (id: number) => ["readHostGroupCompetition", id] as const,
    listHostGroupHosts: (id: number, page?: number, itemsPerPage?: number) => ["listHostGroupHosts", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readHostGroupTeam: (id: number) => ["readHostGroupTeam", id] as const,
    listHost: (page?: number, itemsPerPage?: number) => ["listHost", nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readHost: (id: number) => ["readHost", id] as const,
    readHostCompetition: (id: number) => ["readHostCompetition", id] as const,
    readHostHostGroup: (id: number) => ["readHostHostGroup", id] as const,
    listHostServices: (id: number, page?: number, itemsPerPage?: number) => ["listHostServices", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readHostTeam: (id: number) => ["readHostTeam", id] as const,
    listProperty: (page?: number, itemsPerPage?: number) => ["listProperty", nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readProperty: (id: number) => ["readProperty", id] as const,
    readPropertyCompetition: (id: number) => ["readPropertyCompetition", id] as const,
    readPropertyServices: (id: number) => ["readPropertyServices", id] as const,
    readPropertyTeam: (id: number) => ["readPropertyTeam", id] as const,
    listRound: (page?: number, itemsPerPage?: number) => ["listRound", nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readRound: (id: number) => ["readRound", id] as const,
    listRoundChecks: (id: number, page?: number, itemsPerPage?: number) => ["listRoundChecks", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readRoundCompetition: (id: number) => ["readRoundCompetition", id] as const,
    listService: (page?: number, itemsPerPage?: number) => ["listService", nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readService: (id: number) => ["readService", id] as const,
    listServiceChecks: (id: number, page?: number, itemsPerPage?: number) => ["listServiceChecks", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readServiceCompetition: (id: number) => ["readServiceCompetition", id] as const,
    readServiceHosts: (id: number) => ["readServiceHosts", id] as const,
    listServiceProperties: (id: number, page?: number, itemsPerPage?: number) => ["listServiceProperties", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readServiceTeam: (id: number) => ["readServiceTeam", id] as const,
    listTeam: (page?: number, itemsPerPage?: number) => ["listTeam", nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readTeam: (id: number) => ["readTeam", id] as const,
    readTeamCompetition: (id: number) => ["readTeamCompetition", id] as const,
    listTeamHosts: (id: number, page?: number, itemsPerPage?: number) => ["listTeamHosts", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    listTeamUsers: (id: number, page?: number, itemsPerPage?: number) => ["listTeamUsers", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    listUser: (page?: number, itemsPerPage?: number) => ["listUser", nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    readUser: (id: number) => ["readUser", id] as const,
    listUserCompetitions: (id: number, page?: number, itemsPerPage?: number) => ["listUserCompetitions", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const,
    listUserTeams: (id: number, page?: number, itemsPerPage?: number) => ["listUserTeams", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)] as const
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
            create_time: string;
            update_time: string;
            pause: boolean;
            hidden: boolean;
            competition_id: number;
            log: string;
            error: string;
            passed: boolean;
            competition: number;
            rounds: number;
            services: number;
        }) => axios.request<CheckCreate>({
            method: "post",
            url: `/checks`,
            data: payload
        }).then(res => res.data),
        readCheck: (id: number) => axios.request<CheckRead>({
            method: "get",
            url: `/checks/${id}`
        }).then(res => res.data),
        deleteCheck: (id: number) => axios.request<unknown>({
            method: "delete",
            url: `/checks/${id}`
        }).then(res => res.data),
        updateCheck: (payload: {
            update_time?: string;
            pause?: boolean;
            hidden?: boolean;
            log?: string;
            error?: string;
            passed?: boolean;
            competition?: number;
            rounds?: number;
            services?: number;
        }, id: number) => axios.request<CheckUpdate>({
            method: "patch",
            url: `/checks/${id}`,
            data: payload
        }).then(res => res.data),
        readCheckCompetition: (id: number) => axios.request<CheckCompetitionRead>({
            method: "get",
            url: `/checks/${id}/competition`
        }).then(res => res.data),
        readCheckRounds: (id: number) => axios.request<CheckRoundsRead>({
            method: "get",
            url: `/checks/${id}/rounds`
        }).then(res => res.data),
        readCheckServices: (id: number) => axios.request<CheckServicesRead>({
            method: "get",
            url: `/checks/${id}/services`
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
            create_time: string;
            update_time: string;
            hidden: boolean;
            pause: boolean;
            name: string;
            display_name: string;
            round_duration: number;
            to_be_started_at?: string;
            started_at?: string;
            finished_at?: string;
            teams?: number[];
            users?: number[];
        }) => axios.request<CompetitionCreate>({
            method: "post",
            url: `/competitions`,
            data: payload
        }).then(res => res.data),
        readCompetition: (id: number) => axios.request<CompetitionRead>({
            method: "get",
            url: `/competitions/${id}`
        }).then(res => res.data),
        deleteCompetition: (id: number) => axios.request<unknown>({
            method: "delete",
            url: `/competitions/${id}`
        }).then(res => res.data),
        updateCompetition: (payload: {
            update_time?: string;
            hidden?: boolean;
            pause?: boolean;
            name?: string;
            display_name?: string;
            round_duration?: number;
            to_be_started_at?: string;
            started_at?: string;
            finished_at?: string;
            teams?: number[];
            users?: number[];
        }, id: number) => axios.request<CompetitionUpdate>({
            method: "patch",
            url: `/competitions/${id}`,
            data: payload
        }).then(res => res.data),
        listCompetitionTeams: (id: number, page?: number, itemsPerPage?: number) => axios.request<CompetitionTeamsList[]>({
            method: "get",
            url: `/competitions/${id}/teams`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        listCompetitionUsers: (id: number, page?: number, itemsPerPage?: number) => axios.request<CompetitionUsersList[]>({
            method: "get",
            url: `/competitions/${id}/users`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        listHostGroup: (page?: number, itemsPerPage?: number) => axios.request<HostGroupList[]>({
            method: "get",
            url: `/host-groups`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        createHostGroup: (payload: {
            create_time: string;
            update_time: string;
            pause: boolean;
            hidden: boolean;
            competition_id: number;
            team_id: number;
            name: string;
            competition: number;
            team: number;
            hosts?: number[];
        }) => axios.request<HostGroupCreate>({
            method: "post",
            url: `/host-groups`,
            data: payload
        }).then(res => res.data),
        readHostGroup: (id: number) => axios.request<HostGroupRead>({
            method: "get",
            url: `/host-groups/${id}`
        }).then(res => res.data),
        deleteHostGroup: (id: number) => axios.request<unknown>({
            method: "delete",
            url: `/host-groups/${id}`
        }).then(res => res.data),
        updateHostGroup: (payload: {
            update_time?: string;
            pause?: boolean;
            hidden?: boolean;
            name?: string;
            competition?: number;
            team?: number;
            hosts?: number[];
        }, id: number) => axios.request<HostGroupUpdate>({
            method: "patch",
            url: `/host-groups/${id}`,
            data: payload
        }).then(res => res.data),
        readHostGroupCompetition: (id: number) => axios.request<HostGroupCompetitionRead>({
            method: "get",
            url: `/host-groups/${id}/competition`
        }).then(res => res.data),
        listHostGroupHosts: (id: number, page?: number, itemsPerPage?: number) => axios.request<HostGroupHostsList[]>({
            method: "get",
            url: `/host-groups/${id}/hosts`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        readHostGroupTeam: (id: number) => axios.request<HostGroupTeamRead>({
            method: "get",
            url: `/host-groups/${id}/team`
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
            create_time: string;
            update_time: string;
            pause: boolean;
            hidden: boolean;
            competition_id: number;
            team_id: number;
            address: string;
            address_list_range: string;
            editable: boolean;
            competition: number;
            team: number;
            services?: number[];
            host_group: number;
        }) => axios.request<HostCreate>({
            method: "post",
            url: `/hosts`,
            data: payload
        }).then(res => res.data),
        readHost: (id: number) => axios.request<HostRead>({
            method: "get",
            url: `/hosts/${id}`
        }).then(res => res.data),
        deleteHost: (id: number) => axios.request<unknown>({
            method: "delete",
            url: `/hosts/${id}`
        }).then(res => res.data),
        updateHost: (payload: {
            update_time?: string;
            pause?: boolean;
            hidden?: boolean;
            address?: string;
            address_list_range?: string;
            editable?: boolean;
            competition?: number;
            team?: number;
            services?: number[];
            host_group?: number;
        }, id: number) => axios.request<HostUpdate>({
            method: "patch",
            url: `/hosts/${id}`,
            data: payload
        }).then(res => res.data),
        readHostCompetition: (id: number) => axios.request<HostCompetitionRead>({
            method: "get",
            url: `/hosts/${id}/competition`
        }).then(res => res.data),
        readHostHostGroup: (id: number) => axios.request<HostHostGroupRead>({
            method: "get",
            url: `/hosts/${id}/host-group`
        }).then(res => res.data),
        listHostServices: (id: number, page?: number, itemsPerPage?: number) => axios.request<HostServicesList[]>({
            method: "get",
            url: `/hosts/${id}/services`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        readHostTeam: (id: number) => axios.request<HostTeamRead>({
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
            create_time: string;
            update_time: string;
            competition_id: number;
            team_id: number;
            key: string;
            value: string;
            status: "view" | "edit" | "hide";
            competition: number;
            team: number;
            services: number;
        }) => axios.request<PropertyCreate>({
            method: "post",
            url: `/properties`,
            data: payload
        }).then(res => res.data),
        readProperty: (id: number) => axios.request<PropertyRead>({
            method: "get",
            url: `/properties/${id}`
        }).then(res => res.data),
        deleteProperty: (id: number) => axios.request<unknown>({
            method: "delete",
            url: `/properties/${id}`
        }).then(res => res.data),
        updateProperty: (payload: {
            update_time?: string;
            key?: string;
            value?: string;
            status?: "view" | "edit" | "hide";
            competition?: number;
            team?: number;
            services?: number;
        }, id: number) => axios.request<PropertyUpdate>({
            method: "patch",
            url: `/properties/${id}`,
            data: payload
        }).then(res => res.data),
        readPropertyCompetition: (id: number) => axios.request<PropertyCompetitionRead>({
            method: "get",
            url: `/properties/${id}/competition`
        }).then(res => res.data),
        readPropertyServices: (id: number) => axios.request<PropertyServicesRead>({
            method: "get",
            url: `/properties/${id}/services`
        }).then(res => res.data),
        readPropertyTeam: (id: number) => axios.request<PropertyTeamRead>({
            method: "get",
            url: `/properties/${id}/team`
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
            create_time: string;
            update_time: string;
            competition_id: number;
            round_number: number;
            note: string;
            err: string;
            started_at: string;
            finished_at: string;
            competition: number;
            checks?: number[];
        }) => axios.request<RoundCreate>({
            method: "post",
            url: `/rounds`,
            data: payload
        }).then(res => res.data),
        readRound: (id: number) => axios.request<RoundRead>({
            method: "get",
            url: `/rounds/${id}`
        }).then(res => res.data),
        deleteRound: (id: number) => axios.request<unknown>({
            method: "delete",
            url: `/rounds/${id}`
        }).then(res => res.data),
        updateRound: (payload: {
            update_time?: string;
            round_number?: number;
            note?: string;
            err?: string;
            started_at?: string;
            finished_at?: string;
            competition?: number;
            checks?: number[];
        }, id: number) => axios.request<RoundUpdate>({
            method: "patch",
            url: `/rounds/${id}`,
            data: payload
        }).then(res => res.data),
        listRoundChecks: (id: number, page?: number, itemsPerPage?: number) => axios.request<RoundChecksList[]>({
            method: "get",
            url: `/rounds/${id}/checks`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        readRoundCompetition: (id: number) => axios.request<RoundCompetitionRead>({
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
            create_time: string;
            update_time: string;
            pause: boolean;
            hidden: boolean;
            competition_id: number;
            team_id: number;
            name: string;
            display_name: string;
            weight: number;
            point_boost: number;
            round_units: number;
            round_delay: number;
            competition: number;
            team: number;
            hosts?: number;
            checks?: number[];
            properties?: number[];
        }) => axios.request<ServiceCreate>({
            method: "post",
            url: `/services`,
            data: payload
        }).then(res => res.data),
        readService: (id: number) => axios.request<ServiceRead>({
            method: "get",
            url: `/services/${id}`
        }).then(res => res.data),
        deleteService: (id: number) => axios.request<unknown>({
            method: "delete",
            url: `/services/${id}`
        }).then(res => res.data),
        updateService: (payload: {
            update_time?: string;
            pause?: boolean;
            hidden?: boolean;
            name?: string;
            display_name?: string;
            weight?: number;
            point_boost?: number;
            round_units?: number;
            round_delay?: number;
            competition?: number;
            team?: number;
            hosts?: number;
            checks?: number[];
            properties?: number[];
        }, id: number) => axios.request<ServiceUpdate>({
            method: "patch",
            url: `/services/${id}`,
            data: payload
        }).then(res => res.data),
        listServiceChecks: (id: number, page?: number, itemsPerPage?: number) => axios.request<ServiceChecksList[]>({
            method: "get",
            url: `/services/${id}/checks`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        readServiceCompetition: (id: number) => axios.request<ServiceCompetitionRead>({
            method: "get",
            url: `/services/${id}/competition`
        }).then(res => res.data),
        readServiceHosts: (id: number) => axios.request<ServiceHostsRead>({
            method: "get",
            url: `/services/${id}/hosts`
        }).then(res => res.data),
        listServiceProperties: (id: number, page?: number, itemsPerPage?: number) => axios.request<ServicePropertiesList[]>({
            method: "get",
            url: `/services/${id}/properties`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        readServiceTeam: (id: number) => axios.request<ServiceTeamRead>({
            method: "get",
            url: `/services/${id}/team`
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
            create_time: string;
            update_time: string;
            pause: boolean;
            hidden: boolean;
            competition_id: number;
            name: string;
            index: number;
            competition: number;
            users?: number[];
            hosts?: number[];
        }) => axios.request<TeamCreate>({
            method: "post",
            url: `/teams`,
            data: payload
        }).then(res => res.data),
        readTeam: (id: number) => axios.request<TeamRead>({
            method: "get",
            url: `/teams/${id}`
        }).then(res => res.data),
        deleteTeam: (id: number) => axios.request<unknown>({
            method: "delete",
            url: `/teams/${id}`
        }).then(res => res.data),
        updateTeam: (payload: {
            update_time?: string;
            pause?: boolean;
            hidden?: boolean;
            name?: string;
            index?: number;
            competition?: number;
            users?: number[];
            hosts?: number[];
        }, id: number) => axios.request<TeamUpdate>({
            method: "patch",
            url: `/teams/${id}`,
            data: payload
        }).then(res => res.data),
        readTeamCompetition: (id: number) => axios.request<TeamCompetitionRead>({
            method: "get",
            url: `/teams/${id}/competition`
        }).then(res => res.data),
        listTeamHosts: (id: number, page?: number, itemsPerPage?: number) => axios.request<TeamHostsList[]>({
            method: "get",
            url: `/teams/${id}/hosts`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        listTeamUsers: (id: number, page?: number, itemsPerPage?: number) => axios.request<TeamUsersList[]>({
            method: "get",
            url: `/teams/${id}/users`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        listUser: (page?: number, itemsPerPage?: number) => axios.request<UserList[]>({
            method: "get",
            url: `/users`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        createUser: (payload: {
            create_time: string;
            update_time: string;
            username: string;
            teams?: number[];
            competitions?: number[];
        }) => axios.request<UserCreate>({
            method: "post",
            url: `/users`,
            data: payload
        }).then(res => res.data),
        readUser: (id: number) => axios.request<UserRead>({
            method: "get",
            url: `/users/${id}`
        }).then(res => res.data),
        deleteUser: (id: number) => axios.request<unknown>({
            method: "delete",
            url: `/users/${id}`
        }).then(res => res.data),
        updateUser: (payload: {
            update_time?: string;
            teams?: number[];
            competitions?: number[];
        }, id: number) => axios.request<UserUpdate>({
            method: "patch",
            url: `/users/${id}`,
            data: payload
        }).then(res => res.data),
        listUserCompetitions: (id: number, page?: number, itemsPerPage?: number) => axios.request<UserCompetitionsList[]>({
            method: "get",
            url: `/users/${id}/competitions`,
            params: {
                ...(page !== undefined ? { page } : undefined),
                ...(itemsPerPage !== undefined ? { itemsPerPage } : undefined)
            },
            paramsSerializer: config?.paramsSerializer
        }).then(res => res.data),
        listUserTeams: (id: number, page?: number, itemsPerPage?: number) => axios.request<UserTeamsList[]>({
            method: "get",
            url: `/users/${id}/teams`,
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
        useReadCheck: (id: number, options?: Omit<UseQueryOptions<Response<"readCheck">, unknown, Response<"readCheck">, ReturnType<QueryKeys["readCheck"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readCheck">, unknown> => useQuery({ queryKey: queryKeys.readCheck(id), queryFn: () => requests.readCheck(id), ...options }),
        useReadCheckCompetition: (id: number, options?: Omit<UseQueryOptions<Response<"readCheckCompetition">, unknown, Response<"readCheckCompetition">, ReturnType<QueryKeys["readCheckCompetition"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readCheckCompetition">, unknown> => useQuery({ queryKey: queryKeys.readCheckCompetition(id), queryFn: () => requests.readCheckCompetition(id), ...options }),
        useReadCheckRounds: (id: number, options?: Omit<UseQueryOptions<Response<"readCheckRounds">, unknown, Response<"readCheckRounds">, ReturnType<QueryKeys["readCheckRounds"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readCheckRounds">, unknown> => useQuery({ queryKey: queryKeys.readCheckRounds(id), queryFn: () => requests.readCheckRounds(id), ...options }),
        useReadCheckServices: (id: number, options?: Omit<UseQueryOptions<Response<"readCheckServices">, unknown, Response<"readCheckServices">, ReturnType<QueryKeys["readCheckServices"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readCheckServices">, unknown> => useQuery({ queryKey: queryKeys.readCheckServices(id), queryFn: () => requests.readCheckServices(id), ...options }),
        useListCompetition: (page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listCompetition">, unknown, Response<"listCompetition">, ReturnType<QueryKeys["listCompetition"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listCompetition">, unknown> => useQuery({ queryKey: queryKeys.listCompetition(page, itemsPerPage), queryFn: () => requests.listCompetition(page, itemsPerPage), ...options }),
        useReadCompetition: (id: number, options?: Omit<UseQueryOptions<Response<"readCompetition">, unknown, Response<"readCompetition">, ReturnType<QueryKeys["readCompetition"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readCompetition">, unknown> => useQuery({ queryKey: queryKeys.readCompetition(id), queryFn: () => requests.readCompetition(id), ...options }),
        useListCompetitionTeams: (id: number, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listCompetitionTeams">, unknown, Response<"listCompetitionTeams">, ReturnType<QueryKeys["listCompetitionTeams"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listCompetitionTeams">, unknown> => useQuery({ queryKey: queryKeys.listCompetitionTeams(id, page, itemsPerPage), queryFn: () => requests.listCompetitionTeams(id, page, itemsPerPage), ...options }),
        useListCompetitionUsers: (id: number, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listCompetitionUsers">, unknown, Response<"listCompetitionUsers">, ReturnType<QueryKeys["listCompetitionUsers"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listCompetitionUsers">, unknown> => useQuery({ queryKey: queryKeys.listCompetitionUsers(id, page, itemsPerPage), queryFn: () => requests.listCompetitionUsers(id, page, itemsPerPage), ...options }),
        useListHostGroup: (page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listHostGroup">, unknown, Response<"listHostGroup">, ReturnType<QueryKeys["listHostGroup"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listHostGroup">, unknown> => useQuery({ queryKey: queryKeys.listHostGroup(page, itemsPerPage), queryFn: () => requests.listHostGroup(page, itemsPerPage), ...options }),
        useReadHostGroup: (id: number, options?: Omit<UseQueryOptions<Response<"readHostGroup">, unknown, Response<"readHostGroup">, ReturnType<QueryKeys["readHostGroup"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readHostGroup">, unknown> => useQuery({ queryKey: queryKeys.readHostGroup(id), queryFn: () => requests.readHostGroup(id), ...options }),
        useReadHostGroupCompetition: (id: number, options?: Omit<UseQueryOptions<Response<"readHostGroupCompetition">, unknown, Response<"readHostGroupCompetition">, ReturnType<QueryKeys["readHostGroupCompetition"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readHostGroupCompetition">, unknown> => useQuery({ queryKey: queryKeys.readHostGroupCompetition(id), queryFn: () => requests.readHostGroupCompetition(id), ...options }),
        useListHostGroupHosts: (id: number, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listHostGroupHosts">, unknown, Response<"listHostGroupHosts">, ReturnType<QueryKeys["listHostGroupHosts"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listHostGroupHosts">, unknown> => useQuery({ queryKey: queryKeys.listHostGroupHosts(id, page, itemsPerPage), queryFn: () => requests.listHostGroupHosts(id, page, itemsPerPage), ...options }),
        useReadHostGroupTeam: (id: number, options?: Omit<UseQueryOptions<Response<"readHostGroupTeam">, unknown, Response<"readHostGroupTeam">, ReturnType<QueryKeys["readHostGroupTeam"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readHostGroupTeam">, unknown> => useQuery({ queryKey: queryKeys.readHostGroupTeam(id), queryFn: () => requests.readHostGroupTeam(id), ...options }),
        useListHost: (page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listHost">, unknown, Response<"listHost">, ReturnType<QueryKeys["listHost"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listHost">, unknown> => useQuery({ queryKey: queryKeys.listHost(page, itemsPerPage), queryFn: () => requests.listHost(page, itemsPerPage), ...options }),
        useReadHost: (id: number, options?: Omit<UseQueryOptions<Response<"readHost">, unknown, Response<"readHost">, ReturnType<QueryKeys["readHost"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readHost">, unknown> => useQuery({ queryKey: queryKeys.readHost(id), queryFn: () => requests.readHost(id), ...options }),
        useReadHostCompetition: (id: number, options?: Omit<UseQueryOptions<Response<"readHostCompetition">, unknown, Response<"readHostCompetition">, ReturnType<QueryKeys["readHostCompetition"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readHostCompetition">, unknown> => useQuery({ queryKey: queryKeys.readHostCompetition(id), queryFn: () => requests.readHostCompetition(id), ...options }),
        useReadHostHostGroup: (id: number, options?: Omit<UseQueryOptions<Response<"readHostHostGroup">, unknown, Response<"readHostHostGroup">, ReturnType<QueryKeys["readHostHostGroup"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readHostHostGroup">, unknown> => useQuery({ queryKey: queryKeys.readHostHostGroup(id), queryFn: () => requests.readHostHostGroup(id), ...options }),
        useListHostServices: (id: number, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listHostServices">, unknown, Response<"listHostServices">, ReturnType<QueryKeys["listHostServices"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listHostServices">, unknown> => useQuery({ queryKey: queryKeys.listHostServices(id, page, itemsPerPage), queryFn: () => requests.listHostServices(id, page, itemsPerPage), ...options }),
        useReadHostTeam: (id: number, options?: Omit<UseQueryOptions<Response<"readHostTeam">, unknown, Response<"readHostTeam">, ReturnType<QueryKeys["readHostTeam"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readHostTeam">, unknown> => useQuery({ queryKey: queryKeys.readHostTeam(id), queryFn: () => requests.readHostTeam(id), ...options }),
        useListProperty: (page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listProperty">, unknown, Response<"listProperty">, ReturnType<QueryKeys["listProperty"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listProperty">, unknown> => useQuery({ queryKey: queryKeys.listProperty(page, itemsPerPage), queryFn: () => requests.listProperty(page, itemsPerPage), ...options }),
        useReadProperty: (id: number, options?: Omit<UseQueryOptions<Response<"readProperty">, unknown, Response<"readProperty">, ReturnType<QueryKeys["readProperty"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readProperty">, unknown> => useQuery({ queryKey: queryKeys.readProperty(id), queryFn: () => requests.readProperty(id), ...options }),
        useReadPropertyCompetition: (id: number, options?: Omit<UseQueryOptions<Response<"readPropertyCompetition">, unknown, Response<"readPropertyCompetition">, ReturnType<QueryKeys["readPropertyCompetition"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readPropertyCompetition">, unknown> => useQuery({ queryKey: queryKeys.readPropertyCompetition(id), queryFn: () => requests.readPropertyCompetition(id), ...options }),
        useReadPropertyServices: (id: number, options?: Omit<UseQueryOptions<Response<"readPropertyServices">, unknown, Response<"readPropertyServices">, ReturnType<QueryKeys["readPropertyServices"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readPropertyServices">, unknown> => useQuery({ queryKey: queryKeys.readPropertyServices(id), queryFn: () => requests.readPropertyServices(id), ...options }),
        useReadPropertyTeam: (id: number, options?: Omit<UseQueryOptions<Response<"readPropertyTeam">, unknown, Response<"readPropertyTeam">, ReturnType<QueryKeys["readPropertyTeam"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readPropertyTeam">, unknown> => useQuery({ queryKey: queryKeys.readPropertyTeam(id), queryFn: () => requests.readPropertyTeam(id), ...options }),
        useListRound: (page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listRound">, unknown, Response<"listRound">, ReturnType<QueryKeys["listRound"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listRound">, unknown> => useQuery({ queryKey: queryKeys.listRound(page, itemsPerPage), queryFn: () => requests.listRound(page, itemsPerPage), ...options }),
        useReadRound: (id: number, options?: Omit<UseQueryOptions<Response<"readRound">, unknown, Response<"readRound">, ReturnType<QueryKeys["readRound"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readRound">, unknown> => useQuery({ queryKey: queryKeys.readRound(id), queryFn: () => requests.readRound(id), ...options }),
        useListRoundChecks: (id: number, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listRoundChecks">, unknown, Response<"listRoundChecks">, ReturnType<QueryKeys["listRoundChecks"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listRoundChecks">, unknown> => useQuery({ queryKey: queryKeys.listRoundChecks(id, page, itemsPerPage), queryFn: () => requests.listRoundChecks(id, page, itemsPerPage), ...options }),
        useReadRoundCompetition: (id: number, options?: Omit<UseQueryOptions<Response<"readRoundCompetition">, unknown, Response<"readRoundCompetition">, ReturnType<QueryKeys["readRoundCompetition"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readRoundCompetition">, unknown> => useQuery({ queryKey: queryKeys.readRoundCompetition(id), queryFn: () => requests.readRoundCompetition(id), ...options }),
        useListService: (page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listService">, unknown, Response<"listService">, ReturnType<QueryKeys["listService"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listService">, unknown> => useQuery({ queryKey: queryKeys.listService(page, itemsPerPage), queryFn: () => requests.listService(page, itemsPerPage), ...options }),
        useReadService: (id: number, options?: Omit<UseQueryOptions<Response<"readService">, unknown, Response<"readService">, ReturnType<QueryKeys["readService"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readService">, unknown> => useQuery({ queryKey: queryKeys.readService(id), queryFn: () => requests.readService(id), ...options }),
        useListServiceChecks: (id: number, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listServiceChecks">, unknown, Response<"listServiceChecks">, ReturnType<QueryKeys["listServiceChecks"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listServiceChecks">, unknown> => useQuery({ queryKey: queryKeys.listServiceChecks(id, page, itemsPerPage), queryFn: () => requests.listServiceChecks(id, page, itemsPerPage), ...options }),
        useReadServiceCompetition: (id: number, options?: Omit<UseQueryOptions<Response<"readServiceCompetition">, unknown, Response<"readServiceCompetition">, ReturnType<QueryKeys["readServiceCompetition"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readServiceCompetition">, unknown> => useQuery({ queryKey: queryKeys.readServiceCompetition(id), queryFn: () => requests.readServiceCompetition(id), ...options }),
        useReadServiceHosts: (id: number, options?: Omit<UseQueryOptions<Response<"readServiceHosts">, unknown, Response<"readServiceHosts">, ReturnType<QueryKeys["readServiceHosts"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readServiceHosts">, unknown> => useQuery({ queryKey: queryKeys.readServiceHosts(id), queryFn: () => requests.readServiceHosts(id), ...options }),
        useListServiceProperties: (id: number, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listServiceProperties">, unknown, Response<"listServiceProperties">, ReturnType<QueryKeys["listServiceProperties"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listServiceProperties">, unknown> => useQuery({ queryKey: queryKeys.listServiceProperties(id, page, itemsPerPage), queryFn: () => requests.listServiceProperties(id, page, itemsPerPage), ...options }),
        useReadServiceTeam: (id: number, options?: Omit<UseQueryOptions<Response<"readServiceTeam">, unknown, Response<"readServiceTeam">, ReturnType<QueryKeys["readServiceTeam"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readServiceTeam">, unknown> => useQuery({ queryKey: queryKeys.readServiceTeam(id), queryFn: () => requests.readServiceTeam(id), ...options }),
        useListTeam: (page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listTeam">, unknown, Response<"listTeam">, ReturnType<QueryKeys["listTeam"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listTeam">, unknown> => useQuery({ queryKey: queryKeys.listTeam(page, itemsPerPage), queryFn: () => requests.listTeam(page, itemsPerPage), ...options }),
        useReadTeam: (id: number, options?: Omit<UseQueryOptions<Response<"readTeam">, unknown, Response<"readTeam">, ReturnType<QueryKeys["readTeam"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readTeam">, unknown> => useQuery({ queryKey: queryKeys.readTeam(id), queryFn: () => requests.readTeam(id), ...options }),
        useReadTeamCompetition: (id: number, options?: Omit<UseQueryOptions<Response<"readTeamCompetition">, unknown, Response<"readTeamCompetition">, ReturnType<QueryKeys["readTeamCompetition"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readTeamCompetition">, unknown> => useQuery({ queryKey: queryKeys.readTeamCompetition(id), queryFn: () => requests.readTeamCompetition(id), ...options }),
        useListTeamHosts: (id: number, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listTeamHosts">, unknown, Response<"listTeamHosts">, ReturnType<QueryKeys["listTeamHosts"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listTeamHosts">, unknown> => useQuery({ queryKey: queryKeys.listTeamHosts(id, page, itemsPerPage), queryFn: () => requests.listTeamHosts(id, page, itemsPerPage), ...options }),
        useListTeamUsers: (id: number, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listTeamUsers">, unknown, Response<"listTeamUsers">, ReturnType<QueryKeys["listTeamUsers"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listTeamUsers">, unknown> => useQuery({ queryKey: queryKeys.listTeamUsers(id, page, itemsPerPage), queryFn: () => requests.listTeamUsers(id, page, itemsPerPage), ...options }),
        useListUser: (page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listUser">, unknown, Response<"listUser">, ReturnType<QueryKeys["listUser"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listUser">, unknown> => useQuery({ queryKey: queryKeys.listUser(page, itemsPerPage), queryFn: () => requests.listUser(page, itemsPerPage), ...options }),
        useReadUser: (id: number, options?: Omit<UseQueryOptions<Response<"readUser">, unknown, Response<"readUser">, ReturnType<QueryKeys["readUser"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"readUser">, unknown> => useQuery({ queryKey: queryKeys.readUser(id), queryFn: () => requests.readUser(id), ...options }),
        useListUserCompetitions: (id: number, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listUserCompetitions">, unknown, Response<"listUserCompetitions">, ReturnType<QueryKeys["listUserCompetitions"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listUserCompetitions">, unknown> => useQuery({ queryKey: queryKeys.listUserCompetitions(id, page, itemsPerPage), queryFn: () => requests.listUserCompetitions(id, page, itemsPerPage), ...options }),
        useListUserTeams: (id: number, page?: number, itemsPerPage?: number, options?: Omit<UseQueryOptions<Response<"listUserTeams">, unknown, Response<"listUserTeams">, ReturnType<QueryKeys["listUserTeams"]>>, "queryKey" | "queryFn">): UseQueryResult<Response<"listUserTeams">, unknown> => useQuery({ queryKey: queryKeys.listUserTeams(id, page, itemsPerPage), queryFn: () => requests.listUserTeams(id, page, itemsPerPage), ...options })
    } as const;
}
type MutationConfigs = {
    useCreateCheck?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"createCheck">, unknown, Parameters<Requests["createCheck"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useUpdateCheck?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"updateCheck">, unknown, Parameters<Requests["updateCheck"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useDeleteCheck?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"deleteCheck">, unknown, unknown, unknown>, "onSuccess" | "onSettled" | "onError">;
    useCreateCompetition?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"createCompetition">, unknown, Parameters<Requests["createCompetition"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useUpdateCompetition?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"updateCompetition">, unknown, Parameters<Requests["updateCompetition"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useDeleteCompetition?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"deleteCompetition">, unknown, unknown, unknown>, "onSuccess" | "onSettled" | "onError">;
    useCreateHostGroup?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"createHostGroup">, unknown, Parameters<Requests["createHostGroup"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useUpdateHostGroup?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"updateHostGroup">, unknown, Parameters<Requests["updateHostGroup"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useDeleteHostGroup?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"deleteHostGroup">, unknown, unknown, unknown>, "onSuccess" | "onSettled" | "onError">;
    useCreateHost?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"createHost">, unknown, Parameters<Requests["createHost"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useUpdateHost?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"updateHost">, unknown, Parameters<Requests["updateHost"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useDeleteHost?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"deleteHost">, unknown, unknown, unknown>, "onSuccess" | "onSettled" | "onError">;
    useCreateProperty?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"createProperty">, unknown, Parameters<Requests["createProperty"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useUpdateProperty?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"updateProperty">, unknown, Parameters<Requests["updateProperty"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useDeleteProperty?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"deleteProperty">, unknown, unknown, unknown>, "onSuccess" | "onSettled" | "onError">;
    useCreateRound?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"createRound">, unknown, Parameters<Requests["createRound"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useUpdateRound?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"updateRound">, unknown, Parameters<Requests["updateRound"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useDeleteRound?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"deleteRound">, unknown, unknown, unknown>, "onSuccess" | "onSettled" | "onError">;
    useCreateService?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"createService">, unknown, Parameters<Requests["createService"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useUpdateService?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"updateService">, unknown, Parameters<Requests["updateService"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useDeleteService?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"deleteService">, unknown, unknown, unknown>, "onSuccess" | "onSettled" | "onError">;
    useCreateTeam?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"createTeam">, unknown, Parameters<Requests["createTeam"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useUpdateTeam?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"updateTeam">, unknown, Parameters<Requests["updateTeam"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useDeleteTeam?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"deleteTeam">, unknown, unknown, unknown>, "onSuccess" | "onSettled" | "onError">;
    useCreateUser?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"createUser">, unknown, Parameters<Requests["createUser"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useUpdateUser?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"updateUser">, unknown, Parameters<Requests["updateUser"]>[0], unknown>, "onSuccess" | "onSettled" | "onError">;
    useDeleteUser?: (queryClient: QueryClient) => Pick<UseMutationOptions<Response<"deleteUser">, unknown, unknown, unknown>, "onSuccess" | "onSettled" | "onError">;
};
function makeMutations(requests: Requests, config?: Config["mutations"]) {
    return {
        useCreateCheck: (options?: Omit<UseMutationOptions<Response<"createCheck">, unknown, Parameters<Requests["createCheck"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"createCheck">, unknown, Parameters<Requests["createCheck"]>[0]>(payload => requests.createCheck(payload), config?.useCreateCheck, options),
        useUpdateCheck: (id: number, options?: Omit<UseMutationOptions<Response<"updateCheck">, unknown, Parameters<Requests["updateCheck"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"updateCheck">, unknown, Parameters<Requests["updateCheck"]>[0]>(payload => requests.updateCheck(payload, id), config?.useUpdateCheck, options),
        useDeleteCheck: (id: number, options?: Omit<UseMutationOptions<Response<"deleteCheck">, unknown, unknown, unknown>, "mutationFn">) => useRapiniMutation<Response<"deleteCheck">, unknown, unknown>(() => requests.deleteCheck(id), config?.useDeleteCheck, options),
        useCreateCompetition: (options?: Omit<UseMutationOptions<Response<"createCompetition">, unknown, Parameters<Requests["createCompetition"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"createCompetition">, unknown, Parameters<Requests["createCompetition"]>[0]>(payload => requests.createCompetition(payload), config?.useCreateCompetition, options),
        useUpdateCompetition: (id: number, options?: Omit<UseMutationOptions<Response<"updateCompetition">, unknown, Parameters<Requests["updateCompetition"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"updateCompetition">, unknown, Parameters<Requests["updateCompetition"]>[0]>(payload => requests.updateCompetition(payload, id), config?.useUpdateCompetition, options),
        useDeleteCompetition: (id: number, options?: Omit<UseMutationOptions<Response<"deleteCompetition">, unknown, unknown, unknown>, "mutationFn">) => useRapiniMutation<Response<"deleteCompetition">, unknown, unknown>(() => requests.deleteCompetition(id), config?.useDeleteCompetition, options),
        useCreateHostGroup: (options?: Omit<UseMutationOptions<Response<"createHostGroup">, unknown, Parameters<Requests["createHostGroup"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"createHostGroup">, unknown, Parameters<Requests["createHostGroup"]>[0]>(payload => requests.createHostGroup(payload), config?.useCreateHostGroup, options),
        useUpdateHostGroup: (id: number, options?: Omit<UseMutationOptions<Response<"updateHostGroup">, unknown, Parameters<Requests["updateHostGroup"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"updateHostGroup">, unknown, Parameters<Requests["updateHostGroup"]>[0]>(payload => requests.updateHostGroup(payload, id), config?.useUpdateHostGroup, options),
        useDeleteHostGroup: (id: number, options?: Omit<UseMutationOptions<Response<"deleteHostGroup">, unknown, unknown, unknown>, "mutationFn">) => useRapiniMutation<Response<"deleteHostGroup">, unknown, unknown>(() => requests.deleteHostGroup(id), config?.useDeleteHostGroup, options),
        useCreateHost: (options?: Omit<UseMutationOptions<Response<"createHost">, unknown, Parameters<Requests["createHost"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"createHost">, unknown, Parameters<Requests["createHost"]>[0]>(payload => requests.createHost(payload), config?.useCreateHost, options),
        useUpdateHost: (id: number, options?: Omit<UseMutationOptions<Response<"updateHost">, unknown, Parameters<Requests["updateHost"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"updateHost">, unknown, Parameters<Requests["updateHost"]>[0]>(payload => requests.updateHost(payload, id), config?.useUpdateHost, options),
        useDeleteHost: (id: number, options?: Omit<UseMutationOptions<Response<"deleteHost">, unknown, unknown, unknown>, "mutationFn">) => useRapiniMutation<Response<"deleteHost">, unknown, unknown>(() => requests.deleteHost(id), config?.useDeleteHost, options),
        useCreateProperty: (options?: Omit<UseMutationOptions<Response<"createProperty">, unknown, Parameters<Requests["createProperty"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"createProperty">, unknown, Parameters<Requests["createProperty"]>[0]>(payload => requests.createProperty(payload), config?.useCreateProperty, options),
        useUpdateProperty: (id: number, options?: Omit<UseMutationOptions<Response<"updateProperty">, unknown, Parameters<Requests["updateProperty"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"updateProperty">, unknown, Parameters<Requests["updateProperty"]>[0]>(payload => requests.updateProperty(payload, id), config?.useUpdateProperty, options),
        useDeleteProperty: (id: number, options?: Omit<UseMutationOptions<Response<"deleteProperty">, unknown, unknown, unknown>, "mutationFn">) => useRapiniMutation<Response<"deleteProperty">, unknown, unknown>(() => requests.deleteProperty(id), config?.useDeleteProperty, options),
        useCreateRound: (options?: Omit<UseMutationOptions<Response<"createRound">, unknown, Parameters<Requests["createRound"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"createRound">, unknown, Parameters<Requests["createRound"]>[0]>(payload => requests.createRound(payload), config?.useCreateRound, options),
        useUpdateRound: (id: number, options?: Omit<UseMutationOptions<Response<"updateRound">, unknown, Parameters<Requests["updateRound"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"updateRound">, unknown, Parameters<Requests["updateRound"]>[0]>(payload => requests.updateRound(payload, id), config?.useUpdateRound, options),
        useDeleteRound: (id: number, options?: Omit<UseMutationOptions<Response<"deleteRound">, unknown, unknown, unknown>, "mutationFn">) => useRapiniMutation<Response<"deleteRound">, unknown, unknown>(() => requests.deleteRound(id), config?.useDeleteRound, options),
        useCreateService: (options?: Omit<UseMutationOptions<Response<"createService">, unknown, Parameters<Requests["createService"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"createService">, unknown, Parameters<Requests["createService"]>[0]>(payload => requests.createService(payload), config?.useCreateService, options),
        useUpdateService: (id: number, options?: Omit<UseMutationOptions<Response<"updateService">, unknown, Parameters<Requests["updateService"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"updateService">, unknown, Parameters<Requests["updateService"]>[0]>(payload => requests.updateService(payload, id), config?.useUpdateService, options),
        useDeleteService: (id: number, options?: Omit<UseMutationOptions<Response<"deleteService">, unknown, unknown, unknown>, "mutationFn">) => useRapiniMutation<Response<"deleteService">, unknown, unknown>(() => requests.deleteService(id), config?.useDeleteService, options),
        useCreateTeam: (options?: Omit<UseMutationOptions<Response<"createTeam">, unknown, Parameters<Requests["createTeam"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"createTeam">, unknown, Parameters<Requests["createTeam"]>[0]>(payload => requests.createTeam(payload), config?.useCreateTeam, options),
        useUpdateTeam: (id: number, options?: Omit<UseMutationOptions<Response<"updateTeam">, unknown, Parameters<Requests["updateTeam"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"updateTeam">, unknown, Parameters<Requests["updateTeam"]>[0]>(payload => requests.updateTeam(payload, id), config?.useUpdateTeam, options),
        useDeleteTeam: (id: number, options?: Omit<UseMutationOptions<Response<"deleteTeam">, unknown, unknown, unknown>, "mutationFn">) => useRapiniMutation<Response<"deleteTeam">, unknown, unknown>(() => requests.deleteTeam(id), config?.useDeleteTeam, options),
        useCreateUser: (options?: Omit<UseMutationOptions<Response<"createUser">, unknown, Parameters<Requests["createUser"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"createUser">, unknown, Parameters<Requests["createUser"]>[0]>(payload => requests.createUser(payload), config?.useCreateUser, options),
        useUpdateUser: (id: number, options?: Omit<UseMutationOptions<Response<"updateUser">, unknown, Parameters<Requests["updateUser"]>[0], unknown>, "mutationFn">) => useRapiniMutation<Response<"updateUser">, unknown, Parameters<Requests["updateUser"]>[0]>(payload => requests.updateUser(payload, id), config?.useUpdateUser, options),
        useDeleteUser: (id: number, options?: Omit<UseMutationOptions<Response<"deleteUser">, unknown, unknown, unknown>, "mutationFn">) => useRapiniMutation<Response<"deleteUser">, unknown, unknown>(() => requests.deleteUser(id), config?.useDeleteUser, options)
    } as const;
}
