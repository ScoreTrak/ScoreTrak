package server

import (
	"ScoreTrak/pkg/api"
	"ScoreTrak/pkg/api/handler"
	"ScoreTrak/pkg/check"
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/host"
	"ScoreTrak/pkg/host_group"
	"ScoreTrak/pkg/property"
	"ScoreTrak/pkg/round"
	"ScoreTrak/pkg/service"
	"ScoreTrak/pkg/service_group"
	"ScoreTrak/pkg/team"
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
	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to ScoreTrak!")
}

func (ds *dserver) MapRoutes() {

	routes := Routes{
		Route{
			"Index",
			"GET",
			"/",
			Index,
		},
	}

	routes = append(routes, ds.configRoutes()...)
	routes = append(routes, ds.teamRoutes()...)
	routes = append(routes, ds.checkRoutes()...)
	routes = append(routes, ds.hostRoutes()...)
	routes = append(routes, ds.hostGroupRoutes()...)
	routes = append(routes, ds.propertyRoutes()...)
	routes = append(routes, ds.roundRoutes()...)
	routes = append(routes, ds.scoreRoutes()...)
	routes = append(routes, ds.serviceRoutes()...)
	routes = append(routes, ds.serviceGroupRoutes()...)

	for _, route := range routes {
		var hdler http.Handler
		hdler = route.HandlerFunc
		hdler = api.Logger(hdler, route.Name) //Default Logger

		ds.router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(hdler)
	}
	ds.router.Use(api.TokenVerify)
}

func (ds *dserver) teamRoutes() Routes {

	var teamSvc team.Serv

	ds.cont.Invoke(func(s team.Serv) {
		teamSvc = s
	})
	tm := handler.NewTeamController(ds.logger, teamSvc)

	teamRoutes := Routes{
		Route{
			"AddTeam",
			strings.ToUpper("Post"),
			"/team",
			tm.Store,
		},

		Route{
			"DeleteTeam",
			strings.ToUpper("Delete"),
			"/team/{TeamID}",
			tm.Delete,
		},

		Route{
			"GetTeam",
			strings.ToUpper("Get"),
			"/team/{TeamID}",
			tm.GetByID,
		},

		Route{
			"GetTeams",
			strings.ToUpper("Get"),
			"/team",
			tm.GetAll,
		},

		Route{
			"UpdateTeam",
			strings.ToUpper("Put"),
			"/team/{TeamID}",
			tm.Update,
		},
	}

	return teamRoutes
}

func (ds *dserver) configRoutes() Routes {
	var configSvc config.Serv
	ds.cont.Invoke(func(s config.Serv) {
		configSvc = s
	})

	cfg := handler.NewConfigController(ds.logger, configSvc)
	configRoutes := Routes{
		Route{
			"UpdateConfigProperties",
			strings.ToUpper("Put"),
			"/engine",
			cfg.Update,
		},
		Route{
			"GetEngineProperties",
			strings.ToUpper("Get"),
			"/engine",
			cfg.Get,
		},
	}
	return configRoutes
}

func (ds *dserver) checkRoutes() Routes {
	var checkSvc check.Serv
	ds.cont.Invoke(func(s check.Serv) {
		checkSvc = s
	})
	chk := handler.NewCheckController(ds.logger, checkSvc)
	checkRoutes := Routes{
		Route{
			"GetCheck",
			strings.ToUpper("Get"),
			"/check/{TeamID}/{RoundID}/{ServiceID}",
			chk.GetAllByTeamRoundID,
		},
		Route{
			"GetChecks",
			strings.ToUpper("Get"),
			"/check/{TeamID}/{RoundID}",
			chk.GetByTeamRoundServiceID,
		},
	}
	return checkRoutes
}

func (ds *dserver) hostRoutes() Routes {

	var hostSvc host.Serv
	ds.cont.Invoke(func(s host.Serv) {
		hostSvc = s
	})
	hst := handler.NewHostController(ds.logger, hostSvc)

	hostRoutes := Routes{
		Route{
			"AddHost",
			strings.ToUpper("Post"),
			"/host",
			hst.Store,
		},

		Route{
			"DeleteHost",
			strings.ToUpper("Delete"),
			"/host/{HostID}",
			hst.Delete,
		},

		Route{
			"GetHost",
			strings.ToUpper("Get"),
			"/host/{HostID}",
			hst.GetByID,
		},

		Route{
			"GetHosts",
			strings.ToUpper("Get"),
			"/host",
			hst.GetAll,
		},

		Route{
			"UpdateHost",
			strings.ToUpper("Put"),
			"/host/{HostID}",
			hst.Update,
		},
	}

	return hostRoutes
}

func (ds *dserver) hostGroupRoutes() Routes {

	var hostGroupSvc host_group.Serv
	ds.cont.Invoke(func(s host_group.Serv) {
		hostGroupSvc = s
	})
	hstgrp := handler.NewHostGroupController(ds.logger, hostGroupSvc)
	hostGroupRoutes := Routes{
		Route{
			"AddHostGroup",
			strings.ToUpper("Post"),
			"/host_group",
			hstgrp.Store,
		},

		Route{
			"DeleteHostGroup",
			strings.ToUpper("Delete"),
			"/host_group/{host_groupID}",
			hstgrp.Delete,
		},

		Route{
			"GetHostGroup",
			strings.ToUpper("Get"),
			"/host_group/{host_groupID}",
			hstgrp.GetByID,
		},

		Route{
			"GetHostGroups",
			strings.ToUpper("Get"),
			"/host_group",
			hstgrp.GetAll,
		},

		Route{
			"UpdateHostGroup",
			strings.ToUpper("Put"),
			"/host_group/{host_groupID}",
			hstgrp.Update,
		},
	}

	return hostGroupRoutes
}

func (ds *dserver) propertyRoutes() Routes {

	var propertySvc property.Serv
	ds.cont.Invoke(func(s property.Serv) {
		propertySvc = s
	})

	prop := handler.NewPropertyController(ds.logger, propertySvc)

	propertyRoutes := Routes{
		Route{
			"AddProprty",
			strings.ToUpper("Post"),
			"/property",
			prop.Store,
		},

		Route{
			"DeleteProperty",
			strings.ToUpper("Delete"),
			"/property/{propertyID}",
			prop.Delete,
		},

		Route{
			"GetProperties",
			strings.ToUpper("Get"),
			"/property",
			prop.GetAll,
		},

		Route{
			"GetProperty",
			strings.ToUpper("Get"),
			"/property/{propertyID}",
			prop.Get,
		},

		Route{
			"UpdateProperty",
			strings.ToUpper("Put"),
			"/property/{propertyID}",
			prop.Update,
		},
	}

	return propertyRoutes
}

func (ds *dserver) roundRoutes() Routes {

	var roundSvc round.Serv
	ds.cont.Invoke(func(s round.Serv) {
		roundSvc = s
	})

	rnd := handler.NewRoundController(ds.logger, roundSvc)

	roundRoutes := Routes{
		Route{
			"GetLastRound",
			strings.ToUpper("Get"),
			"/round",
			rnd.GetLastRound,
		},
	}
	return roundRoutes
}

func (ds *dserver) scoreRoutes() Routes {
	s := handler.NewScoreController(ds.logger)
	scoreRoutes := Routes{
		Route{
			"GetScore",
			strings.ToUpper("Get"),
			"/score/{teamID}",
			s.GetScore,
		},

		Route{
			"GetScores",
			strings.ToUpper("Get"),
			"/score",
			s.GetScores,
		},
	}
	return scoreRoutes
}

func (ds *dserver) serviceRoutes() Routes {

	var serviceSvc service.Serv
	ds.cont.Invoke(func(s service.Serv) {
		serviceSvc = s
	})

	srv := handler.NewServiceController(ds.logger, serviceSvc)

	serviceRoutes := Routes{
		Route{
			"AddService",
			strings.ToUpper("Post"),
			"/service",
			srv.Store,
		},

		Route{
			"DeleteService",
			strings.ToUpper("Delete"),
			"/service/{ServiceID}",
			srv.Delete,
		},

		Route{
			"GetService",
			strings.ToUpper("Get"),
			"/service/{ServiceID}",
			srv.GetByID,
		},

		Route{
			"GetServices",
			strings.ToUpper("Get"),
			"/service",
			srv.GetAll,
		},

		Route{
			"UpdateService",
			strings.ToUpper("Put"),
			"/service/{ServiceID}",
			srv.Update,
		},
	}
	return serviceRoutes
}

func (ds *dserver) serviceGroupRoutes() Routes {

	var serviceGroupSvc service_group.Serv
	ds.cont.Invoke(func(s service_group.Serv) {
		serviceGroupSvc = s
	})

	servg := handler.NewServiceGroupController(ds.logger, serviceGroupSvc)

	serviceGroupRoutes := Routes{
		Route{
			"AddServiceGroup",
			strings.ToUpper("Post"),
			"/service_group",
			servg.Store,
		},

		Route{
			"DeleteServiceGroup",
			strings.ToUpper("Delete"),
			"/service_group/{serviceGroupID}",
			servg.Delete,
		},

		Route{
			"GetServiceGroup",
			strings.ToUpper("Get"),
			"/service_group/{serviceGroupID}",
			servg.GetByID,
		},

		Route{
			"GetServiceGroups",
			strings.ToUpper("Get"),
			"/service_group",
			servg.GetAll,
		},

		Route{
			"UpdateServiceGroup",
			strings.ToUpper("Put"),
			"/service_group/{serviceGroupID}",
			servg.Update,
		},
	}

	return serviceGroupRoutes
}
