# Install goreman (or any other procfile runner) to run the procfile. https://github.com/mattn/goreman

# Monitoring
# jaeger: jaeger-all-in-one --log-level=error --collector.otlp.enabled

# Queue
## NSQ
nsqd: nsqd
nsqadmin: nsqadmin --nsqd-http-address=127.0.0.1:4151
## RabbitMQ
## Kafka

# Database
## SQLite
# sqlite:
## MySQL
# mysql:
## PostgreSQL
# postgresql:
cockroachdb: cockroach start-single-node --insecure --listen-addr=localhost:26257

# Proxy
envoy: envoy -c ./configs/envoy/config.yaml

# ScoreTrak
server: go run main.go master --config ./configs/dev-config.yml
#worker: go run main.go worker --config ./configs/dev-config.yml
