//go:build ignore

package main

import (
	"ariga.io/ogent"
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	//_ "github.com/ScoreTrak/ScoreTrak/pkg/entities/runtime"
	"github.com/ogen-go/ogen"
	"log"
	"net/http"
	"strconv"
)

func main() {
	spec := new(ogen.Spec)
	oas, err := entoas.NewExtension(
		entoas.Spec(spec),
		entoas.MinItemsPerPage(10),
		//entoas.SimpleModels(),
		entoas.Mutations(func(graph *gen.Graph, spec *ogen.Spec) error {
			spec.SetInfo(ogen.NewInfo().
				SetTitle("ScoreTrak API").
				SetVersion("3.0.3").
				SetDescription("ScoreTrak API"))
			return nil
		}),
		entoas.Mutations(func(graph *gen.Graph, spec *ogen.Spec) error {
			spec.AddPathItem("/rounds/latest", ogen.NewPathItem().
				SetGet(ogen.NewOperation().
					SetSummary("Gets the latest round").
					SetDescription("Get latest completed round").
					AddTags("Round").
					SetOperationID("readRoundLatest").
					AddResponse(
						strconv.Itoa(http.StatusOK),
						ogen.NewResponse().
							SetDescription("Retrieved latest round").
							SetJSONContent(spec.RefSchema("RoundRead").Schema),
					).
					AddResponse(
						strconv.Itoa(http.StatusNoContent),
						ogen.NewResponse().SetDescription("No rounds"),
					).
					AddNamedResponses(
						spec.RefResponse(strconv.Itoa(http.StatusConflict)),
						spec.RefResponse(strconv.Itoa(http.StatusNotFound)),
						spec.RefResponse(strconv.Itoa(http.StatusInternalServerError)),
					),
				),
			)

			spec.AddPathItem("/rounds/latest/checks", ogen.NewPathItem().
				SetGet(ogen.NewOperation().
					SetSummary("List attached checks").
					SetDescription("List attached checks.").
					AddTags("Round").
					SetOperationID("listRoundChecksLatest").
					AddResponse(
						strconv.Itoa(http.StatusOK),
						ogen.NewResponse().
							SetDescription("result checks list").
							SetJSONContent(spec.RefSchema("Round_ChecksList").Schema.AsArray()),
					).
					AddNamedResponses(
						spec.RefResponse(strconv.Itoa(http.StatusConflict)),
						spec.RefResponse(strconv.Itoa(http.StatusNotFound)),
						spec.RefResponse(strconv.Itoa(http.StatusInternalServerError)),
					),
				),
			)

			return nil
		}),
	)
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}

	oget, err := ogent.NewExtension(spec)
	if err != nil {
		log.Fatalf("creating ogent extension: %v", err)
	}

	err = entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			//gen.FeaturePrivacy,
			//gen.FeatureEntQL,
			gen.FeatureUpsert,
			//gen.FeatureSnapshot,
		},
	}, entc.Extensions(oget, oas))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
