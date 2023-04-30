entities_dir := "./internal/entities/"
schema_dir := entities_dir + "schema"
migration_dir := entities_dir + "migrate/migrations"
openapi_json_path := entities_dir + "openapi.json"

start-server:
  go run main.go serve -c ./configs/dev-config.yml

gen-ent:
  go generate ./internal/entities

gen-entviz:
  go run -mod=mod ariga.io/entviz {{schema_dir}}

create-new-entity +ENTITIES:
  go run -mod=mod entgo.io/ent/cmd/ent new --target {{schema_dir}} {{ENTITIES}}

gen-migration MIGRATION_NAME:
  atlas migrate diff {{MIGRATION_NAME}} \
  --dir "file://{{migration_dir}}" \
  --to "ent://{{schema_dir}}" \
  --dev-url "sqlite://file?mode=memory&_fk=1"

gen-js-client:
  swagger-codegen generate -i {{openapi_json_path}} -o ./web/src/lib/scoretrak-js-client -l javascript

gen-ts-client:
  swagger-codegen generate -i {{openapi_json_path}} -o ./web/src/lib/scoretrak-ts-client -l typescript-axios

gen-ts-fetch-client:
  swagger-codegen generate -i {{openapi_json_path}} -o ./web/src/lib/scoretrak-ts-fetch-client -l typescript-fetch

gen-react-queries:
  rapini react-query v4 -n "scoretrak" -p {{openapi_json_path}} -o ./web/src/lib/scoretrak-queries

build-react-query-library:
  ls

gen-schema: gen-ent gen-react-queries

start-kratos:
  kratos serve -c ./configs/ory/kratos/kratos.yml --dev --watch-courier

