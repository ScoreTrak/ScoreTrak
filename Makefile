

rm-cockroach-data:
	rm -rf cockroach-data

rm-nsqd-data:
	rm *.dat

rm-tmp-data: rm-cockroach-data rm-nsqd-data

build-scoretrak:
	go build -o scoretrak
