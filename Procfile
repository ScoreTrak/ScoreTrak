# Install goreman (or any other procfile runner) to run the procfile. https://github.com/mattn/goreman

jaeger: jaeger-all-in-one --log-level=error --collector.otlp.enabled

# Queue
nsqd: nsqd
nsqadmin: nsqadmin --nsqd-http-address=127.0.0.1:4150

# Database
cockroachdb: cockroach start-single-node --insecure --listen-addr=localhost:26257

#server: go run main.go master --config ./configs/dev-config.yaml
#worker: go run main.go worker --config ./configs/dev-config.yaml
#envoy: envoy --config ./configs/envoy/config.yaml
grpcui: while ! grpcui -plaintext localhost:33333 ; do sleep 1 ; done ;