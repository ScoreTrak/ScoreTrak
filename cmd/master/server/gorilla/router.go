package gorilla

import (
	"encoding/json"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/api/http/handler"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/competition"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/di/repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"time"
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
	routes = append(routes, ds.staticConfigRoutes()...)
	routes = append(routes, ds.teamRoutes()...)
	routes = append(routes, ds.checkRoutes()...)
	routes = append(routes, ds.hostRoutes()...)
	routes = append(routes, ds.hostGroupRoutes()...)
	routes = append(routes, ds.propertyRoutes()...)
	routes = append(routes, ds.roundRoutes()...)
	routes = append(routes, ds.serviceRoutes()...)
	routes = append(routes, ds.serviceGroupRoutes()...)
	routes = append(routes, ds.reportRoutes()...)
	routes = append(routes, ds.competitionRoutes()...)
	routes = append(routes, Route{
		"Time",
		strings.ToUpper("Get"),
		"/time",
		func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(time.Now())
		},
	})

	for _, route := range routes {
		var hdler http.Handler
		hdler = route.HandlerFunc
		hdler = Logger(hdler, route.Name) //Default Logger

		ds.router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(hdler)
	}
	ds.router.Use(JsonHeader)
	ds.router.Use(TokenVerify)
}

func TokenVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("x-access-token")
		json.NewEncoder(w).Encode(r)
		header = strings.TrimSpace(header)

		if header != config.GetToken() {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("The request is either missing or has an incorrect auth token")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func JsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		next.ServeHTTP(w, r)
	})
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func (ds *dserver) reportRoutes() Routes {
	var svc report.Serv
	err := ds.cont.Invoke(func(s report.Serv) {
		svc = s
	})
	if err != nil {
		panic(err)
	}
	return ReportRoutes(ds.logger, svc)
}

func ReportRoutes(l logger.LogInfoFormat, svc report.Serv) Routes {
	ctrl := handler.NewReportController(l, svc)
	reportRoutes := Routes{
		Route{
			"GetReport",
			strings.ToUpper("Get"),
			"/report",
			ctrl.Get,
		},
	}
	return reportRoutes
}

func (ds *dserver) configRoutes() Routes {
	var svc config.Serv
	err := ds.cont.Invoke(func(s config.Serv) {
		svc = s
	})
	if err != nil {
		panic(err)
	}
	return ConfigRoutes(ds.logger, svc)
}

func ConfigRoutes(l logger.LogInfoFormat, svc config.Serv) Routes {
	ctrl := handler.NewConfigController(l, svc)
	configRoutes := Routes{
		Route{
			"UpdateConfigProperties",
			strings.ToUpper("Patch"),
			"/config",
			ctrl.Update,
		},
		Route{
			"GetEngineProperties",
			strings.ToUpper("Get"),
			"/config",
			ctrl.Get,
		},
		Route{
			"ResetScores",
			strings.ToUpper("Delete"),
			"/config/reset_competition",
			ctrl.ResetScores,
		},
		Route{
			"DeleteCompetition",
			strings.ToUpper("Delete"),
			"/config/delete_competition",
			ctrl.DeleteCompetition,
		},
	}
	return configRoutes
}

func (ds *dserver) staticConfigRoutes() Routes {
	var svc config.StaticServ
	err := ds.cont.Invoke(func(s config.StaticServ) {
		svc = s
	})
	if err != nil {
		panic(err)
	}
	return StaticConfigRoutes(ds.logger, svc)
}

func StaticConfigRoutes(l logger.LogInfoFormat, svc config.StaticServ) Routes {
	ctrl := handler.NewStaticConfigController(l, svc)
	configRoutes := Routes{
		Route{
			"GetStaticEngineProperties",
			strings.ToUpper("Get"),
			"/static_config",
			ctrl.Get,
		},
	}
	return configRoutes
}

func (ds *dserver) checkRoutes() Routes {
	var svc check.Serv
	err := ds.cont.Invoke(func(s check.Serv) {
		svc = s
	})
	if err != nil {
		panic(err)
	}
	return CheckRoutes(ds.logger, svc)
}

func CheckRoutes(l logger.LogInfoFormat, svc check.Serv) Routes {
	ctrl := handler.NewCheckController(l, svc)
	checkRoutes := Routes{
		Route{
			"GetCheck",
			strings.ToUpper("Get"),
			"/check/round/{RoundID}/service/{ServiceID}",
			ctrl.GetByRoundServiceID,
		},
		Route{
			"GetChecksByRound",
			strings.ToUpper("Get"),
			"/check/round/{RoundID}",
			ctrl.GetAllByRoundID,
		},
		Route{
			"GetChecksByService",
			strings.ToUpper("Get"),
			"/check/service/{ServiceID}",
			ctrl.GetAllByServiceID,
		},
	}
	return checkRoutes
}

func (ds *dserver) hostRoutes() Routes {
	var svc host.Serv
	err := ds.cont.Invoke(func(s host.Serv) {
		svc = s
	})
	if err != nil {
		panic(err)
	}
	return HostRoutes(ds.logger, svc)
}

func HostRoutes(l logger.LogInfoFormat, svc host.Serv) Routes {
	ctrl := handler.NewHostController(l, svc)
	hostRoutes := Routes{
		Route{
			"AddHost",
			strings.ToUpper("Post"),
			"/host",
			ctrl.Store,
		},

		Route{
			"DeleteHost",
			strings.ToUpper("Delete"),
			"/host/{id}",
			ctrl.Delete,
		},

		Route{
			"GetHost",
			strings.ToUpper("Get"),
			"/host/{id}",
			ctrl.GetByID,
		},

		Route{
			"GetHosts",
			strings.ToUpper("Get"),
			"/host",
			ctrl.GetAll,
		},

		Route{
			"UpdateHost",
			strings.ToUpper("Patch"),
			"/host/{id}",
			ctrl.Update,
		},
	}
	return hostRoutes
}

func (ds *dserver) hostGroupRoutes() Routes {
	var svc host_group.Serv
	err := ds.cont.Invoke(func(s host_group.Serv) {
		svc = s
	})
	if err != nil {
		panic(err)
	}
	return HostGroupRoutes(ds.logger, svc)
}

func HostGroupRoutes(l logger.LogInfoFormat, svc host_group.Serv) Routes {
	ctrl := handler.NewHostGroupController(l, svc)
	hostGroupRoutes := Routes{
		Route{
			"AddHostGroup",
			strings.ToUpper("Post"),
			"/host_group",
			ctrl.Store,
		},

		Route{
			"DeleteHostGroup",
			strings.ToUpper("Delete"),
			"/host_group/{id}",
			ctrl.Delete,
		},

		Route{
			"GetHostGroup",
			strings.ToUpper("Get"),
			"/host_group/{id}",
			ctrl.GetByID,
		},

		Route{
			"GetHostGroups",
			strings.ToUpper("Get"),
			"/host_group",
			ctrl.GetAll,
		},

		Route{
			"UpdateHostGroup",
			strings.ToUpper("Patch"),
			"/host_group/{id}",
			ctrl.Update,
		},
	}

	return hostGroupRoutes
}

func (ds *dserver) propertyRoutes() Routes {
	var svc property.Serv
	err := ds.cont.Invoke(func(s property.Serv) {
		svc = s
	})
	if err != nil {
		panic(err)
	}
	return PropertyRoutes(ds.logger, svc)
}

func PropertyRoutes(l logger.LogInfoFormat, svc property.Serv) Routes {
	ctrl := handler.NewPropertyController(l, svc)
	propertyRoutes := Routes{
		Route{
			"GetProperties",
			strings.ToUpper("Get"),
			"/property",
			ctrl.GetAll,
		},

		Route{
			"AddProperty",
			strings.ToUpper("Post"),
			"/property",
			ctrl.Store,
		},

		Route{
			"GetPropertiesByID",
			strings.ToUpper("Get"),
			"/property/{ServiceID}",
			ctrl.GetAllByServiceID,
		},

		Route{
			"DeleteProperty",
			strings.ToUpper("Delete"),
			"/property/{ServiceID}/{Key}",
			ctrl.Delete,
		},

		Route{
			"GetProperty",
			strings.ToUpper("Get"),
			"/property/{ServiceID}/{Key}",
			ctrl.GetByServiceIDKey,
		},

		Route{
			"UpdateProperty",
			strings.ToUpper("Patch"),
			"/property/{ServiceID}/{Key}",
			ctrl.Update,
		},
	}
	return propertyRoutes
}

func (ds *dserver) roundRoutes() Routes {
	var svc round.Serv
	err := ds.cont.Invoke(func(s round.Serv) {
		svc = s
	})
	if err != nil {
		panic(err)
	}
	return RoundRoutes(ds.logger, svc)
}

func RoundRoutes(l logger.LogInfoFormat, svc round.Serv) Routes {
	ctrl := handler.NewRoundController(l, svc)
	roundRoutes := Routes{
		Route{
			"GetLastNonElapsingRound",
			strings.ToUpper("Get"),
			"/round/last_non_elapsing",
			ctrl.GetLastNonElapsingRound,
		},
		Route{
			"GetLastRound",
			strings.ToUpper("Get"),
			"/round/last",
			ctrl.GetLastRound,
		},
		Route{
			"GetRounds",
			strings.ToUpper("Get"),
			"/round",
			ctrl.GetAll,
		},
		Route{
			"GetRoundByID",
			strings.ToUpper("Get"),
			"/round/{id}",
			ctrl.GetByID,
		},
	}
	return roundRoutes
}

func (ds *dserver) competitionRoutes() Routes {
	svc := competition.NewCompetitionServ(repo.NewStore())
	return CompetitionRoutes(ds.logger, svc)
}

func CompetitionRoutes(l logger.LogInfoFormat, svc competition.Serv) Routes {
	ctrl := handler.NewCompetitionController(l, svc)
	competitionRoutes := Routes{
		Route{
			"LoadCompetition",
			strings.ToUpper("Post"),
			"/competition/upload",
			ctrl.LoadCompetition,
		},

		Route{
			"FetchEntireCompetition",
			strings.ToUpper("Get"),
			"/competition/export_all",
			ctrl.FetchEntireCompetition,
		},

		Route{
			"FetchCoreCompetition",
			strings.ToUpper("Get"),
			"/competition/export_core",
			ctrl.FetchCoreCompetition,
		},
	}
	return competitionRoutes
}

func (ds *dserver) serviceRoutes() Routes {
	var svc service.Serv
	err := ds.cont.Invoke(func(s service.Serv) {
		svc = s
	})
	if err != nil {
		panic(err)
	}
	var q queue.Queue
	err = ds.cont.Invoke(func(s queue.Queue) {
		q = s
	})
	if err != nil {
		panic(err)
	}

	return ServiceRoutes(ds.logger, svc, q, repo.NewStore())
}

func ServiceRoutes(l logger.LogInfoFormat, svc service.Serv, q queue.Queue, repoStore repo.Store) Routes {
	ctrl := handler.NewServiceController(l, svc, q, repoStore)
	serviceRoutes := Routes{
		Route{
			"AddService",
			strings.ToUpper("Post"),
			"/service",
			ctrl.Store,
		},

		Route{
			"DeleteService",
			strings.ToUpper("Delete"),
			"/service/{id}",
			ctrl.Delete,
		},

		Route{
			"GetService",
			strings.ToUpper("Get"),
			"/service/{id}",
			ctrl.GetByID,
		},

		Route{
			"GetServices",
			strings.ToUpper("Get"),
			"/service",
			ctrl.GetAll,
		},

		Route{
			"UpdateService",
			strings.ToUpper("Patch"),
			"/service/{id}",
			ctrl.Update,
		},
		Route{
			"TestService",
			strings.ToUpper("Get"),
			"/service/test/{id}",
			ctrl.TestService,
		},
	}
	return serviceRoutes
}

func (ds *dserver) serviceGroupRoutes() Routes {
	var svc service_group.Serv
	err := ds.cont.Invoke(func(s service_group.Serv) {
		svc = s
	})
	if err != nil {
		panic(err)
	}
	var plt platform.Platform
	err = ds.cont.Invoke(func(p platform.Platform) {
		plt = p
	})
	if err != nil {
		panic(err)
	}
	var q queue.Queue
	err = ds.cont.Invoke(func(s queue.Queue) {
		q = s
	})
	if err != nil {
		panic(err)
	}
	return ServiceGroupRoutes(ds.logger, svc, plt, q)
}

func ServiceGroupRoutes(l logger.LogInfoFormat, svc service_group.Serv, platform platform.Platform, q queue.Queue) Routes {
	ctrl := handler.NewServiceGroupController(l, svc, platform, q)
	serviceGroupRoutes := Routes{
		Route{
			"AddServiceGroup",
			strings.ToUpper("Post"),
			"/service_group",
			ctrl.Store,
		},

		Route{
			"DeleteServiceGroup",
			strings.ToUpper("Delete"),
			"/service_group/{id}",
			ctrl.Delete,
		},

		Route{
			"GetServiceGroup",
			strings.ToUpper("Get"),
			"/service_group/{id}",
			ctrl.GetByID,
		},

		Route{
			"GetServiceGroups",
			strings.ToUpper("Get"),
			"/service_group",
			ctrl.GetAll,
		},

		Route{
			"UpdateServiceGroup",
			strings.ToUpper("Patch"),
			"/service_group/{id}",
			ctrl.Update,
		},
		Route{
			"RedeployWorkers",
			strings.ToUpper("GET"),
			"/service_group/{id}/redeploy",
			ctrl.Redeploy,
		},
	}

	return serviceGroupRoutes
}

func (ds *dserver) teamRoutes() Routes {
	var svc team.Serv
	err := ds.cont.Invoke(func(s team.Serv) {
		svc = s
	})
	if err != nil {
		panic(err)
	}
	return TeamRoutes(ds.logger, svc)
}

func TeamRoutes(l logger.LogInfoFormat, svc team.Serv) Routes {
	ctrl := handler.NewTeamController(l, svc)
	teamRoutes := Routes{
		Route{
			"AddTeam",
			strings.ToUpper("Post"),
			"/team",
			ctrl.Store,
		},

		Route{
			"DeleteByName",
			strings.ToUpper("Delete"),
			"/team/{id}",
			ctrl.Delete,
		},

		Route{
			"GetByName",
			strings.ToUpper("Get"),
			"/team/{id}",
			ctrl.GetByID,
		},

		Route{
			"GetTeams",
			strings.ToUpper("Get"),
			"/team",
			ctrl.GetAll,
		},

		Route{
			"UpdateTeam",
			strings.ToUpper("Patch"),
			"/team/{id}",
			ctrl.Update,
		},
	}

	return teamRoutes
}
