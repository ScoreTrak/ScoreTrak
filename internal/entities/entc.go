//go:build ignore

package main

import (
	"ariga.io/ogent"
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ScoreTrak/ScoreTrak/pkg/version"
	"github.com/ogen-go/ogen"
	"log"
)

func main() {
	info := ogen.NewInfo().SetTitle("ScoreTrak").SetVersion(version.Version).SetDescription("ScoreTrak API")
	spec := ogen.NewSpec().SetInfo(info)
	competitionHeader := ogen.NewParameter().InHeader().SetName("X-Competition-Id")
	oryKratosCookie := ogen.NewParameter().SetName("ory_kratos_session").InCookie().SetSchema(ogen.String())
	spec.AddParameter("ory", oryKratosCookie)
	spec.AddParameter("comp", competitionHeader)
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
