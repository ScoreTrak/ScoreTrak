set dotenv-load

start-ory-tunnel:
  ory tunnel --dev http://localhost:4000

start-server:
  go run main.go master

gen-ent:
  go generate ./internal/entities

gen-entviz: gen-ent
  go run -mod=mod ariga.io/entviz ./internal/entities/schema

create-new-entity +ENTITIES:
  go run -mod=mod entgo.io/ent/cmd/ent new --target ./internal/entities/schema {{ENTITIES}}

gen-migration MIGRATION_NAME:
  atlas migrate diff {{MIGRATION_NAME}} \
  --dir "file://ent/migrate/migrations" \
  --to "ent://internal/entities/schema" \
  --dev-url "sqlite://file?mode=memory&_fk=1"

gen-js-client:
  swagger-codegen generate -i ./internal/entities/openapi.json -o ./web/src/lib/scoretrak-client -l javascript

gen-react-queries:
  rapini react-query v4 -n "scoretrak" -p ./internal/entities/openapi.json -o ./web/src/lib/scoretrak-queries

gen-schema: gen-ent gen-js-client gen-react-queries
