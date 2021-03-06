package auth

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policy_client"
	authpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/auth/v1"
	checkpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/check/v1"
	hostpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/host/v1"
	policypb "github.com/ScoreTrak/ScoreTrak/pkg/proto/policy/v1"
	propertypb "github.com/ScoreTrak/ScoreTrak/pkg/proto/property/v1"
	reportpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/report/v1"
	servicepb "github.com/ScoreTrak/ScoreTrak/pkg/proto/service/v1"
	userpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/user/v1"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Interceptor struct {
	jwtManager      *Manager
	accessibleRoles map[string][]authorizationMap
}

type isAllowedFunc func() bool

func AlwaysAllowFunc() bool {
	return true
}

type authorizationMap struct {
	role      string
	isAllowed isAllowedFunc
}

//NewAuthInterceptor returns an instance of Interceptor. It takes in Manager struct, and policyClient as input. Policy Client allows to dynamically change authorization policies.
func NewAuthInterceptor(jwtManager *Manager, policyClient *policy_client.Client) *Interceptor {
	authMap := map[string][]authorizationMap{}

	authServicePath := fmt.Sprintf("/%s/Login", authpb.AuthService_ServiceDesc.ServiceName)
	authMap[authServicePath] = []authorizationMap{{
		role:      user.Anonymous,
		isAllowed: AlwaysAllowFunc,
	}}

	const grpcReflection = "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo"
	authMap[grpcReflection] = []authorizationMap{{
		role:      user.Anonymous,
		isAllowed: AlwaysAllowFunc,
	}}
	propertyServicePath := fmt.Sprintf("/%s/", propertypb.PropertyService_ServiceDesc.ServiceName)
	authMap[propertyServicePath+"GetByServiceIDKey"] = []authorizationMap{{
		role:      user.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      user.Red,
		isAllowed: AlwaysAllowFunc,
	}}
	authMap[propertyServicePath+"GetAllByServiceID"] = []authorizationMap{{
		role:      user.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      user.Red,
		isAllowed: AlwaysAllowFunc,
	}}

	authMap[propertyServicePath+"Update"] = []authorizationMap{{
		role:      user.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      user.Red,
		isAllowed: AlwaysAllowFunc,
	}}

	serviceServicePath := fmt.Sprintf("/%s/", servicepb.ServiceService_ServiceDesc.ServiceName)
	authMap[serviceServicePath+"GetAll"] = []authorizationMap{{
		role:      user.Red,
		isAllowed: policyClient.GetAllowRedTeamLaunchingServiceTestsManually,
	}}
	authMap[serviceServicePath+"GetByID"] = []authorizationMap{{
		role:      user.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      user.Red,
		isAllowed: AlwaysAllowFunc,
	}}

	authMap[serviceServicePath+"TestService"] = []authorizationMap{{
		role:      user.Red,
		isAllowed: policyClient.GetAllowRedTeamLaunchingServiceTestsManually,
	}}

	hostServicePath := fmt.Sprintf("/%s/", hostpb.HostService_ServiceDesc.ServiceName)
	authMap[hostServicePath+"Update"] = []authorizationMap{{
		role:      user.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      user.Red,
		isAllowed: AlwaysAllowFunc,
	}}
	authMap[hostServicePath+"GetByID"] = []authorizationMap{{
		role:      user.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      user.Red,
		isAllowed: AlwaysAllowFunc,
	}}

	checkServicePath := fmt.Sprintf("/%s/", checkpb.CheckService_ServiceDesc.ServiceName)
	authMap[checkServicePath+"GetByRoundServiceID"] = []authorizationMap{{
		role:      user.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      user.Red,
		isAllowed: AlwaysAllowFunc,
	}}
	authMap[checkServicePath+"GetAllByServiceID"] = []authorizationMap{{
		role:      user.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      user.Red,
		isAllowed: AlwaysAllowFunc,
	}}

	reportServicePath := fmt.Sprintf("/%s/", reportpb.ReportService_ServiceDesc.ServiceName)
	authMap[reportServicePath+"Get"] = []authorizationMap{{
		role:      user.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      user.Red,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      user.Anonymous,
		isAllowed: policyClient.GetAllowUnauthenticatedUsers,
	}}

	policyServicePath := fmt.Sprintf("/%s/", policypb.PolicyService_ServiceDesc.ServiceName)
	authMap[policyServicePath+"Get"] = []authorizationMap{{
		role:      user.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      user.Red,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      user.Anonymous,
		isAllowed: policyClient.GetAllowUnauthenticatedUsers,
	}}

	userServicePath := fmt.Sprintf("/%s/", userpb.UserService_ServiceDesc.ServiceName)
	authMap[userServicePath+"Get"] = []authorizationMap{{
		role:      user.Blue,
		isAllowed: policyClient.GetAllowChangingUsernamesAndPasswords,
	}, {
		role:      user.Red,
		isAllowed: policyClient.GetAllowChangingUsernamesAndPasswords,
	}}

	return &Interceptor{jwtManager, authMap}
}

//Custom Unary( interceptor that adds claim extraction and authorization
func (interceptor *Interceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		claims, err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		if claims != nil {
			ctx = context.WithValue(ctx, KeyClaim, claims)
		}
		return handler(ctx, req)
	}
}

//Custom Stream that allows embedding of user claims for stream grpc (Similar to what describe in: https://stackoverflow.com/questions/60982406/how-to-safely-add-values-to-grpc-serverstream-in-interceptor)
type StreamClaimInjector struct {
	grpc.ServerStream
	Claims *UserClaims
}

func (s StreamClaimInjector) Context() context.Context {
	if s.Claims != nil {
		return context.WithValue(s.ServerStream.Context(), KeyClaim, s.Claims)
	} else {
		return s.ServerStream.Context()
	}
}

//Custom Stream interceptor that adds claim extraction and authorization
func (interceptor *Interceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		claims, err := interceptor.authorize(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}
		return handler(srv, StreamClaimInjector{stream, claims})
	}
}

//authorize takes in context, extracts roles from the context if there are any, and ensures that a given roles has rights to access a given method. If a given role has no access, it returns permission denied error.
func (interceptor *Interceptor) authorize(ctx context.Context, method string) (claims *UserClaims, err error) {
	r := user.Anonymous
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values := md["authorization"]
		if len(values) != 0 {
			accessToken := values[0]
			claims, err = interceptor.jwtManager.Verify(accessToken)
			if err != nil {
				return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
			}
			r = claims.Role
		}
	}
	if r == user.Black {
		return
	} else {
		for i := range interceptor.accessibleRoles[method] {
			if (r == interceptor.accessibleRoles[method][i].role || user.Anonymous == interceptor.accessibleRoles[method][i].role) && interceptor.accessibleRoles[method][i].isAllowed() {
				return
			}
		}
	}
	return nil, status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
