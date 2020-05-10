package api

import (
	"ScoreTrak/pkg/logging"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logging.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.Use(TokenVerify)

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"GetCheck",
		strings.ToUpper("Get"),
		"/check/{TeamId}/{RoundId}/{ServiceID}",
		GetCheck,
	},

	Route{
		"GetChecks",
		strings.ToUpper("Get"),
		"/check/{TeamId}/{RoundId}",
		GetChecks,
	},

	Route{
		"AddEngineProperties",
		strings.ToUpper("Put"),
		"/engine",
		AddConfigProperties,
	},

	Route{
		"GetEngineProperties",
		strings.ToUpper("Get"),
		"/engine",
		GetConfigProperties,
	},

	Route{
		"AddHost",
		strings.ToUpper("Post"),
		"/host",
		AddHost,
	},

	Route{
		"DeleteHost",
		strings.ToUpper("Delete"),
		"/host/{hostId}",
		DeleteHost,
	},

	Route{
		"GetHost",
		strings.ToUpper("Get"),
		"/host/{hostId}",
		GetHost,
	},

	Route{
		"GetHosts",
		strings.ToUpper("Get"),
		"/host",
		GetHosts,
	},

	Route{
		"UpdateHost",
		strings.ToUpper("Put"),
		"/host/{hostId}",
		UpdateHost,
	},

	Route{
		"AddHostGroup",
		strings.ToUpper("Post"),
		"/host_group",
		AddHostGroup,
	},

	Route{
		"DeleteHostGroup",
		strings.ToUpper("Delete"),
		"/host_group/{host_groupId}",
		DeleteHostGroup,
	},

	Route{
		"GetHostGroup",
		strings.ToUpper("Get"),
		"/host_group/{host_groupId}",
		GetHostGroup,
	},

	Route{
		"GetHostGroups",
		strings.ToUpper("Get"),
		"/host_group",
		GetHostGroups,
	},

	Route{
		"UpdateHostGroup",
		strings.ToUpper("Put"),
		"/host_group/{host_groupId}",
		UpdateHostGroup,
	},

	Route{
		"AddProprty",
		strings.ToUpper("Post"),
		"/property",
		AddProprty,
	},

	Route{
		"DeleteProperty",
		strings.ToUpper("Delete"),
		"/property/{propertyId}",
		DeleteProperty,
	},

	Route{
		"GetProperties",
		strings.ToUpper("Get"),
		"/property",
		GetProperties,
	},

	Route{
		"GetProperty",
		strings.ToUpper("Get"),
		"/property/{propertyId}",
		GetProperty,
	},

	Route{
		"UpdateProperty",
		strings.ToUpper("Put"),
		"/property/{propertyId}",
		UpdateProperty,
	},

	Route{
		"GetRoundNumber",
		strings.ToUpper("Get"),
		"/round",
		GetRoundNumber,
	},

	Route{
		"GetScore",
		strings.ToUpper("Get"),
		"/score/{teamID}",
		GetScore,
	},

	Route{
		"GetScores",
		strings.ToUpper("Get"),
		"/score",
		GetScores,
	},

	Route{
		"AddService",
		strings.ToUpper("Post"),
		"/service",
		AddService,
	},

	Route{
		"DeleteService",
		strings.ToUpper("Delete"),
		"/service/{serviceId}",
		DeleteService,
	},

	Route{
		"GetService",
		strings.ToUpper("Get"),
		"/service/{serviceId}",
		GetService,
	},

	Route{
		"GetServices",
		strings.ToUpper("Get"),
		"/service",
		GetServices,
	},

	Route{
		"UpdateService",
		strings.ToUpper("Put"),
		"/service/{serviceId}",
		UpdateService,
	},

	Route{
		"AddServiceGroup",
		strings.ToUpper("Post"),
		"/service_group",
		AddServiceGroup,
	},

	Route{
		"DeleteServiceGroup",
		strings.ToUpper("Delete"),
		"/service_group/{serviceGroupId}",
		DeleteServiceGroup,
	},

	Route{
		"GetServiceGroup",
		strings.ToUpper("Get"),
		"/service_group/{serviceGroupId}",
		GetServiceGroup,
	},

	Route{
		"GetServiceGroups",
		strings.ToUpper("Get"),
		"/service_group",
		GetServiceGroups,
	},

	Route{
		"UpdateServiceGroup",
		strings.ToUpper("Put"),
		"/service_group/{serviceGroupId}",
		UpdateServiceGroup,
	},

	Route{
		"AddTeam",
		strings.ToUpper("Post"),
		"/team",
		AddTeam,
	},

	Route{
		"DelteTeam",
		strings.ToUpper("Delete"),
		"/team/{TeamId}",
		DelteTeam,
	},

	Route{
		"GetTeam",
		strings.ToUpper("Get"),
		"/team/{TeamId}",
		GetTeam,
	},

	Route{
		"GetTeams",
		strings.ToUpper("Get"),
		"/team",
		GetTeams,
	},

	Route{
		"UpdateTeam",
		strings.ToUpper("Put"),
		"/team/{TeamId}",
		UpdateTeam,
	},
}
