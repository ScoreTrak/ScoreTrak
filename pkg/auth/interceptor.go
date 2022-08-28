package auth

import (
	"context"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyclient"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	authv1 "go.buf.build/library/go-grpc/scoretrak/scoretrakapis/scoretrak/auth/v1"
	checkv1 "go.buf.build/library/go-grpc/scoretrak/scoretrakapis/scoretrak/check/v1"
	hostv1 "go.buf.build/library/go-grpc/scoretrak/scoretrakapis/scoretrak/host/v1"
	policyv1 "go.buf.build/library/go-grpc/scoretrak/scoretrakapis/scoretrak/policy/v1"
	propertyv1 "go.buf.build/library/go-grpc/scoretrak/scoretrakapis/scoretrak/property/v1"
	reportv1 "go.buf.build/library/go-grpc/scoretrak/scoretrakapis/scoretrak/report/v1"
	servicev1 "go.buf.build/library/go-grpc/scoretrak/scoretrakapis/scoretrak/service/v1"
	userv1 "go.buf.build/library/go-grpc/scoretrak/scoretrakapis/scoretrak/user/v1"
	"go.opentelemetry.io/otel"
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

// NewAuthInterceptor returns an instance of Interceptor. It takes in Manager struct, and policyClient as input. Policy Client allows to dynamically change authorization policies.
func NewAuthInterceptor(jwtManager *Manager, policyClient *policyclient.Client) *Interceptor {
	authMap := map[string][]authorizationMap{}

	authServicePath := fmt.Sprintf("/%s/Login", authv1.AuthService_ServiceDesc.ServiceName)
	authMap[authServicePath] = []authorizationMap{{
		role:      user.Anonymous,
		isAllowed: AlwaysAllowFunc,
	}}

	const grpcReflection = "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo"
	authMap[grpcReflection] = []authorizationMap{{
		role:      user.Anonymous,
		isAllowed: AlwaysAllowFunc,
	}}
	propertyServicePath := fmt.Sprintf("/%s/", propertyv1.PropertyService_ServiceDesc.ServiceName)
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

	serviceServicePath := fmt.Sprintf("/%s/", servicev1.ServiceService_ServiceDesc.ServiceName)
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

	hostServicePath := fmt.Sprintf("/%s/", hostv1.HostService_ServiceDesc.ServiceName)
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

	checkServicePath := fmt.Sprintf("/%s/", checkv1.CheckService_ServiceDesc.ServiceName)
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

	reportServicePath := fmt.Sprintf("/%s/", reportv1.ReportService_ServiceDesc.ServiceName)
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

	policyServicePath := fmt.Sprintf("/%s/", policyv1.PolicyService_ServiceDesc.ServiceName)
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

	userServicePath := fmt.Sprintf("/%s/", userv1.UserService_ServiceDesc.ServiceName)
	authMap[userServicePath+"Get"] = []authorizationMap{{
		role:      user.Blue,
		isAllowed: policyClient.GetAllowChangingUsernamesAndPasswords,
	}, {
		role:      user.Red,
		isAllowed: policyClient.GetAllowChangingUsernamesAndPasswords,
	}}

	return &Interceptor{jwtManager, authMap}
}

// Custom Unary( interceptor that adds claim extraction and authorization
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

// Custom Stream that allows embedding of user claims for stream grpc (Similar to what describe in: https://stackoverflow.com/questions/60982406/how-to-safely-add-values-to-grpc-serverstream-in-interceptor)
type StreamClaimInjector struct {
	grpc.ServerStream
	Claims *UserClaims
}

func (s StreamClaimInjector) Context() context.Context {
	if s.Claims != nil {
		return context.WithValue(s.ServerStream.Context(), KeyClaim, s.Claims)
	}
	return s.ServerStream.Context()
}

// Custom Stream interceptor that adds claim extraction and authorization
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

// authorize takes in context, extracts roles from the context if there are any, and ensures that a given roles has rights to access a given method. If a given role has no access, it returns permission denied error.
func (interceptor *Interceptor) authorize(ctx context.Context, method string) (claims *UserClaims, err error) {
	_, span := otel.Tracer("scoretrak/master").Start(ctx, "Authorize JWT")
	defer span.End()
	role := user.Anonymous
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values := md["authorization"]
		if len(values) != 0 {
			accessToken := values[0]
			claims, err = interceptor.jwtManager.Verify(accessToken)
			if err != nil {
				return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
			}
			role = claims.Role
		}
	}
	if role == user.Black {
		return
	}
	for i := range interceptor.accessibleRoles[method] {
		if (role == interceptor.accessibleRoles[method][i].role || user.Anonymous == interceptor.accessibleRoles[method][i].role) && interceptor.accessibleRoles[method][i].isAllowed() {
			return
		}
	}
	return nil, status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
