
dev:
	goreman -exit-on-error -logtime -set-ports=false start

seed:
	go run main.go seed --config ./configs/dev-config.yml

migrate:
	go run main.go migrate --config ./configs/dev-config.yml

worker-dev:
	go run main.go worker --config ./configs/dev-config.yml

master-dev:
	go run main.go master --config ./configs/dev-config.yml

test:
	go test -race -covermode=atomic -coverprofile=coverage.out ./pkg/...