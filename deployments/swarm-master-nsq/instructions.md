Make sure cluster clock in synchronized

1) Build scoretrak image. While in Scoretrak directory, run ``docker image build -t scoretrak . -f deployments/master/Dockerfile``
2) Deploy stack. ``docker stack deploy --compose-file deployments/swarm-master-nsq/docker-compose.yml swarm-master-nsq``