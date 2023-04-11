package auth

import (
	"context"

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

// // NewAuthInterceptor returns an instance of Interceptor. It takes in Manager struct, and policyClient as input. Policy Client allows to dynamically change authorization policies.
// func NewAuthInterceptor(jwtManager *Manager, policyClient *policyclient.Client) *Interceptor {
// 	authMap := map[string][]authorizationMap{}

// 	healthServicePath := fmt.Sprintf("/%s", healthv1grpc.Health_ServiceDesc.ServiceName)
// 	authMap[healthServicePath] = []authorizationMap{{
// 		role:      user.Anonymous,
// 		isAllowed: AlwaysAllowFunc,
// 	}}

// 	authV1ServicePath := fmt.Sprintf("/%s/", authv1grpc.AuthService_ServiceDesc.ServiceName)
// 	authMap[authV1ServicePath+"Login"] = []authorizationMap{{
// 		role:      user.Anonymous,
// 		isAllowed: AlwaysAllowFunc,
// 	}}

// 	authV2ServicePath := fmt.Sprintf("/%s/", authv2grpc.AuthService_ServiceDesc.ServiceName)
// 	authMap[authV2ServicePath+"Login"] = []authorizationMap{{
// 		role:      user.Anonymous,
// 		isAllowed: AlwaysAllowFunc,
// 	}}

// 	const grpcReflection = "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo"
// 	authMap[grpcReflection] = []authorizationMap{{
// 		role:      user.Anonymous,
// 		isAllowed: AlwaysAllowFunc,
// 	}}

// 	propertyV1ServicePath := fmt.Sprintf("/%s/", propertyv1grpc.PropertyService_ServiceDesc.ServiceName)
// 	authMap[propertyV1ServicePath+"GetByServiceIDKey"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}}
// 	authMap[propertyV1ServicePath+"GetAllByServiceID"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}}
// 	authMap[propertyV1ServicePath+"Update"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}}

// 	propertyV2ServicePath := fmt.Sprintf("/%s/", propertyv2grpc.PropertyService_ServiceDesc.ServiceName)
// 	authMap[propertyV2ServicePath+"GetByServiceIDKey"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}}
// 	authMap[propertyV2ServicePath+"GetAllByServiceID"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}}
// 	authMap[propertyV2ServicePath+"Update"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}}

// 	serviceV1ServicePath := fmt.Sprintf("/%s/", servicev1grpc.ServiceService_ServiceDesc.ServiceName)
// 	authMap[serviceV1ServicePath+"GetAll"] = []authorizationMap{{
// 		role:      user.Red,
// 		isAllowed: policyClient.GetAllowRedTeamLaunchingServiceTestsManually,
// 	}}
// 	authMap[serviceV1ServicePath+"GetByID"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}}
// 	authMap[serviceV1ServicePath+"TestService"] = []authorizationMap{{
// 		role:      user.Red,
// 		isAllowed: policyClient.GetAllowRedTeamLaunchingServiceTestsManually,
// 	}}

// 	serviceV2ServicePath := fmt.Sprintf("/%s/", servicev2grpc.ServiceService_ServiceDesc.ServiceName)
// 	authMap[serviceV2ServicePath+"GetAll"] = []authorizationMap{{
// 		role:      user.Red,
// 		isAllowed: policyClient.GetAllowRedTeamLaunchingServiceTestsManually,
// 	}}
// 	authMap[serviceV2ServicePath+"GetByID"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}}
// 	authMap[serviceV2ServicePath+"TestService"] = []authorizationMap{{
// 		role:      user.Red,
// 		isAllowed: policyClient.GetAllowRedTeamLaunchingServiceTestsManually,
// 	}}

// 	hostV1ServicePath := fmt.Sprintf("/%s/", hostv1grpc.HostService_ServiceDesc.ServiceName)
// 	authMap[hostV1ServicePath+"Update"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}}
// 	authMap[hostV1ServicePath+"GetByID"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}}

// 	hostV2ServicePath := fmt.Sprintf("/%s/", hostv2grpc.HostService_ServiceDesc.ServiceName)
// 	authMap[hostV2ServicePath+"HostServiceUpdate"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}}
// 	authMap[hostV2ServicePath+"HostServiceGetByID"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}}

// 	checkV1ServicePath := fmt.Sprintf("/%s/", checkv1grpc.CheckService_ServiceDesc.ServiceName)
// 	authMap[checkV1ServicePath+"GetByRoundServiceID"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}}
// 	authMap[checkV1ServicePath+"GetAllByServiceID"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}}

// 	checkV2ServicePath := fmt.Sprintf("/%s/", checkv2grpc.CheckService_ServiceDesc.ServiceName)
// 	authMap[checkV2ServicePath+"GetByRoundServiceID"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}}
// 	authMap[checkV2ServicePath+"GetAllByServiceID"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}}

// 	reportV1ServicePath := fmt.Sprintf("/%s/", reportv1grpc.ReportService_ServiceDesc.ServiceName)
// 	authMap[reportV1ServicePath+"Get"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Anonymous,
// 		isAllowed: policyClient.GetAllowUnauthenticatedUsers,
// 	}}
// 	authMap[reportV1ServicePath+"GetUnary"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Anonymous,
// 		isAllowed: policyClient.GetAllowUnauthenticatedUsers,
// 	}}

// 	reportV2ServicePath := fmt.Sprintf("/%s/", reportv2grpc.ReportService_ServiceDesc.ServiceName)
// 	authMap[reportV2ServicePath+"Get"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Anonymous,
// 		isAllowed: policyClient.GetAllowUnauthenticatedUsers,
// 	}}
// 	authMap[reportV2ServicePath+"GetUnary"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Anonymous,
// 		isAllowed: policyClient.GetAllowUnauthenticatedUsers,
// 	}}

// 	policyV1ServicePath := fmt.Sprintf("/%s/", policyv1grpc.PolicyService_ServiceDesc.ServiceName)
// 	authMap[policyV1ServicePath+"Get"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Anonymous,
// 		isAllowed: policyClient.GetAllowUnauthenticatedUsers,
// 	}}
// 	authMap[policyV1ServicePath+"GetUnary"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Anonymous,
// 		isAllowed: policyClient.GetAllowUnauthenticatedUsers,
// 	}}

// 	policyV2ServicePath := fmt.Sprintf("/%s/", policyv2grpc.PolicyService_ServiceDesc.ServiceName)
// 	authMap[policyV2ServicePath+"Get"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Anonymous,
// 		isAllowed: policyClient.GetAllowUnauthenticatedUsers,
// 	}}
// 	authMap[policyV2ServicePath+"GetUnary"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: AlwaysAllowFunc,
// 	}, {
// 		role:      user.Anonymous,
// 		isAllowed: policyClient.GetAllowUnauthenticatedUsers,
// 	}}

// 	userV1ServicePath := fmt.Sprintf("/%s/", userv1grpc.UserService_ServiceDesc.ServiceName)
// 	authMap[userV1ServicePath+"Get"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: policyClient.GetAllowChangingUsernamesAndPasswords,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: policyClient.GetAllowChangingUsernamesAndPasswords,
// 	}}

// 	userV2ServicePath := fmt.Sprintf("/%s/", userv2grpc.UserService_ServiceDesc.ServiceName)
// 	authMap[userV2ServicePath+"Get"] = []authorizationMap{{
// 		role:      user.Blue,
// 		isAllowed: policyClient.GetAllowChangingUsernamesAndPasswords,
// 	}, {
// 		role:      user.Red,
// 		isAllowed: policyClient.GetAllowChangingUsernamesAndPasswords,
// 	}}

// 	return &Interceptor{jwtManager, authMap}
// }

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
	// role := user.Anonymous
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values := md["authorization"]
		if len(values) != 0 {
			accessToken := values[0]
			claims, err = interceptor.jwtManager.Verify(ctx, accessToken)
			if err != nil {
				return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
			}
			// role = claims.Role
		}
	}
	// if role == user.Black {
	// 	return
	// }
	// for i := range interceptor.accessibleRoles[method] {
	// 	if (role == interceptor.accessibleRoles[method][i].role || user.Anonymous == interceptor.accessibleRoles[method][i].role) && interceptor.accessibleRoles[method][i].isAllowed() {
	// 		return
	// 	}
	// }
	return nil, status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
