container_engine := "docker"
entities_dir := "./internal/entities"
schema_dir := entities_dir + "schema"
migration_dir := entities_dir + "migrate/migrations"

# Tool to automatically bundle and gen api stubs and db schemas while coding
reflex:
  reflex -s -g reflex.conf -- reflex -c reflex.conf

# DevContainer
dc +ARGUMENTS="--help":
  devcontainer --workspace-folder=. {{ARGUMENTS}}

dcexec +CMD="--help":
  devcontainer --workspace-folder=. exec {{CMD}}

# Testing
test-docker:
  {{container_engine}} build --target test .

test-local:
  sudo ginkgo -r --procs=2 --compilers=4 --randomize-all --randomize-suites --fail-on-pending --keep-going --cover --coverprofile=cover.profile --race --trace --json-report=report.json --timeout=10m --poll-progress-after=120s --poll-progress-interval=30s

test-local-simple:
  sudo ginkgo -r -trace

# Linting
lint:
  {{container_engine}} build --target lint .

# Run
run-docker-compose:
  {{container_engine}} compose build
  {{container_engine}} compose up -d

# OpenAPI spec and stub related steps
bundle-spec:
  npx @redocly/cli bundle api/scoretrak/scoretrak.yaml -o api/scoretrak.bundled.yaml

unbundle-spec:
  npx @redocly/cli split api/scoretrak.bundled.yaml --outDir api/scoretrak-split

gen-api-stub:
  go generate internal/api-stub/generate.go

# Start server
start-server:
  go run main.go serve -c ./config/dev-config.yaml

start-flagbearer:
  go run main.go flagbearer -c ./config/dev-config.yaml

start-scorer:
  go run main.go scorer -c ./config/dev-config.yaml

create-new-entity +ENTITIES:
  go run -mod=mod entgo.io/ent/cmd/ent new --target {{schema_dir}} {{ENTITIES}}

gen-ent:
  go generate {{entities_dir}}

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
