package auth

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policy_client"
	"github.com/ScoreTrak/ScoreTrak/pkg/role"
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

func NewAuthInterceptor(jwtManager *Manager, policyClient *policy_client.Client) *Interceptor {
	authMap := map[string][]authorizationMap{}
	const propertyServicePath = "/pkg.property.propertypb.PropertyService/"
	authMap[propertyServicePath+"GetByServiceIDKey"] = []authorizationMap{{
		role:      role.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      role.Red,
		isAllowed: AlwaysAllowFunc,
	}}
	authMap[propertyServicePath+"GetAllByServiceID"] = []authorizationMap{{
		role:      role.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      role.Red,
		isAllowed: AlwaysAllowFunc,
	}}

	authMap[propertyServicePath+"Update"] = []authorizationMap{{
		role:      role.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      role.Red,
		isAllowed: AlwaysAllowFunc,
	}}

	const serviceServicePath = "/pkg.service.servicepb.ServiceService/"
	authMap[serviceServicePath+"GetByID"] = []authorizationMap{{
		role:      role.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      role.Red,
		isAllowed: AlwaysAllowFunc,
	}}

	authMap[serviceServicePath+"TestService"] = []authorizationMap{{
		role:      role.Red,
		isAllowed: policyClient.GetAllowRedTeamLaunchingServiceTestsManually,
	}}

	const hostServicePath = "/pkg.host.hostpb.HostService/"
	authMap[hostServicePath+"Update"] = []authorizationMap{{
		role:      role.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      role.Red,
		isAllowed: AlwaysAllowFunc,
	}}
	authMap[hostServicePath+"GetByID"] = []authorizationMap{{
		role:      role.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      role.Red,
		isAllowed: AlwaysAllowFunc,
	}}

	const checkServicePath = "/pkg.check.checkpb.CheckService/"
	authMap[checkServicePath+"GetByRoundServiceID"] = []authorizationMap{{
		role:      role.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      role.Red,
		isAllowed: AlwaysAllowFunc,
	}}
	authMap[checkServicePath+"GetAllByServiceID"] = []authorizationMap{{
		role:      role.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      role.Red,
		isAllowed: AlwaysAllowFunc,
	}}

	const reportServicePath = "/pkg.report.reportpb.ReportService/"
	authMap[reportServicePath+"Get"] = []authorizationMap{{
		role:      role.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      role.Red,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      role.Anonymous,
		isAllowed: policyClient.GetAllowUnauthenticatedUsers,
	}}

	const policyServicePath = "/pkg.policy.policypb.PolicyService/"
	authMap[policyServicePath+"Get"] = []authorizationMap{{
		role:      role.Blue,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      role.Red,
		isAllowed: AlwaysAllowFunc,
	}, {
		role:      role.Anonymous,
		isAllowed: policyClient.GetAllowUnauthenticatedUsers,
	}}

	const userServicePath = "/pkg.user.userpb.UserService/"
	authMap[userServicePath+"Get"] = []authorizationMap{{
		role:      role.Blue,
		isAllowed: policyClient.GetAllowChangingUsernamesAndPasswords,
	}, {
		role:      role.Red,
		isAllowed: policyClient.GetAllowChangingUsernamesAndPasswords,
	}}

	return &Interceptor{jwtManager, authMap}
}

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
			ctx = context.WithValue(ctx, "claims", claims)
		}
		return handler(ctx, req)
	}
}

//Custom Stream that allows embedding of user claims for stream grpc (Similar to what describe in: https://stackoverflow.com/questions/60982406/how-to-safely-add-values-to-grpc-serverstream-in-interceptor)
type authStream struct {
	grpc.ServerStream
	uClaims *UserClaims
}

func (s authStream) Context() context.Context {
	if s.uClaims != nil {
		return context.WithValue(s.ServerStream.Context(), "claims", s.uClaims)
	} else {
		return s.ServerStream.Context()
	}
}

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
		return handler(srv, authStream{stream, claims})
	}
}

func (interceptor *Interceptor) authorize(ctx context.Context, method string) (claims *UserClaims, err error) {
	var r string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values := md["Authorization"]
		if len(values) != 0 {
			accessToken := values[0]
			claims, err = interceptor.jwtManager.Verify(accessToken)
			if err != nil {
				return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
			}
			r = claims.Role
		}
	}
	if r == role.Black {
		return
	} else {
		for i := range interceptor.accessibleRoles[method] {
			if (r == interceptor.accessibleRoles[method][i].role || role.Anonymous == interceptor.accessibleRoles[method][i].role) && interceptor.accessibleRoles[method][i].isAllowed() {
				return
			}
		}
	}
	return nil, status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
