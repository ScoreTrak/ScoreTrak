package main

import (
	"encoding/json"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/danielgtaylor/huma/schema"
	"reflect"
)

func main() {
	s, err := schema.Generate(reflect.TypeOf(config.StaticConfig{}))
	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
