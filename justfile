entities_dir := "./pkg/entities/"
schema_dir := entities_dir + "schema"
migration_dir := entities_dir + "migrate/migrations"

start-server:
  go run main.go serve -c ./configs/dev-config.yml

create-new-entity +ENTITIES:
  go run -mod=mod entgo.io/ent/cmd/ent new --target {{schema_dir}} {{ENTITIES}}

gen-ent:
  go generate {{entities_dir}}
  cat {{entities_dir}}/openapi.json | yq -P > openapi.yaml
  rm {{entities_dir}}/openapi.json

gen-js-schema:
  npm --prefix web run gen-schema
  npm --prefix web/src/lib/scoretrak-queries run build-esm

gen-entviz:
  go run -mod=mod ariga.io/entviz {{schema_dir}}

gen-migration MIGRATION_NAME:
  atlas migrate diff {{MIGRATION_NAME}} \
  --dir "file://{{migration_dir}}" \
  --to "ent://{{schema_dir}}" \
  --dev-url "sqlite://file?mode=memory&_fk=1"

start-kratos:
  kratos serve -c ./configs/ory/kratos/kratos.yml --dev --watch-courier --sqa-opt-out

import-kratos-identities:
  kratos import identities -e http://localhost:4434 configs/ory/kratos/users.json

start-keto:
  keto serve -c ./configs/ory/keto/keto.yml --sqa-opt-out

start-nats-js:
  nats-server -js -DV
