#!/bin/bash
docker network create -d overlay --attachable swarm-nsq_cockroachdb
docker network create -d overlay --attachable swarm-nsq_nsq

docker service create \
--replicas 1 \
--name swarm-nsq_cockroachdb-1 \
--hostname cockroachdb-1 \
--network name=swarm-nsq_cockroachdb,alias=cockroach,alias=cockroachdb-1 \
--mount type=volume,source=cockroachdb-1,target=/cockroach/cockroach-data,volume-driver=local \
--stop-grace-period 60s \
--publish 8080:8080 \
cockroachdb/cockroach:v20.1.2 start \
--join=cockroachdb-1:26257,cockroachdb-2:26257,cockroachdb-3:26257 \
--cache=.25 \
--max-sql-memory=.25 \
--logtostderr \
--insecure

docker service create \
--replicas 1 \
--name swarm-nsq_cockroachdb-2 \
--hostname cockroachdb-2 \
--network name=swarm-nsq_cockroachdb,alias=cockroach,alias=cockroachdb-2 \
--mount type=volume,source=cockroachdb-2,target=/cockroach/cockroach-data,volume-driver=local \
--stop-grace-period 60s \
cockroachdb/cockroach:v20.1.2 start \
--join=cockroachdb-1:26257,cockroachdb-2:26257,cockroachdb-3:26257 \
--cache=.25 \
--max-sql-memory=.25 \
--logtostderr \
--insecure

docker service create \
--replicas 1 \
--name swarm-nsq_cockroachdb-3 \
--hostname cockroachdb-3 \
--network name=swarm-nsq_cockroachdb,alias=cockroach,alias=cockroachdb-3 \
--mount type=volume,source=cockroachdb-3,target=/cockroach/cockroach-data,volume-driver=local \
--stop-grace-period 60s \
cockroachdb/cockroach:v20.1.2 start \
--join=cockroachdb-1:26257,cockroachdb-2:26257,cockroachdb-3:26257 \
--cache=.25 \
--max-sql-memory=.25 \
--logtostderr \
--insecure

sleep 5s

docker run -it --rm --network=swarm-nsq_cockroachdb cockroachdb/cockroach:v20.1.2 init --host=cockroachdb-1 --insecure

sleep 2s

docker run -it --rm --network=swarm-nsq_cockroachdb cockroachdb/cockroach:v20.1.2 sql --execute="CREATE DATABASE IF NOT EXISTS scoretrak;" --insecure --host cockroachdb-1

curl -L https://downloads.portainer.io/portainer-agent-stack.yml -o portainer-agent-stack.yml
docker stack deploy --compose-file=portainer-agent-stack.yml portainer

docker service create \
--mode global \
--name swarm-nsq_nsqlookupd \
--constraint node.role==manager \
--hostname nsqlookupd \
--network name=swarm-nsq_nsq,alias=nsqlookupd \
nsqio/nsq:latest /nsqlookupd


docker service create \
--replicas 3 \
--name swarm-nsq_nsqd \
--constraint node.role==manager \
--hostname nsqd \
--network name=swarm-nsq_nsq,alias=nsqd \
nsqio/nsq:latest sh -c '/nsqd --broadcast-address=$(hostname -i) --lookupd-tcp-address=nsqlookupd:4160'


docker service create \
--replicas 1 \
--publish 4171:4171 \
--name swarm-nsq_nsqadmin \
--constraint node.role==manager \
--hostname nsqadmin \
--network name=swarm-nsq_nsq,alias=nsqadmin \
nsqio/nsq:latest /nsqadmin --lookupd-http-address=nsqlookupd:4161

docker config create scoretrak-config - < deployments/swarm-nsq/config.yml

docker service create \
--replicas 3 \
--publish 33333:33333 \
--name swarm-nsq_scoretrak \
--constraint node.role==manager \
--hostname scoretrak \
--mount type=bind,src=/var/run/docker.sock,dst=/var/run/docker.sock \
--network swarm-nsq_nsq \
--network swarm-nsq_cockroachdb \
--replicas-max-per-node 1 \
--config src=scoretrak-config,target=/config.yml \
l1ghtman/scoretrak:latest ./master -config /config.yml


