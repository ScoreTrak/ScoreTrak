package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"reflect"

// 	"github.com/ScoreTrak/ScoreTrak/pkg/config"
// 	"github.com/danielgtaylor/huma/schema"
// )

// func main() {
// 	s, err := schema.Generate(reflect.TypeOf(config.Config{}))
// 	if err != nil {
// 		panic(err)
// 	}

// 	b, err := json.MarshalIndent(s, "", "  ")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(string(b))
// }
