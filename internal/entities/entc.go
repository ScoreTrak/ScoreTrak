//go:build ignore

package main

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	_ "github.com/ScoreTrak/ScoreTrak/internal/entities/runtime"
	"github.com/ogen-go/ogen"
	"github.com/thisisibrahimd/ogent"
	"log"
)

func main() {
	spec := ogen.NewSpec().
		SetOpenAPI("3.0.3").
		SetInfo(ogen.NewInfo().
			SetTitle("ScoreTrak API").
			SetVersion("0.1.0").
			SetDescription("ScoreTrak API"),
		).
		AddParameter("competition",
			ogen.NewParameter().
				InHeader().
				SetName("X-Scoretrak-Competition-ID").
				SetSchema(ogen.String()).
				SetDescription("Competition ID"),
		).
		AddParameter("team",
			ogen.NewParameter().
				InHeader().
				SetName("X-Scoretrak-Team-ID").
				SetSchema(ogen.String()).
				SetDescription("Team ID"),
		).
		AddParameter("ory",
			ogen.NewParameter().
				InCookie().
				SetName("ory_kratos_session").
				SetSchema(ogen.String()).
				SetDescription("Ory Kratos Auth Token"),
		)
	oas, err := entoas.NewExtension(
		entoas.Spec(spec),
		entoas.MinItemsPerPage(10),
		entoas.MaxItemsPerPage(100),
		entoas.Mutations(),
	)
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}

	ogent, err := ogent.NewExtension(spec)
	if err != nil {
		log.Fatalf("creating ogent extension: %v", err)
	}

	err = entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			//gen.FeaturePrivacy,
			//gen.FeatureEntQL,
			//gen.FeatureUpsert,
			//gen.FeatureSnapshot,
		},
	}, entc.Extensions(ogent, oas))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
