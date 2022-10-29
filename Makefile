

rm-cockroach-data:
	rm -rf cockroach-data

rm-nsqd-data:
	rm *.dat

rm-tmp-data: rm-cockroach-data rm-nsqd-data

build-scoretrak:
	go build -o scoretrak

create-scoretrak-db:
	cockroach sql --execute="CREATE DATABASE IF NOT EXISTS scoretrak; " --insecure --host localhost --port 26257