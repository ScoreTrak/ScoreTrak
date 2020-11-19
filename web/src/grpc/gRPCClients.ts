import {token} from './token/token'

import {AuthServiceClient} from "./pkg/auth/AuthServiceClientPb"
import {CheckServiceClient} from "./pkg/check/checkpb/CheckServiceClientPb";
import {CompetitionServiceClient} from "./pkg/competition/competitionpb/CompetitionServiceClientPb";
import {DynamicConfigServiceClient, StaticConfigServiceClient} from "./pkg/config/configpb/ConfigServiceClientPb";
import {HostServiceClient} from "./pkg/host/hostpb/HostServiceClientPb";
import {HostGroupServiceClient} from "./pkg/host_group/host_grouppb/Host_groupServiceClientPb";
import {PolicyServiceClient} from "./pkg/policy/policypb/PolicyServiceClientPb";
import {PropertyServiceClient} from "./pkg/property/propertypb/PropertyServiceClientPb";
import {ReportServiceClient} from "./pkg/report/reportpb/ReportServiceClientPb";
import {RoundServiceClient} from "./pkg/round/roundpb/RoundServiceClientPb";
import {ServiceServiceClient} from "./pkg/service/servicepb/ServiceServiceClientPb";
import {ServiceGroupServiceClient} from "./pkg/service_group/service_grouppb/Service_groupServiceClientPb";
import {TeamServiceClient} from "./pkg/team/teampb/TeamServiceClientPb";
import {UserServiceClient} from "./pkg/user/userpb/UserServiceClientPb";

const serverAddress = process.env.PUBLIC_URL

export type GRPCClients = {
    authClient: AuthServiceClient;
    checkClient: CheckServiceClient;
    competitionClient: CompetitionServiceClient;
    dynamicConfigClient: DynamicConfigServiceClient;
    staticConfigClient: StaticConfigServiceClient;
    hostClient: HostServiceClient;
    hostGroupClient: HostGroupServiceClient;
    policyClient: PolicyServiceClient;
    propertyClient: PropertyServiceClient;
    reportClient: ReportServiceClient;
    roundClient: RoundServiceClient;
    serviceClient: ServiceServiceClient;
    serviceGroupClient: ServiceGroupServiceClient;
    teamClient: TeamServiceClient;
    userClient: UserServiceClient;
};

class AuthInterceptor {
    intercept = (request: any, invoker: any) => {
        if (token.isAValidToken()) {
            const metadata = request.getMetadata(undefined, undefined)
            metadata.authorization = token.getToken()
        }
        return invoker(request)
    }
}

const options = {
    unaryInterceptors: [new AuthInterceptor()],
    streamInterceptors: [new AuthInterceptor()]
}

export const gRPCClients: GRPCClients = {
    authClient: new AuthServiceClient(serverAddress, null, options),
    checkClient: new CheckServiceClient(serverAddress, null, options),
    competitionClient: new CompetitionServiceClient(serverAddress, null, options),
    dynamicConfigClient: new DynamicConfigServiceClient(serverAddress, null, options),
    staticConfigClient: new StaticConfigServiceClient(serverAddress, null, options),
    hostClient: new HostServiceClient(serverAddress, null, options),
    hostGroupClient: new HostGroupServiceClient(serverAddress, null, options),
    policyClient: new PolicyServiceClient(serverAddress, null, options),
    propertyClient: new PropertyServiceClient(serverAddress, null, options),
    reportClient: new ReportServiceClient(serverAddress, null, options),
    roundClient: new RoundServiceClient(serverAddress, null, options),
    serviceClient: new ServiceServiceClient(serverAddress, null, options),
    serviceGroupClient: new ServiceGroupServiceClient(serverAddress, null, options),
    teamClient: new TeamServiceClient(serverAddress, null, options),
    userClient: new UserServiceClient(serverAddress, null, options)
};