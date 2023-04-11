var __rest = (this && this.__rest) || function (s, e) {
    var t = {};
    for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p) && e.indexOf(p) < 0)
        t[p] = s[p];
    if (s != null && typeof Object.getOwnPropertySymbols === "function")
        for (var i = 0, p = Object.getOwnPropertySymbols(s); i < p.length; i++) {
            if (e.indexOf(p[i]) < 0 && Object.prototype.propertyIsEnumerable.call(s, p[i]))
                t[p[i]] = s[p[i]];
        }
    return t;
};
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
export function initialize(axios, config) {
    const requests = makeRequests(axios, config === null || config === void 0 ? void 0 : config.axios);
    return {
        requests,
        queries: makeQueries(requests),
        mutations: makeMutations(requests, config === null || config === void 0 ? void 0 : config.mutations)
    };
}
function useRapiniMutation(mutationFn, config, options) {
    const _a = options !== null && options !== void 0 ? options : {}, { onSuccess, onError, onSettled } = _a, rest = __rest(_a, ["onSuccess", "onError", "onSettled"]);
    const queryClient = useQueryClient();
    const conf = config === null || config === void 0 ? void 0 : config(queryClient);
    const mutationOptions = Object.assign({ onSuccess: (data, variables, context) => {
            var _a;
            (_a = conf === null || conf === void 0 ? void 0 : conf.onSuccess) === null || _a === void 0 ? void 0 : _a.call(conf, data, variables, context);
            onSuccess === null || onSuccess === void 0 ? void 0 : onSuccess(data, variables, context);
        }, onError: (error, variables, context) => {
            var _a;
            (_a = conf === null || conf === void 0 ? void 0 : conf.onError) === null || _a === void 0 ? void 0 : _a.call(conf, error, variables, context);
            onError === null || onError === void 0 ? void 0 : onError(error, variables, context);
        }, onSettled: (data, error, variables, context) => {
            var _a;
            (_a = conf === null || conf === void 0 ? void 0 : conf.onSettled) === null || _a === void 0 ? void 0 : _a.call(conf, data, error, variables, context);
            onSettled === null || onSettled === void 0 ? void 0 : onSettled(data, error, variables, context);
        } }, rest);
    return useMutation(Object.assign({ mutationFn }, mutationOptions));
}
function nullIfUndefined(value) {
    return typeof value === "undefined" ? null : value;
}
export const queryKeys = {
    listCheck: (page, itemsPerPage) => ["listCheck", nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readCheck: (id) => ["readCheck", id],
    readCheckCompetition: (id) => ["readCheckCompetition", id],
    readCheckRounds: (id) => ["readCheckRounds", id],
    readCheckServices: (id) => ["readCheckServices", id],
    listCompetition: (page, itemsPerPage) => ["listCompetition", nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readCompetition: (id) => ["readCompetition", id],
    listCompetitionTeams: (id, page, itemsPerPage) => ["listCompetitionTeams", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    listCompetitionUsers: (id, page, itemsPerPage) => ["listCompetitionUsers", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    listHostGroup: (page, itemsPerPage) => ["listHostGroup", nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readHostGroup: (id) => ["readHostGroup", id],
    readHostGroupCompetition: (id) => ["readHostGroupCompetition", id],
    listHostGroupHosts: (id, page, itemsPerPage) => ["listHostGroupHosts", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readHostGroupTeam: (id) => ["readHostGroupTeam", id],
    listHost: (page, itemsPerPage) => ["listHost", nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readHost: (id) => ["readHost", id],
    readHostCompetition: (id) => ["readHostCompetition", id],
    readHostHostGroup: (id) => ["readHostHostGroup", id],
    listHostServices: (id, page, itemsPerPage) => ["listHostServices", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readHostTeam: (id) => ["readHostTeam", id],
    listProperty: (page, itemsPerPage) => ["listProperty", nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readProperty: (id) => ["readProperty", id],
    readPropertyCompetition: (id) => ["readPropertyCompetition", id],
    readPropertyServices: (id) => ["readPropertyServices", id],
    readPropertyTeam: (id) => ["readPropertyTeam", id],
    listRound: (page, itemsPerPage) => ["listRound", nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readRound: (id) => ["readRound", id],
    listRoundChecks: (id, page, itemsPerPage) => ["listRoundChecks", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readRoundCompetition: (id) => ["readRoundCompetition", id],
    listService: (page, itemsPerPage) => ["listService", nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readService: (id) => ["readService", id],
    listServiceChecks: (id, page, itemsPerPage) => ["listServiceChecks", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readServiceCompetition: (id) => ["readServiceCompetition", id],
    readServiceHosts: (id) => ["readServiceHosts", id],
    listServiceProperties: (id, page, itemsPerPage) => ["listServiceProperties", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readServiceTeam: (id) => ["readServiceTeam", id],
    listTeam: (page, itemsPerPage) => ["listTeam", nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readTeam: (id) => ["readTeam", id],
    readTeamCompetition: (id) => ["readTeamCompetition", id],
    listTeamHosts: (id, page, itemsPerPage) => ["listTeamHosts", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    listTeamUsers: (id, page, itemsPerPage) => ["listTeamUsers", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    listUser: (page, itemsPerPage) => ["listUser", nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readUser: (id) => ["readUser", id],
    listUserCompetitions: (id, page, itemsPerPage) => ["listUserCompetitions", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    listUserTeams: (id, page, itemsPerPage) => ["listUserTeams", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)]
};
function makeRequests(axios, config) {
    return {
        listCheck: (page, itemsPerPage) => axios.request({
            method: "get",
            url: `/checks`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        createCheck: (payload) => axios.request({
            method: "post",
            url: `/checks`,
            data: payload
        }).then(res => res.data),
        readCheck: (id) => axios.request({
            method: "get",
            url: `/checks/${id}`
        }).then(res => res.data),
        deleteCheck: (id) => axios.request({
            method: "delete",
            url: `/checks/${id}`
        }).then(res => res.data),
        updateCheck: (payload, id) => axios.request({
            method: "patch",
            url: `/checks/${id}`,
            data: payload
        }).then(res => res.data),
        readCheckCompetition: (id) => axios.request({
            method: "get",
            url: `/checks/${id}/competition`
        }).then(res => res.data),
        readCheckRounds: (id) => axios.request({
            method: "get",
            url: `/checks/${id}/rounds`
        }).then(res => res.data),
        readCheckServices: (id) => axios.request({
            method: "get",
            url: `/checks/${id}/services`
        }).then(res => res.data),
        listCompetition: (page, itemsPerPage) => axios.request({
            method: "get",
            url: `/competitions`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        createCompetition: (payload) => axios.request({
            method: "post",
            url: `/competitions`,
            data: payload
        }).then(res => res.data),
        readCompetition: (id) => axios.request({
            method: "get",
            url: `/competitions/${id}`
        }).then(res => res.data),
        deleteCompetition: (id) => axios.request({
            method: "delete",
            url: `/competitions/${id}`
        }).then(res => res.data),
        updateCompetition: (payload, id) => axios.request({
            method: "patch",
            url: `/competitions/${id}`,
            data: payload
        }).then(res => res.data),
        listCompetitionTeams: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/competitions/${id}/teams`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        listCompetitionUsers: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/competitions/${id}/users`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        listHostGroup: (page, itemsPerPage) => axios.request({
            method: "get",
            url: `/host-groups`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        createHostGroup: (payload) => axios.request({
            method: "post",
            url: `/host-groups`,
            data: payload
        }).then(res => res.data),
        readHostGroup: (id) => axios.request({
            method: "get",
            url: `/host-groups/${id}`
        }).then(res => res.data),
        deleteHostGroup: (id) => axios.request({
            method: "delete",
            url: `/host-groups/${id}`
        }).then(res => res.data),
        updateHostGroup: (payload, id) => axios.request({
            method: "patch",
            url: `/host-groups/${id}`,
            data: payload
        }).then(res => res.data),
        readHostGroupCompetition: (id) => axios.request({
            method: "get",
            url: `/host-groups/${id}/competition`
        }).then(res => res.data),
        listHostGroupHosts: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/host-groups/${id}/hosts`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        readHostGroupTeam: (id) => axios.request({
            method: "get",
            url: `/host-groups/${id}/team`
        }).then(res => res.data),
        listHost: (page, itemsPerPage) => axios.request({
            method: "get",
            url: `/hosts`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        createHost: (payload) => axios.request({
            method: "post",
            url: `/hosts`,
            data: payload
        }).then(res => res.data),
        readHost: (id) => axios.request({
            method: "get",
            url: `/hosts/${id}`
        }).then(res => res.data),
        deleteHost: (id) => axios.request({
            method: "delete",
            url: `/hosts/${id}`
        }).then(res => res.data),
        updateHost: (payload, id) => axios.request({
            method: "patch",
            url: `/hosts/${id}`,
            data: payload
        }).then(res => res.data),
        readHostCompetition: (id) => axios.request({
            method: "get",
            url: `/hosts/${id}/competition`
        }).then(res => res.data),
        readHostHostGroup: (id) => axios.request({
            method: "get",
            url: `/hosts/${id}/host-group`
        }).then(res => res.data),
        listHostServices: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/hosts/${id}/services`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        readHostTeam: (id) => axios.request({
            method: "get",
            url: `/hosts/${id}/team`
        }).then(res => res.data),
        listProperty: (page, itemsPerPage) => axios.request({
            method: "get",
            url: `/properties`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        createProperty: (payload) => axios.request({
            method: "post",
            url: `/properties`,
            data: payload
        }).then(res => res.data),
        readProperty: (id) => axios.request({
            method: "get",
            url: `/properties/${id}`
        }).then(res => res.data),
        deleteProperty: (id) => axios.request({
            method: "delete",
            url: `/properties/${id}`
        }).then(res => res.data),
        updateProperty: (payload, id) => axios.request({
            method: "patch",
            url: `/properties/${id}`,
            data: payload
        }).then(res => res.data),
        readPropertyCompetition: (id) => axios.request({
            method: "get",
            url: `/properties/${id}/competition`
        }).then(res => res.data),
        readPropertyServices: (id) => axios.request({
            method: "get",
            url: `/properties/${id}/services`
        }).then(res => res.data),
        readPropertyTeam: (id) => axios.request({
            method: "get",
            url: `/properties/${id}/team`
        }).then(res => res.data),
        listRound: (page, itemsPerPage) => axios.request({
            method: "get",
            url: `/rounds`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        createRound: (payload) => axios.request({
            method: "post",
            url: `/rounds`,
            data: payload
        }).then(res => res.data),
        readRound: (id) => axios.request({
            method: "get",
            url: `/rounds/${id}`
        }).then(res => res.data),
        deleteRound: (id) => axios.request({
            method: "delete",
            url: `/rounds/${id}`
        }).then(res => res.data),
        updateRound: (payload, id) => axios.request({
            method: "patch",
            url: `/rounds/${id}`,
            data: payload
        }).then(res => res.data),
        listRoundChecks: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/rounds/${id}/checks`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        readRoundCompetition: (id) => axios.request({
            method: "get",
            url: `/rounds/${id}/competition`
        }).then(res => res.data),
        listService: (page, itemsPerPage) => axios.request({
            method: "get",
            url: `/services`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        createService: (payload) => axios.request({
            method: "post",
            url: `/services`,
            data: payload
        }).then(res => res.data),
        readService: (id) => axios.request({
            method: "get",
            url: `/services/${id}`
        }).then(res => res.data),
        deleteService: (id) => axios.request({
            method: "delete",
            url: `/services/${id}`
        }).then(res => res.data),
        updateService: (payload, id) => axios.request({
            method: "patch",
            url: `/services/${id}`,
            data: payload
        }).then(res => res.data),
        listServiceChecks: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/services/${id}/checks`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        readServiceCompetition: (id) => axios.request({
            method: "get",
            url: `/services/${id}/competition`
        }).then(res => res.data),
        readServiceHosts: (id) => axios.request({
            method: "get",
            url: `/services/${id}/hosts`
        }).then(res => res.data),
        listServiceProperties: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/services/${id}/properties`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        readServiceTeam: (id) => axios.request({
            method: "get",
            url: `/services/${id}/team`
        }).then(res => res.data),
        listTeam: (page, itemsPerPage) => axios.request({
            method: "get",
            url: `/teams`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        createTeam: (payload) => axios.request({
            method: "post",
            url: `/teams`,
            data: payload
        }).then(res => res.data),
        readTeam: (id) => axios.request({
            method: "get",
            url: `/teams/${id}`
        }).then(res => res.data),
        deleteTeam: (id) => axios.request({
            method: "delete",
            url: `/teams/${id}`
        }).then(res => res.data),
        updateTeam: (payload, id) => axios.request({
            method: "patch",
            url: `/teams/${id}`,
            data: payload
        }).then(res => res.data),
        readTeamCompetition: (id) => axios.request({
            method: "get",
            url: `/teams/${id}/competition`
        }).then(res => res.data),
        listTeamHosts: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/teams/${id}/hosts`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        listTeamUsers: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/teams/${id}/users`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        listUser: (page, itemsPerPage) => axios.request({
            method: "get",
            url: `/users`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        createUser: (payload) => axios.request({
            method: "post",
            url: `/users`,
            data: payload
        }).then(res => res.data),
        readUser: (id) => axios.request({
            method: "get",
            url: `/users/${id}`
        }).then(res => res.data),
        deleteUser: (id) => axios.request({
            method: "delete",
            url: `/users/${id}`
        }).then(res => res.data),
        updateUser: (payload, id) => axios.request({
            method: "patch",
            url: `/users/${id}`,
            data: payload
        }).then(res => res.data),
        listUserCompetitions: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/users/${id}/competitions`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        listUserTeams: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/users/${id}/teams`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data)
    };
}
function makeQueries(requests) {
    return {
        useListCheck: (page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listCheck(page, itemsPerPage), queryFn: () => requests.listCheck(page, itemsPerPage) }, options)),
        useReadCheck: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readCheck(id), queryFn: () => requests.readCheck(id) }, options)),
        useReadCheckCompetition: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readCheckCompetition(id), queryFn: () => requests.readCheckCompetition(id) }, options)),
        useReadCheckRounds: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readCheckRounds(id), queryFn: () => requests.readCheckRounds(id) }, options)),
        useReadCheckServices: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readCheckServices(id), queryFn: () => requests.readCheckServices(id) }, options)),
        useListCompetition: (page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listCompetition(page, itemsPerPage), queryFn: () => requests.listCompetition(page, itemsPerPage) }, options)),
        useReadCompetition: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readCompetition(id), queryFn: () => requests.readCompetition(id) }, options)),
        useListCompetitionTeams: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listCompetitionTeams(id, page, itemsPerPage), queryFn: () => requests.listCompetitionTeams(id, page, itemsPerPage) }, options)),
        useListCompetitionUsers: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listCompetitionUsers(id, page, itemsPerPage), queryFn: () => requests.listCompetitionUsers(id, page, itemsPerPage) }, options)),
        useListHostGroup: (page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listHostGroup(page, itemsPerPage), queryFn: () => requests.listHostGroup(page, itemsPerPage) }, options)),
        useReadHostGroup: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readHostGroup(id), queryFn: () => requests.readHostGroup(id) }, options)),
        useReadHostGroupCompetition: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readHostGroupCompetition(id), queryFn: () => requests.readHostGroupCompetition(id) }, options)),
        useListHostGroupHosts: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listHostGroupHosts(id, page, itemsPerPage), queryFn: () => requests.listHostGroupHosts(id, page, itemsPerPage) }, options)),
        useReadHostGroupTeam: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readHostGroupTeam(id), queryFn: () => requests.readHostGroupTeam(id) }, options)),
        useListHost: (page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listHost(page, itemsPerPage), queryFn: () => requests.listHost(page, itemsPerPage) }, options)),
        useReadHost: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readHost(id), queryFn: () => requests.readHost(id) }, options)),
        useReadHostCompetition: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readHostCompetition(id), queryFn: () => requests.readHostCompetition(id) }, options)),
        useReadHostHostGroup: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readHostHostGroup(id), queryFn: () => requests.readHostHostGroup(id) }, options)),
        useListHostServices: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listHostServices(id, page, itemsPerPage), queryFn: () => requests.listHostServices(id, page, itemsPerPage) }, options)),
        useReadHostTeam: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readHostTeam(id), queryFn: () => requests.readHostTeam(id) }, options)),
        useListProperty: (page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listProperty(page, itemsPerPage), queryFn: () => requests.listProperty(page, itemsPerPage) }, options)),
        useReadProperty: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readProperty(id), queryFn: () => requests.readProperty(id) }, options)),
        useReadPropertyCompetition: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readPropertyCompetition(id), queryFn: () => requests.readPropertyCompetition(id) }, options)),
        useReadPropertyServices: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readPropertyServices(id), queryFn: () => requests.readPropertyServices(id) }, options)),
        useReadPropertyTeam: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readPropertyTeam(id), queryFn: () => requests.readPropertyTeam(id) }, options)),
        useListRound: (page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listRound(page, itemsPerPage), queryFn: () => requests.listRound(page, itemsPerPage) }, options)),
        useReadRound: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readRound(id), queryFn: () => requests.readRound(id) }, options)),
        useListRoundChecks: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listRoundChecks(id, page, itemsPerPage), queryFn: () => requests.listRoundChecks(id, page, itemsPerPage) }, options)),
        useReadRoundCompetition: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readRoundCompetition(id), queryFn: () => requests.readRoundCompetition(id) }, options)),
        useListService: (page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listService(page, itemsPerPage), queryFn: () => requests.listService(page, itemsPerPage) }, options)),
        useReadService: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readService(id), queryFn: () => requests.readService(id) }, options)),
        useListServiceChecks: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listServiceChecks(id, page, itemsPerPage), queryFn: () => requests.listServiceChecks(id, page, itemsPerPage) }, options)),
        useReadServiceCompetition: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readServiceCompetition(id), queryFn: () => requests.readServiceCompetition(id) }, options)),
        useReadServiceHosts: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readServiceHosts(id), queryFn: () => requests.readServiceHosts(id) }, options)),
        useListServiceProperties: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listServiceProperties(id, page, itemsPerPage), queryFn: () => requests.listServiceProperties(id, page, itemsPerPage) }, options)),
        useReadServiceTeam: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readServiceTeam(id), queryFn: () => requests.readServiceTeam(id) }, options)),
        useListTeam: (page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listTeam(page, itemsPerPage), queryFn: () => requests.listTeam(page, itemsPerPage) }, options)),
        useReadTeam: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readTeam(id), queryFn: () => requests.readTeam(id) }, options)),
        useReadTeamCompetition: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readTeamCompetition(id), queryFn: () => requests.readTeamCompetition(id) }, options)),
        useListTeamHosts: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listTeamHosts(id, page, itemsPerPage), queryFn: () => requests.listTeamHosts(id, page, itemsPerPage) }, options)),
        useListTeamUsers: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listTeamUsers(id, page, itemsPerPage), queryFn: () => requests.listTeamUsers(id, page, itemsPerPage) }, options)),
        useListUser: (page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listUser(page, itemsPerPage), queryFn: () => requests.listUser(page, itemsPerPage) }, options)),
        useReadUser: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readUser(id), queryFn: () => requests.readUser(id) }, options)),
        useListUserCompetitions: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listUserCompetitions(id, page, itemsPerPage), queryFn: () => requests.listUserCompetitions(id, page, itemsPerPage) }, options)),
        useListUserTeams: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listUserTeams(id, page, itemsPerPage), queryFn: () => requests.listUserTeams(id, page, itemsPerPage) }, options))
    };
}
function makeMutations(requests, config) {
    return {
        useCreateCheck: (options) => useRapiniMutation(payload => requests.createCheck(payload), config === null || config === void 0 ? void 0 : config.useCreateCheck, options),
        useUpdateCheck: (id, options) => useRapiniMutation(payload => requests.updateCheck(payload, id), config === null || config === void 0 ? void 0 : config.useUpdateCheck, options),
        useDeleteCheck: (id, options) => useRapiniMutation(() => requests.deleteCheck(id), config === null || config === void 0 ? void 0 : config.useDeleteCheck, options),
        useCreateCompetition: (options) => useRapiniMutation(payload => requests.createCompetition(payload), config === null || config === void 0 ? void 0 : config.useCreateCompetition, options),
        useUpdateCompetition: (id, options) => useRapiniMutation(payload => requests.updateCompetition(payload, id), config === null || config === void 0 ? void 0 : config.useUpdateCompetition, options),
        useDeleteCompetition: (id, options) => useRapiniMutation(() => requests.deleteCompetition(id), config === null || config === void 0 ? void 0 : config.useDeleteCompetition, options),
        useCreateHostGroup: (options) => useRapiniMutation(payload => requests.createHostGroup(payload), config === null || config === void 0 ? void 0 : config.useCreateHostGroup, options),
        useUpdateHostGroup: (id, options) => useRapiniMutation(payload => requests.updateHostGroup(payload, id), config === null || config === void 0 ? void 0 : config.useUpdateHostGroup, options),
        useDeleteHostGroup: (id, options) => useRapiniMutation(() => requests.deleteHostGroup(id), config === null || config === void 0 ? void 0 : config.useDeleteHostGroup, options),
        useCreateHost: (options) => useRapiniMutation(payload => requests.createHost(payload), config === null || config === void 0 ? void 0 : config.useCreateHost, options),
        useUpdateHost: (id, options) => useRapiniMutation(payload => requests.updateHost(payload, id), config === null || config === void 0 ? void 0 : config.useUpdateHost, options),
        useDeleteHost: (id, options) => useRapiniMutation(() => requests.deleteHost(id), config === null || config === void 0 ? void 0 : config.useDeleteHost, options),
        useCreateProperty: (options) => useRapiniMutation(payload => requests.createProperty(payload), config === null || config === void 0 ? void 0 : config.useCreateProperty, options),
        useUpdateProperty: (id, options) => useRapiniMutation(payload => requests.updateProperty(payload, id), config === null || config === void 0 ? void 0 : config.useUpdateProperty, options),
        useDeleteProperty: (id, options) => useRapiniMutation(() => requests.deleteProperty(id), config === null || config === void 0 ? void 0 : config.useDeleteProperty, options),
        useCreateRound: (options) => useRapiniMutation(payload => requests.createRound(payload), config === null || config === void 0 ? void 0 : config.useCreateRound, options),
        useUpdateRound: (id, options) => useRapiniMutation(payload => requests.updateRound(payload, id), config === null || config === void 0 ? void 0 : config.useUpdateRound, options),
        useDeleteRound: (id, options) => useRapiniMutation(() => requests.deleteRound(id), config === null || config === void 0 ? void 0 : config.useDeleteRound, options),
        useCreateService: (options) => useRapiniMutation(payload => requests.createService(payload), config === null || config === void 0 ? void 0 : config.useCreateService, options),
        useUpdateService: (id, options) => useRapiniMutation(payload => requests.updateService(payload, id), config === null || config === void 0 ? void 0 : config.useUpdateService, options),
        useDeleteService: (id, options) => useRapiniMutation(() => requests.deleteService(id), config === null || config === void 0 ? void 0 : config.useDeleteService, options),
        useCreateTeam: (options) => useRapiniMutation(payload => requests.createTeam(payload), config === null || config === void 0 ? void 0 : config.useCreateTeam, options),
        useUpdateTeam: (id, options) => useRapiniMutation(payload => requests.updateTeam(payload, id), config === null || config === void 0 ? void 0 : config.useUpdateTeam, options),
        useDeleteTeam: (id, options) => useRapiniMutation(() => requests.deleteTeam(id), config === null || config === void 0 ? void 0 : config.useDeleteTeam, options),
        useCreateUser: (options) => useRapiniMutation(payload => requests.createUser(payload), config === null || config === void 0 ? void 0 : config.useCreateUser, options),
        useUpdateUser: (id, options) => useRapiniMutation(payload => requests.updateUser(payload, id), config === null || config === void 0 ? void 0 : config.useUpdateUser, options),
        useDeleteUser: (id, options) => useRapiniMutation(() => requests.deleteUser(id), config === null || config === void 0 ? void 0 : config.useDeleteUser, options)
    };
}
