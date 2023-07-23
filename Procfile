kratos: kratos serve -c ./configs/ory/kratos/kratos.yml --dev --watch-courier --sqa-opt-out
import-users: sleep 7 && kratos import identities -e http://localhost:4434 configs/ory/kratos/users.json
keto: keto serve -c ./configs/ory/keto/keto.yml --sqa-opt-out
oathkeeper: oathkeeper serve -c ./configs/ory/oathkeeper/oathkeeper.yml --sqa-opt-out
nats: nats-server -js -DV
web: npm --prefix web run dev
node-01: go run main.go serve -c ./configs/multi/node-01.config.yml
#reflex-node-01: reflex -R '^(web|images|deploy|.github|docs)/' -R 'configs/ory/' -r '(pkg|cmd)/\.go$' -r 'configs/multi/node-01.config.yaml' -- go run main.go serve -c ./configs/multi/node-01.config.yml
#node-02: go run main.go serve -c ./configs/multi/node-02.config.yml
#node-03: go run main.go serve -c ./configs/multi/node-03.config.yml
