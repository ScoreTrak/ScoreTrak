//go:build ignore

package main

import (
	"ariga.io/ogent"
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	//_ "github.com/ScoreTrak/ScoreTrak/internal/entities/runtime"
	"github.com/ScoreTrak/ScoreTrak/pkg/version"
	"github.com/ogen-go/ogen"
	"log"
)

func main() {
	spec := ogen.NewSpec().
		SetInfo(ogen.NewInfo().
			SetTitle("ScoreTrak").
			SetVersion(version.Version).
			SetDescription("ScoreTrak API"),
		).
		AddParameter("comp",
			ogen.NewParameter().
				InHeader().
				SetName("X-Competition-ID"),
		).
		AddParameter("ory",
			ogen.NewParameter().
				InCookie().
				SetName("ory_kratos_session"),
		)
	oas, err := entoas.NewExtension(entoas.Spec(spec))
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}

	ogent, err := ogent.NewExtension(spec)
	if err != nil {
		log.Fatalf("creating ogent extension: %v", err)
	}

	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ogent, oas))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
