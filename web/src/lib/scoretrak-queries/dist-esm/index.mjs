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
    readCheckRounds: (id) => ["readCheckRounds", id],
    readCheckServices: (id) => ["readCheckServices", id],
    readCheckTeam: (id) => ["readCheckTeam", id],
    listCompetition: (page, itemsPerPage) => ["listCompetition", nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readCompetition: (id) => ["readCompetition", id],
    listCompetitionReports: (id, page, itemsPerPage) => ["listCompetitionReports", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    listCompetitionServices: (id, page, itemsPerPage) => ["listCompetitionServices", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    listCompetitionTeams: (id, page, itemsPerPage) => ["listCompetitionTeams", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    listHostService: (page, itemsPerPage) => ["listHostService", nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readHostService: (id) => ["readHostService", id],
    listHostServiceChecks: (id, page, itemsPerPage) => ["listHostServiceChecks", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readHostServiceHosts: (id) => ["readHostServiceHosts", id],
    listHostServiceProperties: (id, page, itemsPerPage) => ["listHostServiceProperties", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readHostServiceTeam: (id) => ["readHostServiceTeam", id],
    listHost: (page, itemsPerPage) => ["listHost", nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readHost: (id) => ["readHost", id],
    listHostHostService: (id, page, itemsPerPage) => ["listHostHostService", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readHostServices: (id) => ["readHostServices", id],
    readHostTeam: (id) => ["readHostTeam", id],
    listProperty: (page, itemsPerPage) => ["listProperty", nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readProperty: (id) => ["readProperty", id],
    readPropertyServices: (id) => ["readPropertyServices", id],
    readPropertyTeam: (id) => ["readPropertyTeam", id],
    listReport: (page, itemsPerPage) => ["listReport", nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readReport: (id) => ["readReport", id],
    listRound: (page, itemsPerPage) => ["listRound", nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readRound: (id) => ["readRound", id],
    listRoundChecks: (id, page, itemsPerPage) => ["listRoundChecks", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    listService: (page, itemsPerPage) => ["listService", nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readService: (id) => ["readService", id],
    readServiceCompetition: (id) => ["readServiceCompetition", id],
    listServiceHosts: (id, page, itemsPerPage) => ["listServiceHosts", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    listTeam: (page, itemsPerPage) => ["listTeam", nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readTeam: (id) => ["readTeam", id],
    listTeamChecks: (id, page, itemsPerPage) => ["listTeamChecks", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    readTeamCompetition: (id) => ["readTeamCompetition", id],
    listTeamHosts: (id, page, itemsPerPage) => ["listTeamHosts", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    listTeamHostservices: (id, page, itemsPerPage) => ["listTeamHostservices", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)],
    listTeamProperties: (id, page, itemsPerPage) => ["listTeamProperties", id, nullIfUndefined(page), nullIfUndefined(itemsPerPage)]
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
        readCheckRounds: (id) => axios.request({
            method: "get",
            url: `/checks/${id}/rounds`
        }).then(res => res.data),
        readCheckServices: (id) => axios.request({
            method: "get",
            url: `/checks/${id}/services`
        }).then(res => res.data),
        readCheckTeam: (id) => axios.request({
            method: "get",
            url: `/checks/${id}/team`
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
        listCompetitionReports: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/competitions/${id}/reports`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        listCompetitionServices: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/competitions/${id}/services`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        listCompetitionTeams: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/competitions/${id}/teams`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        listHostService: (page, itemsPerPage) => axios.request({
            method: "get",
            url: `/host-services`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        createHostService: (payload) => axios.request({
            method: "post",
            url: `/host-services`,
            data: payload
        }).then(res => res.data),
        readHostService: (id) => axios.request({
            method: "get",
            url: `/host-services/${id}`
        }).then(res => res.data),
        deleteHostService: (id) => axios.request({
            method: "delete",
            url: `/host-services/${id}`
        }).then(res => res.data),
        updateHostService: (payload, id) => axios.request({
            method: "patch",
            url: `/host-services/${id}`,
            data: payload
        }).then(res => res.data),
        listHostServiceChecks: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/host-services/${id}/checks`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        readHostServiceHosts: (id) => axios.request({
            method: "get",
            url: `/host-services/${id}/hosts`
        }).then(res => res.data),
        listHostServiceProperties: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/host-services/${id}/properties`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        readHostServiceTeam: (id) => axios.request({
            method: "get",
            url: `/host-services/${id}/team`
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
        listHostHostService: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/hosts/${id}/host-service`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        readHostServices: (id) => axios.request({
            method: "get",
            url: `/hosts/${id}/services`
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
        readPropertyServices: (id) => axios.request({
            method: "get",
            url: `/properties/${id}/services`
        }).then(res => res.data),
        readPropertyTeam: (id) => axios.request({
            method: "get",
            url: `/properties/${id}/team`
        }).then(res => res.data),
        listReport: (page, itemsPerPage) => axios.request({
            method: "get",
            url: `/reports`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        createReport: (payload) => axios.request({
            method: "post",
            url: `/reports`,
            data: payload
        }).then(res => res.data),
        readReport: (id) => axios.request({
            method: "get",
            url: `/reports/${id}`
        }).then(res => res.data),
        deleteReport: (id) => axios.request({
            method: "delete",
            url: `/reports/${id}`
        }).then(res => res.data),
        updateReport: (payload, id) => axios.request({
            method: "patch",
            url: `/reports/${id}`,
            data: payload
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
        readServiceCompetition: (id) => axios.request({
            method: "get",
            url: `/services/${id}/competition`
        }).then(res => res.data),
        listServiceHosts: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/services/${id}/hosts`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
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
        listTeamChecks: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/teams/${id}/checks`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
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
        listTeamHostservices: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/teams/${id}/hostservices`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data),
        listTeamProperties: (id, page, itemsPerPage) => axios.request({
            method: "get",
            url: `/teams/${id}/properties`,
            params: Object.assign(Object.assign({}, (page !== undefined ? { page } : undefined)), (itemsPerPage !== undefined ? { itemsPerPage } : undefined)),
            paramsSerializer: config === null || config === void 0 ? void 0 : config.paramsSerializer
        }).then(res => res.data)
    };
}
function makeQueries(requests) {
    return {
        useListCheck: (page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listCheck(page, itemsPerPage), queryFn: () => requests.listCheck(page, itemsPerPage) }, options)),
        useReadCheck: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readCheck(id), queryFn: () => requests.readCheck(id) }, options)),
        useReadCheckRounds: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readCheckRounds(id), queryFn: () => requests.readCheckRounds(id) }, options)),
        useReadCheckServices: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readCheckServices(id), queryFn: () => requests.readCheckServices(id) }, options)),
        useReadCheckTeam: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readCheckTeam(id), queryFn: () => requests.readCheckTeam(id) }, options)),
        useListCompetition: (page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listCompetition(page, itemsPerPage), queryFn: () => requests.listCompetition(page, itemsPerPage) }, options)),
        useReadCompetition: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readCompetition(id), queryFn: () => requests.readCompetition(id) }, options)),
        useListCompetitionReports: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listCompetitionReports(id, page, itemsPerPage), queryFn: () => requests.listCompetitionReports(id, page, itemsPerPage) }, options)),
        useListCompetitionServices: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listCompetitionServices(id, page, itemsPerPage), queryFn: () => requests.listCompetitionServices(id, page, itemsPerPage) }, options)),
        useListCompetitionTeams: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listCompetitionTeams(id, page, itemsPerPage), queryFn: () => requests.listCompetitionTeams(id, page, itemsPerPage) }, options)),
        useListHostService: (page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listHostService(page, itemsPerPage), queryFn: () => requests.listHostService(page, itemsPerPage) }, options)),
        useReadHostService: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readHostService(id), queryFn: () => requests.readHostService(id) }, options)),
        useListHostServiceChecks: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listHostServiceChecks(id, page, itemsPerPage), queryFn: () => requests.listHostServiceChecks(id, page, itemsPerPage) }, options)),
        useReadHostServiceHosts: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readHostServiceHosts(id), queryFn: () => requests.readHostServiceHosts(id) }, options)),
        useListHostServiceProperties: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listHostServiceProperties(id, page, itemsPerPage), queryFn: () => requests.listHostServiceProperties(id, page, itemsPerPage) }, options)),
        useReadHostServiceTeam: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readHostServiceTeam(id), queryFn: () => requests.readHostServiceTeam(id) }, options)),
        useListHost: (page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listHost(page, itemsPerPage), queryFn: () => requests.listHost(page, itemsPerPage) }, options)),
        useReadHost: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readHost(id), queryFn: () => requests.readHost(id) }, options)),
        useListHostHostService: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listHostHostService(id, page, itemsPerPage), queryFn: () => requests.listHostHostService(id, page, itemsPerPage) }, options)),
        useReadHostServices: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readHostServices(id), queryFn: () => requests.readHostServices(id) }, options)),
        useReadHostTeam: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readHostTeam(id), queryFn: () => requests.readHostTeam(id) }, options)),
        useListProperty: (page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listProperty(page, itemsPerPage), queryFn: () => requests.listProperty(page, itemsPerPage) }, options)),
        useReadProperty: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readProperty(id), queryFn: () => requests.readProperty(id) }, options)),
        useReadPropertyServices: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readPropertyServices(id), queryFn: () => requests.readPropertyServices(id) }, options)),
        useReadPropertyTeam: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readPropertyTeam(id), queryFn: () => requests.readPropertyTeam(id) }, options)),
        useListReport: (page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listReport(page, itemsPerPage), queryFn: () => requests.listReport(page, itemsPerPage) }, options)),
        useReadReport: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readReport(id), queryFn: () => requests.readReport(id) }, options)),
        useListRound: (page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listRound(page, itemsPerPage), queryFn: () => requests.listRound(page, itemsPerPage) }, options)),
        useReadRound: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readRound(id), queryFn: () => requests.readRound(id) }, options)),
        useListRoundChecks: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listRoundChecks(id, page, itemsPerPage), queryFn: () => requests.listRoundChecks(id, page, itemsPerPage) }, options)),
        useListService: (page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listService(page, itemsPerPage), queryFn: () => requests.listService(page, itemsPerPage) }, options)),
        useReadService: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readService(id), queryFn: () => requests.readService(id) }, options)),
        useReadServiceCompetition: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readServiceCompetition(id), queryFn: () => requests.readServiceCompetition(id) }, options)),
        useListServiceHosts: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listServiceHosts(id, page, itemsPerPage), queryFn: () => requests.listServiceHosts(id, page, itemsPerPage) }, options)),
        useListTeam: (page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listTeam(page, itemsPerPage), queryFn: () => requests.listTeam(page, itemsPerPage) }, options)),
        useReadTeam: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readTeam(id), queryFn: () => requests.readTeam(id) }, options)),
        useListTeamChecks: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listTeamChecks(id, page, itemsPerPage), queryFn: () => requests.listTeamChecks(id, page, itemsPerPage) }, options)),
        useReadTeamCompetition: (id, options) => useQuery(Object.assign({ queryKey: queryKeys.readTeamCompetition(id), queryFn: () => requests.readTeamCompetition(id) }, options)),
        useListTeamHosts: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listTeamHosts(id, page, itemsPerPage), queryFn: () => requests.listTeamHosts(id, page, itemsPerPage) }, options)),
        useListTeamHostservices: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listTeamHostservices(id, page, itemsPerPage), queryFn: () => requests.listTeamHostservices(id, page, itemsPerPage) }, options)),
        useListTeamProperties: (id, page, itemsPerPage, options) => useQuery(Object.assign({ queryKey: queryKeys.listTeamProperties(id, page, itemsPerPage), queryFn: () => requests.listTeamProperties(id, page, itemsPerPage) }, options))
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
        useCreateHostService: (options) => useRapiniMutation(payload => requests.createHostService(payload), config === null || config === void 0 ? void 0 : config.useCreateHostService, options),
        useUpdateHostService: (id, options) => useRapiniMutation(payload => requests.updateHostService(payload, id), config === null || config === void 0 ? void 0 : config.useUpdateHostService, options),
        useDeleteHostService: (id, options) => useRapiniMutation(() => requests.deleteHostService(id), config === null || config === void 0 ? void 0 : config.useDeleteHostService, options),
        useCreateHost: (options) => useRapiniMutation(payload => requests.createHost(payload), config === null || config === void 0 ? void 0 : config.useCreateHost, options),
        useUpdateHost: (id, options) => useRapiniMutation(payload => requests.updateHost(payload, id), config === null || config === void 0 ? void 0 : config.useUpdateHost, options),
        useDeleteHost: (id, options) => useRapiniMutation(() => requests.deleteHost(id), config === null || config === void 0 ? void 0 : config.useDeleteHost, options),
        useCreateProperty: (options) => useRapiniMutation(payload => requests.createProperty(payload), config === null || config === void 0 ? void 0 : config.useCreateProperty, options),
        useUpdateProperty: (id, options) => useRapiniMutation(payload => requests.updateProperty(payload, id), config === null || config === void 0 ? void 0 : config.useUpdateProperty, options),
        useDeleteProperty: (id, options) => useRapiniMutation(() => requests.deleteProperty(id), config === null || config === void 0 ? void 0 : config.useDeleteProperty, options),
        useCreateReport: (options) => useRapiniMutation(payload => requests.createReport(payload), config === null || config === void 0 ? void 0 : config.useCreateReport, options),
        useUpdateReport: (id, options) => useRapiniMutation(payload => requests.updateReport(payload, id), config === null || config === void 0 ? void 0 : config.useUpdateReport, options),
        useDeleteReport: (id, options) => useRapiniMutation(() => requests.deleteReport(id), config === null || config === void 0 ? void 0 : config.useDeleteReport, options),
        useCreateRound: (options) => useRapiniMutation(payload => requests.createRound(payload), config === null || config === void 0 ? void 0 : config.useCreateRound, options),
        useUpdateRound: (id, options) => useRapiniMutation(payload => requests.updateRound(payload, id), config === null || config === void 0 ? void 0 : config.useUpdateRound, options),
        useDeleteRound: (id, options) => useRapiniMutation(() => requests.deleteRound(id), config === null || config === void 0 ? void 0 : config.useDeleteRound, options),
        useCreateService: (options) => useRapiniMutation(payload => requests.createService(payload), config === null || config === void 0 ? void 0 : config.useCreateService, options),
        useUpdateService: (id, options) => useRapiniMutation(payload => requests.updateService(payload, id), config === null || config === void 0 ? void 0 : config.useUpdateService, options),
        useDeleteService: (id, options) => useRapiniMutation(() => requests.deleteService(id), config === null || config === void 0 ? void 0 : config.useDeleteService, options),
        useCreateTeam: (options) => useRapiniMutation(payload => requests.createTeam(payload), config === null || config === void 0 ? void 0 : config.useCreateTeam, options),
        useUpdateTeam: (id, options) => useRapiniMutation(payload => requests.updateTeam(payload, id), config === null || config === void 0 ? void 0 : config.useUpdateTeam, options),
        useDeleteTeam: (id, options) => useRapiniMutation(() => requests.deleteTeam(id), config === null || config === void 0 ? void 0 : config.useDeleteTeam, options)
    };
}
